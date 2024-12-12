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

	"github.com/microsoft/fabric-sdk-go/fabric/reflex"
)

// ItemsServer is a fake server for instances of the reflex.ItemsClient type.
type ItemsServer struct {
	// BeginCreateReflex is the fake for method ItemsClient.BeginCreateReflex
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateReflex func(ctx context.Context, workspaceID string, createReflexRequest reflex.CreateReflexRequest, options *reflex.ItemsClientBeginCreateReflexOptions) (resp azfake.PollerResponder[reflex.ItemsClientCreateReflexResponse], errResp azfake.ErrorResponder)

	// DeleteReflex is the fake for method ItemsClient.DeleteReflex
	// HTTP status codes to indicate success: http.StatusOK
	DeleteReflex func(ctx context.Context, workspaceID string, reflexID string, options *reflex.ItemsClientDeleteReflexOptions) (resp azfake.Responder[reflex.ItemsClientDeleteReflexResponse], errResp azfake.ErrorResponder)

	// GetReflex is the fake for method ItemsClient.GetReflex
	// HTTP status codes to indicate success: http.StatusOK
	GetReflex func(ctx context.Context, workspaceID string, reflexID string, options *reflex.ItemsClientGetReflexOptions) (resp azfake.Responder[reflex.ItemsClientGetReflexResponse], errResp azfake.ErrorResponder)

	// BeginGetReflexDefinition is the fake for method ItemsClient.BeginGetReflexDefinition
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGetReflexDefinition func(ctx context.Context, workspaceID string, reflexID string, options *reflex.ItemsClientBeginGetReflexDefinitionOptions) (resp azfake.PollerResponder[reflex.ItemsClientGetReflexDefinitionResponse], errResp azfake.ErrorResponder)

	// NewListReflexesPager is the fake for method ItemsClient.NewListReflexesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListReflexesPager func(workspaceID string, options *reflex.ItemsClientListReflexesOptions) (resp azfake.PagerResponder[reflex.ItemsClientListReflexesResponse])

	// UpdateReflex is the fake for method ItemsClient.UpdateReflex
	// HTTP status codes to indicate success: http.StatusOK
	UpdateReflex func(ctx context.Context, workspaceID string, reflexID string, updateReflexRequest reflex.UpdateReflexRequest, options *reflex.ItemsClientUpdateReflexOptions) (resp azfake.Responder[reflex.ItemsClientUpdateReflexResponse], errResp azfake.ErrorResponder)

	// BeginUpdateReflexDefinition is the fake for method ItemsClient.BeginUpdateReflexDefinition
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginUpdateReflexDefinition func(ctx context.Context, workspaceID string, reflexID string, updateReflexDefinitionRequest reflex.UpdateReflexDefinitionRequest, options *reflex.ItemsClientBeginUpdateReflexDefinitionOptions) (resp azfake.PollerResponder[reflex.ItemsClientUpdateReflexDefinitionResponse], errResp azfake.ErrorResponder)
}

// NewItemsServerTransport creates a new instance of ItemsServerTransport with the provided implementation.
// The returned ItemsServerTransport instance is connected to an instance of reflex.ItemsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewItemsServerTransport(srv *ItemsServer) *ItemsServerTransport {
	return &ItemsServerTransport{
		srv:                         srv,
		beginCreateReflex:           newTracker[azfake.PollerResponder[reflex.ItemsClientCreateReflexResponse]](),
		beginGetReflexDefinition:    newTracker[azfake.PollerResponder[reflex.ItemsClientGetReflexDefinitionResponse]](),
		newListReflexesPager:        newTracker[azfake.PagerResponder[reflex.ItemsClientListReflexesResponse]](),
		beginUpdateReflexDefinition: newTracker[azfake.PollerResponder[reflex.ItemsClientUpdateReflexDefinitionResponse]](),
	}
}

// ItemsServerTransport connects instances of reflex.ItemsClient to instances of ItemsServer.
// Don't use this type directly, use NewItemsServerTransport instead.
type ItemsServerTransport struct {
	srv                         *ItemsServer
	beginCreateReflex           *tracker[azfake.PollerResponder[reflex.ItemsClientCreateReflexResponse]]
	beginGetReflexDefinition    *tracker[azfake.PollerResponder[reflex.ItemsClientGetReflexDefinitionResponse]]
	newListReflexesPager        *tracker[azfake.PagerResponder[reflex.ItemsClientListReflexesResponse]]
	beginUpdateReflexDefinition *tracker[azfake.PollerResponder[reflex.ItemsClientUpdateReflexDefinitionResponse]]
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
			case "ItemsClient.BeginCreateReflex":
				res.resp, res.err = i.dispatchBeginCreateReflex(req)
			case "ItemsClient.DeleteReflex":
				res.resp, res.err = i.dispatchDeleteReflex(req)
			case "ItemsClient.GetReflex":
				res.resp, res.err = i.dispatchGetReflex(req)
			case "ItemsClient.BeginGetReflexDefinition":
				res.resp, res.err = i.dispatchBeginGetReflexDefinition(req)
			case "ItemsClient.NewListReflexesPager":
				res.resp, res.err = i.dispatchNewListReflexesPager(req)
			case "ItemsClient.UpdateReflex":
				res.resp, res.err = i.dispatchUpdateReflex(req)
			case "ItemsClient.BeginUpdateReflexDefinition":
				res.resp, res.err = i.dispatchBeginUpdateReflexDefinition(req)
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

func (i *ItemsServerTransport) dispatchBeginCreateReflex(req *http.Request) (*http.Response, error) {
	if i.srv.BeginCreateReflex == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateReflex not implemented")}
	}
	beginCreateReflex := i.beginCreateReflex.get(req)
	if beginCreateReflex == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reflexes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[reflex.CreateReflexRequest](req)
		if err != nil {
			return nil, err
		}
		workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginCreateReflex(req.Context(), workspaceIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateReflex = &respr
		i.beginCreateReflex.add(req, beginCreateReflex)
	}

	resp, err := server.PollerResponderNext(beginCreateReflex, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		i.beginCreateReflex.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateReflex) {
		i.beginCreateReflex.remove(req)
	}

	return resp, nil
}

func (i *ItemsServerTransport) dispatchDeleteReflex(req *http.Request) (*http.Response, error) {
	if i.srv.DeleteReflex == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteReflex not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reflexes/(?P<reflexId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	reflexIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("reflexId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.DeleteReflex(req.Context(), workspaceIDParam, reflexIDParam, nil)
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

func (i *ItemsServerTransport) dispatchGetReflex(req *http.Request) (*http.Response, error) {
	if i.srv.GetReflex == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetReflex not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reflexes/(?P<reflexId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	reflexIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("reflexId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.GetReflex(req.Context(), workspaceIDParam, reflexIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Reflex, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchBeginGetReflexDefinition(req *http.Request) (*http.Response, error) {
	if i.srv.BeginGetReflexDefinition == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGetReflexDefinition not implemented")}
	}
	beginGetReflexDefinition := i.beginGetReflexDefinition.get(req)
	if beginGetReflexDefinition == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reflexes/(?P<reflexId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getDefinition`
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
		reflexIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("reflexId")])
		if err != nil {
			return nil, err
		}
		formatUnescaped, err := url.QueryUnescape(qp.Get("format"))
		if err != nil {
			return nil, err
		}
		formatParam := getOptional(formatUnescaped)
		var options *reflex.ItemsClientBeginGetReflexDefinitionOptions
		if formatParam != nil {
			options = &reflex.ItemsClientBeginGetReflexDefinitionOptions{
				Format: formatParam,
			}
		}
		respr, errRespr := i.srv.BeginGetReflexDefinition(req.Context(), workspaceIDParam, reflexIDParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGetReflexDefinition = &respr
		i.beginGetReflexDefinition.add(req, beginGetReflexDefinition)
	}

	resp, err := server.PollerResponderNext(beginGetReflexDefinition, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		i.beginGetReflexDefinition.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGetReflexDefinition) {
		i.beginGetReflexDefinition.remove(req)
	}

	return resp, nil
}

func (i *ItemsServerTransport) dispatchNewListReflexesPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListReflexesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListReflexesPager not implemented")}
	}
	newListReflexesPager := i.newListReflexesPager.get(req)
	if newListReflexesPager == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reflexes`
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
		var options *reflex.ItemsClientListReflexesOptions
		if continuationTokenParam != nil {
			options = &reflex.ItemsClientListReflexesOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := i.srv.NewListReflexesPager(workspaceIDParam, options)
		newListReflexesPager = &resp
		i.newListReflexesPager.add(req, newListReflexesPager)
		server.PagerResponderInjectNextLinks(newListReflexesPager, req, func(page *reflex.ItemsClientListReflexesResponse, createLink func() string) {
			page.ContinuationURI = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListReflexesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListReflexesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListReflexesPager) {
		i.newListReflexesPager.remove(req)
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchUpdateReflex(req *http.Request) (*http.Response, error) {
	if i.srv.UpdateReflex == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateReflex not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reflexes/(?P<reflexId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[reflex.UpdateReflexRequest](req)
	if err != nil {
		return nil, err
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	reflexIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("reflexId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.UpdateReflex(req.Context(), workspaceIDParam, reflexIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Reflex, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchBeginUpdateReflexDefinition(req *http.Request) (*http.Response, error) {
	if i.srv.BeginUpdateReflexDefinition == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateReflexDefinition not implemented")}
	}
	beginUpdateReflexDefinition := i.beginUpdateReflexDefinition.get(req)
	if beginUpdateReflexDefinition == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reflexes/(?P<reflexId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateDefinition`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		body, err := server.UnmarshalRequestAsJSON[reflex.UpdateReflexDefinitionRequest](req)
		if err != nil {
			return nil, err
		}
		workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
		if err != nil {
			return nil, err
		}
		reflexIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("reflexId")])
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
		var options *reflex.ItemsClientBeginUpdateReflexDefinitionOptions
		if updateMetadataParam != nil {
			options = &reflex.ItemsClientBeginUpdateReflexDefinitionOptions{
				UpdateMetadata: updateMetadataParam,
			}
		}
		respr, errRespr := i.srv.BeginUpdateReflexDefinition(req.Context(), workspaceIDParam, reflexIDParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateReflexDefinition = &respr
		i.beginUpdateReflexDefinition.add(req, beginUpdateReflexDefinition)
	}

	resp, err := server.PollerResponderNext(beginUpdateReflexDefinition, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		i.beginUpdateReflexDefinition.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateReflexDefinition) {
		i.beginUpdateReflexDefinition.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ItemsServerTransport
var itemsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}