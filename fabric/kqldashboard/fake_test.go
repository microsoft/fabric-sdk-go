// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package kqldashboard_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"

	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"

	"reflect"

	"github.com/stretchr/testify/suite"

	"github.com/microsoft/fabric-sdk-go/fabric/kqldashboard"
	"github.com/microsoft/fabric-sdk-go/fabric/kqldashboard/fake"
)

var err error

type FakeTestSuite struct {
	suite.Suite

	ctx  context.Context
	cred azcore.TokenCredential

	serverFactory *fake.ServerFactory
	clientFactory *kqldashboard.ClientFactory
}

func (testsuite *FakeTestSuite) SetupSuite() {
	testsuite.ctx = context.Background()
	testsuite.cred = &azfake.TokenCredential{}

	testsuite.serverFactory = &fake.ServerFactory{}
	testsuite.clientFactory, err = kqldashboard.NewClientFactory(testsuite.cred, nil, &azcore.ClientOptions{
		Transport: fake.NewServerFactoryTransport(testsuite.serverFactory),
	})
	testsuite.Require().NoError(err, "Failed to create client factory")
}

func TestFakeTest(t *testing.T) {
	suite.Run(t, new(FakeTestSuite))
}

func (testsuite *FakeTestSuite) TestItems_ListKQLDashboards() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"List KQL dashboards in workspace example"},
	})
	var exampleWorkspaceID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"

	exampleRes := kqldashboard.KQLDashboards{
		Value: []kqldashboard.KQLDashboard{
			{
				Type:        to.Ptr(kqldashboard.ItemTypeKQLDashboard),
				Description: to.Ptr("A KQL dashboard description."),
				DisplayName: to.Ptr("KQLDashboard_1"),
				ID:          to.Ptr("3546052c-ae64-4526-b1a8-52af7761426f"),
				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
			},
			{
				Type:        to.Ptr(kqldashboard.ItemTypeKQLDashboard),
				Description: to.Ptr("A KQL dashboard description."),
				DisplayName: to.Ptr("KQLDashboard_2"),
				ID:          to.Ptr("340d91b9-5a39-409c-b9c0-05ba832c476e"),
				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
			}},
	}

	testsuite.serverFactory.ItemsServer.NewListKQLDashboardsPager = func(workspaceID string, options *kqldashboard.ItemsClientListKQLDashboardsOptions) (resp azfake.PagerResponder[kqldashboard.ItemsClientListKQLDashboardsResponse]) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		resp = azfake.PagerResponder[kqldashboard.ItemsClientListKQLDashboardsResponse]{}
		resp.AddPage(http.StatusOK, kqldashboard.ItemsClientListKQLDashboardsResponse{KQLDashboards: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	pager := client.NewListKQLDashboardsPager(exampleWorkspaceID, &kqldashboard.ItemsClientListKQLDashboardsOptions{ContinuationToken: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		testsuite.Require().NoError(err, "Failed to advance page for example ")
		testsuite.Require().True(reflect.DeepEqual(exampleRes, nextResult.KQLDashboards))
		if err == nil {
			break
		}
	}
}

func (testsuite *FakeTestSuite) TestItems_CreateKQLDashboard() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Create a KQL dashboard example"},
	})
	var exampleWorkspaceID string
	var exampleCreateKQLDashboardRequest kqldashboard.CreateKQLDashboardRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleCreateKQLDashboardRequest = kqldashboard.CreateKQLDashboardRequest{
		Description: to.Ptr("A KQL dashboard description"),
		DisplayName: to.Ptr("KQLDashboard_1"),
	}

	testsuite.serverFactory.ItemsServer.BeginCreateKQLDashboard = func(ctx context.Context, workspaceID string, createKQLDashboardRequest kqldashboard.CreateKQLDashboardRequest, options *kqldashboard.ItemsClientBeginCreateKQLDashboardOptions) (resp azfake.PollerResponder[kqldashboard.ItemsClientCreateKQLDashboardResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().True(reflect.DeepEqual(exampleCreateKQLDashboardRequest, createKQLDashboardRequest))
		resp = azfake.PollerResponder[kqldashboard.ItemsClientCreateKQLDashboardResponse]{}
		resp.SetTerminalResponse(http.StatusCreated, kqldashboard.ItemsClientCreateKQLDashboardResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	poller, err := client.BeginCreateKQLDashboard(ctx, exampleWorkspaceID, exampleCreateKQLDashboardRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")

	// From example
	ctx = runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Create a KQL dashboard example with definition"},
	})
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleCreateKQLDashboardRequest = kqldashboard.CreateKQLDashboardRequest{
		Description: to.Ptr("A KQL dashboard description"),
		Definition: &kqldashboard.Definition{
			Parts: []kqldashboard.DefinitionPart{
				{
					Path:        to.Ptr("RealTimeDashboard.json"),
					Payload:     to.Ptr("ewogICJhdXRvUmVmcmVzaCI6IHsKICAgICJlbmFibGVkIjogZmFsc2UKICB9LAogICJiYXNlUXVlcmllcyI6IFtdLAogICJ0aWxlcyI6IFtdLAogICJkYXRhU291cmNlcyI6IFtdLAogICJwYWdlcyI6IFsKICAgIHsKICAgICAgIm5hbWUiOiAiUGFnZSAxIiwKICAgICAgImlkIjogImFiNWYwNjI4LTk3NTAtNDNlNi05YmE0LTIwMTE2ZTg2ODc1NSIKICAgIH0KICBdLAogICJwYXJhbWV0ZXJzIjogWwogICAgewogICAgICAia2luZCI6ICJkdXJhdGlvbiIsCiAgICAgICJpZCI6ICJmNTEwYjlmZi0xZjI3LTQ4YmQtODBlNC1mZDFhNTllODQ5MjQiLAogICAgICAiZGlzcGxheU5hbWUiOiAiVGltZSByYW5nZSIsCiAgICAgICJkZXNjcmlwdGlvbiI6ICIiLAogICAgICAiYmVnaW5WYXJpYWJsZU5hbWUiOiAiX3N0YXJ0VGltZSIsCiAgICAgICJlbmRWYXJpYWJsZU5hbWUiOiAiX2VuZFRpbWUiLAogICAgICAiZGVmYXVsdFZhbHVlIjogewogICAgICAgICJraW5kIjogImR5bmFtaWMiLAogICAgICAgICJjb3VudCI6IDEsCiAgICAgICAgInVuaXQiOiAiaG91cnMiCiAgICAgIH0sCiAgICAgICJzaG93T25QYWdlcyI6IHsKICAgICAgICAia2luZCI6ICJhbGwiCiAgICAgIH0KICAgIH0KICBdLAogICJxdWVyaWVzIjogW10sCiAgInNjaGVtYV92ZXJzaW9uIjogIjUyIiwKICAidGl0bGUiOiAidmNkem5ncnpzbSIKfQ=="),
					PayloadType: to.Ptr(kqldashboard.PayloadTypeInlineBase64),
				},
				{
					Path:        to.Ptr(".platform"),
					Payload:     to.Ptr("ZG90UGxhdGZvcm1CYXNlNjRTdHJpbmc="),
					PayloadType: to.Ptr(kqldashboard.PayloadTypeInlineBase64),
				}},
		},
		DisplayName: to.Ptr("KQL Dashboard 1"),
	}

	testsuite.serverFactory.ItemsServer.BeginCreateKQLDashboard = func(ctx context.Context, workspaceID string, createKQLDashboardRequest kqldashboard.CreateKQLDashboardRequest, options *kqldashboard.ItemsClientBeginCreateKQLDashboardOptions) (resp azfake.PollerResponder[kqldashboard.ItemsClientCreateKQLDashboardResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().True(reflect.DeepEqual(exampleCreateKQLDashboardRequest, createKQLDashboardRequest))
		resp = azfake.PollerResponder[kqldashboard.ItemsClientCreateKQLDashboardResponse]{}
		resp.SetTerminalResponse(http.StatusCreated, kqldashboard.ItemsClientCreateKQLDashboardResponse{}, nil)
		return
	}

	poller, err = client.BeginCreateKQLDashboard(ctx, exampleWorkspaceID, exampleCreateKQLDashboardRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")
}

func (testsuite *FakeTestSuite) TestItems_GetKQLDashboard() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Get a KQL dashboard example"},
	})
	var exampleWorkspaceID string
	var exampleKqlDashboardID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleKqlDashboardID = "5b218778-e7a5-4d73-8187-f10824047715"

	exampleRes := kqldashboard.KQLDashboard{
		Type:        to.Ptr(kqldashboard.ItemTypeKQLDashboard),
		Description: to.Ptr("A KQL dashboard description."),
		DisplayName: to.Ptr("KQLDashboard_1"),
		ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	}

	testsuite.serverFactory.ItemsServer.GetKQLDashboard = func(ctx context.Context, workspaceID string, kqlDashboardID string, options *kqldashboard.ItemsClientGetKQLDashboardOptions) (resp azfake.Responder[kqldashboard.ItemsClientGetKQLDashboardResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleKqlDashboardID, kqlDashboardID)
		resp = azfake.Responder[kqldashboard.ItemsClientGetKQLDashboardResponse]{}
		resp.SetResponse(http.StatusOK, kqldashboard.ItemsClientGetKQLDashboardResponse{KQLDashboard: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	res, err := client.GetKQLDashboard(ctx, exampleWorkspaceID, exampleKqlDashboardID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.KQLDashboard))
}

func (testsuite *FakeTestSuite) TestItems_UpdateKQLDashboard() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Update a KQL dashboard example"},
	})
	var exampleWorkspaceID string
	var exampleKqlDashboardID string
	var exampleUpdateKQLDashboardRequest kqldashboard.UpdateKQLDashboardRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleKqlDashboardID = "5b218778-e7a5-4d73-8187-f10824047715"
	exampleUpdateKQLDashboardRequest = kqldashboard.UpdateKQLDashboardRequest{
		Description: to.Ptr("KQL dashboard new description"),
		DisplayName: to.Ptr("KQL dashboard new name"),
	}

	exampleRes := kqldashboard.KQLDashboard{
		Type:        to.Ptr(kqldashboard.ItemTypeKQLDashboard),
		Description: to.Ptr("KQL dashboard new description"),
		DisplayName: to.Ptr("KQL dashboard new name"),
		ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	}

	testsuite.serverFactory.ItemsServer.UpdateKQLDashboard = func(ctx context.Context, workspaceID string, kqlDashboardID string, updateKQLDashboardRequest kqldashboard.UpdateKQLDashboardRequest, options *kqldashboard.ItemsClientUpdateKQLDashboardOptions) (resp azfake.Responder[kqldashboard.ItemsClientUpdateKQLDashboardResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleKqlDashboardID, kqlDashboardID)
		testsuite.Require().True(reflect.DeepEqual(exampleUpdateKQLDashboardRequest, updateKQLDashboardRequest))
		resp = azfake.Responder[kqldashboard.ItemsClientUpdateKQLDashboardResponse]{}
		resp.SetResponse(http.StatusOK, kqldashboard.ItemsClientUpdateKQLDashboardResponse{KQLDashboard: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	res, err := client.UpdateKQLDashboard(ctx, exampleWorkspaceID, exampleKqlDashboardID, exampleUpdateKQLDashboardRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.KQLDashboard))
}

func (testsuite *FakeTestSuite) TestItems_DeleteKQLDashboard() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Delete a KQL dashboard example"},
	})
	var exampleWorkspaceID string
	var exampleKqlDashboardID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleKqlDashboardID = "5b218778-e7a5-4d73-8187-f10824047715"

	testsuite.serverFactory.ItemsServer.DeleteKQLDashboard = func(ctx context.Context, workspaceID string, kqlDashboardID string, options *kqldashboard.ItemsClientDeleteKQLDashboardOptions) (resp azfake.Responder[kqldashboard.ItemsClientDeleteKQLDashboardResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleKqlDashboardID, kqlDashboardID)
		resp = azfake.Responder[kqldashboard.ItemsClientDeleteKQLDashboardResponse]{}
		resp.SetResponse(http.StatusOK, kqldashboard.ItemsClientDeleteKQLDashboardResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	_, err = client.DeleteKQLDashboard(ctx, exampleWorkspaceID, exampleKqlDashboardID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
}

func (testsuite *FakeTestSuite) TestItems_GetKQLDashboardDefinition() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Get a KQL dashboard definition example"},
	})
	var exampleWorkspaceID string
	var exampleKqlDashboardID string
	exampleWorkspaceID = "6e335e92-a2a2-4b5a-970a-bd6a89fbb765"
	exampleKqlDashboardID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"

	exampleRes := kqldashboard.DefinitionResponse{
		Definition: &kqldashboard.Definition{
			Parts: []kqldashboard.DefinitionPart{
				{
					Path:        to.Ptr("RealTimeDashboard.json"),
					Payload:     to.Ptr("ewogICJhdXRvUmVmcmVzaCI6IHsKICAgICJlbmFibGVkIjogZmFsc2UKICB9LAogICJiYXNlUXVlcmllcyI6IFtdLAogICJ0aWxlcyI6IFtdLAogICJkYXRhU291cmNlcyI6IFtdLAogICJwYWdlcyI6IFsKICAgIHsKICAgICAgIm5hbWUiOiAiUGFnZSAxIiwKICAgICAgImlkIjogImFiNWYwNjI4LTk3NTAtNDNlNi05YmE0LTIwMTE2ZTg2ODc1NSIKICAgIH0KICBdLAogICJwYXJhbWV0ZXJzIjogWwogICAgewogICAgICAia2luZCI6ICJkdXJhdGlvbiIsCiAgICAgICJpZCI6ICJmNTEwYjlmZi0xZjI3LTQ4YmQtODBlNC1mZDFhNTllODQ5MjQiLAogICAgICAiZGlzcGxheU5hbWUiOiAiVGltZSByYW5nZSIsCiAgICAgICJkZXNjcmlwdGlvbiI6ICIiLAogICAgICAiYmVnaW5WYXJpYWJsZU5hbWUiOiAiX3N0YXJ0VGltZSIsCiAgICAgICJlbmRWYXJpYWJsZU5hbWUiOiAiX2VuZFRpbWUiLAogICAgICAiZGVmYXVsdFZhbHVlIjogewogICAgICAgICJraW5kIjogImR5bmFtaWMiLAogICAgICAgICJjb3VudCI6IDEsCiAgICAgICAgInVuaXQiOiAiaG91cnMiCiAgICAgIH0sCiAgICAgICJzaG93T25QYWdlcyI6IHsKICAgICAgICAia2luZCI6ICJhbGwiCiAgICAgIH0KICAgIH0KICBdLAogICJxdWVyaWVzIjogW10sCiAgInNjaGVtYV92ZXJzaW9uIjogIjUyIiwKICAidGl0bGUiOiAidmNkem5ncnpzbSIKfQ=="),
					PayloadType: to.Ptr(kqldashboard.PayloadTypeInlineBase64),
				},
				{
					Path:        to.Ptr(".platform"),
					Payload:     to.Ptr("ewogICIkc2NoZW1hIjogImh0dHBzOi8vZGV2ZWxvcGVyLm1pY3Jvc29mdC5jb20vanNvbi1zY2hlbWFzL2ZhYnJpYy9naXRJbnRlZ3JhdGlvbi9wbGF0Zm9ybVByb3BlcnRpZXMvMi4wLjAvc2NoZW1hLmpzb24iLAogICJtZXRhZGF0YSI6IHsKICAgICJ0eXBlIjogIktRTERhc2hib2FyZCIsCiAgICAiZGlzcGxheU5hbWUiOiAidmNkem5ncnpzbSIsCiAgICAiZGVzY3JpcHRpb24iOiAidmNkem5ncnpzbSIKICB9LAogICJjb25maWciOiB7CiAgICAidmVyc2lvbiI6ICIyLjAiLAogICAgImxvZ2ljYWxJZCI6ICIwMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwMDAiCiAgfQp9"),
					PayloadType: to.Ptr(kqldashboard.PayloadTypeInlineBase64),
				}},
		},
	}

	testsuite.serverFactory.ItemsServer.BeginGetKQLDashboardDefinition = func(ctx context.Context, workspaceID string, kqlDashboardID string, options *kqldashboard.ItemsClientBeginGetKQLDashboardDefinitionOptions) (resp azfake.PollerResponder[kqldashboard.ItemsClientGetKQLDashboardDefinitionResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleKqlDashboardID, kqlDashboardID)
		resp = azfake.PollerResponder[kqldashboard.ItemsClientGetKQLDashboardDefinitionResponse]{}
		resp.SetTerminalResponse(http.StatusOK, kqldashboard.ItemsClientGetKQLDashboardDefinitionResponse{DefinitionResponse: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	poller, err := client.BeginGetKQLDashboardDefinition(ctx, exampleWorkspaceID, exampleKqlDashboardID, &kqldashboard.ItemsClientBeginGetKQLDashboardDefinitionOptions{Format: nil})
	testsuite.Require().NoError(err, "Failed to get result for example ")
	res, err := poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.DefinitionResponse))
}

func (testsuite *FakeTestSuite) TestItems_UpdateKQLDashboardDefinition() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Update a KQL dashboard definition example"},
	})
	var exampleWorkspaceID string
	var exampleKqlDashboardID string
	var exampleUpdateKQLDashboardDefinitionRequest kqldashboard.UpdateKQLDashboardDefinitionRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleKqlDashboardID = "5b218778-e7a5-4d73-8187-f10824047715"
	exampleUpdateKQLDashboardDefinitionRequest = kqldashboard.UpdateKQLDashboardDefinitionRequest{
		Definition: &kqldashboard.Definition{
			Parts: []kqldashboard.DefinitionPart{
				{
					Path:        to.Ptr("RealTimeDashboard.json"),
					Payload:     to.Ptr("ewogICJhdXRvUmVmcmVzaCI6IHsKICAgICJlbmFibGVkIjogZmFsc2UKICB9LAogICJiYXNlUXVlcmllcyI6IFtdLAogICJ0aWxlcyI6IFtdLAogICJkYXRhU291cmNlcyI6IFtdLAogICJwYWdlcyI6IFsKICAgIHsKICAgICAgIm5hbWUiOiAiUGFnZSAxIiwKICAgICAgImlkIjogImFiNWYwNjI4LTk3NTAtNDNlNi05YmE0LTIwMTE2ZTg2ODc1NSIKICAgIH0KICBdLAogICJwYXJhbWV0ZXJzIjogWwogICAgewogICAgICAia2luZCI6ICJkdXJhdGlvbiIsCiAgICAgICJpZCI6ICJmNTEwYjlmZi0xZjI3LTQ4YmQtODBlNC1mZDFhNTllODQ5MjQiLAogICAgICAiZGlzcGxheU5hbWUiOiAiVGltZSByYW5nZSIsCiAgICAgICJkZXNjcmlwdGlvbiI6ICIiLAogICAgICAiYmVnaW5WYXJpYWJsZU5hbWUiOiAiX3N0YXJ0VGltZSIsCiAgICAgICJlbmRWYXJpYWJsZU5hbWUiOiAiX2VuZFRpbWUiLAogICAgICAiZGVmYXVsdFZhbHVlIjogewogICAgICAgICJraW5kIjogImR5bmFtaWMiLAogICAgICAgICJjb3VudCI6IDEsCiAgICAgICAgInVuaXQiOiAiaG91cnMiCiAgICAgIH0sCiAgICAgICJzaG93T25QYWdlcyI6IHsKICAgICAgICAia2luZCI6ICJhbGwiCiAgICAgIH0KICAgIH0KICBdLAogICJxdWVyaWVzIjogW10sCiAgInNjaGVtYV92ZXJzaW9uIjogIjUyIiwKICAidGl0bGUiOiAidmNkem5ncnpzbSIKfQ=="),
					PayloadType: to.Ptr(kqldashboard.PayloadTypeInlineBase64),
				},
				{
					Path:        to.Ptr(".platform"),
					Payload:     to.Ptr("ZG90UGxhdGZvcm1CYXNlNjRTdHJpbmc="),
					PayloadType: to.Ptr(kqldashboard.PayloadTypeInlineBase64),
				}},
		},
	}

	testsuite.serverFactory.ItemsServer.BeginUpdateKQLDashboardDefinition = func(ctx context.Context, workspaceID string, kqlDashboardID string, updateKQLDashboardDefinitionRequest kqldashboard.UpdateKQLDashboardDefinitionRequest, options *kqldashboard.ItemsClientBeginUpdateKQLDashboardDefinitionOptions) (resp azfake.PollerResponder[kqldashboard.ItemsClientUpdateKQLDashboardDefinitionResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleKqlDashboardID, kqlDashboardID)
		testsuite.Require().True(reflect.DeepEqual(exampleUpdateKQLDashboardDefinitionRequest, updateKQLDashboardDefinitionRequest))
		resp = azfake.PollerResponder[kqldashboard.ItemsClientUpdateKQLDashboardDefinitionResponse]{}
		resp.SetTerminalResponse(http.StatusOK, kqldashboard.ItemsClientUpdateKQLDashboardDefinitionResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	poller, err := client.BeginUpdateKQLDashboardDefinition(ctx, exampleWorkspaceID, exampleKqlDashboardID, exampleUpdateKQLDashboardDefinitionRequest, &kqldashboard.ItemsClientBeginUpdateKQLDashboardDefinitionOptions{UpdateMetadata: to.Ptr(true)})
	testsuite.Require().NoError(err, "Failed to get result for example ")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")
}
