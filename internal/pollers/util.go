// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// SPDX-License-Identifier: MIT

package pollers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// the well-known set of LRO status/provisioning state values.
const (
	StatusSucceeded  = "Succeeded"
	StatusCanceled   = "Canceled"
	StatusFailed     = "Failed"
	StatusInProgress = "InProgress"
)

// these are non-conformant states that we've seen in the wild.
// we support them for back-compat.
const (
	StatusCancelled  = "Cancelled"
	StatusCompleted  = "Completed"
	StatusNotStarted = "NotStarted"
	StatusRunning    = "Running"
	StatusUndefined  = "Undefined"
	StatusDeduped    = "Deduped"
)

// IsTerminalState returns true if the LRO's state is terminal.
func IsTerminalState(s string) bool {
	return Failed(s) || Succeeded(s)
}

// Failed returns true if the LRO's state is terminal failure.
func Failed(s string) bool {
	return strings.EqualFold(s, StatusFailed) || strings.EqualFold(s, StatusCanceled) || strings.EqualFold(s, StatusCancelled) || strings.EqualFold(s, StatusUndefined)
}

// Succeeded returns true if the LRO's state is terminal success.
func Succeeded(s string) bool {
	return strings.EqualFold(s, StatusSucceeded) || strings.EqualFold(s, StatusCompleted) || strings.EqualFold(s, StatusDeduped)
}

// InProgress returns true if the LRO's state is in progress.
func InProgress(s string) bool {
	return strings.EqualFold(s, StatusInProgress) || strings.EqualFold(s, StatusRunning) || strings.EqualFold(s, StatusNotStarted)
}

// returns true if the LRO response contains a valid HTTP status code
func StatusCodeValid(resp *http.Response) bool {
	return runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusCreated, http.StatusNoContent)
}

// IsValidURL verifies that the URL is valid and absolute.
func IsValidURL(s string) bool {
	u, err := url.Parse(s)
	return err == nil && u.IsAbs()
}

// ErrNoBody is returned if the response didn't contain a body.
var ErrNoBody = errors.New("the response did not contain a body")

// GetJSON reads the response body into a raw JSON object.
// It returns ErrNoBody if there was no content.
func GetJSON(resp *http.Response) (map[string]any, error) {
	body, err := runtime.Payload(resp)
	if err != nil {
		return nil, err
	}
	if len(body) == 0 {
		return nil, ErrNoBody
	}
	// unmarshall the body to get the value
	var jsonBody map[string]any
	if err = json.Unmarshal(body, &jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// GetStatus returns the LRO's status from the response body.
// Typically used for Azure-AsyncOperation flows.
// If there is no status in the response body the empty string is returned.
func GetStatus(resp *http.Response) (string, error) {
	jsonBody, err := GetJSON(resp)
	if err != nil {
		return "", err
	}
	return status(jsonBody), nil
}

// status returns the status from the response or the empty string.
func status(jsonBody map[string]any) string {
	rawStatus, ok := jsonBody["status"]
	if !ok {
		return ""
	}
	status, ok := rawStatus.(string)
	if !ok {
		return ""
	}
	return status
}

// GetProvisioningState returns the LRO's state from the response body.
// If there is no state in the response body the empty string is returned.
func GetProvisioningState(resp *http.Response) (string, error) {
	jsonBody, err := GetJSON(resp)
	if err != nil {
		return "", err
	}
	return provisioningState(jsonBody), nil
}

// provisioningState returns the provisioning state from the response or the empty string.
func provisioningState(jsonBody map[string]any) string {
	jsonProps, ok := jsonBody["properties"]
	if !ok {
		return ""
	}
	props, ok := jsonProps.(map[string]any)
	if !ok {
		return ""
	}
	rawPs, ok := props["provisioningState"]
	if !ok {
		return ""
	}
	ps, ok := rawPs.(string)
	if !ok {
		return ""
	}
	return ps
}

func PollHelper(ctx context.Context, endpoint string, pl runtime.Pipeline, update func(resp *http.Response) (string, error)) error {
	req, err := runtime.NewRequest(ctx, http.MethodGet, endpoint)
	if err != nil {
		return err
	}
	resp, err := pl.Do(req)
	if err != nil {
		return err
	}
	_, err = update(resp)
	if err != nil {
		return err
	}

	return nil
}

func ResultHelper[T any](resp *http.Response, failed bool, out *T) error {
	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	defer resp.Body.Close()
	if !StatusCodeValid(resp) || failed {
		// the LRO failed. Unmarshall the error and update state
		return runtime.NewResponseError(resp)
	}

	// success case
	payload, err := runtime.Payload(resp)
	if err != nil {
		return err
	}
	if len(payload) == 0 {
		return nil
	}

	if err = json.Unmarshal(payload, out); err != nil {
		return err
	}
	return nil
}

// GetResourceLocation returns the LRO's resourceLocation value from the response body.
// Typically used for Operation-Location flows.
// If there is no resourceLocation in the response body the empty string is returned.
func GetResourceLocation(resp *http.Response) (string, error) {
	jsonBody, err := GetJSON(resp)
	if err != nil {
		return "", err
	}
	v, ok := jsonBody["resourceLocation"]
	if !ok {
		// it might be ok if the field doesn't exist, the caller must make that determination
		return "", nil
	}
	vv, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("the resourceLocation value %v was not in string format", v)
	}
	return vv, nil
}
