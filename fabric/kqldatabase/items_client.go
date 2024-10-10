// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package kqldatabase

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
	"github.com/microsoft/fabric-sdk-go/internal/pollers/locasync"
)

// ItemsClient contains the methods for the Items group.
// Don't use this type directly, use a constructor function instead.
type ItemsClient struct {
	internal *azcore.Client
	endpoint string
}

// BeginCreateKQLDatabase - This API supports long running operations (LRO) [/rest/api/fabric/articles/long-running-operation].
// PERMISSIONS THE CALLER MUST HAVE CONTRIBUTOR OR HIGHER WORKSPACE ROLE.
// REQUIRED DELEGATED SCOPES KQLDatabase.ReadWrite.All or Item.ReadWrite.All
// LIMITATIONS
// * To create a KQL database the workspace must be on a supported Fabric capacity. For more information see: Microsoft Fabric
// license types [/fabric/enterprise/licenses#microsoft-fabric-license-types].
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
//   - createKQLDatabaseRequest - Create item request payload.
//   - options - ItemsClientBeginCreateKQLDatabaseOptions contains the optional parameters for the ItemsClient.BeginCreateKQLDatabase
//     method.
func (client *ItemsClient) BeginCreateKQLDatabase(ctx context.Context, workspaceID string, createKQLDatabaseRequest CreateKQLDatabaseRequest, options *ItemsClientBeginCreateKQLDatabaseOptions) (*runtime.Poller[ItemsClientCreateKQLDatabaseResponse], error) {
	return client.beginCreateKQLDatabase(ctx, workspaceID, createKQLDatabaseRequest, options)
}

// CreateKQLDatabase - This API supports long running operations (LRO) [/rest/api/fabric/articles/long-running-operation].
// PERMISSIONS THE CALLER MUST HAVE CONTRIBUTOR OR HIGHER WORKSPACE ROLE.
// REQUIRED DELEGATED SCOPES KQLDatabase.ReadWrite.All or Item.ReadWrite.All
// LIMITATIONS
// * To create a KQL database the workspace must be on a supported Fabric capacity. For more information see: Microsoft Fabric
// license types [/fabric/enterprise/licenses#microsoft-fabric-license-types].
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | Yes | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
func (client *ItemsClient) createKQLDatabase(ctx context.Context, workspaceID string, createKQLDatabaseRequest CreateKQLDatabaseRequest, options *ItemsClientBeginCreateKQLDatabaseOptions) (*http.Response, error) {
	var err error
	const operationName = "kqldatabase.ItemsClient.BeginCreateKQLDatabase"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createKQLDatabaseCreateRequest(ctx, workspaceID, createKQLDatabaseRequest, options)
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

// createKQLDatabaseCreateRequest creates the CreateKQLDatabase request.
func (client *ItemsClient) createKQLDatabaseCreateRequest(ctx context.Context, workspaceID string, createKQLDatabaseRequest CreateKQLDatabaseRequest, _ *ItemsClientBeginCreateKQLDatabaseOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/kqlDatabases"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, createKQLDatabaseRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// DeleteKQLDatabase - PERMISSIONS The caller must have contributor or higher workspace role.
// REQUIRED DELEGATED SCOPES KQLDatabase.ReadWrite.All or Item.ReadWrite.All
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
//   - kqlDatabaseID - The KQL database ID.
//   - options - ItemsClientDeleteKQLDatabaseOptions contains the optional parameters for the ItemsClient.DeleteKQLDatabase method.
func (client *ItemsClient) DeleteKQLDatabase(ctx context.Context, workspaceID string, kqlDatabaseID string, options *ItemsClientDeleteKQLDatabaseOptions) (ItemsClientDeleteKQLDatabaseResponse, error) {
	var err error
	const operationName = "kqldatabase.ItemsClient.DeleteKQLDatabase"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteKQLDatabaseCreateRequest(ctx, workspaceID, kqlDatabaseID, options)
	if err != nil {
		return ItemsClientDeleteKQLDatabaseResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ItemsClientDeleteKQLDatabaseResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return ItemsClientDeleteKQLDatabaseResponse{}, err
	}
	return ItemsClientDeleteKQLDatabaseResponse{}, nil
}

// deleteKQLDatabaseCreateRequest creates the DeleteKQLDatabase request.
func (client *ItemsClient) deleteKQLDatabaseCreateRequest(ctx context.Context, workspaceID string, kqlDatabaseID string, _ *ItemsClientDeleteKQLDatabaseOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/kqlDatabases/{kqlDatabaseId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if kqlDatabaseID == "" {
		return nil, errors.New("parameter kqlDatabaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kqlDatabaseId}", url.PathEscape(kqlDatabaseID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetKQLDatabase - PERMISSIONS The caller must have viewer or higher workspace role.
// REQUIRED DELEGATED SCOPES KQLDatabase.Read.All or KQLDatabase.ReadWrite.All or Item.Read.All or Item.ReadWrite.All
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
//   - kqlDatabaseID - The KQL database ID.
//   - options - ItemsClientGetKQLDatabaseOptions contains the optional parameters for the ItemsClient.GetKQLDatabase method.
func (client *ItemsClient) GetKQLDatabase(ctx context.Context, workspaceID string, kqlDatabaseID string, options *ItemsClientGetKQLDatabaseOptions) (ItemsClientGetKQLDatabaseResponse, error) {
	var err error
	const operationName = "kqldatabase.ItemsClient.GetKQLDatabase"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getKQLDatabaseCreateRequest(ctx, workspaceID, kqlDatabaseID, options)
	if err != nil {
		return ItemsClientGetKQLDatabaseResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ItemsClientGetKQLDatabaseResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return ItemsClientGetKQLDatabaseResponse{}, err
	}
	resp, err := client.getKQLDatabaseHandleResponse(httpResp)
	return resp, err
}

// getKQLDatabaseCreateRequest creates the GetKQLDatabase request.
func (client *ItemsClient) getKQLDatabaseCreateRequest(ctx context.Context, workspaceID string, kqlDatabaseID string, _ *ItemsClientGetKQLDatabaseOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/kqlDatabases/{kqlDatabaseId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if kqlDatabaseID == "" {
		return nil, errors.New("parameter kqlDatabaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kqlDatabaseId}", url.PathEscape(kqlDatabaseID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getKQLDatabaseHandleResponse handles the GetKQLDatabase response.
func (client *ItemsClient) getKQLDatabaseHandleResponse(resp *http.Response) (ItemsClientGetKQLDatabaseResponse, error) {
	result := ItemsClientGetKQLDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KQLDatabase); err != nil {
		return ItemsClientGetKQLDatabaseResponse{}, err
	}
	return result, nil
}

// NewListKQLDatabasesPager - This API supports pagination [/rest/api/fabric/articles/pagination].
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
//   - options - ItemsClientListKQLDatabasesOptions contains the optional parameters for the ItemsClient.NewListKQLDatabasesPager
//     method.
func (client *ItemsClient) NewListKQLDatabasesPager(workspaceID string, options *ItemsClientListKQLDatabasesOptions) *runtime.Pager[ItemsClientListKQLDatabasesResponse] {
	return runtime.NewPager(runtime.PagingHandler[ItemsClientListKQLDatabasesResponse]{
		More: func(page ItemsClientListKQLDatabasesResponse) bool {
			return page.ContinuationURI != nil && len(*page.ContinuationURI) > 0
		},
		Fetcher: func(ctx context.Context, page *ItemsClientListKQLDatabasesResponse) (ItemsClientListKQLDatabasesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "kqldatabase.ItemsClient.NewListKQLDatabasesPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ContinuationURI
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listKQLDatabasesCreateRequest(ctx, workspaceID, options)
			}, nil)
			if err != nil {
				return ItemsClientListKQLDatabasesResponse{}, err
			}
			return client.listKQLDatabasesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listKQLDatabasesCreateRequest creates the ListKQLDatabases request.
func (client *ItemsClient) listKQLDatabasesCreateRequest(ctx context.Context, workspaceID string, options *ItemsClientListKQLDatabasesOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/kqlDatabases"
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

// listKQLDatabasesHandleResponse handles the ListKQLDatabases response.
func (client *ItemsClient) listKQLDatabasesHandleResponse(resp *http.Response) (ItemsClientListKQLDatabasesResponse, error) {
	result := ItemsClientListKQLDatabasesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KQLDatabases); err != nil {
		return ItemsClientListKQLDatabasesResponse{}, err
	}
	return result, nil
}

// UpdateKQLDatabase - PERMISSIONS The caller must have contributor or higher workspace role.
// REQUIRED DELEGATED SCOPES KQLDatabase.ReadWrite.All or Item.ReadWrite.All
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
//   - kqlDatabaseID - The KQL database ID.
//   - updateKQLDatabaseRequest - Update KQL database request payload.
//   - options - ItemsClientUpdateKQLDatabaseOptions contains the optional parameters for the ItemsClient.UpdateKQLDatabase method.
func (client *ItemsClient) UpdateKQLDatabase(ctx context.Context, workspaceID string, kqlDatabaseID string, updateKQLDatabaseRequest UpdateKQLDatabaseRequest, options *ItemsClientUpdateKQLDatabaseOptions) (ItemsClientUpdateKQLDatabaseResponse, error) {
	var err error
	const operationName = "kqldatabase.ItemsClient.UpdateKQLDatabase"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateKQLDatabaseCreateRequest(ctx, workspaceID, kqlDatabaseID, updateKQLDatabaseRequest, options)
	if err != nil {
		return ItemsClientUpdateKQLDatabaseResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ItemsClientUpdateKQLDatabaseResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return ItemsClientUpdateKQLDatabaseResponse{}, err
	}
	resp, err := client.updateKQLDatabaseHandleResponse(httpResp)
	return resp, err
}

// updateKQLDatabaseCreateRequest creates the UpdateKQLDatabase request.
func (client *ItemsClient) updateKQLDatabaseCreateRequest(ctx context.Context, workspaceID string, kqlDatabaseID string, updateKQLDatabaseRequest UpdateKQLDatabaseRequest, _ *ItemsClientUpdateKQLDatabaseOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/kqlDatabases/{kqlDatabaseId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if kqlDatabaseID == "" {
		return nil, errors.New("parameter kqlDatabaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kqlDatabaseId}", url.PathEscape(kqlDatabaseID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, updateKQLDatabaseRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// updateKQLDatabaseHandleResponse handles the UpdateKQLDatabase response.
func (client *ItemsClient) updateKQLDatabaseHandleResponse(resp *http.Response) (ItemsClientUpdateKQLDatabaseResponse, error) {
	result := ItemsClientUpdateKQLDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KQLDatabase); err != nil {
		return ItemsClientUpdateKQLDatabaseResponse{}, err
	}
	return result, nil
}

// Custom code starts below

// CreateKQLDatabase - returns ItemsClientCreateKQLDatabaseResponse in sync mode.
// This API supports long running operations (LRO) [/rest/api/fabric/articles/long-running-operation].
//
// PERMISSIONS THE CALLER MUST HAVE CONTRIBUTOR OR HIGHER WORKSPACE ROLE.
// REQUIRED DELEGATED SCOPES KQLDatabase.ReadWrite.All or Item.ReadWrite.All
//
// LIMITATIONS
//
//   - To create a KQL database the workspace must be on a supported Fabric capacity. For more information see: Microsoft Fabric license types [/fabric/enterprise/licenses#microsoft-fabric-license-types].
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] | Yes | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
//
// INTERFACE
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - createKQLDatabaseRequest - Create item request payload.
//   - options - ItemsClientBeginCreateKQLDatabaseOptions contains the optional parameters for the ItemsClient.BeginCreateKQLDatabase method.
func (client *ItemsClient) CreateKQLDatabase(ctx context.Context, workspaceID string, createKQLDatabaseRequest CreateKQLDatabaseRequest, options *ItemsClientBeginCreateKQLDatabaseOptions) (ItemsClientCreateKQLDatabaseResponse, error) {
	result, err := iruntime.NewLRO(client.BeginCreateKQLDatabase(ctx, workspaceID, createKQLDatabaseRequest, options)).Sync(ctx)
	if err != nil {
		var azcoreRespError *azcore.ResponseError
		if errors.As(err, &azcoreRespError) {
			return ItemsClientCreateKQLDatabaseResponse{}, core.NewResponseError(azcoreRespError.RawResponse)
		}
		return ItemsClientCreateKQLDatabaseResponse{}, err
	}
	return result, err
}

// beginCreateKQLDatabase creates the createKQLDatabase request.
func (client *ItemsClient) beginCreateKQLDatabase(ctx context.Context, workspaceID string, createKQLDatabaseRequest CreateKQLDatabaseRequest, options *ItemsClientBeginCreateKQLDatabaseOptions) (*runtime.Poller[ItemsClientCreateKQLDatabaseResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createKQLDatabase(ctx, workspaceID, createKQLDatabaseRequest, options)
		if err != nil {
			var azcoreRespError *azcore.ResponseError
			if errors.As(err, &azcoreRespError) {
				return nil, core.NewResponseError(azcoreRespError.RawResponse)
			}
			return nil, err
		}
		handler, err := locasync.NewPollerHandler[ItemsClientCreateKQLDatabaseResponse](client.internal.Pipeline(), resp, runtime.FinalStateViaAzureAsyncOp)
		if err != nil {
			var azcoreRespError *azcore.ResponseError
			if errors.As(err, &azcoreRespError) {
				return nil, core.NewResponseError(azcoreRespError.RawResponse)
			}
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ItemsClientCreateKQLDatabaseResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Handler:       handler,
			Tracer:        client.internal.Tracer(),
		})
	} else {
		handler, err := locasync.NewPollerHandler[ItemsClientCreateKQLDatabaseResponse](client.internal.Pipeline(), nil, runtime.FinalStateViaAzureAsyncOp)
		if err != nil {
			var azcoreRespError *azcore.ResponseError
			if errors.As(err, &azcoreRespError) {
				return nil, core.NewResponseError(azcoreRespError.RawResponse)
			}
			return nil, err
		}
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ItemsClientCreateKQLDatabaseResponse]{
			Handler: handler,
			Tracer:  client.internal.Tracer(),
		})
	}
}

// ListKQLDatabases - returns array of KQLDatabase from all pages.
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
//   - options - ItemsClientListKQLDatabasesOptions contains the optional parameters for the ItemsClient.NewListKQLDatabasesPager method.
func (client *ItemsClient) ListKQLDatabases(ctx context.Context, workspaceID string, options *ItemsClientListKQLDatabasesOptions) ([]KQLDatabase, error) {
	pager := client.NewListKQLDatabasesPager(workspaceID, options)
	mapper := func(resp ItemsClientListKQLDatabasesResponse) []KQLDatabase {
		return resp.Value
	}
	list, err := iruntime.NewPageIterator(ctx, pager, mapper).Get()
	if err != nil {
		var azcoreRespError *azcore.ResponseError
		if errors.As(err, &azcoreRespError) {
			return []KQLDatabase{}, core.NewResponseError(azcoreRespError.RawResponse)
		}
		return []KQLDatabase{}, err
	}
	return list, nil
}
