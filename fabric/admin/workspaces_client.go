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

// WorkspacesClient contains the methods for the Workspaces group.
// Don't use this type directly, use a constructor function instead.
type WorkspacesClient struct {
	internal *azcore.Client
	endpoint string
}

// GetWorkspace - PERMISSIONS The caller must have administrator rights (such as Office 365 Global administrator or Fabric
// administrator) or authenticate using a service principal.
// REQUIRED DELEGATED SCOPES Tenant.Read.All or Tenant.ReadWrite.All
// LIMITATIONS Maximum 200 requests per hour.
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
//   - options - WorkspacesClientGetWorkspaceOptions contains the optional parameters for the WorkspacesClient.GetWorkspace method.
func (client *WorkspacesClient) GetWorkspace(ctx context.Context, workspaceID string, options *WorkspacesClientGetWorkspaceOptions) (WorkspacesClientGetWorkspaceResponse, error) {
	var err error
	const operationName = "admin.WorkspacesClient.GetWorkspace"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getWorkspaceCreateRequest(ctx, workspaceID, options)
	if err != nil {
		return WorkspacesClientGetWorkspaceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspacesClientGetWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return WorkspacesClientGetWorkspaceResponse{}, err
	}
	resp, err := client.getWorkspaceHandleResponse(httpResp)
	return resp, err
}

// getWorkspaceCreateRequest creates the GetWorkspace request.
func (client *WorkspacesClient) getWorkspaceCreateRequest(ctx context.Context, workspaceID string, options *WorkspacesClientGetWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/v1/admin/workspaces/{workspaceId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getWorkspaceHandleResponse handles the GetWorkspace response.
func (client *WorkspacesClient) getWorkspaceHandleResponse(resp *http.Response) (WorkspacesClientGetWorkspaceResponse, error) {
	result := WorkspacesClientGetWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Workspace); err != nil {
		return WorkspacesClientGetWorkspaceResponse{}, err
	}
	return result, nil
}

// NewListGitConnectionsPager - This API supports pagination [/rest/api/fabric/articles/pagination]. A maximum of 1,000 records
// can be returned per request. With the continuous token provided in the response, you can get the next
// 1,000 records.
// PERMISSIONS The caller must be a Fabric administrator or authenticate using a service principal.
// REQUIRED DELEGATED SCOPES Tenant.Read.All or Tenant.ReadWrite.All
// LIMITATIONS Maximum 25 requests per minute.
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | Yes | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes | |
// INTERFACE
//
// Generated from API version v1
//   - options - WorkspacesClientListGitConnectionsOptions contains the optional parameters for the WorkspacesClient.NewListGitConnectionsPager
//     method.
func (client *WorkspacesClient) NewListGitConnectionsPager(options *WorkspacesClientListGitConnectionsOptions) *runtime.Pager[WorkspacesClientListGitConnectionsResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspacesClientListGitConnectionsResponse]{
		More: func(page WorkspacesClientListGitConnectionsResponse) bool {
			return page.ContinuationURI != nil && len(*page.ContinuationURI) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspacesClientListGitConnectionsResponse) (WorkspacesClientListGitConnectionsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "admin.WorkspacesClient.NewListGitConnectionsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ContinuationURI
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listGitConnectionsCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return WorkspacesClientListGitConnectionsResponse{}, err
			}
			return client.listGitConnectionsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listGitConnectionsCreateRequest creates the ListGitConnections request.
func (client *WorkspacesClient) listGitConnectionsCreateRequest(ctx context.Context, options *WorkspacesClientListGitConnectionsOptions) (*policy.Request, error) {
	urlPath := "/v1/admin/workspaces/discoverGitConnections"
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

// listGitConnectionsHandleResponse handles the ListGitConnections response.
func (client *WorkspacesClient) listGitConnectionsHandleResponse(resp *http.Response) (WorkspacesClientListGitConnectionsResponse, error) {
	result := WorkspacesClientListGitConnectionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GitConnections); err != nil {
		return WorkspacesClientListGitConnectionsResponse{}, err
	}
	return result, nil
}

// ListWorkspaceAccessDetails - PERMISSIONS The caller must have administrator rights (such as Office 365 Global administrator
// or Fabric administrator) or authenticate using a service principal.
// REQUIRED DELEGATED SCOPES Tenant.Read.All or Tenant.ReadWrite.All
// LIMITATIONS Maximum 200 requests per hour.
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
//   - options - WorkspacesClientListWorkspaceAccessDetailsOptions contains the optional parameters for the WorkspacesClient.ListWorkspaceAccessDetails
//     method.
func (client *WorkspacesClient) ListWorkspaceAccessDetails(ctx context.Context, workspaceID string, options *WorkspacesClientListWorkspaceAccessDetailsOptions) (WorkspacesClientListWorkspaceAccessDetailsResponse, error) {
	var err error
	const operationName = "admin.WorkspacesClient.ListWorkspaceAccessDetails"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listWorkspaceAccessDetailsCreateRequest(ctx, workspaceID, options)
	if err != nil {
		return WorkspacesClientListWorkspaceAccessDetailsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspacesClientListWorkspaceAccessDetailsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return WorkspacesClientListWorkspaceAccessDetailsResponse{}, err
	}
	resp, err := client.listWorkspaceAccessDetailsHandleResponse(httpResp)
	return resp, err
}

// listWorkspaceAccessDetailsCreateRequest creates the ListWorkspaceAccessDetails request.
func (client *WorkspacesClient) listWorkspaceAccessDetailsCreateRequest(ctx context.Context, workspaceID string, options *WorkspacesClientListWorkspaceAccessDetailsOptions) (*policy.Request, error) {
	urlPath := "/v1/admin/workspaces/{workspaceId}/users"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listWorkspaceAccessDetailsHandleResponse handles the ListWorkspaceAccessDetails response.
func (client *WorkspacesClient) listWorkspaceAccessDetailsHandleResponse(resp *http.Response) (WorkspacesClientListWorkspaceAccessDetailsResponse, error) {
	result := WorkspacesClientListWorkspaceAccessDetailsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceAccessDetailsResponse); err != nil {
		return WorkspacesClientListWorkspaceAccessDetailsResponse{}, err
	}
	return result, nil
}

// NewListWorkspacesPager - This API supports pagination [/rest/api/fabric/articles/pagination]. A maximum of 10,000 records
// can be returned per request. With the continuous token provided in the response, you can get the next
// 10,000 records.
// PERMISSIONS The caller must have administrator rights (such as Office 365 Global administrator or Fabric administrator)
// or authenticate using a service principal.
// REQUIRED DELEGATED SCOPES Tenant.Read.All or Tenant.ReadWrite.All
// LIMITATIONS Maximum 200 requests per hour.
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | Yes | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
//
// Generated from API version v1
//   - options - WorkspacesClientListWorkspacesOptions contains the optional parameters for the WorkspacesClient.NewListWorkspacesPager
//     method.
func (client *WorkspacesClient) NewListWorkspacesPager(options *WorkspacesClientListWorkspacesOptions) *runtime.Pager[WorkspacesClientListWorkspacesResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspacesClientListWorkspacesResponse]{
		More: func(page WorkspacesClientListWorkspacesResponse) bool {
			return page.ContinuationURI != nil && len(*page.ContinuationURI) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspacesClientListWorkspacesResponse) (WorkspacesClientListWorkspacesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "admin.WorkspacesClient.NewListWorkspacesPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ContinuationURI
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listWorkspacesCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return WorkspacesClientListWorkspacesResponse{}, err
			}
			return client.listWorkspacesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listWorkspacesCreateRequest creates the ListWorkspaces request.
func (client *WorkspacesClient) listWorkspacesCreateRequest(ctx context.Context, options *WorkspacesClientListWorkspacesOptions) (*policy.Request, error) {
	urlPath := "/v1/admin/workspaces"
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
	if options != nil && options.Name != nil {
		reqQP.Set("name", *options.Name)
	}
	if options != nil && options.State != nil {
		reqQP.Set("state", *options.State)
	}
	if options != nil && options.Type != nil {
		reqQP.Set("type", *options.Type)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listWorkspacesHandleResponse handles the ListWorkspaces response.
func (client *WorkspacesClient) listWorkspacesHandleResponse(resp *http.Response) (WorkspacesClientListWorkspacesResponse, error) {
	result := WorkspacesClientListWorkspacesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Workspaces); err != nil {
		return WorkspacesClientListWorkspacesResponse{}, err
	}
	return result, nil
}

// Custom code starts below

// ListGitConnections - returns array of GitConnectionDetails from all pages.
// This API supports pagination [/rest/api/fabric/articles/pagination]. A maximum of 1,000 records can be returned per request. With the continuous token provided in the response, you can get the next
// 1,000 records.
//
// PERMISSIONS The caller must be a Fabric administrator or authenticate using a service principal.
//
// # REQUIRED DELEGATED SCOPES Tenant.Read.All or Tenant.ReadWrite.All
//
// LIMITATIONS Maximum 25 requests per minute.
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] | Yes | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes | |
//
// INTERFACE
// Generated from API version v1
//   - options - WorkspacesClientListGitConnectionsOptions contains the optional parameters for the WorkspacesClient.NewListGitConnectionsPager method.
func (client *WorkspacesClient) ListGitConnections(ctx context.Context, options *WorkspacesClientListGitConnectionsOptions) ([]GitConnectionDetails, error) {
	pager := client.NewListGitConnectionsPager(options)
	mapper := func(resp WorkspacesClientListGitConnectionsResponse) []GitConnectionDetails {
		return resp.Value
	}
	list, err := iruntime.NewPageIterator(ctx, pager, mapper).Get()
	if err != nil {
		var azcoreRespError *azcore.ResponseError
		if errors.As(err, &azcoreRespError) {
			return []GitConnectionDetails{}, core.NewResponseError(azcoreRespError.RawResponse)
		}
		return []GitConnectionDetails{}, err
	}
	return list, nil
}

// ListWorkspaces - returns array of Workspace from all pages.
// This API supports pagination [/rest/api/fabric/articles/pagination]. A maximum of 10,000 records can be returned per request. With the continuous token provided in the response, you can get the next
// 10,000 records.
//
// PERMISSIONS The caller must have administrator rights (such as Office 365 Global administrator or Fabric administrator) or authenticate using a service principal.
//
// # REQUIRED DELEGATED SCOPES Tenant.Read.All or Tenant.ReadWrite.All
//
// LIMITATIONS Maximum 200 requests per hour.
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] | Yes | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
//
// INTERFACE
// Generated from API version v1
//   - options - WorkspacesClientListWorkspacesOptions contains the optional parameters for the WorkspacesClient.NewListWorkspacesPager method.
func (client *WorkspacesClient) ListWorkspaces(ctx context.Context, options *WorkspacesClientListWorkspacesOptions) ([]Workspace, error) {
	pager := client.NewListWorkspacesPager(options)
	mapper := func(resp WorkspacesClientListWorkspacesResponse) []Workspace {
		return resp.Workspaces.Workspaces
	}
	list, err := iruntime.NewPageIterator(ctx, pager, mapper).Get()
	if err != nil {
		var azcoreRespError *azcore.ResponseError
		if errors.As(err, &azcoreRespError) {
			return []Workspace{}, core.NewResponseError(azcoreRespError.RawResponse)
		}
		return []Workspace{}, err
	}
	return list, nil
}
