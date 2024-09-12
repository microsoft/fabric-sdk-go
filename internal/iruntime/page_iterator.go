// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// SPDX-License-Identifier: MIT

package iruntime

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// PageIterator is a SDK helper struct to iterate over paginated responses.
type PageIterator[T any, TResult any] struct {
	ctx    context.Context
	pager  *runtime.Pager[T]
	mapper func(resp T) []TResult
}

// NewPageIterator creates a new List for the given pager and mapper.
func NewPageIterator[T any, TResult any](ctx context.Context, pager *runtime.Pager[T], mapper func(resp T) []TResult) PageIterator[T, TResult] {
	return PageIterator[T, TResult]{ctx: ctx, pager: pager, mapper: mapper}
}

// Returns first page only.
func (l PageIterator[T, TResult]) GetFirst() ([]TResult, error) {
	return l.get(true)
}

// Returns all pages.
func (l PageIterator[T, TResult]) Get() ([]TResult, error) {
	return l.get(false)
}

// get returns all pages if firstOnly is false, otherwise returns first page only.
func (l PageIterator[T, TResult]) get(firstOnly bool) ([]TResult, error) {
	var result []TResult
	for l.pager.More() {

		page, err := l.pager.NextPage(l.ctx)
		if err != nil {
			return nil, err
		}

		one := l.mapper(page)
		if len(one) > 0 {
			result = append(result, one...)
		}

		if firstOnly {
			break
		}
	}

	return result, nil
}
