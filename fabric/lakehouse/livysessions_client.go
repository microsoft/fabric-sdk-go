// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package lakehouse

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

// LivySessionsClient contains the methods for the LivySessions group.
// Don't use this type directly, use a constructor function instead.
type LivySessionsClient struct {
	internal *azcore.Client
	endpoint string
}

// GetLivySession - PERMISSIONS The caller must have viewer or higher workspace role.
// REQUIRED DELEGATED SCOPES Lakehouse.Read.All or Lakehouse.ReadWrite.All or Item.Read.All or Item.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - workspaceID - The workspace identifier.
//   - lakehouseID - The lakehouse ID.
//   - livyID - The session identifier.
//   - options - LivySessionsClientGetLivySessionOptions contains the optional parameters for the LivySessionsClient.GetLivySession
//     method.
func (client *LivySessionsClient) GetLivySession(ctx context.Context, workspaceID string, lakehouseID string, livyID string, options *LivySessionsClientGetLivySessionOptions) (LivySessionsClientGetLivySessionResponse, error) {
	var err error
	const operationName = "lakehouse.LivySessionsClient.GetLivySession"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getLivySessionCreateRequest(ctx, workspaceID, lakehouseID, livyID, options)
	if err != nil {
		return LivySessionsClientGetLivySessionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LivySessionsClientGetLivySessionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return LivySessionsClientGetLivySessionResponse{}, err
	}
	resp, err := client.getLivySessionHandleResponse(httpResp)
	return resp, err
}

// getLivySessionCreateRequest creates the GetLivySession request.
func (client *LivySessionsClient) getLivySessionCreateRequest(ctx context.Context, workspaceID string, lakehouseID string, livyID string, _ *LivySessionsClientGetLivySessionOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/lakehouses/{lakehouseId}/livySessions/{livyId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if lakehouseID == "" {
		return nil, errors.New("parameter lakehouseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{lakehouseId}", url.PathEscape(lakehouseID))
	if livyID == "" {
		return nil, errors.New("parameter livyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{livyId}", url.PathEscape(livyID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getLivySessionHandleResponse handles the GetLivySession response.
func (client *LivySessionsClient) getLivySessionHandleResponse(resp *http.Response) (LivySessionsClientGetLivySessionResponse, error) {
	result := LivySessionsClientGetLivySessionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LivySession); err != nil {
		return LivySessionsClientGetLivySessionResponse{}, err
	}
	return result, nil
}

// NewListLivySessionsPager - This API supports pagination [/rest/api/fabric/articles/pagination]. A maximum of 100 records
// can be returned per request. With the URI provided in the response, you can get the next page of records.
// PERMISSIONS The caller must have viewer or higher workspace role.
// REQUIRED DELEGATED SCOPES Lakehouse.Read.All or Lakehouse.ReadWrite.All or Item.Read.All or Item.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
//
// Generated from API version v1
//   - workspaceID - The workspace identifier.
//   - lakehouseID - The lakehouse ID.
//   - options - LivySessionsClientListLivySessionsOptions contains the optional parameters for the LivySessionsClient.NewListLivySessionsPager
//     method.
func (client *LivySessionsClient) NewListLivySessionsPager(workspaceID string, lakehouseID string, options *LivySessionsClientListLivySessionsOptions) *runtime.Pager[LivySessionsClientListLivySessionsResponse] {
	return runtime.NewPager(runtime.PagingHandler[LivySessionsClientListLivySessionsResponse]{
		More: func(page LivySessionsClientListLivySessionsResponse) bool {
			return page.ContinuationURI != nil && len(*page.ContinuationURI) > 0
		},
		Fetcher: func(ctx context.Context, page *LivySessionsClientListLivySessionsResponse) (LivySessionsClientListLivySessionsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "lakehouse.LivySessionsClient.NewListLivySessionsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ContinuationURI
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listLivySessionsCreateRequest(ctx, workspaceID, lakehouseID, options)
			}, nil)
			if err != nil {
				return LivySessionsClientListLivySessionsResponse{}, err
			}
			return client.listLivySessionsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listLivySessionsCreateRequest creates the ListLivySessions request.
func (client *LivySessionsClient) listLivySessionsCreateRequest(ctx context.Context, workspaceID string, lakehouseID string, options *LivySessionsClientListLivySessionsOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/lakehouses/{lakehouseId}/livySessions"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if lakehouseID == "" {
		return nil, errors.New("parameter lakehouseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{lakehouseId}", url.PathEscape(lakehouseID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	if options != nil && options.MaxResults != nil {
		reqQP.Set("maxResults", strconv.FormatInt(int64(*options.MaxResults), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listLivySessionsHandleResponse handles the ListLivySessions response.
func (client *LivySessionsClient) listLivySessionsHandleResponse(resp *http.Response) (LivySessionsClientListLivySessionsResponse, error) {
	result := LivySessionsClientListLivySessionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LivySessions); err != nil {
		return LivySessionsClientListLivySessionsResponse{}, err
	}
	return result, nil
}

// Custom code starts below

// ListLivySessions - returns array of LivySession from all pages.
// This API supports pagination [/rest/api/fabric/articles/pagination]. A maximum of 100 records can be returned per request. With the URI provided in the response, you can get the next page of records.
//
// PERMISSIONS The caller must have viewer or higher workspace role.
//
// # REQUIRED DELEGATED SCOPES Lakehouse.Read.All or Lakehouse.ReadWrite.All or Item.Read.All or Item.ReadWrite.All
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
//
// INTERFACE
// Generated from API version v1
//   - workspaceID - The workspace identifier.
//   - lakehouseID - The lakehouse ID.
//   - options - LivySessionsClientListLivySessionsOptions contains the optional parameters for the LivySessionsClient.NewListLivySessionsPager method.
func (client *LivySessionsClient) ListLivySessions(ctx context.Context, workspaceID string, lakehouseID string, options *LivySessionsClientListLivySessionsOptions) ([]LivySession, error) {
	pager := client.NewListLivySessionsPager(workspaceID, lakehouseID, options)
	mapper := func(resp LivySessionsClientListLivySessionsResponse) []LivySession {
		return resp.Value
	}
	list, err := iruntime.NewPageIterator(ctx, pager, mapper).Get()
	if err != nil {
		var azcoreRespError *azcore.ResponseError
		if errors.As(err, &azcoreRespError) {
			return []LivySession{}, core.NewResponseError(azcoreRespError.RawResponse)
		}
		return []LivySession{}, err
	}
	return list, nil
}
