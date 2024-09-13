// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package datapipeline

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

// BeginCreateDataPipeline - This API supports long running operations (LRO) [/rest/api/fabric/articles/long-running-operation].
// PERMISSIONS THE CALLER MUST HAVE CONTRIBUTOR OR HIGHER WORKSPACE ROLE.
// REQUIRED DELEGATED SCOPES DataPipeline.ReadWrite.All or Item.ReadWrite.All
// LIMITATIONS
// * To create a data pipeline, the workspace must be on a supported Fabric capacity.
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
//   - createDataPipelineRequest - Create item request payload.
//   - options - ItemsClientBeginCreateDataPipelineOptions contains the optional parameters for the ItemsClient.BeginCreateDataPipeline
//     method.
func (client *ItemsClient) BeginCreateDataPipeline(ctx context.Context, workspaceID string, createDataPipelineRequest CreateDataPipelineRequest, options *ItemsClientBeginCreateDataPipelineOptions) (*runtime.Poller[ItemsClientCreateDataPipelineResponse], error) {
	return client.beginCreateDataPipeline(ctx, workspaceID, createDataPipelineRequest, options)
}

// CreateDataPipeline - This API supports long running operations (LRO) [/rest/api/fabric/articles/long-running-operation].
// PERMISSIONS THE CALLER MUST HAVE CONTRIBUTOR OR HIGHER WORKSPACE ROLE.
// REQUIRED DELEGATED SCOPES DataPipeline.ReadWrite.All or Item.ReadWrite.All
// LIMITATIONS
// * To create a data pipeline, the workspace must be on a supported Fabric capacity.
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | No | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | No |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
func (client *ItemsClient) createDataPipeline(ctx context.Context, workspaceID string, createDataPipelineRequest CreateDataPipelineRequest, options *ItemsClientBeginCreateDataPipelineOptions) (*http.Response, error) {
	var err error
	const operationName = "datapipeline.ItemsClient.BeginCreateDataPipeline"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createDataPipelineCreateRequest(ctx, workspaceID, createDataPipelineRequest, options)
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

// createDataPipelineCreateRequest creates the CreateDataPipeline request.
func (client *ItemsClient) createDataPipelineCreateRequest(ctx context.Context, workspaceID string, createDataPipelineRequest CreateDataPipelineRequest, options *ItemsClientBeginCreateDataPipelineOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/dataPipelines"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, createDataPipelineRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// DeleteDataPipeline - PERMISSIONS The caller must have contributor or higher workspace role.
// REQUIRED DELEGATED SCOPES DataPipeline.ReadWrite.All or Item.ReadWrite.All
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
//   - dataPipelineID - The data pipeline ID.
//   - options - ItemsClientDeleteDataPipelineOptions contains the optional parameters for the ItemsClient.DeleteDataPipeline
//     method.
func (client *ItemsClient) DeleteDataPipeline(ctx context.Context, workspaceID string, dataPipelineID string, options *ItemsClientDeleteDataPipelineOptions) (ItemsClientDeleteDataPipelineResponse, error) {
	var err error
	const operationName = "datapipeline.ItemsClient.DeleteDataPipeline"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteDataPipelineCreateRequest(ctx, workspaceID, dataPipelineID, options)
	if err != nil {
		return ItemsClientDeleteDataPipelineResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ItemsClientDeleteDataPipelineResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return ItemsClientDeleteDataPipelineResponse{}, err
	}
	return ItemsClientDeleteDataPipelineResponse{}, nil
}

// deleteDataPipelineCreateRequest creates the DeleteDataPipeline request.
func (client *ItemsClient) deleteDataPipelineCreateRequest(ctx context.Context, workspaceID string, dataPipelineID string, options *ItemsClientDeleteDataPipelineOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/dataPipelines/{dataPipelineId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if dataPipelineID == "" {
		return nil, errors.New("parameter dataPipelineID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataPipelineId}", url.PathEscape(dataPipelineID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetDataPipeline - PERMISSIONS The caller must have viewer or higher workspace role.
// REQUIRED DELEGATED SCOPES DataPipeline.Read.All or DataPipeline.ReadWrite.All or Item.Read.All or Item.ReadWrite.All
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
//   - dataPipelineID - The data pipeline ID.
//   - options - ItemsClientGetDataPipelineOptions contains the optional parameters for the ItemsClient.GetDataPipeline method.
func (client *ItemsClient) GetDataPipeline(ctx context.Context, workspaceID string, dataPipelineID string, options *ItemsClientGetDataPipelineOptions) (ItemsClientGetDataPipelineResponse, error) {
	var err error
	const operationName = "datapipeline.ItemsClient.GetDataPipeline"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getDataPipelineCreateRequest(ctx, workspaceID, dataPipelineID, options)
	if err != nil {
		return ItemsClientGetDataPipelineResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ItemsClientGetDataPipelineResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return ItemsClientGetDataPipelineResponse{}, err
	}
	resp, err := client.getDataPipelineHandleResponse(httpResp)
	return resp, err
}

// getDataPipelineCreateRequest creates the GetDataPipeline request.
func (client *ItemsClient) getDataPipelineCreateRequest(ctx context.Context, workspaceID string, dataPipelineID string, options *ItemsClientGetDataPipelineOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/dataPipelines/{dataPipelineId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if dataPipelineID == "" {
		return nil, errors.New("parameter dataPipelineID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataPipelineId}", url.PathEscape(dataPipelineID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDataPipelineHandleResponse handles the GetDataPipeline response.
func (client *ItemsClient) getDataPipelineHandleResponse(resp *http.Response) (ItemsClientGetDataPipelineResponse, error) {
	result := ItemsClientGetDataPipelineResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataPipeline); err != nil {
		return ItemsClientGetDataPipelineResponse{}, err
	}
	return result, nil
}

// NewListDataPipelinesPager - This API supports pagination [/rest/api/fabric/articles/pagination].
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
//   - options - ItemsClientListDataPipelinesOptions contains the optional parameters for the ItemsClient.NewListDataPipelinesPager
//     method.
func (client *ItemsClient) NewListDataPipelinesPager(workspaceID string, options *ItemsClientListDataPipelinesOptions) *runtime.Pager[ItemsClientListDataPipelinesResponse] {
	return runtime.NewPager(runtime.PagingHandler[ItemsClientListDataPipelinesResponse]{
		More: func(page ItemsClientListDataPipelinesResponse) bool {
			return page.ContinuationURI != nil && len(*page.ContinuationURI) > 0
		},
		Fetcher: func(ctx context.Context, page *ItemsClientListDataPipelinesResponse) (ItemsClientListDataPipelinesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "datapipeline.ItemsClient.NewListDataPipelinesPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ContinuationURI
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listDataPipelinesCreateRequest(ctx, workspaceID, options)
			}, nil)
			if err != nil {
				return ItemsClientListDataPipelinesResponse{}, err
			}
			return client.listDataPipelinesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listDataPipelinesCreateRequest creates the ListDataPipelines request.
func (client *ItemsClient) listDataPipelinesCreateRequest(ctx context.Context, workspaceID string, options *ItemsClientListDataPipelinesOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/dataPipelines"
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

// listDataPipelinesHandleResponse handles the ListDataPipelines response.
func (client *ItemsClient) listDataPipelinesHandleResponse(resp *http.Response) (ItemsClientListDataPipelinesResponse, error) {
	result := ItemsClientListDataPipelinesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataPipelines); err != nil {
		return ItemsClientListDataPipelinesResponse{}, err
	}
	return result, nil
}

// UpdateDataPipeline - PERMISSIONS The caller must have contributor or higher workspace role.
// REQUIRED DELEGATED SCOPES DataPipeline.ReadWrite.All or Item.ReadWrite.All
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
//   - dataPipelineID - The data pipeline ID.
//   - updateDataPipelineRequest - Update data pipeline request payload.
//   - options - ItemsClientUpdateDataPipelineOptions contains the optional parameters for the ItemsClient.UpdateDataPipeline
//     method.
func (client *ItemsClient) UpdateDataPipeline(ctx context.Context, workspaceID string, dataPipelineID string, updateDataPipelineRequest UpdateDataPipelineRequest, options *ItemsClientUpdateDataPipelineOptions) (ItemsClientUpdateDataPipelineResponse, error) {
	var err error
	const operationName = "datapipeline.ItemsClient.UpdateDataPipeline"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateDataPipelineCreateRequest(ctx, workspaceID, dataPipelineID, updateDataPipelineRequest, options)
	if err != nil {
		return ItemsClientUpdateDataPipelineResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ItemsClientUpdateDataPipelineResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return ItemsClientUpdateDataPipelineResponse{}, err
	}
	resp, err := client.updateDataPipelineHandleResponse(httpResp)
	return resp, err
}

// updateDataPipelineCreateRequest creates the UpdateDataPipeline request.
func (client *ItemsClient) updateDataPipelineCreateRequest(ctx context.Context, workspaceID string, dataPipelineID string, updateDataPipelineRequest UpdateDataPipelineRequest, options *ItemsClientUpdateDataPipelineOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/dataPipelines/{dataPipelineId}"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if dataPipelineID == "" {
		return nil, errors.New("parameter dataPipelineID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataPipelineId}", url.PathEscape(dataPipelineID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, updateDataPipelineRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// updateDataPipelineHandleResponse handles the UpdateDataPipeline response.
func (client *ItemsClient) updateDataPipelineHandleResponse(resp *http.Response) (ItemsClientUpdateDataPipelineResponse, error) {
	result := ItemsClientUpdateDataPipelineResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataPipeline); err != nil {
		return ItemsClientUpdateDataPipelineResponse{}, err
	}
	return result, nil
}

// Custom code starts below

// CreateDataPipeline - returns ItemsClientCreateDataPipelineResponse in sync mode.
// This API supports long running operations (LRO) [/rest/api/fabric/articles/long-running-operation].
//
// PERMISSIONS THE CALLER MUST HAVE CONTRIBUTOR OR HIGHER WORKSPACE ROLE.
// REQUIRED DELEGATED SCOPES DataPipeline.ReadWrite.All or Item.ReadWrite.All
//
// LIMITATIONS
//
//   - To create a data pipeline, the workspace must be on a supported Fabric capacity.
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] | No | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | No |
//
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - createDataPipelineRequest - Create item request payload.
//   - options - ItemsClientBeginCreateDataPipelineOptions contains the optional parameters for the ItemsClient.BeginCreateDataPipeline method.
func (client *ItemsClient) CreateDataPipeline(ctx context.Context, workspaceID string, createDataPipelineRequest CreateDataPipelineRequest, options *ItemsClientBeginCreateDataPipelineOptions) (ItemsClientCreateDataPipelineResponse, error) {
	return iruntime.NewLRO(client.BeginCreateDataPipeline(ctx, workspaceID, createDataPipelineRequest, options)).Sync(ctx)
}

// beginCreateDataPipeline creates the createDataPipeline request.
func (client *ItemsClient) beginCreateDataPipeline(ctx context.Context, workspaceID string, createDataPipelineRequest CreateDataPipelineRequest, options *ItemsClientBeginCreateDataPipelineOptions) (*runtime.Poller[ItemsClientCreateDataPipelineResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createDataPipeline(ctx, workspaceID, createDataPipelineRequest, options)
		if err != nil {
			var azcoreRespError *azcore.ResponseError
			if errors.As(err, &azcoreRespError) {
				return nil, core.NewResponseError(azcoreRespError.RawResponse)
			}
			return nil, err
		}
		handler, err := locasync.NewPollerHandler[ItemsClientCreateDataPipelineResponse](client.internal.Pipeline(), resp, runtime.FinalStateViaAzureAsyncOp)
		if err != nil {
			var azcoreRespError *azcore.ResponseError
			if errors.As(err, &azcoreRespError) {
				return nil, core.NewResponseError(azcoreRespError.RawResponse)
			}
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ItemsClientCreateDataPipelineResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Handler:       handler,
			Tracer:        client.internal.Tracer(),
		})
	} else {
		handler, err := locasync.NewPollerHandler[ItemsClientCreateDataPipelineResponse](client.internal.Pipeline(), nil, runtime.FinalStateViaAzureAsyncOp)
		if err != nil {
			var azcoreRespError *azcore.ResponseError
			if errors.As(err, &azcoreRespError) {
				return nil, core.NewResponseError(azcoreRespError.RawResponse)
			}
			return nil, err
		}
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ItemsClientCreateDataPipelineResponse]{
			Handler: handler,
			Tracer:  client.internal.Tracer(),
		})
	}
}

// ListDataPipelines - returns array of DataPipeline from all pages.
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
//   - options - ItemsClientListDataPipelinesOptions contains the optional parameters for the ItemsClient.NewListDataPipelinesPager method.
func (client *ItemsClient) ListDataPipelines(ctx context.Context, workspaceID string, options *ItemsClientListDataPipelinesOptions) ([]DataPipeline, error) {
	pager := client.NewListDataPipelinesPager(workspaceID, options)
	mapper := func(resp ItemsClientListDataPipelinesResponse) []DataPipeline {
		return resp.Value
	}
	list, err := iruntime.NewPageIterator(ctx, pager, mapper).Get()
	if err != nil {
		return nil, err
	}
	return list, nil
}