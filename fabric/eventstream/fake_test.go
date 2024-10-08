// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package eventstream_test

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

	"github.com/microsoft/fabric-sdk-go/fabric/eventstream"
	"github.com/microsoft/fabric-sdk-go/fabric/eventstream/fake"
)

var err error

type FakeTestSuite struct {
	suite.Suite

	ctx  context.Context
	cred azcore.TokenCredential

	serverFactory *fake.ServerFactory
	clientFactory *eventstream.ClientFactory
}

func (testsuite *FakeTestSuite) SetupSuite() {
	testsuite.ctx = context.Background()
	testsuite.cred = &azfake.TokenCredential{}

	testsuite.serverFactory = &fake.ServerFactory{}
	testsuite.clientFactory, err = eventstream.NewClientFactory(testsuite.cred, nil, &azcore.ClientOptions{
		Transport: fake.NewServerFactoryTransport(testsuite.serverFactory),
	})
	testsuite.Require().NoError(err, "Failed to create client factory")
}

func TestFakeTest(t *testing.T) {
	suite.Run(t, new(FakeTestSuite))
}

func (testsuite *FakeTestSuite) TestItems_ListEventstreams() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"List eventstreams in workspace example"},
	})
	var exampleWorkspaceID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"

	exampleRes := eventstream.Eventstreams{
		Value: []eventstream.Eventstream{
			{
				Type:        to.Ptr(eventstream.ItemTypeEventstream),
				Description: to.Ptr("An eventstream description."),
				DisplayName: to.Ptr("Eventstream_1"),
				ID:          to.Ptr("3546052c-ae64-4526-b1a8-52af7761426f"),
				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
			},
			{
				Type:        to.Ptr(eventstream.ItemTypeEventstream),
				Description: to.Ptr("An eventstream description."),
				DisplayName: to.Ptr("Eventstream_2"),
				ID:          to.Ptr("dc307e72-3e89-4425-97c3-a364d86dddfa"),
				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
			},
			{
				Type:        to.Ptr(eventstream.ItemTypeEventstream),
				Description: to.Ptr("An eventstream description."),
				DisplayName: to.Ptr("Eventstream_3"),
				ID:          to.Ptr("b48a1f80-862a-4b97-b672-7217bf064dc0"),
				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
			}},
	}

	testsuite.serverFactory.ItemsServer.NewListEventstreamsPager = func(workspaceID string, options *eventstream.ItemsClientListEventstreamsOptions) (resp azfake.PagerResponder[eventstream.ItemsClientListEventstreamsResponse]) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		resp = azfake.PagerResponder[eventstream.ItemsClientListEventstreamsResponse]{}
		resp.AddPage(http.StatusOK, eventstream.ItemsClientListEventstreamsResponse{Eventstreams: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	pager := client.NewListEventstreamsPager(exampleWorkspaceID, &eventstream.ItemsClientListEventstreamsOptions{ContinuationToken: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		testsuite.Require().NoError(err, "Failed to advance page for example ")
		testsuite.Require().True(reflect.DeepEqual(exampleRes, nextResult.Eventstreams))
		if err == nil {
			break
		}
	}
}

func (testsuite *FakeTestSuite) TestItems_CreateEventstream() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Create an eventstream example"},
	})
	var exampleWorkspaceID string
	var exampleCreateEventstreamRequest eventstream.CreateEventstreamRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleCreateEventstreamRequest = eventstream.CreateEventstreamRequest{
		Description: to.Ptr("An eventstream description."),
		DisplayName: to.Ptr("Eventstream_1"),
	}

	testsuite.serverFactory.ItemsServer.BeginCreateEventstream = func(ctx context.Context, workspaceID string, createEventstreamRequest eventstream.CreateEventstreamRequest, options *eventstream.ItemsClientBeginCreateEventstreamOptions) (resp azfake.PollerResponder[eventstream.ItemsClientCreateEventstreamResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().True(reflect.DeepEqual(exampleCreateEventstreamRequest, createEventstreamRequest))
		resp = azfake.PollerResponder[eventstream.ItemsClientCreateEventstreamResponse]{}
		resp.SetTerminalResponse(http.StatusCreated, eventstream.ItemsClientCreateEventstreamResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	poller, err := client.BeginCreateEventstream(ctx, exampleWorkspaceID, exampleCreateEventstreamRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example ")
}

func (testsuite *FakeTestSuite) TestItems_GetEventstream() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Get an eventstream example"},
	})
	var exampleWorkspaceID string
	var exampleEventstreamID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleEventstreamID = "5b218778-e7a5-4d73-8187-f10824047715"

	exampleRes := eventstream.Eventstream{
		Type:        to.Ptr(eventstream.ItemTypeEventstream),
		Description: to.Ptr("An eventstream description."),
		DisplayName: to.Ptr("Eventstream_1"),
		ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	}

	testsuite.serverFactory.ItemsServer.GetEventstream = func(ctx context.Context, workspaceID string, eventstreamID string, options *eventstream.ItemsClientGetEventstreamOptions) (resp azfake.Responder[eventstream.ItemsClientGetEventstreamResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleEventstreamID, eventstreamID)
		resp = azfake.Responder[eventstream.ItemsClientGetEventstreamResponse]{}
		resp.SetResponse(http.StatusOK, eventstream.ItemsClientGetEventstreamResponse{Eventstream: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	res, err := client.GetEventstream(ctx, exampleWorkspaceID, exampleEventstreamID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.Eventstream))
}

func (testsuite *FakeTestSuite) TestItems_UpdateEventstream() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Update an eventstream example"},
	})
	var exampleWorkspaceID string
	var exampleEventstreamID string
	var exampleUpdateEventstreamRequest eventstream.UpdateEventstreamRequest
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleEventstreamID = "5b218778-e7a5-4d73-8187-f10824047715"
	exampleUpdateEventstreamRequest = eventstream.UpdateEventstreamRequest{
		Description: to.Ptr("A new description for eventstream."),
		DisplayName: to.Ptr("Eventstream_New_Name"),
	}

	exampleRes := eventstream.Eventstream{
		Type:        to.Ptr(eventstream.ItemTypeEventstream),
		Description: to.Ptr("A new description for eventstream."),
		DisplayName: to.Ptr("Eventstream_New_Name"),
		ID:          to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	}

	testsuite.serverFactory.ItemsServer.UpdateEventstream = func(ctx context.Context, workspaceID string, eventstreamID string, updateEventstreamRequest eventstream.UpdateEventstreamRequest, options *eventstream.ItemsClientUpdateEventstreamOptions) (resp azfake.Responder[eventstream.ItemsClientUpdateEventstreamResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleEventstreamID, eventstreamID)
		testsuite.Require().True(reflect.DeepEqual(exampleUpdateEventstreamRequest, updateEventstreamRequest))
		resp = azfake.Responder[eventstream.ItemsClientUpdateEventstreamResponse]{}
		resp.SetResponse(http.StatusOK, eventstream.ItemsClientUpdateEventstreamResponse{Eventstream: exampleRes}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	res, err := client.UpdateEventstream(ctx, exampleWorkspaceID, exampleEventstreamID, exampleUpdateEventstreamRequest, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.Eventstream))
}

func (testsuite *FakeTestSuite) TestItems_DeleteEventstream() {
	// From example
	ctx := runtime.WithHTTPHeader(testsuite.ctx, map[string][]string{
		"example-id": {"Delete an eventstream example"},
	})
	var exampleWorkspaceID string
	var exampleEventstreamID string
	exampleWorkspaceID = "cfafbeb1-8037-4d0c-896e-a46fb27ff229"
	exampleEventstreamID = "5b218778-e7a5-4d73-8187-f10824047715"

	testsuite.serverFactory.ItemsServer.DeleteEventstream = func(ctx context.Context, workspaceID string, eventstreamID string, options *eventstream.ItemsClientDeleteEventstreamOptions) (resp azfake.Responder[eventstream.ItemsClientDeleteEventstreamResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleWorkspaceID, workspaceID)
		testsuite.Require().Equal(exampleEventstreamID, eventstreamID)
		resp = azfake.Responder[eventstream.ItemsClientDeleteEventstreamResponse]{}
		resp.SetResponse(http.StatusOK, eventstream.ItemsClientDeleteEventstreamResponse{}, nil)
		return
	}

	client := testsuite.clientFactory.NewItemsClient()
	_, err = client.DeleteEventstream(ctx, exampleWorkspaceID, exampleEventstreamID, nil)
	testsuite.Require().NoError(err, "Failed to get result for example ")
}
