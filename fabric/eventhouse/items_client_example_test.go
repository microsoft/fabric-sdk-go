// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package eventhouse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"

	"github.com/microsoft/fabric-sdk-go/fabric/eventhouse"
)

// Generated from example definition
func ExampleItemsClient_NewListEventhousesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := eventhouse.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewItemsClient().NewListEventhousesPager("cfafbeb1-8037-4d0c-896e-a46fb27ff229", &eventhouse.ItemsClientListEventhousesOptions{ContinuationToken: nil})
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
		// page.Eventhouses = eventhouse.Eventhouses{
		// 	Value: []eventhouse.Eventhouse{
		// 		{
		// 			Type: to.Ptr(eventhouse.ItemTypeEventhouse),
		// 			Description: to.Ptr("An eventhouse description."),
		// 			DisplayName: to.Ptr("Eventhouse_1"),
		// 			ID: to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		// 			WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		// 			Properties: &eventhouse.Properties{
		// 				DatabasesItemIDs: []string{
		// 					"f1082404-7716-5b21-8778-e7a5e7a54d73",
		// 					"8187f108-2404-4771-6e7a-5b218778e7a5"},
		// 					IngestionServiceURI: to.Ptr("https://ingest-trd-f7k1b2rzuqrjmb3wpd.z5.kusto.fabric.microsoft.com"),
		// 					MinimumConsumptionUnits: to.Ptr[float64](2.25),
		// 					QueryServiceURI: to.Ptr("https://trd-f7k1b2rzuqrjmb3wpd.z5.kusto.fabric.microsoft.com"),
		// 				},
		// 			},
		// 			{
		// 				Type: to.Ptr(eventhouse.ItemTypeEventhouse),
		// 				Description: to.Ptr("An eventhouse description."),
		// 				DisplayName: to.Ptr("Eventhouse_2"),
		// 				ID: to.Ptr("340d91b9-5a39-409c-b9c0-05ba832c476e"),
		// 				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		// 				Properties: &eventhouse.Properties{
		// 					DatabasesItemIDs: []string{
		// 						"4d738187-f108-2404-4771-6e7a5b218778"},
		// 						IngestionServiceURI: to.Ptr("https://ingest-trd-f7k1b2rzuqrjmb3wpd.z5.kusto.fabric.microsoft.com"),
		// 						MinimumConsumptionUnits: to.Ptr[float64](0),
		// 						QueryServiceURI: to.Ptr("https://trd-f7k1b2rzuqrjmb3wpd.z5.kusto.fabric.microsoft.com"),
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition
func ExampleItemsClient_GetEventhouse() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := eventhouse.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewItemsClient().GetEventhouse(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Eventhouse = eventhouse.Eventhouse{
	// 	Type: to.Ptr(eventhouse.ItemTypeEventhouse),
	// 	Description: to.Ptr("An eventhouse description."),
	// 	DisplayName: to.Ptr("Eventhouse_1"),
	// 	ID: to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
	// 	WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	// 	Properties: &eventhouse.Properties{
	// 		DatabasesItemIDs: []string{
	// 			"f1082404-7716-5b21-8778-e7a5e7a54d73",
	// 			"8187f108-2404-4771-6e7a-5b218778e7a5"},
	// 			IngestionServiceURI: to.Ptr("https://ingest-trd-f7k1b2rzuqrjmb3wpd.z5.kusto.fabric.microsoft.com"),
	// 			MinimumConsumptionUnits: to.Ptr[float64](0),
	// 			QueryServiceURI: to.Ptr("https://trd-f7k1b2rzuqrjmb3wpd.z5.kusto.fabric.microsoft.com"),
	// 		},
	// 	}
}

// Generated from example definition
func ExampleItemsClient_UpdateEventhouse() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := eventhouse.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewItemsClient().UpdateEventhouse(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", eventhouse.UpdateEventhouseRequest{
		Description: to.Ptr("A new description for eventhouse."),
		DisplayName: to.Ptr("Eventhouse_New_Name"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Eventhouse = eventhouse.Eventhouse{
	// 	Type: to.Ptr(eventhouse.ItemTypeEventhouse),
	// 	Description: to.Ptr("A new description for eventhouse."),
	// 	DisplayName: to.Ptr("Eventhouse_New_Name"),
	// 	ID: to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
	// 	WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	// }
}

// Generated from example definition
func ExampleItemsClient_DeleteEventhouse() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := eventhouse.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewItemsClient().DeleteEventhouse(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleItemsClient_BeginGetEventhouseDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := eventhouse.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewItemsClient().BeginGetEventhouseDefinition(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", &eventhouse.ItemsClientBeginGetEventhouseDefinitionOptions{Format: nil})
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
	// res.DefinitionResponse = eventhouse.DefinitionResponse{
	// 	Definition: &eventhouse.Definition{
	// 		Parts: []eventhouse.DefinitionPart{
	// 			{
	// 				Path: to.Ptr("EventhouseProperties.json"),
	// 				Payload: to.Ptr("e30="),
	// 				PayloadType: to.Ptr(eventhouse.PayloadTypeInlineBase64),
	// 			},
	// 			{
	// 				Path: to.Ptr(".platform"),
	// 				Payload: to.Ptr("ZG90UGxhdGZvcm1CYXNlNjRTdHJpbmc"),
	// 				PayloadType: to.Ptr(eventhouse.PayloadTypeInlineBase64),
	// 		}},
	// 	},
	// }
}

// Generated from example definition
func ExampleItemsClient_BeginUpdateEventhouseDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := eventhouse.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewItemsClient().BeginUpdateEventhouseDefinition(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", eventhouse.UpdateEventhouseDefinitionRequest{
		Definition: &eventhouse.Definition{
			Parts: []eventhouse.DefinitionPart{
				{
					Path:        to.Ptr("EventhouseProperties.json"),
					Payload:     to.Ptr("e30="),
					PayloadType: to.Ptr(eventhouse.PayloadTypeInlineBase64),
				},
				{
					Path:        to.Ptr(".platform"),
					Payload:     to.Ptr("ZG90UGxhdGZvcm1CYXNlNjRTdHJpbmc="),
					PayloadType: to.Ptr(eventhouse.PayloadTypeInlineBase64),
				}},
		},
	}, &eventhouse.ItemsClientBeginUpdateEventhouseDefinitionOptions{UpdateMetadata: to.Ptr(true)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
