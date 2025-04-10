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

	"github.com/microsoft/fabric-sdk-go/fabric/notebook"
)

// LivySessionsServer is a fake server for instances of the notebook.LivySessionsClient type.
type LivySessionsServer struct {
	// GetLivySession is the fake for method LivySessionsClient.GetLivySession
	// HTTP status codes to indicate success: http.StatusOK
	GetLivySession func(ctx context.Context, workspaceID string, notebookID string, livyID string, options *notebook.LivySessionsClientGetLivySessionOptions) (resp azfake.Responder[notebook.LivySessionsClientGetLivySessionResponse], errResp azfake.ErrorResponder)

	// NewListLivySessionsPager is the fake for method LivySessionsClient.NewListLivySessionsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListLivySessionsPager func(workspaceID string, notebookID string, options *notebook.LivySessionsClientListLivySessionsOptions) (resp azfake.PagerResponder[notebook.LivySessionsClientListLivySessionsResponse])
}

// NewLivySessionsServerTransport creates a new instance of LivySessionsServerTransport with the provided implementation.
// The returned LivySessionsServerTransport instance is connected to an instance of notebook.LivySessionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLivySessionsServerTransport(srv *LivySessionsServer) *LivySessionsServerTransport {
	return &LivySessionsServerTransport{
		srv:                      srv,
		newListLivySessionsPager: newTracker[azfake.PagerResponder[notebook.LivySessionsClientListLivySessionsResponse]](),
	}
}

// LivySessionsServerTransport connects instances of notebook.LivySessionsClient to instances of LivySessionsServer.
// Don't use this type directly, use NewLivySessionsServerTransport instead.
type LivySessionsServerTransport struct {
	srv                      *LivySessionsServer
	newListLivySessionsPager *tracker[azfake.PagerResponder[notebook.LivySessionsClientListLivySessionsResponse]]
}

// Do implements the policy.Transporter interface for LivySessionsServerTransport.
func (l *LivySessionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	parts := strings.Split(method, ".")
	method = parts[1] + "." + parts[2]
	return l.dispatchToMethodFake(req, method)
}

func (l *LivySessionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if livySessionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = livySessionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "LivySessionsClient.GetLivySession":
				res.resp, res.err = l.dispatchGetLivySession(req)
			case "LivySessionsClient.NewListLivySessionsPager":
				res.resp, res.err = l.dispatchNewListLivySessionsPager(req)
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

func (l *LivySessionsServerTransport) dispatchGetLivySession(req *http.Request) (*http.Response, error) {
	if l.srv.GetLivySession == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetLivySession not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notebooks/(?P<notebookId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/livySessions/(?P<livyId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	notebookIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("notebookId")])
	if err != nil {
		return nil, err
	}
	livyIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("livyId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.GetLivySession(req.Context(), workspaceIDParam, notebookIDParam, livyIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LivySession, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LivySessionsServerTransport) dispatchNewListLivySessionsPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListLivySessionsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListLivySessionsPager not implemented")}
	}
	newListLivySessionsPager := l.newListLivySessionsPager.get(req)
	if newListLivySessionsPager == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notebooks/(?P<notebookId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/livySessions`
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
		notebookIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("notebookId")])
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
		var options *notebook.LivySessionsClientListLivySessionsOptions
		if maxResultsParam != nil || continuationTokenParam != nil {
			options = &notebook.LivySessionsClientListLivySessionsOptions{
				MaxResults:        maxResultsParam,
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := l.srv.NewListLivySessionsPager(workspaceIDParam, notebookIDParam, options)
		newListLivySessionsPager = &resp
		l.newListLivySessionsPager.add(req, newListLivySessionsPager)
		server.PagerResponderInjectNextLinks(newListLivySessionsPager, req, func(page *notebook.LivySessionsClientListLivySessionsResponse, createLink func() string) {
			page.ContinuationURI = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListLivySessionsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListLivySessionsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListLivySessionsPager) {
		l.newListLivySessionsPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to LivySessionsServerTransport
var livySessionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
