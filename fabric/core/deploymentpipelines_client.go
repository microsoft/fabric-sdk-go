// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package core

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
	"github.com/microsoft/fabric-sdk-go/internal/pollers/locasync"
)

// DeploymentPipelinesClient contains the methods for the DeploymentPipelines group.
// Don't use this type directly, use a constructor function instead.
type DeploymentPipelinesClient struct {
	internal *azcore.Client
	endpoint string
}

// BeginDeployStageContent - To learn about items that are supported in deployment pipelines, see: Supported items [/fabric/cicd/deployment-pipelines/intro-to-deployment-pipelines#supported-items].
// > [!IMPORTANT] This API supports long running operations (LRO) [/rest/api/fabric/articles/long-running-operation].
// PERMISSIONS The user must be at least a contributor on both source and target deployment workspaces. For more information,
// see: Permissions [https://go.microsoft.com/fwlink/?linkid=2235654].
// REQUIRED DELEGATED SCOPES Pipeline.Deploy
// LIMITATIONS Maximum 300 deployed items per request.
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | Only Power BI Items
// [/rest/api/fabric/articles/item-management/item-management-overview#power-bi] are supported | | Managed identities [/entra/identity/managed-identities-azure-resources/overview]
// | Only Power BI Items
// [/rest/api/fabric/articles/item-management/item-management-overview#power-bi] are supported |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - deploymentPipelineID - The deployment pipeline ID.
//   - deployRequest - The deploy request.
//   - options - DeploymentPipelinesClientBeginDeployStageContentOptions contains the optional parameters for the DeploymentPipelinesClient.BeginDeployStageContent
//     method.
func (client *DeploymentPipelinesClient) BeginDeployStageContent(ctx context.Context, deploymentPipelineID string, deployRequest DeployRequest, options *DeploymentPipelinesClientBeginDeployStageContentOptions) (*runtime.Poller[DeploymentPipelinesClientDeployStageContentResponse], error) {
	return client.beginDeployStageContent(ctx, deploymentPipelineID, deployRequest, options)
}

// DeployStageContent - To learn about items that are supported in deployment pipelines, see: Supported items [/fabric/cicd/deployment-pipelines/intro-to-deployment-pipelines#supported-items].
// > [!IMPORTANT] This API supports long running operations (LRO) [/rest/api/fabric/articles/long-running-operation].
// PERMISSIONS The user must be at least a contributor on both source and target deployment workspaces. For more information,
// see: Permissions [https://go.microsoft.com/fwlink/?linkid=2235654].
// REQUIRED DELEGATED SCOPES Pipeline.Deploy
// LIMITATIONS Maximum 300 deployed items per request.
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | Only Power BI Items
// [/rest/api/fabric/articles/item-management/item-management-overview#power-bi] are supported | | Managed identities [/entra/identity/managed-identities-azure-resources/overview]
// | Only Power BI Items
// [/rest/api/fabric/articles/item-management/item-management-overview#power-bi] are supported |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
func (client *DeploymentPipelinesClient) deployStageContent(ctx context.Context, deploymentPipelineID string, deployRequest DeployRequest, options *DeploymentPipelinesClientBeginDeployStageContentOptions) (*http.Response, error) {
	var err error
	const operationName = "core.DeploymentPipelinesClient.BeginDeployStageContent"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deployStageContentCreateRequest(ctx, deploymentPipelineID, deployRequest, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deployStageContentCreateRequest creates the DeployStageContent request.
func (client *DeploymentPipelinesClient) deployStageContentCreateRequest(ctx context.Context, deploymentPipelineID string, deployRequest DeployRequest, options *DeploymentPipelinesClientBeginDeployStageContentOptions) (*policy.Request, error) {
	urlPath := "/v1/deploymentPipelines/{deploymentPipelineId}/deploy"
	if deploymentPipelineID == "" {
		return nil, errors.New("parameter deploymentPipelineID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentPipelineId}", url.PathEscape(deploymentPipelineID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, deployRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// GetDeploymentPipeline - REQUIRED DELEGATED SCOPES Pipeline.Read.All or Pipeline.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | Yes | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - deploymentPipelineID - The deployment pipeline ID.
//   - options - DeploymentPipelinesClientGetDeploymentPipelineOptions contains the optional parameters for the DeploymentPipelinesClient.GetDeploymentPipeline
//     method.
func (client *DeploymentPipelinesClient) GetDeploymentPipeline(ctx context.Context, deploymentPipelineID string, options *DeploymentPipelinesClientGetDeploymentPipelineOptions) (DeploymentPipelinesClientGetDeploymentPipelineResponse, error) {
	var err error
	const operationName = "core.DeploymentPipelinesClient.GetDeploymentPipeline"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getDeploymentPipelineCreateRequest(ctx, deploymentPipelineID, options)
	if err != nil {
		return DeploymentPipelinesClientGetDeploymentPipelineResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeploymentPipelinesClientGetDeploymentPipelineResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = NewResponseError(httpResp)
		return DeploymentPipelinesClientGetDeploymentPipelineResponse{}, err
	}
	resp, err := client.getDeploymentPipelineHandleResponse(httpResp)
	return resp, err
}

// getDeploymentPipelineCreateRequest creates the GetDeploymentPipeline request.
func (client *DeploymentPipelinesClient) getDeploymentPipelineCreateRequest(ctx context.Context, deploymentPipelineID string, options *DeploymentPipelinesClientGetDeploymentPipelineOptions) (*policy.Request, error) {
	urlPath := "/v1/deploymentPipelines/{deploymentPipelineId}"
	if deploymentPipelineID == "" {
		return nil, errors.New("parameter deploymentPipelineID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentPipelineId}", url.PathEscape(deploymentPipelineID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDeploymentPipelineHandleResponse handles the GetDeploymentPipeline response.
func (client *DeploymentPipelinesClient) getDeploymentPipelineHandleResponse(resp *http.Response) (DeploymentPipelinesClientGetDeploymentPipelineResponse, error) {
	result := DeploymentPipelinesClientGetDeploymentPipelineResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentPipeline); err != nil {
		return DeploymentPipelinesClientGetDeploymentPipelineResponse{}, err
	}
	return result, nil
}

// NewListDeploymentPipelineStageItemsPager - To learn about items that are supported in deployment pipelines, see: Supported
// items [/fabric/cicd/deployment-pipelines/intro-to-deployment-pipelines#supported-items].
// PERMISSIONS The user must be at least a workspace contributor assigned to the specified stage. For more information, see:
// Permissions [https://go.microsoft.com/fwlink/?linkid=2235654].
// REQUIRED DELEGATED SCOPES Pipeline.Read.All or Pipeline.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | Yes | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
//
// Generated from API version v1
//   - deploymentPipelineID - The deployment pipeline ID.
//   - stageID - The deployment pipeline stage ID.
//   - options - DeploymentPipelinesClientListDeploymentPipelineStageItemsOptions contains the optional parameters for the DeploymentPipelinesClient.NewListDeploymentPipelineStageItemsPager
//     method.
func (client *DeploymentPipelinesClient) NewListDeploymentPipelineStageItemsPager(deploymentPipelineID string, stageID string, options *DeploymentPipelinesClientListDeploymentPipelineStageItemsOptions) *runtime.Pager[DeploymentPipelinesClientListDeploymentPipelineStageItemsResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeploymentPipelinesClientListDeploymentPipelineStageItemsResponse]{
		More: func(page DeploymentPipelinesClientListDeploymentPipelineStageItemsResponse) bool {
			return page.ContinuationURI != nil && len(*page.ContinuationURI) > 0
		},
		Fetcher: func(ctx context.Context, page *DeploymentPipelinesClientListDeploymentPipelineStageItemsResponse) (DeploymentPipelinesClientListDeploymentPipelineStageItemsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "core.DeploymentPipelinesClient.NewListDeploymentPipelineStageItemsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ContinuationURI
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listDeploymentPipelineStageItemsCreateRequest(ctx, deploymentPipelineID, stageID, options)
			}, nil)
			if err != nil {
				return DeploymentPipelinesClientListDeploymentPipelineStageItemsResponse{}, err
			}
			return client.listDeploymentPipelineStageItemsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listDeploymentPipelineStageItemsCreateRequest creates the ListDeploymentPipelineStageItems request.
func (client *DeploymentPipelinesClient) listDeploymentPipelineStageItemsCreateRequest(ctx context.Context, deploymentPipelineID string, stageID string, options *DeploymentPipelinesClientListDeploymentPipelineStageItemsOptions) (*policy.Request, error) {
	urlPath := "/v1/deploymentPipelines/{deploymentPipelineId}/stages/{stageId}/items"
	if deploymentPipelineID == "" {
		return nil, errors.New("parameter deploymentPipelineID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentPipelineId}", url.PathEscape(deploymentPipelineID))
	if stageID == "" {
		return nil, errors.New("parameter stageID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{stageId}", url.PathEscape(stageID))
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

// listDeploymentPipelineStageItemsHandleResponse handles the ListDeploymentPipelineStageItems response.
func (client *DeploymentPipelinesClient) listDeploymentPipelineStageItemsHandleResponse(resp *http.Response) (DeploymentPipelinesClientListDeploymentPipelineStageItemsResponse, error) {
	result := DeploymentPipelinesClientListDeploymentPipelineStageItemsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentPipelineStageItems); err != nil {
		return DeploymentPipelinesClientListDeploymentPipelineStageItemsResponse{}, err
	}
	return result, nil
}

// NewListDeploymentPipelineStagesPager - REQUIRED DELEGATED SCOPES Pipeline.Read.All or Pipeline.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | Yes | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
//
// Generated from API version v1
//   - deploymentPipelineID - The deployment pipeline ID.
//   - options - DeploymentPipelinesClientListDeploymentPipelineStagesOptions contains the optional parameters for the DeploymentPipelinesClient.NewListDeploymentPipelineStagesPager
//     method.
func (client *DeploymentPipelinesClient) NewListDeploymentPipelineStagesPager(deploymentPipelineID string, options *DeploymentPipelinesClientListDeploymentPipelineStagesOptions) *runtime.Pager[DeploymentPipelinesClientListDeploymentPipelineStagesResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeploymentPipelinesClientListDeploymentPipelineStagesResponse]{
		More: func(page DeploymentPipelinesClientListDeploymentPipelineStagesResponse) bool {
			return page.ContinuationURI != nil && len(*page.ContinuationURI) > 0
		},
		Fetcher: func(ctx context.Context, page *DeploymentPipelinesClientListDeploymentPipelineStagesResponse) (DeploymentPipelinesClientListDeploymentPipelineStagesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "core.DeploymentPipelinesClient.NewListDeploymentPipelineStagesPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ContinuationURI
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listDeploymentPipelineStagesCreateRequest(ctx, deploymentPipelineID, options)
			}, nil)
			if err != nil {
				return DeploymentPipelinesClientListDeploymentPipelineStagesResponse{}, err
			}
			return client.listDeploymentPipelineStagesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listDeploymentPipelineStagesCreateRequest creates the ListDeploymentPipelineStages request.
func (client *DeploymentPipelinesClient) listDeploymentPipelineStagesCreateRequest(ctx context.Context, deploymentPipelineID string, options *DeploymentPipelinesClientListDeploymentPipelineStagesOptions) (*policy.Request, error) {
	urlPath := "/v1/deploymentPipelines/{deploymentPipelineId}/stages"
	if deploymentPipelineID == "" {
		return nil, errors.New("parameter deploymentPipelineID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentPipelineId}", url.PathEscape(deploymentPipelineID))
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

// listDeploymentPipelineStagesHandleResponse handles the ListDeploymentPipelineStages response.
func (client *DeploymentPipelinesClient) listDeploymentPipelineStagesHandleResponse(resp *http.Response) (DeploymentPipelinesClientListDeploymentPipelineStagesResponse, error) {
	result := DeploymentPipelinesClientListDeploymentPipelineStagesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentPipelineStages); err != nil {
		return DeploymentPipelinesClientListDeploymentPipelineStagesResponse{}, err
	}
	return result, nil
}

// NewListDeploymentPipelinesPager - REQUIRED DELEGATED SCOPES Pipeline.Read.All or Pipeline.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// | Yes | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
//
// Generated from API version v1
//   - options - DeploymentPipelinesClientListDeploymentPipelinesOptions contains the optional parameters for the DeploymentPipelinesClient.NewListDeploymentPipelinesPager
//     method.
func (client *DeploymentPipelinesClient) NewListDeploymentPipelinesPager(options *DeploymentPipelinesClientListDeploymentPipelinesOptions) *runtime.Pager[DeploymentPipelinesClientListDeploymentPipelinesResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeploymentPipelinesClientListDeploymentPipelinesResponse]{
		More: func(page DeploymentPipelinesClientListDeploymentPipelinesResponse) bool {
			return page.ContinuationURI != nil && len(*page.ContinuationURI) > 0
		},
		Fetcher: func(ctx context.Context, page *DeploymentPipelinesClientListDeploymentPipelinesResponse) (DeploymentPipelinesClientListDeploymentPipelinesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "core.DeploymentPipelinesClient.NewListDeploymentPipelinesPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ContinuationURI
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listDeploymentPipelinesCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return DeploymentPipelinesClientListDeploymentPipelinesResponse{}, err
			}
			return client.listDeploymentPipelinesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listDeploymentPipelinesCreateRequest creates the ListDeploymentPipelines request.
func (client *DeploymentPipelinesClient) listDeploymentPipelinesCreateRequest(ctx context.Context, options *DeploymentPipelinesClientListDeploymentPipelinesOptions) (*policy.Request, error) {
	urlPath := "/v1/deploymentPipelines"
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

// listDeploymentPipelinesHandleResponse handles the ListDeploymentPipelines response.
func (client *DeploymentPipelinesClient) listDeploymentPipelinesHandleResponse(resp *http.Response) (DeploymentPipelinesClientListDeploymentPipelinesResponse, error) {
	result := DeploymentPipelinesClientListDeploymentPipelinesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentPipelines); err != nil {
		return DeploymentPipelinesClientListDeploymentPipelinesResponse{}, err
	}
	return result, nil
}

// Custom code starts below

// DeployStageContent - returns DeploymentPipelinesClientDeployStageContentResponse in sync mode.
// To learn about items that are supported in deployment pipelines, see: Supported items [/fabric/cicd/deployment-pipelines/intro-to-deployment-pipelines#supported-items].
//
// >  [!IMPORTANT] This API supports long running operations (LRO) [/rest/api/fabric/articles/long-running-operation].
//
// PERMISSIONS The user must be at least a contributor on both source and target deployment workspaces. For more information, see: Permissions [https://go.microsoft.com/fwlink/?linkid=2235654].
//
// # REQUIRED DELEGATED SCOPES Pipeline.Deploy
//
// LIMITATIONS Maximum 300 deployed items per request.
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] | Only Power BI Items
// [/rest/api/fabric/articles/item-management/item-management-overview#power-bi] are supported | | Managed identities [/entra/identity/managed-identities-azure-resources/overview] | Only Power BI Items
// [/rest/api/fabric/articles/item-management/item-management-overview#power-bi] are supported |
//
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
// Generated from API version v1
//   - deploymentPipelineID - The deployment pipeline ID.
//   - deployRequest - The deploy request.
//   - options - DeploymentPipelinesClientBeginDeployStageContentOptions contains the optional parameters for the DeploymentPipelinesClient.BeginDeployStageContent method.
func (client *DeploymentPipelinesClient) DeployStageContent(ctx context.Context, deploymentPipelineID string, deployRequest DeployRequest, options *DeploymentPipelinesClientBeginDeployStageContentOptions) (DeploymentPipelinesClientDeployStageContentResponse, error) {
	return iruntime.NewLRO(client.BeginDeployStageContent(ctx, deploymentPipelineID, deployRequest, options)).Sync(ctx)
}

// beginDeployStageContent creates the deployStageContent request.
func (client *DeploymentPipelinesClient) beginDeployStageContent(ctx context.Context, deploymentPipelineID string, deployRequest DeployRequest, options *DeploymentPipelinesClientBeginDeployStageContentOptions) (*runtime.Poller[DeploymentPipelinesClientDeployStageContentResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deployStageContent(ctx, deploymentPipelineID, deployRequest, options)
		if err != nil {
			var azcoreRespError *azcore.ResponseError
			if errors.As(err, &azcoreRespError) {
				return nil, NewResponseError(azcoreRespError.RawResponse)
			}
			return nil, err
		}
		handler, err := locasync.NewPollerHandler[DeploymentPipelinesClientDeployStageContentResponse](client.internal.Pipeline(), resp, runtime.FinalStateViaAzureAsyncOp)
		if err != nil {
			var azcoreRespError *azcore.ResponseError
			if errors.As(err, &azcoreRespError) {
				return nil, NewResponseError(azcoreRespError.RawResponse)
			}
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DeploymentPipelinesClientDeployStageContentResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Handler:       handler,
			Tracer:        client.internal.Tracer(),
		})
	} else {
		handler, err := locasync.NewPollerHandler[DeploymentPipelinesClientDeployStageContentResponse](client.internal.Pipeline(), nil, runtime.FinalStateViaAzureAsyncOp)
		if err != nil {
			var azcoreRespError *azcore.ResponseError
			if errors.As(err, &azcoreRespError) {
				return nil, NewResponseError(azcoreRespError.RawResponse)
			}
			return nil, err
		}
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DeploymentPipelinesClientDeployStageContentResponse]{
			Handler: handler,
			Tracer:  client.internal.Tracer(),
		})
	}
}

// ListDeploymentPipelineStageItems - returns array of DeploymentPipelineStageItem from all pages.
// To learn about items that are supported in deployment pipelines, see: Supported items [/fabric/cicd/deployment-pipelines/intro-to-deployment-pipelines#supported-items].
//
// PERMISSIONS The user must be at least a workspace contributor assigned to the specified stage. For more information, see: Permissions [https://go.microsoft.com/fwlink/?linkid=2235654].
//
// # REQUIRED DELEGATED SCOPES Pipeline.Read.All or Pipeline.ReadWrite.All
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] | Yes | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
//
// INTERFACE
// Generated from API version v1
//   - deploymentPipelineID - The deployment pipeline ID.
//   - stageID - The deployment pipeline stage ID.
//   - options - DeploymentPipelinesClientListDeploymentPipelineStageItemsOptions contains the optional parameters for the DeploymentPipelinesClient.NewListDeploymentPipelineStageItemsPager method.
func (client *DeploymentPipelinesClient) ListDeploymentPipelineStageItems(ctx context.Context, deploymentPipelineID string, stageID string, options *DeploymentPipelinesClientListDeploymentPipelineStageItemsOptions) ([]DeploymentPipelineStageItem, error) {
	pager := client.NewListDeploymentPipelineStageItemsPager(deploymentPipelineID, stageID, options)
	mapper := func(resp DeploymentPipelinesClientListDeploymentPipelineStageItemsResponse) []DeploymentPipelineStageItem {
		return resp.Value
	}
	list, err := iruntime.NewPageIterator(ctx, pager, mapper).Get()
	if err != nil {
		return nil, err
	}
	return list, nil
}

// ListDeploymentPipelineStages - returns array of DeploymentPipelineStage from all pages.
// REQUIRED DELEGATED SCOPES Pipeline.Read.All or Pipeline.ReadWrite.All
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] | Yes | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
//
// INTERFACE
// Generated from API version v1
//   - deploymentPipelineID - The deployment pipeline ID.
//   - options - DeploymentPipelinesClientListDeploymentPipelineStagesOptions contains the optional parameters for the DeploymentPipelinesClient.NewListDeploymentPipelineStagesPager method.
func (client *DeploymentPipelinesClient) ListDeploymentPipelineStages(ctx context.Context, deploymentPipelineID string, options *DeploymentPipelinesClientListDeploymentPipelineStagesOptions) ([]DeploymentPipelineStage, error) {
	pager := client.NewListDeploymentPipelineStagesPager(deploymentPipelineID, options)
	mapper := func(resp DeploymentPipelinesClientListDeploymentPipelineStagesResponse) []DeploymentPipelineStage {
		return resp.Value
	}
	list, err := iruntime.NewPageIterator(ctx, pager, mapper).Get()
	if err != nil {
		return nil, err
	}
	return list, nil
}

// ListDeploymentPipelines - returns array of DeploymentPipeline from all pages.
// REQUIRED DELEGATED SCOPES Pipeline.Read.All or Pipeline.ReadWrite.All
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] | Yes | | Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
//
// INTERFACE
// Generated from API version v1
//   - options - DeploymentPipelinesClientListDeploymentPipelinesOptions contains the optional parameters for the DeploymentPipelinesClient.NewListDeploymentPipelinesPager method.
func (client *DeploymentPipelinesClient) ListDeploymentPipelines(ctx context.Context, options *DeploymentPipelinesClientListDeploymentPipelinesOptions) ([]DeploymentPipeline, error) {
	pager := client.NewListDeploymentPipelinesPager(options)
	mapper := func(resp DeploymentPipelinesClientListDeploymentPipelinesResponse) []DeploymentPipeline {
		return resp.Value
	}
	list, err := iruntime.NewPageIterator(ctx, pager, mapper).Get()
	if err != nil {
		return nil, err
	}
	return list, nil
}
