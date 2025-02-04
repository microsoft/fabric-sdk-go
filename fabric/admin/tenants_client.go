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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"

	"github.com/microsoft/fabric-sdk-go/fabric/core"
	"github.com/microsoft/fabric-sdk-go/internal/iruntime"
)

// TenantsClient contains the methods for the Tenants group.
// Don't use this type directly, use a constructor function instead.
type TenantsClient struct {
	internal *azcore.Client
	endpoint string
}

// NewListCapacitiesTenantSettingsOverridesPager - This API supports pagination [/rest/api/fabric/articles/pagination]. A
// maximum of 10,000 records can be returned per request. With the continuation token provided in the response, you can get
// the next
// 10,000 records.
// The user must be a Fabric Service Administrator or authenticate using a service principal.
// REQUIRED DELEGATED SCOPES Tenant.Read.All or Tenant.ReadWrite.All
// LIMITATIONS Maximum 100 requests per hour.
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
//
// Generated from API version v1
//   - options - TenantsClientListCapacitiesTenantSettingsOverridesOptions contains the optional parameters for the TenantsClient.NewListCapacitiesTenantSettingsOverridesPager
//     method.
func (client *TenantsClient) NewListCapacitiesTenantSettingsOverridesPager(options *TenantsClientListCapacitiesTenantSettingsOverridesOptions) *runtime.Pager[TenantsClientListCapacitiesTenantSettingsOverridesResponse] {
	return runtime.NewPager(runtime.PagingHandler[TenantsClientListCapacitiesTenantSettingsOverridesResponse]{
		More: func(page TenantsClientListCapacitiesTenantSettingsOverridesResponse) bool {
			return page.ContinuationURI != nil && len(*page.ContinuationURI) > 0
		},
		Fetcher: func(ctx context.Context, page *TenantsClientListCapacitiesTenantSettingsOverridesResponse) (TenantsClientListCapacitiesTenantSettingsOverridesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "admin.TenantsClient.NewListCapacitiesTenantSettingsOverridesPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ContinuationURI
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCapacitiesTenantSettingsOverridesCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return TenantsClientListCapacitiesTenantSettingsOverridesResponse{}, err
			}
			return client.listCapacitiesTenantSettingsOverridesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCapacitiesTenantSettingsOverridesCreateRequest creates the ListCapacitiesTenantSettingsOverrides request.
func (client *TenantsClient) listCapacitiesTenantSettingsOverridesCreateRequest(ctx context.Context, options *TenantsClientListCapacitiesTenantSettingsOverridesOptions) (*policy.Request, error) {
	urlPath := "/v1/admin/capacities/delegatedTenantSettingOverrides"
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

// listCapacitiesTenantSettingsOverridesHandleResponse handles the ListCapacitiesTenantSettingsOverrides response.
func (client *TenantsClient) listCapacitiesTenantSettingsOverridesHandleResponse(resp *http.Response) (TenantsClientListCapacitiesTenantSettingsOverridesResponse, error) {
	result := TenantsClientListCapacitiesTenantSettingsOverridesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TenantSettingOverrides); err != nil {
		return TenantsClientListCapacitiesTenantSettingsOverridesResponse{}, err
	}
	return result, nil
}

// ListTenantSettings - PERMISSIONS The caller must be a Fabric administrator or authenticate using a service principal.
// REQUIRED DELEGATED SCOPES Tenant.Read.All or Tenant.ReadWrite.All
// LIMITATIONS Maximum 200 requests per hour.
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
//   - options - TenantsClientListTenantSettingsOptions contains the optional parameters for the TenantsClient.ListTenantSettings
//     method.
func (client *TenantsClient) ListTenantSettings(ctx context.Context, options *TenantsClientListTenantSettingsOptions) (TenantsClientListTenantSettingsResponse, error) {
	var err error
	const operationName = "admin.TenantsClient.ListTenantSettings"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listTenantSettingsCreateRequest(ctx, options)
	if err != nil {
		return TenantsClientListTenantSettingsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TenantsClientListTenantSettingsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = core.NewResponseError(httpResp)
		return TenantsClientListTenantSettingsResponse{}, err
	}
	resp, err := client.listTenantSettingsHandleResponse(httpResp)
	return resp, err
}

// listTenantSettingsCreateRequest creates the ListTenantSettings request.
func (client *TenantsClient) listTenantSettingsCreateRequest(ctx context.Context, _ *TenantsClientListTenantSettingsOptions) (*policy.Request, error) {
	urlPath := "/v1/admin/tenantsettings"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listTenantSettingsHandleResponse handles the ListTenantSettings response.
func (client *TenantsClient) listTenantSettingsHandleResponse(resp *http.Response) (TenantsClientListTenantSettingsResponse, error) {
	result := TenantsClientListTenantSettingsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TenantSettings); err != nil {
		return TenantsClientListTenantSettingsResponse{}, err
	}
	return result, nil
}

// Custom code starts below

// ListCapacitiesTenantSettingsOverrides - returns array of TenantSettingOverride from all pages.
// This API supports pagination [/rest/api/fabric/articles/pagination]. A maximum of 10,000 records can be returned per request. With the continuation token provided in the response, you can get the next
// 10,000 records.
//
// The user must be a Fabric Service Administrator or authenticate using a service principal.
//
// # REQUIRED DELEGATED SCOPES Tenant.Read.All or Tenant.ReadWrite.All
//
// LIMITATIONS Maximum 100 requests per hour.
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
//
// INTERFACE
// Generated from API version v1
//   - options - TenantsClientListCapacitiesTenantSettingsOverridesOptions contains the optional parameters for the TenantsClient.NewListCapacitiesTenantSettingsOverridesPager method.
func (client *TenantsClient) ListCapacitiesTenantSettingsOverrides(ctx context.Context, options *TenantsClientListCapacitiesTenantSettingsOverridesOptions) ([]TenantSettingOverride, error) {
	pager := client.NewListCapacitiesTenantSettingsOverridesPager(options)
	mapper := func(resp TenantsClientListCapacitiesTenantSettingsOverridesResponse) []TenantSettingOverride {
		return resp.Overrides
	}
	list, err := iruntime.NewPageIterator(ctx, pager, mapper).Get()
	if err != nil {
		var azcoreRespError *azcore.ResponseError
		if errors.As(err, &azcoreRespError) {
			return []TenantSettingOverride{}, core.NewResponseError(azcoreRespError.RawResponse)
		}
		return []TenantSettingOverride{}, err
	}
	return list, nil
}
