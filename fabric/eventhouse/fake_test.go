// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package eventhouse_test

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

	"github.com/microsoft/fabric-sdk-go/fabric/eventhouse"
	"github.com/microsoft/fabric-sdk-go/fabric/eventhouse/fake"
)

var err error

type FakeTestSuite struct {
	suite.Suite

	ctx  context.Context
	cred azcore.TokenCredential

	serverFactory *fake.ServerFactory
	clientFactory *eventhouse.ClientFactory
}

func (testsuite *FakeTestSuite) SetupSuite() {
	testsuite.ctx = context.Background()
	testsuite.cred = &azfake.TokenCredential{}

	testsuite.serverFactory = &fake.ServerFactory{}
	testsuite.clientFactory, err = eventhouse.NewClientFactory(testsuite.cred, nil, &azcore.ClientOptions{
		Transport: fake.NewServerFactoryTransport(testsuite.serverFactory),
	})
	testsuite.Require().NoError(err, "Failed to create client factory")
}

func TestFakeTest(t *testing.T) {
	suite.Run(t, new(FakeTestSuite))
}

func (testsuite *FakeTestSuite) TestItems_ListEventhouses() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"List eventhouses in workspace example"},
	})
	var exampleWorkspaceID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"

	exampleRes := eventhouse.Eventhouses{
		Value: []eventhouse.Eventhouse{
			{
				Type:        to.Ptr(eventhouse.ItemTypeEventhouse),
				Description: to.Ptr("An eventhouse description."),
				DisplayName: to.Ptr("Eventhouse_1"),
				ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
				Properties: &eventhouse.Properties{
					DatabasesItemIDs: []string{
						"f1082404-7716-5b21-8778-e7a5e7a54d73",
						"8187f108-2404-4771-6e7a-5b218778e7a5"},
					IngestionServiceURI: to.Ptr("https://ingest-trd-f7k1b2rzuqrjmb3wpd.z5.kusto.fabric.microsoft.com"),
					QueryServiceURI:     to.Ptr("https://trd-f7k1b2rzuqrjmb3wpd.z5.kusto.fabric.microsoft.com"),
				},
			},
			{
				Type:        to.Ptr(eventhouse.ItemTypeEventhouse),
				Description: to.Ptr("An eventhouse description."),
				DisplayName: to.Ptr("Eventhouse_2"),
				ID:          to.Ptr("340d91b9-5a39-409c-b9c0-05ba832c476e"),
				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
				Properties: &eventhouse.Properties{
					DatabasesItemIDs: []string{
						"4d738187-f108-2404-4771-6e7a5b218778"},
					IngestionServiceURI: to.Ptr("https://ingest-trd-f7k1b2rzuqrjmb3wpd.z5.kusto.fabric.microsoft.com"),
					QueryServiceURI:     to.Ptr("https://trd-f7k1b2rzuqrjmb3wpd.z5.kusto.fabric.microsoft.com"),
				},
			}},
	}

	testsuite.serverFactory.ItemsServer.NewListEventhousesPager = func(workspaceID string, options *eventhouse.ItemsClientListEventhousesOptions) (resp azfake.PagerResponder[eventhouse.ItemsClientListEventhousesResponse]) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		resp = azfake.PagerResponder[eventhouse.ItemsClientListEventhousesResponse]{}
		resp.AddPage(http.StatusOK, eventhouse.ItemsClientListEventhousesResponse{Eventhouses: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	pager := client.NewListEventhousesPager(exampleWorkspaceID, &eventhouse.ItemsClientListEventhousesOptions{ContinuationToken: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		testsuite.Require().NoError(err, "Failed to advance page for example ")
		testsuite.Require().True(reflect.DeepEqual(exampleRes, nextResult.Eventhouses))
		if err == nil {
			break
		}
	}
}

func (testsuite *FakeTestSuite) TestItems_GetEventhouse() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Get an eventhouse example"},
	})
	var exampleWorkspaceID string
	var exampleEventhouseID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleEventhouseID = "5b218778-e7a5-4d73-8187-f10824047715"

	exampleRes := eventhouse.Eventhouse{
		Type:        to.Ptr(eventhouse.ItemTypeEventhouse),
		Description: to.Ptr("An eventhouse description."),
		DisplayName: to.Ptr("Eventhouse_1"),
		ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		Properties: &eventhouse.Properties{
			DatabasesItemIDs: []string{
				"f1082404-7716-5b21-8778-e7a5e7a54d73",
				"8187f108-2404-4771-6e7a-5b218778e7a5"},
			IngestionServiceURI: to.Ptr("https://ingest-trd-f7k1b2rzuqrjmb3wpd.z5.kusto.fabric.microsoft.com"),
			QueryServiceURI:     to.Ptr("https://trd-f7k1b2rzuqrjmb3wpd.z5.kusto.fabric.microsoft.com"),
		},
	}

	testsuite.serverFactory.ItemsServer.GetEventhouse = func(ctx context.Context, workspaceID string, eventhouseID string, options *eventhouse.ItemsClientGetEventhouseOptions) (resp azfake.Responder[eventhouse.ItemsClientGetEventhouseResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleEventhouseID, eventhouseID)
		resp = azfake.Responder[eventhouse.ItemsClientGetEventhouseResponse]{}
		resp.SetResponse(http.StatusOK, eventhouse.ItemsClientGetEventhouseResponse{Eventhouse: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	res, err := client.GetEventhouse(ctx, exampleWorkspaceID, exampleEventhouseID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.Eventhouse))
}

func (testsuite *FakeTestSuite) TestItems_UpdateEventhouse() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Update an eventhouse example"},
	})
	var exampleWorkspaceID string
	var exampleEventhouseID string
	var exampleUpdateEventhouseRequest eventhouse.UpdateEventhouseRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleEventhouseID = "5b218778-e7a5-4d73-8187-f10824047715"
	exampleUpdateEventhouseRequest = eventhouse.UpdateEventhouseRequest{
		Description: to.Ptr("A new description for eventhouse."),
		DisplayName: to.Ptr("Eventhouse_New_Name"),
	}

	exampleRes := eventhouse.Eventhouse{
		Type:        to.Ptr(eventhouse.ItemTypeEventhouse),
		Description: to.Ptr("A new description for eventhouse."),
		DisplayName: to.Ptr("Eventhouse_New_Name"),
		ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	}

	testsuite.serverFactory.ItemsServer.UpdateEventhouse = func(ctx context.Context, workspaceID string, eventhouseID string, updateEventhouseRequest eventhouse.UpdateEventhouseRequest, options *eventhouse.ItemsClientUpdateEventhouseOptions) (resp azfake.Responder[eventhouse.ItemsClientUpdateEventhouseResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleEventhouseID, eventhouseID)
		testsuite.Require().True(reflect.DeepEqual(exampleUpdateEventhouseRequest, updateEventhouseRequest))
		resp = azfake.Responder[eventhouse.ItemsClientUpdateEventhouseResponse]{}
		resp.SetResponse(http.StatusOK, eventhouse.ItemsClientUpdateEventhouseResponse{Eventhouse: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	res, err := client.UpdateEventhouse(ctx, exampleWorkspaceID, exampleEventhouseID, exampleUpdateEventhouseRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.Eventhouse))
}

func (testsuite *FakeTestSuite) TestItems_DeleteEventhouse() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Delete an eventhouse example"},
	})
	var exampleWorkspaceID string
	var exampleEventhouseID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleEventhouseID = "5b218778-e7a5-4d73-8187-f10824047715"

	testsuite.serverFactory.ItemsServer.DeleteEventhouse = func(ctx context.Context, workspaceID string, eventhouseID string, options *eventhouse.ItemsClientDeleteEventhouseOptions) (resp azfake.Responder[eventhouse.ItemsClientDeleteEventhouseResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleEventhouseID, eventhouseID)
		resp = azfake.Responder[eventhouse.ItemsClientDeleteEventhouseResponse]{}
		resp.SetResponse(http.StatusOK, eventhouse.ItemsClientDeleteEventhouseResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	_, err = client.DeleteEventhouse(ctx, exampleWorkspaceID, exampleEventhouseID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
}

func (testsuite *FakeTestSuite) TestItems_GetEventhouseDefinition() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Get an eventhouse definition example"},
	})
	var exampleWorkspaceID string
	var exampleEventhouseID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleEventhouseID = "5b218778-e7a5-4d73-8187-f10824047715"

	exampleRes := eventhouse.DefinitionResponse{
		Definition: &eventhouse.Definition{
			Parts: []eventhouse.DefinitionPart{
				{
					Path:        to.Ptr("EventhouseProperties.json"),
					Payload:     to.Ptr("e30="),
					PayloadType: to.Ptr(eventhouse.PayloadTypeInlineBase64),
				},
				{
					Path:        to.Ptr(".platform"),
					Payload:     to.Ptr("ZG90UGxhdGZvcm1CYXNlNjRTdHJpbmc"),
					PayloadType: to.Ptr(eventhouse.PayloadTypeInlineBase64),
				}},
		},
	}

	testsuite.serverFactory.ItemsServer.BeginGetEventhouseDefinition = func(ctx context.Context, workspaceID string, eventhouseID string, options *eventhouse.ItemsClientBeginGetEventhouseDefinitionOptions) (resp azfake.PollerResponder[eventhouse.ItemsClientGetEventhouseDefinitionResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleEventhouseID, eventhouseID)
		resp = azfake.PollerResponder[eventhouse.ItemsClientGetEventhouseDefinitionResponse]{}
		resp.SetTerminalResponse(http.StatusOK, eventhouse.ItemsClientGetEventhouseDefinitionResponse{DefinitionResponse: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	poller, err := client.BeginGetEventhouseDefinition(ctx, exampleWorkspaceID, exampleEventhouseID, &eventhouse.ItemsClientBeginGetEventhouseDefinitionOptions{Format: nil})
	testsuite.Require().NoError(err, "Failed to get result for example ")
	res, err := poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.DefinitionResponse))
}

func (testsuite *FakeTestSuite) TestItems_UpdateEventhouseDefinition() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Update an eventhouse definition example"},
	})
	var exampleWorkspaceID string
	var exampleEventhouseID string
	var exampleUpdateEventhouseDefinitionRequest eventhouse.UpdateEventhouseDefinitionRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleEventhouseID = "5b218778-e7a5-4d73-8187-f10824047715"
	exampleUpdateEventhouseDefinitionRequest = eventhouse.UpdateEventhouseDefinitionRequest{
		Definition: &eventhouse.Definition{
			Parts: []eventhouse.DefinitionPart{
				{
					Path:        to.Ptr("EventhouseProperties.json"),
					Payload:     to.Ptr("e30="),
					PayloadType: to.Ptr(eventhouse.PayloadTypeInlineBase64),
				},
				{
					Path:        to.Ptr(".platform"),
					Payload:     to.Ptr("ZG90UGxhdGZvcm1CYXNlNjRTdHJpbmc="),
					PayloadType: to.Ptr(eventhouse.PayloadTypeInlineBase64),
				}},
		},
	}

	testsuite.serverFactory.ItemsServer.BeginUpdateEventhouseDefinition = func(ctx context.Context, workspaceID string, eventhouseID string, updateEventhouseDefinitionRequest eventhouse.UpdateEventhouseDefinitionRequest, options *eventhouse.ItemsClientBeginUpdateEventhouseDefinitionOptions) (resp azfake.PollerResponder[eventhouse.ItemsClientUpdateEventhouseDefinitionResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleEventhouseID, eventhouseID)
		testsuite.Require().True(reflect.DeepEqual(exampleUpdateEventhouseDefinitionRequest, updateEventhouseDefinitionRequest))
		resp = azfake.PollerResponder[eventhouse.ItemsClientUpdateEventhouseDefinitionResponse]{}
		resp.SetTerminalResponse(http.StatusOK, eventhouse.ItemsClientUpdateEventhouseDefinitionResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	poller, err := client.BeginUpdateEventhouseDefinition(ctx, exampleWorkspaceID, exampleEventhouseID, exampleUpdateEventhouseDefinitionRequest, &eventhouse.ItemsClientBeginUpdateEventhouseDefinitionOptions{UpdateMetadata: to.Ptr(true)})
	testsuite.Require().NoError(err, "Failed to get result for example ")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")
}
