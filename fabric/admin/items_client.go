// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package admin

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

// ItemsClient contains the methods for the Items group.
// Don't use this type directly, use a constructor function instead.
type ItemsClient struct {
	internal *azcore.Client
	endpoint string
}

// GetItem - PERMISSIONS The caller must be a Fabric administrator or authenticate using a service principal.
// REQUIRED DELEGATED SCOPES Tenant.Read.All or Tenant.ReadWrite.All
// LIMITATIONS Maximum 200 requests per hour.
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
//   - itemID - The item ID
//   - options - ItemsClientGetItemOptions contains the optional parameters for the ItemsClient.GetItem method.
func (client *ItemsClient) GetItem(ctx context.Context, workspaceID string, itemID string, options *ItemsClientGetItemOptions) (ItemsClientGetItemResponse, error) {
	var err error
	const operationName = "admin.ItemsClient.GetItem"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getItemCreateRequest(ctx, workspaceID, itemID, options)
	if err != nil {
		return ItemsClientGetItemResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ItemsClientGetItemResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return ItemsClientGetItemResponse{}, err
	}
	resp, err := client.getItemHandleResponse(httpResp)
	return resp, err
}

// getItemCreateRequest creates the GetItem request.
func (client *ItemsClient) getItemCreateRequest(ctx context.Context, workspaceID string, itemID string, options *ItemsClientGetItemOptions) (*policy.Request, error) {
	urlPath := "/v1/admin/workspaces/{workspaceId}/items/{itemId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if itemID == "" {
		return nil, errors.New("parameter itemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{itemId}", url.PathEscape(itemID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Type != nil {
		reqQP.Set("type", *options.Type)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getItemHandleResponse handles the GetItem response.
func (client *ItemsClient) getItemHandleResponse(resp *http.Response) (ItemsClientGetItemResponse, error) {
	result := ItemsClientGetItemResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Item); err != nil {
		return ItemsClientGetItemResponse{}, err
	}
	return result, nil
}

// ListItemAccessDetails - PERMISSIONS The caller must be a Fabric administrator or authenticate using a service principal.
// REQUIRED DELEGATED SCOPES Tenant.Read.All or Tenant.ReadWrite.All
// LIMITATIONS Maximum 200 requests per hour.
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
//   - itemID - The item ID.
//   - options - ItemsClientListItemAccessDetailsOptions contains the optional parameters for the ItemsClient.ListItemAccessDetails
//     method.
func (client *ItemsClient) ListItemAccessDetails(ctx context.Context, workspaceID string, itemID string, options *ItemsClientListItemAccessDetailsOptions) (ItemsClientListItemAccessDetailsResponse, error) {
	var err error
	const operationName = "admin.ItemsClient.ListItemAccessDetails"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listItemAccessDetailsCreateRequest(ctx, workspaceID, itemID, options)
	if err != nil {
		return ItemsClientListItemAccessDetailsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ItemsClientListItemAccessDetailsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return ItemsClientListItemAccessDetailsResponse{}, err
	}
	resp, err := client.listItemAccessDetailsHandleResponse(httpResp)
	return resp, err
}

// listItemAccessDetailsCreateRequest creates the ListItemAccessDetails request.
func (client *ItemsClient) listItemAccessDetailsCreateRequest(ctx context.Context, workspaceID string, itemID string, options *ItemsClientListItemAccessDetailsOptions) (*policy.Request, error) {
	urlPath := "/v1/admin/workspaces/{workspaceId}/items/{itemId}/users"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if itemID == "" {
		return nil, errors.New("parameter itemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{itemId}", url.PathEscape(itemID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Type != nil {
		reqQP.Set("type", *options.Type)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listItemAccessDetailsHandleResponse handles the ListItemAccessDetails response.
func (client *ItemsClient) listItemAccessDetailsHandleResponse(resp *http.Response) (ItemsClientListItemAccessDetailsResponse, error) {
	result := ItemsClientListItemAccessDetailsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ItemAccessDetailsResponse); err != nil {
		return ItemsClientListItemAccessDetailsResponse{}, err
	}
	return result, nil
}

// NewListItemsPager - This API supports pagination [/rest/api/fabric/articles/pagination]. A maximum of 10,000 records can
// be returned per request. With the continuous token provided in the response, you can get the next
// 10,000 records.
// Page order:
// 1. Fabric items
// 2. Datamarts
// 3. Reports
// 4. Dashboards
// 5. SemanticModels
// 6. Apps
// 7. Dataflows
// PERMISSIONS The caller must be a Fabric administrator or authenticate using a service principal.
// REQUIRED DELEGATED SCOPES Tenant.Read.All or Tenant.ReadWrite.All
// LIMITATIONS Maximum 200 requests per hour.
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
//
// Generated from API version v1
//   - options - ItemsClientListItemsOptions contains the optional parameters for the ItemsClient.NewListItemsPager method.
func (client *ItemsClient) NewListItemsPager(options *ItemsClientListItemsOptions) *runtime.Pager[ItemsClientListItemsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ItemsClientListItemsResponse]{
		More: func(page ItemsClientListItemsResponse) bool {
			return page.ContinuationURI != nil && len(*page.ContinuationURI) > 0
		},
		Fetcher: func(ctx context.Context, page *ItemsClientListItemsResponse) (ItemsClientListItemsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "admin.ItemsClient.NewListItemsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ContinuationURI
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listItemsCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ItemsClientListItemsResponse{}, err
			}
			return client.listItemsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listItemsCreateRequest creates the ListItems request.
func (client *ItemsClient) listItemsCreateRequest(ctx context.Context, options *ItemsClientListItemsOptions) (*policy.Request, error) {
	urlPath := "/v1/admin/items"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.CapacityID != nil {
		reqQP.Set("capacityId", *options.CapacityID)
	}
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	if options != nil && options.State != nil {
		reqQP.Set("state", *options.State)
	}
	if options != nil && options.Type != nil {
		reqQP.Set("type", *options.Type)
	}
	if options != nil && options.WorkspaceID != nil {
		reqQP.Set("workspaceId", *options.WorkspaceID)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listItemsHandleResponse handles the ListItems response.
func (client *ItemsClient) listItemsHandleResponse(resp *http.Response) (ItemsClientListItemsResponse, error) {
	result := ItemsClientListItemsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Items); err != nil {
		return ItemsClientListItemsResponse{}, err
	}
	return result, nil
}

// Custom code starts below

// ListItems - returns array of Item from all pages.
// This API supports pagination [/rest/api/fabric/articles/pagination]. A maximum of 10,000 records can be returned per request. With the continuous token provided in the response, you can get the next
// 10,000 records.
//
// Page order:
//
//  1. Fabric items
//  2. Datamarts
//  3. Reports
//  4. Dashboards
//  5. SemanticModels
//  6. Apps
//  7. Dataflows
//
// PERMISSIONS The caller must be a Fabric administrator or authenticate using a service principal.
//
// # REQUIRED DELEGATED SCOPES Tenant.Read.All or Tenant.ReadWrite.All
//
// LIMITATIONS Maximum 200 requests per hour.
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
//
// INTERFACE
// Generated from API version v1
//   - options - ItemsClientListItemsOptions contains the optional parameters for the ItemsClient.NewListItemsPager method.
func (client *ItemsClient) ListItems(ctx context.Context, options *ItemsClientListItemsOptions) ([]Item, error) {
	pager := client.NewListItemsPager(options)
	mapper := func(resp ItemsClientListItemsResponse) []Item {
		return resp.ItemEntities
	}
	list, err := iruntime.NewPageIterator(ctx, pager, mapper).Get()
	if err != nil {
		var azcoreRespError *azcore.ResponseError
		if errors.As(err, &azcoreRespError) {
			return []Item{}, core.NewResponseError(azcoreRespError.RawResponse)
		}
		return []Item{}, err
	}
	return list, nil
}
