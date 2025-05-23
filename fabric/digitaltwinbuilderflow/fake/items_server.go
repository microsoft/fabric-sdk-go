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

	"github.com/microsoft/fabric-sdk-go/fabric/digitaltwinbuilderflow"
)

// ItemsServer is a fake server for instances of the digitaltwinbuilderflow.ItemsClient type.
type ItemsServer struct {
	// BeginCreateDigitalTwinBuilderFlow is the fake for method ItemsClient.BeginCreateDigitalTwinBuilderFlow
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateDigitalTwinBuilderFlow func(ctx context.Context, workspaceID string, createDigitalTwinBuilderFlowRequest digitaltwinbuilderflow.CreateDigitalTwinBuilderFlowRequest, options *digitaltwinbuilderflow.ItemsClientBeginCreateDigitalTwinBuilderFlowOptions) (resp azfake.PollerResponder[digitaltwinbuilderflow.ItemsClientCreateDigitalTwinBuilderFlowResponse], errResp azfake.ErrorResponder)

	// DeleteDigitalTwinBuilderFlow is the fake for method ItemsClient.DeleteDigitalTwinBuilderFlow
	// HTTP status codes to indicate success: http.StatusOK
	DeleteDigitalTwinBuilderFlow func(ctx context.Context, workspaceID string, digitalTwinBuilderFlowID string, options *digitaltwinbuilderflow.ItemsClientDeleteDigitalTwinBuilderFlowOptions) (resp azfake.Responder[digitaltwinbuilderflow.ItemsClientDeleteDigitalTwinBuilderFlowResponse], errResp azfake.ErrorResponder)

	// GetDigitalTwinBuilderFlow is the fake for method ItemsClient.GetDigitalTwinBuilderFlow
	// HTTP status codes to indicate success: http.StatusOK
	GetDigitalTwinBuilderFlow func(ctx context.Context, workspaceID string, digitalTwinBuilderFlowID string, options *digitaltwinbuilderflow.ItemsClientGetDigitalTwinBuilderFlowOptions) (resp azfake.Responder[digitaltwinbuilderflow.ItemsClientGetDigitalTwinBuilderFlowResponse], errResp azfake.ErrorResponder)

	// BeginGetDigitalTwinBuilderFlowDefinition is the fake for method ItemsClient.BeginGetDigitalTwinBuilderFlowDefinition
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGetDigitalTwinBuilderFlowDefinition func(ctx context.Context, workspaceID string, digitalTwinBuilderFlowID string, options *digitaltwinbuilderflow.ItemsClientBeginGetDigitalTwinBuilderFlowDefinitionOptions) (resp azfake.PollerResponder[digitaltwinbuilderflow.ItemsClientGetDigitalTwinBuilderFlowDefinitionResponse], errResp azfake.ErrorResponder)

	// NewListDigitalTwinBuilderFlowsPager is the fake for method ItemsClient.NewListDigitalTwinBuilderFlowsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListDigitalTwinBuilderFlowsPager func(workspaceID string, options *digitaltwinbuilderflow.ItemsClientListDigitalTwinBuilderFlowsOptions) (resp azfake.PagerResponder[digitaltwinbuilderflow.ItemsClientListDigitalTwinBuilderFlowsResponse])

	// UpdateDigitalTwinBuilderFlow is the fake for method ItemsClient.UpdateDigitalTwinBuilderFlow
	// HTTP status codes to indicate success: http.StatusOK
	UpdateDigitalTwinBuilderFlow func(ctx context.Context, workspaceID string, digitalTwinBuilderFlowID string, updateDigitalTwinBuilderFlowRequest digitaltwinbuilderflow.UpdateDigitalTwinBuilderFlowRequest, options *digitaltwinbuilderflow.ItemsClientUpdateDigitalTwinBuilderFlowOptions) (resp azfake.Responder[digitaltwinbuilderflow.ItemsClientUpdateDigitalTwinBuilderFlowResponse], errResp azfake.ErrorResponder)

	// BeginUpdateDigitalTwinBuilderFlowDefinition is the fake for method ItemsClient.BeginUpdateDigitalTwinBuilderFlowDefinition
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginUpdateDigitalTwinBuilderFlowDefinition func(ctx context.Context, workspaceID string, digitalTwinBuilderFlowID string, updateDigitalTwinBuilderFlowDefinitionRequest digitaltwinbuilderflow.UpdateDigitalTwinBuilderFlowDefinitionRequest, options *digitaltwinbuilderflow.ItemsClientBeginUpdateDigitalTwinBuilderFlowDefinitionOptions) (resp azfake.PollerResponder[digitaltwinbuilderflow.ItemsClientUpdateDigitalTwinBuilderFlowDefinitionResponse], errResp azfake.ErrorResponder)
}

// NewItemsServerTransport creates a new instance of ItemsServerTransport with the provided implementation.
// The returned ItemsServerTransport instance is connected to an instance of digitaltwinbuilderflow.ItemsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewItemsServerTransport(srv *ItemsServer) *ItemsServerTransport {
	return &ItemsServerTransport{
		srv:                                         srv,
		beginCreateDigitalTwinBuilderFlow:           newTracker[azfake.PollerResponder[digitaltwinbuilderflow.ItemsClientCreateDigitalTwinBuilderFlowResponse]](),
		beginGetDigitalTwinBuilderFlowDefinition:    newTracker[azfake.PollerResponder[digitaltwinbuilderflow.ItemsClientGetDigitalTwinBuilderFlowDefinitionResponse]](),
		newListDigitalTwinBuilderFlowsPager:         newTracker[azfake.PagerResponder[digitaltwinbuilderflow.ItemsClientListDigitalTwinBuilderFlowsResponse]](),
		beginUpdateDigitalTwinBuilderFlowDefinition: newTracker[azfake.PollerResponder[digitaltwinbuilderflow.ItemsClientUpdateDigitalTwinBuilderFlowDefinitionResponse]](),
	}
}

// ItemsServerTransport connects instances of digitaltwinbuilderflow.ItemsClient to instances of ItemsServer.
// Don't use this type directly, use NewItemsServerTransport instead.
type ItemsServerTransport struct {
	srv                                         *ItemsServer
	beginCreateDigitalTwinBuilderFlow           *tracker[azfake.PollerResponder[digitaltwinbuilderflow.ItemsClientCreateDigitalTwinBuilderFlowResponse]]
	beginGetDigitalTwinBuilderFlowDefinition    *tracker[azfake.PollerResponder[digitaltwinbuilderflow.ItemsClientGetDigitalTwinBuilderFlowDefinitionResponse]]
	newListDigitalTwinBuilderFlowsPager         *tracker[azfake.PagerResponder[digitaltwinbuilderflow.ItemsClientListDigitalTwinBuilderFlowsResponse]]
	beginUpdateDigitalTwinBuilderFlowDefinition *tracker[azfake.PollerResponder[digitaltwinbuilderflow.ItemsClientUpdateDigitalTwinBuilderFlowDefinitionResponse]]
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
			case "ItemsClient.BeginCreateDigitalTwinBuilderFlow":
				res.resp, res.err = i.dispatchBeginCreateDigitalTwinBuilderFlow(req)
			case "ItemsClient.DeleteDigitalTwinBuilderFlow":
				res.resp, res.err = i.dispatchDeleteDigitalTwinBuilderFlow(req)
			case "ItemsClient.GetDigitalTwinBuilderFlow":
				res.resp, res.err = i.dispatchGetDigitalTwinBuilderFlow(req)
			case "ItemsClient.BeginGetDigitalTwinBuilderFlowDefinition":
				res.resp, res.err = i.dispatchBeginGetDigitalTwinBuilderFlowDefinition(req)
			case "ItemsClient.NewListDigitalTwinBuilderFlowsPager":
				res.resp, res.err = i.dispatchNewListDigitalTwinBuilderFlowsPager(req)
			case "ItemsClient.UpdateDigitalTwinBuilderFlow":
				res.resp, res.err = i.dispatchUpdateDigitalTwinBuilderFlow(req)
			case "ItemsClient.BeginUpdateDigitalTwinBuilderFlowDefinition":
				res.resp, res.err = i.dispatchBeginUpdateDigitalTwinBuilderFlowDefinition(req)
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

func (i *ItemsServerTransport) dispatchBeginCreateDigitalTwinBuilderFlow(req *http.Request) (*http.Response, error) {
	if i.srv.BeginCreateDigitalTwinBuilderFlow == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateDigitalTwinBuilderFlow not implemented")}
	}
	beginCreateDigitalTwinBuilderFlow := i.beginCreateDigitalTwinBuilderFlow.get(req)
	if beginCreateDigitalTwinBuilderFlow == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/DigitalTwinBuilderFlows`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[digitaltwinbuilderflow.CreateDigitalTwinBuilderFlowRequest](req)
		if err != nil {
			return nil, err
		}
		workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginCreateDigitalTwinBuilderFlow(req.Context(), workspaceIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateDigitalTwinBuilderFlow = &respr
		i.beginCreateDigitalTwinBuilderFlow.add(req, beginCreateDigitalTwinBuilderFlow)
	}

	resp, err := server.PollerResponderNext(beginCreateDigitalTwinBuilderFlow, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		i.beginCreateDigitalTwinBuilderFlow.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateDigitalTwinBuilderFlow) {
		i.beginCreateDigitalTwinBuilderFlow.remove(req)
	}

	return resp, nil
}

func (i *ItemsServerTransport) dispatchDeleteDigitalTwinBuilderFlow(req *http.Request) (*http.Response, error) {
	if i.srv.DeleteDigitalTwinBuilderFlow == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteDigitalTwinBuilderFlow not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/DigitalTwinBuilderFlows/(?P<digitalTwinBuilderFlowId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	digitalTwinBuilderFlowIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("digitalTwinBuilderFlowId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.DeleteDigitalTwinBuilderFlow(req.Context(), workspaceIDParam, digitalTwinBuilderFlowIDParam, nil)
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

func (i *ItemsServerTransport) dispatchGetDigitalTwinBuilderFlow(req *http.Request) (*http.Response, error) {
	if i.srv.GetDigitalTwinBuilderFlow == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDigitalTwinBuilderFlow not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/DigitalTwinBuilderFlows/(?P<digitalTwinBuilderFlowId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	digitalTwinBuilderFlowIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("digitalTwinBuilderFlowId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.GetDigitalTwinBuilderFlow(req.Context(), workspaceIDParam, digitalTwinBuilderFlowIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DigitalTwinBuilderFlow, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchBeginGetDigitalTwinBuilderFlowDefinition(req *http.Request) (*http.Response, error) {
	if i.srv.BeginGetDigitalTwinBuilderFlowDefinition == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGetDigitalTwinBuilderFlowDefinition not implemented")}
	}
	beginGetDigitalTwinBuilderFlowDefinition := i.beginGetDigitalTwinBuilderFlowDefinition.get(req)
	if beginGetDigitalTwinBuilderFlowDefinition == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/DigitalTwinBuilderFlows/(?P<digitalTwinBuilderFlowId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getDefinition`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
		if err != nil {
			return nil, err
		}
		digitalTwinBuilderFlowIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("digitalTwinBuilderFlowId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginGetDigitalTwinBuilderFlowDefinition(req.Context(), workspaceIDParam, digitalTwinBuilderFlowIDParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGetDigitalTwinBuilderFlowDefinition = &respr
		i.beginGetDigitalTwinBuilderFlowDefinition.add(req, beginGetDigitalTwinBuilderFlowDefinition)
	}

	resp, err := server.PollerResponderNext(beginGetDigitalTwinBuilderFlowDefinition, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		i.beginGetDigitalTwinBuilderFlowDefinition.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGetDigitalTwinBuilderFlowDefinition) {
		i.beginGetDigitalTwinBuilderFlowDefinition.remove(req)
	}

	return resp, nil
}

func (i *ItemsServerTransport) dispatchNewListDigitalTwinBuilderFlowsPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListDigitalTwinBuilderFlowsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListDigitalTwinBuilderFlowsPager not implemented")}
	}
	newListDigitalTwinBuilderFlowsPager := i.newListDigitalTwinBuilderFlowsPager.get(req)
	if newListDigitalTwinBuilderFlowsPager == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/DigitalTwinBuilderFlows`
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
		var options *digitaltwinbuilderflow.ItemsClientListDigitalTwinBuilderFlowsOptions
		if continuationTokenParam != nil {
			options = &digitaltwinbuilderflow.ItemsClientListDigitalTwinBuilderFlowsOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := i.srv.NewListDigitalTwinBuilderFlowsPager(workspaceIDParam, options)
		newListDigitalTwinBuilderFlowsPager = &resp
		i.newListDigitalTwinBuilderFlowsPager.add(req, newListDigitalTwinBuilderFlowsPager)
		server.PagerResponderInjectNextLinks(newListDigitalTwinBuilderFlowsPager, req, func(page *digitaltwinbuilderflow.ItemsClientListDigitalTwinBuilderFlowsResponse, createLink func() string) {
			page.ContinuationURI = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListDigitalTwinBuilderFlowsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListDigitalTwinBuilderFlowsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListDigitalTwinBuilderFlowsPager) {
		i.newListDigitalTwinBuilderFlowsPager.remove(req)
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchUpdateDigitalTwinBuilderFlow(req *http.Request) (*http.Response, error) {
	if i.srv.UpdateDigitalTwinBuilderFlow == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateDigitalTwinBuilderFlow not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/DigitalTwinBuilderFlows/(?P<digitalTwinBuilderFlowId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[digitaltwinbuilderflow.UpdateDigitalTwinBuilderFlowRequest](req)
	if err != nil {
		return nil, err
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	digitalTwinBuilderFlowIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("digitalTwinBuilderFlowId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.UpdateDigitalTwinBuilderFlow(req.Context(), workspaceIDParam, digitalTwinBuilderFlowIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DigitalTwinBuilderFlow, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchBeginUpdateDigitalTwinBuilderFlowDefinition(req *http.Request) (*http.Response, error) {
	if i.srv.BeginUpdateDigitalTwinBuilderFlowDefinition == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateDigitalTwinBuilderFlowDefinition not implemented")}
	}
	beginUpdateDigitalTwinBuilderFlowDefinition := i.beginUpdateDigitalTwinBuilderFlowDefinition.get(req)
	if beginUpdateDigitalTwinBuilderFlowDefinition == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/DigitalTwinBuilderFlows/(?P<digitalTwinBuilderFlowId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateDefinition`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		body, err := server.UnmarshalRequestAsJSON[digitaltwinbuilderflow.UpdateDigitalTwinBuilderFlowDefinitionRequest](req)
		if err != nil {
			return nil, err
		}
		workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
		if err != nil {
			return nil, err
		}
		digitalTwinBuilderFlowIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("digitalTwinBuilderFlowId")])
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
		var options *digitaltwinbuilderflow.ItemsClientBeginUpdateDigitalTwinBuilderFlowDefinitionOptions
		if updateMetadataParam != nil {
			options = &digitaltwinbuilderflow.ItemsClientBeginUpdateDigitalTwinBuilderFlowDefinitionOptions{
				UpdateMetadata: updateMetadataParam,
			}
		}
		respr, errRespr := i.srv.BeginUpdateDigitalTwinBuilderFlowDefinition(req.Context(), workspaceIDParam, digitalTwinBuilderFlowIDParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateDigitalTwinBuilderFlowDefinition = &respr
		i.beginUpdateDigitalTwinBuilderFlowDefinition.add(req, beginUpdateDigitalTwinBuilderFlowDefinition)
	}

	resp, err := server.PollerResponderNext(beginUpdateDigitalTwinBuilderFlowDefinition, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		i.beginUpdateDigitalTwinBuilderFlowDefinition.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateDigitalTwinBuilderFlowDefinition) {
		i.beginUpdateDigitalTwinBuilderFlowDefinition.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ItemsServerTransport
var itemsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
