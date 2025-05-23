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

	"github.com/microsoft/fabric-sdk-go/fabric/sqldatabase"
)

// ItemsServer is a fake server for instances of the sqldatabase.ItemsClient type.
type ItemsServer struct {
	// BeginCreateSQLDatabase is the fake for method ItemsClient.BeginCreateSQLDatabase
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateSQLDatabase func(ctx context.Context, workspaceID string, createSQLDatabaseRequest sqldatabase.CreateSQLDatabaseRequest, options *sqldatabase.ItemsClientBeginCreateSQLDatabaseOptions) (resp azfake.PollerResponder[sqldatabase.ItemsClientCreateSQLDatabaseResponse], errResp azfake.ErrorResponder)

	// DeleteSQLDatabase is the fake for method ItemsClient.DeleteSQLDatabase
	// HTTP status codes to indicate success: http.StatusOK
	DeleteSQLDatabase func(ctx context.Context, workspaceID string, sqlDatabaseID string, options *sqldatabase.ItemsClientDeleteSQLDatabaseOptions) (resp azfake.Responder[sqldatabase.ItemsClientDeleteSQLDatabaseResponse], errResp azfake.ErrorResponder)

	// GetSQLDatabase is the fake for method ItemsClient.GetSQLDatabase
	// HTTP status codes to indicate success: http.StatusOK
	GetSQLDatabase func(ctx context.Context, workspaceID string, sqlDatabaseID string, options *sqldatabase.ItemsClientGetSQLDatabaseOptions) (resp azfake.Responder[sqldatabase.ItemsClientGetSQLDatabaseResponse], errResp azfake.ErrorResponder)

	// NewListSQLDatabasesPager is the fake for method ItemsClient.NewListSQLDatabasesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListSQLDatabasesPager func(workspaceID string, options *sqldatabase.ItemsClientListSQLDatabasesOptions) (resp azfake.PagerResponder[sqldatabase.ItemsClientListSQLDatabasesResponse])

	// UpdateSQLDatabase is the fake for method ItemsClient.UpdateSQLDatabase
	// HTTP status codes to indicate success: http.StatusOK
	UpdateSQLDatabase func(ctx context.Context, workspaceID string, sqlDatabaseID string, updateSQLDatabaseRequest sqldatabase.UpdateSQLDatabaseRequest, options *sqldatabase.ItemsClientUpdateSQLDatabaseOptions) (resp azfake.Responder[sqldatabase.ItemsClientUpdateSQLDatabaseResponse], errResp azfake.ErrorResponder)
}

// NewItemsServerTransport creates a new instance of ItemsServerTransport with the provided implementation.
// The returned ItemsServerTransport instance is connected to an instance of sqldatabase.ItemsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewItemsServerTransport(srv *ItemsServer) *ItemsServerTransport {
	return &ItemsServerTransport{
		srv:                      srv,
		beginCreateSQLDatabase:   newTracker[azfake.PollerResponder[sqldatabase.ItemsClientCreateSQLDatabaseResponse]](),
		newListSQLDatabasesPager: newTracker[azfake.PagerResponder[sqldatabase.ItemsClientListSQLDatabasesResponse]](),
	}
}

// ItemsServerTransport connects instances of sqldatabase.ItemsClient to instances of ItemsServer.
// Don't use this type directly, use NewItemsServerTransport instead.
type ItemsServerTransport struct {
	srv                      *ItemsServer
	beginCreateSQLDatabase   *tracker[azfake.PollerResponder[sqldatabase.ItemsClientCreateSQLDatabaseResponse]]
	newListSQLDatabasesPager *tracker[azfake.PagerResponder[sqldatabase.ItemsClientListSQLDatabasesResponse]]
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
			case "ItemsClient.BeginCreateSQLDatabase":
				res.resp, res.err = i.dispatchBeginCreateSQLDatabase(req)
			case "ItemsClient.DeleteSQLDatabase":
				res.resp, res.err = i.dispatchDeleteSQLDatabase(req)
			case "ItemsClient.GetSQLDatabase":
				res.resp, res.err = i.dispatchGetSQLDatabase(req)
			case "ItemsClient.NewListSQLDatabasesPager":
				res.resp, res.err = i.dispatchNewListSQLDatabasesPager(req)
			case "ItemsClient.UpdateSQLDatabase":
				res.resp, res.err = i.dispatchUpdateSQLDatabase(req)
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

func (i *ItemsServerTransport) dispatchBeginCreateSQLDatabase(req *http.Request) (*http.Response, error) {
	if i.srv.BeginCreateSQLDatabase == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateSQLDatabase not implemented")}
	}
	beginCreateSQLDatabase := i.beginCreateSQLDatabase.get(req)
	if beginCreateSQLDatabase == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlDatabases`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[sqldatabase.CreateSQLDatabaseRequest](req)
		if err != nil {
			return nil, err
		}
		workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginCreateSQLDatabase(req.Context(), workspaceIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateSQLDatabase = &respr
		i.beginCreateSQLDatabase.add(req, beginCreateSQLDatabase)
	}

	resp, err := server.PollerResponderNext(beginCreateSQLDatabase, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		i.beginCreateSQLDatabase.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateSQLDatabase) {
		i.beginCreateSQLDatabase.remove(req)
	}

	return resp, nil
}

func (i *ItemsServerTransport) dispatchDeleteSQLDatabase(req *http.Request) (*http.Response, error) {
	if i.srv.DeleteSQLDatabase == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteSQLDatabase not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlDatabases/(?P<SQLDatabaseId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	sqlDatabaseIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("SQLDatabaseId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.DeleteSQLDatabase(req.Context(), workspaceIDParam, sqlDatabaseIDParam, nil)
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

func (i *ItemsServerTransport) dispatchGetSQLDatabase(req *http.Request) (*http.Response, error) {
	if i.srv.GetSQLDatabase == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetSQLDatabase not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlDatabases/(?P<SQLDatabaseId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	sqlDatabaseIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("SQLDatabaseId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.GetSQLDatabase(req.Context(), workspaceIDParam, sqlDatabaseIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SQLDatabase, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchNewListSQLDatabasesPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListSQLDatabasesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListSQLDatabasesPager not implemented")}
	}
	newListSQLDatabasesPager := i.newListSQLDatabasesPager.get(req)
	if newListSQLDatabasesPager == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlDatabases`
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
		var options *sqldatabase.ItemsClientListSQLDatabasesOptions
		if continuationTokenParam != nil {
			options = &sqldatabase.ItemsClientListSQLDatabasesOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := i.srv.NewListSQLDatabasesPager(workspaceIDParam, options)
		newListSQLDatabasesPager = &resp
		i.newListSQLDatabasesPager.add(req, newListSQLDatabasesPager)
		server.PagerResponderInjectNextLinks(newListSQLDatabasesPager, req, func(page *sqldatabase.ItemsClientListSQLDatabasesResponse, createLink func() string) {
			page.ContinuationURI = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListSQLDatabasesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListSQLDatabasesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListSQLDatabasesPager) {
		i.newListSQLDatabasesPager.remove(req)
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchUpdateSQLDatabase(req *http.Request) (*http.Response, error) {
	if i.srv.UpdateSQLDatabase == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateSQLDatabase not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlDatabases/(?P<SQLDatabaseId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[sqldatabase.UpdateSQLDatabaseRequest](req)
	if err != nil {
		return nil, err
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	sqlDatabaseIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("SQLDatabaseId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.UpdateSQLDatabase(req.Context(), workspaceIDParam, sqlDatabaseIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SQLDatabase, req)
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
