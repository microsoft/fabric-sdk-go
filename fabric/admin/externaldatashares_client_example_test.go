// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package admin_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"

	"github.com/microsoft/fabric-sdk-go/fabric/admin"
)

// Generated from example definition
func ExampleExternalDataSharesClient_NewListExternalDataSharesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := admin.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExternalDataSharesClient().NewListExternalDataSharesPager(&admin.ExternalDataSharesClientListExternalDataSharesOptions{ContinuationToken: nil})
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
		// page.ExternalDataShares = admin.ExternalDataShares{
		// 	Value: []admin.ExternalDataShare{
		// 		{
		// 			CreatorPrincipal: &admin.Principal{
		// 				Type: to.Ptr(admin.PrincipalTypeUser),
		// 				DisplayName: to.Ptr("Jacob Hancock"),
		// 				ID: to.Ptr("f3052d1c-61a9-46fb-8df9-0d78916ae041"),
		// 				UserDetails: &admin.PrincipalUserDetails{
		// 					UserPrincipalName: to.Ptr("jacob@contoso.com"),
		// 				},
		// 			},
		// 			ExpirationTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-13T00:00:00.000Z"); return t}()),
		// 			ID: to.Ptr("dccc162f-7a41-4720-83c3-5c7e81187959"),
		// 			InvitationURL: to.Ptr("https://app.fabric.microsoft.com/externaldatasharing/accept?si=VyT5NJ3%2bNkySqEmf368Pjw-dccc162f-7a41-4720-83c3-5c7e81187959"),
		// 			ItemID: to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		// 			Paths: []string{
		// 				"Files/Sales/Contoso_Sales_2023"},
		// 				Recipient: &admin.ExternalDataShareRecipient{
		// 					UserPrincipalName: to.Ptr("lisa@fabrikam.com"),
		// 				},
		// 				Status: to.Ptr(admin.ExternalDataShareStatusPending),
		// 				WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		// 			},
		// 			{
		// 				AcceptedByTenantID: to.Ptr("c51dc03f-268a-4da0-a879-25f24947ab8b"),
		// 				CreatorPrincipal: &admin.Principal{
		// 					Type: to.Ptr(admin.PrincipalTypeUser),
		// 					DisplayName: to.Ptr("Jacob Hancock"),
		// 					ID: to.Ptr("f3052d1c-61a9-46fb-8df9-0d78916ae041"),
		// 					UserDetails: &admin.PrincipalUserDetails{
		// 						UserPrincipalName: to.Ptr("jacob@contoso.com"),
		// 					},
		// 				},
		// 				ExpirationTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-31T00:00:00.000Z"); return t}()),
		// 				ID: to.Ptr("96c21561-65b8-4b23-bb9a-ee8cef945c45"),
		// 				InvitationURL: to.Ptr("https://app.fabric.microsoft.com/externaldatasharing/accept?si=VyT5NJ3%2bNkySqEmf368Pjw-96c21561-65b8-4b23-bb9a-ee8cef945c45"),
		// 				ItemID: to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		// 				Paths: []string{
		// 					"Files/Sales/Contoso_Sales_2023"},
		// 					Recipient: &admin.ExternalDataShareRecipient{
		// 						UserPrincipalName: to.Ptr("lisa@fabrikam.com"),
		// 					},
		// 					Status: to.Ptr(admin.ExternalDataShareStatusActive),
		// 					WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		// 				},
		// 				{
		// 					CreatorPrincipal: &admin.Principal{
		// 						Type: to.Ptr(admin.PrincipalTypeUser),
		// 						DisplayName: to.Ptr("Eric Solomon"),
		// 						ID: to.Ptr("81fac5e1-2a81-421b-a168-110b1c72fa11"),
		// 						UserDetails: &admin.PrincipalUserDetails{
		// 							UserPrincipalName: to.Ptr("eric@contoso.com"),
		// 						},
		// 					},
		// 					ExpirationTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-01T00:00:00.000Z"); return t}()),
		// 					ID: to.Ptr("0f40aeca-8f78-4a6f-a552-e5c45faadc60"),
		// 					InvitationURL: to.Ptr("https://app.fabric.microsoft.com/externaldatasharing/accept?si=VyT5NJ3%2bNkySqEmf368Pjw-0f40aeca-8f78-4a6f-a552-e5c45faadc60"),
		// 					ItemID: to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		// 					Paths: []string{
		// 						"Files/Sales/Contoso_Sales_2023"},
		// 						Recipient: &admin.ExternalDataShareRecipient{
		// 							UserPrincipalName: to.Ptr("lisa@fabrikam.com"),
		// 						},
		// 						Status: to.Ptr(admin.ExternalDataShareStatusInvitationExpired),
		// 						WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		// 					},
		// 					{
		// 						AcceptedByTenantID: to.Ptr("c51dc03f-268a-4da0-a879-25f24947ab8b"),
		// 						CreatorPrincipal: &admin.Principal{
		// 							Type: to.Ptr(admin.PrincipalTypeUser),
		// 							DisplayName: to.Ptr("Eric Solomon"),
		// 							ID: to.Ptr("81fac5e1-2a81-421b-a168-110b1c72fa11"),
		// 							UserDetails: &admin.PrincipalUserDetails{
		// 								UserPrincipalName: to.Ptr("eric@contoso.com"),
		// 							},
		// 						},
		// 						ExpirationTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-01T00:00:00.000Z"); return t}()),
		// 						ID: to.Ptr("89e82a82-0140-4837-8eee-9c919e3e5952"),
		// 						InvitationURL: to.Ptr("https://app.fabric.microsoft.com/externaldatasharing/accept?si=VyT5NJ3%2bNkySqEmf368Pjw-89e82a82-0140-4837-8eee-9c919e3e5952"),
		// 						ItemID: to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
		// 						Paths: []string{
		// 							"Files/Sales/Contoso_Sales_2023"},
		// 							Recipient: &admin.ExternalDataShareRecipient{
		// 								UserPrincipalName: to.Ptr("lisa@fabrikam.com"),
		// 							},
		// 							Status: to.Ptr(admin.ExternalDataShareStatusRevoked),
		// 							WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		// 					}},
		// 				}
	}
}

// Generated from example definition
func ExampleExternalDataSharesClient_RevokeExternalDataShare() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := admin.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewExternalDataSharesClient().RevokeExternalDataShare(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", "dccc162f-7a41-4720-83c3-5c7e81187959", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
