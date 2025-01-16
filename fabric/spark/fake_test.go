// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package spark_test

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

	"github.com/microsoft/fabric-sdk-go/fabric/spark"
	"github.com/microsoft/fabric-sdk-go/fabric/spark/fake"
)

var err error

type FakeTestSuite struct {
	suite.Suite

	ctx  context.Context
	cred azcore.TokenCredential

	serverFactory *fake.ServerFactory
	clientFactory *spark.ClientFactory
}

func (testsuite *FakeTestSuite) SetupSuite() {
	testsuite.ctx = context.Background()
	testsuite.cred = &azfake.TokenCredential{}

	testsuite.serverFactory = &fake.ServerFactory{}
	testsuite.clientFactory, err = spark.NewClientFactory(testsuite.cred, nil, &azcore.ClientOptions{
		Transport: fake.NewServerFactoryTransport(testsuite.serverFactory),
	})
	testsuite.Require().NoError(err, "Failed to create client factory")
}

func TestFakeTest(t *testing.T) {
	suite.Run(t, new(FakeTestSuite))
}

func (testsuite *FakeTestSuite) TestWorkspaceSettings_GetSparkSettings() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Get workspace Spark settings example"},
	})
	var exampleWorkspaceID string
	exampleWorkspaceID = "f089354e-8366-4e18-aea3-4cb4a3a50b48"

	exampleRes := spark.WorkspaceSparkSettings{
		AutomaticLog: &spark.AutomaticLogProperties{
			Enabled: to.Ptr(true),
		},
		Environment: &spark.EnvironmentProperties{
			Name:           to.Ptr("environment1"),
			RuntimeVersion: to.Ptr("1.2"),
		},
		HighConcurrency: &spark.HighConcurrencyProperties{
			NotebookInteractiveRunEnabled: to.Ptr(true),
			NotebookPipelineRunEnabled:    to.Ptr(false),
		},
		Jobs: &spark.JobsProperties{
			ConservativeJobAdmissionEnabled: to.Ptr(false),
			SessionTimeoutInMinutes:         to.Ptr[int32](20),
		},
		Pool: &spark.PoolProperties{
			CustomizeComputeEnabled: to.Ptr(true),
			DefaultPool: &spark.InstancePool{
				Name: to.Ptr("Starter Pool"),
				Type: to.Ptr(spark.CustomPoolTypeWorkspace),
			},
			StarterPool: &spark.StarterPoolProperties{
				MaxExecutors: to.Ptr[int32](1),
				MaxNodeCount: to.Ptr[int32](2),
			},
		},
	}

	testsuite.serverFactory.WorkspaceSettingsServer.GetSparkSettings = func(ctx context.Context, workspaceID string, options *spark.WorkspaceSettingsClientGetSparkSettingsOptions) (resp azfake.Responder[spark.WorkspaceSettingsClientGetSparkSettingsResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		resp = azfake.Responder[spark.WorkspaceSettingsClientGetSparkSettingsResponse]{}
		resp.SetResponse(http.StatusOK, spark.WorkspaceSettingsClientGetSparkSettingsResponse{WorkspaceSparkSettings: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewWorkspaceSettingsClient()
	res, err := client.GetSparkSettings(ctx, exampleWorkspaceID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.WorkspaceSparkSettings))
}

func (testsuite *FakeTestSuite) TestWorkspaceSettings_UpdateSparkSettings() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Update workspace Spark settings example"},
	})
	var exampleWorkspaceID string
	var exampleUpdateWorkspaceSettingsRequest spark.UpdateWorkspaceSparkSettingsRequest
	exampleWorkspaceID = "f089354e-8366-4e18-aea3-4cb4a3a50b48"
	exampleUpdateWorkspaceSettingsRequest = spark.UpdateWorkspaceSparkSettingsRequest{
		AutomaticLog: &spark.AutomaticLogProperties{
			Enabled: to.Ptr(false),
		},
		Environment: &spark.EnvironmentProperties{
			Name:           to.Ptr("environment1"),
			RuntimeVersion: to.Ptr("1.2"),
		},
		HighConcurrency: &spark.HighConcurrencyProperties{
			NotebookInteractiveRunEnabled: to.Ptr(false),
			NotebookPipelineRunEnabled:    to.Ptr(false),
		},
		Jobs: &spark.JobsProperties{
			ConservativeJobAdmissionEnabled: to.Ptr(false),
			SessionTimeoutInMinutes:         to.Ptr[int32](20),
		},
		Pool: &spark.PoolProperties{
			CustomizeComputeEnabled: to.Ptr(false),
			DefaultPool: &spark.InstancePool{
				Name: to.Ptr("Starter Pool"),
				Type: to.Ptr(spark.CustomPoolTypeWorkspace),
			},
			StarterPool: &spark.StarterPoolProperties{
				MaxExecutors: to.Ptr[int32](1),
				MaxNodeCount: to.Ptr[int32](3),
			},
		},
	}

	exampleRes := spark.WorkspaceSparkSettings{
		AutomaticLog: &spark.AutomaticLogProperties{
			Enabled: to.Ptr(false),
		},
		Environment: &spark.EnvironmentProperties{
			Name:           to.Ptr("environment1"),
			RuntimeVersion: to.Ptr("1.2"),
		},
		HighConcurrency: &spark.HighConcurrencyProperties{
			NotebookInteractiveRunEnabled: to.Ptr(false),
			NotebookPipelineRunEnabled:    to.Ptr(false),
		},
		Jobs: &spark.JobsProperties{
			ConservativeJobAdmissionEnabled: to.Ptr(false),
			SessionTimeoutInMinutes:         to.Ptr[int32](20),
		},
		Pool: &spark.PoolProperties{
			CustomizeComputeEnabled: to.Ptr(false),
			DefaultPool: &spark.InstancePool{
				Name: to.Ptr("Starter Pool"),
				Type: to.Ptr(spark.CustomPoolTypeWorkspace),
			},
			StarterPool: &spark.StarterPoolProperties{
				MaxExecutors: to.Ptr[int32](1),
				MaxNodeCount: to.Ptr[int32](3),
			},
		},
	}

	testsuite.serverFactory.WorkspaceSettingsServer.UpdateSparkSettings = func(ctx context.Context, workspaceID string, updateWorkspaceSettingsRequest spark.UpdateWorkspaceSparkSettingsRequest, options *spark.WorkspaceSettingsClientUpdateSparkSettingsOptions) (resp azfake.Responder[spark.WorkspaceSettingsClientUpdateSparkSettingsResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().True(reflect.DeepEqual(exampleUpdateWorkspaceSettingsRequest, updateWorkspaceSettingsRequest))
		resp = azfake.Responder[spark.WorkspaceSettingsClientUpdateSparkSettingsResponse]{}
		resp.SetResponse(http.StatusOK, spark.WorkspaceSettingsClientUpdateSparkSettingsResponse{WorkspaceSparkSettings: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewWorkspaceSettingsClient()
	res, err := client.UpdateSparkSettings(ctx, exampleWorkspaceID, exampleUpdateWorkspaceSettingsRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.WorkspaceSparkSettings))
}

func (testsuite *FakeTestSuite) TestCustomPools_ListWorkspaceCustomPools() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"List custom pools example"},
	})
	var exampleWorkspaceID string
	exampleWorkspaceID = "f089354e-8366-4e18-aea3-4cb4a3a50b48"

	exampleRes := spark.CustomPools{
		Value: []spark.CustomPool{
			{
				Name: to.Ptr("pool1"),
				Type: to.Ptr(spark.CustomPoolTypeWorkspace),
				AutoScale: &spark.AutoScaleProperties{
					Enabled:      to.Ptr(true),
					MaxNodeCount: to.Ptr[int32](4),
					MinNodeCount: to.Ptr[int32](1),
				},
				DynamicExecutorAllocation: &spark.DynamicExecutorAllocationProperties{
					Enabled:      to.Ptr(true),
					MaxExecutors: to.Ptr[int32](2),
					MinExecutors: to.Ptr[int32](1),
				},
				ID:         to.Ptr("2367293d-b70b-4b33-97f2-161b8d04a8d7"),
				NodeFamily: to.Ptr(spark.NodeFamilyMemoryOptimized),
				NodeSize:   to.Ptr(spark.NodeSizeSmall),
			}},
	}

	testsuite.serverFactory.CustomPoolsServer.NewListWorkspaceCustomPoolsPager = func(workspaceID string, options *spark.CustomPoolsClientListWorkspaceCustomPoolsOptions) (resp azfake.PagerResponder[spark.CustomPoolsClientListWorkspaceCustomPoolsResponse]) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		resp = azfake.PagerResponder[spark.CustomPoolsClientListWorkspaceCustomPoolsResponse]{}
		resp.AddPage(http.StatusOK, spark.CustomPoolsClientListWorkspaceCustomPoolsResponse{CustomPools: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewCustomPoolsClient()
	pager := client.NewListWorkspaceCustomPoolsPager(exampleWorkspaceID, &spark.CustomPoolsClientListWorkspaceCustomPoolsOptions{ContinuationToken: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		testsuite.Require().NoError(err, "Failed to advance page for example ")
		testsuite.Require().True(reflect.DeepEqual(exampleRes, nextResult.CustomPools))
		if err == nil {
			break
		}
	}

	// From example
	ctx = runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"List custom pools with continuation example"},
	})
	exampleWorkspaceID = "f089354e-8366-4e18-aea3-4cb4a3a50b48"

	exampleRes = spark.CustomPools{
		ContinuationToken: to.Ptr("LDEsMTAwMDAwLDA%3D"),
		ContinuationURI:   to.Ptr("https://api.fabric.microsoft.com/v1/workspaces/f089354e-8366-4e18-aea3-4cb4a3a50b48/spark/pools?continuationToken=LDEsMTAwMDAwLDA%3D"),
		Value: []spark.CustomPool{
			{
				Name: to.Ptr("pool1"),
				Type: to.Ptr(spark.CustomPoolTypeWorkspace),
				AutoScale: &spark.AutoScaleProperties{
					Enabled:      to.Ptr(true),
					MaxNodeCount: to.Ptr[int32](4),
					MinNodeCount: to.Ptr[int32](1),
				},
				DynamicExecutorAllocation: &spark.DynamicExecutorAllocationProperties{
					Enabled:      to.Ptr(true),
					MaxExecutors: to.Ptr[int32](2),
					MinExecutors: to.Ptr[int32](1),
				},
				ID:         to.Ptr("2367293d-b70b-4b33-97f2-161b8d04a8d7"),
				NodeFamily: to.Ptr(spark.NodeFamilyMemoryOptimized),
				NodeSize:   to.Ptr(spark.NodeSizeSmall),
			}},
	}

	testsuite.serverFactory.CustomPoolsServer.NewListWorkspaceCustomPoolsPager = func(workspaceID string, options *spark.CustomPoolsClientListWorkspaceCustomPoolsOptions) (resp azfake.PagerResponder[spark.CustomPoolsClientListWorkspaceCustomPoolsResponse]) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		resp = azfake.PagerResponder[spark.CustomPoolsClientListWorkspaceCustomPoolsResponse]{}
		resp.AddPage(http.StatusOK, spark.CustomPoolsClientListWorkspaceCustomPoolsResponse{CustomPools: exampleRes}, nil)
		return
	}

	pager = client.NewListWorkspaceCustomPoolsPager(exampleWorkspaceID, &spark.CustomPoolsClientListWorkspaceCustomPoolsOptions{ContinuationToken: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		testsuite.Require().NoError(err, "Failed to advance page for example ")
		testsuite.Require().True(reflect.DeepEqual(exampleRes, nextResult.CustomPools))
		if err == nil {
			break
		}
	}
}

func (testsuite *FakeTestSuite) TestCustomPools_CreateWorkspaceCustomPool() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Create custom pool example"},
	})
	var exampleWorkspaceID string
	var exampleCreateCustomPoolRequest spark.CreateCustomPoolRequest
	exampleWorkspaceID = "f089354e-8366-4e18-aea3-4cb4a3a50b48"
	exampleCreateCustomPoolRequest = spark.CreateCustomPoolRequest{
		Name: to.Ptr("pool1"),
		AutoScale: &spark.AutoScaleProperties{
			Enabled:      to.Ptr(true),
			MaxNodeCount: to.Ptr[int32](2),
			MinNodeCount: to.Ptr[int32](1),
		},
		DynamicExecutorAllocation: &spark.DynamicExecutorAllocationProperties{
			Enabled:      to.Ptr(true),
			MaxExecutors: to.Ptr[int32](1),
			MinExecutors: to.Ptr[int32](1),
		},
		NodeFamily: to.Ptr(spark.NodeFamilyMemoryOptimized),
		NodeSize:   to.Ptr(spark.NodeSizeSmall),
	}

	testsuite.serverFactory.CustomPoolsServer.CreateWorkspaceCustomPool = func(ctx context.Context, workspaceID string, createCustomPoolRequest spark.CreateCustomPoolRequest, options *spark.CustomPoolsClientCreateWorkspaceCustomPoolOptions) (resp azfake.Responder[spark.CustomPoolsClientCreateWorkspaceCustomPoolResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().True(reflect.DeepEqual(exampleCreateCustomPoolRequest, createCustomPoolRequest))
		resp = azfake.Responder[spark.CustomPoolsClientCreateWorkspaceCustomPoolResponse]{}
		resp.SetResponse(http.StatusCreated, spark.CustomPoolsClientCreateWorkspaceCustomPoolResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewCustomPoolsClient()
	_, err = client.CreateWorkspaceCustomPool(ctx, exampleWorkspaceID, exampleCreateCustomPoolRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
}

func (testsuite *FakeTestSuite) TestCustomPools_GetWorkspaceCustomPool() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Get custom pool example"},
	})
	var exampleWorkspaceID string
	var examplePoolID string
	exampleWorkspaceID = "f089354e-8366-4e18-aea3-4cb4a3a50b48"
	examplePoolID = "2367293d-b70b-4b33-97f2-161b8d04a8d7"

	exampleRes := spark.CustomPool{
		Name: to.Ptr("pool1"),
		Type: to.Ptr(spark.CustomPoolTypeWorkspace),
		AutoScale: &spark.AutoScaleProperties{
			Enabled:      to.Ptr(true),
			MaxNodeCount: to.Ptr[int32](4),
			MinNodeCount: to.Ptr[int32](1),
		},
		DynamicExecutorAllocation: &spark.DynamicExecutorAllocationProperties{
			Enabled:      to.Ptr(true),
			MaxExecutors: to.Ptr[int32](2),
			MinExecutors: to.Ptr[int32](1),
		},
		ID:         to.Ptr("2367293d-b70b-4b33-97f2-161b8d04a8d7"),
		NodeFamily: to.Ptr(spark.NodeFamilyMemoryOptimized),
		NodeSize:   to.Ptr(spark.NodeSizeSmall),
	}

	testsuite.serverFactory.CustomPoolsServer.GetWorkspaceCustomPool = func(ctx context.Context, workspaceID string, poolID string, options *spark.CustomPoolsClientGetWorkspaceCustomPoolOptions) (resp azfake.Responder[spark.CustomPoolsClientGetWorkspaceCustomPoolResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(examplePoolID, poolID)
		resp = azfake.Responder[spark.CustomPoolsClientGetWorkspaceCustomPoolResponse]{}
		resp.SetResponse(http.StatusOK, spark.CustomPoolsClientGetWorkspaceCustomPoolResponse{CustomPool: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewCustomPoolsClient()
	res, err := client.GetWorkspaceCustomPool(ctx, exampleWorkspaceID, examplePoolID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.CustomPool))
}

func (testsuite *FakeTestSuite) TestCustomPools_DeleteWorkspaceCustomPool() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Delete custom pool example"},
	})
	var exampleWorkspaceID string
	var examplePoolID string
	exampleWorkspaceID = "f089354e-8366-4e18-aea3-4cb4a3a50b48"
	examplePoolID = "2367293d-b70b-4b33-97f2-161b8d04a8d7"

	testsuite.serverFactory.CustomPoolsServer.DeleteWorkspaceCustomPool = func(ctx context.Context, workspaceID string, poolID string, options *spark.CustomPoolsClientDeleteWorkspaceCustomPoolOptions) (resp azfake.Responder[spark.CustomPoolsClientDeleteWorkspaceCustomPoolResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(examplePoolID, poolID)
		resp = azfake.Responder[spark.CustomPoolsClientDeleteWorkspaceCustomPoolResponse]{}
		resp.SetResponse(http.StatusOK, spark.CustomPoolsClientDeleteWorkspaceCustomPoolResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewCustomPoolsClient()
	_, err = client.DeleteWorkspaceCustomPool(ctx, exampleWorkspaceID, examplePoolID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
}

func (testsuite *FakeTestSuite) TestCustomPools_UpdateWorkspaceCustomPool() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Update custom pool example"},
	})
	var exampleWorkspaceID string
	var examplePoolID string
	var exampleUpdateCustomPoolRequest spark.UpdateCustomPoolRequest
	exampleWorkspaceID = "f089354e-8366-4e18-aea3-4cb4a3a50b48"
	examplePoolID = "2367293d-b70b-4b33-97f2-161b8d04a8d7"
	exampleUpdateCustomPoolRequest = spark.UpdateCustomPoolRequest{
		Name: to.Ptr("pool1"),
		AutoScale: &spark.AutoScaleProperties{
			Enabled:      to.Ptr(true),
			MaxNodeCount: to.Ptr[int32](2),
			MinNodeCount: to.Ptr[int32](1),
		},
		DynamicExecutorAllocation: &spark.DynamicExecutorAllocationProperties{
			Enabled:      to.Ptr(true),
			MaxExecutors: to.Ptr[int32](1),
			MinExecutors: to.Ptr[int32](1),
		},
		NodeFamily: to.Ptr(spark.NodeFamilyMemoryOptimized),
		NodeSize:   to.Ptr(spark.NodeSizeSmall),
	}

	exampleRes := spark.CustomPool{
		Name: to.Ptr("pool1"),
		Type: to.Ptr(spark.CustomPoolTypeWorkspace),
		AutoScale: &spark.AutoScaleProperties{
			Enabled:      to.Ptr(true),
			MaxNodeCount: to.Ptr[int32](2),
			MinNodeCount: to.Ptr[int32](1),
		},
		DynamicExecutorAllocation: &spark.DynamicExecutorAllocationProperties{
			Enabled:      to.Ptr(true),
			MaxExecutors: to.Ptr[int32](1),
			MinExecutors: to.Ptr[int32](1),
		},
		ID:         to.Ptr("2367293d-b70b-4b33-97f2-161b8d04a8d7"),
		NodeFamily: to.Ptr(spark.NodeFamilyMemoryOptimized),
		NodeSize:   to.Ptr(spark.NodeSizeSmall),
	}

	testsuite.serverFactory.CustomPoolsServer.UpdateWorkspaceCustomPool = func(ctx context.Context, workspaceID string, poolID string, updateCustomPoolRequest spark.UpdateCustomPoolRequest, options *spark.CustomPoolsClientUpdateWorkspaceCustomPoolOptions) (resp azfake.Responder[spark.CustomPoolsClientUpdateWorkspaceCustomPoolResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(examplePoolID, poolID)
		testsuite.Require().True(reflect.DeepEqual(exampleUpdateCustomPoolRequest, updateCustomPoolRequest))
		resp = azfake.Responder[spark.CustomPoolsClientUpdateWorkspaceCustomPoolResponse]{}
		resp.SetResponse(http.StatusOK, spark.CustomPoolsClientUpdateWorkspaceCustomPoolResponse{CustomPool: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewCustomPoolsClient()
	res, err := client.UpdateWorkspaceCustomPool(ctx, exampleWorkspaceID, examplePoolID, exampleUpdateCustomPoolRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.CustomPool))
}
