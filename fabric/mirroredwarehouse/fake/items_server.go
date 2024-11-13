// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package fake

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"

	"github.com/microsoft/fabric-sdk-go/fabric/mirroredwarehouse"
)

// ItemsServer is a fake server for instances of the mirroredwarehouse.ItemsClient type.
type ItemsServer struct {
	// NewListMirroredWarehousesPager is the fake for method ItemsClient.NewListMirroredWarehousesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListMirroredWarehousesPager func(workspaceID string, options *mirroredwarehouse.ItemsClientListMirroredWarehousesOptions) (resp azfake.PagerResponder[mirroredwarehouse.ItemsClientListMirroredWarehousesResponse])
}

// NewItemsServerTransport creates a new instance of ItemsServerTransport with the provided implementation.
// The returned ItemsServerTransport instance is connected to an instance of mirroredwarehouse.ItemsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewItemsServerTransport(srv *ItemsServer) *ItemsServerTransport {
	return &ItemsServerTransport{
		srv:                            srv,
		newListMirroredWarehousesPager: newTracker[azfake.PagerResponder[mirroredwarehouse.ItemsClientListMirroredWarehousesResponse]](),
	}
}

// ItemsServerTransport connects instances of mirroredwarehouse.ItemsClient to instances of ItemsServer.
// Don't use this type directly, use NewItemsServerTransport instead.
type ItemsServerTransport struct {
	srv                            *ItemsServer
	newListMirroredWarehousesPager *tracker[azfake.PagerResponder[mirroredwarehouse.ItemsClientListMirroredWarehousesResponse]]
}

// Do implements the policy.Transporter interface for ItemsServerTransport.
func (i *ItemsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	parts := strings.Split(method, ".")
	method = parts[1] + "." + parts[2]
	return i.dispatchToMethodFake(req, method)
}

func (i *ItemsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if itemsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = itemsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ItemsClient.NewListMirroredWarehousesPager":
				res.resp, res.err = i.dispatchNewListMirroredWarehousesPager(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (i *ItemsServerTransport) dispatchNewListMirroredWarehousesPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListMirroredWarehousesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListMirroredWarehousesPager not implemented")}
	}
	newListMirroredWarehousesPager := i.newListMirroredWarehousesPager.get(req)
	if newListMirroredWarehousesPager == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/mirroredWarehouses`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
		if err != nil {
			return nil, err
		}
		continuationTokenUnescaped, err := url.QueryUnescape(qp.Get("continuationToken"))
		if err != nil {
			return nil, err
		}
		continuationTokenParam := getOptional(continuationTokenUnescaped)
		var options *mirroredwarehouse.ItemsClientListMirroredWarehousesOptions
		if continuationTokenParam != nil {
			options = &mirroredwarehouse.ItemsClientListMirroredWarehousesOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := i.srv.NewListMirroredWarehousesPager(workspaceIDParam, options)
		newListMirroredWarehousesPager = &resp
		i.newListMirroredWarehousesPager.add(req, newListMirroredWarehousesPager)
		server.PagerResponderInjectNextLinks(newListMirroredWarehousesPager, req, func(page *mirroredwarehouse.ItemsClientListMirroredWarehousesResponse, createLink func() string) {
			page.ContinuationURI = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListMirroredWarehousesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListMirroredWarehousesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListMirroredWarehousesPager) {
		i.newListMirroredWarehousesPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ItemsServerTransport
var itemsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
