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

// ExternalDataSharesClient contains the methods for the ExternalDataShares group.
// Don't use this type directly, use a constructor function instead.
type ExternalDataSharesClient struct {
	internal *azcore.Client
	endpoint string
}

// NewListExternalDataSharesPager - This API supports pagination [/rest/api/fabric/articles/pagination]. A maximum of 10,000
// records can be returned per request. With the continuous token provided in the response, you can get the next
// 10,000 records.
// PERMISSIONS The caller must have administrator rights (such as Microsoft 365 global administrator or Microsoft Fabric administrator)
// or authenticate using a service principal.
// REQUIRED DELEGATED SCOPES Tenant.ReadWrite.All
// REQUIRED TENANT SETTINGS To use this API, enable the 'External data sharing' admin switch for the calling principal.
// LIMITATIONS Maximum 10 requests per minute.
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | No | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | No |
// INTERFACE
//
// Generated from API version v1
//   - options - ExternalDataSharesClientListExternalDataSharesOptions contains the optional parameters for the ExternalDataSharesClient.NewListExternalDataSharesPager
//     method.
func (client *ExternalDataSharesClient) NewListExternalDataSharesPager(options *ExternalDataSharesClientListExternalDataSharesOptions) *runtime.Pager[ExternalDataSharesClientListExternalDataSharesResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExternalDataSharesClientListExternalDataSharesResponse]{
		More: func(page ExternalDataSharesClientListExternalDataSharesResponse) bool {
			return page.ContinuationURI != nil && len(*page.ContinuationURI) > 0
		},
		Fetcher: func(ctx context.Context, page *ExternalDataSharesClientListExternalDataSharesResponse) (ExternalDataSharesClientListExternalDataSharesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "admin.ExternalDataSharesClient.NewListExternalDataSharesPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ContinuationURI
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listExternalDataSharesCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ExternalDataSharesClientListExternalDataSharesResponse{}, err
			}
			return client.listExternalDataSharesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listExternalDataSharesCreateRequest creates the ListExternalDataShares request.
func (client *ExternalDataSharesClient) listExternalDataSharesCreateRequest(ctx context.Context, options *ExternalDataSharesClientListExternalDataSharesOptions) (*policy.Request, error) {
	urlPath := "/v1/admin/items/externalDataShares"
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

// listExternalDataSharesHandleResponse handles the ListExternalDataShares response.
func (client *ExternalDataSharesClient) listExternalDataSharesHandleResponse(resp *http.Response) (ExternalDataSharesClientListExternalDataSharesResponse, error) {
	result := ExternalDataSharesClientListExternalDataSharesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExternalDataShares); err != nil {
		return ExternalDataSharesClientListExternalDataSharesResponse{}, err
	}
	return result, nil
}

// RevokeExternalDataShare - PERMISSIONS The caller must have administrator rights (such as Microsoft 365 global administrator
// or Microsoft Fabric administrator).
// REQUIRED DELEGATED SCOPES Tenant.ReadWrite.All
// REQUIRED TENANT SETTINGS To use this API, enable the 'External data sharing' admin switch for the calling principal.
// LIMITATIONS Maximum 10 requests per minute.
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
//   - externalDataShareID - The external data share ID.
//   - options - ExternalDataSharesClientRevokeExternalDataShareOptions contains the optional parameters for the ExternalDataSharesClient.RevokeExternalDataShare
//     method.
func (client *ExternalDataSharesClient) RevokeExternalDataShare(ctx context.Context, workspaceID string, itemID string, externalDataShareID string, options *ExternalDataSharesClientRevokeExternalDataShareOptions) (ExternalDataSharesClientRevokeExternalDataShareResponse, error) {
	var err error
	const operationName = "admin.ExternalDataSharesClient.RevokeExternalDataShare"
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
		err = core.NewResponseError(httpResp)
		return ExternalDataSharesClientRevokeExternalDataShareResponse{}, err
	}
	return ExternalDataSharesClientRevokeExternalDataShareResponse{}, nil
}

// revokeExternalDataShareCreateRequest creates the RevokeExternalDataShare request.
func (client *ExternalDataSharesClient) revokeExternalDataShareCreateRequest(ctx context.Context, workspaceID string, itemID string, externalDataShareID string, options *ExternalDataSharesClientRevokeExternalDataShareOptions) (*policy.Request, error) {
	urlPath := "/v1/admin/workspaces/{workspaceId}/items/{itemId}/externalDataShares/{externalDataShareId}/revoke"
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

// ListExternalDataShares - returns array of ExternalDataShare from all pages.
// This API supports pagination [/rest/api/fabric/articles/pagination]. A maximum of 10,000 records can be returned per request. With the continuous token provided in the response, you can get the next
// 10,000 records.
//
// PERMISSIONS The caller must have administrator rights (such as Microsoft 365 global administrator or Microsoft Fabric administrator) or authenticate using a service principal.
//
// # REQUIRED DELEGATED SCOPES Tenant.ReadWrite.All
//
// REQUIRED TENANT SETTINGS To use this API, enable the 'External data sharing' admin switch for the calling principal.
//
// LIMITATIONS Maximum 10 requests per minute.
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] | No | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | No |
//
// INTERFACE
// Generated from API version v1
//   - options - ExternalDataSharesClientListExternalDataSharesOptions contains the optional parameters for the ExternalDataSharesClient.NewListExternalDataSharesPager method.
func (client *ExternalDataSharesClient) ListExternalDataShares(ctx context.Context, options *ExternalDataSharesClientListExternalDataSharesOptions) ([]ExternalDataShare, error) {
	pager := client.NewListExternalDataSharesPager(options)
	mapper := func(resp ExternalDataSharesClientListExternalDataSharesResponse) []ExternalDataShare {
		return resp.Value
	}
	list, err := iruntime.NewPageIterator(ctx, pager, mapper).Get()
	if err != nil {
		return nil, err
	}
	return list, nil
}
