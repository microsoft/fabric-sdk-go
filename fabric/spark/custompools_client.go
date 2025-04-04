// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package spark

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"

	"github.com/microsoft/fabric-sdk-go/fabric/core"
	"github.com/microsoft/fabric-sdk-go/internal/iruntime"
)

// CustomPoolsClient contains the methods for the CustomPools group.
// Don't use this type directly, use a constructor function instead.
type CustomPoolsClient struct {
	internal *azcore.Client
	endpoint string
}

// CreateWorkspaceCustomPool - PERMISSIONS The caller must have admin workspace role.
// REQUIRED DELEGATED SCOPES Workspace.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - createCustomPoolRequest - Create custom pool request payload.
//   - options - CustomPoolsClientCreateWorkspaceCustomPoolOptions contains the optional parameters for the CustomPoolsClient.CreateWorkspaceCustomPool
//     method.
func (client *CustomPoolsClient) CreateWorkspaceCustomPool(ctx context.Context, workspaceID string, createCustomPoolRequest CreateCustomPoolRequest, options *CustomPoolsClientCreateWorkspaceCustomPoolOptions) (CustomPoolsClientCreateWorkspaceCustomPoolResponse, error) {
	var err error
	const operationName = "spark.CustomPoolsClient.CreateWorkspaceCustomPool"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createWorkspaceCustomPoolCreateRequest(ctx, workspaceID, createCustomPoolRequest, options)
	if err != nil {
		return CustomPoolsClientCreateWorkspaceCustomPoolResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomPoolsClientCreateWorkspaceCustomPoolResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = core.NewResponseError(httpResp)
		return CustomPoolsClientCreateWorkspaceCustomPoolResponse{}, err
	}
	resp, err := client.createWorkspaceCustomPoolHandleResponse(httpResp)
	return resp, err
}

// createWorkspaceCustomPoolCreateRequest creates the CreateWorkspaceCustomPool request.
func (client *CustomPoolsClient) createWorkspaceCustomPoolCreateRequest(ctx context.Context, workspaceID string, createCustomPoolRequest CreateCustomPoolRequest, _ *CustomPoolsClientCreateWorkspaceCustomPoolOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/spark/pools"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, createCustomPoolRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// createWorkspaceCustomPoolHandleResponse handles the CreateWorkspaceCustomPool response.
func (client *CustomPoolsClient) createWorkspaceCustomPoolHandleResponse(resp *http.Response) (CustomPoolsClientCreateWorkspaceCustomPoolResponse, error) {
	result := CustomPoolsClientCreateWorkspaceCustomPoolResponse{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomPool); err != nil {
		return CustomPoolsClientCreateWorkspaceCustomPoolResponse{}, err
	}
	return result, nil
}

// DeleteWorkspaceCustomPool - PERMISSIONS The caller must have admin workspace role.
// REQUIRED DELEGATED SCOPES Workspace.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - poolID - The custom pool ID.
//   - options - CustomPoolsClientDeleteWorkspaceCustomPoolOptions contains the optional parameters for the CustomPoolsClient.DeleteWorkspaceCustomPool
//     method.
func (client *CustomPoolsClient) DeleteWorkspaceCustomPool(ctx context.Context, workspaceID string, poolID string, options *CustomPoolsClientDeleteWorkspaceCustomPoolOptions) (CustomPoolsClientDeleteWorkspaceCustomPoolResponse, error) {
	var err error
	const operationName = "spark.CustomPoolsClient.DeleteWorkspaceCustomPool"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteWorkspaceCustomPoolCreateRequest(ctx, workspaceID, poolID, options)
	if err != nil {
		return CustomPoolsClientDeleteWorkspaceCustomPoolResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomPoolsClientDeleteWorkspaceCustomPoolResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return CustomPoolsClientDeleteWorkspaceCustomPoolResponse{}, err
	}
	return CustomPoolsClientDeleteWorkspaceCustomPoolResponse{}, nil
}

// deleteWorkspaceCustomPoolCreateRequest creates the DeleteWorkspaceCustomPool request.
func (client *CustomPoolsClient) deleteWorkspaceCustomPoolCreateRequest(ctx context.Context, workspaceID string, poolID string, _ *CustomPoolsClientDeleteWorkspaceCustomPoolOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/spark/pools/{poolId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if poolID == "" {
		return nil, errors.New("parameter poolID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolId}", url.PathEscape(poolID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetWorkspaceCustomPool - PERMISSIONS The caller must have viewer or higher workspace role.
// REQUIRED DELEGATED SCOPES Workspace.Read.All or Workspace.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - poolID - The custom pool ID.
//   - options - CustomPoolsClientGetWorkspaceCustomPoolOptions contains the optional parameters for the CustomPoolsClient.GetWorkspaceCustomPool
//     method.
func (client *CustomPoolsClient) GetWorkspaceCustomPool(ctx context.Context, workspaceID string, poolID string, options *CustomPoolsClientGetWorkspaceCustomPoolOptions) (CustomPoolsClientGetWorkspaceCustomPoolResponse, error) {
	var err error
	const operationName = "spark.CustomPoolsClient.GetWorkspaceCustomPool"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getWorkspaceCustomPoolCreateRequest(ctx, workspaceID, poolID, options)
	if err != nil {
		return CustomPoolsClientGetWorkspaceCustomPoolResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomPoolsClientGetWorkspaceCustomPoolResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return CustomPoolsClientGetWorkspaceCustomPoolResponse{}, err
	}
	resp, err := client.getWorkspaceCustomPoolHandleResponse(httpResp)
	return resp, err
}

// getWorkspaceCustomPoolCreateRequest creates the GetWorkspaceCustomPool request.
func (client *CustomPoolsClient) getWorkspaceCustomPoolCreateRequest(ctx context.Context, workspaceID string, poolID string, _ *CustomPoolsClientGetWorkspaceCustomPoolOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/spark/pools/{poolId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if poolID == "" {
		return nil, errors.New("parameter poolID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolId}", url.PathEscape(poolID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getWorkspaceCustomPoolHandleResponse handles the GetWorkspaceCustomPool response.
func (client *CustomPoolsClient) getWorkspaceCustomPoolHandleResponse(resp *http.Response) (CustomPoolsClientGetWorkspaceCustomPoolResponse, error) {
	result := CustomPoolsClientGetWorkspaceCustomPoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomPool); err != nil {
		return CustomPoolsClientGetWorkspaceCustomPoolResponse{}, err
	}
	return result, nil
}

// NewListWorkspaceCustomPoolsPager - PERMISSIONS The caller must have viewer or higher workspace role.
// REQUIRED DELEGATED SCOPES Workspace.Read.All or Workspace.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - options - CustomPoolsClientListWorkspaceCustomPoolsOptions contains the optional parameters for the CustomPoolsClient.NewListWorkspaceCustomPoolsPager
//     method.
func (client *CustomPoolsClient) NewListWorkspaceCustomPoolsPager(workspaceID string, options *CustomPoolsClientListWorkspaceCustomPoolsOptions) *runtime.Pager[CustomPoolsClientListWorkspaceCustomPoolsResponse] {
	return runtime.NewPager(runtime.PagingHandler[CustomPoolsClientListWorkspaceCustomPoolsResponse]{
		More: func(page CustomPoolsClientListWorkspaceCustomPoolsResponse) bool {
			return page.ContinuationURI != nil && len(*page.ContinuationURI) > 0
		},
		Fetcher: func(ctx context.Context, page *CustomPoolsClientListWorkspaceCustomPoolsResponse) (CustomPoolsClientListWorkspaceCustomPoolsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "spark.CustomPoolsClient.NewListWorkspaceCustomPoolsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ContinuationURI
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listWorkspaceCustomPoolsCreateRequest(ctx, workspaceID, options)
			}, nil)
			if err != nil {
				return CustomPoolsClientListWorkspaceCustomPoolsResponse{}, err
			}
			return client.listWorkspaceCustomPoolsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listWorkspaceCustomPoolsCreateRequest creates the ListWorkspaceCustomPools request.
func (client *CustomPoolsClient) listWorkspaceCustomPoolsCreateRequest(ctx context.Context, workspaceID string, options *CustomPoolsClientListWorkspaceCustomPoolsOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/spark/pools"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listWorkspaceCustomPoolsHandleResponse handles the ListWorkspaceCustomPools response.
func (client *CustomPoolsClient) listWorkspaceCustomPoolsHandleResponse(resp *http.Response) (CustomPoolsClientListWorkspaceCustomPoolsResponse, error) {
	result := CustomPoolsClientListWorkspaceCustomPoolsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomPools); err != nil {
		return CustomPoolsClientListWorkspaceCustomPoolsResponse{}, err
	}
	return result, nil
}

// UpdateWorkspaceCustomPool - PERMISSIONS The caller must have admin workspace role.
// REQUIRED DELEGATED SCOPES Workspace.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - poolID - The custom pool ID.
//   - updateCustomPoolRequest - Update custom pool request payload.
//   - options - CustomPoolsClientUpdateWorkspaceCustomPoolOptions contains the optional parameters for the CustomPoolsClient.UpdateWorkspaceCustomPool
//     method.
func (client *CustomPoolsClient) UpdateWorkspaceCustomPool(ctx context.Context, workspaceID string, poolID string, updateCustomPoolRequest UpdateCustomPoolRequest, options *CustomPoolsClientUpdateWorkspaceCustomPoolOptions) (CustomPoolsClientUpdateWorkspaceCustomPoolResponse, error) {
	var err error
	const operationName = "spark.CustomPoolsClient.UpdateWorkspaceCustomPool"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateWorkspaceCustomPoolCreateRequest(ctx, workspaceID, poolID, updateCustomPoolRequest, options)
	if err != nil {
		return CustomPoolsClientUpdateWorkspaceCustomPoolResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomPoolsClientUpdateWorkspaceCustomPoolResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return CustomPoolsClientUpdateWorkspaceCustomPoolResponse{}, err
	}
	resp, err := client.updateWorkspaceCustomPoolHandleResponse(httpResp)
	return resp, err
}

// updateWorkspaceCustomPoolCreateRequest creates the UpdateWorkspaceCustomPool request.
func (client *CustomPoolsClient) updateWorkspaceCustomPoolCreateRequest(ctx context.Context, workspaceID string, poolID string, updateCustomPoolRequest UpdateCustomPoolRequest, _ *CustomPoolsClientUpdateWorkspaceCustomPoolOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/spark/pools/{poolId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if poolID == "" {
		return nil, errors.New("parameter poolID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolId}", url.PathEscape(poolID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, updateCustomPoolRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// updateWorkspaceCustomPoolHandleResponse handles the UpdateWorkspaceCustomPool response.
func (client *CustomPoolsClient) updateWorkspaceCustomPoolHandleResponse(resp *http.Response) (CustomPoolsClientUpdateWorkspaceCustomPoolResponse, error) {
	result := CustomPoolsClientUpdateWorkspaceCustomPoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomPool); err != nil {
		return CustomPoolsClientUpdateWorkspaceCustomPoolResponse{}, err
	}
	return result, nil
}

// Custom code starts below

// ListWorkspaceCustomPools - returns array of CustomPool from all pages.
// PERMISSIONS The caller must have viewer or higher workspace role.
//
// # REQUIRED DELEGATED SCOPES Workspace.Read.All or Workspace.ReadWrite.All
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
//
// INTERFACE
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - options - CustomPoolsClientListWorkspaceCustomPoolsOptions contains the optional parameters for the CustomPoolsClient.NewListWorkspaceCustomPoolsPager method.
func (client *CustomPoolsClient) ListWorkspaceCustomPools(ctx context.Context, workspaceID string, options *CustomPoolsClientListWorkspaceCustomPoolsOptions) ([]CustomPool, error) {
	pager := client.NewListWorkspaceCustomPoolsPager(workspaceID, options)
	mapper := func(resp CustomPoolsClientListWorkspaceCustomPoolsResponse) []CustomPool {
		return resp.Value
	}
	list, err := iruntime.NewPageIterator(ctx, pager, mapper).Get()
	if err != nil {
		var azcoreRespError *azcore.ResponseError
		if errors.As(err, &azcoreRespError) {
			return []CustomPool{}, core.NewResponseError(azcoreRespError.RawResponse)
		}
		return []CustomPool{}, err
	}
	return list, nil
}
