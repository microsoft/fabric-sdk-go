// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package spark

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
)

// WorkspaceSettingsClient contains the methods for the WorkspaceSettings group.
// Don't use this type directly, use a constructor function instead.
type WorkspaceSettingsClient struct {
	internal *azcore.Client
	endpoint string
}

// GetSparkSettings - PERMISSIONS The caller must have viewer or higher workspace role.
// REQUIRED DELEGATED SCOPES Workspace.Read.All or Workspace.ReadWrite.All
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
//   - options - WorkspaceSettingsClientGetSparkSettingsOptions contains the optional parameters for the WorkspaceSettingsClient.GetSparkSettings
//     method.
func (client *WorkspaceSettingsClient) GetSparkSettings(ctx context.Context, workspaceID string, options *WorkspaceSettingsClientGetSparkSettingsOptions) (WorkspaceSettingsClientGetSparkSettingsResponse, error) {
	var err error
	const operationName = "spark.WorkspaceSettingsClient.GetSparkSettings"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getSparkSettingsCreateRequest(ctx, workspaceID, options)
	if err != nil {
		return WorkspaceSettingsClientGetSparkSettingsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceSettingsClientGetSparkSettingsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return WorkspaceSettingsClientGetSparkSettingsResponse{}, err
	}
	resp, err := client.getSparkSettingsHandleResponse(httpResp)
	return resp, err
}

// getSparkSettingsCreateRequest creates the GetSparkSettings request.
func (client *WorkspaceSettingsClient) getSparkSettingsCreateRequest(ctx context.Context, workspaceID string, _ *WorkspaceSettingsClientGetSparkSettingsOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/spark/settings"
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

// getSparkSettingsHandleResponse handles the GetSparkSettings response.
func (client *WorkspaceSettingsClient) getSparkSettingsHandleResponse(resp *http.Response) (WorkspaceSettingsClientGetSparkSettingsResponse, error) {
	result := WorkspaceSettingsClientGetSparkSettingsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceSparkSettings); err != nil {
		return WorkspaceSettingsClientGetSparkSettingsResponse{}, err
	}
	return result, nil
}

// UpdateSparkSettings - PERMISSIONS The caller must have admin workspace role.
// REQUIRED DELEGATED SCOPES Workspace.ReadWrite.All
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
//   - updateWorkspaceSettingsRequest - Update workspace Spark settings request payload.
//   - options - WorkspaceSettingsClientUpdateSparkSettingsOptions contains the optional parameters for the WorkspaceSettingsClient.UpdateSparkSettings
//     method.
func (client *WorkspaceSettingsClient) UpdateSparkSettings(ctx context.Context, workspaceID string, updateWorkspaceSettingsRequest UpdateWorkspaceSparkSettingsRequest, options *WorkspaceSettingsClientUpdateSparkSettingsOptions) (WorkspaceSettingsClientUpdateSparkSettingsResponse, error) {
	var err error
	const operationName = "spark.WorkspaceSettingsClient.UpdateSparkSettings"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateSparkSettingsCreateRequest(ctx, workspaceID, updateWorkspaceSettingsRequest, options)
	if err != nil {
		return WorkspaceSettingsClientUpdateSparkSettingsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceSettingsClientUpdateSparkSettingsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return WorkspaceSettingsClientUpdateSparkSettingsResponse{}, err
	}
	resp, err := client.updateSparkSettingsHandleResponse(httpResp)
	return resp, err
}

// updateSparkSettingsCreateRequest creates the UpdateSparkSettings request.
func (client *WorkspaceSettingsClient) updateSparkSettingsCreateRequest(ctx context.Context, workspaceID string, updateWorkspaceSettingsRequest UpdateWorkspaceSparkSettingsRequest, _ *WorkspaceSettingsClientUpdateSparkSettingsOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/spark/settings"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, updateWorkspaceSettingsRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// updateSparkSettingsHandleResponse handles the UpdateSparkSettings response.
func (client *WorkspaceSettingsClient) updateSparkSettingsHandleResponse(resp *http.Response) (WorkspaceSettingsClientUpdateSparkSettingsResponse, error) {
	result := WorkspaceSettingsClientUpdateSparkSettingsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceSparkSettings); err != nil {
		return WorkspaceSettingsClientUpdateSparkSettingsResponse{}, err
	}
	return result, nil
}

// Custom code starts below
