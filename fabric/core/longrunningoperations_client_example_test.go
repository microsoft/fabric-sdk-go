// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package core_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"

	"github.com/microsoft/fabric-sdk-go/fabric/core"
)

// Generated from example definition
func ExampleLongRunningOperationsClient_GetOperationState_getActiveLongRunningOperationExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLongRunningOperationsClient().GetOperationState(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff227", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationState = core.OperationState{
	// 	CreatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-13T14:56:18.477Z"); return t}()),
	// 	LastUpdatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-13T15:01:10.532Z"); return t}()),
	// 	PercentComplete: to.Ptr[int32](25),
	// 	Status: to.Ptr(core.LongRunningOperationStatusRunning),
	// }
}

// Generated from example definition
func ExampleLongRunningOperationsClient_GetOperationState_getCompletedLongRunningOperationExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLongRunningOperationsClient().GetOperationState(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff227", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationState = core.OperationState{
	// 	CreatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-13T14:56:18.477Z"); return t}()),
	// 	LastUpdatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-13T15:01:10.532Z"); return t}()),
	// 	PercentComplete: to.Ptr[int32](100),
	// 	Status: to.Ptr(core.LongRunningOperationStatusSucceeded),
	// }
}

// Generated from example definition
func ExampleLongRunningOperationsClient_GetOperationResult() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewLongRunningOperationsClient().GetOperationResult(ctx, "431e8d7b-4a95-4c02-8ccd-6faef5ba1bd7", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
