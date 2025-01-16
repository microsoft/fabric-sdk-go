// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package lakehouse

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
	"github.com/microsoft/fabric-sdk-go/internal/iruntime"
	"github.com/microsoft/fabric-sdk-go/internal/pollers/locasync"
)

// TablesClient contains the methods for the Tables group.
// Don't use this type directly, use a constructor function instead.
type TablesClient struct {
	internal *azcore.Client
	endpoint string
}

// NewListTablesPager - This API supports pagination [/rest/api/fabric/articles/pagination]. A maximum of 100 records can
// be returned per request. With the URI provided in the response, you can get the next page of records.
// REQUIRED DELEGATED SCOPES Lakehouse.Read.All or Lakehouse.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
//
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - lakehouseID - The lakehouse ID.
//   - options - TablesClientListTablesOptions contains the optional parameters for the TablesClient.NewListTablesPager method.
func (client *TablesClient) NewListTablesPager(workspaceID string, lakehouseID string, options *TablesClientListTablesOptions) *runtime.Pager[TablesClientListTablesResponse] {
	return runtime.NewPager(runtime.PagingHandler[TablesClientListTablesResponse]{
		More: func(page TablesClientListTablesResponse) bool {
			return page.ContinuationURI != nil && len(*page.ContinuationURI) > 0
		},
		Fetcher: func(ctx context.Context, page *TablesClientListTablesResponse) (TablesClientListTablesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "lakehouse.TablesClient.NewListTablesPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ContinuationURI
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listTablesCreateRequest(ctx, workspaceID, lakehouseID, options)
			}, nil)
			if err != nil {
				return TablesClientListTablesResponse{}, err
			}
			return client.listTablesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listTablesCreateRequest creates the ListTables request.
func (client *TablesClient) listTablesCreateRequest(ctx context.Context, workspaceID string, lakehouseID string, options *TablesClientListTablesOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/lakehouses/{lakehouseId}/tables"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if lakehouseID == "" {
		return nil, errors.New("parameter lakehouseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{lakehouseId}", url.PathEscape(lakehouseID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	if options != nil && options.MaxResults != nil {
		reqQP.Set("maxResults", strconv.FormatInt(int64(*options.MaxResults), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listTablesHandleResponse handles the ListTables response.
func (client *TablesClient) listTablesHandleResponse(resp *http.Response) (TablesClientListTablesResponse, error) {
	result := TablesClientListTablesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Tables); err != nil {
		return TablesClientListTablesResponse{}, err
	}
	return result, nil
}

// BeginLoadTable - This API supports long running operations (LRO) [/rest/api/fabric/articles/long-running-operation].
// PERMISSIONS Write permission to the lakehouse item.
// REQUIRED DELEGATED SCOPES Lakehouse.ReadWrite.All
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
//   - lakehouseID - The lakehouse item ID.
//   - tableName - The table name.
//   - loadTableRequest - The load table request payload.
//   - options - TablesClientBeginLoadTableOptions contains the optional parameters for the TablesClient.BeginLoadTable method.
func (client *TablesClient) BeginLoadTable(ctx context.Context, workspaceID string, lakehouseID string, tableName string, loadTableRequest LoadTableRequest, options *TablesClientBeginLoadTableOptions) (*runtime.Poller[TablesClientLoadTableResponse], error) {
	return client.beginLoadTable(ctx, workspaceID, lakehouseID, tableName, loadTableRequest, options)
}

// LoadTable - This API supports long running operations (LRO) [/rest/api/fabric/articles/long-running-operation].
// PERMISSIONS Write permission to the lakehouse item.
// REQUIRED DELEGATED SCOPES Lakehouse.ReadWrite.All
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support]
// listed in this section.
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object]
// and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
// INTERFACE
// If the operation fails it returns an *core.ResponseError type.
//
// Generated from API version v1
func (client *TablesClient) loadTable(ctx context.Context, workspaceID string, lakehouseID string, tableName string, loadTableRequest LoadTableRequest, options *TablesClientBeginLoadTableOptions) (*http.Response, error) {
	var err error
	const operationName = "lakehouse.TablesClient.BeginLoadTable"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.loadTableCreateRequest(ctx, workspaceID, lakehouseID, tableName, loadTableRequest, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = core.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// loadTableCreateRequest creates the LoadTable request.
func (client *TablesClient) loadTableCreateRequest(ctx context.Context, workspaceID string, lakehouseID string, tableName string, loadTableRequest LoadTableRequest, _ *TablesClientBeginLoadTableOptions) (*policy.Request, error) {
	urlPath := "/v1/workspaces/{workspaceId}/lakehouses/{lakehouseId}/tables/{tableName}/load"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if lakehouseID == "" {
		return nil, errors.New("parameter lakehouseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{lakehouseId}", url.PathEscape(lakehouseID))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, loadTableRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// Custom code starts below

// LoadTable - returns TablesClientLoadTableResponse in sync mode.
// This API supports long running operations (LRO) [/rest/api/fabric/articles/long-running-operation].
//
// PERMISSIONS Write permission to the lakehouse item.
//
// # REQUIRED DELEGATED SCOPES Lakehouse.ReadWrite.All
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
//
// INTERFACE
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - lakehouseID - The lakehouse item ID.
//   - tableName - The table name.
//   - loadTableRequest - The load table request payload.
//   - options - TablesClientBeginLoadTableOptions contains the optional parameters for the TablesClient.BeginLoadTable method.
func (client *TablesClient) LoadTable(ctx context.Context, workspaceID string, lakehouseID string, tableName string, loadTableRequest LoadTableRequest, options *TablesClientBeginLoadTableOptions) (TablesClientLoadTableResponse, error) {
	result, err := iruntime.NewLRO(client.BeginLoadTable(ctx, workspaceID, lakehouseID, tableName, loadTableRequest, options)).Sync(ctx)
	if err != nil {
		var azcoreRespError *azcore.ResponseError
		if errors.As(err, &azcoreRespError) {
			return TablesClientLoadTableResponse{}, core.NewResponseError(azcoreRespError.RawResponse)
		}
		return TablesClientLoadTableResponse{}, err
	}
	return result, err
}

// beginLoadTable creates the loadTable request.
func (client *TablesClient) beginLoadTable(ctx context.Context, workspaceID string, lakehouseID string, tableName string, loadTableRequest LoadTableRequest, options *TablesClientBeginLoadTableOptions) (*runtime.Poller[TablesClientLoadTableResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.loadTable(ctx, workspaceID, lakehouseID, tableName, loadTableRequest, options)
		if err != nil {
			var azcoreRespError *azcore.ResponseError
			if errors.As(err, &azcoreRespError) {
				return nil, core.NewResponseError(azcoreRespError.RawResponse)
			}
			return nil, err
		}
		handler, err := locasync.NewPollerHandler[TablesClientLoadTableResponse](client.internal.Pipeline(), resp, runtime.FinalStateViaAzureAsyncOp)
		if err != nil {
			var azcoreRespError *azcore.ResponseError
			if errors.As(err, &azcoreRespError) {
				return nil, core.NewResponseError(azcoreRespError.RawResponse)
			}
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TablesClientLoadTableResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Handler:       handler,
			Tracer:        client.internal.Tracer(),
		})
	} else {
		handler, err := locasync.NewPollerHandler[TablesClientLoadTableResponse](client.internal.Pipeline(), nil, runtime.FinalStateViaAzureAsyncOp)
		if err != nil {
			var azcoreRespError *azcore.ResponseError
			if errors.As(err, &azcoreRespError) {
				return nil, core.NewResponseError(azcoreRespError.RawResponse)
			}
			return nil, err
		}
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TablesClientLoadTableResponse]{
			Handler: handler,
			Tracer:  client.internal.Tracer(),
		})
	}
}

// ListTables - returns array of Table from all pages.
// This API supports pagination [/rest/api/fabric/articles/pagination]. A maximum of 100 records can be returned per request. With the URI provided in the response, you can get the next page of records.
//
// # REQUIRED DELEGATED SCOPES Lakehouse.Read.All or Lakehouse.ReadWrite.All
//
// MICROSOFT ENTRA SUPPORTED IDENTITIES This API supports the Microsoft identities [/rest/api/fabric/articles/identity-support] listed in this section.
//
// | Identity | Support | |-|-| | User | Yes | | Service principal [/entra/identity-platform/app-objects-and-service-principals#service-principal-object] and Managed identities
// [/entra/identity/managed-identities-azure-resources/overview] | Yes |
//
// INTERFACE
// Generated from API version v1
//   - workspaceID - The workspace ID.
//   - lakehouseID - The lakehouse ID.
//   - options - TablesClientListTablesOptions contains the optional parameters for the TablesClient.NewListTablesPager method.
func (client *TablesClient) ListTables(ctx context.Context, workspaceID string, lakehouseID string, options *TablesClientListTablesOptions) ([]Table, error) {
	pager := client.NewListTablesPager(workspaceID, lakehouseID, options)
	mapper := func(resp TablesClientListTablesResponse) []Table {
		return resp.Data
	}
	list, err := iruntime.NewPageIterator(ctx, pager, mapper).Get()
	if err != nil {
		var azcoreRespError *azcore.ResponseError
		if errors.As(err, &azcoreRespError) {
			return []Table{}, core.NewResponseError(azcoreRespError.RawResponse)
		}
		return []Table{}, err
	}
	return list, nil
}
