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

	"github.com/microsoft/fabric-sdk-go/fabric/environment"
)

// SparkComputeServer is a fake server for instances of the environment.SparkComputeClient type.
type SparkComputeServer struct {
	// GetPublishedSettings is the fake for method SparkComputeClient.GetPublishedSettings
	// HTTP status codes to indicate success: http.StatusOK
	GetPublishedSettings func(ctx context.Context, workspaceID string, environmentID string, options *environment.SparkComputeClientGetPublishedSettingsOptions) (resp azfake.Responder[environment.SparkComputeClientGetPublishedSettingsResponse], errResp azfake.ErrorResponder)

	// GetStagingSettings is the fake for method SparkComputeClient.GetStagingSettings
	// HTTP status codes to indicate success: http.StatusOK
	GetStagingSettings func(ctx context.Context, workspaceID string, environmentID string, options *environment.SparkComputeClientGetStagingSettingsOptions) (resp azfake.Responder[environment.SparkComputeClientGetStagingSettingsResponse], errResp azfake.ErrorResponder)

	// UpdateStagingSettings is the fake for method SparkComputeClient.UpdateStagingSettings
	// HTTP status codes to indicate success: http.StatusOK
	UpdateStagingSettings func(ctx context.Context, workspaceID string, environmentID string, updateEnvironmentSparkComputeRequest environment.UpdateEnvironmentSparkComputeRequest, options *environment.SparkComputeClientUpdateStagingSettingsOptions) (resp azfake.Responder[environment.SparkComputeClientUpdateStagingSettingsResponse], errResp azfake.ErrorResponder)
}

// NewSparkComputeServerTransport creates a new instance of SparkComputeServerTransport with the provided implementation.
// The returned SparkComputeServerTransport instance is connected to an instance of environment.SparkComputeClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSparkComputeServerTransport(srv *SparkComputeServer) *SparkComputeServerTransport {
	return &SparkComputeServerTransport{srv: srv}
}

// SparkComputeServerTransport connects instances of environment.SparkComputeClient to instances of SparkComputeServer.
// Don't use this type directly, use NewSparkComputeServerTransport instead.
type SparkComputeServerTransport struct {
	srv *SparkComputeServer
}

// Do implements the policy.Transporter interface for SparkComputeServerTransport.
func (s *SparkComputeServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	parts := strings.Split(method, ".")
	method = parts[1] + "." + parts[2]
	return s.dispatchToMethodFake(req, method)
}

func (s *SparkComputeServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if sparkComputeServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = sparkComputeServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "SparkComputeClient.GetPublishedSettings":
				res.resp, res.err = s.dispatchGetPublishedSettings(req)
			case "SparkComputeClient.GetStagingSettings":
				res.resp, res.err = s.dispatchGetStagingSettings(req)
			case "SparkComputeClient.UpdateStagingSettings":
				res.resp, res.err = s.dispatchUpdateStagingSettings(req)
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

func (s *SparkComputeServerTransport) dispatchGetPublishedSettings(req *http.Request) (*http.Response, error) {
	if s.srv.GetPublishedSettings == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetPublishedSettings not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/environments/(?P<environmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sparkcompute`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	environmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetPublishedSettings(req.Context(), workspaceIDParam, environmentIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SparkCompute, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SparkComputeServerTransport) dispatchGetStagingSettings(req *http.Request) (*http.Response, error) {
	if s.srv.GetStagingSettings == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetStagingSettings not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/environments/(?P<environmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/staging/sparkcompute`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	environmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetStagingSettings(req.Context(), workspaceIDParam, environmentIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SparkCompute, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SparkComputeServerTransport) dispatchUpdateStagingSettings(req *http.Request) (*http.Response, error) {
	if s.srv.UpdateStagingSettings == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateStagingSettings not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/environments/(?P<environmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/staging/sparkcompute`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[environment.UpdateEnvironmentSparkComputeRequest](req)
	if err != nil {
		return nil, err
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	environmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.UpdateStagingSettings(req.Context(), workspaceIDParam, environmentIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SparkCompute, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to SparkComputeServerTransport
var sparkComputeServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
