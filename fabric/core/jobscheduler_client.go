// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package core

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// JobSchedulerClient contains the methods for the JobScheduler group.
// Don't use this type directly, use a constructor function instead.
type JobSchedulerClient struct {
	internal *azcore.Client
	endpoint string
}

// CancelItemJobInstance - REQUIRED DELEGATED SCOPES For item APIs use these scope types:
// * Generic scope: Item.Execute.All
//
// * Specific scope: itemType.Execute.All (for example: Notebook.Execute.All)
//
// for more information about scopes, see scopes article [/rest/api/fabric/articles/scopes].
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | No | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | No |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - itemID - The item ID.
//   - jobInstanceID - The job instance ID.
//   - options - JobSchedulerClientCancelItemJobInstanceOptions contains the optional parameters for the JobSchedulerClient.CancelItemJobInstance
//     method.
func (client *JobSchedulerClient) CancelItemJobInstance(ctx context.Context, workspaceID string, itemID string, jobInstanceID string, options *JobSchedulerClientCancelItemJobInstanceOptions) (JobSchedulerClientCancelItemJobInstanceResponse, error) {
	var err error
	const operationName = "core.JobSchedulerClient.CancelItemJobInstance"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.cancelItemJobInstanceCreateRequest(ctx, workspaceID, itemID, jobInstanceID, options)
	if err != nil {
		return JobSchedulerClientCancelItemJobInstanceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JobSchedulerClientCancelItemJobInstanceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = NewResponseError(httpResp)
		return JobSchedulerClientCancelItemJobInstanceResponse{}, err
	}
	resp, err := client.cancelItemJobInstanceHandleResponse(httpResp)
	return resp, err
}

// cancelItemJobInstanceCreateRequest creates the CancelItemJobInstance request.
func (client *JobSchedulerClient) cancelItemJobInstanceCreateRequest(ctx context.Context, workspaceID string, itemID string, jobInstanceID string, options *JobSchedulerClientCancelItemJobInstanceOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/items/{itemId}/jobs/instances/{jobInstanceId}/cancel"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if itemID == "" {
		return nil, errors.New("parameter itemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{itemId}", url.PathEscape(itemID))
	if jobInstanceID == "" {
		return nil, errors.New("parameter jobInstanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobInstanceId}", url.PathEscape(jobInstanceID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// cancelItemJobInstanceHandleResponse handles the CancelItemJobInstance response.
func (client *JobSchedulerClient) cancelItemJobInstanceHandleResponse(resp *http.Response) (JobSchedulerClientCancelItemJobInstanceResponse, error) {
	result := JobSchedulerClientCancelItemJobInstanceResponse{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return JobSchedulerClientCancelItemJobInstanceResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	return result, nil
}

// CreateItemSchedule - REQUIRED DELEGATED SCOPES: Item.Execute.All and Item.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | No | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | No |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - itemID - The item ID.
//   - jobType - The job type.
//   - createScheduleRequest - A item schedule create request.
//   - options - JobSchedulerClientCreateItemScheduleOptions contains the optional parameters for the JobSchedulerClient.CreateItemSchedule
//     method.
func (client *JobSchedulerClient) CreateItemSchedule(ctx context.Context, workspaceID string, itemID string, jobType string, createScheduleRequest CreateScheduleRequest, options *JobSchedulerClientCreateItemScheduleOptions) (JobSchedulerClientCreateItemScheduleResponse, error) {
	var err error
	const operationName = "core.JobSchedulerClient.CreateItemSchedule"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createItemScheduleCreateRequest(ctx, workspaceID, itemID, jobType, createScheduleRequest, options)
	if err != nil {
		return JobSchedulerClientCreateItemScheduleResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JobSchedulerClientCreateItemScheduleResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = NewResponseError(httpResp)
		return JobSchedulerClientCreateItemScheduleResponse{}, err
	}
	resp, err := client.createItemScheduleHandleResponse(httpResp)
	return resp, err
}

// createItemScheduleCreateRequest creates the CreateItemSchedule request.
func (client *JobSchedulerClient) createItemScheduleCreateRequest(ctx context.Context, workspaceID string, itemID string, jobType string, createScheduleRequest CreateScheduleRequest, options *JobSchedulerClientCreateItemScheduleOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/items/{itemId}/jobs/{jobType}/schedules"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if itemID == "" {
		return nil, errors.New("parameter itemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{itemId}", url.PathEscape(itemID))
	if jobType == "" {
		return nil, errors.New("parameter jobType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobType}", url.PathEscape(jobType))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, createScheduleRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// createItemScheduleHandleResponse handles the CreateItemSchedule response.
func (client *JobSchedulerClient) createItemScheduleHandleResponse(resp *http.Response) (JobSchedulerClientCreateItemScheduleResponse, error) {
	result := JobSchedulerClientCreateItemScheduleResponse{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ItemSchedule); err != nil {
		return JobSchedulerClientCreateItemScheduleResponse{}, err
	}
	return result, nil
}

// GetItemJobInstance - REQUIRED DELEGATED SCOPES For item APIs use these scope types:
// * Generic scope: Item.ReadWrite.All or Item.Read.All
//
// * Specific scope: itemType.ReadWrite.All or itemType.Read.All (for example: Notebook.ReadWrite.All)
//
// for more information about scopes, see scopes article [/rest/api/fabric/articles/scopes].
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | No | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | No |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - itemID - The item ID.
//   - jobInstanceID - The job instance ID.
//   - options - JobSchedulerClientGetItemJobInstanceOptions contains the optional parameters for the JobSchedulerClient.GetItemJobInstance
//     method.
func (client *JobSchedulerClient) GetItemJobInstance(ctx context.Context, workspaceID string, itemID string, jobInstanceID string, options *JobSchedulerClientGetItemJobInstanceOptions) (JobSchedulerClientGetItemJobInstanceResponse, error) {
	var err error
	const operationName = "core.JobSchedulerClient.GetItemJobInstance"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getItemJobInstanceCreateRequest(ctx, workspaceID, itemID, jobInstanceID, options)
	if err != nil {
		return JobSchedulerClientGetItemJobInstanceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JobSchedulerClientGetItemJobInstanceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = NewResponseError(httpResp)
		return JobSchedulerClientGetItemJobInstanceResponse{}, err
	}
	resp, err := client.getItemJobInstanceHandleResponse(httpResp)
	return resp, err
}

// getItemJobInstanceCreateRequest creates the GetItemJobInstance request.
func (client *JobSchedulerClient) getItemJobInstanceCreateRequest(ctx context.Context, workspaceID string, itemID string, jobInstanceID string, options *JobSchedulerClientGetItemJobInstanceOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/items/{itemId}/jobs/instances/{jobInstanceId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if itemID == "" {
		return nil, errors.New("parameter itemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{itemId}", url.PathEscape(itemID))
	if jobInstanceID == "" {
		return nil, errors.New("parameter jobInstanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobInstanceId}", url.PathEscape(jobInstanceID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getItemJobInstanceHandleResponse handles the GetItemJobInstance response.
func (client *JobSchedulerClient) getItemJobInstanceHandleResponse(resp *http.Response) (JobSchedulerClientGetItemJobInstanceResponse, error) {
	result := JobSchedulerClientGetItemJobInstanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ItemJobInstance); err != nil {
		return JobSchedulerClientGetItemJobInstanceResponse{}, err
	}
	return result, nil
}

// GetItemSchedule - REQUIRED DELEGATED SCOPES For item APIs use these scope types:
// * Generic scope: Item.Read.All or Item.ReadWrite.All
//
// * Specific scope: itemType.Read.All or itemType.ReadWrite.All (for example: Notebook.Read.All)
//
// for more information about scopes, see scopes article [/rest/api/fabric/articles/scopes].
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | No | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | No |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - itemID - The item ID.
//   - jobType - The job type.
//   - scheduleID - The item schedule ID.
//   - options - JobSchedulerClientGetItemScheduleOptions contains the optional parameters for the JobSchedulerClient.GetItemSchedule
//     method.
func (client *JobSchedulerClient) GetItemSchedule(ctx context.Context, workspaceID string, itemID string, jobType string, scheduleID string, options *JobSchedulerClientGetItemScheduleOptions) (JobSchedulerClientGetItemScheduleResponse, error) {
	var err error
	const operationName = "core.JobSchedulerClient.GetItemSchedule"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getItemScheduleCreateRequest(ctx, workspaceID, itemID, jobType, scheduleID, options)
	if err != nil {
		return JobSchedulerClientGetItemScheduleResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JobSchedulerClientGetItemScheduleResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = NewResponseError(httpResp)
		return JobSchedulerClientGetItemScheduleResponse{}, err
	}
	resp, err := client.getItemScheduleHandleResponse(httpResp)
	return resp, err
}

// getItemScheduleCreateRequest creates the GetItemSchedule request.
func (client *JobSchedulerClient) getItemScheduleCreateRequest(ctx context.Context, workspaceID string, itemID string, jobType string, scheduleID string, options *JobSchedulerClientGetItemScheduleOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/items/{itemId}/jobs/{jobType}/schedules/{scheduleId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if itemID == "" {
		return nil, errors.New("parameter itemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{itemId}", url.PathEscape(itemID))
	if jobType == "" {
		return nil, errors.New("parameter jobType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobType}", url.PathEscape(jobType))
	if scheduleID == "" {
		return nil, errors.New("parameter scheduleID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleId}", url.PathEscape(scheduleID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getItemScheduleHandleResponse handles the GetItemSchedule response.
func (client *JobSchedulerClient) getItemScheduleHandleResponse(resp *http.Response) (JobSchedulerClientGetItemScheduleResponse, error) {
	result := JobSchedulerClientGetItemScheduleResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ItemSchedule); err != nil {
		return JobSchedulerClientGetItemScheduleResponse{}, err
	}
	return result, nil
}

// ListItemSchedules - This API supports pagination [/rest/api/fabric/articles/pagination].
// REQUIRED DELEGATED SCOPES For item APIs use these scope types:
// * Generic scope: Item.ReadWrite.All or Item.Read.All
//
// * Specific scope: itemType.ReadWrite.All or itemType.Read.All (for example: Notebook.ReadWrite.All)
//
// for more information about scopes, see scopes article [/rest/api/fabric/articles/scopes].
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | No | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | No |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - itemID - The item ID.
//   - jobType - The job type.
//   - options - JobSchedulerClientListItemSchedulesOptions contains the optional parameters for the JobSchedulerClient.ListItemSchedules
//     method.
func (client *JobSchedulerClient) ListItemSchedules(ctx context.Context, workspaceID string, itemID string, jobType string, options *JobSchedulerClientListItemSchedulesOptions) (JobSchedulerClientListItemSchedulesResponse, error) {
	var err error
	const operationName = "core.JobSchedulerClient.ListItemSchedules"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listItemSchedulesCreateRequest(ctx, workspaceID, itemID, jobType, options)
	if err != nil {
		return JobSchedulerClientListItemSchedulesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JobSchedulerClientListItemSchedulesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = NewResponseError(httpResp)
		return JobSchedulerClientListItemSchedulesResponse{}, err
	}
	resp, err := client.listItemSchedulesHandleResponse(httpResp)
	return resp, err
}

// listItemSchedulesCreateRequest creates the ListItemSchedules request.
func (client *JobSchedulerClient) listItemSchedulesCreateRequest(ctx context.Context, workspaceID string, itemID string, jobType string, options *JobSchedulerClientListItemSchedulesOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/items/{itemId}/jobs/{jobType}/schedules"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if itemID == "" {
		return nil, errors.New("parameter itemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{itemId}", url.PathEscape(itemID))
	if jobType == "" {
		return nil, errors.New("parameter jobType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobType}", url.PathEscape(jobType))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listItemSchedulesHandleResponse handles the ListItemSchedules response.
func (client *JobSchedulerClient) listItemSchedulesHandleResponse(resp *http.Response) (JobSchedulerClientListItemSchedulesResponse, error) {
	result := JobSchedulerClientListItemSchedulesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ItemSchedules); err != nil {
		return JobSchedulerClientListItemSchedulesResponse{}, err
	}
	return result, nil
}

// RunOnDemandItemJob - REQUIRED DELEGATED SCOPES For item APIs use these scope types:
// * Generic scope: Item.Execute.All
//
// * Specific scope: itemType.Execute.All (for example: Notebook.Execute.All)
//
// for more information about scopes, see: scopes article [/rest/api/fabric/articles/scopes].
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | No | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | No |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - itemID - The item ID.
//   - jobType - Job type
//   - options - JobSchedulerClientRunOnDemandItemJobOptions contains the optional parameters for the JobSchedulerClient.RunOnDemandItemJob
//     method.
func (client *JobSchedulerClient) RunOnDemandItemJob(ctx context.Context, workspaceID string, itemID string, jobType string, options *JobSchedulerClientRunOnDemandItemJobOptions) (JobSchedulerClientRunOnDemandItemJobResponse, error) {
	var err error
	const operationName = "core.JobSchedulerClient.RunOnDemandItemJob"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.runOnDemandItemJobCreateRequest(ctx, workspaceID, itemID, jobType, options)
	if err != nil {
		return JobSchedulerClientRunOnDemandItemJobResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JobSchedulerClientRunOnDemandItemJobResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = NewResponseError(httpResp)
		return JobSchedulerClientRunOnDemandItemJobResponse{}, err
	}
	resp, err := client.runOnDemandItemJobHandleResponse(httpResp)
	return resp, err
}

// runOnDemandItemJobCreateRequest creates the RunOnDemandItemJob request.
func (client *JobSchedulerClient) runOnDemandItemJobCreateRequest(ctx context.Context, workspaceID string, itemID string, jobType string, options *JobSchedulerClientRunOnDemandItemJobOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/items/{itemId}/jobs/instances"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if itemID == "" {
		return nil, errors.New("parameter itemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{itemId}", url.PathEscape(itemID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("jobType", jobType)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.RunOnDemandItemJobRequest != nil {
		if err := runtime.MarshalAsJSON(req, *options.RunOnDemandItemJobRequest); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// runOnDemandItemJobHandleResponse handles the RunOnDemandItemJob response.
func (client *JobSchedulerClient) runOnDemandItemJobHandleResponse(resp *http.Response) (JobSchedulerClientRunOnDemandItemJobResponse, error) {
	result := JobSchedulerClientRunOnDemandItemJobResponse{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return JobSchedulerClientRunOnDemandItemJobResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	return result, nil
}

// UpdateItemSchedule - REQUIRED DELEGATED SCOPES: Item.Execute.All and Item.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | No | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | No |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - itemID - The item ID.
//   - jobType - The job type.
//   - scheduleID - The item schedule ID.
//   - updateScheduleRequest - A item schedule update request.
//   - options - JobSchedulerClientUpdateItemScheduleOptions contains the optional parameters for the JobSchedulerClient.UpdateItemSchedule
//     method.
func (client *JobSchedulerClient) UpdateItemSchedule(ctx context.Context, workspaceID string, itemID string, jobType string, scheduleID string, updateScheduleRequest UpdateScheduleRequest, options *JobSchedulerClientUpdateItemScheduleOptions) (JobSchedulerClientUpdateItemScheduleResponse, error) {
	var err error
	const operationName = "core.JobSchedulerClient.UpdateItemSchedule"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateItemScheduleCreateRequest(ctx, workspaceID, itemID, jobType, scheduleID, updateScheduleRequest, options)
	if err != nil {
		return JobSchedulerClientUpdateItemScheduleResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JobSchedulerClientUpdateItemScheduleResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = NewResponseError(httpResp)
		return JobSchedulerClientUpdateItemScheduleResponse{}, err
	}
	resp, err := client.updateItemScheduleHandleResponse(httpResp)
	return resp, err
}

// updateItemScheduleCreateRequest creates the UpdateItemSchedule request.
func (client *JobSchedulerClient) updateItemScheduleCreateRequest(ctx context.Context, workspaceID string, itemID string, jobType string, scheduleID string, updateScheduleRequest UpdateScheduleRequest, options *JobSchedulerClientUpdateItemScheduleOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/items/{itemId}/jobs/{jobType}/schedules/{scheduleId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if itemID == "" {
		return nil, errors.New("parameter itemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{itemId}", url.PathEscape(itemID))
	if jobType == "" {
		return nil, errors.New("parameter jobType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobType}", url.PathEscape(jobType))
	if scheduleID == "" {
		return nil, errors.New("parameter scheduleID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleId}", url.PathEscape(scheduleID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, updateScheduleRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// updateItemScheduleHandleResponse handles the UpdateItemSchedule response.
func (client *JobSchedulerClient) updateItemScheduleHandleResponse(resp *http.Response) (JobSchedulerClientUpdateItemScheduleResponse, error) {
	result := JobSchedulerClientUpdateItemScheduleResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ItemSchedule); err != nil {
		return JobSchedulerClientUpdateItemScheduleResponse{}, err
	}
	return result, nil
}

// Custom code starts below