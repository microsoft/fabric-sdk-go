// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// SPDX-License-Identifier: MIT

package oploc

// This file contains handwritten additions to the generated code
// custom poller handler for LRO operations: Operation Location

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"

	"github.com/microsoft/fabric-sdk-go/internal/pollers"
)

// PollerHandler is an LRO poller that uses the Location pattern.
type PollerHandler[T any] struct {
	pl   runtime.Pipeline
	resp *http.Response

	OpLocURL string `json:"oplocURL"`
	OrigURL  string `json:"origURL"`
	Method   string `json:"method"`

	FinalState runtime.FinalStateVia `json:"finalState"`
	CurState   string                `json:"state"`
}

// NewPollerHandler creates a new PollerHandler from the provided initial response.
// Pass nil for response to create an empty Poller for rehydration.
func NewPollerHandler[T any](pl runtime.Pipeline, resp *http.Response, finalState runtime.FinalStateVia) (*PollerHandler[T], error) {
	if resp == nil {
		return &PollerHandler[T]{pl: pl}, nil
	}

	locURL := resp.Header.Get("Location")
	if locURL == "" {
		return nil, errors.New("response is missing Location header")
	}
	if !pollers.IsValidURL(locURL) {
		return nil, fmt.Errorf("invalid polling URL %s", locURL)
	}

	curState := pollers.StatusNotStarted
	status, err := pollers.GetStatus(resp)
	if err != nil && !errors.Is(err, pollers.ErrNoBody) {
		return nil, err
	}
	if status != "" {
		curState = status
	}

	return &PollerHandler[T]{
		pl:         pl,
		resp:       resp,
		OpLocURL:   locURL,
		OrigURL:    resp.Request.URL.String(),
		Method:     resp.Request.Method,
		FinalState: finalState,
		CurState:   curState,
	}, nil
}

// Done returns true if the LRO is in a terminal state.
// InProgress is considered a terminal state for the purposes of this method.
func (p *PollerHandler[T]) Done() bool {
	return pollers.IsTerminalState(p.CurState) || pollers.InProgress(p.CurState)
}

// Poll retrieves the current state of the LRO.
func (p *PollerHandler[T]) Poll(ctx context.Context) (*http.Response, error) {
	err := pollers.PollHelper(ctx, p.OpLocURL, p.pl, func(resp *http.Response) (string, error) {
		if !pollers.StatusCodeValid(resp) {
			p.resp = resp
			return "", runtime.NewResponseError(resp)
		}
		state, err := pollers.GetStatus(resp)
		if err != nil {
			return "", err
		} else if state == "" {
			return "", errors.New("the response did not contain a status")
		}
		p.resp = resp
		p.CurState = state
		return p.CurState, nil
	})
	if err != nil {
		return nil, err
	}
	return p.resp, nil
}

// Result retrieves the final result of the LRO.
func (p *PollerHandler[T]) Result(ctx context.Context, out *T) error {
	if p.resp.StatusCode == http.StatusNoContent {
		return nil
	} else if !pollers.StatusCodeValid(p.resp) || pollers.Failed(p.CurState) {
		return runtime.NewResponseError(p.resp)
	}

	locURL := p.resp.Header.Get("Location")
	if (pollers.Succeeded(p.CurState) || pollers.InProgress(p.CurState)) && locURL != "" {
		if !pollers.IsValidURL(locURL) {
			return fmt.Errorf("invalid result URL %s", locURL)
		}
		req, err := runtime.NewRequest(ctx, http.MethodGet, locURL)
		if err != nil {
			return err
		}
		resp, err := p.pl.Do(req)
		if err != nil {
			return err
		}
		p.resp = resp
	}
	return pollers.ResultHelper(p.resp, pollers.Failed(p.CurState), out)
}
