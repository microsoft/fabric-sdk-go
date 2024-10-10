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

	"github.com/microsoft/fabric-sdk-go/fabric/lakehouse"
)

// BackgroundJobsServer is a fake server for instances of the lakehouse.BackgroundJobsClient type.
type BackgroundJobsServer struct {
	// RunOnDemandTableMaintenance is the fake for method BackgroundJobsClient.RunOnDemandTableMaintenance
	// HTTP status codes to indicate success: http.StatusAccepted
	RunOnDemandTableMaintenance func(ctx context.Context, workspaceID string, lakehouseID string, jobType string, runOnDemandTableMaintenanceRequest lakehouse.RunOnDemandTableMaintenanceRequest, options *lakehouse.BackgroundJobsClientRunOnDemandTableMaintenanceOptions) (resp azfake.Responder[lakehouse.BackgroundJobsClientRunOnDemandTableMaintenanceResponse], errResp azfake.ErrorResponder)
}

// NewBackgroundJobsServerTransport creates a new instance of BackgroundJobsServerTransport with the provided implementation.
// The returned BackgroundJobsServerTransport instance is connected to an instance of lakehouse.BackgroundJobsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBackgroundJobsServerTransport(srv *BackgroundJobsServer) *BackgroundJobsServerTransport {
	return &BackgroundJobsServerTransport{srv: srv}
}

// BackgroundJobsServerTransport connects instances of lakehouse.BackgroundJobsClient to instances of BackgroundJobsServer.
// Don't use this type directly, use NewBackgroundJobsServerTransport instead.
type BackgroundJobsServerTransport struct {
	srv *BackgroundJobsServer
}

// Do implements the policy.Transporter interface for BackgroundJobsServerTransport.
func (b *BackgroundJobsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	parts := strings.Split(method, ".")
	method = parts[1] + "." + parts[2]
	return b.dispatchToMethodFake(req, method)
}

func (b *BackgroundJobsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "BackgroundJobsClient.RunOnDemandTableMaintenance":
			res.resp, res.err = b.dispatchRunOnDemandTableMaintenance(req)
		default:
			res.err = fmt.Errorf("unhandled API %s", method)
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

func (b *BackgroundJobsServerTransport) dispatchRunOnDemandTableMaintenance(req *http.Request) (*http.Response, error) {
	if b.srv.RunOnDemandTableMaintenance == nil {
		return nil, &nonRetriableError{errors.New("fake for method RunOnDemandTableMaintenance not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/lakehouses/(?P<lakehouseId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/instances`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[lakehouse.RunOnDemandTableMaintenanceRequest](req)
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
	jobTypeParam, err := url.QueryUnescape(qp.Get("jobType"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.RunOnDemandTableMaintenance(req.Context(), workspaceIDParam, lakehouseIDParam, jobTypeParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
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
