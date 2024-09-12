// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package kqldatabase_test

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

	"github.com/microsoft/fabric-sdk-go/fabric/kqldatabase"
	"github.com/microsoft/fabric-sdk-go/fabric/kqldatabase/fake"
)

var err error

type FakeTestSuite struct {
	suite.Suite

	ctx  context.Context
	cred azcore.TokenCredential

	serverFactory *fake.ServerFactory
	clientFactory *kqldatabase.ClientFactory
}

func (testsuite *FakeTestSuite) SetupSuite() {
	testsuite.ctx = context.Background()
	testsuite.cred = &azfake.TokenCredential{}

	testsuite.serverFactory = &fake.ServerFactory{}
	testsuite.clientFactory, err = kqldatabase.NewClientFactory(testsuite.cred, nil, &azcore.ClientOptions{
		Transport: fake.NewServerFactoryTransport(testsuite.serverFactory),
	})
	testsuite.Require().NoError(err, "Failed to create client factory")
}

func TestFakeTest(t *testing.T) {
	suite.Run(t, new(FakeTestSuite))
}

func (testsuite *FakeTestSuite) TestItems_ListKQLDatabases() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"List KQL databases in workspace example"},
	})
	var exampleWorkspaceID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"

	exampleRes := kqldatabase.KQLDatabases{
		Value: []kqldatabase.KQLDatabase{
			{
				Type:        to.Ptr(kqldatabase.ItemTypeKQLDatabase),
				Description: to.Ptr("A KQL database description."),
				DisplayName: to.Ptr("KQLDatabase_1"),
				ID:          to.Ptr("3546052c-ae64-4526-b1a8-52af7761426f"),
				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
				Properties: &kqldatabase.Properties{
					DatabaseType:           to.Ptr(kqldatabase.TypeReadWrite),
					IngestionServiceURI:    to.Ptr("https://ingest-trd-f7k1b2rzuqrjmb3wpd.z5.kusto.fabric.microsoft.com"),
					ParentEventhouseItemID: to.Ptr("6a437a7c-1a28-4fd0-a362-11308b94c79b"),
					QueryServiceURI:        to.Ptr("https://trd-f7k1b2rzuqrjmb3wpd.z5.kusto.fabric.microsoft.com"),
				},
			},
			{
				Type:        to.Ptr(kqldatabase.ItemTypeKQLDatabase),
				Description: to.Ptr("A KQL database description."),
				DisplayName: to.Ptr("KQLDatabase_2"),
				ID:          to.Ptr("340d91b9-5a39-409c-b9c0-05ba832c476e"),
				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
				Properties: &kqldatabase.Properties{
					DatabaseType:           to.Ptr(kqldatabase.TypeReadWrite),
					IngestionServiceURI:    to.Ptr("https://ingest-trd-f7k1b2rzuqrjmb3wpd.z5.kusto.fabric.microsoft.com"),
					ParentEventhouseItemID: to.Ptr("9add9a4d-079a-432a-b43a-70c899f2087b"),
					QueryServiceURI:        to.Ptr("https://trd-f7k1b2rzuqrjmb3wpd.z5.kusto.fabric.microsoft.com"),
				},
			}},
	}

	testsuite.serverFactory.ItemsServer.NewListKQLDatabasesPager = func(workspaceID string, options *kqldatabase.ItemsClientListKQLDatabasesOptions) (resp azfake.PagerResponder[kqldatabase.ItemsClientListKQLDatabasesResponse]) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		resp = azfake.PagerResponder[kqldatabase.ItemsClientListKQLDatabasesResponse]{}
		resp.AddPage(http.StatusOK, kqldatabase.ItemsClientListKQLDatabasesResponse{KQLDatabases: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	pager := client.NewListKQLDatabasesPager(exampleWorkspaceID, &kqldatabase.ItemsClientListKQLDatabasesOptions{ContinuationToken: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		testsuite.Require().NoError(err, "Failed to advance page for example ")
		testsuite.Require().True(reflect.DeepEqual(exampleRes, nextResult.KQLDatabases))
		if err == nil {
			break
		}
	}
}

func (testsuite *FakeTestSuite) TestItems_CreateKQLDatabase() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Create a ReadWrite KQL database example"},
	})
	var exampleWorkspaceID string
	var exampleCreateKQLDatabaseRequest kqldatabase.CreateKQLDatabaseRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleCreateKQLDatabaseRequest = kqldatabase.CreateKQLDatabaseRequest{
		Description: to.Ptr("A KQL database description."),
		CreationPayload: &kqldatabase.ReadWriteDatabaseCreationPayload{
			DatabaseType:           to.Ptr(kqldatabase.TypeReadWrite),
			ParentEventhouseItemID: to.Ptr("5b218778-e7a5-4d73-8187-f10824047836"),
		},
		DisplayName: to.Ptr("KQLDatabase_1"),
	}

	testsuite.serverFactory.ItemsServer.BeginCreateKQLDatabase = func(ctx context.Context, workspaceID string, createKQLDatabaseRequest kqldatabase.CreateKQLDatabaseRequest, options *kqldatabase.ItemsClientBeginCreateKQLDatabaseOptions) (resp azfake.PollerResponder[kqldatabase.ItemsClientCreateKQLDatabaseResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().True(reflect.DeepEqual(exampleCreateKQLDatabaseRequest, createKQLDatabaseRequest))
		resp = azfake.PollerResponder[kqldatabase.ItemsClientCreateKQLDatabaseResponse]{}
		resp.SetTerminalResponse(http.StatusOK, kqldatabase.ItemsClientCreateKQLDatabaseResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	poller, err := client.BeginCreateKQLDatabase(ctx, exampleWorkspaceID, exampleCreateKQLDatabaseRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")

	// From example
	ctx = runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Create a Shortcut KQL database to source Azure Data Explorer cluster example"},
	})
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleCreateKQLDatabaseRequest = kqldatabase.CreateKQLDatabaseRequest{
		Description: to.Ptr("A KQL database description."),
		CreationPayload: &kqldatabase.ShortcutDatabaseCreationPayload{
			DatabaseType:           to.Ptr(kqldatabase.TypeShortcut),
			ParentEventhouseItemID: to.Ptr("5b218778-e7a5-4d73-8187-f10824047836"),
			SourceClusterURI:       to.Ptr("https://adxcluster.westus.kusto.windows.net"),
			SourceDatabaseName:     to.Ptr("MyDatabase"),
		},
		DisplayName: to.Ptr("KQLDatabase_1"),
	}

	testsuite.serverFactory.ItemsServer.BeginCreateKQLDatabase = func(ctx context.Context, workspaceID string, createKQLDatabaseRequest kqldatabase.CreateKQLDatabaseRequest, options *kqldatabase.ItemsClientBeginCreateKQLDatabaseOptions) (resp azfake.PollerResponder[kqldatabase.ItemsClientCreateKQLDatabaseResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().True(reflect.DeepEqual(exampleCreateKQLDatabaseRequest, createKQLDatabaseRequest))
		resp = azfake.PollerResponder[kqldatabase.ItemsClientCreateKQLDatabaseResponse]{}
		resp.SetTerminalResponse(http.StatusOK, kqldatabase.ItemsClientCreateKQLDatabaseResponse{}, nil)
		return
	}

	poller, err = client.BeginCreateKQLDatabase(ctx, exampleWorkspaceID, exampleCreateKQLDatabaseRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")

	// From example
	ctx = runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Create a Shortcut KQL database to source Azure Data Explorer cluster with invitation token example"},
	})
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleCreateKQLDatabaseRequest = kqldatabase.CreateKQLDatabaseRequest{
		Description: to.Ptr("A KQL database description."),
		CreationPayload: &kqldatabase.ShortcutDatabaseCreationPayload{
			DatabaseType:           to.Ptr(kqldatabase.TypeShortcut),
			ParentEventhouseItemID: to.Ptr("5b218778-e7a5-4d73-8187-f10824047836"),
			InvitationToken:        to.Ptr("eyJ0eXAiOiJKVInvitationToken"),
		},
		DisplayName: to.Ptr("KQLDatabase_1"),
	}

	testsuite.serverFactory.ItemsServer.BeginCreateKQLDatabase = func(ctx context.Context, workspaceID string, createKQLDatabaseRequest kqldatabase.CreateKQLDatabaseRequest, options *kqldatabase.ItemsClientBeginCreateKQLDatabaseOptions) (resp azfake.PollerResponder[kqldatabase.ItemsClientCreateKQLDatabaseResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().True(reflect.DeepEqual(exampleCreateKQLDatabaseRequest, createKQLDatabaseRequest))
		resp = azfake.PollerResponder[kqldatabase.ItemsClientCreateKQLDatabaseResponse]{}
		resp.SetTerminalResponse(http.StatusOK, kqldatabase.ItemsClientCreateKQLDatabaseResponse{}, nil)
		return
	}

	poller, err = client.BeginCreateKQLDatabase(ctx, exampleWorkspaceID, exampleCreateKQLDatabaseRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")

	// From example
	ctx = runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Create a Shortcut KQL database to source KQL database example"},
	})
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleCreateKQLDatabaseRequest = kqldatabase.CreateKQLDatabaseRequest{
		Description: to.Ptr("A KQL database description."),
		CreationPayload: &kqldatabase.ShortcutDatabaseCreationPayload{
			DatabaseType:           to.Ptr(kqldatabase.TypeShortcut),
			ParentEventhouseItemID: to.Ptr("5b218778-e7a5-4d73-8187-f10824047836"),
			SourceDatabaseName:     to.Ptr("ac542109-abd1-4ee3-aec5-86282c01ee24"),
		},
		DisplayName: to.Ptr("KQLDatabase_1"),
	}

	testsuite.serverFactory.ItemsServer.BeginCreateKQLDatabase = func(ctx context.Context, workspaceID string, createKQLDatabaseRequest kqldatabase.CreateKQLDatabaseRequest, options *kqldatabase.ItemsClientBeginCreateKQLDatabaseOptions) (resp azfake.PollerResponder[kqldatabase.ItemsClientCreateKQLDatabaseResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().True(reflect.DeepEqual(exampleCreateKQLDatabaseRequest, createKQLDatabaseRequest))
		resp = azfake.PollerResponder[kqldatabase.ItemsClientCreateKQLDatabaseResponse]{}
		resp.SetTerminalResponse(http.StatusOK, kqldatabase.ItemsClientCreateKQLDatabaseResponse{}, nil)
		return
	}

	poller, err = client.BeginCreateKQLDatabase(ctx, exampleWorkspaceID, exampleCreateKQLDatabaseRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")
}

func (testsuite *FakeTestSuite) TestItems_GetKQLDatabase() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Get a KQL database example"},
	})
	var exampleWorkspaceID string
	var exampleKqlDatabaseID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleKqlDatabaseID = "5b218778-e7a5-4d73-8187-f10824047715"

	exampleRes := kqldatabase.KQLDatabase{
		Type:        to.Ptr(kqldatabase.ItemTypeKQLDatabase),
		Description: to.Ptr("A KQL database description."),
		DisplayName: to.Ptr("KQLDatabase_1"),
		ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		Properties: &kqldatabase.Properties{
			DatabaseType:           to.Ptr(kqldatabase.TypeReadWrite),
			IngestionServiceURI:    to.Ptr("https://ingest-trd-f7k1b2rzuqrjmb3wpd.z5.kusto.fabric.microsoft.com"),
			ParentEventhouseItemID: to.Ptr("6a437a7c-1a28-4fd0-a362-11308b94c79b"),
			QueryServiceURI:        to.Ptr("https://trd-f7k1b2rzuqrjmb3wpd.z5.kusto.fabric.microsoft.com"),
		},
	}

	testsuite.serverFactory.ItemsServer.GetKQLDatabase = func(ctx context.Context, workspaceID string, kqlDatabaseID string, options *kqldatabase.ItemsClientGetKQLDatabaseOptions) (resp azfake.Responder[kqldatabase.ItemsClientGetKQLDatabaseResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleKqlDatabaseID, kqlDatabaseID)
		resp = azfake.Responder[kqldatabase.ItemsClientGetKQLDatabaseResponse]{}
		resp.SetResponse(http.StatusOK, kqldatabase.ItemsClientGetKQLDatabaseResponse{KQLDatabase: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	res, err := client.GetKQLDatabase(ctx, exampleWorkspaceID, exampleKqlDatabaseID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.KQLDatabase))
}

func (testsuite *FakeTestSuite) TestItems_UpdateKQLDatabase() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Update a KQL database example"},
	})
	var exampleWorkspaceID string
	var exampleKqlDatabaseID string
	var exampleUpdateKQLDatabaseRequest kqldatabase.UpdateKQLDatabaseRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleKqlDatabaseID = "5b218778-e7a5-4d73-8187-f10824047715"
	exampleUpdateKQLDatabaseRequest = kqldatabase.UpdateKQLDatabaseRequest{
		Description: to.Ptr("A new description for KQL database."),
		DisplayName: to.Ptr("KQLDatabase_New_Name"),
	}

	exampleRes := kqldatabase.KQLDatabase{
		Type:        to.Ptr(kqldatabase.ItemTypeKQLDatabase),
		Description: to.Ptr("A new description for KQL database."),
		DisplayName: to.Ptr("KQLDatabase_New_Name"),
		ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	}

	testsuite.serverFactory.ItemsServer.UpdateKQLDatabase = func(ctx context.Context, workspaceID string, kqlDatabaseID string, updateKQLDatabaseRequest kqldatabase.UpdateKQLDatabaseRequest, options *kqldatabase.ItemsClientUpdateKQLDatabaseOptions) (resp azfake.Responder[kqldatabase.ItemsClientUpdateKQLDatabaseResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleKqlDatabaseID, kqlDatabaseID)
		testsuite.Require().True(reflect.DeepEqual(exampleUpdateKQLDatabaseRequest, updateKQLDatabaseRequest))
		resp = azfake.Responder[kqldatabase.ItemsClientUpdateKQLDatabaseResponse]{}
		resp.SetResponse(http.StatusOK, kqldatabase.ItemsClientUpdateKQLDatabaseResponse{KQLDatabase: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	res, err := client.UpdateKQLDatabase(ctx, exampleWorkspaceID, exampleKqlDatabaseID, exampleUpdateKQLDatabaseRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.KQLDatabase))
}

func (testsuite *FakeTestSuite) TestItems_DeleteKQLDatabase() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Delete a KQL database example"},
	})
	var exampleWorkspaceID string
	var exampleKqlDatabaseID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleKqlDatabaseID = "5b218778-e7a5-4d73-8187-f10824047715"

	testsuite.serverFactory.ItemsServer.DeleteKQLDatabase = func(ctx context.Context, workspaceID string, kqlDatabaseID string, options *kqldatabase.ItemsClientDeleteKQLDatabaseOptions) (resp azfake.Responder[kqldatabase.ItemsClientDeleteKQLDatabaseResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleKqlDatabaseID, kqlDatabaseID)
		resp = azfake.Responder[kqldatabase.ItemsClientDeleteKQLDatabaseResponse]{}
		resp.SetResponse(http.StatusOK, kqldatabase.ItemsClientDeleteKQLDatabaseResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	_, err = client.DeleteKQLDatabase(ctx, exampleWorkspaceID, exampleKqlDatabaseID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
}
