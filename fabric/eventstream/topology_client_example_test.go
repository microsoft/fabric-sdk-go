// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package eventstream_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"

	"github.com/microsoft/fabric-sdk-go/fabric/eventstream"
)

// Generated from example definition
func ExampleTopologyClient_GetEventstreamTopology() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := eventstream.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTopologyClient().GetEventstreamTopology(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "8c500070-073f-4a88-b478-8fabe1941c52", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleTopologyClient_PauseEventstream() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := eventstream.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTopologyClient().PauseEventstream(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "8c500070-073f-4a88-b478-8fabe1941c52", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleTopologyClient_ResumeEventstream() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := eventstream.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTopologyClient().ResumeEventstream(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "8c500070-073f-4a88-b478-8fabe1941c52", eventstream.DataSourceStartRequest{
		StartType: to.Ptr(eventstream.DataSourceStartTypeNow),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleTopologyClient_GetEventstreamSource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := eventstream.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTopologyClient().GetEventstreamSource(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "8c500070-073f-4a88-b478-8fabe1941c52", "e2886002-d696-4c05-969c-51361365cc24", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleTopologyClient_GetEventstreamDestination() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := eventstream.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTopologyClient().GetEventstreamDestination(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "8c500070-073f-4a88-b478-8fabe1941c52", "2e4c91e7-0c4a-4cc4-abe3-cc7ba4310a37", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleTopologyClient_GetEventstreamSourceConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := eventstream.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTopologyClient().GetEventstreamSourceConnection(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "8c500070-073f-4a88-b478-8fabe1941c52", "f344e2e0-e846-4991-ac26-d27dfb6a73c2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleTopologyClient_GetEventstreamDestinationConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := eventstream.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTopologyClient().GetEventstreamDestinationConnection(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "8c500070-073f-4a88-b478-8fabe1941c52", "2e4c91e7-0c4a-4cc4-abe3-cc7ba4310a37", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleTopologyClient_PauseEventstreamSource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := eventstream.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTopologyClient().PauseEventstreamSource(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "8c500070-073f-4a88-b478-8fabe1941c52", "e2886002-d696-4c05-969c-51361365cc24", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleTopologyClient_ResumeEventstreamSource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := eventstream.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTopologyClient().ResumeEventstreamSource(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "8c500070-073f-4a88-b478-8fabe1941c52", "e2886002-d696-4c05-969c-51361365cc24", eventstream.DataSourceStartRequest{
		StartType: to.Ptr(eventstream.DataSourceStartTypeWhenLastStopped),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleTopologyClient_PauseEventstreamDestination() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := eventstream.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTopologyClient().PauseEventstreamDestination(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "8c500070-073f-4a88-b478-8fabe1941c52", "2e4c91e7-0c4a-4cc4-abe3-cc7ba4310a37", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleTopologyClient_ResumeEventstreamDestination() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := eventstream.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTopologyClient().ResumeEventstreamDestination(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "8c500070-073f-4a88-b478-8fabe1941c52", "2e4c91e7-0c4a-4cc4-abe3-cc7ba4310a37", eventstream.DataSourceStartRequest{
		CustomStartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-23T16:22:20.000Z"); return t }()),
		StartType:           to.Ptr(eventstream.DataSourceStartTypeCustomTime),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
