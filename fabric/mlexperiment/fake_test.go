// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package mlexperiment_test

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

	"github.com/microsoft/fabric-sdk-go/fabric/mlexperiment"
	"github.com/microsoft/fabric-sdk-go/fabric/mlexperiment/fake"
)

var err error

type FakeTestSuite struct {
	suite.Suite

	ctx  context.Context
	cred azcore.TokenCredential

	serverFactory *fake.ServerFactory
	clientFactory *mlexperiment.ClientFactory
}

func (testsuite *FakeTestSuite) SetupSuite() {
	testsuite.ctx = context.Background()
	testsuite.cred = &azfake.TokenCredential{}

	testsuite.serverFactory = &fake.ServerFactory{}
	testsuite.clientFactory, err = mlexperiment.NewClientFactory(testsuite.cred, nil, &azcore.ClientOptions{
		Transport: fake.NewServerFactoryTransport(testsuite.serverFactory),
	})
	testsuite.Require().NoError(err, "Failed to create client factory")
}

func TestFakeTest(t *testing.T) {
	suite.Run(t, new(FakeTestSuite))
}

func (testsuite *FakeTestSuite) TestItems_ListMLExperiments() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"List machine learning experiments in workspace example"},
	})
	var exampleWorkspaceID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"

	exampleRes := mlexperiment.MLExperiments{
		Value: []mlexperiment.MLExperiment{
			{
				Type:        to.Ptr(mlexperiment.ItemTypeMLExperiment),
				Description: to.Ptr("A machine learning experiment description."),
				DisplayName: to.Ptr("MLExperiment_1"),
				ID:          to.Ptr("3546052c-ae64-4526-b1a8-52af7761426f"),
				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
			},
			{
				Type:        to.Ptr(mlexperiment.ItemTypeMLExperiment),
				Description: to.Ptr("A machine learning experiment description."),
				DisplayName: to.Ptr("MLExperiment_2"),
				ID:          to.Ptr("6c192553-2375-4ba1-a61c-dec7729bdccc"),
				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
			}},
	}

	testsuite.serverFactory.ItemsServer.NewListMLExperimentsPager = func(workspaceID string, options *mlexperiment.ItemsClientListMLExperimentsOptions) (resp azfake.PagerResponder[mlexperiment.ItemsClientListMLExperimentsResponse]) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		resp = azfake.PagerResponder[mlexperiment.ItemsClientListMLExperimentsResponse]{}
		resp.AddPage(http.StatusOK, mlexperiment.ItemsClientListMLExperimentsResponse{MLExperiments: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	pager := client.NewListMLExperimentsPager(exampleWorkspaceID, &mlexperiment.ItemsClientListMLExperimentsOptions{ContinuationToken: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		testsuite.Require().NoError(err, "Failed to advance page for example ")
		testsuite.Require().True(reflect.DeepEqual(exampleRes, nextResult.MLExperiments))
		if err == nil {
			break
		}
	}
}

func (testsuite *FakeTestSuite) TestItems_CreateMLExperiment() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Create a machine learning experiment example"},
	})
	var exampleWorkspaceID string
	var exampleCreateMLExperimentRequest mlexperiment.CreateMLExperimentRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleCreateMLExperimentRequest = mlexperiment.CreateMLExperimentRequest{
		Description: to.Ptr("A machine learning experiment description"),
		DisplayName: to.Ptr("MLExperiment_1"),
	}

	testsuite.serverFactory.ItemsServer.BeginCreateMLExperiment = func(ctx context.Context, workspaceID string, createMLExperimentRequest mlexperiment.CreateMLExperimentRequest, options *mlexperiment.ItemsClientBeginCreateMLExperimentOptions) (resp azfake.PollerResponder[mlexperiment.ItemsClientCreateMLExperimentResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().True(reflect.DeepEqual(exampleCreateMLExperimentRequest, createMLExperimentRequest))
		resp = azfake.PollerResponder[mlexperiment.ItemsClientCreateMLExperimentResponse]{}
		resp.SetTerminalResponse(http.StatusCreated, mlexperiment.ItemsClientCreateMLExperimentResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	poller, err := client.BeginCreateMLExperiment(ctx, exampleWorkspaceID, exampleCreateMLExperimentRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")
}

func (testsuite *FakeTestSuite) TestItems_GetMLExperiment() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Get a machine learning experiment example"},
	})
	var exampleWorkspaceID string
	var exampleMlExperimentID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleMlExperimentID = "5b218778-e7a5-4d73-8187-f10824047715"

	exampleRes := mlexperiment.MLExperiment{
		Type:        to.Ptr(mlexperiment.ItemTypeMLExperiment),
		Description: to.Ptr("A machine learning experiment description"),
		DisplayName: to.Ptr("MLExperiment_1"),
		ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	}

	testsuite.serverFactory.ItemsServer.GetMLExperiment = func(ctx context.Context, workspaceID string, mlExperimentID string, options *mlexperiment.ItemsClientGetMLExperimentOptions) (resp azfake.Responder[mlexperiment.ItemsClientGetMLExperimentResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleMlExperimentID, mlExperimentID)
		resp = azfake.Responder[mlexperiment.ItemsClientGetMLExperimentResponse]{}
		resp.SetResponse(http.StatusOK, mlexperiment.ItemsClientGetMLExperimentResponse{MLExperiment: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	res, err := client.GetMLExperiment(ctx, exampleWorkspaceID, exampleMlExperimentID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.MLExperiment))
}

func (testsuite *FakeTestSuite) TestItems_UpdateMLExperiment() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Update an machine learning experiment example"},
	})
	var exampleWorkspaceID string
	var exampleMlExperimentID string
	var exampleUpdateMLExperimentRequest mlexperiment.UpdateMLExperimentRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleMlExperimentID = "5b218778-e7a5-4d73-8187-f10824047715"
	exampleUpdateMLExperimentRequest = mlexperiment.UpdateMLExperimentRequest{
		Description: to.Ptr("A new description for machine learning experiment."),
		DisplayName: to.Ptr("MLExperiment_New_Name"),
	}

	exampleRes := mlexperiment.MLExperiment{
		Type:        to.Ptr(mlexperiment.ItemTypeMLExperiment),
		Description: to.Ptr("A new description for machine learning experiment."),
		DisplayName: to.Ptr("MLExperiment_New_Name"),
		ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	}

	testsuite.serverFactory.ItemsServer.UpdateMLExperiment = func(ctx context.Context, workspaceID string, mlExperimentID string, updateMLExperimentRequest mlexperiment.UpdateMLExperimentRequest, options *mlexperiment.ItemsClientUpdateMLExperimentOptions) (resp azfake.Responder[mlexperiment.ItemsClientUpdateMLExperimentResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleMlExperimentID, mlExperimentID)
		testsuite.Require().True(reflect.DeepEqual(exampleUpdateMLExperimentRequest, updateMLExperimentRequest))
		resp = azfake.Responder[mlexperiment.ItemsClientUpdateMLExperimentResponse]{}
		resp.SetResponse(http.StatusOK, mlexperiment.ItemsClientUpdateMLExperimentResponse{MLExperiment: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	res, err := client.UpdateMLExperiment(ctx, exampleWorkspaceID, exampleMlExperimentID, exampleUpdateMLExperimentRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.MLExperiment))
}

func (testsuite *FakeTestSuite) TestItems_DeleteMLExperiment() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Delete a machine learning experiment example"},
	})
	var exampleWorkspaceID string
	var exampleMlExperimentID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleMlExperimentID = "5b218778-e7a5-4d73-8187-f10824047715"

	testsuite.serverFactory.ItemsServer.DeleteMLExperiment = func(ctx context.Context, workspaceID string, mlExperimentID string, options *mlexperiment.ItemsClientDeleteMLExperimentOptions) (resp azfake.Responder[mlexperiment.ItemsClientDeleteMLExperimentResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleMlExperimentID, mlExperimentID)
		resp = azfake.Responder[mlexperiment.ItemsClientDeleteMLExperimentResponse]{}
		resp.SetResponse(http.StatusOK, mlexperiment.ItemsClientDeleteMLExperimentResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	_, err = client.DeleteMLExperiment(ctx, exampleWorkspaceID, exampleMlExperimentID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
}
