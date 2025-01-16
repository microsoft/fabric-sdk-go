// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package mirroreddatabase_test

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

	"github.com/microsoft/fabric-sdk-go/fabric/mirroreddatabase"
	"github.com/microsoft/fabric-sdk-go/fabric/mirroreddatabase/fake"
)

var err error

type FakeTestSuite struct {
	suite.Suite

	ctx  context.Context
	cred azcore.TokenCredential

	serverFactory *fake.ServerFactory
	clientFactory *mirroreddatabase.ClientFactory
}

func (testsuite *FakeTestSuite) SetupSuite() {
	testsuite.ctx = context.Background()
	testsuite.cred = &azfake.TokenCredential{}

	testsuite.serverFactory = &fake.ServerFactory{}
	testsuite.clientFactory, err = mirroreddatabase.NewClientFactory(testsuite.cred, nil, &azcore.ClientOptions{
		Transport: fake.NewServerFactoryTransport(testsuite.serverFactory),
	})
	testsuite.Require().NoError(err, "Failed to create client factory")
}

func TestFakeTest(t *testing.T) {
	suite.Run(t, new(FakeTestSuite))
}

func (testsuite *FakeTestSuite) TestItems_ListMirroredDatabases() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"List mirrored databases in workspace example"},
	})
	var exampleWorkspaceID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"

	exampleRes := mirroreddatabase.MirroredDatabases{
		Value: []mirroreddatabase.MirroredDatabase{
			{
				Type:        to.Ptr(mirroreddatabase.ItemTypeMirroredDatabase),
				Description: to.Ptr("A mirrored database description."),
				DisplayName: to.Ptr("Mirrored database 1"),
				ID:          to.Ptr("3546052c-ae64-4526-b1a8-52af7761426f"),
				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
				Properties: &mirroreddatabase.Properties{
					DefaultSchema:     to.Ptr("dbo"),
					OneLakeTablesPath: to.Ptr("https://onelake.dfs.fabric.microsoft.com/7ba2c2f8-750a-47d8-9e95-99585107e23a/c3889489-6074-490e-90ed-e5a9e46142c6/Tables"),
					SQLEndpointProperties: &mirroreddatabase.SQLEndpointProperties{
						ConnectionString:   to.Ptr("x6eps4xrq2xudenlfv6naeo3i4-7dbke6ykoxmephuvtfmfcb7chi.msit-datawarehouse.fabric.microsoft.com"),
						ID:                 to.Ptr("84c9ad4a-ba9e-42d9-9025-b40f8e38c025"),
						ProvisioningStatus: to.Ptr(mirroreddatabase.SQLEndpointProvisioningStatusSuccess),
					},
				},
			}},
	}

	testsuite.serverFactory.ItemsServer.NewListMirroredDatabasesPager = func(workspaceID string, options *mirroreddatabase.ItemsClientListMirroredDatabasesOptions) (resp azfake.PagerResponder[mirroreddatabase.ItemsClientListMirroredDatabasesResponse]) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		resp = azfake.PagerResponder[mirroreddatabase.ItemsClientListMirroredDatabasesResponse]{}
		resp.AddPage(http.StatusOK, mirroreddatabase.ItemsClientListMirroredDatabasesResponse{MirroredDatabases: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	pager := client.NewListMirroredDatabasesPager(exampleWorkspaceID, &mirroreddatabase.ItemsClientListMirroredDatabasesOptions{ContinuationToken: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		testsuite.Require().NoError(err, "Failed to advance page for example ")
		testsuite.Require().True(reflect.DeepEqual(exampleRes, nextResult.MirroredDatabases))
		if err == nil {
			break
		}
	}
}

func (testsuite *FakeTestSuite) TestItems_CreateMirroredDatabase() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Create a mirrored database with public definition example"},
	})
	var exampleWorkspaceID string
	var exampleCreateMirroredDatabaseRequest mirroreddatabase.CreateMirroredDatabaseRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleCreateMirroredDatabaseRequest = mirroreddatabase.CreateMirroredDatabaseRequest{
		Description: to.Ptr("A mirrored database description"),
		Definition: &mirroreddatabase.Definition{
			Parts: []mirroreddatabase.DefinitionPart{
				{
					Path:        to.Ptr("mirroredDatabase.json"),
					Payload:     to.Ptr("eyAicHJvcGVydGllcy..WJsZSIgfSB9IH0gXSB9IH0"),
					PayloadType: to.Ptr(mirroreddatabase.PayloadTypeInlineBase64),
				}},
		},
		DisplayName: to.Ptr("Mirrored database 1"),
	}

	testsuite.serverFactory.ItemsServer.CreateMirroredDatabase = func(ctx context.Context, workspaceID string, createMirroredDatabaseRequest mirroreddatabase.CreateMirroredDatabaseRequest, options *mirroreddatabase.ItemsClientCreateMirroredDatabaseOptions) (resp azfake.Responder[mirroreddatabase.ItemsClientCreateMirroredDatabaseResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().True(reflect.DeepEqual(exampleCreateMirroredDatabaseRequest, createMirroredDatabaseRequest))
		resp = azfake.Responder[mirroreddatabase.ItemsClientCreateMirroredDatabaseResponse]{}
		resp.SetResponse(http.StatusCreated, mirroreddatabase.ItemsClientCreateMirroredDatabaseResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	_, err = client.CreateMirroredDatabase(ctx, exampleWorkspaceID, exampleCreateMirroredDatabaseRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
}

func (testsuite *FakeTestSuite) TestItems_GetMirroredDatabase() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Get a mirrored database example"},
	})
	var exampleWorkspaceID string
	var exampleMirroredDatabaseID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleMirroredDatabaseID = "5b218778-e7a5-4d73-8187-f10824047715"

	exampleRes := mirroreddatabase.MirroredDatabase{
		Type:        to.Ptr(mirroreddatabase.ItemTypeMirroredDatabase),
		Description: to.Ptr("A mirrored database description."),
		DisplayName: to.Ptr("Mirrored database 1"),
		ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		Properties: &mirroreddatabase.Properties{
			DefaultSchema:     to.Ptr("dbo"),
			OneLakeTablesPath: to.Ptr("https://onelake.dfs.fabric.microsoft.com/cfafbeb1-8037-4d0c-896e-a46fb27ff229/5b218778-e7a5-4d73-8187-f10824047715/Tables"),
			SQLEndpointProperties: &mirroreddatabase.SQLEndpointProperties{
				ConnectionString:   to.Ptr("xxx.datawarehouse.fabric.microsoft.com"),
				ID:                 to.Ptr("84c9ad4a-ba9e-42d9-9025-b40f8e38c025"),
				ProvisioningStatus: to.Ptr(mirroreddatabase.SQLEndpointProvisioningStatusSuccess),
			},
		},
	}

	testsuite.serverFactory.ItemsServer.GetMirroredDatabase = func(ctx context.Context, workspaceID string, mirroredDatabaseID string, options *mirroreddatabase.ItemsClientGetMirroredDatabaseOptions) (resp azfake.Responder[mirroreddatabase.ItemsClientGetMirroredDatabaseResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleMirroredDatabaseID, mirroredDatabaseID)
		resp = azfake.Responder[mirroreddatabase.ItemsClientGetMirroredDatabaseResponse]{}
		resp.SetResponse(http.StatusOK, mirroreddatabase.ItemsClientGetMirroredDatabaseResponse{MirroredDatabase: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	res, err := client.GetMirroredDatabase(ctx, exampleWorkspaceID, exampleMirroredDatabaseID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.MirroredDatabase))
}

func (testsuite *FakeTestSuite) TestItems_UpdateMirroredDatabase() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Update a mirrored database example"},
	})
	var exampleWorkspaceID string
	var exampleMirroredDatabaseID string
	var exampleUpdateMirroredDatabaseRequest mirroreddatabase.UpdateMirroredDatabaseRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleMirroredDatabaseID = "5b218778-e7a5-4d73-8187-f10824047715"
	exampleUpdateMirroredDatabaseRequest = mirroreddatabase.UpdateMirroredDatabaseRequest{
		Description: to.Ptr("A new description for mirrored database."),
		DisplayName: to.Ptr("MirroredDatabase's New name"),
	}

	exampleRes := mirroreddatabase.MirroredDatabase{
		Type:        to.Ptr(mirroreddatabase.ItemTypeMirroredDatabase),
		Description: to.Ptr("A new description for mirrored database."),
		DisplayName: to.Ptr("MirroredDatabase's New name"),
		ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	}

	testsuite.serverFactory.ItemsServer.UpdateMirroredDatabase = func(ctx context.Context, workspaceID string, mirroredDatabaseID string, updateMirroredDatabaseRequest mirroreddatabase.UpdateMirroredDatabaseRequest, options *mirroreddatabase.ItemsClientUpdateMirroredDatabaseOptions) (resp azfake.Responder[mirroreddatabase.ItemsClientUpdateMirroredDatabaseResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleMirroredDatabaseID, mirroredDatabaseID)
		testsuite.Require().True(reflect.DeepEqual(exampleUpdateMirroredDatabaseRequest, updateMirroredDatabaseRequest))
		resp = azfake.Responder[mirroreddatabase.ItemsClientUpdateMirroredDatabaseResponse]{}
		resp.SetResponse(http.StatusOK, mirroreddatabase.ItemsClientUpdateMirroredDatabaseResponse{MirroredDatabase: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	res, err := client.UpdateMirroredDatabase(ctx, exampleWorkspaceID, exampleMirroredDatabaseID, exampleUpdateMirroredDatabaseRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.MirroredDatabase))
}

func (testsuite *FakeTestSuite) TestItems_DeleteMirroredDatabase() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Delete a mirroed database example"},
	})
	var exampleWorkspaceID string
	var exampleMirroredDatabaseID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleMirroredDatabaseID = "5b218778-e7a5-4d73-8187-f10824047715"

	testsuite.serverFactory.ItemsServer.DeleteMirroredDatabase = func(ctx context.Context, workspaceID string, mirroredDatabaseID string, options *mirroreddatabase.ItemsClientDeleteMirroredDatabaseOptions) (resp azfake.Responder[mirroreddatabase.ItemsClientDeleteMirroredDatabaseResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleMirroredDatabaseID, mirroredDatabaseID)
		resp = azfake.Responder[mirroreddatabase.ItemsClientDeleteMirroredDatabaseResponse]{}
		resp.SetResponse(http.StatusOK, mirroreddatabase.ItemsClientDeleteMirroredDatabaseResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	_, err = client.DeleteMirroredDatabase(ctx, exampleWorkspaceID, exampleMirroredDatabaseID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
}

func (testsuite *FakeTestSuite) TestItems_GetMirroredDatabaseDefinition() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Get a mirrored database definition example"},
	})
	var exampleWorkspaceID string
	var exampleMirroredDatabaseID string
	exampleWorkspaceID = "6e335e92-a2a2-4b5a-970a-bd6a89fbb765"
	exampleMirroredDatabaseID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"

	exampleRes := mirroreddatabase.DefinitionResponse{
		Definition: &mirroreddatabase.Definition{
			Parts: []mirroreddatabase.DefinitionPart{
				{
					Path:        to.Ptr("mirroredDatabase.json"),
					Payload:     to.Ptr("eyAicHJvcGVydGllcy..WJsZSIgfSB9IH0gXSB9IH0"),
					PayloadType: to.Ptr(mirroreddatabase.PayloadTypeInlineBase64),
				},
				{
					Path:        to.Ptr(".platform"),
					Payload:     to.Ptr("ZG90UGxhdGZvcm1CYXNlNjRTdHJpbmc="),
					PayloadType: to.Ptr(mirroreddatabase.PayloadTypeInlineBase64),
				}},
		},
	}

	testsuite.serverFactory.ItemsServer.GetMirroredDatabaseDefinition = func(ctx context.Context, workspaceID string, mirroredDatabaseID string, options *mirroreddatabase.ItemsClientGetMirroredDatabaseDefinitionOptions) (resp azfake.Responder[mirroreddatabase.ItemsClientGetMirroredDatabaseDefinitionResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleMirroredDatabaseID, mirroredDatabaseID)
		resp = azfake.Responder[mirroreddatabase.ItemsClientGetMirroredDatabaseDefinitionResponse]{}
		resp.SetResponse(http.StatusOK, mirroreddatabase.ItemsClientGetMirroredDatabaseDefinitionResponse{DefinitionResponse: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	res, err := client.GetMirroredDatabaseDefinition(ctx, exampleWorkspaceID, exampleMirroredDatabaseID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.DefinitionResponse))
}

func (testsuite *FakeTestSuite) TestItems_UpdateMirroredDatabaseDefinition() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Update a mirrored database definition example"},
	})
	var exampleWorkspaceID string
	var exampleMirroredDatabaseID string
	var exampleUpdateMirroredDatabaseDefinitionRequest mirroreddatabase.UpdateMirroredDatabaseDefinitionRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleMirroredDatabaseID = "5b218778-e7a5-4d73-8187-f10824047715"
	exampleUpdateMirroredDatabaseDefinitionRequest = mirroreddatabase.UpdateMirroredDatabaseDefinitionRequest{
		Definition: &mirroreddatabase.Definition{
			Parts: []mirroreddatabase.DefinitionPart{
				{
					Path:        to.Ptr("mirroredDatabase.json"),
					Payload:     to.Ptr("eyAicHJvcGVydGllcy..WJsZSIgfSB9IH0gXSB9IH0"),
					PayloadType: to.Ptr(mirroreddatabase.PayloadTypeInlineBase64),
				},
				{
					Path:        to.Ptr(".platform"),
					Payload:     to.Ptr("ZG90UGxhdGZvcm1CYXNlNjRTdHJpbmc="),
					PayloadType: to.Ptr(mirroreddatabase.PayloadTypeInlineBase64),
				}},
		},
	}

	testsuite.serverFactory.ItemsServer.UpdateMirroredDatabaseDefinition = func(ctx context.Context, workspaceID string, mirroredDatabaseID string, updateMirroredDatabaseDefinitionRequest mirroreddatabase.UpdateMirroredDatabaseDefinitionRequest, options *mirroreddatabase.ItemsClientUpdateMirroredDatabaseDefinitionOptions) (resp azfake.Responder[mirroreddatabase.ItemsClientUpdateMirroredDatabaseDefinitionResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleMirroredDatabaseID, mirroredDatabaseID)
		testsuite.Require().True(reflect.DeepEqual(exampleUpdateMirroredDatabaseDefinitionRequest, updateMirroredDatabaseDefinitionRequest))
		resp = azfake.Responder[mirroreddatabase.ItemsClientUpdateMirroredDatabaseDefinitionResponse]{}
		resp.SetResponse(http.StatusOK, mirroreddatabase.ItemsClientUpdateMirroredDatabaseDefinitionResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	_, err = client.UpdateMirroredDatabaseDefinition(ctx, exampleWorkspaceID, exampleMirroredDatabaseID, exampleUpdateMirroredDatabaseDefinitionRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
}

func (testsuite *FakeTestSuite) TestMirroring_StartMirroring() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Start mirroring example"},
	})
	var exampleWorkspaceID string
	var exampleMirroredDatabaseID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleMirroredDatabaseID = "5b218778-e7a5-4d73-8187-f10824047715"

	testsuite.serverFactory.MirroringServer.StartMirroring = func(ctx context.Context, workspaceID string, mirroredDatabaseID string, options *mirroreddatabase.MirroringClientStartMirroringOptions) (resp azfake.Responder[mirroreddatabase.MirroringClientStartMirroringResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleMirroredDatabaseID, mirroredDatabaseID)
		resp = azfake.Responder[mirroreddatabase.MirroringClientStartMirroringResponse]{}
		resp.SetResponse(http.StatusOK, mirroreddatabase.MirroringClientStartMirroringResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewMirroringClient()
	_, err = client.StartMirroring(ctx, exampleWorkspaceID, exampleMirroredDatabaseID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
}

func (testsuite *FakeTestSuite) TestMirroring_StopMirroring() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Stop mirroring example"},
	})
	var exampleWorkspaceID string
	var exampleMirroredDatabaseID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleMirroredDatabaseID = "5b218778-e7a5-4d73-8187-f10824047715"

	testsuite.serverFactory.MirroringServer.StopMirroring = func(ctx context.Context, workspaceID string, mirroredDatabaseID string, options *mirroreddatabase.MirroringClientStopMirroringOptions) (resp azfake.Responder[mirroreddatabase.MirroringClientStopMirroringResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleMirroredDatabaseID, mirroredDatabaseID)
		resp = azfake.Responder[mirroreddatabase.MirroringClientStopMirroringResponse]{}
		resp.SetResponse(http.StatusOK, mirroreddatabase.MirroringClientStopMirroringResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewMirroringClient()
	_, err = client.StopMirroring(ctx, exampleWorkspaceID, exampleMirroredDatabaseID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
}

func (testsuite *FakeTestSuite) TestMirroring_GetMirroringStatus() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Get mirroring status example"},
	})
	var exampleWorkspaceID string
	var exampleMirroredDatabaseID string
	exampleWorkspaceID = "6e335e92-a2a2-4b5a-970a-bd6a89fbb765"
	exampleMirroredDatabaseID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"

	exampleRes := mirroreddatabase.MirroringStatusResponse{
		Status: to.Ptr(mirroreddatabase.MirroringStatusRunning),
	}

	testsuite.serverFactory.MirroringServer.GetMirroringStatus = func(ctx context.Context, workspaceID string, mirroredDatabaseID string, options *mirroreddatabase.MirroringClientGetMirroringStatusOptions) (resp azfake.Responder[mirroreddatabase.MirroringClientGetMirroringStatusResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleMirroredDatabaseID, mirroredDatabaseID)
		resp = azfake.Responder[mirroreddatabase.MirroringClientGetMirroringStatusResponse]{}
		resp.SetResponse(http.StatusOK, mirroreddatabase.MirroringClientGetMirroringStatusResponse{MirroringStatusResponse: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewMirroringClient()
	res, err := client.GetMirroringStatus(ctx, exampleWorkspaceID, exampleMirroredDatabaseID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.MirroringStatusResponse))
}

func (testsuite *FakeTestSuite) TestMirroring_GetTablesMirroringStatus() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Get tables mirroring status example"},
	})
	var exampleWorkspaceID string
	var exampleMirroredDatabaseID string
	exampleWorkspaceID = "6e335e92-a2a2-4b5a-970a-bd6a89fbb765"
	exampleMirroredDatabaseID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"

	exampleRes := mirroreddatabase.TablesMirroringStatusResponse{
		Data: []mirroreddatabase.TableMirroringStatusResponse{
			{
				Metrics: &mirroreddatabase.TableMirroringMetrics{
					LastSyncDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-08T05:07:11.066Z"); return t }()),
					ProcessedBytes:   to.Ptr[int64](1247),
					ProcessedRows:    to.Ptr[int64](6),
				},
				SourceSchemaName: to.Ptr("dbo"),
				SourceTableName:  to.Ptr("test"),
				Status:           to.Ptr(mirroreddatabase.TableMirroringStatusReplicating),
			}},
	}

	testsuite.serverFactory.MirroringServer.GetTablesMirroringStatus = func(ctx context.Context, workspaceID string, mirroredDatabaseID string, options *mirroreddatabase.MirroringClientGetTablesMirroringStatusOptions) (resp azfake.Responder[mirroreddatabase.MirroringClientGetTablesMirroringStatusResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleMirroredDatabaseID, mirroredDatabaseID)
		resp = azfake.Responder[mirroreddatabase.MirroringClientGetTablesMirroringStatusResponse]{}
		resp.SetResponse(http.StatusOK, mirroreddatabase.MirroringClientGetTablesMirroringStatusResponse{TablesMirroringStatusResponse: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewMirroringClient()
	res, err := client.GetTablesMirroringStatus(ctx, exampleWorkspaceID, exampleMirroredDatabaseID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.TablesMirroringStatusResponse))
}
