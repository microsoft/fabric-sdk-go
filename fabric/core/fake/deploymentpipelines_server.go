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

	"github.com/microsoft/fabric-sdk-go/fabric/core"
)

// DeploymentPipelinesServer is a fake server for instances of the core.DeploymentPipelinesClient type.
type DeploymentPipelinesServer struct {
	// BeginDeployStageContent is the fake for method DeploymentPipelinesClient.BeginDeployStageContent
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginDeployStageContent func(ctx context.Context, deploymentPipelineID string, deployRequest core.DeployRequest, options *core.DeploymentPipelinesClientBeginDeployStageContentOptions) (resp azfake.PollerResponder[core.DeploymentPipelinesClientDeployStageContentResponse], errResp azfake.ErrorResponder)

	// GetDeploymentPipeline is the fake for method DeploymentPipelinesClient.GetDeploymentPipeline
	// HTTP status codes to indicate success: http.StatusOK
	GetDeploymentPipeline func(ctx context.Context, deploymentPipelineID string, options *core.DeploymentPipelinesClientGetDeploymentPipelineOptions) (resp azfake.Responder[core.DeploymentPipelinesClientGetDeploymentPipelineResponse], errResp azfake.ErrorResponder)

	// NewListDeploymentPipelineStageItemsPager is the fake for method DeploymentPipelinesClient.NewListDeploymentPipelineStageItemsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListDeploymentPipelineStageItemsPager func(deploymentPipelineID string, stageID string, options *core.DeploymentPipelinesClientListDeploymentPipelineStageItemsOptions) (resp azfake.PagerResponder[core.DeploymentPipelinesClientListDeploymentPipelineStageItemsResponse])

	// NewListDeploymentPipelineStagesPager is the fake for method DeploymentPipelinesClient.NewListDeploymentPipelineStagesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListDeploymentPipelineStagesPager func(deploymentPipelineID string, options *core.DeploymentPipelinesClientListDeploymentPipelineStagesOptions) (resp azfake.PagerResponder[core.DeploymentPipelinesClientListDeploymentPipelineStagesResponse])

	// NewListDeploymentPipelinesPager is the fake for method DeploymentPipelinesClient.NewListDeploymentPipelinesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListDeploymentPipelinesPager func(options *core.DeploymentPipelinesClientListDeploymentPipelinesOptions) (resp azfake.PagerResponder[core.DeploymentPipelinesClientListDeploymentPipelinesResponse])
}

// NewDeploymentPipelinesServerTransport creates a new instance of DeploymentPipelinesServerTransport with the provided implementation.
// The returned DeploymentPipelinesServerTransport instance is connected to an instance of core.DeploymentPipelinesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDeploymentPipelinesServerTransport(srv *DeploymentPipelinesServer) *DeploymentPipelinesServerTransport {
	return &DeploymentPipelinesServerTransport{
		srv:                                      srv,
		beginDeployStageContent:                  newTracker[azfake.PollerResponder[core.DeploymentPipelinesClientDeployStageContentResponse]](),
		newListDeploymentPipelineStageItemsPager: newTracker[azfake.PagerResponder[core.DeploymentPipelinesClientListDeploymentPipelineStageItemsResponse]](),
		newListDeploymentPipelineStagesPager:     newTracker[azfake.PagerResponder[core.DeploymentPipelinesClientListDeploymentPipelineStagesResponse]](),
		newListDeploymentPipelinesPager:          newTracker[azfake.PagerResponder[core.DeploymentPipelinesClientListDeploymentPipelinesResponse]](),
	}
}

// DeploymentPipelinesServerTransport connects instances of core.DeploymentPipelinesClient to instances of DeploymentPipelinesServer.
// Don't use this type directly, use NewDeploymentPipelinesServerTransport instead.
type DeploymentPipelinesServerTransport struct {
	srv                                      *DeploymentPipelinesServer
	beginDeployStageContent                  *tracker[azfake.PollerResponder[core.DeploymentPipelinesClientDeployStageContentResponse]]
	newListDeploymentPipelineStageItemsPager *tracker[azfake.PagerResponder[core.DeploymentPipelinesClientListDeploymentPipelineStageItemsResponse]]
	newListDeploymentPipelineStagesPager     *tracker[azfake.PagerResponder[core.DeploymentPipelinesClientListDeploymentPipelineStagesResponse]]
	newListDeploymentPipelinesPager          *tracker[azfake.PagerResponder[core.DeploymentPipelinesClientListDeploymentPipelinesResponse]]
}

// Do implements the policy.Transporter interface for DeploymentPipelinesServerTransport.
func (d *DeploymentPipelinesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	parts := strings.Split(method, ".")
	method = parts[1] + "." + parts[2]
	return d.dispatchToMethodFake(req, method)
}

func (d *DeploymentPipelinesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if deploymentPipelinesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = deploymentPipelinesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "DeploymentPipelinesClient.BeginDeployStageContent":
				res.resp, res.err = d.dispatchBeginDeployStageContent(req)
			case "DeploymentPipelinesClient.GetDeploymentPipeline":
				res.resp, res.err = d.dispatchGetDeploymentPipeline(req)
			case "DeploymentPipelinesClient.NewListDeploymentPipelineStageItemsPager":
				res.resp, res.err = d.dispatchNewListDeploymentPipelineStageItemsPager(req)
			case "DeploymentPipelinesClient.NewListDeploymentPipelineStagesPager":
				res.resp, res.err = d.dispatchNewListDeploymentPipelineStagesPager(req)
			case "DeploymentPipelinesClient.NewListDeploymentPipelinesPager":
				res.resp, res.err = d.dispatchNewListDeploymentPipelinesPager(req)
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

func (d *DeploymentPipelinesServerTransport) dispatchBeginDeployStageContent(req *http.Request) (*http.Response, error) {
	if d.srv.BeginDeployStageContent == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDeployStageContent not implemented")}
	}
	beginDeployStageContent := d.beginDeployStageContent.get(req)
	if beginDeployStageContent == nil {
		const regexStr = `/v1/deploymentPipelines/(?P<deploymentPipelineId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deploy`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[core.DeployRequest](req)
		if err != nil {
			return nil, err
		}
		deploymentPipelineIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("deploymentPipelineId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginDeployStageContent(req.Context(), deploymentPipelineIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDeployStageContent = &respr
		d.beginDeployStageContent.add(req, beginDeployStageContent)
	}

	resp, err := server.PollerResponderNext(beginDeployStageContent, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginDeployStageContent.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDeployStageContent) {
		d.beginDeployStageContent.remove(req)
	}

	return resp, nil
}

func (d *DeploymentPipelinesServerTransport) dispatchGetDeploymentPipeline(req *http.Request) (*http.Response, error) {
	if d.srv.GetDeploymentPipeline == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDeploymentPipeline not implemented")}
	}
	const regexStr = `/v1/deploymentPipelines/(?P<deploymentPipelineId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	deploymentPipelineIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("deploymentPipelineId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.GetDeploymentPipeline(req.Context(), deploymentPipelineIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeploymentPipeline, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DeploymentPipelinesServerTransport) dispatchNewListDeploymentPipelineStageItemsPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListDeploymentPipelineStageItemsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListDeploymentPipelineStageItemsPager not implemented")}
	}
	newListDeploymentPipelineStageItemsPager := d.newListDeploymentPipelineStageItemsPager.get(req)
	if newListDeploymentPipelineStageItemsPager == nil {
		const regexStr = `/v1/deploymentPipelines/(?P<deploymentPipelineId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stages/(?P<stageId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		deploymentPipelineIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("deploymentPipelineId")])
		if err != nil {
			return nil, err
		}
		stageIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("stageId")])
		if err != nil {
			return nil, err
		}
		continuationTokenUnescaped, err := url.QueryUnescape(qp.Get("continuationToken"))
		if err != nil {
			return nil, err
		}
		continuationTokenParam := getOptional(continuationTokenUnescaped)
		var options *core.DeploymentPipelinesClientListDeploymentPipelineStageItemsOptions
		if continuationTokenParam != nil {
			options = &core.DeploymentPipelinesClientListDeploymentPipelineStageItemsOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := d.srv.NewListDeploymentPipelineStageItemsPager(deploymentPipelineIDParam, stageIDParam, options)
		newListDeploymentPipelineStageItemsPager = &resp
		d.newListDeploymentPipelineStageItemsPager.add(req, newListDeploymentPipelineStageItemsPager)
		server.PagerResponderInjectNextLinks(newListDeploymentPipelineStageItemsPager, req, func(page *core.DeploymentPipelinesClientListDeploymentPipelineStageItemsResponse, createLink func() string) {
			page.ContinuationURI = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListDeploymentPipelineStageItemsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListDeploymentPipelineStageItemsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListDeploymentPipelineStageItemsPager) {
		d.newListDeploymentPipelineStageItemsPager.remove(req)
	}
	return resp, nil
}

func (d *DeploymentPipelinesServerTransport) dispatchNewListDeploymentPipelineStagesPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListDeploymentPipelineStagesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListDeploymentPipelineStagesPager not implemented")}
	}
	newListDeploymentPipelineStagesPager := d.newListDeploymentPipelineStagesPager.get(req)
	if newListDeploymentPipelineStagesPager == nil {
		const regexStr = `/v1/deploymentPipelines/(?P<deploymentPipelineId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stages`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		deploymentPipelineIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("deploymentPipelineId")])
		if err != nil {
			return nil, err
		}
		continuationTokenUnescaped, err := url.QueryUnescape(qp.Get("continuationToken"))
		if err != nil {
			return nil, err
		}
		continuationTokenParam := getOptional(continuationTokenUnescaped)
		var options *core.DeploymentPipelinesClientListDeploymentPipelineStagesOptions
		if continuationTokenParam != nil {
			options = &core.DeploymentPipelinesClientListDeploymentPipelineStagesOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := d.srv.NewListDeploymentPipelineStagesPager(deploymentPipelineIDParam, options)
		newListDeploymentPipelineStagesPager = &resp
		d.newListDeploymentPipelineStagesPager.add(req, newListDeploymentPipelineStagesPager)
		server.PagerResponderInjectNextLinks(newListDeploymentPipelineStagesPager, req, func(page *core.DeploymentPipelinesClientListDeploymentPipelineStagesResponse, createLink func() string) {
			page.ContinuationURI = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListDeploymentPipelineStagesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListDeploymentPipelineStagesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListDeploymentPipelineStagesPager) {
		d.newListDeploymentPipelineStagesPager.remove(req)
	}
	return resp, nil
}

func (d *DeploymentPipelinesServerTransport) dispatchNewListDeploymentPipelinesPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListDeploymentPipelinesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListDeploymentPipelinesPager not implemented")}
	}
	newListDeploymentPipelinesPager := d.newListDeploymentPipelinesPager.get(req)
	if newListDeploymentPipelinesPager == nil {
		qp := req.URL.Query()
		continuationTokenUnescaped, err := url.QueryUnescape(qp.Get("continuationToken"))
		if err != nil {
			return nil, err
		}
		continuationTokenParam := getOptional(continuationTokenUnescaped)
		var options *core.DeploymentPipelinesClientListDeploymentPipelinesOptions
		if continuationTokenParam != nil {
			options = &core.DeploymentPipelinesClientListDeploymentPipelinesOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := d.srv.NewListDeploymentPipelinesPager(options)
		newListDeploymentPipelinesPager = &resp
		d.newListDeploymentPipelinesPager.add(req, newListDeploymentPipelinesPager)
		server.PagerResponderInjectNextLinks(newListDeploymentPipelinesPager, req, func(page *core.DeploymentPipelinesClientListDeploymentPipelinesResponse, createLink func() string) {
			page.ContinuationURI = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListDeploymentPipelinesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListDeploymentPipelinesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListDeploymentPipelinesPager) {
		d.newListDeploymentPipelinesPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to DeploymentPipelinesServerTransport
var deploymentPipelinesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
