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

// TagsClient contains the methods for the Tags group.
// Don't use this type directly, use a constructor function instead.
type TagsClient struct {
	internal *azcore.Client
	endpoint string
}

// BulkCreateTags - PERMISSIONS The caller must have administrator rights such as Fabric administrator.
// REQUIRED DELEGATED SCOPES Tenant.ReadWrite.All.
// LIMITATIONS Maximum 25 requests per one minute per principal.
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - createTagsRequest - The request payload for creating tags.
//   - options - TagsClientBulkCreateTagsOptions contains the optional parameters for the TagsClient.BulkCreateTags method.
func (client *TagsClient) BulkCreateTags(ctx context.Context, createTagsRequest CreateTagsRequest, options *TagsClientBulkCreateTagsOptions) (TagsClientBulkCreateTagsResponse, error) {
	var err error
	const operationName = "admin.TagsClient.BulkCreateTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.bulkCreateTagsCreateRequest(ctx, createTagsRequest, options)
	if err != nil {
		return TagsClientBulkCreateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TagsClientBulkCreateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = core.NewResponseError(httpResp)
		return TagsClientBulkCreateTagsResponse{}, err
	}
	resp, err := client.bulkCreateTagsHandleResponse(httpResp)
	return resp, err
}

// bulkCreateTagsCreateRequest creates the BulkCreateTags request.
func (client *TagsClient) bulkCreateTagsCreateRequest(ctx context.Context, createTagsRequest CreateTagsRequest, _ *TagsClientBulkCreateTagsOptions) (*policy.Request, error) {
	urlPath := "/v1/admin/tags/bulkCreateTags"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, createTagsRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// bulkCreateTagsHandleResponse handles the BulkCreateTags response.
func (client *TagsClient) bulkCreateTagsHandleResponse(resp *http.Response) (TagsClientBulkCreateTagsResponse, error) {
	result := TagsClientBulkCreateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CreateTagsResponse); err != nil {
		return TagsClientBulkCreateTagsResponse{}, err
	}
	return result, nil
}

// DeleteTag - PERMISSIONS The caller must have administrator rights such as Fabric administrator.
// REQUIRED DELEGATED SCOPES Tenant.ReadWrite.All.
// LIMITATIONS Maximum 10 requests per one minute per principal.
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - tagID - The tag ID.
//   - options - TagsClientDeleteTagOptions contains the optional parameters for the TagsClient.DeleteTag method.
func (client *TagsClient) DeleteTag(ctx context.Context, tagID string, options *TagsClientDeleteTagOptions) (TagsClientDeleteTagResponse, error) {
	var err error
	const operationName = "admin.TagsClient.DeleteTag"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteTagCreateRequest(ctx, tagID, options)
	if err != nil {
		return TagsClientDeleteTagResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TagsClientDeleteTagResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return TagsClientDeleteTagResponse{}, err
	}
	return TagsClientDeleteTagResponse{}, nil
}

// deleteTagCreateRequest creates the DeleteTag request.
func (client *TagsClient) deleteTagCreateRequest(ctx context.Context, tagID string, _ *TagsClientDeleteTagOptions) (*policy.Request, error) {
	urlPath := "/v1/admin/tags/{tagId}"
	if tagID == "" {
		return nil, errors.New("parameter tagID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagId}", url.PathEscape(tagID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// NewListTagsPager - REQUIRED DELEGATED SCOPES Tenant.Read.All or Tenant.ReadWrite.All
// LIMITATIONS Maximum 25 requests per one minute per principal.
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
//
// Generated from API version v1
//   - options - TagsClientListTagsOptions contains the optional parameters for the TagsClient.NewListTagsPager method.
func (client *TagsClient) NewListTagsPager(options *TagsClientListTagsOptions) *runtime.Pager[TagsClientListTagsResponse] {
	return runtime.NewPager(runtime.PagingHandler[TagsClientListTagsResponse]{
		More: func(page TagsClientListTagsResponse) bool {
			return page.ContinuationURI != nil && len(*page.ContinuationURI) > 0
		},
		Fetcher: func(ctx context.Context, page *TagsClientListTagsResponse) (TagsClientListTagsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "admin.TagsClient.NewListTagsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ContinuationURI
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listTagsCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return TagsClientListTagsResponse{}, err
			}
			return client.listTagsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listTagsCreateRequest creates the ListTags request.
func (client *TagsClient) listTagsCreateRequest(ctx context.Context, options *TagsClientListTagsOptions) (*policy.Request, error) {
	urlPath := "/v1/admin/tags"
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

// listTagsHandleResponse handles the ListTags response.
func (client *TagsClient) listTagsHandleResponse(resp *http.Response) (TagsClientListTagsResponse, error) {
	result := TagsClientListTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagsInfo); err != nil {
		return TagsClientListTagsResponse{}, err
	}
	return result, nil
}

// UpdateTag - PERMISSIONS The caller must have administrator rights such as Fabric administrator.
// REQUIRED DELEGATED SCOPES Tenant.ReadWrite.All.
// LIMITATIONS Maximum 25 requests per one minute per principal.
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - tagID - The tag ID.
//   - updateTagRequest - The request payload for updating the tag.
//   - options - TagsClientUpdateTagOptions contains the optional parameters for the TagsClient.UpdateTag method.
func (client *TagsClient) UpdateTag(ctx context.Context, tagID string, updateTagRequest UpdateTagRequest, options *TagsClientUpdateTagOptions) (TagsClientUpdateTagResponse, error) {
	var err error
	const operationName = "admin.TagsClient.UpdateTag"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateTagCreateRequest(ctx, tagID, updateTagRequest, options)
	if err != nil {
		return TagsClientUpdateTagResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TagsClientUpdateTagResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return TagsClientUpdateTagResponse{}, err
	}
	resp, err := client.updateTagHandleResponse(httpResp)
	return resp, err
}

// updateTagCreateRequest creates the UpdateTag request.
func (client *TagsClient) updateTagCreateRequest(ctx context.Context, tagID string, updateTagRequest UpdateTagRequest, _ *TagsClientUpdateTagOptions) (*policy.Request, error) {
	urlPath := "/v1/admin/tags/{tagId}"
	if tagID == "" {
		return nil, errors.New("parameter tagID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagId}", url.PathEscape(tagID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, updateTagRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// updateTagHandleResponse handles the UpdateTag response.
func (client *TagsClient) updateTagHandleResponse(resp *http.Response) (TagsClientUpdateTagResponse, error) {
	result := TagsClientUpdateTagResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Tag); err != nil {
		return TagsClientUpdateTagResponse{}, err
	}
	return result, nil
}

// Custom code starts below

// ListTags - returns array of TagInfo from all pages.
// REQUIRED DELEGATED SCOPES Tenant.Read.All or Tenant.ReadWrite.All
//
// LIMITATIONS Maximum 25 requests per one minute per principal.
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
//
// INTERFACE
// Generated from API version v1
//   - options - TagsClientListTagsOptions contains the optional parameters for the TagsClient.NewListTagsPager method.
func (client *TagsClient) ListTags(ctx context.Context, options *TagsClientListTagsOptions) ([]TagInfo, error) {
	pager := client.NewListTagsPager(options)
	mapper := func(resp TagsClientListTagsResponse) []TagInfo {
		return resp.Value
	}
	list, err := iruntime.NewPageIterator(ctx, pager, mapper).Get()
	if err != nil {
		var azcoreRespError *azcore.ResponseError
		if errors.As(err, &azcoreRespError) {
			return []TagInfo{}, core.NewResponseError(azcoreRespError.RawResponse)
		}
		return []TagInfo{}, err
	}
	return list, nil
}
