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
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"

	"github.com/microsoft/fabric-sdk-go/internal/iruntime"
)

// ExternalDataSharesClient contains the methods for the ExternalDataShares group.
// Don't use this type directly, use a constructor function instead.
type ExternalDataSharesClient struct {
	internal *azcore.Client
	endpoint string
}

// CreateExternalDataShare - PERMISSIONS The caller must have Read and Reshare permissions on the item.
// REQUIRED DELEGATED SCOPES For item APIs you can have one of the 2 types of scopes in the token:
// * Generic scope: Item.ExternalDataShare.All
//
// * Specific scope: itemType.ExternalDataShare.All (for example: Lakehouse.ExternalDataShare.All)
//
// For more information about scopes, see scopes article [/rest/api/fabric/articles/scopes].
//
// REQUIRED TENANT SETTINGS To use this API, enable the 'External data sharing' admin switch for the calling principal.
// LIMITATIONS Maximum 10 requests per minute.
// Supported item types: See External data sharing in Microsoft Fabric - Supported item types [/fabric/governance/external-data-sharing-overview#supported-fabric-item-types].
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
//   - createExternalDataShareRequest - The request payload for creating an external data share.
//   - options - ExternalDataSharesClientCreateExternalDataShareOptions contains the optional parameters for the ExternalDataSharesClient.CreateExternalDataShare
//     method.
func (client *ExternalDataSharesClient) CreateExternalDataShare(ctx context.Context, workspaceID string, itemID string, createExternalDataShareRequest CreateExternalDataShareRequest, options *ExternalDataSharesClientCreateExternalDataShareOptions) (ExternalDataSharesClientCreateExternalDataShareResponse, error) {
	var err error
	const operationName = "core.ExternalDataSharesClient.CreateExternalDataShare"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createExternalDataShareCreateRequest(ctx, workspaceID, itemID, createExternalDataShareRequest, options)
	if err != nil {
		return ExternalDataSharesClientCreateExternalDataShareResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExternalDataSharesClientCreateExternalDataShareResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = NewResponseError(httpResp)
		return ExternalDataSharesClientCreateExternalDataShareResponse{}, err
	}
	resp, err := client.createExternalDataShareHandleResponse(httpResp)
	return resp, err
}

// createExternalDataShareCreateRequest creates the CreateExternalDataShare request.
func (client *ExternalDataSharesClient) createExternalDataShareCreateRequest(ctx context.Context, workspaceID string, itemID string, createExternalDataShareRequest CreateExternalDataShareRequest, _ *ExternalDataSharesClientCreateExternalDataShareOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/items/{itemId}/externalDataShares"
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, createExternalDataShareRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// createExternalDataShareHandleResponse handles the CreateExternalDataShare response.
func (client *ExternalDataSharesClient) createExternalDataShareHandleResponse(resp *http.Response) (ExternalDataSharesClientCreateExternalDataShareResponse, error) {
	result := ExternalDataSharesClientCreateExternalDataShareResponse{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExternalDataShare); err != nil {
		return ExternalDataSharesClientCreateExternalDataShareResponse{}, err
	}
	return result, nil
}

// GetExternalDataShare - PERMISSIONS The caller must have Read and Reshare permissions on the item.
// REQUIRED DELEGATED SCOPES For item APIs you can have one of the 2 types of scopes in the token:
// * Generic scope: Item.ExternalDataShare.All
//
// * Specific scope: itemType.ExternalDataShare.All (for example: Lakehouse.ExternalDataShare.All)
//
// For more information about scopes, see scopes article [/rest/api/fabric/articles/scopes].
//
// REQUIRED TENANT SETTINGS To use this API, enable the 'External data sharing' admin switch for the calling principal.
// LIMITATIONS Maximum 10 requests per minute.
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
//   - externalDataShareID - The external data share ID.
//   - options - ExternalDataSharesClientGetExternalDataShareOptions contains the optional parameters for the ExternalDataSharesClient.GetExternalDataShare
//     method.
func (client *ExternalDataSharesClient) GetExternalDataShare(ctx context.Context, workspaceID string, itemID string, externalDataShareID string, options *ExternalDataSharesClientGetExternalDataShareOptions) (ExternalDataSharesClientGetExternalDataShareResponse, error) {
	var err error
	const operationName = "core.ExternalDataSharesClient.GetExternalDataShare"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getExternalDataShareCreateRequest(ctx, workspaceID, itemID, externalDataShareID, options)
	if err != nil {
		return ExternalDataSharesClientGetExternalDataShareResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExternalDataSharesClientGetExternalDataShareResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = NewResponseError(httpResp)
		return ExternalDataSharesClientGetExternalDataShareResponse{}, err
	}
	resp, err := client.getExternalDataShareHandleResponse(httpResp)
	return resp, err
}

// getExternalDataShareCreateRequest creates the GetExternalDataShare request.
func (client *ExternalDataSharesClient) getExternalDataShareCreateRequest(ctx context.Context, workspaceID string, itemID string, externalDataShareID string, _ *ExternalDataSharesClientGetExternalDataShareOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/items/{itemId}/externalDataShares/{externalDataShareId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if itemID == "" {
		return nil, errors.New("parameter itemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{itemId}", url.PathEscape(itemID))
	if externalDataShareID == "" {
		return nil, errors.New("parameter externalDataShareID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{externalDataShareId}", url.PathEscape(externalDataShareID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getExternalDataShareHandleResponse handles the GetExternalDataShare response.
func (client *ExternalDataSharesClient) getExternalDataShareHandleResponse(resp *http.Response) (ExternalDataSharesClientGetExternalDataShareResponse, error) {
	result := ExternalDataSharesClientGetExternalDataShareResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExternalDataShare); err != nil {
		return ExternalDataSharesClientGetExternalDataShareResponse{}, err
	}
	return result, nil
}

// NewListExternalDataSharesInItemPager - This API supports pagination [/rest/api/fabric/articles/pagination].
// PERMISSIONS The caller must have Read and Reshare permissions on the item.
// REQUIRED DELEGATED SCOPES For item APIs you can have one of the 2 types of scopes in the token:
// * Generic scope: Item.ExternalDataShare.All
//
// * Specific scope: itemType.ExternalDataShare.All (for example: Lakehouse.ExternalDataShare.All)
//
// For more information about scopes, see scopes article [/rest/api/fabric/articles/scopes].
//
// REQUIRED TENANT SETTINGS To use this API, enable the 'External data sharing' admin switch for the calling principal.
// LIMITATIONS Maximum 10 requests per minute.
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - itemID - The item ID.
//   - options - ExternalDataSharesClientListExternalDataSharesInItemOptions contains the optional parameters for the ExternalDataSharesClient.NewListExternalDataSharesInItemPager
//     method.
func (client *ExternalDataSharesClient) NewListExternalDataSharesInItemPager(workspaceID string, itemID string, options *ExternalDataSharesClientListExternalDataSharesInItemOptions) *runtime.Pager[ExternalDataSharesClientListExternalDataSharesInItemResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExternalDataSharesClientListExternalDataSharesInItemResponse]{
		More: func(page ExternalDataSharesClientListExternalDataSharesInItemResponse) bool {
			return page.ContinuationURI != nil && len(*page.ContinuationURI) > 0
		},
		Fetcher: func(ctx context.Context, page *ExternalDataSharesClientListExternalDataSharesInItemResponse) (ExternalDataSharesClientListExternalDataSharesInItemResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "core.ExternalDataSharesClient.NewListExternalDataSharesInItemPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ContinuationURI
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listExternalDataSharesInItemCreateRequest(ctx, workspaceID, itemID, options)
			}, nil)
			if err != nil {
				return ExternalDataSharesClientListExternalDataSharesInItemResponse{}, err
			}
			return client.listExternalDataSharesInItemHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listExternalDataSharesInItemCreateRequest creates the ListExternalDataSharesInItem request.
func (client *ExternalDataSharesClient) listExternalDataSharesInItemCreateRequest(ctx context.Context, workspaceID string, itemID string, options *ExternalDataSharesClientListExternalDataSharesInItemOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/items/{itemId}/externalDataShares"
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
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listExternalDataSharesInItemHandleResponse handles the ListExternalDataSharesInItem response.
func (client *ExternalDataSharesClient) listExternalDataSharesInItemHandleResponse(resp *http.Response) (ExternalDataSharesClientListExternalDataSharesInItemResponse, error) {
	result := ExternalDataSharesClientListExternalDataSharesInItemResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExternalDataShares); err != nil {
		return ExternalDataSharesClientListExternalDataSharesInItemResponse{}, err
	}
	return result, nil
}

// RevokeExternalDataShare - PERMISSIONS The caller must have Read and Reshare permissions on the item.
// REQUIRED DELEGATED SCOPES For item APIs you can have one of the 2 types of scopes in the token:
// * Generic scope: Item.ExternalDataShare.All
//
// * Specific scope: itemType.ExternalDataShare.All (for example: Lakehouse.ExternalDataShare.All)
//
// For more information about scopes, see scopes article [/rest/api/fabric/articles/scopes].
//
// REQUIRED TENANT SETTINGS To use this API, enable the 'External data sharing' admin switch for the calling principal.
// LIMITATIONS Maximum 10 requests per minute.
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
//   - externalDataShareID - The external data share ID.
//   - options - ExternalDataSharesClientRevokeExternalDataShareOptions contains the optional parameters for the ExternalDataSharesClient.RevokeExternalDataShare
//     method.
func (client *ExternalDataSharesClient) RevokeExternalDataShare(ctx context.Context, workspaceID string, itemID string, externalDataShareID string, options *ExternalDataSharesClientRevokeExternalDataShareOptions) (ExternalDataSharesClientRevokeExternalDataShareResponse, error) {
	var err error
	const operationName = "core.ExternalDataSharesClient.RevokeExternalDataShare"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.revokeExternalDataShareCreateRequest(ctx, workspaceID, itemID, externalDataShareID, options)
	if err != nil {
		return ExternalDataSharesClientRevokeExternalDataShareResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExternalDataSharesClientRevokeExternalDataShareResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = NewResponseError(httpResp)
		return ExternalDataSharesClientRevokeExternalDataShareResponse{}, err
	}
	return ExternalDataSharesClientRevokeExternalDataShareResponse{}, nil
}

// revokeExternalDataShareCreateRequest creates the RevokeExternalDataShare request.
func (client *ExternalDataSharesClient) revokeExternalDataShareCreateRequest(ctx context.Context, workspaceID string, itemID string, externalDataShareID string, _ *ExternalDataSharesClientRevokeExternalDataShareOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/items/{itemId}/externalDataShares/{externalDataShareId}/revoke"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if itemID == "" {
		return nil, errors.New("parameter itemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{itemId}", url.PathEscape(itemID))
	if externalDataShareID == "" {
		return nil, errors.New("parameter externalDataShareID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{externalDataShareId}", url.PathEscape(externalDataShareID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Custom code starts below

// ListExternalDataSharesInItem - returns array of ExternalDataShare from all pages.
// This API supports pagination [/rest/api/fabric/articles/pagination].
//
// PERMISSIONS The caller must have Read and Reshare permissions on the item.
//
// REQUIRED DELEGATED SCOPES For item APIs you can have one of the 2 types of scopes in the token:
//
//   - Generic scope: Item.ExternalDataShare.All
//
//   - Specific scope: itemType.ExternalDataShare.All (for example: Lakehouse.ExternalDataShare.All)
//
//     For more information about scopes, see scopes article [/rest/api/fabric/articles/scopes].
//
// REQUIRED TENANT SETTINGS To use this API, enable the 'External data sharing' admin switch for the calling principal.
//
// LIMITATIONS Maximum 10 requests per minute.
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
//
// INTERFACE
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - itemID - The item ID.
//   - options - ExternalDataSharesClientListExternalDataSharesInItemOptions contains the optional parameters for the ExternalDataSharesClient.NewListExternalDataSharesInItemPager method.
func (client *ExternalDataSharesClient) ListExternalDataSharesInItem(ctx context.Context, workspaceID string, itemID string, options *ExternalDataSharesClientListExternalDataSharesInItemOptions) ([]ExternalDataShare, error) {
	pager := client.NewListExternalDataSharesInItemPager(workspaceID, itemID, options)
	mapper := func(resp ExternalDataSharesClientListExternalDataSharesInItemResponse) []ExternalDataShare {
		return resp.Value
	}
	list, err := iruntime.NewPageIterator(ctx, pager, mapper).Get()
	if err != nil {
		var azcoreRespError *azcore.ResponseError
		if errors.As(err, &azcoreRespError) {
			return []ExternalDataShare{}, NewResponseError(azcoreRespError.RawResponse)
		}
		return []ExternalDataShare{}, err
	}
	return list, nil
}
