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

	"github.com/microsoft/fabric-sdk-go/fabric/admin"
)

// ExternalDataSharesServer is a fake server for instances of the admin.ExternalDataSharesClient type.
type ExternalDataSharesServer struct {
	// NewListExternalDataSharesPager is the fake for method ExternalDataSharesClient.NewListExternalDataSharesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListExternalDataSharesPager func(options *admin.ExternalDataSharesClientListExternalDataSharesOptions) (resp azfake.PagerResponder[admin.ExternalDataSharesClientListExternalDataSharesResponse])

	// RevokeExternalDataShare is the fake for method ExternalDataSharesClient.RevokeExternalDataShare
	// HTTP status codes to indicate success: http.StatusOK
	RevokeExternalDataShare func(ctx context.Context, workspaceID string, itemID string, externalDataShareID string, options *admin.ExternalDataSharesClientRevokeExternalDataShareOptions) (resp azfake.Responder[admin.ExternalDataSharesClientRevokeExternalDataShareResponse], errResp azfake.ErrorResponder)
}

// NewExternalDataSharesServerTransport creates a new instance of ExternalDataSharesServerTransport with the provided implementation.
// The returned ExternalDataSharesServerTransport instance is connected to an instance of admin.ExternalDataSharesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExternalDataSharesServerTransport(srv *ExternalDataSharesServer) *ExternalDataSharesServerTransport {
	return &ExternalDataSharesServerTransport{
		srv:                            srv,
		newListExternalDataSharesPager: newTracker[azfake.PagerResponder[admin.ExternalDataSharesClientListExternalDataSharesResponse]](),
	}
}

// ExternalDataSharesServerTransport connects instances of admin.ExternalDataSharesClient to instances of ExternalDataSharesServer.
// Don't use this type directly, use NewExternalDataSharesServerTransport instead.
type ExternalDataSharesServerTransport struct {
	srv                            *ExternalDataSharesServer
	newListExternalDataSharesPager *tracker[azfake.PagerResponder[admin.ExternalDataSharesClientListExternalDataSharesResponse]]
}

// Do implements the policy.Transporter interface for ExternalDataSharesServerTransport.
func (e *ExternalDataSharesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	parts := strings.Split(method, ".")
	method = parts[1] + "." + parts[2]
	return e.dispatchToMethodFake(req, method)
}

func (e *ExternalDataSharesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "ExternalDataSharesClient.NewListExternalDataSharesPager":
		resp, err = e.dispatchNewListExternalDataSharesPager(req)
	case "ExternalDataSharesClient.RevokeExternalDataShare":
		resp, err = e.dispatchRevokeExternalDataShare(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (e *ExternalDataSharesServerTransport) dispatchNewListExternalDataSharesPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListExternalDataSharesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListExternalDataSharesPager not implemented")}
	}
	newListExternalDataSharesPager := e.newListExternalDataSharesPager.get(req)
	if newListExternalDataSharesPager == nil {
		qp := req.URL.Query()
		continuationTokenUnescaped, err := url.QueryUnescape(qp.Get("continuationToken"))
		if err != nil {
			return nil, err
		}
		continuationTokenParam := getOptional(continuationTokenUnescaped)
		var options *admin.ExternalDataSharesClientListExternalDataSharesOptions
		if continuationTokenParam != nil {
			options = &admin.ExternalDataSharesClientListExternalDataSharesOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := e.srv.NewListExternalDataSharesPager(options)
		newListExternalDataSharesPager = &resp
		e.newListExternalDataSharesPager.add(req, newListExternalDataSharesPager)
		server.PagerResponderInjectNextLinks(newListExternalDataSharesPager, req, func(page *admin.ExternalDataSharesClientListExternalDataSharesResponse, createLink func() string) {
			page.ContinuationURI = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListExternalDataSharesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListExternalDataSharesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListExternalDataSharesPager) {
		e.newListExternalDataSharesPager.remove(req)
	}
	return resp, nil
}

func (e *ExternalDataSharesServerTransport) dispatchRevokeExternalDataShare(req *http.Request) (*http.Response, error) {
	if e.srv.RevokeExternalDataShare == nil {
		return nil, &nonRetriableError{errors.New("fake for method RevokeExternalDataShare not implemented")}
	}
	const regexStr = `/v1/admin/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items/(?P<itemId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/externalDataShares/(?P<externalDataShareId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/revoke`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
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
	externalDataShareIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("externalDataShareId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.RevokeExternalDataShare(req.Context(), workspaceIDParam, itemIDParam, externalDataShareIDParam, nil)
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
