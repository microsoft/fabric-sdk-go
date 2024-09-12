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

	"github.com/microsoft/fabric-sdk-go/fabric/spark"
)

// WorkspaceSettingsServer is a fake server for instances of the spark.WorkspaceSettingsClient type.
type WorkspaceSettingsServer struct {
	// GetSparkSettings is the fake for method WorkspaceSettingsClient.GetSparkSettings
	// HTTP status codes to indicate success: http.StatusOK
	GetSparkSettings func(ctx context.Context, workspaceID string, options *spark.WorkspaceSettingsClientGetSparkSettingsOptions) (resp azfake.Responder[spark.WorkspaceSettingsClientGetSparkSettingsResponse], errResp azfake.ErrorResponder)

	// UpdateSparkSettings is the fake for method WorkspaceSettingsClient.UpdateSparkSettings
	// HTTP status codes to indicate success: http.StatusOK
	UpdateSparkSettings func(ctx context.Context, workspaceID string, updateWorkspaceSettingsRequest spark.UpdateWorkspaceSparkSettingsRequest, options *spark.WorkspaceSettingsClientUpdateSparkSettingsOptions) (resp azfake.Responder[spark.WorkspaceSettingsClientUpdateSparkSettingsResponse], errResp azfake.ErrorResponder)
}

// NewWorkspaceSettingsServerTransport creates a new instance of WorkspaceSettingsServerTransport with the provided implementation.
// The returned WorkspaceSettingsServerTransport instance is connected to an instance of spark.WorkspaceSettingsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkspaceSettingsServerTransport(srv *WorkspaceSettingsServer) *WorkspaceSettingsServerTransport {
	return &WorkspaceSettingsServerTransport{srv: srv}
}

// WorkspaceSettingsServerTransport connects instances of spark.WorkspaceSettingsClient to instances of WorkspaceSettingsServer.
// Don't use this type directly, use NewWorkspaceSettingsServerTransport instead.
type WorkspaceSettingsServerTransport struct {
	srv *WorkspaceSettingsServer
}

// Do implements the policy.Transporter interface for WorkspaceSettingsServerTransport.
func (w *WorkspaceSettingsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	parts := strings.Split(method, ".")
	method = parts[1] + "." + parts[2]
	return w.dispatchToMethodFake(req, method)
}

func (w *WorkspaceSettingsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "WorkspaceSettingsClient.GetSparkSettings":
		resp, err = w.dispatchGetSparkSettings(req)
	case "WorkspaceSettingsClient.UpdateSparkSettings":
		resp, err = w.dispatchUpdateSparkSettings(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (w *WorkspaceSettingsServerTransport) dispatchGetSparkSettings(req *http.Request) (*http.Response, error) {
	if w.srv.GetSparkSettings == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetSparkSettings not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/spark/settings`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.GetSparkSettings(req.Context(), workspaceIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WorkspaceSparkSettings, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkspaceSettingsServerTransport) dispatchUpdateSparkSettings(req *http.Request) (*http.Response, error) {
	if w.srv.UpdateSparkSettings == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateSparkSettings not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/spark/settings`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[spark.UpdateWorkspaceSparkSettingsRequest](req)
	if err != nil {
		return nil, err
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.UpdateSparkSettings(req.Context(), workspaceIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WorkspaceSparkSettings, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
