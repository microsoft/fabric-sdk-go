// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package apacheairflowjob_test

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

	"github.com/microsoft/fabric-sdk-go/fabric/apacheairflowjob"
	"github.com/microsoft/fabric-sdk-go/fabric/apacheairflowjob/fake"
)

var err error

type FakeTestSuite struct {
	suite.Suite

	ctx  context.Context
	cred azcore.TokenCredential

	serverFactory *fake.ServerFactory
	clientFactory *apacheairflowjob.ClientFactory
}

func (testsuite *FakeTestSuite) SetupSuite() {
	testsuite.ctx = context.Background()
	testsuite.cred = &azfake.TokenCredential{}

	testsuite.serverFactory = &fake.ServerFactory{}
	testsuite.clientFactory, err = apacheairflowjob.NewClientFactory(testsuite.cred, nil, &azcore.ClientOptions{
		Transport: fake.NewServerFactoryTransport(testsuite.serverFactory),
	})
	testsuite.Require().NoError(err, "Failed to create client factory")
}

func TestFakeTest(t *testing.T) {
	suite.Run(t, new(FakeTestSuite))
}

func (testsuite *FakeTestSuite) TestItems_ListApacheAirflowJobs() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"List s in workspace example"},
	})
	var exampleWorkspaceID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"

	exampleRes := apacheairflowjob.ApacheAirflowJobs{
		Value: []apacheairflowjob.ApacheAirflowJob{
			{
				Type:        to.Ptr(apacheairflowjob.ItemTypeApacheAirflowJob),
				Description: to.Ptr("An Apache Airflow job description."),
				DisplayName: to.Ptr("ApacheAirflowJobName1"),
				ID:          to.Ptr("3546052c-ae64-4526-b1a8-52af7761426f"),
				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
			},
			{
				Type:        to.Ptr(apacheairflowjob.ItemTypeApacheAirflowJob),
				Description: to.Ptr("An Apache Airflow job description."),
				DisplayName: to.Ptr("ApacheAirflowJob2"),
				ID:          to.Ptr("f697fb63-abd4-4399-9548-be7e3c3c0dac"),
				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
			}},
	}

	testsuite.serverFactory.ItemsServer.NewListApacheAirflowJobsPager = func(workspaceID string, options *apacheairflowjob.ItemsClientListApacheAirflowJobsOptions) (resp azfake.PagerResponder[apacheairflowjob.ItemsClientListApacheAirflowJobsResponse]) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		resp = azfake.PagerResponder[apacheairflowjob.ItemsClientListApacheAirflowJobsResponse]{}
		resp.AddPage(http.StatusOK, apacheairflowjob.ItemsClientListApacheAirflowJobsResponse{ApacheAirflowJobs: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	pager := client.NewListApacheAirflowJobsPager(exampleWorkspaceID, &apacheairflowjob.ItemsClientListApacheAirflowJobsOptions{ContinuationToken: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		testsuite.Require().NoError(err, "Failed to advance page for example ")
		testsuite.Require().True(reflect.DeepEqual(exampleRes, nextResult.ApacheAirflowJobs))
		if err == nil {
			break
		}
	}
}

func (testsuite *FakeTestSuite) TestItems_CreateApacheAirflowJob() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Create an Apache Airflow job with public definition example"},
	})
	var exampleWorkspaceID string
	var exampleCreateApacheAirflowJobRequest apacheairflowjob.CreateApacheAirflowJobRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleCreateApacheAirflowJobRequest = apacheairflowjob.CreateApacheAirflowJobRequest{
		Description: to.Ptr("An Apache Airflow job description."),
		DisplayName: to.Ptr("ApacheAirflowJob1"),
	}

	testsuite.serverFactory.ItemsServer.BeginCreateApacheAirflowJob = func(ctx context.Context, workspaceID string, createApacheAirflowJobRequest apacheairflowjob.CreateApacheAirflowJobRequest, options *apacheairflowjob.ItemsClientBeginCreateApacheAirflowJobOptions) (resp azfake.PollerResponder[apacheairflowjob.ItemsClientCreateApacheAirflowJobResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().True(reflect.DeepEqual(exampleCreateApacheAirflowJobRequest, createApacheAirflowJobRequest))
		resp = azfake.PollerResponder[apacheairflowjob.ItemsClientCreateApacheAirflowJobResponse]{}
		resp.SetTerminalResponse(http.StatusCreated, apacheairflowjob.ItemsClientCreateApacheAirflowJobResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	poller, err := client.BeginCreateApacheAirflowJob(ctx, exampleWorkspaceID, exampleCreateApacheAirflowJobRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")
}

func (testsuite *FakeTestSuite) TestItems_GetApacheAirflowJob() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Get an Apache Airflow job example"},
	})
	var exampleWorkspaceID string
	var exampleApacheAirflowJobID string
	exampleWorkspaceID = "f089354e-8366-4e18-aea3-4cb4a3a50b48"
	exampleApacheAirflowJobID = "41ce06d1-d81b-4ea0-bc6d-2ce3dd2f8e87"

	exampleRes := apacheairflowjob.ApacheAirflowJob{
		Type:        to.Ptr(apacheairflowjob.ItemTypeApacheAirflowJob),
		Description: to.Ptr("An Apache Airflow job description."),
		DisplayName: to.Ptr("ApacheAirflowJob1"),
		ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	}

	testsuite.serverFactory.ItemsServer.GetApacheAirflowJob = func(ctx context.Context, workspaceID string, apacheAirflowJobID string, options *apacheairflowjob.ItemsClientGetApacheAirflowJobOptions) (resp azfake.Responder[apacheairflowjob.ItemsClientGetApacheAirflowJobResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleApacheAirflowJobID, apacheAirflowJobID)
		resp = azfake.Responder[apacheairflowjob.ItemsClientGetApacheAirflowJobResponse]{}
		resp.SetResponse(http.StatusOK, apacheairflowjob.ItemsClientGetApacheAirflowJobResponse{ApacheAirflowJob: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	res, err := client.GetApacheAirflowJob(ctx, exampleWorkspaceID, exampleApacheAirflowJobID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.ApacheAirflowJob))
}

func (testsuite *FakeTestSuite) TestItems_UpdateApacheAirflowJob() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Update an Apache Airflow job example"},
	})
	var exampleWorkspaceID string
	var exampleApacheAirflowJobID string
	var exampleUpdateApacheAirflowJobRequest apacheairflowjob.UpdateApacheAirflowJobRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleApacheAirflowJobID = "5b218778-e7a5-4d73-8187-f10824047715"
	exampleUpdateApacheAirflowJobRequest = apacheairflowjob.UpdateApacheAirflowJobRequest{
		Description: to.Ptr("Apache Airflow job's New description"),
	}

	exampleRes := apacheairflowjob.ApacheAirflowJob{
		Type:        to.Ptr(apacheairflowjob.ItemTypeApacheAirflowJob),
		Description: to.Ptr("Apache Airflow job's New description"),
		DisplayName: to.Ptr("ApacheAirflowJob1"),
		ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	}

	testsuite.serverFactory.ItemsServer.UpdateApacheAirflowJob = func(ctx context.Context, workspaceID string, apacheAirflowJobID string, updateApacheAirflowJobRequest apacheairflowjob.UpdateApacheAirflowJobRequest, options *apacheairflowjob.ItemsClientUpdateApacheAirflowJobOptions) (resp azfake.Responder[apacheairflowjob.ItemsClientUpdateApacheAirflowJobResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleApacheAirflowJobID, apacheAirflowJobID)
		testsuite.Require().True(reflect.DeepEqual(exampleUpdateApacheAirflowJobRequest, updateApacheAirflowJobRequest))
		resp = azfake.Responder[apacheairflowjob.ItemsClientUpdateApacheAirflowJobResponse]{}
		resp.SetResponse(http.StatusOK, apacheairflowjob.ItemsClientUpdateApacheAirflowJobResponse{ApacheAirflowJob: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	res, err := client.UpdateApacheAirflowJob(ctx, exampleWorkspaceID, exampleApacheAirflowJobID, exampleUpdateApacheAirflowJobRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.ApacheAirflowJob))
}

func (testsuite *FakeTestSuite) TestItems_DeleteApacheAirflowJob() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Delete an Apache Airflow job example"},
	})
	var exampleWorkspaceID string
	var exampleApacheAirflowJobID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleApacheAirflowJobID = "5b218778-e7a5-4d73-8187-f10824047715"

	testsuite.serverFactory.ItemsServer.DeleteApacheAirflowJob = func(ctx context.Context, workspaceID string, apacheAirflowJobID string, options *apacheairflowjob.ItemsClientDeleteApacheAirflowJobOptions) (resp azfake.Responder[apacheairflowjob.ItemsClientDeleteApacheAirflowJobResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleApacheAirflowJobID, apacheAirflowJobID)
		resp = azfake.Responder[apacheairflowjob.ItemsClientDeleteApacheAirflowJobResponse]{}
		resp.SetResponse(http.StatusOK, apacheairflowjob.ItemsClientDeleteApacheAirflowJobResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	_, err = client.DeleteApacheAirflowJob(ctx, exampleWorkspaceID, exampleApacheAirflowJobID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
}

func (testsuite *FakeTestSuite) TestItems_GetApacheAirflowJobDefinition() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Get an Apache Airflow job public definition example"},
	})
	var exampleWorkspaceID string
	var exampleApacheAirflowJobID string
	exampleWorkspaceID = "6e335e92-a2a2-4b5a-970a-bd6a89fbb765"
	exampleApacheAirflowJobID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"

	exampleRes := apacheairflowjob.DefinitionResponse{
		Definition: &apacheairflowjob.Definition{
			Parts: []apacheairflowjob.DefinitionPart{
				{
					Path:        to.Ptr("ApacheAirflowJobV1.json"),
					Payload:     to.Ptr("ewogICJwcm9wZXJ0aWVzIjogewogICAgInR5cGUiOiAiQXBhY2hlQWlyZmxvd0pvYiIsCiAgICAidHlwZVByb3BlcnRpZXMiOiB7CiAgICAgICJhaXJmbG93UHJvcGVydGllcyI6IHsKICAgICAgICAiYWlyZmxvd1ZlcnNpb24iOiAiMi42LjMiLAogICAgICAgICJweXRob25WZXJzaW9uIjogIjMuOCIsCiAgICAgICAgImVudmlyb25tZW50VmFyaWFibGVzIjoge30sCiAgICAgICAgImFpcmZsb3dDb25maWd1cmF0aW9uT3ZlcnJpZGVzIjoge30sCiAgICAgICAgImFpcmZsb3dSZXF1aXJlbWVudHMiOiBbCiAgICAgICAgICAiZmxhc2stYmNyeXB0PT0wLjcuMSIKICAgICAgICBdLAogICAgICAgICJwYWNrYWdlUHJvdmlkZXJQYXRoIjogInBsdWdpbnMiLAogICAgICAgICJlbmFibGVBQURJbnRlZ3JhdGlvbiI6IHRydWUsCiAgICAgICAgImVuYWJsZVRyaWdnZXJlcnMiOiBmYWxzZSwKICAgICAgICAic2VjcmV0cyI6IFtdCiAgICAgIH0sCiAgICAgICJjb21wdXRlUHJvcGVydGllcyI6IHsKICAgICAgICAiY29tcHV0ZVBvb2wiOiAiU3RhcnRlclBvb2wiLAogICAgICAgICJsb2NhdGlvbiI6ICJDZW50cmFsIFVTIiwKICAgICAgICAiY29tcHV0ZVNpemUiOiAiU21hbGwiLAogICAgICAgICJleHRyYU5vZGVzIjogMCwKICAgICAgICAiZW5hYmxlQXZhaWxhYmlsaXR5Wm9uZXMiOiBmYWxzZSwKICAgICAgICAiZW5hYmxlQXV0b3NjYWxlIjogZmFsc2UKICAgICAgfQogICAgfQogIH0KfQ=="),
					PayloadType: to.Ptr(apacheairflowjob.PayloadTypeInlineBase64),
				},
				{
					Path:        to.Ptr(".platform"),
					Payload:     to.Ptr("ZG90UGxhdGZvcm1CYXNlNjRTdHJpbmc="),
					PayloadType: to.Ptr(apacheairflowjob.PayloadTypeInlineBase64),
				}},
		},
	}

	testsuite.serverFactory.ItemsServer.BeginGetApacheAirflowJobDefinition = func(ctx context.Context, workspaceID string, apacheAirflowJobID string, options *apacheairflowjob.ItemsClientBeginGetApacheAirflowJobDefinitionOptions) (resp azfake.PollerResponder[apacheairflowjob.ItemsClientGetApacheAirflowJobDefinitionResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleApacheAirflowJobID, apacheAirflowJobID)
		resp = azfake.PollerResponder[apacheairflowjob.ItemsClientGetApacheAirflowJobDefinitionResponse]{}
		resp.SetTerminalResponse(http.StatusOK, apacheairflowjob.ItemsClientGetApacheAirflowJobDefinitionResponse{DefinitionResponse: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	poller, err := client.BeginGetApacheAirflowJobDefinition(ctx, exampleWorkspaceID, exampleApacheAirflowJobID, &apacheairflowjob.ItemsClientBeginGetApacheAirflowJobDefinitionOptions{Format: nil})
	testsuite.Require().NoError(err, "Failed to get result for example ")
	res, err := poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.DefinitionResponse))
}

func (testsuite *FakeTestSuite) TestItems_UpdateApacheAirflowJobDefinition() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Update an Apache Airflow job public definition example"},
	})
	var exampleWorkspaceID string
	var exampleApacheAirflowJobID string
	var exampleUpdateApacheAirflowJobDefinitionRequest apacheairflowjob.UpdateApacheAirflowJobDefinitionRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleApacheAirflowJobID = "5b218778-e7a5-4d73-8187-f10824047715"
	exampleUpdateApacheAirflowJobDefinitionRequest = apacheairflowjob.UpdateApacheAirflowJobDefinitionRequest{
		Definition: &apacheairflowjob.Definition{
			Parts: []apacheairflowjob.DefinitionPart{
				{
					Path:        to.Ptr("ApacheAirflowJobV1.json"),
					Payload:     to.Ptr("ewogICJwcm9wZXJ0aWVzIjogewogICAgInR5cGUiOiAiQXBhY2hlQWlyZmxvd0pvYiIsCiAgICAidHlwZVByb3BlcnRpZXMiOiB7CiAgICAgICJhaXJmbG93UHJvcGVydGllcyI6IHsKICAgICAgICAiYWlyZmxvd1ZlcnNpb24iOiAiMi42LjMiLAogICAgICAgICJweXRob25WZXJzaW9uIjogIjMuOCIsCiAgICAgICAgImVudmlyb25tZW50VmFyaWFibGVzIjoge30sCiAgICAgICAgImFpcmZsb3dDb25maWd1cmF0aW9uT3ZlcnJpZGVzIjoge30sCiAgICAgICAgImFpcmZsb3dSZXF1aXJlbWVudHMiOiBbCiAgICAgICAgICAiZmxhc2stYmNyeXB0PT0wLjcuMSIKICAgICAgICBdLAogICAgICAgICJwYWNrYWdlUHJvdmlkZXJQYXRoIjogInBsdWdpbnMiLAogICAgICAgICJlbmFibGVBQURJbnRlZ3JhdGlvbiI6IHRydWUsCiAgICAgICAgImVuYWJsZVRyaWdnZXJlcnMiOiBmYWxzZSwKICAgICAgICAic2VjcmV0cyI6IFtdCiAgICAgIH0sCiAgICAgICJjb21wdXRlUHJvcGVydGllcyI6IHsKICAgICAgICAiY29tcHV0ZVBvb2wiOiAiU3RhcnRlclBvb2wiLAogICAgICAgICJsb2NhdGlvbiI6ICJDZW50cmFsIFVTIiwKICAgICAgICAiY29tcHV0ZVNpemUiOiAiU21hbGwiLAogICAgICAgICJleHRyYU5vZGVzIjogMCwKICAgICAgICAiZW5hYmxlQXZhaWxhYmlsaXR5Wm9uZXMiOiBmYWxzZSwKICAgICAgICAiZW5hYmxlQXV0b3NjYWxlIjogZmFsc2UKICAgICAgfQogICAgfQogIH0KfQ=="),
					PayloadType: to.Ptr(apacheairflowjob.PayloadTypeInlineBase64),
				},
				{
					Path:        to.Ptr(".platform"),
					Payload:     to.Ptr("ZG90UGxhdGZvcm1CYXNlNjRTdHJpbmc="),
					PayloadType: to.Ptr(apacheairflowjob.PayloadTypeInlineBase64),
				}},
		},
	}

	testsuite.serverFactory.ItemsServer.BeginUpdateApacheAirflowJobDefinition = func(ctx context.Context, workspaceID string, apacheAirflowJobID string, updateApacheAirflowJobDefinitionRequest apacheairflowjob.UpdateApacheAirflowJobDefinitionRequest, options *apacheairflowjob.ItemsClientBeginUpdateApacheAirflowJobDefinitionOptions) (resp azfake.PollerResponder[apacheairflowjob.ItemsClientUpdateApacheAirflowJobDefinitionResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleApacheAirflowJobID, apacheAirflowJobID)
		testsuite.Require().True(reflect.DeepEqual(exampleUpdateApacheAirflowJobDefinitionRequest, updateApacheAirflowJobDefinitionRequest))
		resp = azfake.PollerResponder[apacheairflowjob.ItemsClientUpdateApacheAirflowJobDefinitionResponse]{}
		resp.SetTerminalResponse(http.StatusOK, apacheairflowjob.ItemsClientUpdateApacheAirflowJobDefinitionResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	poller, err := client.BeginUpdateApacheAirflowJobDefinition(ctx, exampleWorkspaceID, exampleApacheAirflowJobID, exampleUpdateApacheAirflowJobDefinitionRequest, &apacheairflowjob.ItemsClientBeginUpdateApacheAirflowJobDefinitionOptions{UpdateMetadata: to.Ptr(true)})
	testsuite.Require().NoError(err, "Failed to get result for example ")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")
}
