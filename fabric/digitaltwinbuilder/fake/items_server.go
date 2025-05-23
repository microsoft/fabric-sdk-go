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

	"github.com/microsoft/fabric-sdk-go/fabric/digitaltwinbuilder"
)

// ItemsServer is a fake server for instances of the digitaltwinbuilder.ItemsClient type.
type ItemsServer struct {
	// BeginCreateDigitalTwinBuilder is the fake for method ItemsClient.BeginCreateDigitalTwinBuilder
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateDigitalTwinBuilder func(ctx context.Context, workspaceID string, createDigitalTwinBuilderRequest digitaltwinbuilder.CreateDigitalTwinBuilderRequest, options *digitaltwinbuilder.ItemsClientBeginCreateDigitalTwinBuilderOptions) (resp azfake.PollerResponder[digitaltwinbuilder.ItemsClientCreateDigitalTwinBuilderResponse], errResp azfake.ErrorResponder)

	// DeleteDigitalTwinBuilder is the fake for method ItemsClient.DeleteDigitalTwinBuilder
	// HTTP status codes to indicate success: http.StatusOK
	DeleteDigitalTwinBuilder func(ctx context.Context, workspaceID string, digitaltwinbuilderID string, options *digitaltwinbuilder.ItemsClientDeleteDigitalTwinBuilderOptions) (resp azfake.Responder[digitaltwinbuilder.ItemsClientDeleteDigitalTwinBuilderResponse], errResp azfake.ErrorResponder)

	// GetDigitalTwinBuilder is the fake for method ItemsClient.GetDigitalTwinBuilder
	// HTTP status codes to indicate success: http.StatusOK
	GetDigitalTwinBuilder func(ctx context.Context, workspaceID string, digitaltwinbuilderID string, options *digitaltwinbuilder.ItemsClientGetDigitalTwinBuilderOptions) (resp azfake.Responder[digitaltwinbuilder.ItemsClientGetDigitalTwinBuilderResponse], errResp azfake.ErrorResponder)

	// BeginGetDigitalTwinBuilderDefinition is the fake for method ItemsClient.BeginGetDigitalTwinBuilderDefinition
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGetDigitalTwinBuilderDefinition func(ctx context.Context, workspaceID string, digitaltwinbuilderID string, options *digitaltwinbuilder.ItemsClientBeginGetDigitalTwinBuilderDefinitionOptions) (resp azfake.PollerResponder[digitaltwinbuilder.ItemsClientGetDigitalTwinBuilderDefinitionResponse], errResp azfake.ErrorResponder)

	// NewListDigitalTwinBuildersPager is the fake for method ItemsClient.NewListDigitalTwinBuildersPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListDigitalTwinBuildersPager func(workspaceID string, options *digitaltwinbuilder.ItemsClientListDigitalTwinBuildersOptions) (resp azfake.PagerResponder[digitaltwinbuilder.ItemsClientListDigitalTwinBuildersResponse])

	// UpdateDigitalTwinBuilder is the fake for method ItemsClient.UpdateDigitalTwinBuilder
	// HTTP status codes to indicate success: http.StatusOK
	UpdateDigitalTwinBuilder func(ctx context.Context, workspaceID string, digitaltwinbuilderID string, updateDigitalTwinBuilderRequest digitaltwinbuilder.UpdateDigitalTwinBuilderRequest, options *digitaltwinbuilder.ItemsClientUpdateDigitalTwinBuilderOptions) (resp azfake.Responder[digitaltwinbuilder.ItemsClientUpdateDigitalTwinBuilderResponse], errResp azfake.ErrorResponder)

	// BeginUpdateDigitalTwinBuilderDefinition is the fake for method ItemsClient.BeginUpdateDigitalTwinBuilderDefinition
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginUpdateDigitalTwinBuilderDefinition func(ctx context.Context, workspaceID string, digitaltwinbuilderID string, updateDigitalTwinBuilderDefinitionRequest digitaltwinbuilder.UpdateDigitalTwinBuilderDefinitionRequest, options *digitaltwinbuilder.ItemsClientBeginUpdateDigitalTwinBuilderDefinitionOptions) (resp azfake.PollerResponder[digitaltwinbuilder.ItemsClientUpdateDigitalTwinBuilderDefinitionResponse], errResp azfake.ErrorResponder)
}

// NewItemsServerTransport creates a new instance of ItemsServerTransport with the provided implementation.
// The returned ItemsServerTransport instance is connected to an instance of digitaltwinbuilder.ItemsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewItemsServerTransport(srv *ItemsServer) *ItemsServerTransport {
	return &ItemsServerTransport{
		srv:                                     srv,
		beginCreateDigitalTwinBuilder:           newTracker[azfake.PollerResponder[digitaltwinbuilder.ItemsClientCreateDigitalTwinBuilderResponse]](),
		beginGetDigitalTwinBuilderDefinition:    newTracker[azfake.PollerResponder[digitaltwinbuilder.ItemsClientGetDigitalTwinBuilderDefinitionResponse]](),
		newListDigitalTwinBuildersPager:         newTracker[azfake.PagerResponder[digitaltwinbuilder.ItemsClientListDigitalTwinBuildersResponse]](),
		beginUpdateDigitalTwinBuilderDefinition: newTracker[azfake.PollerResponder[digitaltwinbuilder.ItemsClientUpdateDigitalTwinBuilderDefinitionResponse]](),
	}
}

// ItemsServerTransport connects instances of digitaltwinbuilder.ItemsClient to instances of ItemsServer.
// Don't use this type directly, use NewItemsServerTransport instead.
type ItemsServerTransport struct {
	srv                                     *ItemsServer
	beginCreateDigitalTwinBuilder           *tracker[azfake.PollerResponder[digitaltwinbuilder.ItemsClientCreateDigitalTwinBuilderResponse]]
	beginGetDigitalTwinBuilderDefinition    *tracker[azfake.PollerResponder[digitaltwinbuilder.ItemsClientGetDigitalTwinBuilderDefinitionResponse]]
	newListDigitalTwinBuildersPager         *tracker[azfake.PagerResponder[digitaltwinbuilder.ItemsClientListDigitalTwinBuildersResponse]]
	beginUpdateDigitalTwinBuilderDefinition *tracker[azfake.PollerResponder[digitaltwinbuilder.ItemsClientUpdateDigitalTwinBuilderDefinitionResponse]]
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
			case "ItemsClient.BeginCreateDigitalTwinBuilder":
				res.resp, res.err = i.dispatchBeginCreateDigitalTwinBuilder(req)
			case "ItemsClient.DeleteDigitalTwinBuilder":
				res.resp, res.err = i.dispatchDeleteDigitalTwinBuilder(req)
			case "ItemsClient.GetDigitalTwinBuilder":
				res.resp, res.err = i.dispatchGetDigitalTwinBuilder(req)
			case "ItemsClient.BeginGetDigitalTwinBuilderDefinition":
				res.resp, res.err = i.dispatchBeginGetDigitalTwinBuilderDefinition(req)
			case "ItemsClient.NewListDigitalTwinBuildersPager":
				res.resp, res.err = i.dispatchNewListDigitalTwinBuildersPager(req)
			case "ItemsClient.UpdateDigitalTwinBuilder":
				res.resp, res.err = i.dispatchUpdateDigitalTwinBuilder(req)
			case "ItemsClient.BeginUpdateDigitalTwinBuilderDefinition":
				res.resp, res.err = i.dispatchBeginUpdateDigitalTwinBuilderDefinition(req)
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

func (i *ItemsServerTransport) dispatchBeginCreateDigitalTwinBuilder(req *http.Request) (*http.Response, error) {
	if i.srv.BeginCreateDigitalTwinBuilder == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateDigitalTwinBuilder not implemented")}
	}
	beginCreateDigitalTwinBuilder := i.beginCreateDigitalTwinBuilder.get(req)
	if beginCreateDigitalTwinBuilder == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/digitaltwinbuilders`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[digitaltwinbuilder.CreateDigitalTwinBuilderRequest](req)
		if err != nil {
			return nil, err
		}
		workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginCreateDigitalTwinBuilder(req.Context(), workspaceIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateDigitalTwinBuilder = &respr
		i.beginCreateDigitalTwinBuilder.add(req, beginCreateDigitalTwinBuilder)
	}

	resp, err := server.PollerResponderNext(beginCreateDigitalTwinBuilder, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		i.beginCreateDigitalTwinBuilder.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateDigitalTwinBuilder) {
		i.beginCreateDigitalTwinBuilder.remove(req)
	}

	return resp, nil
}

func (i *ItemsServerTransport) dispatchDeleteDigitalTwinBuilder(req *http.Request) (*http.Response, error) {
	if i.srv.DeleteDigitalTwinBuilder == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteDigitalTwinBuilder not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/digitaltwinbuilders/(?P<digitaltwinbuilderId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	digitaltwinbuilderIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("digitaltwinbuilderId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.DeleteDigitalTwinBuilder(req.Context(), workspaceIDParam, digitaltwinbuilderIDParam, nil)
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

func (i *ItemsServerTransport) dispatchGetDigitalTwinBuilder(req *http.Request) (*http.Response, error) {
	if i.srv.GetDigitalTwinBuilder == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDigitalTwinBuilder not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/digitaltwinbuilders/(?P<digitaltwinbuilderId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	digitaltwinbuilderIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("digitaltwinbuilderId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.GetDigitalTwinBuilder(req.Context(), workspaceIDParam, digitaltwinbuilderIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DigitalTwinBuilder, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchBeginGetDigitalTwinBuilderDefinition(req *http.Request) (*http.Response, error) {
	if i.srv.BeginGetDigitalTwinBuilderDefinition == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGetDigitalTwinBuilderDefinition not implemented")}
	}
	beginGetDigitalTwinBuilderDefinition := i.beginGetDigitalTwinBuilderDefinition.get(req)
	if beginGetDigitalTwinBuilderDefinition == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/digitaltwinbuilders/(?P<digitaltwinbuilderId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getDefinition`
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
		digitaltwinbuilderIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("digitaltwinbuilderId")])
		if err != nil {
			return nil, err
		}
		formatUnescaped, err := url.QueryUnescape(qp.Get("format"))
		if err != nil {
			return nil, err
		}
		formatParam := getOptional(formatUnescaped)
		var options *digitaltwinbuilder.ItemsClientBeginGetDigitalTwinBuilderDefinitionOptions
		if formatParam != nil {
			options = &digitaltwinbuilder.ItemsClientBeginGetDigitalTwinBuilderDefinitionOptions{
				Format: formatParam,
			}
		}
		respr, errRespr := i.srv.BeginGetDigitalTwinBuilderDefinition(req.Context(), workspaceIDParam, digitaltwinbuilderIDParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGetDigitalTwinBuilderDefinition = &respr
		i.beginGetDigitalTwinBuilderDefinition.add(req, beginGetDigitalTwinBuilderDefinition)
	}

	resp, err := server.PollerResponderNext(beginGetDigitalTwinBuilderDefinition, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		i.beginGetDigitalTwinBuilderDefinition.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGetDigitalTwinBuilderDefinition) {
		i.beginGetDigitalTwinBuilderDefinition.remove(req)
	}

	return resp, nil
}

func (i *ItemsServerTransport) dispatchNewListDigitalTwinBuildersPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListDigitalTwinBuildersPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListDigitalTwinBuildersPager not implemented")}
	}
	newListDigitalTwinBuildersPager := i.newListDigitalTwinBuildersPager.get(req)
	if newListDigitalTwinBuildersPager == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/digitaltwinbuilders`
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
		var options *digitaltwinbuilder.ItemsClientListDigitalTwinBuildersOptions
		if continuationTokenParam != nil {
			options = &digitaltwinbuilder.ItemsClientListDigitalTwinBuildersOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := i.srv.NewListDigitalTwinBuildersPager(workspaceIDParam, options)
		newListDigitalTwinBuildersPager = &resp
		i.newListDigitalTwinBuildersPager.add(req, newListDigitalTwinBuildersPager)
		server.PagerResponderInjectNextLinks(newListDigitalTwinBuildersPager, req, func(page *digitaltwinbuilder.ItemsClientListDigitalTwinBuildersResponse, createLink func() string) {
			page.ContinuationURI = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListDigitalTwinBuildersPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListDigitalTwinBuildersPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListDigitalTwinBuildersPager) {
		i.newListDigitalTwinBuildersPager.remove(req)
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchUpdateDigitalTwinBuilder(req *http.Request) (*http.Response, error) {
	if i.srv.UpdateDigitalTwinBuilder == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateDigitalTwinBuilder not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/digitaltwinbuilders/(?P<digitaltwinbuilderId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[digitaltwinbuilder.UpdateDigitalTwinBuilderRequest](req)
	if err != nil {
		return nil, err
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	digitaltwinbuilderIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("digitaltwinbuilderId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.UpdateDigitalTwinBuilder(req.Context(), workspaceIDParam, digitaltwinbuilderIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DigitalTwinBuilder, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ItemsServerTransport) dispatchBeginUpdateDigitalTwinBuilderDefinition(req *http.Request) (*http.Response, error) {
	if i.srv.BeginUpdateDigitalTwinBuilderDefinition == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateDigitalTwinBuilderDefinition not implemented")}
	}
	beginUpdateDigitalTwinBuilderDefinition := i.beginUpdateDigitalTwinBuilderDefinition.get(req)
	if beginUpdateDigitalTwinBuilderDefinition == nil {
		const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/digitaltwinbuilders/(?P<digitaltwinbuilderId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateDefinition`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		body, err := server.UnmarshalRequestAsJSON[digitaltwinbuilder.UpdateDigitalTwinBuilderDefinitionRequest](req)
		if err != nil {
			return nil, err
		}
		workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
		if err != nil {
			return nil, err
		}
		digitaltwinbuilderIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("digitaltwinbuilderId")])
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
		var options *digitaltwinbuilder.ItemsClientBeginUpdateDigitalTwinBuilderDefinitionOptions
		if updateMetadataParam != nil {
			options = &digitaltwinbuilder.ItemsClientBeginUpdateDigitalTwinBuilderDefinitionOptions{
				UpdateMetadata: updateMetadataParam,
			}
		}
		respr, errRespr := i.srv.BeginUpdateDigitalTwinBuilderDefinition(req.Context(), workspaceIDParam, digitaltwinbuilderIDParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateDigitalTwinBuilderDefinition = &respr
		i.beginUpdateDigitalTwinBuilderDefinition.add(req, beginUpdateDigitalTwinBuilderDefinition)
	}

	resp, err := server.PollerResponderNext(beginUpdateDigitalTwinBuilderDefinition, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		i.beginUpdateDigitalTwinBuilderDefinition.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateDigitalTwinBuilderDefinition) {
		i.beginUpdateDigitalTwinBuilderDefinition.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ItemsServerTransport
var itemsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
