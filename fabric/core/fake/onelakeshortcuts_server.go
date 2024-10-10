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

	"github.com/microsoft/fabric-sdk-go/fabric/core"
)

// OneLakeShortcutsServer is a fake server for instances of the core.OneLakeShortcutsClient type.
type OneLakeShortcutsServer struct {
	// CreateShortcut is the fake for method OneLakeShortcutsClient.CreateShortcut
	// HTTP status codes to indicate success: http.StatusCreated
	CreateShortcut func(ctx context.Context, workspaceID string, itemID string, createShortcutRequest core.CreateShortcutRequest, options *core.OneLakeShortcutsClientCreateShortcutOptions) (resp azfake.Responder[core.OneLakeShortcutsClientCreateShortcutResponse], errResp azfake.ErrorResponder)

	// DeleteShortcut is the fake for method OneLakeShortcutsClient.DeleteShortcut
	// HTTP status codes to indicate success: http.StatusOK
	DeleteShortcut func(ctx context.Context, workspaceID string, itemID string, shortcutPath string, shortcutName string, options *core.OneLakeShortcutsClientDeleteShortcutOptions) (resp azfake.Responder[core.OneLakeShortcutsClientDeleteShortcutResponse], errResp azfake.ErrorResponder)

	// GetShortcut is the fake for method OneLakeShortcutsClient.GetShortcut
	// HTTP status codes to indicate success: http.StatusOK
	GetShortcut func(ctx context.Context, workspaceID string, itemID string, shortcutPath string, shortcutName string, options *core.OneLakeShortcutsClientGetShortcutOptions) (resp azfake.Responder[core.OneLakeShortcutsClientGetShortcutResponse], errResp azfake.ErrorResponder)

	// NewListShortcutsPager is the fake for method OneLakeShortcutsClient.NewListShortcutsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListShortcutsPager func(workspaceID string, itemID string, options *core.OneLakeShortcutsClientListShortcutsOptions) (resp azfake.PagerResponder[core.OneLakeShortcutsClientListShortcutsResponse])
}

// NewOneLakeShortcutsServerTransport creates a new instance of OneLakeShortcutsServerTransport with the provided implementation.
// The returned OneLakeShortcutsServerTransport instance is connected to an instance of core.OneLakeShortcutsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOneLakeShortcutsServerTransport(srv *OneLakeShortcutsServer) *OneLakeShortcutsServerTransport {
	return &OneLakeShortcutsServerTransport{
		srv:                   srv,
		newListShortcutsPager: newTracker[azfake.PagerResponder[core.OneLakeShortcutsClientListShortcutsResponse]](),
	}
}

// OneLakeShortcutsServerTransport connects instances of core.OneLakeShortcutsClient to instances of OneLakeShortcutsServer.
// Don't use this type directly, use NewOneLakeShortcutsServerTransport instead.
type OneLakeShortcutsServerTransport struct {
	srv                   *OneLakeShortcutsServer
	newListShortcutsPager *tracker[azfake.PagerResponder[core.OneLakeShortcutsClientListShortcutsResponse]]
}

// Do implements the policy.Transporter interface for OneLakeShortcutsServerTransport.
func (o *OneLakeShortcutsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	parts := strings.Split(method, ".")
	method = parts[1] + "." + parts[2]
	return o.dispatchToMethodFake(req, method)
}

func (o *OneLakeShortcutsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "OneLakeShortcutsClient.CreateShortcut":
			res.resp, res.err = o.dispatchCreateShortcut(req)
		case "OneLakeShortcutsClient.DeleteShortcut":
			res.resp, res.err = o.dispatchDeleteShortcut(req)
		case "OneLakeShortcutsClient.GetShortcut":
			res.resp, res.err = o.dispatchGetShortcut(req)
		case "OneLakeShortcutsClient.NewListShortcutsPager":
			res.resp, res.err = o.dispatchNewListShortcutsPager(req)
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

func (o *OneLakeShortcutsServerTransport) dispatchCreateShortcut(req *http.Request) (*http.Response, error) {
	if o.srv.CreateShortcut == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateShortcut not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items/(?P<itemId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shortcuts`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[core.CreateShortcutRequest](req)
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
	shortcutConflictPolicyUnescaped, err := url.QueryUnescape(qp.Get("shortcutConflictPolicy"))
	if err != nil {
		return nil, err
	}
	shortcutConflictPolicyParam := getOptional(core.ShortcutConflictPolicy(shortcutConflictPolicyUnescaped))
	var options *core.OneLakeShortcutsClientCreateShortcutOptions
	if shortcutConflictPolicyParam != nil {
		options = &core.OneLakeShortcutsClientCreateShortcutOptions{
			ShortcutConflictPolicy: shortcutConflictPolicyParam,
		}
	}
	respr, errRespr := o.srv.CreateShortcut(req.Context(), workspaceIDParam, itemIDParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Shortcut, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).Location; val != nil {
		resp.Header.Set("Location", *val)
	}
	return resp, nil
}

func (o *OneLakeShortcutsServerTransport) dispatchDeleteShortcut(req *http.Request) (*http.Response, error) {
	if o.srv.DeleteShortcut == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteShortcut not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items/(?P<itemId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shortcuts/(?P<shortcutPath>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<shortcutName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
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
	shortcutPathParam, err := url.PathUnescape(matches[regex.SubexpIndex("shortcutPath")])
	if err != nil {
		return nil, err
	}
	shortcutNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("shortcutName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.DeleteShortcut(req.Context(), workspaceIDParam, itemIDParam, shortcutPathParam, shortcutNameParam, nil)
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

func (o *OneLakeShortcutsServerTransport) dispatchGetShortcut(req *http.Request) (*http.Response, error) {
	if o.srv.GetShortcut == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetShortcut not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items/(?P<itemId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shortcuts/(?P<shortcutPath>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<shortcutName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
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
	shortcutPathParam, err := url.PathUnescape(matches[regex.SubexpIndex("shortcutPath")])
	if err != nil {
		return nil, err
	}
	shortcutNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("shortcutName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.GetShortcut(req.Context(), workspaceIDParam, itemIDParam, shortcutPathParam, shortcutNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Shortcut, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OneLakeShortcutsServerTransport) dispatchNewListShortcutsPager(req *http.Request) (*http.Response, error) {
	if o.srv.NewListShortcutsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListShortcutsPager not implemented")}
	}
	newListShortcutsPager := o.newListShortcutsPager.get(req)
	if newListShortcutsPager == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items/(?P<itemId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shortcuts`
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
		parentPathUnescaped, err := url.QueryUnescape(qp.Get("parentPath"))
		if err != nil {
			return nil, err
		}
		parentPathParam := getOptional(parentPathUnescaped)
		continuationTokenUnescaped, err := url.QueryUnescape(qp.Get("continuationToken"))
		if err != nil {
			return nil, err
		}
		continuationTokenParam := getOptional(continuationTokenUnescaped)
		var options *core.OneLakeShortcutsClientListShortcutsOptions
		if parentPathParam != nil || continuationTokenParam != nil {
			options = &core.OneLakeShortcutsClientListShortcutsOptions{
				ParentPath:        parentPathParam,
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := o.srv.NewListShortcutsPager(workspaceIDParam, itemIDParam, options)
		newListShortcutsPager = &resp
		o.newListShortcutsPager.add(req, newListShortcutsPager)
		server.PagerResponderInjectNextLinks(newListShortcutsPager, req, func(page *core.OneLakeShortcutsClientListShortcutsResponse, createLink func() string) {
			page.ContinuationURI = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListShortcutsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		o.newListShortcutsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListShortcutsPager) {
		o.newListShortcutsPager.remove(req)
	}
	return resp, nil
}
