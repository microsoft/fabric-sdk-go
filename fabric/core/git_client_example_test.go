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
func ExampleGitClient_GetConnection_getGitConnectionDetailsExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGitClient().GetConnection(ctx, "1455b6a2-c120-4c1c-dda7-92bafe99bec3", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GitConnection = core.GitConnection{
	// 	GitConnectionState: to.Ptr(core.GitConnectionStateConnectedAndInitialized),
	// 	GitProviderDetails: &core.AzureDevOpsDetails{
	// 		BranchName: to.Ptr("Test Branch"),
	// 		DirectoryName: to.Ptr(""),
	// 		GitProviderType: to.Ptr(core.GitProviderTypeAzureDevOps),
	// 		RepositoryName: to.Ptr("Test Repo"),
	// 		OrganizationName: to.Ptr("Test Organization"),
	// 		ProjectName: to.Ptr("Test Project"),
	// 	},
	// 	GitSyncDetails: &core.GitSyncDetails{
	// 		Head: to.Ptr("eaa737b48cda41b37ffefac772ea48f6fed3eac4"),
	// 		LastSyncTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-20T09:26:43.153Z"); return t}()),
	// 	},
	// }
}

// Generated from example definition
func ExampleGitClient_GetConnection_workspaceNotConnectedToGitExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGitClient().GetConnection(ctx, "1455b6a2-c120-4c1c-dda7-92bafe99bec3", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GitConnection = core.GitConnection{
	// 	GitConnectionState: to.Ptr(core.GitConnectionStateNotConnected),
	// }
}

// Generated from example definition
func ExampleGitClient_Connect_connectAWorkspaceToAzureDevOpsExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewGitClient().Connect(ctx, "1565e6a3-c020-4c0c-dda7-92bafe99eec5", core.GitConnectRequest{
		GitProviderDetails: &core.AzureDevOpsDetails{
			BranchName:       to.Ptr("Test Branch"),
			DirectoryName:    to.Ptr("Test Directory"),
			GitProviderType:  to.Ptr(core.GitProviderTypeAzureDevOps),
			RepositoryName:   to.Ptr("Test Repo"),
			OrganizationName: to.Ptr("Test Organization"),
			ProjectName:      to.Ptr("Test Project"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleGitClient_Connect_connectAWorkspaceToAzureDevOpsUsingConfiguredConnectionExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewGitClient().Connect(ctx, "1565e6a3-c020-4c0c-dda7-92bafe99eec5", core.GitConnectRequest{
		GitProviderDetails: &core.AzureDevOpsDetails{
			BranchName:       to.Ptr("Test Branch"),
			DirectoryName:    to.Ptr("Test Directory/Test Subdirectory"),
			GitProviderType:  to.Ptr(core.GitProviderTypeAzureDevOps),
			RepositoryName:   to.Ptr("Test Repo"),
			OrganizationName: to.Ptr("Test Organization"),
			ProjectName:      to.Ptr("Test Project"),
		},
		MyGitCredentials: &core.ConfiguredConnectionGitCredentials{
			Source:       to.Ptr(core.GitCredentialsSourceConfiguredConnection),
			ConnectionID: to.Ptr("3f2504e0-4f89-11d3-9a0c-0305e82c3301"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleGitClient_Connect_connectAWorkspaceToGitHubExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewGitClient().Connect(ctx, "1565e6a3-c020-4c0c-dda7-92bafe99eec5", core.GitConnectRequest{
		GitProviderDetails: &core.GitHubDetails{
			BranchName:      to.Ptr("Test Branch"),
			DirectoryName:   to.Ptr("Test Directory/Test Subdirectory"),
			GitProviderType: to.Ptr(core.GitProviderTypeGitHub),
			RepositoryName:  to.Ptr("Test Repo"),
			OwnerName:       to.Ptr("Test Owner"),
		},
		MyGitCredentials: &core.ConfiguredConnectionGitCredentials{
			Source:       to.Ptr(core.GitCredentialsSourceConfiguredConnection),
			ConnectionID: to.Ptr("3f2504e0-4f89-11d3-9a0c-0305e82c3301"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleGitClient_Disconnect() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewGitClient().Disconnect(ctx, "1565e6a3-c020-4c0c-dda7-92bafe99eec5", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleGitClient_BeginInitializeConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGitClient().BeginInitializeConnection(ctx, "1565e6a3-c020-4c0c-dda7-92bafe99eec5", &core.GitClientBeginInitializeConnectionOptions{GitInitializeConnectionRequest: nil})
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
	// res.InitializeGitConnectionResponse = core.InitializeGitConnectionResponse{
	// 	RemoteCommitHash: to.Ptr("7d03b2918bf6aa62f96d0a4307293f3853201705"),
	// 	RequiredAction: to.Ptr(core.RequiredActionUpdateFromGit),
	// 	WorkspaceHead: to.Ptr("eaa737b48cda41b37ffefac772ea48f6fed3eac4"),
	// }
}

// Generated from example definition
func ExampleGitClient_BeginUpdateFromGit() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGitClient().BeginUpdateFromGit(ctx, "1565e6a3-c020-4c0c-dda7-92bafe99eec5", core.UpdateFromGitRequest{
		ConflictResolution: &core.WorkspaceConflictResolution{
			ConflictResolutionPolicy: to.Ptr(core.ConflictResolutionPolicyPreferWorkspace),
			ConflictResolutionType:   to.Ptr(core.ConflictResolutionTypeWorkspace),
		},
		Options: &core.UpdateOptions{
			AllowOverrideItems: to.Ptr(true),
		},
		RemoteCommitHash: to.Ptr("7d03b2918bf6aa62f96d0a4307293f3853201705"),
		WorkspaceHead:    to.Ptr("eaa737b48cda41b37ffefac772ea48f6fed3eac4"),
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
func ExampleGitClient_BeginGetStatus_getStatusExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGitClient().BeginGetStatus(ctx, "1455b6a2-c120-4c1c-dda7-92bafe99bec3", nil)
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
	// res.GitStatusResponse = core.GitStatusResponse{
	// 	Changes: []core.ItemChange{
	// 		{
	// 			ConflictType: to.Ptr(core.ConflictTypeNone),
	// 			ItemMetadata: &core.ItemMetadata{
	// 				DisplayName: to.Ptr("My new dataset in the workspace"),
	// 				ItemIdentifier: &core.ItemIdentifier{
	// 					ObjectID: to.Ptr("7753f3b4-dbb8-44c1-a94f-6ae4d776369e"),
	// 				},
	// 				ItemType: to.Ptr(core.ItemTypeSemanticModel),
	// 			},
	// 			WorkspaceChange: to.Ptr(core.ChangeTypeAdded),
	// 		},
	// 		{
	// 			ConflictType: to.Ptr(core.ConflictTypeNone),
	// 			ItemMetadata: &core.ItemMetadata{
	// 				DisplayName: to.Ptr("My deleted report in Git"),
	// 				ItemIdentifier: &core.ItemIdentifier{
	// 					LogicalID: to.Ptr("1423f3b4-dba5-44c1-a94f-6ae4d776369a"),
	// 				},
	// 				ItemType: to.Ptr(core.ItemTypeReport),
	// 			},
	// 			RemoteChange: to.Ptr(core.ChangeTypeDeleted),
	// 		},
	// 		{
	// 			ConflictType: to.Ptr(core.ConflictTypeNone),
	// 			ItemMetadata: &core.ItemMetadata{
	// 				DisplayName: to.Ptr("Modified dataset in the workspace"),
	// 				ItemIdentifier: &core.ItemIdentifier{
	// 					LogicalID: to.Ptr("111e8d7b-4a95-4c02-8ccd-6faef5ba1bd1"),
	// 					ObjectID: to.Ptr("1153f3b4-dbb8-33c1-a84f-6ae4d776362d"),
	// 				},
	// 				ItemType: to.Ptr(core.ItemTypeSemanticModel),
	// 			},
	// 			WorkspaceChange: to.Ptr(core.ChangeTypeModified),
	// 	}},
	// 	RemoteCommitHash: to.Ptr("7d03b2918bf6aa62f96d0a4307293f3853201705"),
	// 	WorkspaceHead: to.Ptr("eaa737b48cda41b37ffefac772ea48f6fed3eac4"),
	// }
}

// Generated from example definition
func ExampleGitClient_BeginGetStatus_getStatusNoChangesExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGitClient().BeginGetStatus(ctx, "1455b6a2-c120-4c1c-dda7-92bafe99bec3", nil)
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
	// res.GitStatusResponse = core.GitStatusResponse{
	// 	Changes: []core.ItemChange{
	// 	},
	// 	RemoteCommitHash: to.Ptr("eaa737b48cda41b37ffefac772ea48f6fed3eac4"),
	// 	WorkspaceHead: to.Ptr("eaa737b48cda41b37ffefac772ea48f6fed3eac4"),
	// }
}

// Generated from example definition
func ExampleGitClient_BeginGetStatus_getStatusWithConflictExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGitClient().BeginGetStatus(ctx, "1455b6a2-c120-4c1c-dda7-92bafe99bec3", nil)
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
	// res.GitStatusResponse = core.GitStatusResponse{
	// 	Changes: []core.ItemChange{
	// 		{
	// 			ConflictType: to.Ptr(core.ConflictTypeConflict),
	// 			ItemMetadata: &core.ItemMetadata{
	// 				DisplayName: to.Ptr("Modified report on both sides"),
	// 				ItemIdentifier: &core.ItemIdentifier{
	// 					LogicalID: to.Ptr("222e8d7b-4a95-4c02-8ccd-6faef5ba1bd2"),
	// 					ObjectID: to.Ptr("8853f3b4-dbb8-33c1-a84f-6ae4d776362a"),
	// 				},
	// 				ItemType: to.Ptr(core.ItemTypeReport),
	// 			},
	// 			RemoteChange: to.Ptr(core.ChangeTypeModified),
	// 			WorkspaceChange: to.Ptr(core.ChangeTypeModified),
	// 	}},
	// 	RemoteCommitHash: to.Ptr("7d03b2918bf6aa62f96d0a4307293f3853201705"),
	// 	WorkspaceHead: to.Ptr("eaa737b48cda41b37ffefac772ea48f6fed3eac4"),
	// }
}

// Generated from example definition
func ExampleGitClient_BeginCommitToGit_commitAllToGitExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGitClient().BeginCommitToGit(ctx, "1565e6a3-c020-4c0c-dda7-92bafe99eec5", core.CommitToGitRequest{
		Comment:       to.Ptr("I'm committing all my changes."),
		Mode:          to.Ptr(core.CommitModeAll),
		WorkspaceHead: to.Ptr("eaa737b48cda41b37ffefac772ea48f6fed3eac4"),
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
func ExampleGitClient_BeginCommitToGit_commitSelectiveItemsToGitExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGitClient().BeginCommitToGit(ctx, "1565e6a3-c020-4c0c-dda7-92bafe99eec5", core.CommitToGitRequest{
		Comment: to.Ptr("I'm committing specific changes."),
		Items: []core.ItemIdentifier{
			{
				LogicalID: to.Ptr("111e8d7b-4a95-4c02-8ccd-6faef5ba1bd1"),
				ObjectID:  to.Ptr("1153f3b4-dbb8-33c1-a84f-6ae4d776362d"),
			},
			{
				ObjectID: to.Ptr("7753f3b4-dbb8-44c1-a94f-6ae4d776369e"),
			}},
		Mode:          to.Ptr(core.CommitModeSelective),
		WorkspaceHead: to.Ptr("eaa737b48cda41b37ffefac772ea48f6fed3eac4"),
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
func ExampleGitClient_GetMyGitCredentials_getTheUsersGitCredentialsConfigurationForAzureDevOpsWhenItIsAutomaticExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGitClient().GetMyGitCredentials(ctx, "1565e6a3-c020-4c0c-dda7-92bafe99eec5", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = core.GitClientGetMyGitCredentialsResponse{
	// 	                            GitCredentialsConfigurationResponseClassification: &core.AutomaticGitCredentialsResponse{
	// 		Source: to.Ptr(core.GitCredentialsSourceAutomatic),
	// 	},
	// 	                        }
}

// Generated from example definition
func ExampleGitClient_GetMyGitCredentials_getTheUsersGitCredentialsConfigurationWhenItIsConfiguredByConnectionExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGitClient().GetMyGitCredentials(ctx, "1565e6a3-c020-4c0c-dda7-92bafe99eec5", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = core.GitClientGetMyGitCredentialsResponse{
	// 	                            GitCredentialsConfigurationResponseClassification: &core.ConfiguredConnectionGitCredentialsResponse{
	// 		Source: to.Ptr(core.GitCredentialsSourceConfiguredConnection),
	// 		ConnectionID: to.Ptr("3f2504e0-4f89-11d3-9a0c-0305e82c3301"),
	// 	},
	// 	                        }
}

// Generated from example definition
func ExampleGitClient_GetMyGitCredentials_getTheUsersGitCredentialsConfigurationWhenItIsNotConfiguredExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGitClient().GetMyGitCredentials(ctx, "1565e6a3-c020-4c0c-dda7-92bafe99eec5", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = core.GitClientGetMyGitCredentialsResponse{
	// 	                            GitCredentialsConfigurationResponseClassification: &core.NoneGitCredentialsResponse{
	// 		Source: to.Ptr(core.GitCredentialsSourceNone),
	// 	},
	// 	                        }
}

// Generated from example definition
func ExampleGitClient_UpdateMyGitCredentials_updateUsersGitCredentialsToAutomaticExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGitClient().UpdateMyGitCredentials(ctx, "1565e6a3-c020-4c0c-dda7-92bafe99eec5", &core.UpdateGitCredentialsToAutomaticRequest{
		Source: to.Ptr(core.GitCredentialsSourceAutomatic),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = core.GitClientUpdateMyGitCredentialsResponse{
	// 	                            GitCredentialsConfigurationResponseClassification: &core.AutomaticGitCredentialsResponse{
	// 		Source: to.Ptr(core.GitCredentialsSourceAutomatic),
	// 	},
	// 	                        }
}

// Generated from example definition
func ExampleGitClient_UpdateMyGitCredentials_updateUsersGitCredentialsToConfiguredConnectionExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGitClient().UpdateMyGitCredentials(ctx, "1565e6a3-c020-4c0c-dda7-92bafe99eec5", &core.UpdateGitCredentialsToConfiguredConnectionRequest{
		Source:       to.Ptr(core.GitCredentialsSourceConfiguredConnection),
		ConnectionID: to.Ptr("3f2504e0-4f89-11d3-9a0c-0305e82c3301"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = core.GitClientUpdateMyGitCredentialsResponse{
	// 	                            GitCredentialsConfigurationResponseClassification: &core.ConfiguredConnectionGitCredentialsResponse{
	// 		Source: to.Ptr(core.GitCredentialsSourceConfiguredConnection),
	// 		ConnectionID: to.Ptr("3f2504e0-4f89-11d3-9a0c-0305e82c3301"),
	// 	},
	// 	                        }
}

// Generated from example definition
func ExampleGitClient_UpdateMyGitCredentials_updateUsersGitCredentialsToNoneExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGitClient().UpdateMyGitCredentials(ctx, "1565e6a3-c020-4c0c-dda7-92bafe99eec5", &core.UpdateGitCredentialsToNoneRequest{
		Source: to.Ptr(core.GitCredentialsSourceNone),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = core.GitClientUpdateMyGitCredentialsResponse{
	// 	                            GitCredentialsConfigurationResponseClassification: &core.NoneGitCredentialsResponse{
	// 		Source: to.Ptr(core.GitCredentialsSourceNone),
	// 	},
	// 	                        }
}
