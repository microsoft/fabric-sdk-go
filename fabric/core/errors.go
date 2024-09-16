// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// SPDX-License-Identifier: MIT

package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ResponseError struct {
	ErrorCode     string
	StatusCode    int
	ErrorResponse *ErrorResponse
	RawResponse   *http.Response
}

func NewResponseError(resp *http.Response) error {
	if resp.Body == nil {
		// this shouldn't happen in real-world scenarios as a
		// response with no body should set it to http.NoBody
		return nil
	}

	respBody, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return err
	}

	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(respBody, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling error: %v", err)
	}

	// handle LRO result error
	if rawMsg["error"] != nil && rawMsg["status"] != nil {
		var errorMap map[string]any
		if err := json.Unmarshal(rawMsg["error"], &errorMap); err != nil {
			return fmt.Errorf("unmarshalling error field: %v", err)
		}

		errorMap["requestId"] = resp.Header.Get("requestId")
		errorMap["moreDetails"] = []map[string]any{
			{
				"errorCode": string(rawMsg["status"]),
				"message":   "Long Running Operation: " + string(rawMsg["status"]),
			},
		}

		respBody, err = json.Marshal(errorMap)
		if err != nil {
			return fmt.Errorf("marshalling error field: %v", err)
		}
	}

	var errResp ErrorResponse
	if err := errResp.UnmarshalJSON(respBody); err != nil {
		return err
	}

	return &ResponseError{
		ErrorCode:     *errResp.ErrorCode,
		StatusCode:    resp.StatusCode,
		ErrorResponse: &errResp,
		RawResponse:   resp,
	}
}

func (e *ResponseError) Error() string {
	const separator = "--------------------------------------------------------------------------------"

	msg := &bytes.Buffer{}

	if e.RawResponse != nil {
		if e.RawResponse.Request != nil {
			fmt.Fprintf(msg, "%s %s://%s%s\n", e.RawResponse.Request.Method, e.RawResponse.Request.URL.Scheme, e.RawResponse.Request.URL.Host, e.RawResponse.Request.URL.Path)
		} else {
			fmt.Fprintln(msg, "Request information not available")
		}
		fmt.Fprintln(msg, separator)
		fmt.Fprintf(msg, "RESPONSE %d: %s\n", e.RawResponse.StatusCode, e.RawResponse.Status)
	} else {
		fmt.Fprintln(msg, "Missing RawResponse")
		fmt.Fprintln(msg, separator)
	}

	if e.ErrorResponse != nil {
		if e.ErrorResponse.ErrorCode != nil {
			fmt.Fprintf(msg, "ERROR CODE: %s\n", *e.ErrorResponse.ErrorCode)
		} else {
			fmt.Fprintln(msg, "ERROR CODE UNAVAILABLE")
		}

		errResp, err := e.ErrorResponse.MarshalJSON()
		if err != nil {
			fmt.Fprintf(msg, "Error reading ErrorResponse: %v\n", err)
		} else if len(errResp) > 0 {
			if err := json.Indent(msg, errResp, "", "  "); err != nil {
				// failed to pretty-print so just dump it verbatim
				fmt.Fprintln(msg, string(errResp))
			}
		} else {
			fmt.Fprintln(msg, "Error response contained no body")
		}

		fmt.Fprintf(msg, "\n%s", separator)
	}

	return msg.String()
}

func (e *ResponseError) Is(target error) bool {
	if e.ErrorResponse == nil {
		return false
	}

	if e.ErrorResponse.ErrorCode == nil {
		return false
	}

	return *e.ErrorResponse.ErrorCode == target.Error()
}
