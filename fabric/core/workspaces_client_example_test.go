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
func ExampleWorkspacesClient_CreateWorkspace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewWorkspacesClient().CreateWorkspace(ctx, core.CreateWorkspaceRequest{
		DisplayName: to.Ptr("New workspace"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleWorkspacesClient_NewListWorkspacesPager_listWorkspacesExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacesClient().NewListWorkspacesPager(&core.WorkspacesClientListWorkspacesOptions{Roles: nil,
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
		// page.Workspaces = core.Workspaces{
		// 	Value: []core.Workspace{
		// 		{
		// 			Type: to.Ptr(core.WorkspaceTypePersonal),
		// 			Description: to.Ptr(""),
		// 			DisplayName: to.Ptr("My workspace"),
		// 			ID: to.Ptr("fa9ad228-3e6b-44d4-b5f4-e275f337afa9"),
		// 		},
		// 		{
		// 			Type: to.Ptr(core.WorkspaceTypeWorkspace),
		// 			Description: to.Ptr("A  workspace used by the marketing team"),
		// 			DisplayName: to.Ptr("Marketing"),
		// 			ID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff227"),
		// 		},
		// 		{
		// 			Type: to.Ptr(core.WorkspaceTypeWorkspace),
		// 			Description: to.Ptr("A workspace used by the finance team"),
		// 			CapacityID: to.Ptr("171018af-1531-4e61-942a-74f024b7f9fd"),
		// 			DisplayName: to.Ptr("Finance"),
		// 			ID: to.Ptr("f2d70dc6-8f3e-4f2c-b00e-e2d336d7d711"),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleWorkspacesClient_NewListWorkspacesPager_listWorkspacesWithContinuationExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacesClient().NewListWorkspacesPager(&core.WorkspacesClientListWorkspacesOptions{Roles: nil,
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
		// page.Workspaces = core.Workspaces{
		// 	ContinuationToken: to.Ptr("LDEsMTAwMDAwLDA%3D"),
		// 	ContinuationURI: to.Ptr("https://api.fabric.microsoft.com/v1/workspaces?continuationToken=LDEsMTAwMDAwLDA%3D"),
		// 	Value: []core.Workspace{
		// 		{
		// 			Type: to.Ptr(core.WorkspaceTypePersonal),
		// 			Description: to.Ptr(""),
		// 			DisplayName: to.Ptr("My workspace"),
		// 			ID: to.Ptr("fa9ad228-3e6b-44d4-b5f4-e275f337afa9"),
		// 		},
		// 		{
		// 			Type: to.Ptr(core.WorkspaceTypeWorkspace),
		// 			Description: to.Ptr("A  workspace used by the marketing team"),
		// 			DisplayName: to.Ptr("Marketing"),
		// 			ID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff227"),
		// 		},
		// 		{
		// 			Type: to.Ptr(core.WorkspaceTypeWorkspace),
		// 			Description: to.Ptr("A workspace used by the finance team"),
		// 			CapacityID: to.Ptr("171018af-1531-4e61-942a-74f024b7f9fd"),
		// 			DisplayName: to.Ptr("Finance"),
		// 			ID: to.Ptr("f2d70dc6-8f3e-4f2c-b00e-e2d336d7d711"),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleWorkspacesClient_NewListWorkspacesPager_listWorkspacesWithRolesFilterExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacesClient().NewListWorkspacesPager(&core.WorkspacesClientListWorkspacesOptions{Roles: to.Ptr("Admin,Member,Contributor,Viewer"),
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
		// page.Workspaces = core.Workspaces{
		// 	Value: []core.Workspace{
		// 		{
		// 			Type: to.Ptr(core.WorkspaceTypePersonal),
		// 			Description: to.Ptr("A workspace for admins"),
		// 			DisplayName: to.Ptr("Admins workspace"),
		// 			ID: to.Ptr("fa9ad228-3e6b-44d4-b5f4-e275f337afa9"),
		// 		},
		// 		{
		// 			Type: to.Ptr(core.WorkspaceTypeWorkspace),
		// 			Description: to.Ptr("A workspace for members"),
		// 			DisplayName: to.Ptr("Members workspace"),
		// 			ID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff227"),
		// 		},
		// 		{
		// 			Type: to.Ptr(core.WorkspaceTypeWorkspace),
		// 			Description: to.Ptr("A workspace for contributors"),
		// 			DisplayName: to.Ptr("Contributors workspace"),
		// 			ID: to.Ptr("0c02a0cd-71bc-410f-aa05-5a7bc98765f7"),
		// 		},
		// 		{
		// 			Type: to.Ptr(core.WorkspaceTypeWorkspace),
		// 			Description: to.Ptr("A workspace for viewers"),
		// 			DisplayName: to.Ptr("Viewers workspace"),
		// 			ID: to.Ptr("99d58687-8903-4dbd-8a78-40f95dca7177"),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleWorkspacesClient_GetWorkspace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().GetWorkspace(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff227", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkspaceInfo = core.WorkspaceInfo{
	// 	Type: to.Ptr(core.WorkspaceTypeWorkspace),
	// 	Description: to.Ptr("New workspace description"),
	// 	CapacityID: to.Ptr("56bac802-080d-4f73-8a42-1b406eb1fcac"),
	// 	DisplayName: to.Ptr("New workspace"),
	// 	ID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff227"),
	// 	CapacityAssignmentProgress: to.Ptr(core.CapacityAssignmentProgressCompleted),
	// 	CapacityRegion: to.Ptr(core.CapacityRegionEastUS),
	// 	OneLakeEndpoints: &core.OneLakeEndpoints{
	// 		BlobEndpoint: to.Ptr("https://eastus-onelake.blob.fabric.microsoft.com"),
	// 		DfsEndpoint: to.Ptr("https://eastus-onelake.dfs.fabric.microsoft.com"),
	// 	},
	// 	WorkspaceIdentity: &core.WorkspaceIdentity{
	// 		ApplicationID: to.Ptr("00a4a8f9-78d3-41b3-b87a-6ae5271c8d0d"),
	// 		ServicePrincipalID: to.Ptr("5ba4ae58-d402-45c6-a848-0253e834fd78"),
	// 	},
	// }
}

// Generated from example definition
func ExampleWorkspacesClient_UpdateWorkspace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().UpdateWorkspace(ctx, "33bae707-5fe7-4352-89bd-061a1318b60a", core.UpdateWorkspaceRequest{
		Description: to.Ptr("Workspace New description"),
		DisplayName: to.Ptr("Workspace New displayName"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Workspace = core.Workspace{
	// 	Type: to.Ptr(core.WorkspaceTypeWorkspace),
	// 	Description: to.Ptr("Workspace New description"),
	// 	DisplayName: to.Ptr("Workspace New displayName"),
	// 	ID: to.Ptr("33bae707-5fe7-4352-89bd-061a1318b60a"),
	// }
}

// Generated from example definition
func ExampleWorkspacesClient_DeleteWorkspace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewWorkspacesClient().DeleteWorkspace(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff222", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleWorkspacesClient_AddWorkspaceRoleAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewWorkspacesClient().AddWorkspaceRoleAssignment(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff512", core.AddWorkspaceRoleAssignmentRequest{
		Principal: &core.Principal{
			Type: to.Ptr(core.PrincipalTypeUser),
			ID:   to.Ptr("8eedb1b0-3af8-4b17-8e7e-663e61e12211"),
		},
		Role: to.Ptr(core.WorkspaceRoleMember),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleWorkspacesClient_NewListWorkspaceRoleAssignmentsPager_getWorkspaceRoleAssignmentsExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacesClient().NewListWorkspaceRoleAssignmentsPager("e4ae4765-02a0-4cd8-bbef-65be17dd5a22", &core.WorkspacesClientListWorkspaceRoleAssignmentsOptions{ContinuationToken: nil})
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
		// page.WorkspaceRoleAssignments = core.WorkspaceRoleAssignments{
		// 	Value: []core.WorkspaceRoleAssignment{
		// 		{
		// 			ID: to.Ptr("81fac5e1-2a81-421b-a168-110b1c72fa11"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeUser),
		// 				DisplayName: to.Ptr("Eric Solomon"),
		// 				ID: to.Ptr("81fac5e1-2a81-421b-a168-110b1c72fa11"),
		// 				UserDetails: &core.PrincipalUserDetails{
		// 					UserPrincipalName: to.Ptr("eric@microsoft.com"),
		// 				},
		// 			},
		// 			Role: to.Ptr(core.WorkspaceRoleAdmin),
		// 		},
		// 		{
		// 			ID: to.Ptr("dbc4f130-681f-46b9-b19a-ca19ea5daa31"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeServicePrincipal),
		// 				DisplayName: to.Ptr("ServicePrincipal"),
		// 				ID: to.Ptr("dbc4f130-681f-46b9-b19a-ca19ea5daa31"),
		// 				ServicePrincipalDetails: &core.PrincipalServicePrincipalDetails{
		// 					AADAppID: to.Ptr("7ac9c70b-69f1-48c5-bf5b-69ac50578a55"),
		// 				},
		// 			},
		// 			Role: to.Ptr(core.WorkspaceRoleMember),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleWorkspacesClient_NewListWorkspaceRoleAssignmentsPager_getWorkspaceRoleAssignmentsWithContinuationExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacesClient().NewListWorkspaceRoleAssignmentsPager("e4ae4765-02a0-4cd8-bbef-65be17dd5a22", &core.WorkspacesClientListWorkspaceRoleAssignmentsOptions{ContinuationToken: nil})
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
		// page.WorkspaceRoleAssignments = core.WorkspaceRoleAssignments{
		// 	ContinuationToken: to.Ptr("LDEsMTAwMDAwLDA%3D"),
		// 	ContinuationURI: to.Ptr("https://api.fabric.microsoft.com/v1/workspaces/e4ae4765-02a0-4cd8-bbef-65be17dd5a22/roleAssignments?continuationToken=LDEsMTAwMDAwLDA%3D"),
		// 	Value: []core.WorkspaceRoleAssignment{
		// 		{
		// 			ID: to.Ptr("81fac5e1-2a81-421b-a168-110b1c72fa11"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeUser),
		// 				DisplayName: to.Ptr("Eric Solomon"),
		// 				ID: to.Ptr("81fac5e1-2a81-421b-a168-110b1c72fa11"),
		// 				UserDetails: &core.PrincipalUserDetails{
		// 					UserPrincipalName: to.Ptr("eric@microsoft.com"),
		// 				},
		// 			},
		// 			Role: to.Ptr(core.WorkspaceRoleAdmin),
		// 		},
		// 		{
		// 			ID: to.Ptr("dbc4f130-681f-46b9-b19a-ca19ea5daa31"),
		// 			Principal: &core.Principal{
		// 				Type: to.Ptr(core.PrincipalTypeServicePrincipal),
		// 				DisplayName: to.Ptr("ServicePrincipal"),
		// 				ID: to.Ptr("dbc4f130-681f-46b9-b19a-ca19ea5daa31"),
		// 				ServicePrincipalDetails: &core.PrincipalServicePrincipalDetails{
		// 					AADAppID: to.Ptr("7ac9c70b-69f1-48c5-bf5b-69ac50578a55"),
		// 				},
		// 			},
		// 			Role: to.Ptr(core.WorkspaceRoleMember),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleWorkspacesClient_DeleteWorkspaceRoleAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewWorkspacesClient().DeleteWorkspaceRoleAssignment(ctx, "ba33a98a-2b66-49f1-bb71-80c38e4b3756", "6e335e92-a2a2-4b5a-970a-bd6a89fbb765", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleWorkspacesClient_UpdateWorkspaceRoleAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().UpdateWorkspaceRoleAssignment(ctx, "0ac682f5-aee3-4968-9d21-692eb3fd4056", "0218b8c4-f5a2-4a1e-bbbd-a986dd8aeb81", core.UpdateWorkspaceRoleAssignmentRequest{
		Role: to.Ptr(core.WorkspaceRoleContributor),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkspaceRoleAssignment = core.WorkspaceRoleAssignment{
	// 	ID: to.Ptr("0218b8c4-f5a2-4a1e-bbbd-a986dd8aeb81"),
	// 	Principal: &core.Principal{
	// 		Type: to.Ptr(core.PrincipalTypeUser),
	// 		DisplayName: to.Ptr("user1"),
	// 		ID: to.Ptr("0218b8c4-f5a2-4a1e-bbbd-a986dd8aeb81"),
	// 		UserDetails: &core.PrincipalUserDetails{
	// 			UserPrincipalName: to.Ptr("user1@microsoft.com"),
	// 		},
	// 	},
	// 	Role: to.Ptr(core.WorkspaceRoleContributor),
	// }
}

// Generated from example definition
func ExampleWorkspacesClient_GetWorkspaceRoleAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().GetWorkspaceRoleAssignment(ctx, "0ac682f5-aee3-4968-9d21-692eb3fd4056", "259b6674-74cf-4197-ac05-1bf391800ec2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkspaceRoleAssignment = core.WorkspaceRoleAssignment{
	// 	ID: to.Ptr("259b6674-74cf-4197-ac05-1bf391800ec2"),
	// 	Principal: &core.Principal{
	// 		Type: to.Ptr(core.PrincipalTypeUser),
	// 		ID: to.Ptr("259b6674-74cf-4197-ac05-1bf391800ec2"),
	// 	},
	// 	Role: to.Ptr(core.WorkspaceRoleMember),
	// }
}

// Generated from example definition
func ExampleWorkspacesClient_AssignToCapacity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewWorkspacesClient().AssignToCapacity(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff512", core.AssignWorkspaceToCapacityRequest{
		CapacityID: to.Ptr("0f084df7-c13d-451b-af5f-ed0c466403b2"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleWorkspacesClient_UnassignFromCapacity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewWorkspacesClient().UnassignFromCapacity(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff512", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleWorkspacesClient_BeginProvisionIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspacesClient().BeginProvisionIdentity(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff227", nil)
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
	// res.WorkspaceIdentity = core.WorkspaceIdentity{
	// 	ApplicationID: to.Ptr("00a4a8f9-78d3-41b3-b87a-6ae5271c8d0d"),
	// 	ServicePrincipalID: to.Ptr("5ba4ae58-d402-45c6-a848-0253e834fd78"),
	// }
}
