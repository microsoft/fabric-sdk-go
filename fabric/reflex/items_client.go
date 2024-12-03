// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package reflex

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
	"github.com/microsoft/fabric-sdk-go/internal/pollers/locasync"
)

// ItemsClient contains the methods for the Items group.
// Don't use this type directly, use a constructor function instead.
type ItemsClient struct {
	internal *azcore.Client
	endpoint string
}

// BeginCreateReflex - This API supports long running operations (LRO) [/rest/api/fabric/articles/long-running-operation].
// To create notebook with definition, refer to Reflex definition [/rest/api/fabric/articles/item-management/definitions/reflex-definition]
// article.
// PERMISSIONS THE CALLER MUST HAVE CONTRIBUTOR OR HIGHER WORKSPACE ROLE.
// REQUIRED DELEGATED SCOPES Reflex.ReadWrite.All or Item.ReadWrite.All
// LIMITATIONS
// * To create a Reflex the workspace must be on a supported Fabric capacity. For more information see: Microsoft Fabric license
// types [/fabric/enterprise/licenses#microsoft-fabric-license-types].
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
//   - createReflexRequest - Create item request payload.
//   - options - ItemsClientBeginCreateReflexOptions contains the optional parameters for the ItemsClient.BeginCreateReflex method.
func (client *ItemsClient) BeginCreateReflex(ctx context.Context, workspaceID string, createReflexRequest CreateReflexRequest, options *ItemsClientBeginCreateReflexOptions) (*runtime.Poller[ItemsClientCreateReflexResponse], error) {
	return client.beginCreateReflex(ctx, workspaceID, createReflexRequest, options)
}

// CreateReflex - This API supports long running operations (LRO) [/rest/api/fabric/articles/long-running-operation].
// To create notebook with definition, refer to Reflex definition [/rest/api/fabric/articles/item-management/definitions/reflex-definition]
// article.
// PERMISSIONS THE CALLER MUST HAVE CONTRIBUTOR OR HIGHER WORKSPACE ROLE.
// REQUIRED DELEGATED SCOPES Reflex.ReadWrite.All or Item.ReadWrite.All
// LIMITATIONS
// * To create a Reflex the workspace must be on a supported Fabric capacity. For more information see: Microsoft Fabric license
// types [/fabric/enterprise/licenses#microsoft-fabric-license-types].
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | No | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | No |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
func (client *ItemsClient) createReflex(ctx context.Context, workspaceID string, createReflexRequest CreateReflexRequest, options *ItemsClientBeginCreateReflexOptions) (*http.Response, error) {
	var err error
	const operationName = "reflex.ItemsClient.BeginCreateReflex"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createReflexCreateRequest(ctx, workspaceID, createReflexRequest, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated, http.StatusAccepted) {
		err = core.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createReflexCreateRequest creates the CreateReflex request.
func (client *ItemsClient) createReflexCreateRequest(ctx context.Context, workspaceID string, createReflexRequest CreateReflexRequest, _ *ItemsClientBeginCreateReflexOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/reflexes"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, createReflexRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// DeleteReflex - PERMISSIONS The caller must have contributor or higher workspace role.
// REQUIRED DELEGATED SCOPES Reflex.ReadWrite.All or Item.ReadWrite.All
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
//   - reflexID - The Reflex ID.
//   - options - ItemsClientDeleteReflexOptions contains the optional parameters for the ItemsClient.DeleteReflex method.
func (client *ItemsClient) DeleteReflex(ctx context.Context, workspaceID string, reflexID string, options *ItemsClientDeleteReflexOptions) (ItemsClientDeleteReflexResponse, error) {
	var err error
	const operationName = "reflex.ItemsClient.DeleteReflex"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteReflexCreateRequest(ctx, workspaceID, reflexID, options)
	if err != nil {
		return ItemsClientDeleteReflexResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ItemsClientDeleteReflexResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return ItemsClientDeleteReflexResponse{}, err
	}
	return ItemsClientDeleteReflexResponse{}, nil
}

// deleteReflexCreateRequest creates the DeleteReflex request.
func (client *ItemsClient) deleteReflexCreateRequest(ctx context.Context, workspaceID string, reflexID string, _ *ItemsClientDeleteReflexOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/reflexes/{reflexId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if reflexID == "" {
		return nil, errors.New("parameter reflexID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reflexId}", url.PathEscape(reflexID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetReflex - PERMISSIONS The caller must have viewer or higher workspace role.
// REQUIRED DELEGATED SCOPES Reflex.Read.All or Reflex.ReadWrite.All or Item.Read.All or Item.ReadWrite.All
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
//   - reflexID - The Reflex ID.
//   - options - ItemsClientGetReflexOptions contains the optional parameters for the ItemsClient.GetReflex method.
func (client *ItemsClient) GetReflex(ctx context.Context, workspaceID string, reflexID string, options *ItemsClientGetReflexOptions) (ItemsClientGetReflexResponse, error) {
	var err error
	const operationName = "reflex.ItemsClient.GetReflex"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getReflexCreateRequest(ctx, workspaceID, reflexID, options)
	if err != nil {
		return ItemsClientGetReflexResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ItemsClientGetReflexResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return ItemsClientGetReflexResponse{}, err
	}
	resp, err := client.getReflexHandleResponse(httpResp)
	return resp, err
}

// getReflexCreateRequest creates the GetReflex request.
func (client *ItemsClient) getReflexCreateRequest(ctx context.Context, workspaceID string, reflexID string, _ *ItemsClientGetReflexOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/reflexes/{reflexId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if reflexID == "" {
		return nil, errors.New("parameter reflexID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reflexId}", url.PathEscape(reflexID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getReflexHandleResponse handles the GetReflex response.
func (client *ItemsClient) getReflexHandleResponse(resp *http.Response) (ItemsClientGetReflexResponse, error) {
	result := ItemsClientGetReflexResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Reflex); err != nil {
		return ItemsClientGetReflexResponse{}, err
	}
	return result, nil
}

// BeginGetReflexDefinition - This API supports long running operations (LRO) [/rest/api/fabric/articles/long-running-operation].
// When you get a Reflex's public definition, the sensitivity label is not a part of the definition.
// PERMISSIONS The caller must have contributor or higher workspace role.
// REQUIRED DELEGATED SCOPES Reflex.ReadWrite.All or Item.ReadWrite.All
// LIMITATIONS This API is blocked for a Reflex with an encrypted sensitivity label.
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
//   - reflexID - The Reflex ID.
//   - options - ItemsClientBeginGetReflexDefinitionOptions contains the optional parameters for the ItemsClient.BeginGetReflexDefinition
//     method.
func (client *ItemsClient) BeginGetReflexDefinition(ctx context.Context, workspaceID string, reflexID string, options *ItemsClientBeginGetReflexDefinitionOptions) (*runtime.Poller[ItemsClientGetReflexDefinitionResponse], error) {
	return client.beginGetReflexDefinition(ctx, workspaceID, reflexID, options)
}

// GetReflexDefinition - This API supports long running operations (LRO) [/rest/api/fabric/articles/long-running-operation].
// When you get a Reflex's public definition, the sensitivity label is not a part of the definition.
// PERMISSIONS The caller must have contributor or higher workspace role.
// REQUIRED DELEGATED SCOPES Reflex.ReadWrite.All or Item.ReadWrite.All
// LIMITATIONS This API is blocked for a Reflex with an encrypted sensitivity label.
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | No | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | No |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
func (client *ItemsClient) getReflexDefinition(ctx context.Context, workspaceID string, reflexID string, options *ItemsClientBeginGetReflexDefinitionOptions) (*http.Response, error) {
	var err error
	const operationName = "reflex.ItemsClient.BeginGetReflexDefinition"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getReflexDefinitionCreateRequest(ctx, workspaceID, reflexID, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = core.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// getReflexDefinitionCreateRequest creates the GetReflexDefinition request.
func (client *ItemsClient) getReflexDefinitionCreateRequest(ctx context.Context, workspaceID string, reflexID string, options *ItemsClientBeginGetReflexDefinitionOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/reflexes/{reflexId}/getDefinition"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if reflexID == "" {
		return nil, errors.New("parameter reflexID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reflexId}", url.PathEscape(reflexID))
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

// NewListReflexesPager - This API supports pagination [/rest/api/fabric/articles/pagination].
// PERMISSIONS The caller must have viewer or higher workspace role.
// REQUIRED DELEGATED SCOPES Workspace.Read.All or Workspace.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | No | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | No |
// INTERFACE
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - options - ItemsClientListReflexesOptions contains the optional parameters for the ItemsClient.NewListReflexesPager method.
func (client *ItemsClient) NewListReflexesPager(workspaceID string, options *ItemsClientListReflexesOptions) *runtime.Pager[ItemsClientListReflexesResponse] {
	return runtime.NewPager(runtime.PagingHandler[ItemsClientListReflexesResponse]{
		More: func(page ItemsClientListReflexesResponse) bool {
			return page.ContinuationURI != nil && len(*page.ContinuationURI) > 0
		},
		Fetcher: func(ctx context.Context, page *ItemsClientListReflexesResponse) (ItemsClientListReflexesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "reflex.ItemsClient.NewListReflexesPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ContinuationURI
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listReflexesCreateRequest(ctx, workspaceID, options)
			}, nil)
			if err != nil {
				return ItemsClientListReflexesResponse{}, err
			}
			return client.listReflexesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listReflexesCreateRequest creates the ListReflexes request.
func (client *ItemsClient) listReflexesCreateRequest(ctx context.Context, workspaceID string, options *ItemsClientListReflexesOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/reflexes"
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

// listReflexesHandleResponse handles the ListReflexes response.
func (client *ItemsClient) listReflexesHandleResponse(resp *http.Response) (ItemsClientListReflexesResponse, error) {
	result := ItemsClientListReflexesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Reflexes); err != nil {
		return ItemsClientListReflexesResponse{}, err
	}
	return result, nil
}

// UpdateReflex - PERMISSIONS The caller must have contributor or higher workspace role.
// REQUIRED DELEGATED SCOPES Reflex.ReadWrite.All or Item.ReadWrite.All
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
//   - reflexID - The Reflex ID.
//   - updateReflexRequest - Update Reflex request payload.
//   - options - ItemsClientUpdateReflexOptions contains the optional parameters for the ItemsClient.UpdateReflex method.
func (client *ItemsClient) UpdateReflex(ctx context.Context, workspaceID string, reflexID string, updateReflexRequest UpdateReflexRequest, options *ItemsClientUpdateReflexOptions) (ItemsClientUpdateReflexResponse, error) {
	var err error
	const operationName = "reflex.ItemsClient.UpdateReflex"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateReflexCreateRequest(ctx, workspaceID, reflexID, updateReflexRequest, options)
	if err != nil {
		return ItemsClientUpdateReflexResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ItemsClientUpdateReflexResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return ItemsClientUpdateReflexResponse{}, err
	}
	resp, err := client.updateReflexHandleResponse(httpResp)
	return resp, err
}

// updateReflexCreateRequest creates the UpdateReflex request.
func (client *ItemsClient) updateReflexCreateRequest(ctx context.Context, workspaceID string, reflexID string, updateReflexRequest UpdateReflexRequest, _ *ItemsClientUpdateReflexOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/reflexes/{reflexId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if reflexID == "" {
		return nil, errors.New("parameter reflexID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reflexId}", url.PathEscape(reflexID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, updateReflexRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// updateReflexHandleResponse handles the UpdateReflex response.
func (client *ItemsClient) updateReflexHandleResponse(resp *http.Response) (ItemsClientUpdateReflexResponse, error) {
	result := ItemsClientUpdateReflexResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Reflex); err != nil {
		return ItemsClientUpdateReflexResponse{}, err
	}
	return result, nil
}

// BeginUpdateReflexDefinition - This API supports long running operations (LRO) [/rest/api/fabric/articles/long-running-operation].
// Updating the Reflex's definition, does not affect its sensitivity label.
// PERMISSIONS The API caller must have contributor or higher workspace role.
// REQUIRED DELEGATED SCOPES Reflex.ReadWrite.All or Item.ReadWrite.All
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
//   - reflexID - The Reflex ID.
//   - updateReflexDefinitionRequest - Update Reflex definition request payload.
//   - options - ItemsClientBeginUpdateReflexDefinitionOptions contains the optional parameters for the ItemsClient.BeginUpdateReflexDefinition
//     method.
func (client *ItemsClient) BeginUpdateReflexDefinition(ctx context.Context, workspaceID string, reflexID string, updateReflexDefinitionRequest UpdateReflexDefinitionRequest, options *ItemsClientBeginUpdateReflexDefinitionOptions) (*runtime.Poller[ItemsClientUpdateReflexDefinitionResponse], error) {
	return client.beginUpdateReflexDefinition(ctx, workspaceID, reflexID, updateReflexDefinitionRequest, options)
}

// UpdateReflexDefinition - This API supports long running operations (LRO) [/rest/api/fabric/articles/long-running-operation].
// Updating the Reflex's definition, does not affect its sensitivity label.
// PERMISSIONS The API caller must have contributor or higher workspace role.
// REQUIRED DELEGATED SCOPES Reflex.ReadWrite.All or Item.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | No | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | No |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
func (client *ItemsClient) updateReflexDefinition(ctx context.Context, workspaceID string, reflexID string, updateReflexDefinitionRequest UpdateReflexDefinitionRequest, options *ItemsClientBeginUpdateReflexDefinitionOptions) (*http.Response, error) {
	var err error
	const operationName = "reflex.ItemsClient.BeginUpdateReflexDefinition"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateReflexDefinitionCreateRequest(ctx, workspaceID, reflexID, updateReflexDefinitionRequest, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = core.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateReflexDefinitionCreateRequest creates the UpdateReflexDefinition request.
func (client *ItemsClient) updateReflexDefinitionCreateRequest(ctx context.Context, workspaceID string, reflexID string, updateReflexDefinitionRequest UpdateReflexDefinitionRequest, options *ItemsClientBeginUpdateReflexDefinitionOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/reflexes/{reflexId}/updateDefinition"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if reflexID == "" {
		return nil, errors.New("parameter reflexID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reflexId}", url.PathEscape(reflexID))
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
	if err := runtime.MarshalAsJSON(req, updateReflexDefinitionRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// Custom code starts below

// CreateReflex - returns ItemsClientCreateReflexResponse in sync mode.
// This API supports long running operations (LRO) [/rest/api/fabric/articles/long-running-operation].
//
// To create notebook with definition, refer to Reflex definition [/rest/api/fabric/articles/item-management/definitions/reflex-definition] article.
//
// PERMISSIONS THE CALLER MUST HAVE CONTRIBUTOR OR HIGHER WORKSPACE ROLE.
// REQUIRED DELEGATED SCOPES Reflex.ReadWrite.All or Item.ReadWrite.All
//
// LIMITATIONS
//
//   - To create a Reflex the workspace must be on a supported Fabric capacity. For more information see: Microsoft Fabric license types [/fabric/enterprise/licenses#microsoft-fabric-license-types].
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] | No | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | No |
//
// INTERFACE
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - createReflexRequest - Create item request payload.
//   - options - ItemsClientBeginCreateReflexOptions contains the optional parameters for the ItemsClient.BeginCreateReflex method.
func (client *ItemsClient) CreateReflex(ctx context.Context, workspaceID string, createReflexRequest CreateReflexRequest, options *ItemsClientBeginCreateReflexOptions) (ItemsClientCreateReflexResponse, error) {
	result, err := iruntime.NewLRO(client.BeginCreateReflex(ctx, workspaceID, createReflexRequest, options)).Sync(ctx)
	if err != nil {
		var azcoreRespError *azcore.ResponseError
		if errors.As(err, &azcoreRespError) {
			return ItemsClientCreateReflexResponse{}, core.NewResponseError(azcoreRespError.RawResponse)
		}
		return ItemsClientCreateReflexResponse{}, err
	}
	return result, err
}

// beginCreateReflex creates the createReflex request.
func (client *ItemsClient) beginCreateReflex(ctx context.Context, workspaceID string, createReflexRequest CreateReflexRequest, options *ItemsClientBeginCreateReflexOptions) (*runtime.Poller[ItemsClientCreateReflexResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createReflex(ctx, workspaceID, createReflexRequest, options)
		if err != nil {
			var azcoreRespError *azcore.ResponseError
			if errors.As(err, &azcoreRespError) {
				return nil, core.NewResponseError(azcoreRespError.RawResponse)
			}
			return nil, err
		}
		handler, err := locasync.NewPollerHandler[ItemsClientCreateReflexResponse](client.internal.Pipeline(), resp, runtime.FinalStateViaAzureAsyncOp)
		if err != nil {
			var azcoreRespError *azcore.ResponseError
			if errors.As(err, &azcoreRespError) {
				return nil, core.NewResponseError(azcoreRespError.RawResponse)
			}
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ItemsClientCreateReflexResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Handler:       handler,
			Tracer:        client.internal.Tracer(),
		})
	} else {
		handler, err := locasync.NewPollerHandler[ItemsClientCreateReflexResponse](client.internal.Pipeline(), nil, runtime.FinalStateViaAzureAsyncOp)
		if err != nil {
			var azcoreRespError *azcore.ResponseError
			if errors.As(err, &azcoreRespError) {
				return nil, core.NewResponseError(azcoreRespError.RawResponse)
			}
			return nil, err
		}
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ItemsClientCreateReflexResponse]{
			Handler: handler,
			Tracer:  client.internal.Tracer(),
		})
	}
}

// GetReflexDefinition - returns ItemsClientGetReflexDefinitionResponse in sync mode.
// This API supports long running operations (LRO) [/rest/api/fabric/articles/long-running-operation].
//
// When you get a Reflex's public definition, the sensitivity label is not a part of the definition.
//
// PERMISSIONS The caller must have contributor or higher workspace role.
//
// # REQUIRED DELEGATED SCOPES Reflex.ReadWrite.All or Item.ReadWrite.All
//
// LIMITATIONS This API is blocked for a Reflex with an encrypted sensitivity label.
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] | No | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | No |
//
// INTERFACE
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - reflexID - The Reflex ID.
//   - options - ItemsClientBeginGetReflexDefinitionOptions contains the optional parameters for the ItemsClient.BeginGetReflexDefinition method.
func (client *ItemsClient) GetReflexDefinition(ctx context.Context, workspaceID string, reflexID string, options *ItemsClientBeginGetReflexDefinitionOptions) (ItemsClientGetReflexDefinitionResponse, error) {
	result, err := iruntime.NewLRO(client.BeginGetReflexDefinition(ctx, workspaceID, reflexID, options)).Sync(ctx)
	if err != nil {
		var azcoreRespError *azcore.ResponseError
		if errors.As(err, &azcoreRespError) {
			return ItemsClientGetReflexDefinitionResponse{}, core.NewResponseError(azcoreRespError.RawResponse)
		}
		return ItemsClientGetReflexDefinitionResponse{}, err
	}
	return result, err
}

// beginGetReflexDefinition creates the getReflexDefinition request.
func (client *ItemsClient) beginGetReflexDefinition(ctx context.Context, workspaceID string, reflexID string, options *ItemsClientBeginGetReflexDefinitionOptions) (*runtime.Poller[ItemsClientGetReflexDefinitionResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.getReflexDefinition(ctx, workspaceID, reflexID, options)
		if err != nil {
			var azcoreRespError *azcore.ResponseError
			if errors.As(err, &azcoreRespError) {
				return nil, core.NewResponseError(azcoreRespError.RawResponse)
			}
			return nil, err
		}
		handler, err := locasync.NewPollerHandler[ItemsClientGetReflexDefinitionResponse](client.internal.Pipeline(), resp, runtime.FinalStateViaAzureAsyncOp)
		if err != nil {
			var azcoreRespError *azcore.ResponseError
			if errors.As(err, &azcoreRespError) {
				return nil, core.NewResponseError(azcoreRespError.RawResponse)
			}
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ItemsClientGetReflexDefinitionResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Handler:       handler,
			Tracer:        client.internal.Tracer(),
		})
	} else {
		handler, err := locasync.NewPollerHandler[ItemsClientGetReflexDefinitionResponse](client.internal.Pipeline(), nil, runtime.FinalStateViaAzureAsyncOp)
		if err != nil {
			var azcoreRespError *azcore.ResponseError
			if errors.As(err, &azcoreRespError) {
				return nil, core.NewResponseError(azcoreRespError.RawResponse)
			}
			return nil, err
		}
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ItemsClientGetReflexDefinitionResponse]{
			Handler: handler,
			Tracer:  client.internal.Tracer(),
		})
	}
}

// UpdateReflexDefinition - returns ItemsClientUpdateReflexDefinitionResponse in sync mode.
// This API supports long running operations (LRO) [/rest/api/fabric/articles/long-running-operation].
//
// Updating the Reflex's definition, does not affect its sensitivity label.
//
// PERMISSIONS The API caller must have contributor or higher workspace role.
//
// # REQUIRED DELEGATED SCOPES Reflex.ReadWrite.All or Item.ReadWrite.All
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] | No | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | No |
//
// INTERFACE
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - reflexID - The Reflex ID.
//   - updateReflexDefinitionRequest - Update Reflex definition request payload.
//   - options - ItemsClientBeginUpdateReflexDefinitionOptions contains the optional parameters for the ItemsClient.BeginUpdateReflexDefinition method.
func (client *ItemsClient) UpdateReflexDefinition(ctx context.Context, workspaceID string, reflexID string, updateReflexDefinitionRequest UpdateReflexDefinitionRequest, options *ItemsClientBeginUpdateReflexDefinitionOptions) (ItemsClientUpdateReflexDefinitionResponse, error) {
	result, err := iruntime.NewLRO(client.BeginUpdateReflexDefinition(ctx, workspaceID, reflexID, updateReflexDefinitionRequest, options)).Sync(ctx)
	if err != nil {
		var azcoreRespError *azcore.ResponseError
		if errors.As(err, &azcoreRespError) {
			return ItemsClientUpdateReflexDefinitionResponse{}, core.NewResponseError(azcoreRespError.RawResponse)
		}
		return ItemsClientUpdateReflexDefinitionResponse{}, err
	}
	return result, err
}

// beginUpdateReflexDefinition creates the updateReflexDefinition request.
func (client *ItemsClient) beginUpdateReflexDefinition(ctx context.Context, workspaceID string, reflexID string, updateReflexDefinitionRequest UpdateReflexDefinitionRequest, options *ItemsClientBeginUpdateReflexDefinitionOptions) (*runtime.Poller[ItemsClientUpdateReflexDefinitionResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateReflexDefinition(ctx, workspaceID, reflexID, updateReflexDefinitionRequest, options)
		if err != nil {
			var azcoreRespError *azcore.ResponseError
			if errors.As(err, &azcoreRespError) {
				return nil, core.NewResponseError(azcoreRespError.RawResponse)
			}
			return nil, err
		}
		handler, err := locasync.NewPollerHandler[ItemsClientUpdateReflexDefinitionResponse](client.internal.Pipeline(), resp, runtime.FinalStateViaAzureAsyncOp)
		if err != nil {
			var azcoreRespError *azcore.ResponseError
			if errors.As(err, &azcoreRespError) {
				return nil, core.NewResponseError(azcoreRespError.RawResponse)
			}
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ItemsClientUpdateReflexDefinitionResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Handler:       handler,
			Tracer:        client.internal.Tracer(),
		})
	} else {
		handler, err := locasync.NewPollerHandler[ItemsClientUpdateReflexDefinitionResponse](client.internal.Pipeline(), nil, runtime.FinalStateViaAzureAsyncOp)
		if err != nil {
			var azcoreRespError *azcore.ResponseError
			if errors.As(err, &azcoreRespError) {
				return nil, core.NewResponseError(azcoreRespError.RawResponse)
			}
			return nil, err
		}
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ItemsClientUpdateReflexDefinitionResponse]{
			Handler: handler,
			Tracer:  client.internal.Tracer(),
		})
	}
}

// ListReflexes - returns array of Reflex from all pages.
// This API supports pagination [/rest/api/fabric/articles/pagination].
//
// PERMISSIONS The caller must have viewer or higher workspace role.
//
// # REQUIRED DELEGATED SCOPES Workspace.Read.All or Workspace.ReadWrite.All
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] | No | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | No |
//
// INTERFACE
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - options - ItemsClientListReflexesOptions contains the optional parameters for the ItemsClient.NewListReflexesPager method.
func (client *ItemsClient) ListReflexes(ctx context.Context, workspaceID string, options *ItemsClientListReflexesOptions) ([]Reflex, error) {
	pager := client.NewListReflexesPager(workspaceID, options)
	mapper := func(resp ItemsClientListReflexesResponse) []Reflex {
		return resp.Value
	}
	list, err := iruntime.NewPageIterator(ctx, pager, mapper).Get()
	if err != nil {
		var azcoreRespError *azcore.ResponseError
		if errors.As(err, &azcoreRespError) {
			return []Reflex{}, core.NewResponseError(azcoreRespError.RawResponse)
		}
		return []Reflex{}, err
	}
	return list, nil
}
