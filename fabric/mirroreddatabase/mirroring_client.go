// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package mirroreddatabase

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

// MirroringClient contains the methods for the Mirroring group.
// Don't use this type directly, use a constructor function instead.
type MirroringClient struct {
	internal *azcore.Client
	endpoint string
}

// GetMirroringStatus - PERMISSIONS The caller must have viewer or higher workspace role.
// REQUIRED DELEGATED SCOPES MirroredDatabase.Read.All or MirroredDatabase.ReadWrite.All or Item.Read.All or Item.ReadWrite.All
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
//   - mirroredDatabaseID - The mirrored database ID.
//   - options - MirroringClientGetMirroringStatusOptions contains the optional parameters for the MirroringClient.GetMirroringStatus
//     method.
func (client *MirroringClient) GetMirroringStatus(ctx context.Context, workspaceID string, mirroredDatabaseID string, options *MirroringClientGetMirroringStatusOptions) (MirroringClientGetMirroringStatusResponse, error) {
	var err error
	const operationName = "mirroreddatabase.MirroringClient.GetMirroringStatus"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getMirroringStatusCreateRequest(ctx, workspaceID, mirroredDatabaseID, options)
	if err != nil {
		return MirroringClientGetMirroringStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MirroringClientGetMirroringStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return MirroringClientGetMirroringStatusResponse{}, err
	}
	resp, err := client.getMirroringStatusHandleResponse(httpResp)
	return resp, err
}

// getMirroringStatusCreateRequest creates the GetMirroringStatus request.
func (client *MirroringClient) getMirroringStatusCreateRequest(ctx context.Context, workspaceID string, mirroredDatabaseID string, _ *MirroringClientGetMirroringStatusOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/mirroredDatabases/{mirroredDatabaseId}/getMirroringStatus"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if mirroredDatabaseID == "" {
		return nil, errors.New("parameter mirroredDatabaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mirroredDatabaseId}", url.PathEscape(mirroredDatabaseID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getMirroringStatusHandleResponse handles the GetMirroringStatus response.
func (client *MirroringClient) getMirroringStatusHandleResponse(resp *http.Response) (MirroringClientGetMirroringStatusResponse, error) {
	result := MirroringClientGetMirroringStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MirroringStatusResponse); err != nil {
		return MirroringClientGetMirroringStatusResponse{}, err
	}
	return result, nil
}

// GetTablesMirroringStatus - his API supports pagination [/rest/api/fabric/articles/pagination].
// PERMISSIONS The caller must have viewer or higher workspace role.
// REQUIRED DELEGATED SCOPES MirroredDatabase.Read.All or MirroredDatabase.ReadWrite.All or Item.Read.All or Item.ReadWrite.All
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
//   - mirroredDatabaseID - The mirrored database ID.
//   - options - MirroringClientGetTablesMirroringStatusOptions contains the optional parameters for the MirroringClient.GetTablesMirroringStatus
//     method.
func (client *MirroringClient) GetTablesMirroringStatus(ctx context.Context, workspaceID string, mirroredDatabaseID string, options *MirroringClientGetTablesMirroringStatusOptions) (MirroringClientGetTablesMirroringStatusResponse, error) {
	var err error
	const operationName = "mirroreddatabase.MirroringClient.GetTablesMirroringStatus"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getTablesMirroringStatusCreateRequest(ctx, workspaceID, mirroredDatabaseID, options)
	if err != nil {
		return MirroringClientGetTablesMirroringStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MirroringClientGetTablesMirroringStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return MirroringClientGetTablesMirroringStatusResponse{}, err
	}
	resp, err := client.getTablesMirroringStatusHandleResponse(httpResp)
	return resp, err
}

// getTablesMirroringStatusCreateRequest creates the GetTablesMirroringStatus request.
func (client *MirroringClient) getTablesMirroringStatusCreateRequest(ctx context.Context, workspaceID string, mirroredDatabaseID string, _ *MirroringClientGetTablesMirroringStatusOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/mirroredDatabases/{mirroredDatabaseId}/getTablesMirroringStatus"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if mirroredDatabaseID == "" {
		return nil, errors.New("parameter mirroredDatabaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mirroredDatabaseId}", url.PathEscape(mirroredDatabaseID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getTablesMirroringStatusHandleResponse handles the GetTablesMirroringStatus response.
func (client *MirroringClient) getTablesMirroringStatusHandleResponse(resp *http.Response) (MirroringClientGetTablesMirroringStatusResponse, error) {
	result := MirroringClientGetTablesMirroringStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TablesMirroringStatusResponse); err != nil {
		return MirroringClientGetTablesMirroringStatusResponse{}, err
	}
	return result, nil
}

// StartMirroring - PERMISSIONS The API caller must have contributor or higher workspace role.
// REQUIRED DELEGATED SCOPES MirroredDatabase.ReadWrite.All or Item.ReadWrite.All
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
//   - mirroredDatabaseID - The mirrored database ID.
//   - options - MirroringClientStartMirroringOptions contains the optional parameters for the MirroringClient.StartMirroring
//     method.
func (client *MirroringClient) StartMirroring(ctx context.Context, workspaceID string, mirroredDatabaseID string, options *MirroringClientStartMirroringOptions) (MirroringClientStartMirroringResponse, error) {
	var err error
	const operationName = "mirroreddatabase.MirroringClient.StartMirroring"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.startMirroringCreateRequest(ctx, workspaceID, mirroredDatabaseID, options)
	if err != nil {
		return MirroringClientStartMirroringResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MirroringClientStartMirroringResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return MirroringClientStartMirroringResponse{}, err
	}
	return MirroringClientStartMirroringResponse{}, nil
}

// startMirroringCreateRequest creates the StartMirroring request.
func (client *MirroringClient) startMirroringCreateRequest(ctx context.Context, workspaceID string, mirroredDatabaseID string, _ *MirroringClientStartMirroringOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/mirroredDatabases/{mirroredDatabaseId}/startMirroring"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if mirroredDatabaseID == "" {
		return nil, errors.New("parameter mirroredDatabaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mirroredDatabaseId}", url.PathEscape(mirroredDatabaseID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// StopMirroring - PERMISSIONS The API caller must have contributor or higher workspace role.
// REQUIRED DELEGATED SCOPES MirroredDatabase.ReadWrite.All or Item.ReadWrite.All
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
//   - mirroredDatabaseID - The mirrored database ID.
//   - options - MirroringClientStopMirroringOptions contains the optional parameters for the MirroringClient.StopMirroring method.
func (client *MirroringClient) StopMirroring(ctx context.Context, workspaceID string, mirroredDatabaseID string, options *MirroringClientStopMirroringOptions) (MirroringClientStopMirroringResponse, error) {
	var err error
	const operationName = "mirroreddatabase.MirroringClient.StopMirroring"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.stopMirroringCreateRequest(ctx, workspaceID, mirroredDatabaseID, options)
	if err != nil {
		return MirroringClientStopMirroringResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MirroringClientStopMirroringResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return MirroringClientStopMirroringResponse{}, err
	}
	return MirroringClientStopMirroringResponse{}, nil
}

// stopMirroringCreateRequest creates the StopMirroring request.
func (client *MirroringClient) stopMirroringCreateRequest(ctx context.Context, workspaceID string, mirroredDatabaseID string, _ *MirroringClientStopMirroringOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/mirroredDatabases/{mirroredDatabaseId}/stopMirroring"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if mirroredDatabaseID == "" {
		return nil, errors.New("parameter mirroredDatabaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mirroredDatabaseId}", url.PathEscape(mirroredDatabaseID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Custom code starts below