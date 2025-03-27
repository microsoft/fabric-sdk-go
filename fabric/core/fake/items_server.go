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
	"strconv"
	"strings"

	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"

	"github.com/microsoft/fabric-sdk-go/fabric/core"
)

// ItemsServer is a fake server for instances of the core.ItemsClient type.
type ItemsServer struct {
	// BeginCreateItem is the fake for method ItemsClient.BeginCreateItem
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateItem func(ctx context.Context, workspaceID string, createItemRequest core.CreateItemRequest, options *core.ItemsClientBeginCreateItemOptions) (resp azfake.PollerResponder[core.ItemsClientCreateItemResponse], errResp azfake.ErrorResponder)

	// DeleteItem is the fake for method ItemsClient.DeleteItem
	// HTTP status codes to indicate success: http.StatusOK
	DeleteItem func(ctx context.Context, workspaceID string, itemID string, options *core.ItemsClientDeleteItemOptions) (resp azfake.Responder[core.ItemsClientDeleteItemResponse], errResp azfake.ErrorResponder)

	// GetItem is the fake for method ItemsClient.GetItem
	// HTTP status codes to indicate success: http.StatusOK
	GetItem func(ctx context.Context, workspaceID string, itemID string, options *core.ItemsClientGetItemOptions) (resp azfake.Responder[core.ItemsClientGetItemResponse], errResp azfake.ErrorResponder)

	// BeginGetItemDefinition is the fake for method ItemsClient.BeginGetItemDefinition
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGetItemDefinition func(ctx context.Context, workspaceID string, itemID string, options *core.ItemsClientBeginGetItemDefinitionOptions) (resp azfake.PollerResponder[core.ItemsClientGetItemDefinitionResponse], errResp azfake.ErrorResponder)

	// NewListItemConnectionsPager is the fake for method ItemsClient.NewListItemConnectionsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListItemConnectionsPager func(workspaceID string, itemID string, options *core.ItemsClientListItemConnectionsOptions) (resp azfake.PagerResponder[core.ItemsClientListItemConnectionsResponse])

	// NewListItemsPager is the fake for method ItemsClient.NewListItemsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListItemsPager func(workspaceID string, options *core.ItemsClientListItemsOptions) (resp azfake.PagerResponder[core.ItemsClientListItemsResponse])

	// UpdateItem is the fake for method ItemsClient.UpdateItem
	// HTTP status codes to indicate success: http.StatusOK
	UpdateItem func(ctx context.Context, workspaceID string, itemID string, updateItemRequest core.UpdateItemRequest, options *core.ItemsClientUpdateItemOptions) (resp azfake.Responder[core.ItemsClientUpdateItemResponse], errResp azfake.ErrorResponder)

	// BeginUpdateItemDefinition is the fake for method ItemsClient.BeginUpdateItemDefinition
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginUpdateItemDefinition func(ctx context.Context, workspaceID string, itemID string, updateItemDefinitionRequest core.UpdateItemDefinitionRequest, options *core.ItemsClientBeginUpdateItemDefinitionOptions) (resp azfake.PollerResponder[core.ItemsClientUpdateItemDefinitionResponse], errResp azfake.ErrorResponder)
}

// NewItemsServerTransport creates a new instance of ItemsServerTransport with the provided implementation.
// The returned ItemsServerTransport instance is connected to an instance of core.ItemsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewItemsServerTransport(srv *ItemsServer) *ItemsServerTransport {
	return &ItemsServerTransport{
		srv:                         srv,
		beginCreateItem:             newTracker[azfake.PollerResponder[core.ItemsClientCreateItemResponse]](),
		beginGetItemDefinition:      newTracker[azfake.PollerResponder[core.ItemsClientGetItemDefinitionResponse]](),
		newListItemConnectionsPager: newTracker[azfake.PagerResponder[core.ItemsClientListItemConnectionsResponse]](),
		newListItemsPager:           newTracker[azfake.PagerResponder[core.ItemsClientListItemsResponse]](),
		beginUpdateItemDefinition:   newTracker[azfake.PollerResponder[core.ItemsClientUpdateItemDefinitionResponse]](),
	}
}

// ItemsServerTransport connects instances of core.ItemsClient to instances of ItemsServer.
// Don't use this type directly, use NewItemsServerTransport instead.
type ItemsServerTransport struct {
	srv                         *ItemsServer
	beginCreateItem             *tracker[azfake.PollerResponder[core.ItemsClientCreateItemResponse]]
	beginGetItemDefinition      *tracker[azfake.PollerResponder[core.ItemsClientGetItemDefinitionResponse]]
	newListItemConnectionsPager *tracker[azfake.PagerResponder[core.ItemsClientListItemConnectionsResponse]]
	newListItemsPager           *tracker[azfake.PagerResponder[core.ItemsClientListItemsResponse]]
	beginUpdateItemDefinition   *tracker[azfake.PollerResponder[core.ItemsClientUpdateItemDefinitionResponse]]
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
			case "ItemsClient.BeginCreateItem":
				res.resp, res.err = i.dispatchBeginCreateItem(req)
			case "ItemsClient.DeleteItem":
				res.resp, res.err = i.dispatchDeleteItem(req)
			case "ItemsClient.GetItem":
				res.resp, res.err = i.dispatchGetItem(req)
			case "ItemsClient.BeginGetItemDefinition":
				res.resp, res.err = i.dispatchBeginGetItemDefinition(req)
			case "ItemsClient.NewListItemConnectionsPager":
				res.resp, res.err = i.dispatchNewListItemConnectionsPager(req)
			case "ItemsClient.NewListItemsPager":
				res.resp, res.err = i.dispatchNewListItemsPager(req)
			case "ItemsClient.UpdateItem":
				res.resp, res.err = i.dispatchUpdateItem(req)
			case "ItemsClient.BeginUpdateItemDefinition":
				res.resp, res.err = i.dispatchBeginUpdateItemDefinition(req)
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

func (i *ItemsServerTransport) dispatchBeginCreateItem(req *http.Request) (*http.Response, error) {
	if i.srv.BeginCreateItem == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateItem not implemented")}
	}
	beginCreateItem := i.beginCreateItem.get(req)
	if beginCreateItem == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[core.CreateItemRequest](req)
		if err != nil {
			return nil, err
		}
		workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginCreateItem(req.Context(), workspaceIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateItem = &respr
		i.beginCreateItem.add(req, beginCreateItem)
	}

	resp, err := server.PollerResponderNext(beginCreateItem, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		i.beginCreateItem.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateItem) {
		i.beginCreateItem.remove(req)
	}

	return resp, nil
}

func (i *ItemsServerTransport) dispatchDeleteItem(req *http.Request) (*http.Response, error) {
	if i.srv.DeleteItem == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteItem not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items/(?P<itemId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	itemIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("itemId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.DeleteItem(req.Context(), workspaceIDParam, itemIDParam, nil)
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

func (i *ItemsServerTransport) dispatchGetItem(req *http.Request) (*http.Response, error) {
	if i.srv.GetItem == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetItem not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items/(?P<itemId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	itemIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("itemId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.GetItem(req.Context(), workspaceIDParam, itemIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Item, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchBeginGetItemDefinition(req *http.Request) (*http.Response, error) {
	if i.srv.BeginGetItemDefinition == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGetItemDefinition not implemented")}
	}
	beginGetItemDefinition := i.beginGetItemDefinition.get(req)
	if beginGetItemDefinition == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items/(?P<itemId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getDefinition`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
		if err != nil {
			return nil, err
		}
		itemIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("itemId")])
		if err != nil {
			return nil, err
		}
		formatUnescaped, err := url.QueryUnescape(qp.Get("format"))
		if err != nil {
			return nil, err
		}
		formatParam := getOptional(formatUnescaped)
		var options *core.ItemsClientBeginGetItemDefinitionOptions
		if formatParam != nil {
			options = &core.ItemsClientBeginGetItemDefinitionOptions{
				Format: formatParam,
			}
		}
		respr, errRespr := i.srv.BeginGetItemDefinition(req.Context(), workspaceIDParam, itemIDParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGetItemDefinition = &respr
		i.beginGetItemDefinition.add(req, beginGetItemDefinition)
	}

	resp, err := server.PollerResponderNext(beginGetItemDefinition, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		i.beginGetItemDefinition.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGetItemDefinition) {
		i.beginGetItemDefinition.remove(req)
	}

	return resp, nil
}

func (i *ItemsServerTransport) dispatchNewListItemConnectionsPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListItemConnectionsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListItemConnectionsPager not implemented")}
	}
	newListItemConnectionsPager := i.newListItemConnectionsPager.get(req)
	if newListItemConnectionsPager == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items/(?P<itemId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
		if err != nil {
			return nil, err
		}
		itemIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("itemId")])
		if err != nil {
			return nil, err
		}
		continuationTokenUnescaped, err := url.QueryUnescape(qp.Get("continuationToken"))
		if err != nil {
			return nil, err
		}
		continuationTokenParam := getOptional(continuationTokenUnescaped)
		var options *core.ItemsClientListItemConnectionsOptions
		if continuationTokenParam != nil {
			options = &core.ItemsClientListItemConnectionsOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := i.srv.NewListItemConnectionsPager(workspaceIDParam, itemIDParam, options)
		newListItemConnectionsPager = &resp
		i.newListItemConnectionsPager.add(req, newListItemConnectionsPager)
		server.PagerResponderInjectNextLinks(newListItemConnectionsPager, req, func(page *core.ItemsClientListItemConnectionsResponse, createLink func() string) {
			page.ContinuationURI = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListItemConnectionsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListItemConnectionsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListItemConnectionsPager) {
		i.newListItemConnectionsPager.remove(req)
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchNewListItemsPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListItemsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListItemsPager not implemented")}
	}
	newListItemsPager := i.newListItemsPager.get(req)
	if newListItemsPager == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items`
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
		typeUnescaped, err := url.QueryUnescape(qp.Get("type"))
		if err != nil {
			return nil, err
		}
		typeParam := getOptional(typeUnescaped)
		recursiveUnescaped, err := url.QueryUnescape(qp.Get("recursive"))
		if err != nil {
			return nil, err
		}
		recursiveParam, err := parseOptional(recursiveUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		rootFolderIDUnescaped, err := url.QueryUnescape(qp.Get("rootFolderId"))
		if err != nil {
			return nil, err
		}
		rootFolderIDParam := getOptional(rootFolderIDUnescaped)
		continuationTokenUnescaped, err := url.QueryUnescape(qp.Get("continuationToken"))
		if err != nil {
			return nil, err
		}
		continuationTokenParam := getOptional(continuationTokenUnescaped)
		var options *core.ItemsClientListItemsOptions
		if typeParam != nil || recursiveParam != nil || rootFolderIDParam != nil || continuationTokenParam != nil {
			options = &core.ItemsClientListItemsOptions{
				Type:              typeParam,
				Recursive:         recursiveParam,
				RootFolderID:      rootFolderIDParam,
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := i.srv.NewListItemsPager(workspaceIDParam, options)
		newListItemsPager = &resp
		i.newListItemsPager.add(req, newListItemsPager)
		server.PagerResponderInjectNextLinks(newListItemsPager, req, func(page *core.ItemsClientListItemsResponse, createLink func() string) {
			page.ContinuationURI = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListItemsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListItemsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListItemsPager) {
		i.newListItemsPager.remove(req)
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchUpdateItem(req *http.Request) (*http.Response, error) {
	if i.srv.UpdateItem == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateItem not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items/(?P<itemId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[core.UpdateItemRequest](req)
	if err != nil {
		return nil, err
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	itemIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("itemId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.UpdateItem(req.Context(), workspaceIDParam, itemIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Item, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchBeginUpdateItemDefinition(req *http.Request) (*http.Response, error) {
	if i.srv.BeginUpdateItemDefinition == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateItemDefinition not implemented")}
	}
	beginUpdateItemDefinition := i.beginUpdateItemDefinition.get(req)
	if beginUpdateItemDefinition == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items/(?P<itemId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateDefinition`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		body, err := server.UnmarshalRequestAsJSON[core.UpdateItemDefinitionRequest](req)
		if err != nil {
			return nil, err
		}
		workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
		if err != nil {
			return nil, err
		}
		itemIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("itemId")])
		if err != nil {
			return nil, err
		}
		updateMetadataUnescaped, err := url.QueryUnescape(qp.Get("updateMetadata"))
		if err != nil {
			return nil, err
		}
		updateMetadataParam, err := parseOptional(updateMetadataUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		var options *core.ItemsClientBeginUpdateItemDefinitionOptions
		if updateMetadataParam != nil {
			options = &core.ItemsClientBeginUpdateItemDefinitionOptions{
				UpdateMetadata: updateMetadataParam,
			}
		}
		respr, errRespr := i.srv.BeginUpdateItemDefinition(req.Context(), workspaceIDParam, itemIDParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateItemDefinition = &respr
		i.beginUpdateItemDefinition.add(req, beginUpdateItemDefinition)
	}

	resp, err := server.PollerResponderNext(beginUpdateItemDefinition, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		i.beginUpdateItemDefinition.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateItemDefinition) {
		i.beginUpdateItemDefinition.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ItemsServerTransport
var itemsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
