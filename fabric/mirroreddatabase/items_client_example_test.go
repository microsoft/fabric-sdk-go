// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package mirroreddatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"

	"github.com/microsoft/fabric-sdk-go/fabric/mirroreddatabase"
)

// Generated from example definition
func ExampleItemsClient_NewListMirroredDatabasesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := mirroreddatabase.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewItemsClient().NewListMirroredDatabasesPager("cfafbeb1-8037-4d0c-896e-a46fb27ff229", &mirroreddatabase.ItemsClientListMirroredDatabasesOptions{ContinuationToken: nil})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.MirroredDatabases = mirroreddatabase.MirroredDatabases{
		// 	Value: []mirroreddatabase.MirroredDatabase{
		// 		{
		// 			Type: to.Ptr(mirroreddatabase.ItemTypeMirroredDatabase),
		// 			Description: to.Ptr("A mirrored database description."),
		// 			DisplayName: to.Ptr("Mirrored database 1"),
		// 			ID: to.Ptr("3546052c-ae64-4526-b1a8-52af7761426f"),
		// 			WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		// 			Properties: &mirroreddatabase.Properties{
		// 				DefaultSchema: to.Ptr("dbo"),
		// 				OneLakeTablesPath: to.Ptr("https://onelake.dfs.fabric.microsoft.com/7ba2c2f8-750a-47d8-9e95-99585107e23a/c3889489-6074-490e-90ed-e5a9e46142c6/Tables"),
		// 				SQLEndpointProperties: &mirroreddatabase.SQLEndpointProperties{
		// 					ConnectionString: to.Ptr("x6eps4xrq2xudenlfv6naeo3i4-7dbke6ykoxmephuvtfmfcb7chi.msit-datawarehouse.fabric.microsoft.com"),
		// 					ID: to.Ptr("84c9ad4a-ba9e-42d9-9025-b40f8e38c025"),
		// 					ProvisioningStatus: to.Ptr(mirroreddatabase.SQLEndpointProvisioningStatusSuccess),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleItemsClient_CreateMirroredDatabase() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := mirroreddatabase.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewItemsClient().CreateMirroredDatabase(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", mirroreddatabase.CreateMirroredDatabaseRequest{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleItemsClient_GetMirroredDatabase() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := mirroreddatabase.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewItemsClient().GetMirroredDatabase(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MirroredDatabase = mirroreddatabase.MirroredDatabase{
	// 	Type: to.Ptr(mirroreddatabase.ItemTypeMirroredDatabase),
	// 	Description: to.Ptr("A mirrored database description."),
	// 	DisplayName: to.Ptr("Mirrored database 1"),
	// 	ID: to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
	// 	WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	// 	Properties: &mirroreddatabase.Properties{
	// 		DefaultSchema: to.Ptr("dbo"),
	// 		OneLakeTablesPath: to.Ptr("https://onelake.dfs.fabric.microsoft.com/cfafbeb1-8037-4d0c-896e-a46fb27ff229/5b218778-e7a5-4d73-8187-f10824047715/Tables"),
	// 		SQLEndpointProperties: &mirroreddatabase.SQLEndpointProperties{
	// 			ConnectionString: to.Ptr("xxx.datawarehouse.fabric.microsoft.com"),
	// 			ID: to.Ptr("84c9ad4a-ba9e-42d9-9025-b40f8e38c025"),
	// 			ProvisioningStatus: to.Ptr(mirroreddatabase.SQLEndpointProvisioningStatusSuccess),
	// 		},
	// 	},
	// }
}

// Generated from example definition
func ExampleItemsClient_UpdateMirroredDatabase() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := mirroreddatabase.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewItemsClient().UpdateMirroredDatabase(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", mirroreddatabase.UpdateMirroredDatabaseRequest{
		Description: to.Ptr("A new description for mirrored database."),
		DisplayName: to.Ptr("MirroredDatabase's New name"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MirroredDatabase = mirroreddatabase.MirroredDatabase{
	// 	Type: to.Ptr(mirroreddatabase.ItemTypeMirroredDatabase),
	// 	Description: to.Ptr("A new description for mirrored database."),
	// 	DisplayName: to.Ptr("MirroredDatabase's New name"),
	// 	ID: to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
	// 	WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	// }
}

// Generated from example definition
func ExampleItemsClient_DeleteMirroredDatabase() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := mirroreddatabase.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewItemsClient().DeleteMirroredDatabase(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleItemsClient_GetMirroredDatabaseDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := mirroreddatabase.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewItemsClient().GetMirroredDatabaseDefinition(ctx, "6e335e92-a2a2-4b5a-970a-bd6a89fbb765", "cfafbeb1-8037-4d0c-896e-a46fb27ff229", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DefinitionResponse = mirroreddatabase.DefinitionResponse{
	// 	Definition: &mirroreddatabase.Definition{
	// 		Parts: []mirroreddatabase.DefinitionPart{
	// 			{
	// 				Path: to.Ptr("mirroredDatabase.json"),
	// 				Payload: to.Ptr("eyAicHJvcGVydGllcy..WJsZSIgfSB9IH0gXSB9IH0"),
	// 				PayloadType: to.Ptr(mirroreddatabase.PayloadTypeInlineBase64),
	// 			},
	// 			{
	// 				Path: to.Ptr(".platform"),
	// 				Payload: to.Ptr("ZG90UGxhdGZvcm1CYXNlNjRTdHJpbmc="),
	// 				PayloadType: to.Ptr(mirroreddatabase.PayloadTypeInlineBase64),
	// 		}},
	// 	},
	// }
}

// Generated from example definition
func ExampleItemsClient_UpdateMirroredDatabaseDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := mirroreddatabase.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewItemsClient().UpdateMirroredDatabaseDefinition(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", mirroreddatabase.UpdateMirroredDatabaseDefinitionRequest{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
