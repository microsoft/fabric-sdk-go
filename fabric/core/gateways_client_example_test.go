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
func ExampleGatewaysClient_NewListGatewaysPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGatewaysClient().NewListGatewaysPager(&core.GatewaysClientListGatewaysOptions{ContinuationToken: nil})
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
		// page.ListGatewaysResponse = core.ListGatewaysResponse{
		// 	ContinuationToken: to.Ptr("LDEsMTAwMDAwLDA%3D"),
		// 	ContinuationURI: to.Ptr("https://api.fabric.microsoft.com/v1/connections?continuationToken=LDEsMTAwMDAwLDA%3D"),
		// 	Value: []core.GatewayClassification{
		// 		&core.OnPremisesGateway{
		// 			Type: to.Ptr(core.GatewayTypeOnPremises),
		// 			ID: to.Ptr("8e41c4dd-a382-4937-9bf3-695ab881f7c2"),
		// 			AllowCloudConnectionRefresh: to.Ptr(true),
		// 			AllowCustomConnectors: to.Ptr(true),
		// 			DisplayName: to.Ptr("ContosoOnPremisesGateway"),
		// 			LoadBalancingSetting: to.Ptr(core.LoadBalancingSettingDistributeEvenly),
		// 			NumberOfMemberGateways: to.Ptr[int32](2),
		// 			PublicKey: &core.PublicKey{
		// 				Exponent: to.Ptr("AQGB"),
		// 				Modulus: to.Ptr("od9b...90Jp1Q=="),
		// 			},
		// 			Version: to.Ptr("3000.1.1"),
		// 		},
		// 		&core.OnPremisesGatewayPersonal{
		// 			Type: to.Ptr(core.GatewayTypeOnPremisesPersonal),
		// 			ID: to.Ptr("ca8979ff-4238-4489-ad13-2e1bd69a8412"),
		// 			PublicKey: &core.PublicKey{
		// 				Exponent: to.Ptr("AQOV"),
		// 				Modulus: to.Ptr("pt9b...87Jp1Q=="),
		// 			},
		// 			Version: to.Ptr("3000.1.1"),
		// 		},
		// 		&core.VirtualNetworkGateway{
		// 			Type: to.Ptr(core.GatewayTypeVirtualNetwork),
		// 			ID: to.Ptr("271c5c9a-0860-4927-b1da-ce49008d6565"),
		// 			CapacityID: to.Ptr("ed26b6f3-7bc5-44b0-9565-a8942619ef4c"),
		// 			DisplayName: to.Ptr("ContosoVirtualNetworkGateway"),
		// 			InactivityMinutesBeforeSleep: to.Ptr[int32](1440),
		// 			NumberOfMemberGateways: to.Ptr[int32](3),
		// 			VirtualNetworkAzureResource: &core.VirtualNetworkAzureResource{
		// 				ResourceGroupName: to.Ptr("ContosoResourceGroup"),
		// 				SubscriptionID: to.Ptr("879b4ba0-ed17-4ff2-851e-4a2228e00b70"),
		// 				SubnetName: to.Ptr("ContosoSubnet"),
		// 				VirtualNetworkName: to.Ptr("ContosoVirtualNetwork"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleGatewaysClient_CreateGateway() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewGatewaysClient().CreateGateway(ctx, &core.CreateVirtualNetworkGatewayRequest{
		Type:                         to.Ptr(core.GatewayTypeVirtualNetwork),
		CapacityID:                   to.Ptr("ed26b6f3-7bc5-44b0-9565-a8942619ef4c"),
		DisplayName:                  to.Ptr("ContosoVirtualNetworkGateway"),
		InactivityMinutesBeforeSleep: to.Ptr[int32](120),
		NumberOfMemberGateways:       to.Ptr[int32](3),
		VirtualNetworkAzureResource: &core.VirtualNetworkAzureResource{
			ResourceGroupName:  to.Ptr("ContosoResourceGroup"),
			SubscriptionID:     to.Ptr("879b4ba0-ed17-4ff2-851e-4a2228e00b70"),
			SubnetName:         to.Ptr("ContosoSubnet"),
			VirtualNetworkName: to.Ptr("ContosoVirtualNetwork"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleGatewaysClient_GetGateway() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGatewaysClient().GetGateway(ctx, "8e41c4dd-a382-4937-9bf3-695ab881f7c2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = core.GatewaysClientGetGatewayResponse{
	// 	                            GatewayClassification: &core.OnPremisesGateway{
	// 		Type: to.Ptr(core.GatewayTypeOnPremises),
	// 		ID: to.Ptr("8e41c4dd-a382-4937-9bf3-695ab881f7c2"),
	// 		AllowCloudConnectionRefresh: to.Ptr(true),
	// 		AllowCustomConnectors: to.Ptr(true),
	// 		DisplayName: to.Ptr("ContosoOnPremisesGateway"),
	// 		LoadBalancingSetting: to.Ptr(core.LoadBalancingSettingDistributeEvenly),
	// 		NumberOfMemberGateways: to.Ptr[int32](2),
	// 		PublicKey: &core.PublicKey{
	// 			Exponent: to.Ptr("AQGB"),
	// 			Modulus: to.Ptr("od9b...90Jp1Q=="),
	// 		},
	// 		Version: to.Ptr("3000.1.1"),
	// 	},
	// 	                        }
}

// Generated from example definition
func ExampleGatewaysClient_UpdateGateway_onPremisesGatewayExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGatewaysClient().UpdateGateway(ctx, "3d1290e1-e3ed-4bd6-93bc-2bbd5b49a789", &core.UpdateOnPremisesGatewayRequest{
		Type:                        to.Ptr(core.GatewayTypeOnPremises),
		DisplayName:                 to.Ptr("ContosoGatewayCluster1"),
		AllowCloudConnectionRefresh: to.Ptr(false),
		AllowCustomConnectors:       to.Ptr(false),
		LoadBalancingSetting:        to.Ptr(core.LoadBalancingSettingFailover),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = core.GatewaysClientUpdateGatewayResponse{
	// 	                            GatewayClassification: &core.OnPremisesGateway{
	// 		Type: to.Ptr(core.GatewayTypeOnPremises),
	// 		ID: to.Ptr("3d1290e1-e3ed-4bd6-93bc-2bbd5b49a789"),
	// 		AllowCloudConnectionRefresh: to.Ptr(false),
	// 		AllowCustomConnectors: to.Ptr(false),
	// 		DisplayName: to.Ptr("ContosoGatewayCluster1"),
	// 		LoadBalancingSetting: to.Ptr(core.LoadBalancingSettingFailover),
	// 		NumberOfMemberGateways: to.Ptr[int32](2),
	// 		PublicKey: &core.PublicKey{
	// 			Exponent: to.Ptr("AQGB"),
	// 			Modulus: to.Ptr("od9b...90Jp1Q=="),
	// 		},
	// 		Version: to.Ptr("3000.1.2"),
	// 	},
	// 	                        }
}

// Generated from example definition
func ExampleGatewaysClient_UpdateGateway_virtualNetworkGatewayExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGatewaysClient().UpdateGateway(ctx, "7015263e-885f-455b-80f7-bbf862899176", &core.UpdateVirtualNetworkGatewayRequest{
		Type:                         to.Ptr(core.GatewayTypeVirtualNetwork),
		DisplayName:                  to.Ptr("ContosoVirtualNetworkGateway1"),
		CapacityID:                   to.Ptr("7cf7181f-9457-4178-b488-e7472b02faf4"),
		InactivityMinutesBeforeSleep: to.Ptr[int32](720),
		NumberOfMemberGateways:       to.Ptr[int32](5),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = core.GatewaysClientUpdateGatewayResponse{
	// 	                            GatewayClassification: &core.VirtualNetworkGateway{
	// 		Type: to.Ptr(core.GatewayTypeVirtualNetwork),
	// 		ID: to.Ptr("7015263e-885f-455b-80f7-bbf862899176"),
	// 		CapacityID: to.Ptr("7cf7181f-9457-4178-b488-e7472b02faf4"),
	// 		DisplayName: to.Ptr("ContosoVirtualNetworkGateway1"),
	// 		InactivityMinutesBeforeSleep: to.Ptr[int32](720),
	// 		NumberOfMemberGateways: to.Ptr[int32](5),
	// 		VirtualNetworkAzureResource: &core.VirtualNetworkAzureResource{
	// 			ResourceGroupName: to.Ptr("ContosoResourceGroup"),
	// 			SubscriptionID: to.Ptr("879b4ba0-ed17-4ff2-851e-4a2228e00b70"),
	// 			SubnetName: to.Ptr("ContosoSubnet"),
	// 			VirtualNetworkName: to.Ptr("ContosoVirtualNetwork"),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition
func ExampleGatewaysClient_DeleteGateway() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewGatewaysClient().DeleteGateway(ctx, "411a04d3-7c15-4b69-9dd8-de6e80df1009", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleGatewaysClient_ListGatewayMembers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGatewaysClient().ListGatewayMembers(ctx, "8e41c4dd-a382-4937-9bf3-695ab881f7c2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ListGatewayMembersResponse = core.ListGatewayMembersResponse{
	// 	Value: []core.OnPremisesGatewayMember{
	// 		{
	// 			DisplayName: to.Ptr("ContosoPrimaryMemberGateway"),
	// 			Enabled: to.Ptr(true),
	// 			ID: to.Ptr("8e41c4dd-a382-4937-9bf3-695ab881f7c2"),
	// 			PublicKey: &core.PublicKey{
	// 				Exponent: to.Ptr("AQGB"),
	// 				Modulus: to.Ptr("od9b...90Jp1Q=="),
	// 			},
	// 			Version: to.Ptr("3000.1.1"),
	// 		},
	// 		{
	// 			DisplayName: to.Ptr("ContosoSecondaryMemberGateway"),
	// 			Enabled: to.Ptr(false),
	// 			ID: to.Ptr("5d225cda-42d5-43d3-bc40-218f746e1d58"),
	// 			PublicKey: &core.PublicKey{
	// 				Exponent: to.Ptr("AQCB"),
	// 				Modulus: to.Ptr("o57c...90Jh1P=="),
	// 			},
	// 			Version: to.Ptr("3000.1.1"),
	// 	}},
	// }
}

// Generated from example definition
func ExampleGatewaysClient_UpdateGatewayMember() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGatewaysClient().UpdateGatewayMember(ctx, "41d198df-5fb2-4463-8157-2af5153e6503", "1e624ffb-f74a-4dfb-9aa8-ea5056da37da", core.UpdateGatewayMemberRequest{
		DisplayName: to.Ptr("ContosoGatewayMember1"),
		Enabled:     to.Ptr(false),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OnPremisesGatewayMember = core.OnPremisesGatewayMember{
	// 	DisplayName: to.Ptr("ContosoGatewayMember1"),
	// 	Enabled: to.Ptr(false),
	// 	ID: to.Ptr("1e624ffb-f74a-4dfb-9aa8-ea5056da37da"),
	// 	PublicKey: &core.PublicKey{
	// 		Exponent: to.Ptr("AQGB"),
	// 		Modulus: to.Ptr("od9b...90Jp1Q=="),
	// 	},
	// 	Version: to.Ptr("3000.1.1"),
	// }
}

// Generated from example definition
func ExampleGatewaysClient_DeleteGatewayMember() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewGatewaysClient().DeleteGatewayMember(ctx, "ad457a0a-1fc4-4218-9867-1d84661ca4b8", "f921ee6b-8feb-4595-8aa7-3ed34338e8b6", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleGatewaysClient_NewListGatewayRoleAssignmentsPager_listGatewayRoleAssignmentExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGatewaysClient().NewListGatewayRoleAssignmentsPager("8e41c4dd-a382-4937-9bf3-695ab881f7c2", &core.GatewaysClientListGatewayRoleAssignmentsOptions{ContinuationToken: nil})
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
		// page.GatewayRoleAssignments = core.GatewayRoleAssignments{
		// 	Value: []core.GatewayRoleAssignment{
		// 		{
		// 			ID: to.Ptr("ef543eb8-969c-46b0-a5a1-3f93351b6b31"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeUser),
		// 				ID: to.Ptr("ef543eb8-969c-46b0-a5a1-3f93351b6b31"),
		// 			},
		// 			Role: to.Ptr(core.GatewayRoleAdmin),
		// 		},
		// 		{
		// 			ID: to.Ptr("5cf7d203-9123-4dff-a87f-7097dc4b5d60"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeUser),
		// 				ID: to.Ptr("5cf7d203-9123-4dff-a87f-7097dc4b5d60"),
		// 			},
		// 			Role: to.Ptr(core.GatewayRoleConnectionCreatorWithResharing),
		// 		},
		// 		{
		// 			ID: to.Ptr("5931cd21-857f-42a5-beaf-0120e8b36542"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeUser),
		// 				ID: to.Ptr("5931cd21-857f-42a5-beaf-0120e8b36542"),
		// 			},
		// 			Role: to.Ptr(core.GatewayRoleConnectionCreator),
		// 		},
		// 		{
		// 			ID: to.Ptr("97614f04-507c-4f6c-8dbc-da1845f582ef"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeUser),
		// 				ID: to.Ptr("97614f04-507c-4f6c-8dbc-da1845f582ef"),
		// 			},
		// 			Role: to.Ptr(core.GatewayRoleConnectionCreator),
		// 		},
		// 		{
		// 			ID: to.Ptr("a5d9f30d-a15c-4fb9-b8ff-e2a884c9fd82"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeGroup),
		// 				ID: to.Ptr("a5d9f30d-a15c-4fb9-b8ff-e2a884c9fd82"),
		// 			},
		// 			Role: to.Ptr(core.GatewayRoleConnectionCreator),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleGatewaysClient_NewListGatewayRoleAssignmentsPager_listGatewayRoleAssignmentWithContinuationExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGatewaysClient().NewListGatewayRoleAssignmentsPager("8e41c4dd-a382-4937-9bf3-695ab881f7c2", &core.GatewaysClientListGatewayRoleAssignmentsOptions{ContinuationToken: nil})
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
		// page.GatewayRoleAssignments = core.GatewayRoleAssignments{
		// 	ContinuationToken: to.Ptr("LDEsMTAwMDAwLDA%3D"),
		// 	ContinuationURI: to.Ptr("https://api.fabric.microsoft.com/v1/gateways/8e41c4dd-a382-4937-9bf3-695ab881f7c2/roleAssignments?continuationToken=LDEsMTAwMDAwLDA%3D"),
		// 	Value: []core.GatewayRoleAssignment{
		// 		{
		// 			ID: to.Ptr("ef543eb8-969c-46b0-a5a1-3f93351b6b31"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeUser),
		// 				ID: to.Ptr("ef543eb8-969c-46b0-a5a1-3f93351b6b31"),
		// 			},
		// 			Role: to.Ptr(core.GatewayRoleAdmin),
		// 		},
		// 		{
		// 			ID: to.Ptr("5cf7d203-9123-4dff-a87f-7097dc4b5d60"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeUser),
		// 				ID: to.Ptr("5cf7d203-9123-4dff-a87f-7097dc4b5d60"),
		// 			},
		// 			Role: to.Ptr(core.GatewayRoleConnectionCreatorWithResharing),
		// 		},
		// 		{
		// 			ID: to.Ptr("5931cd21-857f-42a5-beaf-0120e8b36542"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeUser),
		// 				ID: to.Ptr("5931cd21-857f-42a5-beaf-0120e8b36542"),
		// 			},
		// 			Role: to.Ptr(core.GatewayRoleConnectionCreator),
		// 		},
		// 		{
		// 			ID: to.Ptr("97614f04-507c-4f6c-8dbc-da1845f582ef"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeUser),
		// 				ID: to.Ptr("97614f04-507c-4f6c-8dbc-da1845f582ef"),
		// 			},
		// 			Role: to.Ptr(core.GatewayRoleConnectionCreator),
		// 		},
		// 		{
		// 			ID: to.Ptr("a5d9f30d-a15c-4fb9-b8ff-e2a884c9fd82"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeGroup),
		// 				ID: to.Ptr("a5d9f30d-a15c-4fb9-b8ff-e2a884c9fd82"),
		// 			},
		// 			Role: to.Ptr(core.GatewayRoleConnectionCreator),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleGatewaysClient_AddGatewayRoleAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewGatewaysClient().AddGatewayRoleAssignment(ctx, "d12d139f-4141-467c-9f53-80787b198843", core.AddGatewayRoleAssignmentRequest{
		Principal: &core.Principal{
			Type: to.Ptr(core.PrincipalTypeUser),
			ID:   to.Ptr("6a002b3d-e4ec-43df-8c08-e8eb7547d9dd"),
		},
		Role: to.Ptr(core.GatewayRoleConnectionCreator),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleGatewaysClient_GetGatewayRoleAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGatewaysClient().GetGatewayRoleAssignment(ctx, "8e41c4dd-a382-4937-9bf3-695ab881f7c2", "056afb37-8f6c-4fd8-9aa5-64ba7f1974e7", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GatewayRoleAssignment = core.GatewayRoleAssignment{
	// 	ID: to.Ptr("056afb37-8f6c-4fd8-9aa5-64ba7f1974e7"),
	// 	Principal: &core.Principal{
	// 		Type: to.Ptr(core.PrincipalTypeUser),
	// 		ID: to.Ptr("056afb37-8f6c-4fd8-9aa5-64ba7f1974e7"),
	// 	},
	// 	Role: to.Ptr(core.GatewayRoleAdmin),
	// }
}

// Generated from example definition
func ExampleGatewaysClient_UpdateGatewayRoleAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGatewaysClient().UpdateGatewayRoleAssignment(ctx, "8e41c4dd-a382-4937-9bf3-695ab881f7c2", "43970761-afc9-4428-ae6e-3b08bef098ff", core.UpdateGatewayRoleAssignmentRequest{
		Role: to.Ptr(core.GatewayRoleConnectionCreator),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GatewayRoleAssignment = core.GatewayRoleAssignment{
	// 	ID: to.Ptr("43970761-afc9-4428-ae6e-3b08bef098ff"),
	// 	Principal: &core.Principal{
	// 		Type: to.Ptr(core.PrincipalTypeUser),
	// 		ID: to.Ptr("43970761-afc9-4428-ae6e-3b08bef098ff"),
	// 	},
	// 	Role: to.Ptr(core.GatewayRoleConnectionCreator),
	// }
}

// Generated from example definition
func ExampleGatewaysClient_DeleteGatewayRoleAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewGatewaysClient().DeleteGatewayRoleAssignment(ctx, "8e41c4dd-a382-4937-9bf3-695ab881f7c2", "056afb37-8f6c-4fd8-9aa5-64ba7f1974e7", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
