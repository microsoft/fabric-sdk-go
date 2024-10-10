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

	"github.com/microsoft/fabric-sdk-go/fabric/sqlendpoint"
)

// ItemsServer is a fake server for instances of the sqlendpoint.ItemsClient type.
type ItemsServer struct {
	// NewListSQLEndpointsPager is the fake for method ItemsClient.NewListSQLEndpointsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListSQLEndpointsPager func(workspaceID string, options *sqlendpoint.ItemsClientListSQLEndpointsOptions) (resp azfake.PagerResponder[sqlendpoint.ItemsClientListSQLEndpointsResponse])
}

// NewItemsServerTransport creates a new instance of ItemsServerTransport with the provided implementation.
// The returned ItemsServerTransport instance is connected to an instance of sqlendpoint.ItemsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewItemsServerTransport(srv *ItemsServer) *ItemsServerTransport {
	return &ItemsServerTransport{
		srv:                      srv,
		newListSQLEndpointsPager: newTracker[azfake.PagerResponder[sqlendpoint.ItemsClientListSQLEndpointsResponse]](),
	}
}

// ItemsServerTransport connects instances of sqlendpoint.ItemsClient to instances of ItemsServer.
// Don't use this type directly, use NewItemsServerTransport instead.
type ItemsServerTransport struct {
	srv                      *ItemsServer
	newListSQLEndpointsPager *tracker[azfake.PagerResponder[sqlendpoint.ItemsClientListSQLEndpointsResponse]]
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
		case "ItemsClient.NewListSQLEndpointsPager":
			res.resp, res.err = i.dispatchNewListSQLEndpointsPager(req)
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

func (i *ItemsServerTransport) dispatchNewListSQLEndpointsPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListSQLEndpointsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListSQLEndpointsPager not implemented")}
	}
	newListSQLEndpointsPager := i.newListSQLEndpointsPager.get(req)
	if newListSQLEndpointsPager == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlEndpoints`
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
		var options *sqlendpoint.ItemsClientListSQLEndpointsOptions
		if continuationTokenParam != nil {
			options = &sqlendpoint.ItemsClientListSQLEndpointsOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := i.srv.NewListSQLEndpointsPager(workspaceIDParam, options)
		newListSQLEndpointsPager = &resp
		i.newListSQLEndpointsPager.add(req, newListSQLEndpointsPager)
		server.PagerResponderInjectNextLinks(newListSQLEndpointsPager, req, func(page *sqlendpoint.ItemsClientListSQLEndpointsResponse, createLink func() string) {
			page.ContinuationURI = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListSQLEndpointsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListSQLEndpointsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListSQLEndpointsPager) {
		i.newListSQLEndpointsPager.remove(req)
	}
	return resp, nil
}
