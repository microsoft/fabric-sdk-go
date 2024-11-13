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
func ExampleConnectionsClient_NewListConnectionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectionsClient().NewListConnectionsPager(&core.ConnectionsClientListConnectionsOptions{ContinuationToken: nil})
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
		// page.ListConnectionsResponse = core.ListConnectionsResponse{
		// 	ContinuationToken: to.Ptr("LDEsMTAwMDAwLDA%3D"),
		// 	ContinuationURI: to.Ptr("https://api.fabric.microsoft.com/v1/connections?continuationToken=LDEsMTAwMDAwLDA%3D"),
		// 	Value: []core.Connection{
		// 		{
		// 			ConnectionDetails: &core.ListConnectionDetails{
		// 				Type: to.Ptr("Web"),
		// 				Path: to.Ptr("https://www.contoso.com"),
		// 			},
		// 			ConnectivityType: to.Ptr(core.ConnectivityTypeShareableCloud),
		// 			CredentialDetails: &core.ListCredentialDetails{
		// 				ConnectionEncryption: to.Ptr(core.ConnectionEncryptionNotEncrypted),
		// 				SingleSignOnType: to.Ptr(core.SingleSignOnTypeNone),
		// 				SkipTestConnection: to.Ptr(false),
		// 				CredentialType: to.Ptr(core.CredentialTypeAnonymous),
		// 			},
		// 			DisplayName: to.Ptr("ContosoConnection1"),
		// 			ID: to.Ptr("6952a7b2-aea3-414f-9d85-6c0fe5d34539"),
		// 			PrivacyLevel: to.Ptr(core.PrivacyLevelPublic),
		// 		},
		// 		{
		// 			ConnectionDetails: &core.ListConnectionDetails{
		// 				Type: to.Ptr("SQL"),
		// 				Path: to.Ptr("contoso.database.windows.net;sales"),
		// 			},
		// 			ConnectivityType: to.Ptr(core.ConnectivityTypeOnPremisesGateway),
		// 			CredentialDetails: &core.ListCredentialDetails{
		// 				ConnectionEncryption: to.Ptr(core.ConnectionEncryptionAny),
		// 				SingleSignOnType: to.Ptr(core.SingleSignOnTypeNone),
		// 				SkipTestConnection: to.Ptr(false),
		// 				CredentialType: to.Ptr(core.CredentialTypeBasic),
		// 			},
		// 			DisplayName: to.Ptr("ContosoConnection2"),
		// 			GatewayID: to.Ptr("58376c10-5f61-4024-887e-748df4beae45"),
		// 			ID: to.Ptr("f6a39b76-9816-4e4b-b93a-f42e405017b7"),
		// 			PrivacyLevel: to.Ptr(core.PrivacyLevelOrganizational),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleConnectionsClient_CreateConnection_cloudExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewConnectionsClient().CreateConnection(ctx, &core.CreateCloudConnectionRequest{
		ConnectionDetails: &core.CreateConnectionDetails{
			Type:           to.Ptr("SQL"),
			CreationMethod: to.Ptr("SQL"),
			Parameters: []core.ConnectionDetailsParameterClassification{
				&core.ConnectionDetailsTextParameter{
					Name:     to.Ptr("server"),
					DataType: to.Ptr(core.DataTypeText),
					Value:    to.Ptr("contoso.database.windows.net"),
				},
				&core.ConnectionDetailsTextParameter{
					Name:     to.Ptr("database"),
					DataType: to.Ptr(core.DataTypeText),
					Value:    to.Ptr("sales"),
				}},
		},
		ConnectivityType: to.Ptr(core.ConnectivityTypeShareableCloud),
		DisplayName:      to.Ptr("ContosoCloudConnection"),
		PrivacyLevel:     to.Ptr(core.PrivacyLevelOrganizational),
		CredentialDetails: &core.CreateCredentialDetails{
			ConnectionEncryption: to.Ptr(core.ConnectionEncryptionNotEncrypted),
			SingleSignOnType:     to.Ptr(core.SingleSignOnTypeNone),
			SkipTestConnection:   to.Ptr(false),
			Credentials: &core.BasicCredentials{
				CredentialType: to.Ptr(core.CredentialTypeBasic),
				Password:       to.Ptr("********"),
				Username:       to.Ptr("admin"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleConnectionsClient_CreateConnection_virtualNetworkGatewayExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewConnectionsClient().CreateConnection(ctx, &core.CreateVirtualNetworkGatewayConnectionRequest{
		ConnectionDetails: &core.CreateConnectionDetails{
			Type:           to.Ptr("SQL"),
			CreationMethod: to.Ptr("SQL"),
			Parameters: []core.ConnectionDetailsParameterClassification{
				&core.ConnectionDetailsTextParameter{
					Name:     to.Ptr("server"),
					DataType: to.Ptr(core.DataTypeText),
					Value:    to.Ptr("contoso.database.windows.net"),
				},
				&core.ConnectionDetailsTextParameter{
					Name:     to.Ptr("database"),
					DataType: to.Ptr(core.DataTypeText),
					Value:    to.Ptr("sales"),
				}},
		},
		ConnectivityType: to.Ptr(core.ConnectivityTypeVirtualNetworkGateway),
		DisplayName:      to.Ptr("ContosoVirtualNetworkGatewayConnection"),
		PrivacyLevel:     to.Ptr(core.PrivacyLevelOrganizational),
		CredentialDetails: &core.CreateCredentialDetails{
			ConnectionEncryption: to.Ptr(core.ConnectionEncryptionEncrypted),
			SingleSignOnType:     to.Ptr(core.SingleSignOnTypeNone),
			SkipTestConnection:   to.Ptr(false),
			Credentials: &core.BasicCredentials{
				CredentialType: to.Ptr(core.CredentialTypeBasic),
				Password:       to.Ptr("*********"),
				Username:       to.Ptr("admin"),
			},
		},
		GatewayID: to.Ptr("93491300-cfbd-402f-bf17-9ace59a92354"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleConnectionsClient_GetConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectionsClient().GetConnection(ctx, "f6a39b76-9816-4e4b-b93a-f42e405017b7", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Connection = core.Connection{
	// 	ConnectionDetails: &core.ListConnectionDetails{
	// 		Type: to.Ptr("SQL"),
	// 		Path: to.Ptr("contoso.database.windows.net;sales"),
	// 	},
	// 	ConnectivityType: to.Ptr(core.ConnectivityTypeOnPremisesGateway),
	// 	CredentialDetails: &core.ListCredentialDetails{
	// 		ConnectionEncryption: to.Ptr(core.ConnectionEncryptionNotEncrypted),
	// 		SingleSignOnType: to.Ptr(core.SingleSignOnTypeNone),
	// 		SkipTestConnection: to.Ptr(false),
	// 		CredentialType: to.Ptr(core.CredentialTypeBasic),
	// 	},
	// 	DisplayName: to.Ptr("ContosoConnection"),
	// 	GatewayID: to.Ptr("58376c10-5f61-4024-887e-748df4beae45"),
	// 	ID: to.Ptr("f6a39b76-9816-4e4b-b93a-f42e405017b7"),
	// 	PrivacyLevel: to.Ptr(core.PrivacyLevelOrganizational),
	// }
}

// Generated from example definition
func ExampleConnectionsClient_UpdateConnection_personalCloudExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectionsClient().UpdateConnection(ctx, "7a0369b2-58c4-4b67-b3f3-92156a95f1cd", &core.UpdatePersonalCloudConnectionRequest{
		ConnectivityType: to.Ptr(core.ConnectivityTypePersonalCloud),
		PrivacyLevel:     to.Ptr(core.PrivacyLevelOrganizational),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Connection = core.Connection{
	// 	ConnectionDetails: &core.ListConnectionDetails{
	// 		Type: to.Ptr("SQL"),
	// 		Path: to.Ptr("contoso.database.windows.net;finances"),
	// 	},
	// 	ConnectivityType: to.Ptr(core.ConnectivityTypePersonalCloud),
	// 	CredentialDetails: &core.ListCredentialDetails{
	// 		ConnectionEncryption: to.Ptr(core.ConnectionEncryptionNotEncrypted),
	// 		SingleSignOnType: to.Ptr(core.SingleSignOnTypeNone),
	// 		SkipTestConnection: to.Ptr(false),
	// 		CredentialType: to.Ptr(core.CredentialTypeOAuth2),
	// 	},
	// 	ID: to.Ptr("7a0369b2-58c4-4b67-b3f3-92156a95f1cd"),
	// 	PrivacyLevel: to.Ptr(core.PrivacyLevelOrganizational),
	// }
}

// Generated from example definition
func ExampleConnectionsClient_UpdateConnection_shareableCloudExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectionsClient().UpdateConnection(ctx, "fa968eee-8075-48f6-8c6d-41260ee1396d", &core.UpdateShareableCloudConnectionRequest{
		ConnectivityType: to.Ptr(core.ConnectivityTypeShareableCloud),
		DisplayName:      to.Ptr("ContosoCloudConnection"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Connection = core.Connection{
	// 	ConnectionDetails: &core.ListConnectionDetails{
	// 		Type: to.Ptr("SQL"),
	// 		Path: to.Ptr("contoso.database.windows.net;networks"),
	// 	},
	// 	ConnectivityType: to.Ptr(core.ConnectivityTypeShareableCloud),
	// 	CredentialDetails: &core.ListCredentialDetails{
	// 		ConnectionEncryption: to.Ptr(core.ConnectionEncryptionNotEncrypted),
	// 		SingleSignOnType: to.Ptr(core.SingleSignOnTypeNone),
	// 		SkipTestConnection: to.Ptr(true),
	// 		CredentialType: to.Ptr(core.CredentialTypeBasic),
	// 	},
	// 	DisplayName: to.Ptr("ContosoCloudConnection"),
	// 	ID: to.Ptr("fa968eee-8075-48f6-8c6d-41260ee1396d"),
	// 	PrivacyLevel: to.Ptr(core.PrivacyLevelPublic),
	// }
}

// Generated from example definition
func ExampleConnectionsClient_UpdateConnection_virtualNetworkGatewayExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectionsClient().UpdateConnection(ctx, "6b571614-2e98-4bfd-b9ed-1cb8d3ffc396", &core.UpdateVirtualNetworkGatewayConnectionRequest{
		ConnectivityType: to.Ptr(core.ConnectivityTypeVirtualNetworkGateway),
		PrivacyLevel:     to.Ptr(core.PrivacyLevelOrganizational),
		CredentialDetails: &core.UpdateCredentialDetails{
			SingleSignOnType: to.Ptr(core.SingleSignOnTypeNone),
		},
		DisplayName: to.Ptr("ContosoMarketingVirtualNetworkGatewayConnection"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Connection = core.Connection{
	// 	ConnectionDetails: &core.ListConnectionDetails{
	// 		Type: to.Ptr("SQL"),
	// 		Path: to.Ptr("contoso.database.windows.net;marketing"),
	// 	},
	// 	ConnectivityType: to.Ptr(core.ConnectivityTypeVirtualNetworkGateway),
	// 	CredentialDetails: &core.ListCredentialDetails{
	// 		ConnectionEncryption: to.Ptr(core.ConnectionEncryptionNotEncrypted),
	// 		SingleSignOnType: to.Ptr(core.SingleSignOnTypeNone),
	// 		SkipTestConnection: to.Ptr(false),
	// 		CredentialType: to.Ptr(core.CredentialTypeBasic),
	// 	},
	// 	DisplayName: to.Ptr("ContosoMarketingVirtualNetworkGatewayConnection"),
	// 	GatewayID: to.Ptr("befccff4-3ee6-40d7-b8f1-a0a9fd684a85"),
	// 	ID: to.Ptr("6b571614-2e98-4bfd-b9ed-1cb8d3ffc396"),
	// 	PrivacyLevel: to.Ptr(core.PrivacyLevelOrganizational),
	// }
}

// Generated from example definition
func ExampleConnectionsClient_DeleteConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewConnectionsClient().DeleteConnection(ctx, "536f7c95-076c-40b5-8fe0-2179536e4161", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleConnectionsClient_NewListSupportedConnectionTypesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectionsClient().NewListSupportedConnectionTypesPager(&core.ConnectionsClientListSupportedConnectionTypesOptions{GatewayID: to.Ptr("6d824cb9-6bfb-4bdb-a702-238e172a8743"),
		ShowAllCreationMethods: nil,
		ContinuationToken:      nil,
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
		// page.ListSupportedConnectionTypesResponse = core.ListSupportedConnectionTypesResponse{
		// 	ContinuationToken: to.Ptr("LDEsMTAwMDAwLDA%3D"),
		// 	ContinuationURI: to.Ptr("https://api.fabric.microsoft.com/v1/connections/supportedConnections?connectivityType=ShareableCloud&continuationToken=LDEsMTAwMDAwLDA%3D"),
		// 	Value: []core.ConnectionCreationMetadata{
		// 		{
		// 			Type: to.Ptr("SQL"),
		// 			CreationMethods: []core.ConnectionCreationMethod{
		// 				{
		// 					Name: to.Ptr("SQL"),
		// 					Parameters: []core.ConnectionCreationParameter{
		// 						{
		// 							Name: to.Ptr("server"),
		// 							DataType: to.Ptr(core.DataTypeText),
		// 							Required: to.Ptr(true),
		// 						},
		// 						{
		// 							Name: to.Ptr("database"),
		// 							DataType: to.Ptr(core.DataTypeText),
		// 							Required: to.Ptr(false),
		// 					}},
		// 			}},
		// 			SupportedConnectionEncryptionTypes: []core.ConnectionEncryption{
		// 				core.ConnectionEncryptionEncrypted,
		// 				core.ConnectionEncryptionNotEncrypted},
		// 				SupportedCredentialTypes: []core.CredentialType{
		// 					core.CredentialTypeBasic,
		// 					core.CredentialTypeOAuth2},
		// 					SupportsSkipTestConnection: to.Ptr(true),
		// 			}},
		// 		}
	}
}

// Generated from example definition
func ExampleConnectionsClient_NewListConnectionRoleAssignmentsPager_listConnectionRoleAssignmentExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectionsClient().NewListConnectionRoleAssignmentsPager("9558d649-84f1-4b7e-850a-59b5d0ae95eb", &core.ConnectionsClientListConnectionRoleAssignmentsOptions{ContinuationToken: nil})
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
		// page.ConnectionRoleAssignments = core.ConnectionRoleAssignments{
		// 	Value: []core.ConnectionRoleAssignment{
		// 		{
		// 			ID: to.Ptr("1f227c77-826d-40eb-a2b7-27a325afb900"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeUser),
		// 				ID: to.Ptr("1c15c348-dd88-4065-8f25-57581c216bcf"),
		// 			},
		// 			Role: to.Ptr(core.ConnectionRoleOwner),
		// 		},
		// 		{
		// 			ID: to.Ptr("c8a395c6-f7ad-4caa-8ab1-0cc5684a3966"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeGroup),
		// 				ID: to.Ptr("d3a7dbf7-6641-48f2-851e-d71bbf9d90c4"),
		// 			},
		// 			Role: to.Ptr(core.ConnectionRoleOwner),
		// 		},
		// 		{
		// 			ID: to.Ptr("f42d1536-e8d3-4f69-8eab-8509bef50315"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeGroup),
		// 				ID: to.Ptr("48ba22da-4431-4da4-8b70-3401685bf9e5"),
		// 			},
		// 			Role: to.Ptr(core.ConnectionRoleUser),
		// 		},
		// 		{
		// 			ID: to.Ptr("40ac84af-e80b-4e6d-8b38-91541be3845f"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeUser),
		// 				ID: to.Ptr("1dfa1747-ce76-4caf-99c8-360b95f9f17a"),
		// 			},
		// 			Role: to.Ptr(core.ConnectionRoleUserWithReshare),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleConnectionsClient_NewListConnectionRoleAssignmentsPager_listConnectionRoleAssignmentWithContinuationExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectionsClient().NewListConnectionRoleAssignmentsPager("9558d649-84f1-4b7e-850a-59b5d0ae95eb", &core.ConnectionsClientListConnectionRoleAssignmentsOptions{ContinuationToken: nil})
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
		// page.ConnectionRoleAssignments = core.ConnectionRoleAssignments{
		// 	ContinuationToken: to.Ptr("LDEsMTAwMDAwLDA%3D"),
		// 	ContinuationURI: to.Ptr("https://api.fabric.microsoft.com/v1/connections/9558d649-84f1-4b7e-850a-59b5d0ae95eb/roleAssignments?continuationToken=LDEsMTAwMDAwLDA%3D"),
		// 	Value: []core.ConnectionRoleAssignment{
		// 		{
		// 			ID: to.Ptr("1f227c77-826d-40eb-a2b7-27a325afb900"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeUser),
		// 				ID: to.Ptr("1c15c348-dd88-4065-8f25-57581c216bcf"),
		// 			},
		// 			Role: to.Ptr(core.ConnectionRoleOwner),
		// 		},
		// 		{
		// 			ID: to.Ptr("c8a395c6-f7ad-4caa-8ab1-0cc5684a3966"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeGroup),
		// 				ID: to.Ptr("d3a7dbf7-6641-48f2-851e-d71bbf9d90c4"),
		// 			},
		// 			Role: to.Ptr(core.ConnectionRoleOwner),
		// 		},
		// 		{
		// 			ID: to.Ptr("f42d1536-e8d3-4f69-8eab-8509bef50315"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeGroup),
		// 				ID: to.Ptr("48ba22da-4431-4da4-8b70-3401685bf9e5"),
		// 			},
		// 			Role: to.Ptr(core.ConnectionRoleUser),
		// 		},
		// 		{
		// 			ID: to.Ptr("40ac84af-e80b-4e6d-8b38-91541be3845f"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeUser),
		// 				ID: to.Ptr("1dfa1747-ce76-4caf-99c8-360b95f9f17a"),
		// 			},
		// 			Role: to.Ptr(core.ConnectionRoleUserWithReshare),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleConnectionsClient_AddConnectionRoleAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewConnectionsClient().AddConnectionRoleAssignment(ctx, "f3a2e6af-d048-4f85-94d9-b3d16140df05", core.AddConnectionRoleAssignmentRequest{
		Principal: &core.Principal{
			Type: to.Ptr(core.PrincipalTypeUser),
			ID:   to.Ptr("6a002b3d-e4ec-43df-8c08-e8eb7547d9dd"),
		},
		Role: to.Ptr(core.ConnectionRoleOwner),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleConnectionsClient_GetConnectionRoleAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectionsClient().GetConnectionRoleAssignment(ctx, "4bbb41c0-1dcc-4a8c-aff5-c681a3d10208", "43970761-afc9-4428-ae6e-3b08bef098ff", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectionRoleAssignment = core.ConnectionRoleAssignment{
	// 	ID: to.Ptr("c28d68d6-2984-4d36-9a6b-82751093d3f1"),
	// 	Principal: &core.Principal{
	// 		Type: to.Ptr(core.PrincipalTypeUser),
	// 		ID: to.Ptr("43970761-afc9-4428-ae6e-3b08bef098ff"),
	// 	},
	// 	Role: to.Ptr(core.ConnectionRoleOwner),
	// }
}

// Generated from example definition
func ExampleConnectionsClient_UpdateConnectionRoleAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectionsClient().UpdateConnectionRoleAssignment(ctx, "fe8e181d-dbb8-471a-99f0-fdbf0a2ad4fd", "449a8a88-7f31-40c1-aece-8e128adb14a8", core.UpdateConnectionRoleAssignmentRequest{
		Role: to.Ptr(core.ConnectionRoleUserWithReshare),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectionRoleAssignment = core.ConnectionRoleAssignment{
	// 	ID: to.Ptr("43970761-afc9-4428-ae6e-3b08bef098ff"),
	// 	Principal: &core.Principal{
	// 		Type: to.Ptr(core.PrincipalTypeUser),
	// 		ID: to.Ptr("43970761-afc9-4428-ae6e-3b08bef098ff"),
	// 	},
	// 	Role: to.Ptr(core.ConnectionRoleUserWithReshare),
	// }
}

// Generated from example definition
func ExampleConnectionsClient_DeleteConnectionRoleAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewConnectionsClient().DeleteConnectionRoleAssignment(ctx, "a06fbff3-09e1-4958-a92b-2a7550459762", "b7439ed7-7331-4c59-b2bb-f4f917e61979", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
