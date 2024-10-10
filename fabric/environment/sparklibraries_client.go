// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package environment

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
)

// SparkLibrariesClient contains the methods for the SparkLibraries group.
// Don't use this type directly, use a constructor function instead.
type SparkLibrariesClient struct {
	internal *azcore.Client
	endpoint string
}

// CancelPublish - PERMISSIONS Write permission for the environment item.
// REQUIRED DELEGATED SCOPES Environment.ReadWrite.All
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
//   - environmentID - The environment ID.
//   - options - SparkLibrariesClientCancelPublishOptions contains the optional parameters for the SparkLibrariesClient.CancelPublish
//     method.
func (client *SparkLibrariesClient) CancelPublish(ctx context.Context, workspaceID string, environmentID string, options *SparkLibrariesClientCancelPublishOptions) (SparkLibrariesClientCancelPublishResponse, error) {
	var err error
	const operationName = "environment.SparkLibrariesClient.CancelPublish"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.cancelPublishCreateRequest(ctx, workspaceID, environmentID, options)
	if err != nil {
		return SparkLibrariesClientCancelPublishResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SparkLibrariesClientCancelPublishResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return SparkLibrariesClientCancelPublishResponse{}, err
	}
	resp, err := client.cancelPublishHandleResponse(httpResp)
	return resp, err
}

// cancelPublishCreateRequest creates the CancelPublish request.
func (client *SparkLibrariesClient) cancelPublishCreateRequest(ctx context.Context, workspaceID string, environmentID string, _ *SparkLibrariesClientCancelPublishOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/environments/{environmentId}/staging/cancelPublish"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if environmentID == "" {
		return nil, errors.New("parameter environmentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentId}", url.PathEscape(environmentID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// cancelPublishHandleResponse handles the CancelPublish response.
func (client *SparkLibrariesClient) cancelPublishHandleResponse(resp *http.Response) (SparkLibrariesClientCancelPublishResponse, error) {
	result := SparkLibrariesClientCancelPublishResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PublishInfo); err != nil {
		return SparkLibrariesClientCancelPublishResponse{}, err
	}
	return result, nil
}

// DeleteStagingLibrary - PERMISSIONS Write permission for the environment item.
// REQUIRED DELEGATED SCOPES Environment.ReadWrite.All
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
//   - environmentID - The environment ID.
//   - libraryToDelete - The library file to be deleted. The library name needs to include its extension, for example samplefile.jar.
//   - options - SparkLibrariesClientDeleteStagingLibraryOptions contains the optional parameters for the SparkLibrariesClient.DeleteStagingLibrary
//     method.
func (client *SparkLibrariesClient) DeleteStagingLibrary(ctx context.Context, workspaceID string, environmentID string, libraryToDelete string, options *SparkLibrariesClientDeleteStagingLibraryOptions) (SparkLibrariesClientDeleteStagingLibraryResponse, error) {
	var err error
	const operationName = "environment.SparkLibrariesClient.DeleteStagingLibrary"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteStagingLibraryCreateRequest(ctx, workspaceID, environmentID, libraryToDelete, options)
	if err != nil {
		return SparkLibrariesClientDeleteStagingLibraryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SparkLibrariesClientDeleteStagingLibraryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return SparkLibrariesClientDeleteStagingLibraryResponse{}, err
	}
	return SparkLibrariesClientDeleteStagingLibraryResponse{}, nil
}

// deleteStagingLibraryCreateRequest creates the DeleteStagingLibrary request.
func (client *SparkLibrariesClient) deleteStagingLibraryCreateRequest(ctx context.Context, workspaceID string, environmentID string, libraryToDelete string, _ *SparkLibrariesClientDeleteStagingLibraryOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/environments/{environmentId}/staging/libraries"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if environmentID == "" {
		return nil, errors.New("parameter environmentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentId}", url.PathEscape(environmentID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("libraryToDelete", libraryToDelete)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetPublishedLibraries - PERMISSIONS Read permission for the environment item.
// REQUIRED DELEGATED SCOPES Environment.Read.All or Environment.ReadWrite.All
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
//   - environmentID - The environment ID.
//   - options - SparkLibrariesClientGetPublishedLibrariesOptions contains the optional parameters for the SparkLibrariesClient.GetPublishedLibraries
//     method.
func (client *SparkLibrariesClient) GetPublishedLibraries(ctx context.Context, workspaceID string, environmentID string, options *SparkLibrariesClientGetPublishedLibrariesOptions) (SparkLibrariesClientGetPublishedLibrariesResponse, error) {
	var err error
	const operationName = "environment.SparkLibrariesClient.GetPublishedLibraries"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getPublishedLibrariesCreateRequest(ctx, workspaceID, environmentID, options)
	if err != nil {
		return SparkLibrariesClientGetPublishedLibrariesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SparkLibrariesClientGetPublishedLibrariesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return SparkLibrariesClientGetPublishedLibrariesResponse{}, err
	}
	resp, err := client.getPublishedLibrariesHandleResponse(httpResp)
	return resp, err
}

// getPublishedLibrariesCreateRequest creates the GetPublishedLibraries request.
func (client *SparkLibrariesClient) getPublishedLibrariesCreateRequest(ctx context.Context, workspaceID string, environmentID string, _ *SparkLibrariesClientGetPublishedLibrariesOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/environments/{environmentId}/libraries"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if environmentID == "" {
		return nil, errors.New("parameter environmentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentId}", url.PathEscape(environmentID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getPublishedLibrariesHandleResponse handles the GetPublishedLibraries response.
func (client *SparkLibrariesClient) getPublishedLibrariesHandleResponse(resp *http.Response) (SparkLibrariesClientGetPublishedLibrariesResponse, error) {
	result := SparkLibrariesClientGetPublishedLibrariesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Libraries); err != nil {
		return SparkLibrariesClientGetPublishedLibrariesResponse{}, err
	}
	return result, nil
}

// GetStagingLibraries - PERMISSIONS Read permission for the environment item.
// REQUIRED DELEGATED SCOPES Environment.Read.All or Environment.ReadWrite.All
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
//   - environmentID - The environment ID.
//   - options - SparkLibrariesClientGetStagingLibrariesOptions contains the optional parameters for the SparkLibrariesClient.GetStagingLibraries
//     method.
func (client *SparkLibrariesClient) GetStagingLibraries(ctx context.Context, workspaceID string, environmentID string, options *SparkLibrariesClientGetStagingLibrariesOptions) (SparkLibrariesClientGetStagingLibrariesResponse, error) {
	var err error
	const operationName = "environment.SparkLibrariesClient.GetStagingLibraries"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getStagingLibrariesCreateRequest(ctx, workspaceID, environmentID, options)
	if err != nil {
		return SparkLibrariesClientGetStagingLibrariesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SparkLibrariesClientGetStagingLibrariesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return SparkLibrariesClientGetStagingLibrariesResponse{}, err
	}
	resp, err := client.getStagingLibrariesHandleResponse(httpResp)
	return resp, err
}

// getStagingLibrariesCreateRequest creates the GetStagingLibraries request.
func (client *SparkLibrariesClient) getStagingLibrariesCreateRequest(ctx context.Context, workspaceID string, environmentID string, _ *SparkLibrariesClientGetStagingLibrariesOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/environments/{environmentId}/staging/libraries"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if environmentID == "" {
		return nil, errors.New("parameter environmentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentId}", url.PathEscape(environmentID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getStagingLibrariesHandleResponse handles the GetStagingLibraries response.
func (client *SparkLibrariesClient) getStagingLibrariesHandleResponse(resp *http.Response) (SparkLibrariesClientGetStagingLibrariesResponse, error) {
	result := SparkLibrariesClientGetStagingLibrariesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Libraries); err != nil {
		return SparkLibrariesClientGetStagingLibrariesResponse{}, err
	}
	return result, nil
}

// PublishEnvironment - PERMISSIONS Write permission for the environment item.
// REQUIRED DELEGATED SCOPES Environment.ReadWrite.All
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
//   - environmentID - The environment ID.
//   - options - SparkLibrariesClientPublishEnvironmentOptions contains the optional parameters for the SparkLibrariesClient.PublishEnvironment
//     method.
func (client *SparkLibrariesClient) PublishEnvironment(ctx context.Context, workspaceID string, environmentID string, options *SparkLibrariesClientPublishEnvironmentOptions) (SparkLibrariesClientPublishEnvironmentResponse, error) {
	var err error
	const operationName = "environment.SparkLibrariesClient.PublishEnvironment"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.publishEnvironmentCreateRequest(ctx, workspaceID, environmentID, options)
	if err != nil {
		return SparkLibrariesClientPublishEnvironmentResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SparkLibrariesClientPublishEnvironmentResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = core.NewResponseError(httpResp)
		return SparkLibrariesClientPublishEnvironmentResponse{}, err
	}
	resp, err := client.publishEnvironmentHandleResponse(httpResp)
	return resp, err
}

// publishEnvironmentCreateRequest creates the PublishEnvironment request.
func (client *SparkLibrariesClient) publishEnvironmentCreateRequest(ctx context.Context, workspaceID string, environmentID string, _ *SparkLibrariesClientPublishEnvironmentOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/environments/{environmentId}/staging/publish"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if environmentID == "" {
		return nil, errors.New("parameter environmentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentId}", url.PathEscape(environmentID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// publishEnvironmentHandleResponse handles the PublishEnvironment response.
func (client *SparkLibrariesClient) publishEnvironmentHandleResponse(resp *http.Response) (SparkLibrariesClientPublishEnvironmentResponse, error) {
	result := SparkLibrariesClientPublishEnvironmentResponse{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return SparkLibrariesClientPublishEnvironmentResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PublishInfo); err != nil {
		return SparkLibrariesClientPublishEnvironmentResponse{}, err
	}
	return result, nil
}

// UploadStagingLibrary - This API allows one file upload at a time. The supported file formats are .jar, .py, .whl, .tar.gz,
// or environment.yml.
// LIMITATIONS The maximum request size allowed is 100MB.
// Network issues may impact the upload size and may result in timeouts.
// PERMISSIONS Write permission for the environment item.
// REQUIRED DELEGATED SCOPES Environment.ReadWrite.All
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
//   - environmentID - The environment ID.
//   - options - SparkLibrariesClientUploadStagingLibraryOptions contains the optional parameters for the SparkLibrariesClient.UploadStagingLibrary
//     method.
func (client *SparkLibrariesClient) UploadStagingLibrary(ctx context.Context, workspaceID string, environmentID string, options *SparkLibrariesClientUploadStagingLibraryOptions) (SparkLibrariesClientUploadStagingLibraryResponse, error) {
	var err error
	const operationName = "environment.SparkLibrariesClient.UploadStagingLibrary"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.uploadStagingLibraryCreateRequest(ctx, workspaceID, environmentID, options)
	if err != nil {
		return SparkLibrariesClientUploadStagingLibraryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SparkLibrariesClientUploadStagingLibraryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return SparkLibrariesClientUploadStagingLibraryResponse{}, err
	}
	return SparkLibrariesClientUploadStagingLibraryResponse{}, nil
}

// uploadStagingLibraryCreateRequest creates the UploadStagingLibrary request.
func (client *SparkLibrariesClient) uploadStagingLibraryCreateRequest(ctx context.Context, workspaceID string, environmentID string, _ *SparkLibrariesClientUploadStagingLibraryOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/environments/{environmentId}/staging/libraries"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if environmentID == "" {
		return nil, errors.New("parameter environmentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentId}", url.PathEscape(environmentID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Custom code starts below
