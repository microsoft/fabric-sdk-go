// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package sqlendpoint

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

// ItemsClient contains the methods for the Items group.
// Don't use this type directly, use a constructor function instead.
type ItemsClient struct {
	internal *azcore.Client
	endpoint string
}

// NewListSQLEndpointsPager - This API supports pagination [/rest/api/fabric/articles/pagination].
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
//   - options - ItemsClientListSQLEndpointsOptions contains the optional parameters for the ItemsClient.NewListSQLEndpointsPager
//     method.
func (client *ItemsClient) NewListSQLEndpointsPager(workspaceID string, options *ItemsClientListSQLEndpointsOptions) *runtime.Pager[ItemsClientListSQLEndpointsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ItemsClientListSQLEndpointsResponse]{
		More: func(page ItemsClientListSQLEndpointsResponse) bool {
			return page.ContinuationURI != nil && len(*page.ContinuationURI) > 0
		},
		Fetcher: func(ctx context.Context, page *ItemsClientListSQLEndpointsResponse) (ItemsClientListSQLEndpointsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "sqlendpoint.ItemsClient.NewListSQLEndpointsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ContinuationURI
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listSQLEndpointsCreateRequest(ctx, workspaceID, options)
			}, nil)
			if err != nil {
				return ItemsClientListSQLEndpointsResponse{}, err
			}
			return client.listSQLEndpointsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listSQLEndpointsCreateRequest creates the ListSQLEndpoints request.
func (client *ItemsClient) listSQLEndpointsCreateRequest(ctx context.Context, workspaceID string, options *ItemsClientListSQLEndpointsOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/sqlEndpoints"
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

// listSQLEndpointsHandleResponse handles the ListSQLEndpoints response.
func (client *ItemsClient) listSQLEndpointsHandleResponse(resp *http.Response) (ItemsClientListSQLEndpointsResponse, error) {
	result := ItemsClientListSQLEndpointsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLEndpoints); err != nil {
		return ItemsClientListSQLEndpointsResponse{}, err
	}
	return result, nil
}

// Custom code starts below

// ListSQLEndpoints - returns array of SQLEndpoint from all pages.
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
//   - options - ItemsClientListSQLEndpointsOptions contains the optional parameters for the ItemsClient.NewListSQLEndpointsPager method.
func (client *ItemsClient) ListSQLEndpoints(ctx context.Context, workspaceID string, options *ItemsClientListSQLEndpointsOptions) ([]SQLEndpoint, error) {
	pager := client.NewListSQLEndpointsPager(workspaceID, options)
	mapper := func(resp ItemsClientListSQLEndpointsResponse) []SQLEndpoint {
		return resp.Value
	}
	list, err := iruntime.NewPageIterator(ctx, pager, mapper).Get()
	if err != nil {
		return nil, err
	}
	return list, nil
}