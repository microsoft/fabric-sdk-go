// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package fake

import (
	"context"
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

	"github.com/microsoft/fabric-sdk-go/fabric/warehouse"
)

// ItemsServer is a fake server for instances of the warehouse.ItemsClient type.
type ItemsServer struct {
	// BeginCreateWarehouse is the fake for method ItemsClient.BeginCreateWarehouse
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateWarehouse func(ctx context.Context, workspaceID string, createWarehouseRequest warehouse.CreateWarehouseRequest, options *warehouse.ItemsClientBeginCreateWarehouseOptions) (resp azfake.PollerResponder[warehouse.ItemsClientCreateWarehouseResponse], errResp azfake.ErrorResponder)

	// DeleteWarehouse is the fake for method ItemsClient.DeleteWarehouse
	// HTTP status codes to indicate success: http.StatusOK
	DeleteWarehouse func(ctx context.Context, workspaceID string, warehouseID string, options *warehouse.ItemsClientDeleteWarehouseOptions) (resp azfake.Responder[warehouse.ItemsClientDeleteWarehouseResponse], errResp azfake.ErrorResponder)

	// GetWarehouse is the fake for method ItemsClient.GetWarehouse
	// HTTP status codes to indicate success: http.StatusOK
	GetWarehouse func(ctx context.Context, workspaceID string, warehouseID string, options *warehouse.ItemsClientGetWarehouseOptions) (resp azfake.Responder[warehouse.ItemsClientGetWarehouseResponse], errResp azfake.ErrorResponder)

	// NewListWarehousesPager is the fake for method ItemsClient.NewListWarehousesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListWarehousesPager func(workspaceID string, options *warehouse.ItemsClientListWarehousesOptions) (resp azfake.PagerResponder[warehouse.ItemsClientListWarehousesResponse])

	// UpdateWarehouse is the fake for method ItemsClient.UpdateWarehouse
	// HTTP status codes to indicate success: http.StatusOK
	UpdateWarehouse func(ctx context.Context, workspaceID string, warehouseID string, updateWarehouseRequest warehouse.UpdateWarehouseRequest, options *warehouse.ItemsClientUpdateWarehouseOptions) (resp azfake.Responder[warehouse.ItemsClientUpdateWarehouseResponse], errResp azfake.ErrorResponder)
}

// NewItemsServerTransport creates a new instance of ItemsServerTransport with the provided implementation.
// The returned ItemsServerTransport instance is connected to an instance of warehouse.ItemsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewItemsServerTransport(srv *ItemsServer) *ItemsServerTransport {
	return &ItemsServerTransport{
		srv:                    srv,
		beginCreateWarehouse:   newTracker[azfake.PollerResponder[warehouse.ItemsClientCreateWarehouseResponse]](),
		newListWarehousesPager: newTracker[azfake.PagerResponder[warehouse.ItemsClientListWarehousesResponse]](),
	}
}

// ItemsServerTransport connects instances of warehouse.ItemsClient to instances of ItemsServer.
// Don't use this type directly, use NewItemsServerTransport instead.
type ItemsServerTransport struct {
	srv                    *ItemsServer
	beginCreateWarehouse   *tracker[azfake.PollerResponder[warehouse.ItemsClientCreateWarehouseResponse]]
	newListWarehousesPager *tracker[azfake.PagerResponder[warehouse.ItemsClientListWarehousesResponse]]
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
			case "ItemsClient.BeginCreateWarehouse":
				res.resp, res.err = i.dispatchBeginCreateWarehouse(req)
			case "ItemsClient.DeleteWarehouse":
				res.resp, res.err = i.dispatchDeleteWarehouse(req)
			case "ItemsClient.GetWarehouse":
				res.resp, res.err = i.dispatchGetWarehouse(req)
			case "ItemsClient.NewListWarehousesPager":
				res.resp, res.err = i.dispatchNewListWarehousesPager(req)
			case "ItemsClient.UpdateWarehouse":
				res.resp, res.err = i.dispatchUpdateWarehouse(req)
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

func (i *ItemsServerTransport) dispatchBeginCreateWarehouse(req *http.Request) (*http.Response, error) {
	if i.srv.BeginCreateWarehouse == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateWarehouse not implemented")}
	}
	beginCreateWarehouse := i.beginCreateWarehouse.get(req)
	if beginCreateWarehouse == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/warehouses`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[warehouse.CreateWarehouseRequest](req)
		if err != nil {
			return nil, err
		}
		workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginCreateWarehouse(req.Context(), workspaceIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateWarehouse = &respr
		i.beginCreateWarehouse.add(req, beginCreateWarehouse)
	}

	resp, err := server.PollerResponderNext(beginCreateWarehouse, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		i.beginCreateWarehouse.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateWarehouse) {
		i.beginCreateWarehouse.remove(req)
	}

	return resp, nil
}

func (i *ItemsServerTransport) dispatchDeleteWarehouse(req *http.Request) (*http.Response, error) {
	if i.srv.DeleteWarehouse == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteWarehouse not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/warehouses/(?P<warehouseId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	warehouseIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("warehouseId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.DeleteWarehouse(req.Context(), workspaceIDParam, warehouseIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchGetWarehouse(req *http.Request) (*http.Response, error) {
	if i.srv.GetWarehouse == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetWarehouse not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/warehouses/(?P<warehouseId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	warehouseIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("warehouseId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.GetWarehouse(req.Context(), workspaceIDParam, warehouseIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Warehouse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchNewListWarehousesPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListWarehousesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListWarehousesPager not implemented")}
	}
	newListWarehousesPager := i.newListWarehousesPager.get(req)
	if newListWarehousesPager == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/warehouses`
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
		var options *warehouse.ItemsClientListWarehousesOptions
		if continuationTokenParam != nil {
			options = &warehouse.ItemsClientListWarehousesOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := i.srv.NewListWarehousesPager(workspaceIDParam, options)
		newListWarehousesPager = &resp
		i.newListWarehousesPager.add(req, newListWarehousesPager)
		server.PagerResponderInjectNextLinks(newListWarehousesPager, req, func(page *warehouse.ItemsClientListWarehousesResponse, createLink func() string) {
			page.ContinuationURI = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListWarehousesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListWarehousesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListWarehousesPager) {
		i.newListWarehousesPager.remove(req)
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchUpdateWarehouse(req *http.Request) (*http.Response, error) {
	if i.srv.UpdateWarehouse == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateWarehouse not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/warehouses/(?P<warehouseId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[warehouse.UpdateWarehouseRequest](req)
	if err != nil {
		return nil, err
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	warehouseIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("warehouseId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.UpdateWarehouse(req.Context(), workspaceIDParam, warehouseIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Warehouse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ItemsServerTransport
var itemsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
