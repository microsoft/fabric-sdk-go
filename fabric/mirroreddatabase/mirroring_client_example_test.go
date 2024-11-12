// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package mirroreddatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"

	"github.com/microsoft/fabric-sdk-go/fabric/mirroreddatabase"
)

// Generated from example definition
func ExampleMirroringClient_StartMirroring() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := mirroreddatabase.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewMirroringClient().StartMirroring(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleMirroringClient_StopMirroring() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := mirroreddatabase.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewMirroringClient().StopMirroring(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleMirroringClient_GetMirroringStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := mirroreddatabase.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMirroringClient().GetMirroringStatus(ctx, "6e335e92-a2a2-4b5a-970a-bd6a89fbb765", "cfafbeb1-8037-4d0c-896e-a46fb27ff229", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MirroringStatusResponse = mirroreddatabase.MirroringStatusResponse{
	// 	Status: to.Ptr(mirroreddatabase.MirroringStatusRunning),
	// }
}

// Generated from example definition
func ExampleMirroringClient_GetTablesMirroringStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := mirroreddatabase.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMirroringClient().GetTablesMirroringStatus(ctx, "6e335e92-a2a2-4b5a-970a-bd6a89fbb765", "cfafbeb1-8037-4d0c-896e-a46fb27ff229", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TablesMirroringStatusResponse = mirroreddatabase.TablesMirroringStatusResponse{
	// 	Data: []mirroreddatabase.TableMirroringStatusResponse{
	// 		{
	// 			Metrics: &mirroreddatabase.TableMirroringMetrics{
	// 				LastSyncDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-08T05:07:11.066Z"); return t}()),
	// 				ProcessedBytes: to.Ptr[int64](1247),
	// 				ProcessedRows: to.Ptr[int64](6),
	// 			},
	// 			SourceSchemaName: to.Ptr("dbo"),
	// 			SourceTableName: to.Ptr("test"),
	// 			Status: to.Ptr(mirroreddatabase.TableMirroringStatusReplicating),
	// 	}},
	// }
}