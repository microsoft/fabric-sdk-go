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

	"github.com/microsoft/fabric-sdk-go/fabric/environment"
)

// SparkLibrariesServer is a fake server for instances of the environment.SparkLibrariesClient type.
type SparkLibrariesServer struct {
	// CancelPublish is the fake for method SparkLibrariesClient.CancelPublish
	// HTTP status codes to indicate success: http.StatusOK
	CancelPublish func(ctx context.Context, workspaceID string, environmentID string, options *environment.SparkLibrariesClientCancelPublishOptions) (resp azfake.Responder[environment.SparkLibrariesClientCancelPublishResponse], errResp azfake.ErrorResponder)

	// DeleteStagingLibrary is the fake for method SparkLibrariesClient.DeleteStagingLibrary
	// HTTP status codes to indicate success: http.StatusOK
	DeleteStagingLibrary func(ctx context.Context, workspaceID string, environmentID string, libraryToDelete string, options *environment.SparkLibrariesClientDeleteStagingLibraryOptions) (resp azfake.Responder[environment.SparkLibrariesClientDeleteStagingLibraryResponse], errResp azfake.ErrorResponder)

	// GetPublishedLibraries is the fake for method SparkLibrariesClient.GetPublishedLibraries
	// HTTP status codes to indicate success: http.StatusOK
	GetPublishedLibraries func(ctx context.Context, workspaceID string, environmentID string, options *environment.SparkLibrariesClientGetPublishedLibrariesOptions) (resp azfake.Responder[environment.SparkLibrariesClientGetPublishedLibrariesResponse], errResp azfake.ErrorResponder)

	// GetStagingLibraries is the fake for method SparkLibrariesClient.GetStagingLibraries
	// HTTP status codes to indicate success: http.StatusOK
	GetStagingLibraries func(ctx context.Context, workspaceID string, environmentID string, options *environment.SparkLibrariesClientGetStagingLibrariesOptions) (resp azfake.Responder[environment.SparkLibrariesClientGetStagingLibrariesResponse], errResp azfake.ErrorResponder)

	// PublishEnvironment is the fake for method SparkLibrariesClient.PublishEnvironment
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	PublishEnvironment func(ctx context.Context, workspaceID string, environmentID string, options *environment.SparkLibrariesClientPublishEnvironmentOptions) (resp azfake.Responder[environment.SparkLibrariesClientPublishEnvironmentResponse], errResp azfake.ErrorResponder)

	// UploadStagingLibrary is the fake for method SparkLibrariesClient.UploadStagingLibrary
	// HTTP status codes to indicate success: http.StatusOK
	UploadStagingLibrary func(ctx context.Context, workspaceID string, environmentID string, options *environment.SparkLibrariesClientUploadStagingLibraryOptions) (resp azfake.Responder[environment.SparkLibrariesClientUploadStagingLibraryResponse], errResp azfake.ErrorResponder)
}

// NewSparkLibrariesServerTransport creates a new instance of SparkLibrariesServerTransport with the provided implementation.
// The returned SparkLibrariesServerTransport instance is connected to an instance of environment.SparkLibrariesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSparkLibrariesServerTransport(srv *SparkLibrariesServer) *SparkLibrariesServerTransport {
	return &SparkLibrariesServerTransport{srv: srv}
}

// SparkLibrariesServerTransport connects instances of environment.SparkLibrariesClient to instances of SparkLibrariesServer.
// Don't use this type directly, use NewSparkLibrariesServerTransport instead.
type SparkLibrariesServerTransport struct {
	srv *SparkLibrariesServer
}

// Do implements the policy.Transporter interface for SparkLibrariesServerTransport.
func (s *SparkLibrariesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	parts := strings.Split(method, ".")
	method = parts[1] + "." + parts[2]
	return s.dispatchToMethodFake(req, method)
}

func (s *SparkLibrariesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "SparkLibrariesClient.CancelPublish":
		resp, err = s.dispatchCancelPublish(req)
	case "SparkLibrariesClient.DeleteStagingLibrary":
		resp, err = s.dispatchDeleteStagingLibrary(req)
	case "SparkLibrariesClient.GetPublishedLibraries":
		resp, err = s.dispatchGetPublishedLibraries(req)
	case "SparkLibrariesClient.GetStagingLibraries":
		resp, err = s.dispatchGetStagingLibraries(req)
	case "SparkLibrariesClient.PublishEnvironment":
		resp, err = s.dispatchPublishEnvironment(req)
	case "SparkLibrariesClient.UploadStagingLibrary":
		resp, err = s.dispatchUploadStagingLibrary(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (s *SparkLibrariesServerTransport) dispatchCancelPublish(req *http.Request) (*http.Response, error) {
	if s.srv.CancelPublish == nil {
		return nil, &nonRetriableError{errors.New("fake for method CancelPublish not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/environments/(?P<environmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/staging/cancelPublish`
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
	respr, errRespr := s.srv.CancelPublish(req.Context(), workspaceIDParam, environmentIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PublishInfo, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SparkLibrariesServerTransport) dispatchDeleteStagingLibrary(req *http.Request) (*http.Response, error) {
	if s.srv.DeleteStagingLibrary == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteStagingLibrary not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/environments/(?P<environmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/staging/libraries`
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
	environmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentId")])
	if err != nil {
		return nil, err
	}
	libraryToDeleteParam, err := url.QueryUnescape(qp.Get("libraryToDelete"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.DeleteStagingLibrary(req.Context(), workspaceIDParam, environmentIDParam, libraryToDeleteParam, nil)
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

func (s *SparkLibrariesServerTransport) dispatchGetPublishedLibraries(req *http.Request) (*http.Response, error) {
	if s.srv.GetPublishedLibraries == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetPublishedLibraries not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/environments/(?P<environmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/libraries`
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
	respr, errRespr := s.srv.GetPublishedLibraries(req.Context(), workspaceIDParam, environmentIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Libraries, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SparkLibrariesServerTransport) dispatchGetStagingLibraries(req *http.Request) (*http.Response, error) {
	if s.srv.GetStagingLibraries == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetStagingLibraries not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/environments/(?P<environmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/staging/libraries`
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
	respr, errRespr := s.srv.GetStagingLibraries(req.Context(), workspaceIDParam, environmentIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Libraries, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SparkLibrariesServerTransport) dispatchPublishEnvironment(req *http.Request) (*http.Response, error) {
	if s.srv.PublishEnvironment == nil {
		return nil, &nonRetriableError{errors.New("fake for method PublishEnvironment not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/environments/(?P<environmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/staging/publish`
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
	respr, errRespr := s.srv.PublishEnvironment(req.Context(), workspaceIDParam, environmentIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PublishInfo, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).Location; val != nil {
		resp.Header.Set("Location", *val)
	}
	if val := server.GetResponse(respr).RetryAfter; val != nil {
		resp.Header.Set("Retry-After", strconv.FormatInt(int64(*val), 10))
	}
	return resp, nil
}

func (s *SparkLibrariesServerTransport) dispatchUploadStagingLibrary(req *http.Request) (*http.Response, error) {
	if s.srv.UploadStagingLibrary == nil {
		return nil, &nonRetriableError{errors.New("fake for method UploadStagingLibrary not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/environments/(?P<environmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/staging/libraries`
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
	respr, errRespr := s.srv.UploadStagingLibrary(req.Context(), workspaceIDParam, environmentIDParam, nil)
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
