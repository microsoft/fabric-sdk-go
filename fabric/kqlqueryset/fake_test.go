// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package kqlqueryset_test

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

	"github.com/microsoft/fabric-sdk-go/fabric/kqlqueryset"
	"github.com/microsoft/fabric-sdk-go/fabric/kqlqueryset/fake"
)

var err error

type FakeTestSuite struct {
	suite.Suite

	ctx  context.Context
	cred azcore.TokenCredential

	serverFactory *fake.ServerFactory
	clientFactory *kqlqueryset.ClientFactory
}

func (testsuite *FakeTestSuite) SetupSuite() {
	testsuite.ctx = context.Background()
	testsuite.cred = &azfake.TokenCredential{}

	testsuite.serverFactory = &fake.ServerFactory{}
	testsuite.clientFactory, err = kqlqueryset.NewClientFactory(testsuite.cred, nil, &azcore.ClientOptions{
		Transport: fake.NewServerFactoryTransport(testsuite.serverFactory),
	})
	testsuite.Require().NoError(err, "Failed to create client factory")
}

func TestFakeTest(t *testing.T) {
	suite.Run(t, new(FakeTestSuite))
}

func (testsuite *FakeTestSuite) TestItems_ListKQLQuerysets() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"List KQL querysets in workspace example"},
	})
	var exampleWorkspaceID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"

	exampleRes := kqlqueryset.KQLQuerysets{
		Value: []kqlqueryset.KQLQueryset{
			{
				Type:        to.Ptr(kqlqueryset.ItemTypeKQLQueryset),
				Description: to.Ptr("A KQL queryset description."),
				DisplayName: to.Ptr("KQLQueryset_1"),
				ID:          to.Ptr("3546052c-ae64-4526-b1a8-52af7761426f"),
				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
			},
			{
				Type:        to.Ptr(kqlqueryset.ItemTypeKQLQueryset),
				Description: to.Ptr("A KQL queryset description."),
				DisplayName: to.Ptr("KQLQueryset_2"),
				ID:          to.Ptr("4c9adc22-ffb1-491f-baaa-9c9987745591"),
				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
			},
			{
				Type:        to.Ptr(kqlqueryset.ItemTypeKQLQueryset),
				Description: to.Ptr("A KQL queryset description."),
				DisplayName: to.Ptr("KQLQueryset_3"),
				ID:          to.Ptr("8b681594-894d-4adf-8ae8-aed415dd1de6"),
				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
			}},
	}

	testsuite.serverFactory.ItemsServer.NewListKQLQuerysetsPager = func(workspaceID string, options *kqlqueryset.ItemsClientListKQLQuerysetsOptions) (resp azfake.PagerResponder[kqlqueryset.ItemsClientListKQLQuerysetsResponse]) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		resp = azfake.PagerResponder[kqlqueryset.ItemsClientListKQLQuerysetsResponse]{}
		resp.AddPage(http.StatusOK, kqlqueryset.ItemsClientListKQLQuerysetsResponse{KQLQuerysets: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	pager := client.NewListKQLQuerysetsPager(exampleWorkspaceID, &kqlqueryset.ItemsClientListKQLQuerysetsOptions{ContinuationToken: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		testsuite.Require().NoError(err, "Failed to advance page for example ")
		testsuite.Require().True(reflect.DeepEqual(exampleRes, nextResult.KQLQuerysets))
		if err == nil {
			break
		}
	}
}

func (testsuite *FakeTestSuite) TestItems_CreateKQLQueryset() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Create a KQL queryset example"},
	})
	var exampleWorkspaceID string
	var exampleCreateKQLQuerysetRequest kqlqueryset.CreateKQLQuerysetRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleCreateKQLQuerysetRequest = kqlqueryset.CreateKQLQuerysetRequest{
		Description: to.Ptr("A KQL queryset description"),
		DisplayName: to.Ptr("KQLQueryset_1"),
	}

	testsuite.serverFactory.ItemsServer.BeginCreateKQLQueryset = func(ctx context.Context, workspaceID string, createKQLQuerysetRequest kqlqueryset.CreateKQLQuerysetRequest, options *kqlqueryset.ItemsClientBeginCreateKQLQuerysetOptions) (resp azfake.PollerResponder[kqlqueryset.ItemsClientCreateKQLQuerysetResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().True(reflect.DeepEqual(exampleCreateKQLQuerysetRequest, createKQLQuerysetRequest))
		resp = azfake.PollerResponder[kqlqueryset.ItemsClientCreateKQLQuerysetResponse]{}
		resp.SetTerminalResponse(http.StatusCreated, kqlqueryset.ItemsClientCreateKQLQuerysetResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	poller, err := client.BeginCreateKQLQueryset(ctx, exampleWorkspaceID, exampleCreateKQLQuerysetRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")

	// From example
	ctx = runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Create a KQL queryset with definition example"},
	})
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleCreateKQLQuerysetRequest = kqlqueryset.CreateKQLQuerysetRequest{
		Description: to.Ptr("A KQL queryset description"),
		Definition: &kqlqueryset.Definition{
			Parts: []kqlqueryset.DefinitionPart{
				{
					Path:        to.Ptr("RealTimeQueryset.json"),
					Payload:     to.Ptr("ewogICAgInF1ZXJ5c2V0IjogewogICAgICAgICJ2ZXJzaW9uIjogIjEuMC4wIiwKICAgICAgICAiZGF0YVNvdXJjZXMiOiBbewogICAgICAgICAgICAgICAgImlkIjogImMyNDM0YmY4LTI1YmItNGFhMC04NzQ2LWRiNDcwNTMzYWRhZiIsCiAgICAgICAgICAgICAgICAiY2x1c3RlclVyaSI6ICJodHRwczovL2hlbHAua3VzdG8ud2luZG93cy5uZXQvIiwKICAgICAgICAgICAgICAgICJ0eXBlIjogIkF6dXJlRGF0YUV4cGxvcmVyIiwKICAgICAgICAgICAgICAgICJkYXRhYmFzZU5hbWUiOiAiU2FtcGxlcyIKICAgICAgICAgICAgfQogICAgICAgIF0sCiAgICAgICAgInRhYnMiOiBbewogICAgICAgICAgICAgICAgImlkIjogImNjZDdiOTBjLTUxZmUtNDI5Zi1hODUzLTM4NWIwMmJkNzRjOSIsCiAgICAgICAgICAgICAgICAiY29udGVudCI6ICJTdG9ybUV2ZW50c1xcXFxufCBjb3VudCIsCiAgICAgICAgICAgICAgICAidGl0bGUiOiAiVGFiMU5hbWUiLAogICAgICAgICAgICAgICAgImRhdGFTb3VyY2VJZCI6ICJjMjQzNGJmOC0yNWJiLTRhYTAtODc0Ni1kYjQ3MDUzM2FkYWYiCiAgICAgICAgICAgIH0KICAgICAgICBdCiAgICB9Cn0="),
					PayloadType: to.Ptr(kqlqueryset.PayloadTypeInlineBase64),
				}},
		},
		DisplayName: to.Ptr("KQLQueryset_1"),
	}

	testsuite.serverFactory.ItemsServer.BeginCreateKQLQueryset = func(ctx context.Context, workspaceID string, createKQLQuerysetRequest kqlqueryset.CreateKQLQuerysetRequest, options *kqlqueryset.ItemsClientBeginCreateKQLQuerysetOptions) (resp azfake.PollerResponder[kqlqueryset.ItemsClientCreateKQLQuerysetResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().True(reflect.DeepEqual(exampleCreateKQLQuerysetRequest, createKQLQuerysetRequest))
		resp = azfake.PollerResponder[kqlqueryset.ItemsClientCreateKQLQuerysetResponse]{}
		resp.SetTerminalResponse(http.StatusCreated, kqlqueryset.ItemsClientCreateKQLQuerysetResponse{}, nil)
		return
	}

	poller, err = client.BeginCreateKQLQueryset(ctx, exampleWorkspaceID, exampleCreateKQLQuerysetRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")
}

func (testsuite *FakeTestSuite) TestItems_GetKQLQueryset() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Get a KQL queryset example"},
	})
	var exampleWorkspaceID string
	var exampleKqlQuerysetID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleKqlQuerysetID = "5b218778-e7a5-4d73-8187-f10824047715"

	exampleRes := kqlqueryset.KQLQueryset{
		Type:        to.Ptr(kqlqueryset.ItemTypeKQLQueryset),
		Description: to.Ptr("A KQL queryset description."),
		DisplayName: to.Ptr("KQLQueryset_1"),
		ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	}

	testsuite.serverFactory.ItemsServer.GetKQLQueryset = func(ctx context.Context, workspaceID string, kqlQuerysetID string, options *kqlqueryset.ItemsClientGetKQLQuerysetOptions) (resp azfake.Responder[kqlqueryset.ItemsClientGetKQLQuerysetResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleKqlQuerysetID, kqlQuerysetID)
		resp = azfake.Responder[kqlqueryset.ItemsClientGetKQLQuerysetResponse]{}
		resp.SetResponse(http.StatusOK, kqlqueryset.ItemsClientGetKQLQuerysetResponse{KQLQueryset: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	res, err := client.GetKQLQueryset(ctx, exampleWorkspaceID, exampleKqlQuerysetID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.KQLQueryset))
}

func (testsuite *FakeTestSuite) TestItems_UpdateKQLQueryset() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Update a KQL queryset example"},
	})
	var exampleWorkspaceID string
	var exampleKqlQuerysetID string
	var exampleUpdateKQLQuerysetRequest kqlqueryset.UpdateKQLQuerysetRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleKqlQuerysetID = "5b218778-e7a5-4d73-8187-f10824047715"
	exampleUpdateKQLQuerysetRequest = kqlqueryset.UpdateKQLQuerysetRequest{
		Description: to.Ptr("A new description for KQL queryset."),
		DisplayName: to.Ptr("KQLQueryset_New_Name"),
	}

	exampleRes := kqlqueryset.KQLQueryset{
		Type:        to.Ptr(kqlqueryset.ItemTypeKQLQueryset),
		Description: to.Ptr("A new description for KQL queryset."),
		DisplayName: to.Ptr("KQLQueryset_New_Name"),
		ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	}

	testsuite.serverFactory.ItemsServer.UpdateKQLQueryset = func(ctx context.Context, workspaceID string, kqlQuerysetID string, updateKQLQuerysetRequest kqlqueryset.UpdateKQLQuerysetRequest, options *kqlqueryset.ItemsClientUpdateKQLQuerysetOptions) (resp azfake.Responder[kqlqueryset.ItemsClientUpdateKQLQuerysetResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleKqlQuerysetID, kqlQuerysetID)
		testsuite.Require().True(reflect.DeepEqual(exampleUpdateKQLQuerysetRequest, updateKQLQuerysetRequest))
		resp = azfake.Responder[kqlqueryset.ItemsClientUpdateKQLQuerysetResponse]{}
		resp.SetResponse(http.StatusOK, kqlqueryset.ItemsClientUpdateKQLQuerysetResponse{KQLQueryset: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	res, err := client.UpdateKQLQueryset(ctx, exampleWorkspaceID, exampleKqlQuerysetID, exampleUpdateKQLQuerysetRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.KQLQueryset))
}

func (testsuite *FakeTestSuite) TestItems_DeleteKQLQueryset() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Delete a KQL queryset example"},
	})
	var exampleWorkspaceID string
	var exampleKqlQuerysetID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleKqlQuerysetID = "5b218778-e7a5-4d73-8187-f10824047715"

	testsuite.serverFactory.ItemsServer.DeleteKQLQueryset = func(ctx context.Context, workspaceID string, kqlQuerysetID string, options *kqlqueryset.ItemsClientDeleteKQLQuerysetOptions) (resp azfake.Responder[kqlqueryset.ItemsClientDeleteKQLQuerysetResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleKqlQuerysetID, kqlQuerysetID)
		resp = azfake.Responder[kqlqueryset.ItemsClientDeleteKQLQuerysetResponse]{}
		resp.SetResponse(http.StatusOK, kqlqueryset.ItemsClientDeleteKQLQuerysetResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	_, err = client.DeleteKQLQueryset(ctx, exampleWorkspaceID, exampleKqlQuerysetID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
}

func (testsuite *FakeTestSuite) TestItems_GetKQLQuerysetDefinition() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Get a KQL queryset definition example"},
	})
	var exampleWorkspaceID string
	var exampleKqlQuerysetID string
	exampleWorkspaceID = "314fff62-7c47-4225-9a6c-1a2220f4ce32"
	exampleKqlQuerysetID = "9b7de20c-a62f-470e-931d-e7e53f373c0c"

	exampleRes := kqlqueryset.DefinitionResponse{
		Definition: &kqlqueryset.Definition{
			Parts: []kqlqueryset.DefinitionPart{
				{
					Path:        to.Ptr("RealTimeQueryset.json"),
					Payload:     to.Ptr("ewogICAgInF1ZXJ5c2V0IjogewogICAgICAgICJ2ZXJzaW9uIjogIjEuMC4wIiwKICAgICAgICAiZGF0YVNvdXJjZXMiOiBbewogICAgICAgICAgICAgICAgImlkIjogImMyNDM0YmY4LTI1YmItNGFhMC04NzQ2LWRiNDcwNTMzYWRhZiIsCiAgICAgICAgICAgICAgICAiY2x1c3RlclVyaSI6ICJodHRwczovL2hlbHAua3VzdG8ud2luZG93cy5uZXQvIiwKICAgICAgICAgICAgICAgICJ0eXBlIjogIkF6dXJlRGF0YUV4cGxvcmVyIiwKICAgICAgICAgICAgICAgICJkYXRhYmFzZU5hbWUiOiAiU2FtcGxlcyIKICAgICAgICAgICAgfQogICAgICAgIF0sCiAgICAgICAgInRhYnMiOiBbewogICAgICAgICAgICAgICAgImlkIjogImNjZDdiOTBjLTUxZmUtNDI5Zi1hODUzLTM4NWIwMmJkNzRjOSIsCiAgICAgICAgICAgICAgICAiY29udGVudCI6ICJTdG9ybUV2ZW50c1xcXFxufCBjb3VudCIsCiAgICAgICAgICAgICAgICAidGl0bGUiOiAiVGFiMU5hbWUiLAogICAgICAgICAgICAgICAgImRhdGFTb3VyY2VJZCI6ICJjMjQzNGJmOC0yNWJiLTRhYTAtODc0Ni1kYjQ3MDUzM2FkYWYiCiAgICAgICAgICAgIH0KICAgICAgICBdCiAgICB9Cn0="),
					PayloadType: to.Ptr(kqlqueryset.PayloadTypeInlineBase64),
				},
				{
					Path:        to.Ptr(".platform"),
					Payload:     to.Ptr("ewogICIkc2NoZW1hIjogImh0dHBzOi8vZGV2ZWxvcGVyLm1pY3Jvc29mdC5jb20vanNvbi1zY2hlbWFzL2ZhYnJpYy9naXRJbnRlZ3JhdGlvbi9wbGF0Zm9ybVByb3BlcnRpZXMvMi4wLjAvc2NoZW1hLmpzb24iLAogICJtZXRhZGF0YSI6IHsKICAgICJ0eXBlIjogIktRTFF1ZXJ5c2V0IiwKICAgICJkaXNwbGF5TmFtZSI6ICJjeHp2ZWFyZ2giLAogICAgImRlc2NyaXB0aW9uIjogImN4enZlYXJnaCIKICB9LAogICJjb25maWciOiB7CiAgICAidmVyc2lvbiI6ICIyLjAiLAogICAgImxvZ2ljYWxJZCI6ICIwMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwMDAiCiAgfQp9"),
					PayloadType: to.Ptr(kqlqueryset.PayloadTypeInlineBase64),
				}},
		},
	}

	testsuite.serverFactory.ItemsServer.BeginGetKQLQuerysetDefinition = func(ctx context.Context, workspaceID string, kqlQuerysetID string, options *kqlqueryset.ItemsClientBeginGetKQLQuerysetDefinitionOptions) (resp azfake.PollerResponder[kqlqueryset.ItemsClientGetKQLQuerysetDefinitionResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleKqlQuerysetID, kqlQuerysetID)
		resp = azfake.PollerResponder[kqlqueryset.ItemsClientGetKQLQuerysetDefinitionResponse]{}
		resp.SetTerminalResponse(http.StatusOK, kqlqueryset.ItemsClientGetKQLQuerysetDefinitionResponse{DefinitionResponse: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	poller, err := client.BeginGetKQLQuerysetDefinition(ctx, exampleWorkspaceID, exampleKqlQuerysetID, &kqlqueryset.ItemsClientBeginGetKQLQuerysetDefinitionOptions{Format: nil})
	testsuite.Require().NoError(err, "Failed to get result for example ")
	res, err := poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.DefinitionResponse))
}

func (testsuite *FakeTestSuite) TestItems_UpdateKQLQuerysetDefinition() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Update a KQL queryset definition example"},
	})
	var exampleWorkspaceID string
	var exampleKqlQuerysetID string
	var exampleUpdateKQLQuerysetDefinitionRequest kqlqueryset.UpdateKQLQuerysetDefinitionRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleKqlQuerysetID = "5b218778-e7a5-4d73-8187-f10824047715"
	exampleUpdateKQLQuerysetDefinitionRequest = kqlqueryset.UpdateKQLQuerysetDefinitionRequest{
		Definition: &kqlqueryset.Definition{
			Parts: []kqlqueryset.DefinitionPart{
				{
					Path:        to.Ptr("RealTimeQueryset.json"),
					Payload:     to.Ptr("ewogICAgInF1ZXJ5c2V0IjogewogICAgICAgICJ2ZXJzaW9uIjogIjEuMC4wIiwKICAgICAgICAiZGF0YVNvdXJjZXMiOiBbewogICAgICAgICAgICAgICAgImlkIjogImMyNDM0YmY4LTI1YmItNGFhMC04NzQ2LWRiNDcwNTMzYWRhZiIsCiAgICAgICAgICAgICAgICAiY2x1c3RlclVyaSI6ICJodHRwczovL2hlbHAua3VzdG8ud2luZG93cy5uZXQvIiwKICAgICAgICAgICAgICAgICJ0eXBlIjogIkF6dXJlRGF0YUV4cGxvcmVyIiwKICAgICAgICAgICAgICAgICJkYXRhYmFzZU5hbWUiOiAiU2FtcGxlcyIKICAgICAgICAgICAgfQogICAgICAgIF0sCiAgICAgICAgInRhYnMiOiBbewogICAgICAgICAgICAgICAgImlkIjogImNjZDdiOTBjLTUxZmUtNDI5Zi1hODUzLTM4NWIwMmJkNzRjOSIsCiAgICAgICAgICAgICAgICAiY29udGVudCI6ICJTdG9ybUV2ZW50c1xcXFxufCBjb3VudCIsCiAgICAgICAgICAgICAgICAidGl0bGUiOiAiVGFiMU5hbWUiLAogICAgICAgICAgICAgICAgImRhdGFTb3VyY2VJZCI6ICJjMjQzNGJmOC0yNWJiLTRhYTAtODc0Ni1kYjQ3MDUzM2FkYWYiCiAgICAgICAgICAgIH0KICAgICAgICBdCiAgICB9Cn0="),
					PayloadType: to.Ptr(kqlqueryset.PayloadTypeInlineBase64),
				},
				{
					Path:        to.Ptr(".platform"),
					Payload:     to.Ptr("ZG90UGxhdGZvcm1CYXNlNjRTdHJpbmc="),
					PayloadType: to.Ptr(kqlqueryset.PayloadTypeInlineBase64),
				}},
		},
	}

	testsuite.serverFactory.ItemsServer.BeginUpdateKQLQuerysetDefinition = func(ctx context.Context, workspaceID string, kqlQuerysetID string, updateKQLQuerysetDefinitionRequest kqlqueryset.UpdateKQLQuerysetDefinitionRequest, options *kqlqueryset.ItemsClientBeginUpdateKQLQuerysetDefinitionOptions) (resp azfake.PollerResponder[kqlqueryset.ItemsClientUpdateKQLQuerysetDefinitionResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleKqlQuerysetID, kqlQuerysetID)
		testsuite.Require().True(reflect.DeepEqual(exampleUpdateKQLQuerysetDefinitionRequest, updateKQLQuerysetDefinitionRequest))
		resp = azfake.PollerResponder[kqlqueryset.ItemsClientUpdateKQLQuerysetDefinitionResponse]{}
		resp.SetTerminalResponse(http.StatusOK, kqlqueryset.ItemsClientUpdateKQLQuerysetDefinitionResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	poller, err := client.BeginUpdateKQLQuerysetDefinition(ctx, exampleWorkspaceID, exampleKqlQuerysetID, exampleUpdateKQLQuerysetDefinitionRequest, &kqlqueryset.ItemsClientBeginUpdateKQLQuerysetDefinitionOptions{UpdateMetadata: to.Ptr(true)})
	testsuite.Require().NoError(err, "Failed to get result for example ")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")
}
