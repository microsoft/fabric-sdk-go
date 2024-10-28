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

// ManagedPrivateEndpointsServer is a fake server for instances of the core.ManagedPrivateEndpointsClient type.
type ManagedPrivateEndpointsServer struct {
	// CreateWorkspaceManagedPrivateEndpoint is the fake for method ManagedPrivateEndpointsClient.CreateWorkspaceManagedPrivateEndpoint
	// HTTP status codes to indicate success: http.StatusCreated
	CreateWorkspaceManagedPrivateEndpoint func(ctx context.Context, workspaceID string, createManagedPrivateEndpointRequest core.CreateManagedPrivateEndpointRequest, options *core.ManagedPrivateEndpointsClientCreateWorkspaceManagedPrivateEndpointOptions) (resp azfake.Responder[core.ManagedPrivateEndpointsClientCreateWorkspaceManagedPrivateEndpointResponse], errResp azfake.ErrorResponder)

	// DeleteWorkspaceManagedPrivateEndpoint is the fake for method ManagedPrivateEndpointsClient.DeleteWorkspaceManagedPrivateEndpoint
	// HTTP status codes to indicate success: http.StatusOK
	DeleteWorkspaceManagedPrivateEndpoint func(ctx context.Context, workspaceID string, managedPrivateEndpointID string, options *core.ManagedPrivateEndpointsClientDeleteWorkspaceManagedPrivateEndpointOptions) (resp azfake.Responder[core.ManagedPrivateEndpointsClientDeleteWorkspaceManagedPrivateEndpointResponse], errResp azfake.ErrorResponder)

	// GetWorkspaceManagedPrivateEndpoint is the fake for method ManagedPrivateEndpointsClient.GetWorkspaceManagedPrivateEndpoint
	// HTTP status codes to indicate success: http.StatusOK
	GetWorkspaceManagedPrivateEndpoint func(ctx context.Context, workspaceID string, managedPrivateEndpointID string, options *core.ManagedPrivateEndpointsClientGetWorkspaceManagedPrivateEndpointOptions) (resp azfake.Responder[core.ManagedPrivateEndpointsClientGetWorkspaceManagedPrivateEndpointResponse], errResp azfake.ErrorResponder)

	// NewListWorkspaceManagedPrivateEndpointsPager is the fake for method ManagedPrivateEndpointsClient.NewListWorkspaceManagedPrivateEndpointsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListWorkspaceManagedPrivateEndpointsPager func(workspaceID string, options *core.ManagedPrivateEndpointsClientListWorkspaceManagedPrivateEndpointsOptions) (resp azfake.PagerResponder[core.ManagedPrivateEndpointsClientListWorkspaceManagedPrivateEndpointsResponse])
}

// NewManagedPrivateEndpointsServerTransport creates a new instance of ManagedPrivateEndpointsServerTransport with the provided implementation.
// The returned ManagedPrivateEndpointsServerTransport instance is connected to an instance of core.ManagedPrivateEndpointsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedPrivateEndpointsServerTransport(srv *ManagedPrivateEndpointsServer) *ManagedPrivateEndpointsServerTransport {
	return &ManagedPrivateEndpointsServerTransport{
		srv: srv,
		newListWorkspaceManagedPrivateEndpointsPager: newTracker[azfake.PagerResponder[core.ManagedPrivateEndpointsClientListWorkspaceManagedPrivateEndpointsResponse]](),
	}
}

// ManagedPrivateEndpointsServerTransport connects instances of core.ManagedPrivateEndpointsClient to instances of ManagedPrivateEndpointsServer.
// Don't use this type directly, use NewManagedPrivateEndpointsServerTransport instead.
type ManagedPrivateEndpointsServerTransport struct {
	srv                                          *ManagedPrivateEndpointsServer
	newListWorkspaceManagedPrivateEndpointsPager *tracker[azfake.PagerResponder[core.ManagedPrivateEndpointsClientListWorkspaceManagedPrivateEndpointsResponse]]
}

// Do implements the policy.Transporter interface for ManagedPrivateEndpointsServerTransport.
func (m *ManagedPrivateEndpointsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	parts := strings.Split(method, ".")
	method = parts[1] + "." + parts[2]
	return m.dispatchToMethodFake(req, method)
}

func (m *ManagedPrivateEndpointsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "ManagedPrivateEndpointsClient.CreateWorkspaceManagedPrivateEndpoint":
			res.resp, res.err = m.dispatchCreateWorkspaceManagedPrivateEndpoint(req)
		case "ManagedPrivateEndpointsClient.DeleteWorkspaceManagedPrivateEndpoint":
			res.resp, res.err = m.dispatchDeleteWorkspaceManagedPrivateEndpoint(req)
		case "ManagedPrivateEndpointsClient.GetWorkspaceManagedPrivateEndpoint":
			res.resp, res.err = m.dispatchGetWorkspaceManagedPrivateEndpoint(req)
		case "ManagedPrivateEndpointsClient.NewListWorkspaceManagedPrivateEndpointsPager":
			res.resp, res.err = m.dispatchNewListWorkspaceManagedPrivateEndpointsPager(req)
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

func (m *ManagedPrivateEndpointsServerTransport) dispatchCreateWorkspaceManagedPrivateEndpoint(req *http.Request) (*http.Response, error) {
	if m.srv.CreateWorkspaceManagedPrivateEndpoint == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateWorkspaceManagedPrivateEndpoint not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedPrivateEndpoints`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[core.CreateManagedPrivateEndpointRequest](req)
	if err != nil {
		return nil, err
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.CreateWorkspaceManagedPrivateEndpoint(req.Context(), workspaceIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedPrivateEndpoint, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).Location; val != nil {
		resp.Header.Set("Location", *val)
	}
	return resp, nil
}

func (m *ManagedPrivateEndpointsServerTransport) dispatchDeleteWorkspaceManagedPrivateEndpoint(req *http.Request) (*http.Response, error) {
	if m.srv.DeleteWorkspaceManagedPrivateEndpoint == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteWorkspaceManagedPrivateEndpoint not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedPrivateEndpoints/(?P<managedPrivateEndpointId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	managedPrivateEndpointIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedPrivateEndpointId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.DeleteWorkspaceManagedPrivateEndpoint(req.Context(), workspaceIDParam, managedPrivateEndpointIDParam, nil)
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

func (m *ManagedPrivateEndpointsServerTransport) dispatchGetWorkspaceManagedPrivateEndpoint(req *http.Request) (*http.Response, error) {
	if m.srv.GetWorkspaceManagedPrivateEndpoint == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetWorkspaceManagedPrivateEndpoint not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedPrivateEndpoints/(?P<managedPrivateEndpointId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	managedPrivateEndpointIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedPrivateEndpointId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.GetWorkspaceManagedPrivateEndpoint(req.Context(), workspaceIDParam, managedPrivateEndpointIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedPrivateEndpoint, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedPrivateEndpointsServerTransport) dispatchNewListWorkspaceManagedPrivateEndpointsPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListWorkspaceManagedPrivateEndpointsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListWorkspaceManagedPrivateEndpointsPager not implemented")}
	}
	newListWorkspaceManagedPrivateEndpointsPager := m.newListWorkspaceManagedPrivateEndpointsPager.get(req)
	if newListWorkspaceManagedPrivateEndpointsPager == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedPrivateEndpoints`
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
		var options *core.ManagedPrivateEndpointsClientListWorkspaceManagedPrivateEndpointsOptions
		if continuationTokenParam != nil {
			options = &core.ManagedPrivateEndpointsClientListWorkspaceManagedPrivateEndpointsOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := m.srv.NewListWorkspaceManagedPrivateEndpointsPager(workspaceIDParam, options)
		newListWorkspaceManagedPrivateEndpointsPager = &resp
		m.newListWorkspaceManagedPrivateEndpointsPager.add(req, newListWorkspaceManagedPrivateEndpointsPager)
		server.PagerResponderInjectNextLinks(newListWorkspaceManagedPrivateEndpointsPager, req, func(page *core.ManagedPrivateEndpointsClientListWorkspaceManagedPrivateEndpointsResponse, createLink func() string) {
			page.ContinuationURI = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListWorkspaceManagedPrivateEndpointsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListWorkspaceManagedPrivateEndpointsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListWorkspaceManagedPrivateEndpointsPager) {
		m.newListWorkspaceManagedPrivateEndpointsPager.remove(req)
	}
	return resp, nil
}
