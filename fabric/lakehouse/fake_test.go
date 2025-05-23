// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package lakehouse_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"

	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"

	"reflect"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/microsoft/fabric-sdk-go/fabric/lakehouse"
	"github.com/microsoft/fabric-sdk-go/fabric/lakehouse/fake"
)

var err error

type FakeTestSuite struct {
	suite.Suite

	ctx  context.Context
	cred azcore.TokenCredential

	serverFactory *fake.ServerFactory
	clientFactory *lakehouse.ClientFactory
}

func (testsuite *FakeTestSuite) SetupSuite() {
	testsuite.ctx = context.Background()
	testsuite.cred = &azfake.TokenCredential{}

	testsuite.serverFactory = &fake.ServerFactory{}
	testsuite.clientFactory, err = lakehouse.NewClientFactory(testsuite.cred, nil, &azcore.ClientOptions{
		Transport: fake.NewServerFactoryTransport(testsuite.serverFactory),
	})
	testsuite.Require().NoError(err, "Failed to create client factory")
}

func TestFakeTest(t *testing.T) {
	suite.Run(t, new(FakeTestSuite))
}

func (testsuite *FakeTestSuite) TestItems_ListLakehouses() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"List lakehouses in workspace example"},
	})
	var exampleWorkspaceID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"

	exampleRes := lakehouse.Lakehouses{
		Value: []lakehouse.Lakehouse{
			{
				Type:        to.Ptr(lakehouse.ItemTypeLakehouse),
				Description: to.Ptr("A lakehouse description."),
				DisplayName: to.Ptr("Lakehouse_1"),
				ID:          to.Ptr("3546052c-ae64-4526-b1a8-52af7761426f"),
				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
				Properties: &lakehouse.Properties{
					OneLakeFilesPath:  to.Ptr("https://onelake.dfs.fabric.microsoft.com/2382cdf5-d577-44d0-a1fc-42184f29a7eb/e5fb215b-1934-413e-b33a-debaf844afde/Files"),
					OneLakeTablesPath: to.Ptr("https://onelake.dfs.fabric.microsoft.com/2382cdf5-d577-44d0-a1fc-42184f29a7eb/e5fb215b-1934-413e-b33a-debaf844afde/Tables"),
					SQLEndpointProperties: &lakehouse.SQLEndpointProperties{
						ConnectionString:   to.Ptr("qvrmbuxie7we7glrekxgy6npqu-6xgyei3x2xiejip4iime6knh5m.datawarehouse.fabric.microsoft.com"),
						ID:                 to.Ptr("37dc8a41-dea9-465d-b528-3e95043b2356"),
						ProvisioningStatus: to.Ptr(lakehouse.SQLEndpointProvisioningStatusSuccess),
					},
				},
			},
			{
				Type:        to.Ptr(lakehouse.ItemTypeLakehouse),
				Description: to.Ptr("A lakehouse description."),
				DisplayName: to.Ptr("Lakehouse_2"),
				ID:          to.Ptr("a8a1bffa-7eea-49dc-a1d2-6281c1d031f1"),
				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
				Properties: &lakehouse.Properties{
					OneLakeFilesPath:  to.Ptr("https://onelake.dfs.fabric.microsoft.com/fc5d0537-1b22-4de1-a5e9-9b8bb58ed1e1/6dc325f6-46f6-4a2a-930b-10b96a463566/Files"),
					OneLakeTablesPath: to.Ptr("https://onelake.dfs.fabric.microsoft.com/fc5d0537-1b22-4de1-a5e9-9b8bb58ed1e1/6dc325f6-46f6-4a2a-930b-10b96a463566/Tables"),
					SQLEndpointProperties: &lakehouse.SQLEndpointProperties{
						ConnectionString:   to.Ptr("qvrmbuxie7we7glrekxgy6npqu-6xgyei3x2xiejip4iime6knh5m.datawarehouse.fabric.microsoft.com"),
						ID:                 to.Ptr("37dc8a41-dea9-465d-b528-3e95043b2356"),
						ProvisioningStatus: to.Ptr(lakehouse.SQLEndpointProvisioningStatusSuccess),
					},
				},
			}},
	}

	testsuite.serverFactory.ItemsServer.NewListLakehousesPager = func(workspaceID string, options *lakehouse.ItemsClientListLakehousesOptions) (resp azfake.PagerResponder[lakehouse.ItemsClientListLakehousesResponse]) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		resp = azfake.PagerResponder[lakehouse.ItemsClientListLakehousesResponse]{}
		resp.AddPage(http.StatusOK, lakehouse.ItemsClientListLakehousesResponse{Lakehouses: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	pager := client.NewListLakehousesPager(exampleWorkspaceID, &lakehouse.ItemsClientListLakehousesOptions{ContinuationToken: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		testsuite.Require().NoError(err, "Failed to advance page for example ")
		testsuite.Require().True(reflect.DeepEqual(exampleRes, nextResult.Lakehouses))
		if err == nil {
			break
		}
	}
}

func (testsuite *FakeTestSuite) TestItems_CreateLakehouse() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Create a lakehouse example"},
	})
	var exampleWorkspaceID string
	var exampleCreateLakehouseRequest lakehouse.CreateLakehouseRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleCreateLakehouseRequest = lakehouse.CreateLakehouseRequest{
		Description: to.Ptr("A lakehouse description"),
		DisplayName: to.Ptr("Lakehouse_1"),
	}

	testsuite.serverFactory.ItemsServer.BeginCreateLakehouse = func(ctx context.Context, workspaceID string, createLakehouseRequest lakehouse.CreateLakehouseRequest, options *lakehouse.ItemsClientBeginCreateLakehouseOptions) (resp azfake.PollerResponder[lakehouse.ItemsClientCreateLakehouseResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().True(reflect.DeepEqual(exampleCreateLakehouseRequest, createLakehouseRequest))
		resp = azfake.PollerResponder[lakehouse.ItemsClientCreateLakehouseResponse]{}
		resp.SetTerminalResponse(http.StatusCreated, lakehouse.ItemsClientCreateLakehouseResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	poller, err := client.BeginCreateLakehouse(ctx, exampleWorkspaceID, exampleCreateLakehouseRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")

	// From example
	ctx = runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Create a lakehouse with schema example"},
	})
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleCreateLakehouseRequest = lakehouse.CreateLakehouseRequest{
		Description: to.Ptr("A schema enabled lakehouse."),
		CreationPayload: &lakehouse.CreationPayload{
			EnableSchemas: to.Ptr(true),
		},
		DisplayName: to.Ptr("Lakehouse_created_with_schema"),
	}

	testsuite.serverFactory.ItemsServer.BeginCreateLakehouse = func(ctx context.Context, workspaceID string, createLakehouseRequest lakehouse.CreateLakehouseRequest, options *lakehouse.ItemsClientBeginCreateLakehouseOptions) (resp azfake.PollerResponder[lakehouse.ItemsClientCreateLakehouseResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().True(reflect.DeepEqual(exampleCreateLakehouseRequest, createLakehouseRequest))
		resp = azfake.PollerResponder[lakehouse.ItemsClientCreateLakehouseResponse]{}
		resp.SetTerminalResponse(http.StatusCreated, lakehouse.ItemsClientCreateLakehouseResponse{}, nil)
		return
	}

	poller, err = client.BeginCreateLakehouse(ctx, exampleWorkspaceID, exampleCreateLakehouseRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")
}

func (testsuite *FakeTestSuite) TestItems_GetLakehouse() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Get a lakehouse example"},
	})
	var exampleWorkspaceID string
	var exampleLakehouseID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleLakehouseID = "5b218778-e7a5-4d73-8187-f10824047715"

	exampleRes := lakehouse.Lakehouse{
		Type:        to.Ptr(lakehouse.ItemTypeLakehouse),
		Description: to.Ptr("A lakehouse description."),
		DisplayName: to.Ptr("Lakehouse_1"),
		ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		Properties: &lakehouse.Properties{
			OneLakeFilesPath:  to.Ptr("https://onelake.dfs.fabric.microsoft.com/2382cdf5-d577-44d0-a1fc-42184f29a7eb/e5fb215b-1934-413e-b33a-debaf844afde/Files"),
			OneLakeTablesPath: to.Ptr("https://onelake.dfs.fabric.microsoft.com/2382cdf5-d577-44d0-a1fc-42184f29a7eb/e5fb215b-1934-413e-b33a-debaf844afde/Tables"),
			SQLEndpointProperties: &lakehouse.SQLEndpointProperties{
				ConnectionString:   to.Ptr("qvrmbuxie7we7glrekxgy6npqu-6xgyei3x2xiejip4iime6knh5m.datawarehouse.fabric.microsoft.com"),
				ID:                 to.Ptr("37dc8a41-dea9-465d-b528-3e95043b2356"),
				ProvisioningStatus: to.Ptr(lakehouse.SQLEndpointProvisioningStatusSuccess),
			},
		},
	}

	testsuite.serverFactory.ItemsServer.GetLakehouse = func(ctx context.Context, workspaceID string, lakehouseID string, options *lakehouse.ItemsClientGetLakehouseOptions) (resp azfake.Responder[lakehouse.ItemsClientGetLakehouseResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleLakehouseID, lakehouseID)
		resp = azfake.Responder[lakehouse.ItemsClientGetLakehouseResponse]{}
		resp.SetResponse(http.StatusOK, lakehouse.ItemsClientGetLakehouseResponse{Lakehouse: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	res, err := client.GetLakehouse(ctx, exampleWorkspaceID, exampleLakehouseID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.Lakehouse))

	// From example
	ctx = runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Get a lakehouse with schema example"},
	})
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleLakehouseID = "5b218778-e7a5-4d73-8187-f10824047715"

	exampleRes = lakehouse.Lakehouse{
		Type:        to.Ptr(lakehouse.ItemTypeLakehouse),
		Description: to.Ptr("A schema enabled lakehouse."),
		DisplayName: to.Ptr("Lakehouse_created_with_schema"),
		ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		Properties: &lakehouse.Properties{
			DefaultSchema:     to.Ptr("dbo"),
			OneLakeFilesPath:  to.Ptr("https://onelake.dfs.fabric.microsoft.com/2382cdf5-d577-44d0-a1fc-42184f29a7eb/e5fb215b-1934-413e-b33a-debaf844afde/Files"),
			OneLakeTablesPath: to.Ptr("https://onelake.dfs.fabric.microsoft.com/2382cdf5-d577-44d0-a1fc-42184f29a7eb/e5fb215b-1934-413e-b33a-debaf844afde/Tables"),
			SQLEndpointProperties: &lakehouse.SQLEndpointProperties{
				ConnectionString:   to.Ptr("qvrmbuxie7we7glrekxgy6npqu-6xgyei3x2xiejip4iime6knh5m.datawarehouse.fabric.microsoft.com"),
				ID:                 to.Ptr("37dc8a41-dea9-465d-b528-3e95043b2356"),
				ProvisioningStatus: to.Ptr(lakehouse.SQLEndpointProvisioningStatusSuccess),
			},
		},
	}

	testsuite.serverFactory.ItemsServer.GetLakehouse = func(ctx context.Context, workspaceID string, lakehouseID string, options *lakehouse.ItemsClientGetLakehouseOptions) (resp azfake.Responder[lakehouse.ItemsClientGetLakehouseResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleLakehouseID, lakehouseID)
		resp = azfake.Responder[lakehouse.ItemsClientGetLakehouseResponse]{}
		resp.SetResponse(http.StatusOK, lakehouse.ItemsClientGetLakehouseResponse{Lakehouse: exampleRes}, nil)
		return
	}

	res, err = client.GetLakehouse(ctx, exampleWorkspaceID, exampleLakehouseID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.Lakehouse))
}

func (testsuite *FakeTestSuite) TestItems_UpdateLakehouse() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Update a lakehouse example"},
	})
	var exampleWorkspaceID string
	var exampleLakehouseID string
	var exampleUpdateLakehouseRequest lakehouse.UpdateLakehouseRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleLakehouseID = "5b218778-e7a5-4d73-8187-f10824047715"
	exampleUpdateLakehouseRequest = lakehouse.UpdateLakehouseRequest{
		Description: to.Ptr("A new description for lakehouse."),
		DisplayName: to.Ptr("Lakehouse_New_Name"),
	}

	exampleRes := lakehouse.Lakehouse{
		Type:        to.Ptr(lakehouse.ItemTypeLakehouse),
		Description: to.Ptr("A new description for lakehouse."),
		DisplayName: to.Ptr("Lakehouse_New_Name"),
		ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	}

	testsuite.serverFactory.ItemsServer.UpdateLakehouse = func(ctx context.Context, workspaceID string, lakehouseID string, updateLakehouseRequest lakehouse.UpdateLakehouseRequest, options *lakehouse.ItemsClientUpdateLakehouseOptions) (resp azfake.Responder[lakehouse.ItemsClientUpdateLakehouseResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleLakehouseID, lakehouseID)
		testsuite.Require().True(reflect.DeepEqual(exampleUpdateLakehouseRequest, updateLakehouseRequest))
		resp = azfake.Responder[lakehouse.ItemsClientUpdateLakehouseResponse]{}
		resp.SetResponse(http.StatusOK, lakehouse.ItemsClientUpdateLakehouseResponse{Lakehouse: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	res, err := client.UpdateLakehouse(ctx, exampleWorkspaceID, exampleLakehouseID, exampleUpdateLakehouseRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.Lakehouse))
}

func (testsuite *FakeTestSuite) TestItems_DeleteLakehouse() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Delete a lakehouse example"},
	})
	var exampleWorkspaceID string
	var exampleLakehouseID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleLakehouseID = "5b218778-e7a5-4d73-8187-f10824047715"

	testsuite.serverFactory.ItemsServer.DeleteLakehouse = func(ctx context.Context, workspaceID string, lakehouseID string, options *lakehouse.ItemsClientDeleteLakehouseOptions) (resp azfake.Responder[lakehouse.ItemsClientDeleteLakehouseResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleLakehouseID, lakehouseID)
		resp = azfake.Responder[lakehouse.ItemsClientDeleteLakehouseResponse]{}
		resp.SetResponse(http.StatusOK, lakehouse.ItemsClientDeleteLakehouseResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	_, err = client.DeleteLakehouse(ctx, exampleWorkspaceID, exampleLakehouseID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
}

func (testsuite *FakeTestSuite) TestTables_ListTables() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"List tables example"},
	})
	var exampleWorkspaceID string
	var exampleLakehouseID string
	exampleWorkspaceID = "f089354e-8366-4e18-aea3-4cb4a3a50b48"
	exampleLakehouseID = "41ce06d1-d81b-4ea0-bc6d-2ce3dd2f8e87"

	exampleRes := lakehouse.Tables{
		ContinuationToken: to.Ptr("+RID:~HTsuAOseYicH-GcAAAAAAA==#RT:1#TRC:1#ISV:2#IEO:65567#QCF:8#FPC:AgKfAZ8BnwEEAAe8eoA="),
		ContinuationURI:   to.Ptr("https://api.fabric.microsoft.com/v1/workspaces/f089354e-8366-4e18-aea3-4cb4a3a50b48/lakehouses/41ce06d1-d81b-4ea0-bc6d-2ce3dd2f8e87/tables?continuationToken=%2BRID%3A~HTsuAOseYicH-GcAAAAAAA%3D%3D%23RT%3A1%23TRC%3A1%23ISV%3A2%23IEO%3A65567%23QCF%3A8%23FPC%3AAgKfAZ8BnwEEAAe8eoA%3D"),
		Data: []lakehouse.Table{
			{
				Name:     to.Ptr("Table1"),
				Type:     to.Ptr(lakehouse.TableTypeManaged),
				Format:   to.Ptr("Delta"),
				Location: to.Ptr("abfss://f089354e-8366-4e18-aea3-4cb4a3a50b48@onelake.dfs.fabric.microsoft.com/41ce06d1-d81b-4ea0-bc6d-2ce3dd2f8e87/Tables/Table1"),
			}},
	}

	testsuite.serverFactory.TablesServer.NewListTablesPager = func(workspaceID string, lakehouseID string, options *lakehouse.TablesClientListTablesOptions) (resp azfake.PagerResponder[lakehouse.TablesClientListTablesResponse]) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleLakehouseID, lakehouseID)
		resp = azfake.PagerResponder[lakehouse.TablesClientListTablesResponse]{}
		resp.AddPage(http.StatusOK, lakehouse.TablesClientListTablesResponse{Tables: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewTablesClient()
	pager := client.NewListTablesPager(exampleWorkspaceID, exampleLakehouseID, &lakehouse.TablesClientListTablesOptions{MaxResults: to.Ptr[int32](10),
		ContinuationToken: to.Ptr(""),
	})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		testsuite.Require().NoError(err, "Failed to advance page for example ")
		testsuite.Require().True(reflect.DeepEqual(exampleRes, nextResult.Tables))
		if err == nil {
			break
		}
	}
}

func (testsuite *FakeTestSuite) TestBackgroundJobs_RunOnDemandTableMaintenance() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Run table maintenance with optimize Z-Order and vacuum enabled for schema enabled lakehouse."},
	})
	var exampleWorkspaceID string
	var exampleLakehouseID string
	var exampleJobType string
	var exampleRunOnDemandTableMaintenanceRequest lakehouse.RunOnDemandTableMaintenanceRequest
	exampleWorkspaceID = "4b218778-e7a5-4d73-8187-f10824047715"
	exampleLakehouseID = "431e8d7b-4a95-4c02-8ccd-6faef5ba1bd7"
	exampleJobType = "TableMaintenance"
	exampleRunOnDemandTableMaintenanceRequest = lakehouse.RunOnDemandTableMaintenanceRequest{
		ExecutionData: &lakehouse.TableMaintenanceExecutionData{
			OptimizeSettings: &lakehouse.OptimizeSettings{
				VOrder: to.Ptr(true),
				ZOrderBy: []string{
					"tipAmount"},
			},
			SchemaName: to.Ptr("dbo"),
			TableName:  to.Ptr("table1"),
			VacuumSettings: &lakehouse.VacuumSettings{
				RetentionPeriod: to.Ptr("7:01:00:00"),
			},
		},
	}

	testsuite.serverFactory.BackgroundJobsServer.RunOnDemandTableMaintenance = func(ctx context.Context, workspaceID string, lakehouseID string, jobType string, runOnDemandTableMaintenanceRequest lakehouse.RunOnDemandTableMaintenanceRequest, options *lakehouse.BackgroundJobsClientRunOnDemandTableMaintenanceOptions) (resp azfake.Responder[lakehouse.BackgroundJobsClientRunOnDemandTableMaintenanceResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleLakehouseID, lakehouseID)
		testsuite.Require().Equal(exampleJobType, jobType)
		testsuite.Require().True(reflect.DeepEqual(exampleRunOnDemandTableMaintenanceRequest, runOnDemandTableMaintenanceRequest))
		resp = azfake.Responder[lakehouse.BackgroundJobsClientRunOnDemandTableMaintenanceResponse]{}
		resp.SetResponse(http.StatusAccepted, lakehouse.BackgroundJobsClientRunOnDemandTableMaintenanceResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewBackgroundJobsClient()
	_, err = client.RunOnDemandTableMaintenance(ctx, exampleWorkspaceID, exampleLakehouseID, exampleJobType, exampleRunOnDemandTableMaintenanceRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")

	// From example
	ctx = runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Run table maintenance with optimize Z-Order and vacuum enabled."},
	})
	exampleWorkspaceID = "4b218778-e7a5-4d73-8187-f10824047715"
	exampleLakehouseID = "431e8d7b-4a95-4c02-8ccd-6faef5ba1bd7"
	exampleJobType = "TableMaintenance"
	exampleRunOnDemandTableMaintenanceRequest = lakehouse.RunOnDemandTableMaintenanceRequest{
		ExecutionData: &lakehouse.TableMaintenanceExecutionData{
			OptimizeSettings: &lakehouse.OptimizeSettings{
				VOrder: to.Ptr(true),
				ZOrderBy: []string{
					"tipAmount"},
			},
			TableName: to.Ptr("table1"),
			VacuumSettings: &lakehouse.VacuumSettings{
				RetentionPeriod: to.Ptr("7:01:00:00"),
			},
		},
	}

	testsuite.serverFactory.BackgroundJobsServer.RunOnDemandTableMaintenance = func(ctx context.Context, workspaceID string, lakehouseID string, jobType string, runOnDemandTableMaintenanceRequest lakehouse.RunOnDemandTableMaintenanceRequest, options *lakehouse.BackgroundJobsClientRunOnDemandTableMaintenanceOptions) (resp azfake.Responder[lakehouse.BackgroundJobsClientRunOnDemandTableMaintenanceResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleLakehouseID, lakehouseID)
		testsuite.Require().Equal(exampleJobType, jobType)
		testsuite.Require().True(reflect.DeepEqual(exampleRunOnDemandTableMaintenanceRequest, runOnDemandTableMaintenanceRequest))
		resp = azfake.Responder[lakehouse.BackgroundJobsClientRunOnDemandTableMaintenanceResponse]{}
		resp.SetResponse(http.StatusAccepted, lakehouse.BackgroundJobsClientRunOnDemandTableMaintenanceResponse{}, nil)
		return
	}

	_, err = client.RunOnDemandTableMaintenance(ctx, exampleWorkspaceID, exampleLakehouseID, exampleJobType, exampleRunOnDemandTableMaintenanceRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")

	// From example
	ctx = runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Run table maintenance with optimize enabled and vacuum disabled."},
	})
	exampleWorkspaceID = "4b218778-e7a5-4d73-8187-f10824047715"
	exampleLakehouseID = "431e8d7b-4a95-4c02-8ccd-6faef5ba1bd7"
	exampleJobType = "TableMaintenance"
	exampleRunOnDemandTableMaintenanceRequest = lakehouse.RunOnDemandTableMaintenanceRequest{
		ExecutionData: &lakehouse.TableMaintenanceExecutionData{
			OptimizeSettings: &lakehouse.OptimizeSettings{},
			TableName:        to.Ptr("table1"),
		},
	}

	testsuite.serverFactory.BackgroundJobsServer.RunOnDemandTableMaintenance = func(ctx context.Context, workspaceID string, lakehouseID string, jobType string, runOnDemandTableMaintenanceRequest lakehouse.RunOnDemandTableMaintenanceRequest, options *lakehouse.BackgroundJobsClientRunOnDemandTableMaintenanceOptions) (resp azfake.Responder[lakehouse.BackgroundJobsClientRunOnDemandTableMaintenanceResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleLakehouseID, lakehouseID)
		testsuite.Require().Equal(exampleJobType, jobType)
		testsuite.Require().True(reflect.DeepEqual(exampleRunOnDemandTableMaintenanceRequest, runOnDemandTableMaintenanceRequest))
		resp = azfake.Responder[lakehouse.BackgroundJobsClientRunOnDemandTableMaintenanceResponse]{}
		resp.SetResponse(http.StatusAccepted, lakehouse.BackgroundJobsClientRunOnDemandTableMaintenanceResponse{}, nil)
		return
	}

	_, err = client.RunOnDemandTableMaintenance(ctx, exampleWorkspaceID, exampleLakehouseID, exampleJobType, exampleRunOnDemandTableMaintenanceRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
}

func (testsuite *FakeTestSuite) TestLivySessions_ListLivySessions() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"List all livy sessions example"},
	})
	var exampleWorkspaceID string
	var exampleLakehouseID string
	exampleWorkspaceID = "f8113ba8-dd81-443e-811a-b385340f3f05"
	exampleLakehouseID = "8cee7699-2e81-4121-9a53-cc9025046193"

	exampleRes := lakehouse.LivySessions{
		Value: []lakehouse.LivySession{
			{
				AttemptNumber:      to.Ptr[int32](1),
				CancellationReason: to.Ptr("User cancelled the Spark batch"),
				CapacityID:         to.Ptr("3c0cd366-dc28-4b6d-a525-4d415a8666e7"),
				CreatorItem: &lakehouse.ItemReferenceByID{
					ReferenceType: to.Ptr(lakehouse.ItemReferenceTypeByID),
					ItemID:        to.Ptr("8cee7699-2e81-4121-9a53-cc9025046193"),
					WorkspaceID:   to.Ptr("f8113ba8-dd81-443e-811a-b385340f3f05"),
				},
				EndDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-31T15:37:30.000Z"); return t }()),
				Item: &lakehouse.ItemReferenceByID{
					ReferenceType: to.Ptr(lakehouse.ItemReferenceTypeByID),
					ItemID:        to.Ptr("8cee7699-2e81-4121-9a53-cc9025046193"),
					WorkspaceID:   to.Ptr("f8113ba8-dd81-443e-811a-b385340f3f05"),
				},
				ItemName:                   to.Ptr("lh_itemName"),
				ItemType:                   to.Ptr(lakehouse.ItemTypeLakehouse),
				JobInstanceID:              to.Ptr("c2baabbd-5327-430c-87a6-ff4f98285601"),
				JobType:                    to.Ptr(lakehouse.JobTypeSparkBatch),
				LivyID:                     to.Ptr("9611f500-bf44-42e0-a0de-78dacb374398"),
				LivyName:                   to.Ptr("random_test_name_app"),
				LivySessionItemResourceURI: to.Ptr(""),
				MaxNumberOfAttempts:        to.Ptr[int32](1),
				OperationName:              to.Ptr("Batch Livy Run"),
				Origin:                     to.Ptr(lakehouse.OriginSubmittedJob),
				QueuedDuration: &lakehouse.Duration{
					TimeUnit: to.Ptr(lakehouse.TimeUnitSeconds),
					Value:    to.Ptr[float32](1),
				},
				RunningDuration: &lakehouse.Duration{
					TimeUnit: to.Ptr(lakehouse.TimeUnitSeconds),
					Value:    to.Ptr[float32](180),
				},
				RuntimeVersion:     to.Ptr("1.3"),
				SparkApplicationID: to.Ptr("application_1730933685452_0001"),
				StartDateTime:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-31T15:34:11.000Z"); return t }()),
				State:              to.Ptr(lakehouse.StateCancelled),
				SubmittedDateTime:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-31T15:32:03.000Z"); return t }()),
				Submitter: &lakehouse.Principal{
					Type: to.Ptr(lakehouse.PrincipalTypeUser),
					ID:   to.Ptr("6f23a8a6-d954-4550-b91a-4df73ccd0311"),
				},
				TotalDuration: &lakehouse.Duration{
					TimeUnit: to.Ptr(lakehouse.TimeUnitSeconds),
					Value:    to.Ptr[float32](360),
				},
			}},
	}

	testsuite.serverFactory.LivySessionsServer.NewListLivySessionsPager = func(workspaceID string, lakehouseID string, options *lakehouse.LivySessionsClientListLivySessionsOptions) (resp azfake.PagerResponder[lakehouse.LivySessionsClientListLivySessionsResponse]) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleLakehouseID, lakehouseID)
		resp = azfake.PagerResponder[lakehouse.LivySessionsClientListLivySessionsResponse]{}
		resp.AddPage(http.StatusOK, lakehouse.LivySessionsClientListLivySessionsResponse{LivySessions: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewLivySessionsClient()
	pager := client.NewListLivySessionsPager(exampleWorkspaceID, exampleLakehouseID, &lakehouse.LivySessionsClientListLivySessionsOptions{ContinuationToken: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		testsuite.Require().NoError(err, "Failed to advance page for example ")
		testsuite.Require().True(reflect.DeepEqual(exampleRes, nextResult.LivySessions))
		if err == nil {
			break
		}
	}
}

func (testsuite *FakeTestSuite) TestLivySessions_GetLivySession() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Get a livy session example"},
	})
	var exampleWorkspaceID string
	var exampleLakehouseID string
	var exampleLivyID string
	exampleWorkspaceID = "f8113ba8-dd81-443e-811a-b385340f3f05"
	exampleLakehouseID = "8cee7699-2e81-4121-9a53-cc9025046193"
	exampleLivyID = "9611f500-bf44-42e0-a0de-78dacb374398"

	exampleRes := lakehouse.LivySession{
		AttemptNumber:      to.Ptr[int32](1),
		CancellationReason: to.Ptr("User cancelled the Spark batch"),
		CapacityID:         to.Ptr("3c0cd366-dc28-4b6d-a525-4d415a8666e7"),
		CreatorItem: &lakehouse.ItemReferenceByID{
			ReferenceType: to.Ptr(lakehouse.ItemReferenceTypeByID),
			ItemID:        to.Ptr("8cee7699-2e81-4121-9a53-cc9025046193"),
			WorkspaceID:   to.Ptr("f8113ba8-dd81-443e-811a-b385340f3f05"),
		},
		EndDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-31T15:37:30.000Z"); return t }()),
		Item: &lakehouse.ItemReferenceByID{
			ReferenceType: to.Ptr(lakehouse.ItemReferenceTypeByID),
			ItemID:        to.Ptr("8cee7699-2e81-4121-9a53-cc9025046193"),
			WorkspaceID:   to.Ptr("f8113ba8-dd81-443e-811a-b385340f3f05"),
		},
		ItemName:                   to.Ptr("lh_itemName"),
		ItemType:                   to.Ptr(lakehouse.ItemTypeLakehouse),
		JobInstanceID:              to.Ptr("c2baabbd-5327-430c-87a6-ff4f98285601"),
		JobType:                    to.Ptr(lakehouse.JobTypeSparkBatch),
		LivyID:                     to.Ptr("9611f500-bf44-42e0-a0de-78dacb374398"),
		LivyName:                   to.Ptr("random_test_name_app"),
		LivySessionItemResourceURI: to.Ptr(""),
		MaxNumberOfAttempts:        to.Ptr[int32](1),
		OperationName:              to.Ptr("Batch Livy Run"),
		Origin:                     to.Ptr(lakehouse.OriginSubmittedJob),
		QueuedDuration: &lakehouse.Duration{
			TimeUnit: to.Ptr(lakehouse.TimeUnitSeconds),
			Value:    to.Ptr[float32](1),
		},
		RunningDuration: &lakehouse.Duration{
			TimeUnit: to.Ptr(lakehouse.TimeUnitSeconds),
			Value:    to.Ptr[float32](180),
		},
		RuntimeVersion:     to.Ptr("1.3"),
		SparkApplicationID: to.Ptr("application_1730933685452_0001"),
		StartDateTime:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-31T15:34:11.000Z"); return t }()),
		State:              to.Ptr(lakehouse.StateCancelled),
		SubmittedDateTime:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-31T15:32:03.000Z"); return t }()),
		Submitter: &lakehouse.Principal{
			Type: to.Ptr(lakehouse.PrincipalTypeUser),
			ID:   to.Ptr("6f23a8a6-d954-4550-b91a-4df73ccd0311"),
		},
		TotalDuration: &lakehouse.Duration{
			TimeUnit: to.Ptr(lakehouse.TimeUnitSeconds),
			Value:    to.Ptr[float32](360),
		},
	}

	testsuite.serverFactory.LivySessionsServer.GetLivySession = func(ctx context.Context, workspaceID string, lakehouseID string, livyID string, options *lakehouse.LivySessionsClientGetLivySessionOptions) (resp azfake.Responder[lakehouse.LivySessionsClientGetLivySessionResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleLakehouseID, lakehouseID)
		testsuite.Require().Equal(exampleLivyID, livyID)
		resp = azfake.Responder[lakehouse.LivySessionsClientGetLivySessionResponse]{}
		resp.SetResponse(http.StatusOK, lakehouse.LivySessionsClientGetLivySessionResponse{LivySession: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewLivySessionsClient()
	res, err := client.GetLivySession(ctx, exampleWorkspaceID, exampleLakehouseID, exampleLivyID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.LivySession))
}
