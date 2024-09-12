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
	"reflect"
	"regexp"
	"strconv"
	"strings"

	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"

	"github.com/microsoft/fabric-sdk-go/fabric/core"
)

// JobSchedulerServer is a fake server for instances of the core.JobSchedulerClient type.
type JobSchedulerServer struct {
	// CancelItemJobInstance is the fake for method JobSchedulerClient.CancelItemJobInstance
	// HTTP status codes to indicate success: http.StatusAccepted
	CancelItemJobInstance func(ctx context.Context, workspaceID string, itemID string, jobInstanceID string, options *core.JobSchedulerClientCancelItemJobInstanceOptions) (resp azfake.Responder[core.JobSchedulerClientCancelItemJobInstanceResponse], errResp azfake.ErrorResponder)

	// CreateItemSchedule is the fake for method JobSchedulerClient.CreateItemSchedule
	// HTTP status codes to indicate success: http.StatusCreated
	CreateItemSchedule func(ctx context.Context, workspaceID string, itemID string, jobType string, createScheduleRequest core.CreateScheduleRequest, options *core.JobSchedulerClientCreateItemScheduleOptions) (resp azfake.Responder[core.JobSchedulerClientCreateItemScheduleResponse], errResp azfake.ErrorResponder)

	// GetItemJobInstance is the fake for method JobSchedulerClient.GetItemJobInstance
	// HTTP status codes to indicate success: http.StatusOK
	GetItemJobInstance func(ctx context.Context, workspaceID string, itemID string, jobInstanceID string, options *core.JobSchedulerClientGetItemJobInstanceOptions) (resp azfake.Responder[core.JobSchedulerClientGetItemJobInstanceResponse], errResp azfake.ErrorResponder)

	// GetItemSchedule is the fake for method JobSchedulerClient.GetItemSchedule
	// HTTP status codes to indicate success: http.StatusOK
	GetItemSchedule func(ctx context.Context, workspaceID string, itemID string, jobType string, scheduleID string, options *core.JobSchedulerClientGetItemScheduleOptions) (resp azfake.Responder[core.JobSchedulerClientGetItemScheduleResponse], errResp azfake.ErrorResponder)

	// ListItemSchedules is the fake for method JobSchedulerClient.ListItemSchedules
	// HTTP status codes to indicate success: http.StatusOK
	ListItemSchedules func(ctx context.Context, workspaceID string, itemID string, jobType string, options *core.JobSchedulerClientListItemSchedulesOptions) (resp azfake.Responder[core.JobSchedulerClientListItemSchedulesResponse], errResp azfake.ErrorResponder)

	// RunOnDemandItemJob is the fake for method JobSchedulerClient.RunOnDemandItemJob
	// HTTP status codes to indicate success: http.StatusAccepted
	RunOnDemandItemJob func(ctx context.Context, workspaceID string, itemID string, jobType string, options *core.JobSchedulerClientRunOnDemandItemJobOptions) (resp azfake.Responder[core.JobSchedulerClientRunOnDemandItemJobResponse], errResp azfake.ErrorResponder)

	// UpdateItemSchedule is the fake for method JobSchedulerClient.UpdateItemSchedule
	// HTTP status codes to indicate success: http.StatusOK
	UpdateItemSchedule func(ctx context.Context, workspaceID string, itemID string, jobType string, scheduleID string, updateScheduleRequest core.UpdateScheduleRequest, options *core.JobSchedulerClientUpdateItemScheduleOptions) (resp azfake.Responder[core.JobSchedulerClientUpdateItemScheduleResponse], errResp azfake.ErrorResponder)
}

// NewJobSchedulerServerTransport creates a new instance of JobSchedulerServerTransport with the provided implementation.
// The returned JobSchedulerServerTransport instance is connected to an instance of core.JobSchedulerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewJobSchedulerServerTransport(srv *JobSchedulerServer) *JobSchedulerServerTransport {
	return &JobSchedulerServerTransport{srv: srv}
}

// JobSchedulerServerTransport connects instances of core.JobSchedulerClient to instances of JobSchedulerServer.
// Don't use this type directly, use NewJobSchedulerServerTransport instead.
type JobSchedulerServerTransport struct {
	srv *JobSchedulerServer
}

// Do implements the policy.Transporter interface for JobSchedulerServerTransport.
func (j *JobSchedulerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	parts := strings.Split(method, ".")
	method = parts[1] + "." + parts[2]
	return j.dispatchToMethodFake(req, method)
}

func (j *JobSchedulerServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "JobSchedulerClient.CancelItemJobInstance":
		resp, err = j.dispatchCancelItemJobInstance(req)
	case "JobSchedulerClient.CreateItemSchedule":
		resp, err = j.dispatchCreateItemSchedule(req)
	case "JobSchedulerClient.GetItemJobInstance":
		resp, err = j.dispatchGetItemJobInstance(req)
	case "JobSchedulerClient.GetItemSchedule":
		resp, err = j.dispatchGetItemSchedule(req)
	case "JobSchedulerClient.ListItemSchedules":
		resp, err = j.dispatchListItemSchedules(req)
	case "JobSchedulerClient.RunOnDemandItemJob":
		resp, err = j.dispatchRunOnDemandItemJob(req)
	case "JobSchedulerClient.UpdateItemSchedule":
		resp, err = j.dispatchUpdateItemSchedule(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (j *JobSchedulerServerTransport) dispatchCancelItemJobInstance(req *http.Request) (*http.Response, error) {
	if j.srv.CancelItemJobInstance == nil {
		return nil, &nonRetriableError{errors.New("fake for method CancelItemJobInstance not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items/(?P<itemId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/instances/(?P<jobInstanceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cancel`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	itemIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("itemId")])
	if err != nil {
		return nil, err
	}
	jobInstanceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobInstanceId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.CancelItemJobInstance(req.Context(), workspaceIDParam, itemIDParam, jobInstanceIDParam, nil)
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

func (j *JobSchedulerServerTransport) dispatchCreateItemSchedule(req *http.Request) (*http.Response, error) {
	if j.srv.CreateItemSchedule == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateItemSchedule not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items/(?P<itemId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schedules`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[core.CreateScheduleRequest](req)
	if err != nil {
		return nil, err
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	itemIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("itemId")])
	if err != nil {
		return nil, err
	}
	jobTypeParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobType")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.CreateItemSchedule(req.Context(), workspaceIDParam, itemIDParam, jobTypeParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ItemSchedule, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).Location; val != nil {
		resp.Header.Set("Location", *val)
	}
	return resp, nil
}

func (j *JobSchedulerServerTransport) dispatchGetItemJobInstance(req *http.Request) (*http.Response, error) {
	if j.srv.GetItemJobInstance == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetItemJobInstance not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items/(?P<itemId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/instances/(?P<jobInstanceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	itemIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("itemId")])
	if err != nil {
		return nil, err
	}
	jobInstanceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobInstanceId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.GetItemJobInstance(req.Context(), workspaceIDParam, itemIDParam, jobInstanceIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ItemJobInstance, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobSchedulerServerTransport) dispatchGetItemSchedule(req *http.Request) (*http.Response, error) {
	if j.srv.GetItemSchedule == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetItemSchedule not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items/(?P<itemId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schedules/(?P<scheduleId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	itemIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("itemId")])
	if err != nil {
		return nil, err
	}
	jobTypeParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobType")])
	if err != nil {
		return nil, err
	}
	scheduleIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("scheduleId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.GetItemSchedule(req.Context(), workspaceIDParam, itemIDParam, jobTypeParam, scheduleIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ItemSchedule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobSchedulerServerTransport) dispatchListItemSchedules(req *http.Request) (*http.Response, error) {
	if j.srv.ListItemSchedules == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListItemSchedules not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items/(?P<itemId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schedules`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	itemIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("itemId")])
	if err != nil {
		return nil, err
	}
	jobTypeParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobType")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.ListItemSchedules(req.Context(), workspaceIDParam, itemIDParam, jobTypeParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ItemSchedules, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobSchedulerServerTransport) dispatchRunOnDemandItemJob(req *http.Request) (*http.Response, error) {
	if j.srv.RunOnDemandItemJob == nil {
		return nil, &nonRetriableError{errors.New("fake for method RunOnDemandItemJob not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items/(?P<itemId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/instances`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[core.RunOnDemandItemJobRequest](req)
	if err != nil {
		return nil, err
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	itemIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("itemId")])
	if err != nil {
		return nil, err
	}
	jobTypeParam, err := url.QueryUnescape(qp.Get("jobType"))
	if err != nil {
		return nil, err
	}
	var options *core.JobSchedulerClientRunOnDemandItemJobOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &core.JobSchedulerClientRunOnDemandItemJobOptions{
			RunOnDemandItemJobRequest: &body,
		}
	}
	respr, errRespr := j.srv.RunOnDemandItemJob(req.Context(), workspaceIDParam, itemIDParam, jobTypeParam, options)
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

func (j *JobSchedulerServerTransport) dispatchUpdateItemSchedule(req *http.Request) (*http.Response, error) {
	if j.srv.UpdateItemSchedule == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateItemSchedule not implemented")}
	}
	const regexStr = `/v1/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items/(?P<itemId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schedules/(?P<scheduleId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[core.UpdateScheduleRequest](req)
	if err != nil {
		return nil, err
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	itemIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("itemId")])
	if err != nil {
		return nil, err
	}
	jobTypeParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobType")])
	if err != nil {
		return nil, err
	}
	scheduleIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("scheduleId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.UpdateItemSchedule(req.Context(), workspaceIDParam, itemIDParam, jobTypeParam, scheduleIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ItemSchedule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
