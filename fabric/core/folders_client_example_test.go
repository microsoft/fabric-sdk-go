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
func ExampleFoldersClient_NewListFoldersPager_listAllFoldersInWorkspaceExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFoldersClient().NewListFoldersPager("aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb", &core.FoldersClientListFoldersOptions{RootFolderID: nil,
		Recursive:         nil,
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
		// page.Folders = core.Folders{
		// 	Value: []core.Folder{
		// 		{
		// 			DisplayName: to.Ptr("Sales"),
		// 			ID: to.Ptr("aaaaaaaa-6666-7777-8888-bbbbbbbbbbbb"),
		// 			WorkspaceID: to.Ptr("aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb"),
		// 		},
		// 		{
		// 			DisplayName: to.Ptr("Y2024"),
		// 			ID: to.Ptr("bbbbbbbb-1111-2222-3333-cccccccccccc"),
		// 			ParentFolderID: to.Ptr("aaaaaaaa-6666-7777-8888-bbbbbbbbbbbb"),
		// 			WorkspaceID: to.Ptr("aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb"),
		// 		},
		// 		{
		// 			DisplayName: to.Ptr("Q1"),
		// 			ID: to.Ptr("cccccccc-8888-9999-0000-dddddddddddd"),
		// 			ParentFolderID: to.Ptr("bbbbbbbb-1111-2222-3333-cccccccccccc"),
		// 			WorkspaceID: to.Ptr("aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb"),
		// 		},
		// 		{
		// 			DisplayName: to.Ptr("Q2"),
		// 			ID: to.Ptr("dddddddd-9999-0000-1111-eeeeeeeeeeee"),
		// 			ParentFolderID: to.Ptr("bbbbbbbb-1111-2222-3333-cccccccccccc"),
		// 			WorkspaceID: to.Ptr("aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb"),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleFoldersClient_NewListFoldersPager_listAllFoldersInWorkspaceWithContinuationExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFoldersClient().NewListFoldersPager("aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb", &core.FoldersClientListFoldersOptions{RootFolderID: nil,
		Recursive:         nil,
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
		// page.Folders = core.Folders{
		// 	ContinuationToken: to.Ptr("MAEsMTbwMDAwLDA%5D"),
		// 	ContinuationURI: to.Ptr("https://api.fabric.microsoft.com/v1/workspaces/aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb/folders?continuationToken=MAEsMTbwMDAwLDA%5D"),
		// 	Value: []core.Folder{
		// 		{
		// 			DisplayName: to.Ptr("Sales"),
		// 			ID: to.Ptr("aaaaaaaa-6666-7777-8888-bbbbbbbbbbbb"),
		// 			WorkspaceID: to.Ptr("aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb"),
		// 		},
		// 		{
		// 			DisplayName: to.Ptr("Y2024"),
		// 			ID: to.Ptr("bbbbbbbb-1111-2222-3333-cccccccccccc"),
		// 			ParentFolderID: to.Ptr("aaaaaaaa-6666-7777-8888-bbbbbbbbbbbb"),
		// 			WorkspaceID: to.Ptr("aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb"),
		// 		},
		// 		{
		// 			DisplayName: to.Ptr("Q1"),
		// 			ID: to.Ptr("cccccccc-8888-9999-0000-dddddddddddd"),
		// 			ParentFolderID: to.Ptr("bbbbbbbb-1111-2222-3333-cccccccccccc"),
		// 			WorkspaceID: to.Ptr("aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb"),
		// 		},
		// 		{
		// 			DisplayName: to.Ptr("Q2"),
		// 			ID: to.Ptr("dddddddd-9999-0000-1111-eeeeeeeeeeee"),
		// 			ParentFolderID: to.Ptr("bbbbbbbb-1111-2222-3333-cccccccccccc"),
		// 			WorkspaceID: to.Ptr("aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb"),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleFoldersClient_NewListFoldersPager_listDirectChildrenFoldersUnderParentFolderExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFoldersClient().NewListFoldersPager("aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb", &core.FoldersClientListFoldersOptions{RootFolderID: to.Ptr("aaaaaaaa-6666-7777-8888-bbbbbbbbbbbb"),
		Recursive:         to.Ptr(false),
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
		// page.Folders = core.Folders{
		// 	Value: []core.Folder{
		// 		{
		// 			DisplayName: to.Ptr("Y2024"),
		// 			ID: to.Ptr("bbbbbbbb-1111-2222-3333-cccccccccccc"),
		// 			ParentFolderID: to.Ptr("aaaaaaaa-6666-7777-8888-bbbbbbbbbbbb"),
		// 			WorkspaceID: to.Ptr("aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb"),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleFoldersClient_NewListFoldersPager_listFoldersUnderParentFolderRecursivelyExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFoldersClient().NewListFoldersPager("aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb", &core.FoldersClientListFoldersOptions{RootFolderID: to.Ptr("aaaaaaaa-6666-7777-8888-bbbbbbbbbbbb"),
		Recursive:         to.Ptr(true),
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
		// page.Folders = core.Folders{
		// 	Value: []core.Folder{
		// 		{
		// 			DisplayName: to.Ptr("Y2024"),
		// 			ID: to.Ptr("bbbbbbbb-1111-2222-3333-cccccccccccc"),
		// 			ParentFolderID: to.Ptr("aaaaaaaa-6666-7777-8888-bbbbbbbbbbbb"),
		// 			WorkspaceID: to.Ptr("aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb"),
		// 		},
		// 		{
		// 			DisplayName: to.Ptr("Q1"),
		// 			ID: to.Ptr("cccccccc-8888-9999-0000-dddddddddddd"),
		// 			ParentFolderID: to.Ptr("bbbbbbbb-1111-2222-3333-cccccccccccc"),
		// 			WorkspaceID: to.Ptr("aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb"),
		// 		},
		// 		{
		// 			DisplayName: to.Ptr("Q2"),
		// 			ID: to.Ptr("dddddddd-9999-0000-1111-eeeeeeeeeeee"),
		// 			ParentFolderID: to.Ptr("bbbbbbbb-1111-2222-3333-cccccccccccc"),
		// 			WorkspaceID: to.Ptr("aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb"),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleFoldersClient_CreateFolder_createAFolderUnderAnotherFolderExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewFoldersClient().CreateFolder(ctx, "aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb", core.CreateFolderRequest{
		DisplayName:    to.Ptr("Q3"),
		ParentFolderID: to.Ptr("bbbbbbbb-1111-2222-3333-cccccccccccc"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleFoldersClient_CreateFolder_createAFolderWithTheWorkspaceAsItsParentFolderExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewFoldersClient().CreateFolder(ctx, "aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb", core.CreateFolderRequest{
		DisplayName: to.Ptr("A folder"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleFoldersClient_GetFolder() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFoldersClient().GetFolder(ctx, "aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb", "bbbbbbbb-1111-2222-3333-cccccccccccc", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Folder = core.Folder{
	// 	DisplayName: to.Ptr("A nested folder"),
	// 	ID: to.Ptr("bbbbbbbb-1111-2222-3333-cccccccccccc"),
	// 	ParentFolderID: to.Ptr("aaaaaaaa-6666-7777-8888-bbbbbbbbbbbb"),
	// 	WorkspaceID: to.Ptr("aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb"),
	// }
}

// Generated from example definition
func ExampleFoldersClient_UpdateFolder() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFoldersClient().UpdateFolder(ctx, "aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb", "aaaaaaaa-6666-7777-8888-bbbbbbbbbbbb", core.UpdateFolderRequest{
		DisplayName: to.Ptr("A new Folder name"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Folder = core.Folder{
	// 	DisplayName: to.Ptr("A new Folder name"),
	// 	ID: to.Ptr("aaaaaaaa-6666-7777-8888-bbbbbbbbbbbb"),
	// 	WorkspaceID: to.Ptr("aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb"),
	// }
}

// Generated from example definition
func ExampleFoldersClient_DeleteFolder() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewFoldersClient().DeleteFolder(ctx, "aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb", "bbbbbbbb-1111-2222-3333-cccccccccccc", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleFoldersClient_MoveFolder_moveAFolderIntoAnotherFolderExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFoldersClient().MoveFolder(ctx, "aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb", "dddddddd-9999-0000-1111-eeeeeeeeeeee", core.MoveFolderRequest{
		TargetFolderID: to.Ptr("cccccccc-8888-9999-0000-dddddddddddd"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Folder = core.Folder{
	// 	DisplayName: to.Ptr("Q2"),
	// 	ID: to.Ptr("dddddddd-9999-0000-1111-eeeeeeeeeeee"),
	// 	ParentFolderID: to.Ptr("cccccccc-8888-9999-0000-dddddddddddd"),
	// 	WorkspaceID: to.Ptr("aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb"),
	// }
}

// Generated from example definition
func ExampleFoldersClient_MoveFolder_moveAFolderWithTheWorkspaceAsTheDestinationExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := core.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFoldersClient().MoveFolder(ctx, "aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb", "dddddddd-9999-0000-1111-eeeeeeeeeeee", core.MoveFolderRequest{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Folder = core.Folder{
	// 	DisplayName: to.Ptr("Q2"),
	// 	ID: to.Ptr("dddddddd-9999-0000-1111-eeeeeeeeeeee"),
	// 	WorkspaceID: to.Ptr("aaaaaaaa-0000-1111-2222-bbbbbbbbbbbb"),
	// }
}
