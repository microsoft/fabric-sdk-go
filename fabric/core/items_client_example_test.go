// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package core_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"

	"github.com/microsoft/fabric-sdk-go/fabric/core"
)

// Generated from example definition
func ExampleItemsClient_NewListItemsPager_listItemInWorkspaceWithContinuationExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewItemsClient().NewListItemsPager("cfafbeb1-8037-4d0c-896e-a46fb27ff229", &core.ItemsClientListItemsOptions{Type: nil,
		ContinuationToken: nil,
	})
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
		// page.Items = core.Items{
		// 	ContinuationToken: to.Ptr("LDEsMTAwMDAwLDA%3D"),
		// 	ContinuationURI: to.Ptr("https://api.fabric.microsoft.com/v1/workspaces/cfafbeb1-8037-4d0c-896e-a46fb27ff229/items?continuationToken=LDEsMTAwMDAwLDA%3D"),
		// 	Value: []core.Item{
		// 		{
		// 			Type: to.Ptr(core.ItemTypeLakehouse),
		// 			Description: to.Ptr("A lakehouse used by the analytics team."),
		// 			DisplayName: to.Ptr("Lakehouse"),
		// 			ID: to.Ptr("3546052c-ae64-4526-b1a8-52af7761426f"),
		// 			WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		// 		},
		// 		{
		// 			Type: to.Ptr(core.ItemType("KustoDashboard")),
		// 			Description: to.Ptr("A notebook for refining medical data analysis through machine learning algorithms."),
		// 			DisplayName: to.Ptr("Notebook"),
		// 			ID: to.Ptr("58fa1eac-9694-4a6b-ba25-3520288e8fea"),
		// 			WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleItemsClient_NewListItemsPager_listItemsInWorkspaceByTypeQueryParameterExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewItemsClient().NewListItemsPager("cfafbeb1-8037-4d0c-896e-a46fb27ff229", &core.ItemsClientListItemsOptions{Type: to.Ptr("Lakehouse"),
		ContinuationToken: nil,
	})
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
		// page.Items = core.Items{
		// 	Value: []core.Item{
		// 		{
		// 			Type: to.Ptr(core.ItemTypeLakehouse),
		// 			Description: to.Ptr("A lakehouse used by the analytics team."),
		// 			DisplayName: to.Ptr("Lakehouse Name 1"),
		// 			ID: to.Ptr("3546052c-ae64-4526-b1a8-52af7761426f"),
		// 			WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleItemsClient_NewListItemsPager_listItemsInWorkspaceExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewItemsClient().NewListItemsPager("cfafbeb1-8037-4d0c-896e-a46fb27ff229", &core.ItemsClientListItemsOptions{Type: nil,
		ContinuationToken: nil,
	})
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
		// page.Items = core.Items{
		// 	Value: []core.Item{
		// 		{
		// 			Type: to.Ptr(core.ItemTypeLakehouse),
		// 			Description: to.Ptr("A lakehouse used by the analytics team."),
		// 			DisplayName: to.Ptr("Lakehouse"),
		// 			ID: to.Ptr("3546052c-ae64-4526-b1a8-52af7761426f"),
		// 			WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		// 		},
		// 		{
		// 			Type: to.Ptr(core.ItemType("KustoDashboard")),
		// 			Description: to.Ptr("A notebook for refining medical data analysis through machine learning algorithms."),
		// 			DisplayName: to.Ptr("Notebook"),
		// 			ID: to.Ptr("58fa1eac-9694-4a6b-ba25-3520288e8fea"),
		// 			WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleItemsClient_BeginCreateItem() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewItemsClient().BeginCreateItem(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", core.CreateItemRequest{
		Type:        to.Ptr(core.ItemTypeLakehouse),
		DisplayName: to.Ptr("Item 1"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition
func ExampleItemsClient_GetItem() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewItemsClient().GetItem(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Item = core.Item{
	// 	Type: to.Ptr(core.ItemTypeLakehouse),
	// 	Description: to.Ptr("Item 1 description"),
	// 	DisplayName: to.Ptr("Item 1"),
	// 	ID: to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
	// 	WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	// }
}

// Generated from example definition
func ExampleItemsClient_UpdateItem() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewItemsClient().UpdateItem(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", core.UpdateItemRequest{
		Description: to.Ptr("Item's New description"),
		DisplayName: to.Ptr("Item's New name"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Item = core.Item{
	// 	Type: to.Ptr(core.ItemTypeLakehouse),
	// 	Description: to.Ptr("Item's New description"),
	// 	DisplayName: to.Ptr("Item's New name"),
	// 	ID: to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
	// 	WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	// }
}

// Generated from example definition
func ExampleItemsClient_DeleteItem() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewItemsClient().DeleteItem(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleItemsClient_BeginGetItemDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewItemsClient().BeginGetItemDefinition(ctx, "6e335e92-a2a2-4b5a-970a-bd6a89fbb765", "cfafbeb1-8037-4d0c-896e-a46fb27ff229", &core.ItemsClientBeginGetItemDefinitionOptions{Format: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ItemDefinitionResponse = core.ItemDefinitionResponse{
	// 	Definition: &core.ItemDefinition{
	// 		Parts: []core.ItemDefinitionPart{
	// 			{
	// 				Path: to.Ptr("report.json"),
	// 				Payload: to.Ptr("QmFzZTY0U3RyaW5n"),
	// 				PayloadType: to.Ptr(core.PayloadTypeInlineBase64),
	// 			},
	// 			{
	// 				Path: to.Ptr("definition.pbir"),
	// 				Payload: to.Ptr("QW5vdGhlckJhc2U2NFN0cmluZw"),
	// 				PayloadType: to.Ptr(core.PayloadTypeInlineBase64),
	// 			},
	// 			{
	// 				Path: to.Ptr(".platform"),
	// 				Payload: to.Ptr("ZG90UGxhdGZvcm1CYXNlNjRTdHJpbmc="),
	// 				PayloadType: to.Ptr(core.PayloadTypeInlineBase64),
	// 		}},
	// 	},
	// }
}

// Generated from example definition
func ExampleItemsClient_BeginUpdateItemDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewItemsClient().BeginUpdateItemDefinition(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", core.UpdateItemDefinitionRequest{
		Definition: &core.ItemDefinition{
			Parts: []core.ItemDefinitionPart{
				{
					Path:        to.Ptr("report.json"),
					Payload:     to.Ptr("QmFzZTY0U3RyaW5n"),
					PayloadType: to.Ptr(core.PayloadTypeInlineBase64),
				},
				{
					Path:        to.Ptr("definition.pbir"),
					Payload:     to.Ptr("QW5vdGhlckJhc2U2NFN0cmluZw"),
					PayloadType: to.Ptr(core.PayloadTypeInlineBase64),
				},
				{
					Path:        to.Ptr(".platform"),
					Payload:     to.Ptr("ZG90UGxhdGZvcm1CYXNlNjRTdHJpbmc="),
					PayloadType: to.Ptr(core.PayloadTypeInlineBase64),
				}},
		},
	}, &core.ItemsClientBeginUpdateItemDefinitionOptions{UpdateMetadata: to.Ptr(true)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition
func ExampleItemsClient_NewListItemConnectionsPager_listItemConnectionsSemanticModelDirectLakeExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewItemsClient().NewListItemConnectionsPager("4b218778-e7a5-4d73-8187-f10824047715", "431e8d7b-4a95-4c02-8ccd-6faef5ba1bd7", &core.ItemsClientListItemConnectionsOptions{ContinuationToken: nil})
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
		// page.ItemConnections = core.ItemConnections{
		// 	Value: []core.ItemConnection{
		// 		{
		// 			ConnectionDetails: &core.ListConnectionDetails{
		// 				Type: to.Ptr("SQL"),
		// 				Path: to.Ptr("xqoruksalslrtkdxe2bvrlwgsi-5j2iqepw1i3ucdp6bepe62hcii.datawarehouse.fabric.microsoft.com;532183f5-ac60-4d12-0fc5-4094532f14b5"),
		// 			},
		// 			ConnectivityType: to.Ptr(core.ConnectivityType("DefaultSetting")),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleItemsClient_NewListItemConnectionsPager_listItemConnectionsExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewItemsClient().NewListItemConnectionsPager("4b218778-e7a5-4d73-8187-f10824047715", "431e8d7b-4a95-4c02-8ccd-6faef5ba1bd7", &core.ItemsClientListItemConnectionsOptions{ContinuationToken: nil})
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
		// page.ItemConnections = core.ItemConnections{
		// 	Value: []core.ItemConnection{
		// 		{
		// 			ConnectionDetails: &core.ListConnectionDetails{
		// 				Type: to.Ptr("Web"),
		// 				Path: to.Ptr("https://www.contoso.com"),
		// 			},
		// 			ConnectivityType: to.Ptr(core.ConnectivityTypeShareableCloud),
		// 			DisplayName: to.Ptr("ContosoConnection1"),
		// 			ID: to.Ptr("6952a7b2-aea3-414f-9d85-6c0fe5d34539"),
		// 		},
		// 		{
		// 			ConnectionDetails: &core.ListConnectionDetails{
		// 				Type: to.Ptr("SQL"),
		// 				Path: to.Ptr("contoso.database.windows.net;sales"),
		// 			},
		// 			ConnectivityType: to.Ptr(core.ConnectivityType("OnPremisesDataGateway")),
		// 			DisplayName: to.Ptr("ContosoConnection2"),
		// 			GatewayID: to.Ptr("58376c10-5f61-4024-887e-748df4beae45"),
		// 			ID: to.Ptr("0b9af1bd-e974-4893-8947-d89d5a560385"),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleItemsClient_NewListItemConnectionsPager_listItemConnectionsWithContinuationExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewItemsClient().NewListItemConnectionsPager("4b218778-e7a5-4d73-8187-f10824047715", "431e8d7b-4a95-4c02-8ccd-6faef5ba1bd7", &core.ItemsClientListItemConnectionsOptions{ContinuationToken: nil})
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
		// page.ItemConnections = core.ItemConnections{
		// 	Value: []core.ItemConnection{
		// 		{
		// 			ConnectionDetails: &core.ListConnectionDetails{
		// 				Type: to.Ptr("SQL"),
		// 				Path: to.Ptr("xqoruksalslrtkdxe2bvrlwgsi-5j2iqepw1i3ucdp6bepe62hcii.datawarehouse.fabric.microsoft.com;532183f5-ac60-4d12-0fc5-4094532f14b5"),
		// 			},
		// 			ConnectivityType: to.Ptr(core.ConnectivityType("DefaultSetting")),
		// 	}},
		// }
	}
}