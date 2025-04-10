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

	"github.com/microsoft/fabric-sdk-go/fabric/lakehouse"
)

// TablesServer is a fake server for instances of the lakehouse.TablesClient type.
type TablesServer struct {
	// NewListTablesPager is the fake for method TablesClient.NewListTablesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListTablesPager func(workspaceID string, lakehouseID string, options *lakehouse.TablesClientListTablesOptions) (resp azfake.PagerResponder[lakehouse.TablesClientListTablesResponse])

	// BeginLoadTable is the fake for method TablesClient.BeginLoadTable
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginLoadTable func(ctx context.Context, workspaceID string, lakehouseID string, tableName string, loadTableRequest lakehouse.LoadTableRequest, options *lakehouse.TablesClientBeginLoadTableOptions) (resp azfake.PollerResponder[lakehouse.TablesClientLoadTableResponse], errResp azfake.ErrorResponder)
}

// NewTablesServerTransport creates a new instance of TablesServerTransport with the provided implementation.
// The returned TablesServerTransport instance is connected to an instance of lakehouse.TablesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTablesServerTransport(srv *TablesServer) *TablesServerTransport {
	return &TablesServerTransport{
		srv:                srv,
		newListTablesPager: newTracker[azfake.PagerResponder[lakehouse.TablesClientListTablesResponse]](),
		beginLoadTable:     newTracker[azfake.PollerResponder[lakehouse.TablesClientLoadTableResponse]](),
	}
}

// TablesServerTransport connects instances of lakehouse.TablesClient to instances of TablesServer.
// Don't use this type directly, use NewTablesServerTransport instead.
type TablesServerTransport struct {
	srv                *TablesServer
	newListTablesPager *tracker[azfake.PagerResponder[lakehouse.TablesClientListTablesResponse]]
	beginLoadTable     *tracker[azfake.PollerResponder[lakehouse.TablesClientLoadTableResponse]]
}

// Do implements the policy.Transporter interface for TablesServerTransport.
func (t *TablesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	parts := strings.Split(method, ".")
	method = parts[1] + "." + parts[2]
	return t.dispatchToMethodFake(req, method)
}

func (t *TablesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if tablesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = tablesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "TablesClient.NewListTablesPager":
				res.resp, res.err = t.dispatchNewListTablesPager(req)
			case "TablesClient.BeginLoadTable":
				res.resp, res.err = t.dispatchBeginLoadTable(req)
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

func (t *TablesServerTransport) dispatchNewListTablesPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListTablesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListTablesPager not implemented")}
	}
	newListTablesPager := t.newListTablesPager.get(req)
	if newListTablesPager == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/lakehouses/(?P<lakehouseId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tables`
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
		lakehouseIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("lakehouseId")])
		if err != nil {
			return nil, err
		}
		maxResultsUnescaped, err := url.QueryUnescape(qp.Get("maxResults"))
		if err != nil {
			return nil, err
		}
		maxResultsParam, err := parseOptional(maxResultsUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		continuationTokenUnescaped, err := url.QueryUnescape(qp.Get("continuationToken"))
		if err != nil {
			return nil, err
		}
		continuationTokenParam := getOptional(continuationTokenUnescaped)
		var options *lakehouse.TablesClientListTablesOptions
		if maxResultsParam != nil || continuationTokenParam != nil {
			options = &lakehouse.TablesClientListTablesOptions{
				MaxResults:        maxResultsParam,
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := t.srv.NewListTablesPager(workspaceIDParam, lakehouseIDParam, options)
		newListTablesPager = &resp
		t.newListTablesPager.add(req, newListTablesPager)
		server.PagerResponderInjectNextLinks(newListTablesPager, req, func(page *lakehouse.TablesClientListTablesResponse, createLink func() string) {
			page.ContinuationURI = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListTablesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListTablesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListTablesPager) {
		t.newListTablesPager.remove(req)
	}
	return resp, nil
}

func (t *TablesServerTransport) dispatchBeginLoadTable(req *http.Request) (*http.Response, error) {
	if t.srv.BeginLoadTable == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginLoadTable not implemented")}
	}
	beginLoadTable := t.beginLoadTable.get(req)
	if beginLoadTable == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/lakehouses/(?P<lakehouseId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tables/(?P<tableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/load`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[lakehouse.LoadTableRequest](req)
		if err != nil {
			return nil, err
		}
		workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
		if err != nil {
			return nil, err
		}
		lakehouseIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("lakehouseId")])
		if err != nil {
			return nil, err
		}
		tableNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("tableName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := t.srv.BeginLoadTable(req.Context(), workspaceIDParam, lakehouseIDParam, tableNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginLoadTable = &respr
		t.beginLoadTable.add(req, beginLoadTable)
	}

	resp, err := server.PollerResponderNext(beginLoadTable, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		t.beginLoadTable.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginLoadTable) {
		t.beginLoadTable.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to TablesServerTransport
var tablesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
