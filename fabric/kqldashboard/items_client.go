// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package kqldashboard

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

	"github.com/microsoft/fabric-sdk-go/fabric/core"
	"github.com/microsoft/fabric-sdk-go/internal/iruntime"
)

// ItemsClient contains the methods for the Items group.
// Don't use this type directly, use a constructor function instead.
type ItemsClient struct {
	internal *azcore.Client
	endpoint string
}

// CreateKQLDashboard - To create a KQL dashboard with definition, refer to the KQL dashboard definition article [/rest/api/fabric/articles/item-management/definitions/kql-dashboard-definition].
// PERMISSIONS
// The caller must have contributor or higher workspace role.
// REQUIRED DELEGATED SCOPES KQLDashboard.ReadWrite.All or Item.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | Yes | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - createKQLDashboardRequest - Create KQL dashboard definition request payload.
//   - options - ItemsClientCreateKQLDashboardOptions contains the optional parameters for the ItemsClient.CreateKQLDashboard
//     method.
func (client *ItemsClient) CreateKQLDashboard(ctx context.Context, workspaceID string, createKQLDashboardRequest CreateKQLDashboardRequest, options *ItemsClientCreateKQLDashboardOptions) (ItemsClientCreateKQLDashboardResponse, error) {
	var err error
	const operationName = "kqldashboard.ItemsClient.CreateKQLDashboard"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createKQLDashboardCreateRequest(ctx, workspaceID, createKQLDashboardRequest, options)
	if err != nil {
		return ItemsClientCreateKQLDashboardResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ItemsClientCreateKQLDashboardResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = core.NewResponseError(httpResp)
		return ItemsClientCreateKQLDashboardResponse{}, err
	}
	resp, err := client.createKQLDashboardHandleResponse(httpResp)
	return resp, err
}

// createKQLDashboardCreateRequest creates the CreateKQLDashboard request.
func (client *ItemsClient) createKQLDashboardCreateRequest(ctx context.Context, workspaceID string, createKQLDashboardRequest CreateKQLDashboardRequest, _ *ItemsClientCreateKQLDashboardOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/kqlDashboards"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, createKQLDashboardRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// createKQLDashboardHandleResponse handles the CreateKQLDashboard response.
func (client *ItemsClient) createKQLDashboardHandleResponse(resp *http.Response) (ItemsClientCreateKQLDashboardResponse, error) {
	result := ItemsClientCreateKQLDashboardResponse{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.KQLDashboard); err != nil {
		return ItemsClientCreateKQLDashboardResponse{}, err
	}
	return result, nil
}

// DeleteKQLDashboard - PERMISSIONS The caller must have contributor or higher workspace role.
// REQUIRED DELEGATED SCOPES KQLDashboard.ReadWrite.All or Item.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | Yes | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - kqlDashboardID - The KQL dashboard ID.
//   - options - ItemsClientDeleteKQLDashboardOptions contains the optional parameters for the ItemsClient.DeleteKQLDashboard
//     method.
func (client *ItemsClient) DeleteKQLDashboard(ctx context.Context, workspaceID string, kqlDashboardID string, options *ItemsClientDeleteKQLDashboardOptions) (ItemsClientDeleteKQLDashboardResponse, error) {
	var err error
	const operationName = "kqldashboard.ItemsClient.DeleteKQLDashboard"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteKQLDashboardCreateRequest(ctx, workspaceID, kqlDashboardID, options)
	if err != nil {
		return ItemsClientDeleteKQLDashboardResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ItemsClientDeleteKQLDashboardResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return ItemsClientDeleteKQLDashboardResponse{}, err
	}
	return ItemsClientDeleteKQLDashboardResponse{}, nil
}

// deleteKQLDashboardCreateRequest creates the DeleteKQLDashboard request.
func (client *ItemsClient) deleteKQLDashboardCreateRequest(ctx context.Context, workspaceID string, kqlDashboardID string, _ *ItemsClientDeleteKQLDashboardOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/kqlDashboards/{kqlDashboardId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if kqlDashboardID == "" {
		return nil, errors.New("parameter kqlDashboardID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kqlDashboardId}", url.PathEscape(kqlDashboardID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetKQLDashboard - PERMISSIONS The caller must have viewer or higher workspace role.
// REQUIRED DELEGATED SCOPES KQLDashboard.Read.All or KQLDashboard.ReadWrite.All or Item.Read.All or Item.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | Yes | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - kqlDashboardID - The KQL dashboard ID.
//   - options - ItemsClientGetKQLDashboardOptions contains the optional parameters for the ItemsClient.GetKQLDashboard method.
func (client *ItemsClient) GetKQLDashboard(ctx context.Context, workspaceID string, kqlDashboardID string, options *ItemsClientGetKQLDashboardOptions) (ItemsClientGetKQLDashboardResponse, error) {
	var err error
	const operationName = "kqldashboard.ItemsClient.GetKQLDashboard"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getKQLDashboardCreateRequest(ctx, workspaceID, kqlDashboardID, options)
	if err != nil {
		return ItemsClientGetKQLDashboardResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ItemsClientGetKQLDashboardResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return ItemsClientGetKQLDashboardResponse{}, err
	}
	resp, err := client.getKQLDashboardHandleResponse(httpResp)
	return resp, err
}

// getKQLDashboardCreateRequest creates the GetKQLDashboard request.
func (client *ItemsClient) getKQLDashboardCreateRequest(ctx context.Context, workspaceID string, kqlDashboardID string, _ *ItemsClientGetKQLDashboardOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/kqlDashboards/{kqlDashboardId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if kqlDashboardID == "" {
		return nil, errors.New("parameter kqlDashboardID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kqlDashboardId}", url.PathEscape(kqlDashboardID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getKQLDashboardHandleResponse handles the GetKQLDashboard response.
func (client *ItemsClient) getKQLDashboardHandleResponse(resp *http.Response) (ItemsClientGetKQLDashboardResponse, error) {
	result := ItemsClientGetKQLDashboardResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KQLDashboard); err != nil {
		return ItemsClientGetKQLDashboardResponse{}, err
	}
	return result, nil
}

// GetKQLDashboardDefinition - PERMISSIONS The caller must have contributor or higher workspace role.
// REQUIRED DELEGATED SCOPES KQLDashboard.ReadWrite.All or Item.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | Yes | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - kqlDashboardID - The KQL dashboard ID.
//   - options - ItemsClientGetKQLDashboardDefinitionOptions contains the optional parameters for the ItemsClient.GetKQLDashboardDefinition
//     method.
func (client *ItemsClient) GetKQLDashboardDefinition(ctx context.Context, workspaceID string, kqlDashboardID string, options *ItemsClientGetKQLDashboardDefinitionOptions) (ItemsClientGetKQLDashboardDefinitionResponse, error) {
	var err error
	const operationName = "kqldashboard.ItemsClient.GetKQLDashboardDefinition"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getKQLDashboardDefinitionCreateRequest(ctx, workspaceID, kqlDashboardID, options)
	if err != nil {
		return ItemsClientGetKQLDashboardDefinitionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ItemsClientGetKQLDashboardDefinitionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return ItemsClientGetKQLDashboardDefinitionResponse{}, err
	}
	resp, err := client.getKQLDashboardDefinitionHandleResponse(httpResp)
	return resp, err
}

// getKQLDashboardDefinitionCreateRequest creates the GetKQLDashboardDefinition request.
func (client *ItemsClient) getKQLDashboardDefinitionCreateRequest(ctx context.Context, workspaceID string, kqlDashboardID string, options *ItemsClientGetKQLDashboardDefinitionOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/kqlDashboards/{kqlDashboardId}/getDefinition"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if kqlDashboardID == "" {
		return nil, errors.New("parameter kqlDashboardID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kqlDashboardId}", url.PathEscape(kqlDashboardID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Format != nil {
		reqQP.Set("format", *options.Format)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getKQLDashboardDefinitionHandleResponse handles the GetKQLDashboardDefinition response.
func (client *ItemsClient) getKQLDashboardDefinitionHandleResponse(resp *http.Response) (ItemsClientGetKQLDashboardDefinitionResponse, error) {
	result := ItemsClientGetKQLDashboardDefinitionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DefinitionResponse); err != nil {
		return ItemsClientGetKQLDashboardDefinitionResponse{}, err
	}
	return result, nil
}

// NewListKQLDashboardsPager - This API supports pagination [/rest/api/fabric/articles/pagination].
// PERMISSIONS The caller must have viewer or higher workspace role.
// REQUIRED DELEGATED SCOPES Workspace.Read.All or Workspace.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | Yes | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - options - ItemsClientListKQLDashboardsOptions contains the optional parameters for the ItemsClient.NewListKQLDashboardsPager
//     method.
func (client *ItemsClient) NewListKQLDashboardsPager(workspaceID string, options *ItemsClientListKQLDashboardsOptions) *runtime.Pager[ItemsClientListKQLDashboardsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ItemsClientListKQLDashboardsResponse]{
		More: func(page ItemsClientListKQLDashboardsResponse) bool {
			return page.ContinuationURI != nil && len(*page.ContinuationURI) > 0
		},
		Fetcher: func(ctx context.Context, page *ItemsClientListKQLDashboardsResponse) (ItemsClientListKQLDashboardsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "kqldashboard.ItemsClient.NewListKQLDashboardsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ContinuationURI
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listKQLDashboardsCreateRequest(ctx, workspaceID, options)
			}, nil)
			if err != nil {
				return ItemsClientListKQLDashboardsResponse{}, err
			}
			return client.listKQLDashboardsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listKQLDashboardsCreateRequest creates the ListKQLDashboards request.
func (client *ItemsClient) listKQLDashboardsCreateRequest(ctx context.Context, workspaceID string, options *ItemsClientListKQLDashboardsOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/kqlDashboards"
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

// listKQLDashboardsHandleResponse handles the ListKQLDashboards response.
func (client *ItemsClient) listKQLDashboardsHandleResponse(resp *http.Response) (ItemsClientListKQLDashboardsResponse, error) {
	result := ItemsClientListKQLDashboardsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KQLDashboards); err != nil {
		return ItemsClientListKQLDashboardsResponse{}, err
	}
	return result, nil
}

// UpdateKQLDashboard - PERMISSIONS The caller must have contributor or higher workspace role.
// REQUIRED DELEGATED SCOPES KQLDashboard.ReadWrite.All or Item.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | Yes | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - kqlDashboardID - The KQL dashboard ID.
//   - updateKQLDashboardRequest - Update KQL dashboard request payload.
//   - options - ItemsClientUpdateKQLDashboardOptions contains the optional parameters for the ItemsClient.UpdateKQLDashboard
//     method.
func (client *ItemsClient) UpdateKQLDashboard(ctx context.Context, workspaceID string, kqlDashboardID string, updateKQLDashboardRequest UpdateKQLDashboardRequest, options *ItemsClientUpdateKQLDashboardOptions) (ItemsClientUpdateKQLDashboardResponse, error) {
	var err error
	const operationName = "kqldashboard.ItemsClient.UpdateKQLDashboard"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateKQLDashboardCreateRequest(ctx, workspaceID, kqlDashboardID, updateKQLDashboardRequest, options)
	if err != nil {
		return ItemsClientUpdateKQLDashboardResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ItemsClientUpdateKQLDashboardResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return ItemsClientUpdateKQLDashboardResponse{}, err
	}
	resp, err := client.updateKQLDashboardHandleResponse(httpResp)
	return resp, err
}

// updateKQLDashboardCreateRequest creates the UpdateKQLDashboard request.
func (client *ItemsClient) updateKQLDashboardCreateRequest(ctx context.Context, workspaceID string, kqlDashboardID string, updateKQLDashboardRequest UpdateKQLDashboardRequest, _ *ItemsClientUpdateKQLDashboardOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/kqlDashboards/{kqlDashboardId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if kqlDashboardID == "" {
		return nil, errors.New("parameter kqlDashboardID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kqlDashboardId}", url.PathEscape(kqlDashboardID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, updateKQLDashboardRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// updateKQLDashboardHandleResponse handles the UpdateKQLDashboard response.
func (client *ItemsClient) updateKQLDashboardHandleResponse(resp *http.Response) (ItemsClientUpdateKQLDashboardResponse, error) {
	result := ItemsClientUpdateKQLDashboardResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KQLDashboard); err != nil {
		return ItemsClientUpdateKQLDashboardResponse{}, err
	}
	return result, nil
}

// UpdateKQLDashboardDefinition - PERMISSIONS The caller must have contributor or higher workspace role.
// REQUIRED DELEGATED SCOPES KQLDashboard.ReadWrite.All or Item.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | Yes | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - kqlDashboardID - The KQL dashboard ID.
//   - updateKQLDashboardDefinitionRequest - Update KQL dashboard definition request payload.
//   - options - ItemsClientUpdateKQLDashboardDefinitionOptions contains the optional parameters for the ItemsClient.UpdateKQLDashboardDefinition
//     method.
func (client *ItemsClient) UpdateKQLDashboardDefinition(ctx context.Context, workspaceID string, kqlDashboardID string, updateKQLDashboardDefinitionRequest UpdateKQLDashboardDefinitionRequest, options *ItemsClientUpdateKQLDashboardDefinitionOptions) (ItemsClientUpdateKQLDashboardDefinitionResponse, error) {
	var err error
	const operationName = "kqldashboard.ItemsClient.UpdateKQLDashboardDefinition"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateKQLDashboardDefinitionCreateRequest(ctx, workspaceID, kqlDashboardID, updateKQLDashboardDefinitionRequest, options)
	if err != nil {
		return ItemsClientUpdateKQLDashboardDefinitionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ItemsClientUpdateKQLDashboardDefinitionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return ItemsClientUpdateKQLDashboardDefinitionResponse{}, err
	}
	return ItemsClientUpdateKQLDashboardDefinitionResponse{}, nil
}

// updateKQLDashboardDefinitionCreateRequest creates the UpdateKQLDashboardDefinition request.
func (client *ItemsClient) updateKQLDashboardDefinitionCreateRequest(ctx context.Context, workspaceID string, kqlDashboardID string, updateKQLDashboardDefinitionRequest UpdateKQLDashboardDefinitionRequest, options *ItemsClientUpdateKQLDashboardDefinitionOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/kqlDashboards/{kqlDashboardId}/updateDefinition"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if kqlDashboardID == "" {
		return nil, errors.New("parameter kqlDashboardID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kqlDashboardId}", url.PathEscape(kqlDashboardID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.UpdateMetadata != nil {
		reqQP.Set("updateMetadata", strconv.FormatBool(*options.UpdateMetadata))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, updateKQLDashboardDefinitionRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// Custom code starts below

// ListKQLDashboards - returns array of KQLDashboard from all pages.
// This API supports pagination [/rest/api/fabric/articles/pagination].
//
// PERMISSIONS The caller must have viewer or higher workspace role.
//
// # REQUIRED DELEGATED SCOPES Workspace.Read.All or Workspace.ReadWrite.All
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] | Yes | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
//
// INTERFACE
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - options - ItemsClientListKQLDashboardsOptions contains the optional parameters for the ItemsClient.NewListKQLDashboardsPager method.
func (client *ItemsClient) ListKQLDashboards(ctx context.Context, workspaceID string, options *ItemsClientListKQLDashboardsOptions) ([]KQLDashboard, error) {
	pager := client.NewListKQLDashboardsPager(workspaceID, options)
	mapper := func(resp ItemsClientListKQLDashboardsResponse) []KQLDashboard {
		return resp.Value
	}
	list, err := iruntime.NewPageIterator(ctx, pager, mapper).Get()
	if err != nil {
		var azcoreRespError *azcore.ResponseError
		if errors.As(err, &azcoreRespError) {
			return []KQLDashboard{}, core.NewResponseError(azcoreRespError.RawResponse)
		}
		return []KQLDashboard{}, err
	}
	return list, nil
}
