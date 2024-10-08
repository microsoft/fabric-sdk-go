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

	"github.com/microsoft/fabric-sdk-go/fabric/eventhouse"
)

// ItemsServer is a fake server for instances of the eventhouse.ItemsClient type.
type ItemsServer struct {
	// BeginCreateEventhouse is the fake for method ItemsClient.BeginCreateEventhouse
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateEventhouse func(ctx context.Context, workspaceID string, createEventhouseRequest eventhouse.CreateEventhouseRequest, options *eventhouse.ItemsClientBeginCreateEventhouseOptions) (resp azfake.PollerResponder[eventhouse.ItemsClientCreateEventhouseResponse], errResp azfake.ErrorResponder)

	// DeleteEventhouse is the fake for method ItemsClient.DeleteEventhouse
	// HTTP status codes to indicate success: http.StatusOK
	DeleteEventhouse func(ctx context.Context, workspaceID string, eventhouseID string, options *eventhouse.ItemsClientDeleteEventhouseOptions) (resp azfake.Responder[eventhouse.ItemsClientDeleteEventhouseResponse], errResp azfake.ErrorResponder)

	// GetEventhouse is the fake for method ItemsClient.GetEventhouse
	// HTTP status codes to indicate success: http.StatusOK
	GetEventhouse func(ctx context.Context, workspaceID string, eventhouseID string, options *eventhouse.ItemsClientGetEventhouseOptions) (resp azfake.Responder[eventhouse.ItemsClientGetEventhouseResponse], errResp azfake.ErrorResponder)

	// NewListEventhousesPager is the fake for method ItemsClient.NewListEventhousesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListEventhousesPager func(workspaceID string, options *eventhouse.ItemsClientListEventhousesOptions) (resp azfake.PagerResponder[eventhouse.ItemsClientListEventhousesResponse])

	// UpdateEventhouse is the fake for method ItemsClient.UpdateEventhouse
	// HTTP status codes to indicate success: http.StatusOK
	UpdateEventhouse func(ctx context.Context, workspaceID string, eventhouseID string, updateEventhouseRequest eventhouse.UpdateEventhouseRequest, options *eventhouse.ItemsClientUpdateEventhouseOptions) (resp azfake.Responder[eventhouse.ItemsClientUpdateEventhouseResponse], errResp azfake.ErrorResponder)
}

// NewItemsServerTransport creates a new instance of ItemsServerTransport with the provided implementation.
// The returned ItemsServerTransport instance is connected to an instance of eventhouse.ItemsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewItemsServerTransport(srv *ItemsServer) *ItemsServerTransport {
	return &ItemsServerTransport{
		srv:                     srv,
		beginCreateEventhouse:   newTracker[azfake.PollerResponder[eventhouse.ItemsClientCreateEventhouseResponse]](),
		newListEventhousesPager: newTracker[azfake.PagerResponder[eventhouse.ItemsClientListEventhousesResponse]](),
	}
}

// ItemsServerTransport connects instances of eventhouse.ItemsClient to instances of ItemsServer.
// Don't use this type directly, use NewItemsServerTransport instead.
type ItemsServerTransport struct {
	srv                     *ItemsServer
	beginCreateEventhouse   *tracker[azfake.PollerResponder[eventhouse.ItemsClientCreateEventhouseResponse]]
	newListEventhousesPager *tracker[azfake.PagerResponder[eventhouse.ItemsClientListEventhousesResponse]]
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
		var res result
		switch method {
		case "ItemsClient.BeginCreateEventhouse":
			res.resp, res.err = i.dispatchBeginCreateEventhouse(req)
		case "ItemsClient.DeleteEventhouse":
			res.resp, res.err = i.dispatchDeleteEventhouse(req)
		case "ItemsClient.GetEventhouse":
			res.resp, res.err = i.dispatchGetEventhouse(req)
		case "ItemsClient.NewListEventhousesPager":
			res.resp, res.err = i.dispatchNewListEventhousesPager(req)
		case "ItemsClient.UpdateEventhouse":
			res.resp, res.err = i.dispatchUpdateEventhouse(req)
		default:
			res.err = fmt.Errorf("unhandled API %s", method)
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

func (i *ItemsServerTransport) dispatchBeginCreateEventhouse(req *http.Request) (*http.Response, error) {
	if i.srv.BeginCreateEventhouse == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateEventhouse not implemented")}
	}
	beginCreateEventhouse := i.beginCreateEventhouse.get(req)
	if beginCreateEventhouse == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventhouses`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[eventhouse.CreateEventhouseRequest](req)
		if err != nil {
			return nil, err
		}
		workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginCreateEventhouse(req.Context(), workspaceIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateEventhouse = &respr
		i.beginCreateEventhouse.add(req, beginCreateEventhouse)
	}

	resp, err := server.PollerResponderNext(beginCreateEventhouse, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		i.beginCreateEventhouse.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateEventhouse) {
		i.beginCreateEventhouse.remove(req)
	}

	return resp, nil
}

func (i *ItemsServerTransport) dispatchDeleteEventhouse(req *http.Request) (*http.Response, error) {
	if i.srv.DeleteEventhouse == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteEventhouse not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventhouses/(?P<eventhouseId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	eventhouseIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventhouseId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.DeleteEventhouse(req.Context(), workspaceIDParam, eventhouseIDParam, nil)
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

func (i *ItemsServerTransport) dispatchGetEventhouse(req *http.Request) (*http.Response, error) {
	if i.srv.GetEventhouse == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetEventhouse not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventhouses/(?P<eventhouseId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	eventhouseIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventhouseId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.GetEventhouse(req.Context(), workspaceIDParam, eventhouseIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Eventhouse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchNewListEventhousesPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListEventhousesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListEventhousesPager not implemented")}
	}
	newListEventhousesPager := i.newListEventhousesPager.get(req)
	if newListEventhousesPager == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventhouses`
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
		var options *eventhouse.ItemsClientListEventhousesOptions
		if continuationTokenParam != nil {
			options = &eventhouse.ItemsClientListEventhousesOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := i.srv.NewListEventhousesPager(workspaceIDParam, options)
		newListEventhousesPager = &resp
		i.newListEventhousesPager.add(req, newListEventhousesPager)
		server.PagerResponderInjectNextLinks(newListEventhousesPager, req, func(page *eventhouse.ItemsClientListEventhousesResponse, createLink func() string) {
			page.ContinuationURI = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListEventhousesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListEventhousesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListEventhousesPager) {
		i.newListEventhousesPager.remove(req)
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchUpdateEventhouse(req *http.Request) (*http.Response, error) {
	if i.srv.UpdateEventhouse == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateEventhouse not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventhouses/(?P<eventhouseId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[eventhouse.UpdateEventhouseRequest](req)
	if err != nil {
		return nil, err
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	eventhouseIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventhouseId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.UpdateEventhouse(req.Context(), workspaceIDParam, eventhouseIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Eventhouse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
