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

	"github.com/microsoft/fabric-sdk-go/fabric/kqldashboard"
)

// ItemsServer is a fake server for instances of the kqldashboard.ItemsClient type.
type ItemsServer struct {
	// BeginCreateKQLDashboard is the fake for method ItemsClient.BeginCreateKQLDashboard
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateKQLDashboard func(ctx context.Context, workspaceID string, createKQLDashboardRequest kqldashboard.CreateKQLDashboardRequest, options *kqldashboard.ItemsClientBeginCreateKQLDashboardOptions) (resp azfake.PollerResponder[kqldashboard.ItemsClientCreateKQLDashboardResponse], errResp azfake.ErrorResponder)

	// DeleteKQLDashboard is the fake for method ItemsClient.DeleteKQLDashboard
	// HTTP status codes to indicate success: http.StatusOK
	DeleteKQLDashboard func(ctx context.Context, workspaceID string, kqlDashboardID string, options *kqldashboard.ItemsClientDeleteKQLDashboardOptions) (resp azfake.Responder[kqldashboard.ItemsClientDeleteKQLDashboardResponse], errResp azfake.ErrorResponder)

	// GetKQLDashboard is the fake for method ItemsClient.GetKQLDashboard
	// HTTP status codes to indicate success: http.StatusOK
	GetKQLDashboard func(ctx context.Context, workspaceID string, kqlDashboardID string, options *kqldashboard.ItemsClientGetKQLDashboardOptions) (resp azfake.Responder[kqldashboard.ItemsClientGetKQLDashboardResponse], errResp azfake.ErrorResponder)

	// BeginGetKQLDashboardDefinition is the fake for method ItemsClient.BeginGetKQLDashboardDefinition
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGetKQLDashboardDefinition func(ctx context.Context, workspaceID string, kqlDashboardID string, options *kqldashboard.ItemsClientBeginGetKQLDashboardDefinitionOptions) (resp azfake.PollerResponder[kqldashboard.ItemsClientGetKQLDashboardDefinitionResponse], errResp azfake.ErrorResponder)

	// NewListKQLDashboardsPager is the fake for method ItemsClient.NewListKQLDashboardsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListKQLDashboardsPager func(workspaceID string, options *kqldashboard.ItemsClientListKQLDashboardsOptions) (resp azfake.PagerResponder[kqldashboard.ItemsClientListKQLDashboardsResponse])

	// UpdateKQLDashboard is the fake for method ItemsClient.UpdateKQLDashboard
	// HTTP status codes to indicate success: http.StatusOK
	UpdateKQLDashboard func(ctx context.Context, workspaceID string, kqlDashboardID string, updateKQLDashboardRequest kqldashboard.UpdateKQLDashboardRequest, options *kqldashboard.ItemsClientUpdateKQLDashboardOptions) (resp azfake.Responder[kqldashboard.ItemsClientUpdateKQLDashboardResponse], errResp azfake.ErrorResponder)

	// BeginUpdateKQLDashboardDefinition is the fake for method ItemsClient.BeginUpdateKQLDashboardDefinition
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginUpdateKQLDashboardDefinition func(ctx context.Context, workspaceID string, kqlDashboardID string, updateKQLDashboardDefinitionRequest kqldashboard.UpdateKQLDashboardDefinitionRequest, options *kqldashboard.ItemsClientBeginUpdateKQLDashboardDefinitionOptions) (resp azfake.PollerResponder[kqldashboard.ItemsClientUpdateKQLDashboardDefinitionResponse], errResp azfake.ErrorResponder)
}

// NewItemsServerTransport creates a new instance of ItemsServerTransport with the provided implementation.
// The returned ItemsServerTransport instance is connected to an instance of kqldashboard.ItemsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewItemsServerTransport(srv *ItemsServer) *ItemsServerTransport {
	return &ItemsServerTransport{
		srv:                               srv,
		beginCreateKQLDashboard:           newTracker[azfake.PollerResponder[kqldashboard.ItemsClientCreateKQLDashboardResponse]](),
		beginGetKQLDashboardDefinition:    newTracker[azfake.PollerResponder[kqldashboard.ItemsClientGetKQLDashboardDefinitionResponse]](),
		newListKQLDashboardsPager:         newTracker[azfake.PagerResponder[kqldashboard.ItemsClientListKQLDashboardsResponse]](),
		beginUpdateKQLDashboardDefinition: newTracker[azfake.PollerResponder[kqldashboard.ItemsClientUpdateKQLDashboardDefinitionResponse]](),
	}
}

// ItemsServerTransport connects instances of kqldashboard.ItemsClient to instances of ItemsServer.
// Don't use this type directly, use NewItemsServerTransport instead.
type ItemsServerTransport struct {
	srv                               *ItemsServer
	beginCreateKQLDashboard           *tracker[azfake.PollerResponder[kqldashboard.ItemsClientCreateKQLDashboardResponse]]
	beginGetKQLDashboardDefinition    *tracker[azfake.PollerResponder[kqldashboard.ItemsClientGetKQLDashboardDefinitionResponse]]
	newListKQLDashboardsPager         *tracker[azfake.PagerResponder[kqldashboard.ItemsClientListKQLDashboardsResponse]]
	beginUpdateKQLDashboardDefinition *tracker[azfake.PollerResponder[kqldashboard.ItemsClientUpdateKQLDashboardDefinitionResponse]]
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
			case "ItemsClient.BeginCreateKQLDashboard":
				res.resp, res.err = i.dispatchBeginCreateKQLDashboard(req)
			case "ItemsClient.DeleteKQLDashboard":
				res.resp, res.err = i.dispatchDeleteKQLDashboard(req)
			case "ItemsClient.GetKQLDashboard":
				res.resp, res.err = i.dispatchGetKQLDashboard(req)
			case "ItemsClient.BeginGetKQLDashboardDefinition":
				res.resp, res.err = i.dispatchBeginGetKQLDashboardDefinition(req)
			case "ItemsClient.NewListKQLDashboardsPager":
				res.resp, res.err = i.dispatchNewListKQLDashboardsPager(req)
			case "ItemsClient.UpdateKQLDashboard":
				res.resp, res.err = i.dispatchUpdateKQLDashboard(req)
			case "ItemsClient.BeginUpdateKQLDashboardDefinition":
				res.resp, res.err = i.dispatchBeginUpdateKQLDashboardDefinition(req)
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

func (i *ItemsServerTransport) dispatchBeginCreateKQLDashboard(req *http.Request) (*http.Response, error) {
	if i.srv.BeginCreateKQLDashboard == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateKQLDashboard not implemented")}
	}
	beginCreateKQLDashboard := i.beginCreateKQLDashboard.get(req)
	if beginCreateKQLDashboard == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/kqlDashboards`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[kqldashboard.CreateKQLDashboardRequest](req)
		if err != nil {
			return nil, err
		}
		workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginCreateKQLDashboard(req.Context(), workspaceIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateKQLDashboard = &respr
		i.beginCreateKQLDashboard.add(req, beginCreateKQLDashboard)
	}

	resp, err := server.PollerResponderNext(beginCreateKQLDashboard, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		i.beginCreateKQLDashboard.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateKQLDashboard) {
		i.beginCreateKQLDashboard.remove(req)
	}

	return resp, nil
}

func (i *ItemsServerTransport) dispatchDeleteKQLDashboard(req *http.Request) (*http.Response, error) {
	if i.srv.DeleteKQLDashboard == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteKQLDashboard not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/kqlDashboards/(?P<kqlDashboardId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	kqlDashboardIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("kqlDashboardId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.DeleteKQLDashboard(req.Context(), workspaceIDParam, kqlDashboardIDParam, nil)
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

func (i *ItemsServerTransport) dispatchGetKQLDashboard(req *http.Request) (*http.Response, error) {
	if i.srv.GetKQLDashboard == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetKQLDashboard not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/kqlDashboards/(?P<kqlDashboardId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	kqlDashboardIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("kqlDashboardId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.GetKQLDashboard(req.Context(), workspaceIDParam, kqlDashboardIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).KQLDashboard, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchBeginGetKQLDashboardDefinition(req *http.Request) (*http.Response, error) {
	if i.srv.BeginGetKQLDashboardDefinition == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGetKQLDashboardDefinition not implemented")}
	}
	beginGetKQLDashboardDefinition := i.beginGetKQLDashboardDefinition.get(req)
	if beginGetKQLDashboardDefinition == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/kqlDashboards/(?P<kqlDashboardId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getDefinition`
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
		kqlDashboardIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("kqlDashboardId")])
		if err != nil {
			return nil, err
		}
		formatUnescaped, err := url.QueryUnescape(qp.Get("format"))
		if err != nil {
			return nil, err
		}
		formatParam := getOptional(formatUnescaped)
		var options *kqldashboard.ItemsClientBeginGetKQLDashboardDefinitionOptions
		if formatParam != nil {
			options = &kqldashboard.ItemsClientBeginGetKQLDashboardDefinitionOptions{
				Format: formatParam,
			}
		}
		respr, errRespr := i.srv.BeginGetKQLDashboardDefinition(req.Context(), workspaceIDParam, kqlDashboardIDParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGetKQLDashboardDefinition = &respr
		i.beginGetKQLDashboardDefinition.add(req, beginGetKQLDashboardDefinition)
	}

	resp, err := server.PollerResponderNext(beginGetKQLDashboardDefinition, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		i.beginGetKQLDashboardDefinition.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGetKQLDashboardDefinition) {
		i.beginGetKQLDashboardDefinition.remove(req)
	}

	return resp, nil
}

func (i *ItemsServerTransport) dispatchNewListKQLDashboardsPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListKQLDashboardsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListKQLDashboardsPager not implemented")}
	}
	newListKQLDashboardsPager := i.newListKQLDashboardsPager.get(req)
	if newListKQLDashboardsPager == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/kqlDashboards`
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
		var options *kqldashboard.ItemsClientListKQLDashboardsOptions
		if continuationTokenParam != nil {
			options = &kqldashboard.ItemsClientListKQLDashboardsOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := i.srv.NewListKQLDashboardsPager(workspaceIDParam, options)
		newListKQLDashboardsPager = &resp
		i.newListKQLDashboardsPager.add(req, newListKQLDashboardsPager)
		server.PagerResponderInjectNextLinks(newListKQLDashboardsPager, req, func(page *kqldashboard.ItemsClientListKQLDashboardsResponse, createLink func() string) {
			page.ContinuationURI = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListKQLDashboardsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListKQLDashboardsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListKQLDashboardsPager) {
		i.newListKQLDashboardsPager.remove(req)
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchUpdateKQLDashboard(req *http.Request) (*http.Response, error) {
	if i.srv.UpdateKQLDashboard == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateKQLDashboard not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/kqlDashboards/(?P<kqlDashboardId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[kqldashboard.UpdateKQLDashboardRequest](req)
	if err != nil {
		return nil, err
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	kqlDashboardIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("kqlDashboardId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.UpdateKQLDashboard(req.Context(), workspaceIDParam, kqlDashboardIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).KQLDashboard, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchBeginUpdateKQLDashboardDefinition(req *http.Request) (*http.Response, error) {
	if i.srv.BeginUpdateKQLDashboardDefinition == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateKQLDashboardDefinition not implemented")}
	}
	beginUpdateKQLDashboardDefinition := i.beginUpdateKQLDashboardDefinition.get(req)
	if beginUpdateKQLDashboardDefinition == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/kqlDashboards/(?P<kqlDashboardId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateDefinition`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		body, err := server.UnmarshalRequestAsJSON[kqldashboard.UpdateKQLDashboardDefinitionRequest](req)
		if err != nil {
			return nil, err
		}
		workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
		if err != nil {
			return nil, err
		}
		kqlDashboardIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("kqlDashboardId")])
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
		var options *kqldashboard.ItemsClientBeginUpdateKQLDashboardDefinitionOptions
		if updateMetadataParam != nil {
			options = &kqldashboard.ItemsClientBeginUpdateKQLDashboardDefinitionOptions{
				UpdateMetadata: updateMetadataParam,
			}
		}
		respr, errRespr := i.srv.BeginUpdateKQLDashboardDefinition(req.Context(), workspaceIDParam, kqlDashboardIDParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateKQLDashboardDefinition = &respr
		i.beginUpdateKQLDashboardDefinition.add(req, beginUpdateKQLDashboardDefinition)
	}

	resp, err := server.PollerResponderNext(beginUpdateKQLDashboardDefinition, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		i.beginUpdateKQLDashboardDefinition.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateKQLDashboardDefinition) {
		i.beginUpdateKQLDashboardDefinition.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ItemsServerTransport
var itemsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
