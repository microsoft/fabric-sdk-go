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

// ExternalDataSharesProviderClient contains the methods for the ExternalDataSharesProvider group.
// Don't use this type directly, use a constructor function instead.
type ExternalDataSharesProviderClient struct {
	internal *azcore.Client
	endpoint string
}

// CreateExternalDataShare - PERMISSIONS The caller must have read and reshare permissions on the item.
// REQUIRED DELEGATED SCOPES Item APIs can have one of these scopes in their token:
// * Generic scope: Item.ExternalDataShare.All
//
// * Specific scope: itemType.ExternalDataShare.All, for example: Lakehouse.ExternalDataShare.All
//
// For more information about scopes, see scopes article [/rest/api/fabric/articles/scopes].
//
// REQUIRED TENANT SETTINGS To use this API, enable the External data sharing admin switch for the calling principal.
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
//   - options - ExternalDataSharesProviderClientCreateExternalDataShareOptions contains the optional parameters for the ExternalDataSharesProviderClient.CreateExternalDataShare
//     method.
func (client *ExternalDataSharesProviderClient) CreateExternalDataShare(ctx context.Context, workspaceID string, itemID string, createExternalDataShareRequest CreateExternalDataShareRequest, options *ExternalDataSharesProviderClientCreateExternalDataShareOptions) (ExternalDataSharesProviderClientCreateExternalDataShareResponse, error) {
	var err error
	const operationName = "core.ExternalDataSharesProviderClient.CreateExternalDataShare"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createExternalDataShareCreateRequest(ctx, workspaceID, itemID, createExternalDataShareRequest, options)
	if err != nil {
		return ExternalDataSharesProviderClientCreateExternalDataShareResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExternalDataSharesProviderClientCreateExternalDataShareResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = NewResponseError(httpResp)
		return ExternalDataSharesProviderClientCreateExternalDataShareResponse{}, err
	}
	resp, err := client.createExternalDataShareHandleResponse(httpResp)
	return resp, err
}

// createExternalDataShareCreateRequest creates the CreateExternalDataShare request.
func (client *ExternalDataSharesProviderClient) createExternalDataShareCreateRequest(ctx context.Context, workspaceID string, itemID string, createExternalDataShareRequest CreateExternalDataShareRequest, _ *ExternalDataSharesProviderClientCreateExternalDataShareOptions) (*policy.Request, error) {
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
func (client *ExternalDataSharesProviderClient) createExternalDataShareHandleResponse(resp *http.Response) (ExternalDataSharesProviderClientCreateExternalDataShareResponse, error) {
	result := ExternalDataSharesProviderClientCreateExternalDataShareResponse{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExternalDataShare); err != nil {
		return ExternalDataSharesProviderClientCreateExternalDataShareResponse{}, err
	}
	return result, nil
}

// DeleteExternalDataShare - PERMISSIONS The caller must have read and reshare permissions on the item.
// REQUIRED DELEGATED SCOPES Item APIs can have one of these scopes in their token:
// * Generic scope: Item.ExternalDataShare.All
//
// * Specific scope: itemType.ExternalDataShare.All, for example: Lakehouse.ExternalDataShare.All
//
// For more information about scopes, see scopes article [/rest/api/fabric/articles/scopes].
//
// REQUIRED TENANT SETTINGS To use this API, enable the External data sharing admin switch for the calling principal.
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
//   - options - ExternalDataSharesProviderClientDeleteExternalDataShareOptions contains the optional parameters for the ExternalDataSharesProviderClient.DeleteExternalDataShare
//     method.
func (client *ExternalDataSharesProviderClient) DeleteExternalDataShare(ctx context.Context, workspaceID string, itemID string, externalDataShareID string, options *ExternalDataSharesProviderClientDeleteExternalDataShareOptions) (ExternalDataSharesProviderClientDeleteExternalDataShareResponse, error) {
	var err error
	const operationName = "core.ExternalDataSharesProviderClient.DeleteExternalDataShare"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteExternalDataShareCreateRequest(ctx, workspaceID, itemID, externalDataShareID, options)
	if err != nil {
		return ExternalDataSharesProviderClientDeleteExternalDataShareResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExternalDataSharesProviderClientDeleteExternalDataShareResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = NewResponseError(httpResp)
		return ExternalDataSharesProviderClientDeleteExternalDataShareResponse{}, err
	}
	return ExternalDataSharesProviderClientDeleteExternalDataShareResponse{}, nil
}

// deleteExternalDataShareCreateRequest creates the DeleteExternalDataShare request.
func (client *ExternalDataSharesProviderClient) deleteExternalDataShareCreateRequest(ctx context.Context, workspaceID string, itemID string, externalDataShareID string, _ *ExternalDataSharesProviderClientDeleteExternalDataShareOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetExternalDataShare - PERMISSIONS The caller must have read and reshare permissions on the item.
// REQUIRED DELEGATED SCOPES Item APIs can have one of these scopes in their token:
// * Generic scope: Item.ExternalDataShare.All
//
// * Specific scope: itemType.ExternalDataShare.All, for example: Lakehouse.ExternalDataShare.All
//
// For more information about scopes, see scopes article [/rest/api/fabric/articles/scopes].
//
// REQUIRED TENANT SETTINGS To use this API, enable the External data sharing admin switch for the calling principal.
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
//   - options - ExternalDataSharesProviderClientGetExternalDataShareOptions contains the optional parameters for the ExternalDataSharesProviderClient.GetExternalDataShare
//     method.
func (client *ExternalDataSharesProviderClient) GetExternalDataShare(ctx context.Context, workspaceID string, itemID string, externalDataShareID string, options *ExternalDataSharesProviderClientGetExternalDataShareOptions) (ExternalDataSharesProviderClientGetExternalDataShareResponse, error) {
	var err error
	const operationName = "core.ExternalDataSharesProviderClient.GetExternalDataShare"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getExternalDataShareCreateRequest(ctx, workspaceID, itemID, externalDataShareID, options)
	if err != nil {
		return ExternalDataSharesProviderClientGetExternalDataShareResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExternalDataSharesProviderClientGetExternalDataShareResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = NewResponseError(httpResp)
		return ExternalDataSharesProviderClientGetExternalDataShareResponse{}, err
	}
	resp, err := client.getExternalDataShareHandleResponse(httpResp)
	return resp, err
}

// getExternalDataShareCreateRequest creates the GetExternalDataShare request.
func (client *ExternalDataSharesProviderClient) getExternalDataShareCreateRequest(ctx context.Context, workspaceID string, itemID string, externalDataShareID string, _ *ExternalDataSharesProviderClientGetExternalDataShareOptions) (*policy.Request, error) {
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
func (client *ExternalDataSharesProviderClient) getExternalDataShareHandleResponse(resp *http.Response) (ExternalDataSharesProviderClientGetExternalDataShareResponse, error) {
	result := ExternalDataSharesProviderClientGetExternalDataShareResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExternalDataShare); err != nil {
		return ExternalDataSharesProviderClientGetExternalDataShareResponse{}, err
	}
	return result, nil
}

// NewListExternalDataSharesInItemPager - This API supports pagination [/rest/api/fabric/articles/pagination].
// PERMISSIONS The caller must have read and reshare permissions on the item.
// REQUIRED DELEGATED SCOPES Item APIs can have one of these scopes in their token:
// * Generic scope: Item.ExternalDataShare.All
//
// * Specific scope: itemType.ExternalDataShare.All, for example: Lakehouse.ExternalDataShare.All
//
// For more information about scopes, see scopes article [/rest/api/fabric/articles/scopes].
//
// REQUIRED TENANT SETTINGS To use this API, enable the External data sharing admin switch for the calling principal.
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
//   - options - ExternalDataSharesProviderClientListExternalDataSharesInItemOptions contains the optional parameters for the
//     ExternalDataSharesProviderClient.NewListExternalDataSharesInItemPager method.
func (client *ExternalDataSharesProviderClient) NewListExternalDataSharesInItemPager(workspaceID string, itemID string, options *ExternalDataSharesProviderClientListExternalDataSharesInItemOptions) *runtime.Pager[ExternalDataSharesProviderClientListExternalDataSharesInItemResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExternalDataSharesProviderClientListExternalDataSharesInItemResponse]{
		More: func(page ExternalDataSharesProviderClientListExternalDataSharesInItemResponse) bool {
			return page.ContinuationURI != nil && len(*page.ContinuationURI) > 0
		},
		Fetcher: func(ctx context.Context, page *ExternalDataSharesProviderClientListExternalDataSharesInItemResponse) (ExternalDataSharesProviderClientListExternalDataSharesInItemResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "core.ExternalDataSharesProviderClient.NewListExternalDataSharesInItemPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ContinuationURI
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listExternalDataSharesInItemCreateRequest(ctx, workspaceID, itemID, options)
			}, nil)
			if err != nil {
				return ExternalDataSharesProviderClientListExternalDataSharesInItemResponse{}, err
			}
			return client.listExternalDataSharesInItemHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listExternalDataSharesInItemCreateRequest creates the ListExternalDataSharesInItem request.
func (client *ExternalDataSharesProviderClient) listExternalDataSharesInItemCreateRequest(ctx context.Context, workspaceID string, itemID string, options *ExternalDataSharesProviderClientListExternalDataSharesInItemOptions) (*policy.Request, error) {
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
func (client *ExternalDataSharesProviderClient) listExternalDataSharesInItemHandleResponse(resp *http.Response) (ExternalDataSharesProviderClientListExternalDataSharesInItemResponse, error) {
	result := ExternalDataSharesProviderClientListExternalDataSharesInItemResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExternalDataShares); err != nil {
		return ExternalDataSharesProviderClientListExternalDataSharesInItemResponse{}, err
	}
	return result, nil
}

// RevokeExternalDataShare - PERMISSIONS The caller must have read and reshare permissions on the item.
// REQUIRED DELEGATED SCOPES Item APIs can have one of these scopes in their token:
// * Generic scope: Item.ExternalDataShare.All
//
// * Specific scope: itemType.ExternalDataShare.All, for example: Lakehouse.ExternalDataShare.All
//
// For more information about scopes, see scopes article [/rest/api/fabric/articles/scopes].
//
// REQUIRED TENANT SETTINGS To use this API, enable the External data sharing admin switch for the calling principal.
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
//   - options - ExternalDataSharesProviderClientRevokeExternalDataShareOptions contains the optional parameters for the ExternalDataSharesProviderClient.RevokeExternalDataShare
//     method.
func (client *ExternalDataSharesProviderClient) RevokeExternalDataShare(ctx context.Context, workspaceID string, itemID string, externalDataShareID string, options *ExternalDataSharesProviderClientRevokeExternalDataShareOptions) (ExternalDataSharesProviderClientRevokeExternalDataShareResponse, error) {
	var err error
	const operationName = "core.ExternalDataSharesProviderClient.RevokeExternalDataShare"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.revokeExternalDataShareCreateRequest(ctx, workspaceID, itemID, externalDataShareID, options)
	if err != nil {
		return ExternalDataSharesProviderClientRevokeExternalDataShareResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExternalDataSharesProviderClientRevokeExternalDataShareResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = NewResponseError(httpResp)
		return ExternalDataSharesProviderClientRevokeExternalDataShareResponse{}, err
	}
	return ExternalDataSharesProviderClientRevokeExternalDataShareResponse{}, nil
}

// revokeExternalDataShareCreateRequest creates the RevokeExternalDataShare request.
func (client *ExternalDataSharesProviderClient) revokeExternalDataShareCreateRequest(ctx context.Context, workspaceID string, itemID string, externalDataShareID string, _ *ExternalDataSharesProviderClientRevokeExternalDataShareOptions) (*policy.Request, error) {
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
// PERMISSIONS The caller must have read and reshare permissions on the item.
//
// REQUIRED DELEGATED SCOPES Item APIs can have one of these scopes in their token:
//
//   - Generic scope: Item.ExternalDataShare.All
//
//   - Specific scope: itemType.ExternalDataShare.All, for example: Lakehouse.ExternalDataShare.All
//
//     For more information about scopes, see scopes article [/rest/api/fabric/articles/scopes].
//
// REQUIRED TENANT SETTINGS To use this API, enable the External data sharing admin switch for the calling principal.
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
//   - options - ExternalDataSharesProviderClientListExternalDataSharesInItemOptions contains the optional parameters for the ExternalDataSharesProviderClient.NewListExternalDataSharesInItemPager method.
func (client *ExternalDataSharesProviderClient) ListExternalDataSharesInItem(ctx context.Context, workspaceID string, itemID string, options *ExternalDataSharesProviderClientListExternalDataSharesInItemOptions) ([]ExternalDataShare, error) {
	pager := client.NewListExternalDataSharesInItemPager(workspaceID, itemID, options)
	mapper := func(resp ExternalDataSharesProviderClientListExternalDataSharesInItemResponse) []ExternalDataShare {
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
