// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package admin_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"

	"github.com/microsoft/fabric-sdk-go/fabric/admin"
)

// Generated from example definition
func ExampleItemsClient_NewListItemsPager_getAListOfItemsUsingTypeQueryParameterExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := admin.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewItemsClient().NewListItemsPager(&admin.ItemsClientListItemsOptions{WorkspaceID: nil,
		CapacityID:        nil,
		State:             nil,
		Type:              to.Ptr("Report"),
		ContinuationToken: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.ItemEntities {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.Items = admin.Items{
		// 	ItemEntities: []admin.Item{
		// 		{
		// 			Name: to.Ptr("Test Report"),
		// 			Type: to.Ptr(admin.ItemTypeReport),
		// 			CapacityID: to.Ptr("D5E336D6-D919-4ECC-B424-1F771A506851"),
		// 			CreatorPrincipal: &admin.Principal{
		// 				Type: to.Ptr(admin.PrincipalTypeUser),
		// 				DisplayName: to.Ptr("Jacob Hancock"),
		// 				ID: to.Ptr("f3052d1c-61a9-46fb-8df9-0d78916ae041"),
		// 				UserDetails: &admin.PrincipalUserDetails{
		// 					UserPrincipalName: to.Ptr("Jacob@example.com"),
		// 				},
		// 			},
		// 			ID: to.Ptr("b1a7e572-2585-4650-98ae-b92356f4460b"),
		// 			LastUpdatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-06-27T16:55:04.893Z"); return t}()),
		// 			State: to.Ptr(admin.ItemStateActive),
		// 			WorkspaceID: to.Ptr("7f4496db-9929-47bd-89c0-d7eb2f517a98"),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleItemsClient_NewListItemsPager_getAllItemsInTheTenantExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := admin.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewItemsClient().NewListItemsPager(&admin.ItemsClientListItemsOptions{WorkspaceID: nil,
		CapacityID:        nil,
		State:             nil,
		Type:              nil,
		ContinuationToken: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.ItemEntities {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.Items = admin.Items{
		// 	ContinuationToken: to.Ptr("MSwxMDAwMCww"),
		// 	ContinuationURI: to.Ptr("https://api.fabric.microsoft.com/v1/admin/items?continuationToken=MSwxMDAwMCww"),
		// 	ItemEntities: []admin.Item{
		// 		{
		// 			Name: to.Ptr("Test"),
		// 			Type: to.Ptr(admin.ItemTypeNotebook),
		// 			Description: to.Ptr("Test notebook."),
		// 			CapacityID: to.Ptr("D5E336D6-D919-4ECC-B424-1F771A506851"),
		// 			CreatorPrincipal: &admin.Principal{
		// 				Type: to.Ptr(admin.PrincipalTypeUser),
		// 				DisplayName: to.Ptr("Caleb Foster"),
		// 				ID: to.Ptr("f3052d1c-61a9-46fb-8df9-0d78916ae041"),
		// 				UserDetails: &admin.PrincipalUserDetails{
		// 					UserPrincipalName: to.Ptr("caleb@example.com"),
		// 				},
		// 			},
		// 			ID: to.Ptr("17d8929d-ab32-46d1-858b-fdea74e93bff"),
		// 			LastUpdatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-29T17:47:29.986Z"); return t}()),
		// 			State: to.Ptr(admin.ItemStateActive),
		// 			WorkspaceID: to.Ptr("7f4496db-9929-47bd-89c0-d7eb2f517a95"),
		// 		},
		// 		{
		// 			Name: to.Ptr("TestKusto"),
		// 			Type: to.Ptr(admin.ItemTypeKQLDatabase),
		// 			Description: to.Ptr("Test KQL database."),
		// 			CapacityID: to.Ptr("D5E336D6-D919-4ECC-B424-1F881A506851"),
		// 			CreatorPrincipal: &admin.Principal{
		// 				Type: to.Ptr(admin.PrincipalTypeUser),
		// 				DisplayName: to.Ptr("Jacob Hancock"),
		// 				ID: to.Ptr("f3052d1c-61a9-46fb-8df9-0d78916ae041"),
		// 				UserDetails: &admin.PrincipalUserDetails{
		// 					UserPrincipalName: to.Ptr("jacob@example.com"),
		// 				},
		// 			},
		// 			ID: to.Ptr("37d8929d-ab32-46d1-858b-fdea74e93bff"),
		// 			LastUpdatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-06-30T17:47:29.986Z"); return t}()),
		// 			State: to.Ptr(admin.ItemStateActive),
		// 			WorkspaceID: to.Ptr("8f4496db-9929-47bd-89c0-d7eb2f517a95"),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleItemsClient_NewListItemsPager_getListOfDatamartsUsingTypeQueryParameterExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := admin.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewItemsClient().NewListItemsPager(&admin.ItemsClientListItemsOptions{WorkspaceID: nil,
		CapacityID:        nil,
		State:             nil,
		Type:              to.Ptr("Lakehouse"),
		ContinuationToken: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.ItemEntities {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.Items = admin.Items{
		// 	ItemEntities: []admin.Item{
		// 		{
		// 			Name: to.Ptr("Lakehouse 2022-03-16T21:42:38.442Z"),
		// 			Type: to.Ptr(admin.ItemTypeLakehouse),
		// 			CapacityID: to.Ptr("D5E336D6-D919-4ECC-B424-1F771A506851"),
		// 			CreatorPrincipal: &admin.Principal{
		// 				Type: to.Ptr(admin.PrincipalTypeUser),
		// 				DisplayName: to.Ptr("Jacob Hancock"),
		// 				ID: to.Ptr("f3052d1c-61a9-46fb-8df9-0d78916ae041"),
		// 				UserDetails: &admin.PrincipalUserDetails{
		// 					UserPrincipalName: to.Ptr("jacob@example.com"),
		// 				},
		// 			},
		// 			ID: to.Ptr("b1a7e572-2585-4650-98ae-b92356f4460b"),
		// 			LastUpdatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-06-27T16:55:04.893Z"); return t}()),
		// 			State: to.Ptr(admin.ItemStateActive),
		// 			WorkspaceID: to.Ptr("7f4496db-9929-47bd-89c0-d7eb2f517a98"),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleItemsClient_GetItem_getItemDetailsByIdAndTypeExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := admin.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewItemsClient().GetItem(ctx, "7f4496db-9929-47bd-89c0-d7eb2f517a98", "f089354e-8366-4e18-aea3-4cb4a3a50b48", &admin.ItemsClientGetItemOptions{Type: to.Ptr("Report")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Item = admin.Item{
	// 	Name: to.Ptr("Test"),
	// 	Type: to.Ptr(admin.ItemTypeReport),
	// 	Description: to.Ptr("Test Jacob's Report,"),
	// 	CapacityID: to.Ptr("D5E336D6-D919-4ECC-B424-1F771A506851"),
	// 	CreatorPrincipal: &admin.Principal{
	// 		Type: to.Ptr(admin.PrincipalTypeUser),
	// 		DisplayName: to.Ptr("Jacob Hancock"),
	// 		ID: to.Ptr("f3052d1c-61a9-46fb-8df9-0d78916ae041"),
	// 		UserDetails: &admin.PrincipalUserDetails{
	// 			UserPrincipalName: to.Ptr("Jacob@example.com"),
	// 		},
	// 	},
	// 	ID: to.Ptr("17d8929d-ab32-46d1-858b-fdea74e93bf2"),
	// 	LastUpdatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-29T17:47:29.986Z"); return t}()),
	// 	State: to.Ptr(admin.ItemStateActive),
	// 	WorkspaceID: to.Ptr("7f4496db-9929-47bd-89c0-d7eb2f517a98"),
	// }
}

// Generated from example definition
func ExampleItemsClient_GetItem_getItemDetailsByIdExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := admin.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewItemsClient().GetItem(ctx, "7f4496db-9929-47bd-89c0-d7eb2f517a98", "f089354e-8366-4e18-aea3-4cb4a3a50b48", &admin.ItemsClientGetItemOptions{Type: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Item = admin.Item{
	// 	Name: to.Ptr("Test"),
	// 	Type: to.Ptr(admin.ItemType("Kusto")),
	// 	Description: to.Ptr("Test Jacob's notebook."),
	// 	CapacityID: to.Ptr("D5E336D6-D919-4ECC-B424-1F771A506851"),
	// 	CreatorPrincipal: &admin.Principal{
	// 		Type: to.Ptr(admin.PrincipalTypeUser),
	// 		DisplayName: to.Ptr("Jacob Hancock"),
	// 		ID: to.Ptr("f3052d1c-61a9-46fb-8df9-0d78916ae041"),
	// 		UserDetails: &admin.PrincipalUserDetails{
	// 			UserPrincipalName: to.Ptr("Jacob@example.com"),
	// 		},
	// 	},
	// 	ID: to.Ptr("17d8929d-ab32-46d1-858b-fdea74e93bf2"),
	// 	LastUpdatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-29T17:47:29.986Z"); return t}()),
	// 	State: to.Ptr(admin.ItemStateActive),
	// 	WorkspaceID: to.Ptr("7f4496db-9929-47bd-89c0-d7eb2f517a98"),
	// }
}

// Generated from example definition
func ExampleItemsClient_ListItemAccessDetails_listOfUsersForGivenItemIdAndTypeExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := admin.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewItemsClient().ListItemAccessDetails(ctx, "7f4496db-9929-47bd-89c0-d7eb2f517a98", "f089354e-8366-4e18-aea3-4cb4a3a50b48", &admin.ItemsClientListItemAccessDetailsOptions{Type: to.Ptr("Report")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ItemAccessDetailsResponse = admin.ItemAccessDetailsResponse{
	// 	AccessDetails: []admin.ItemAccessDetails{
	// 		{
	// 			ItemAccessDetails: &admin.ItemAccessDetail{
	// 				Type: to.Ptr(admin.ItemTypeReport),
	// 				AdditionalPermissions: []string{
	// 					"ReadAll"},
	// 					Permissions: []admin.ItemPermissions{
	// 						admin.ItemPermissionsRead,
	// 						admin.ItemPermissionsReshare},
	// 					},
	// 					Principal: &admin.Principal{
	// 						Type: to.Ptr(admin.PrincipalTypeUser),
	// 						DisplayName: to.Ptr("Jacob Hancock"),
	// 						ID: to.Ptr("f3052d1c-61a9-46fb-8df9-0d78916ae041"),
	// 						UserDetails: &admin.PrincipalUserDetails{
	// 							UserPrincipalName: to.Ptr("jacob@example.com"),
	// 						},
	// 					},
	// 			}},
	// 		}
}

// Generated from example definition
func ExampleItemsClient_ListItemAccessDetails_listOfUsersForGivenItemIdExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := admin.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewItemsClient().ListItemAccessDetails(ctx, "7f4496db-9929-47bd-89c0-d7eb2f517a98", "f089354e-8366-4e18-aea3-4cb4a3a50b48", &admin.ItemsClientListItemAccessDetailsOptions{Type: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ItemAccessDetailsResponse = admin.ItemAccessDetailsResponse{
	// 	AccessDetails: []admin.ItemAccessDetails{
	// 		{
	// 			ItemAccessDetails: &admin.ItemAccessDetail{
	// 				Type: to.Ptr(admin.ItemTypeNotebook),
	// 				AdditionalPermissions: []string{
	// 					"ReadAll",
	// 					"viewOutput"},
	// 					Permissions: []admin.ItemPermissions{
	// 						admin.ItemPermissionsRead,
	// 						admin.ItemPermissionsReshare},
	// 					},
	// 					Principal: &admin.Principal{
	// 						Type: to.Ptr(admin.PrincipalTypeUser),
	// 						DisplayName: to.Ptr("Jacob Hancock"),
	// 						ID: to.Ptr("f3052d1c-61a9-46fb-8df9-0d78916ae041"),
	// 						UserDetails: &admin.PrincipalUserDetails{
	// 							UserPrincipalName: to.Ptr("jacob@example.com"),
	// 						},
	// 					},
	// 				},
	// 				{
	// 					ItemAccessDetails: &admin.ItemAccessDetail{
	// 						Type: to.Ptr(admin.ItemTypeNotebook),
	// 						AdditionalPermissions: []string{
	// 							"ReadAll"},
	// 							Permissions: []admin.ItemPermissions{
	// 								admin.ItemPermissionsRead,
	// 								admin.ItemPermissionsReshare,
	// 								admin.ItemPermissionsExplore},
	// 							},
	// 							Principal: &admin.Principal{
	// 								Type: to.Ptr(admin.PrincipalTypeUser),
	// 								DisplayName: to.Ptr("Eric Solomon"),
	// 								ID: to.Ptr("c7db8e03-c8cb-4d4c-9f64-1dcd327c9d3c"),
	// 								UserDetails: &admin.PrincipalUserDetails{
	// 									UserPrincipalName: to.Ptr("eric@example.com"),
	// 								},
	// 							},
	// 						},
	// 						{
	// 							ItemAccessDetails: &admin.ItemAccessDetail{
	// 								Type: to.Ptr(admin.ItemTypeNotebook),
	// 								AdditionalPermissions: []string{
	// 								},
	// 								Permissions: []admin.ItemPermissions{
	// 									admin.ItemPermissionsRead,
	// 									admin.ItemPermissionsReshare},
	// 								},
	// 								Principal: &admin.Principal{
	// 									Type: to.Ptr(admin.PrincipalTypeGroup),
	// 									DisplayName: to.Ptr("TestSecurityGroup"),
	// 									GroupDetails: &admin.PrincipalGroupDetails{
	// 										GroupType: to.Ptr(admin.GroupTypeSecurityGroup),
	// 									},
	// 									ID: to.Ptr("f51b705f-a409-4d40-9197-c5d5f349e2f0"),
	// 								},
	// 						}},
	// 					}
}