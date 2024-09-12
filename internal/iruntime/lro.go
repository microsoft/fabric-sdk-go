// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// SPDX-License-Identifier: MIT

package iruntime

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// LRO is a SDK helper struct to handle Long Running Operations.
type LRO[TResult any] struct {
	poller *runtime.Poller[TResult]
	err    error
}

// NewLRO creates a new LRO.
func NewLRO[TResult any](poller *runtime.Poller[TResult], err error) LRO[TResult] {
	return LRO[TResult]{poller: poller, err: err}
}

// Wait blocks until the LRO is done.
func (lro LRO[TResult]) Sync(ctx context.Context) (TResult, error) {
	if lro.err != nil {
		return *new(TResult), lro.err
	}
	res, err := lro.poller.PollUntilDone(ctx, nil)
	return res, err
}
