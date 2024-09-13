// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package datapipeline_test

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

	"github.com/microsoft/fabric-sdk-go/fabric/datapipeline"
	"github.com/microsoft/fabric-sdk-go/fabric/datapipeline/fake"
)

var err error

type FakeTestSuite struct {
	suite.Suite

	ctx  context.Context
	cred azcore.TokenCredential

	serverFactory *fake.ServerFactory
	clientFactory *datapipeline.ClientFactory
}

func (testsuite *FakeTestSuite) SetupSuite() {
	testsuite.ctx = context.Background()
	testsuite.cred = &azfake.TokenCredential{}

	testsuite.serverFactory = &fake.ServerFactory{}
	testsuite.clientFactory, err = datapipeline.NewClientFactory(testsuite.cred, nil, &azcore.ClientOptions{
		Transport: fake.NewServerFactoryTransport(testsuite.serverFactory),
	})
	testsuite.Require().NoError(err, "Failed to create client factory")
}

func TestFakeTest(t *testing.T) {
	suite.Run(t, new(FakeTestSuite))
}

func (testsuite *FakeTestSuite) TestItems_ListDataPipelines() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"List data pipelines in workspace example"},
	})
	var exampleWorkspaceID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"

	exampleRes := datapipeline.DataPipelines{
		Value: []datapipeline.DataPipeline{
			{
				Type:        to.Ptr(datapipeline.ItemTypeDataPipeline),
				Description: to.Ptr("A data pipeline description."),
				DisplayName: to.Ptr("DataPipeline Name 1"),
				ID:          to.Ptr("3546052c-ae64-4526-b1a8-52af7761426f"),
				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
			}},
	}

	testsuite.serverFactory.ItemsServer.NewListDataPipelinesPager = func(workspaceID string, options *datapipeline.ItemsClientListDataPipelinesOptions) (resp azfake.PagerResponder[datapipeline.ItemsClientListDataPipelinesResponse]) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		resp = azfake.PagerResponder[datapipeline.ItemsClientListDataPipelinesResponse]{}
		resp.AddPage(http.StatusOK, datapipeline.ItemsClientListDataPipelinesResponse{DataPipelines: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	pager := client.NewListDataPipelinesPager(exampleWorkspaceID, &datapipeline.ItemsClientListDataPipelinesOptions{ContinuationToken: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		testsuite.Require().NoError(err, "Failed to advance page for example ")
		testsuite.Require().True(reflect.DeepEqual(exampleRes, nextResult.DataPipelines))
		if err == nil {
			break
		}
	}
}

func (testsuite *FakeTestSuite) TestItems_CreateDataPipeline() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Create data pipeline example"},
	})
	var exampleWorkspaceID string
	var exampleCreateDataPipelineRequest datapipeline.CreateDataPipelineRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleCreateDataPipelineRequest = datapipeline.CreateDataPipelineRequest{
		Description: to.Ptr("A data pipeline description"),
		DisplayName: to.Ptr("DataPipeline 1"),
	}

	testsuite.serverFactory.ItemsServer.BeginCreateDataPipeline = func(ctx context.Context, workspaceID string, createDataPipelineRequest datapipeline.CreateDataPipelineRequest, options *datapipeline.ItemsClientBeginCreateDataPipelineOptions) (resp azfake.PollerResponder[datapipeline.ItemsClientCreateDataPipelineResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().True(reflect.DeepEqual(exampleCreateDataPipelineRequest, createDataPipelineRequest))
		resp = azfake.PollerResponder[datapipeline.ItemsClientCreateDataPipelineResponse]{}
		resp.SetTerminalResponse(http.StatusCreated, datapipeline.ItemsClientCreateDataPipelineResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	poller, err := client.BeginCreateDataPipeline(ctx, exampleWorkspaceID, exampleCreateDataPipelineRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")
}

func (testsuite *FakeTestSuite) TestItems_GetDataPipeline() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Get a data pipeline example"},
	})
	var exampleWorkspaceID string
	var exampleDataPipelineID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleDataPipelineID = "5b218778-e7a5-4d73-8187-f10824047715"

	exampleRes := datapipeline.DataPipeline{
		Type:        to.Ptr(datapipeline.ItemTypeDataPipeline),
		Description: to.Ptr("A data pipeline description."),
		DisplayName: to.Ptr("DataPipeline 1"),
		ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	}

	testsuite.serverFactory.ItemsServer.GetDataPipeline = func(ctx context.Context, workspaceID string, dataPipelineID string, options *datapipeline.ItemsClientGetDataPipelineOptions) (resp azfake.Responder[datapipeline.ItemsClientGetDataPipelineResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleDataPipelineID, dataPipelineID)
		resp = azfake.Responder[datapipeline.ItemsClientGetDataPipelineResponse]{}
		resp.SetResponse(http.StatusOK, datapipeline.ItemsClientGetDataPipelineResponse{DataPipeline: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	res, err := client.GetDataPipeline(ctx, exampleWorkspaceID, exampleDataPipelineID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.DataPipeline))
}

func (testsuite *FakeTestSuite) TestItems_UpdateDataPipeline() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Update a data pipeline example"},
	})
	var exampleWorkspaceID string
	var exampleDataPipelineID string
	var exampleUpdateDataPipelineRequest datapipeline.UpdateDataPipelineRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleDataPipelineID = "5b218778-e7a5-4d73-8187-f10824047715"
	exampleUpdateDataPipelineRequest = datapipeline.UpdateDataPipelineRequest{
		Description: to.Ptr("A new description for the data pipeline."),
		DisplayName: to.Ptr("A new name for the DataPipeline"),
	}

	exampleRes := datapipeline.DataPipeline{
		Type:        to.Ptr(datapipeline.ItemTypeDataPipeline),
		Description: to.Ptr("A new description for the data pipeline."),
		DisplayName: to.Ptr("A new name for the DataPipeline"),
		ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	}

	testsuite.serverFactory.ItemsServer.UpdateDataPipeline = func(ctx context.Context, workspaceID string, dataPipelineID string, updateDataPipelineRequest datapipeline.UpdateDataPipelineRequest, options *datapipeline.ItemsClientUpdateDataPipelineOptions) (resp azfake.Responder[datapipeline.ItemsClientUpdateDataPipelineResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleDataPipelineID, dataPipelineID)
		testsuite.Require().True(reflect.DeepEqual(exampleUpdateDataPipelineRequest, updateDataPipelineRequest))
		resp = azfake.Responder[datapipeline.ItemsClientUpdateDataPipelineResponse]{}
		resp.SetResponse(http.StatusOK, datapipeline.ItemsClientUpdateDataPipelineResponse{DataPipeline: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	res, err := client.UpdateDataPipeline(ctx, exampleWorkspaceID, exampleDataPipelineID, exampleUpdateDataPipelineRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.DataPipeline))
}

func (testsuite *FakeTestSuite) TestItems_DeleteDataPipeline() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Delete a data pipeline example"},
	})
	var exampleWorkspaceID string
	var exampleDataPipelineID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleDataPipelineID = "5b218778-e7a5-4d73-8187-f10824047715"

	testsuite.serverFactory.ItemsServer.DeleteDataPipeline = func(ctx context.Context, workspaceID string, dataPipelineID string, options *datapipeline.ItemsClientDeleteDataPipelineOptions) (resp azfake.Responder[datapipeline.ItemsClientDeleteDataPipelineResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleDataPipelineID, dataPipelineID)
		resp = azfake.Responder[datapipeline.ItemsClientDeleteDataPipelineResponse]{}
		resp.SetResponse(http.StatusOK, datapipeline.ItemsClientDeleteDataPipelineResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	_, err = client.DeleteDataPipeline(ctx, exampleWorkspaceID, exampleDataPipelineID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
}