// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package kqlqueryset_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"

	"github.com/microsoft/fabric-sdk-go/fabric/kqlqueryset"
)

// Generated from example definition
func ExampleItemsClient_NewListKQLQuerysetsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := kqlqueryset.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewItemsClient().NewListKQLQuerysetsPager("cfafbeb1-8037-4d0c-896e-a46fb27ff229", &kqlqueryset.ItemsClientListKQLQuerysetsOptions{ContinuationToken: nil})
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
		// page.KQLQuerysets = kqlqueryset.KQLQuerysets{
		// 	Value: []kqlqueryset.KQLQueryset{
		// 		{
		// 			Type: to.Ptr(kqlqueryset.ItemTypeKQLQueryset),
		// 			Description: to.Ptr("A KQL queryset description."),
		// 			DisplayName: to.Ptr("KQLQueryset_1"),
		// 			ID: to.Ptr("3546052c-ae64-4526-b1a8-52af7761426f"),
		// 			WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		// 		},
		// 		{
		// 			Type: to.Ptr(kqlqueryset.ItemTypeKQLQueryset),
		// 			Description: to.Ptr("A KQL queryset description."),
		// 			DisplayName: to.Ptr("KQLQueryset_2"),
		// 			ID: to.Ptr("4c9adc22-ffb1-491f-baaa-9c9987745591"),
		// 			WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		// 		},
		// 		{
		// 			Type: to.Ptr(kqlqueryset.ItemTypeKQLQueryset),
		// 			Description: to.Ptr("A KQL queryset description."),
		// 			DisplayName: to.Ptr("KQLQueryset_3"),
		// 			ID: to.Ptr("8b681594-894d-4adf-8ae8-aed415dd1de6"),
		// 			WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleItemsClient_GetKQLQueryset() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := kqlqueryset.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewItemsClient().GetKQLQueryset(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.KQLQueryset = kqlqueryset.KQLQueryset{
	// 	Type: to.Ptr(kqlqueryset.ItemTypeKQLQueryset),
	// 	Description: to.Ptr("A KQL queryset description."),
	// 	DisplayName: to.Ptr("KQLQueryset_1"),
	// 	ID: to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
	// 	WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	// }
}

// Generated from example definition
func ExampleItemsClient_UpdateKQLQueryset() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := kqlqueryset.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewItemsClient().UpdateKQLQueryset(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", kqlqueryset.UpdateKQLQuerysetRequest{
		Description: to.Ptr("A new description for KQL queryset."),
		DisplayName: to.Ptr("KQLQueryset_New_Name"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.KQLQueryset = kqlqueryset.KQLQueryset{
	// 	Type: to.Ptr(kqlqueryset.ItemTypeKQLQueryset),
	// 	Description: to.Ptr("A new description for KQL queryset."),
	// 	DisplayName: to.Ptr("KQLQueryset_New_Name"),
	// 	ID: to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
	// 	WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	// }
}

// Generated from example definition
func ExampleItemsClient_DeleteKQLQueryset() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := kqlqueryset.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewItemsClient().DeleteKQLQueryset(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}