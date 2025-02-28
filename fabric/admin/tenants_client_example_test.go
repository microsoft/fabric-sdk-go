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
func ExampleTenantsClient_NewListTenantSettingsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := admin.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTenantsClient().NewListTenantSettingsPager(&admin.TenantsClientListTenantSettingsOptions{ContinuationToken: nil})
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
		// page.TenantSettings = admin.TenantSettings{
		// 	Value: []admin.TenantSetting{
		// 		{
		// 			CanSpecifySecurityGroups: to.Ptr(true),
		// 			Enabled: to.Ptr(true),
		// 			SettingName: to.Ptr("AdminApisIncludeDetailedMetadata"),
		// 			TenantSettingGroup: to.Ptr("AdminApiSettings"),
		// 			Title: to.Ptr("Enhance admin APIs responses with detailed metadata"),
		// 		},
		// 		{
		// 			CanSpecifySecurityGroups: to.Ptr(true),
		// 			Enabled: to.Ptr(true),
		// 			EnabledSecurityGroups: []admin.TenantSettingSecurityGroup{
		// 				{
		// 					Name: to.Ptr("TestComputeCdsa"),
		// 					GraphID: to.Ptr("f51b705f-a409-4d40-9197-c5d5f349e2f0"),
		// 				},
		// 				{
		// 					Name: to.Ptr("TestComputeGroup2"),
		// 					GraphID: to.Ptr("1fecf19f-6e33-41b3-89fa-de8c821f3b79"),
		// 				},
		// 				{
		// 					Name: to.Ptr("TestCertifiers"),
		// 					GraphID: to.Ptr("64bc10f1-1f1b-4a7e-b7a0-c87d89cba2b4"),
		// 			}},
		// 			SettingName: to.Ptr("DatamartTenant"),
		// 			TenantSettingGroup: to.Ptr("DatamartSettings"),
		// 			Title: to.Ptr("Create Datamarts (Preview)"),
		// 		},
		// 		{
		// 			CanSpecifySecurityGroups: to.Ptr(true),
		// 			Enabled: to.Ptr(true),
		// 			SettingName: to.Ptr("CertifyDatasets"),
		// 			TenantSettingGroup: to.Ptr("ExportAndSharing"),
		// 			Title: to.Ptr("Certification"),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleTenantsClient_UpdateTenantSetting() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := admin.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTenantsClient().UpdateTenantSetting(ctx, "PublishToWeb", admin.UpdateTenantSettingRequest{
		Enabled: to.Ptr(true),
		EnabledSecurityGroups: []admin.TenantSettingSecurityGroup{
			{
				Name:    to.Ptr("TestComputeCdsa"),
				GraphID: to.Ptr("f51b705f-a409-4d40-9197-c5d5f349e2f0"),
			}},
		Properties: []admin.TenantSettingProperty{
			{
				Name:  to.Ptr("CreateP2w"),
				Type:  to.Ptr(admin.TenantSettingPropertyTypeBoolean),
				Value: to.Ptr("true"),
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UpdateTenantSettingResponse = admin.UpdateTenantSettingResponse{
	// 	TenantSettings: []admin.TenantSetting{
	// 		{
	// 			CanSpecifySecurityGroups: to.Ptr(true),
	// 			Enabled: to.Ptr(true),
	// 			EnabledSecurityGroups: []admin.TenantSettingSecurityGroup{
	// 				{
	// 					Name: to.Ptr("TestComputeCdsa"),
	// 					GraphID: to.Ptr("f51b705f-a409-4d40-9197-c5d5f349e2f0"),
	// 			}},
	// 			Properties: []admin.TenantSettingProperty{
	// 				{
	// 					Name: to.Ptr("CreateP2w"),
	// 					Type: to.Ptr(admin.TenantSettingPropertyTypeBoolean),
	// 					Value: to.Ptr("true"),
	// 			}},
	// 			SettingName: to.Ptr("PublishToWeb"),
	// 			TenantSettingGroup: to.Ptr("TestSetting"),
	// 			Title: to.Ptr("Sample test tenant setting"),
	// 	}},
	// }
}

// Generated from example definition
func ExampleTenantsClient_NewListCapacitiesTenantSettingsOverridesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := admin.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTenantsClient().NewListCapacitiesTenantSettingsOverridesPager(&admin.TenantsClientListCapacitiesTenantSettingsOverridesOptions{ContinuationToken: nil})
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
		// page.CapacityTenantSettingOverrides = admin.CapacityTenantSettingOverrides{
		// 	ContinuationToken: to.Ptr("MSwxMDAwMCww"),
		// 	ContinuationURI: to.Ptr("https://api.fabric.microsoft.com/v1/admin/capacities/delegatedTenantSettingOverrides?continuationToken=MSwxMDAwMCww"),
		// 	Value: []admin.CapacityTenantSettingOverride{
		// 		{
		// 			ID: to.Ptr("f51b705f-a409-4d40-9197-c5d5f349e2ef"),
		// 			TenantSettings: []admin.CapacityTenantSetting{
		// 				{
		// 					CanSpecifySecurityGroups: to.Ptr(true),
		// 					Enabled: to.Ptr(true),
		// 					EnabledSecurityGroups: []admin.TenantSettingSecurityGroup{
		// 						{
		// 							Name: to.Ptr("TestComputeCdsa"),
		// 							GraphID: to.Ptr("f51b705f-a409-4d40-9197-c5d5f349e2f0"),
		// 						},
		// 						{
		// 							Name: to.Ptr("TestComputeGroup2"),
		// 							GraphID: to.Ptr("1fecf19f-6e33-41b3-89fa-de8c821f3b79"),
		// 						},
		// 						{
		// 							Name: to.Ptr("TestCertifiers"),
		// 							GraphID: to.Ptr("64bc10f1-1f1b-4a7e-b7a0-c87d89cba2b4"),
		// 					}},
		// 					Properties: []admin.TenantSettingProperty{
		// 						{
		// 							Name: to.Ptr("testIntProp"),
		// 							Type: to.Ptr(admin.TenantSettingPropertyTypeInteger),
		// 							Value: to.Ptr("5"),
		// 					}},
		// 					SettingName: to.Ptr("TenantSettingForCapacityDelegatedSwitch"),
		// 					TenantSettingGroup: to.Ptr("Delegation testing"),
		// 					Title: to.Ptr("Capacity delegation test settings"),
		// 					DelegateToWorkspace: to.Ptr(false),
		// 					DelegatedFrom: to.Ptr(admin.DelegatedFromTenant),
		// 			}},
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleTenantsClient_NewListCapacityTenantSettingsOverridesByCapacityIDPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := admin.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTenantsClient().NewListCapacityTenantSettingsOverridesByCapacityIDPager("f51b705f-a409-4d40-9197-c5d5f349e2ef", &admin.TenantsClientListCapacityTenantSettingsOverridesByCapacityIDOptions{ContinuationToken: nil})
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
		// page.CapacityTenantSettingsByCapacityIDResponse = admin.CapacityTenantSettingsByCapacityIDResponse{
		// 	Value: []admin.CapacityTenantSetting{
		// 		{
		// 			CanSpecifySecurityGroups: to.Ptr(true),
		// 			Enabled: to.Ptr(true),
		// 			EnabledSecurityGroups: []admin.TenantSettingSecurityGroup{
		// 				{
		// 					Name: to.Ptr("TestComputeCdsa"),
		// 					GraphID: to.Ptr("f51b705f-a409-4d40-9197-c5d5f349e2f0"),
		// 				},
		// 				{
		// 					Name: to.Ptr("TestComputeGroup2"),
		// 					GraphID: to.Ptr("1fecf19f-6e33-41b3-89fa-de8c821f3b79"),
		// 				},
		// 				{
		// 					Name: to.Ptr("TestCertifiers"),
		// 					GraphID: to.Ptr("64bc10f1-1f1b-4a7e-b7a0-c87d89cba2b4"),
		// 			}},
		// 			Properties: []admin.TenantSettingProperty{
		// 				{
		// 					Name: to.Ptr("testIntProp"),
		// 					Type: to.Ptr(admin.TenantSettingPropertyTypeInteger),
		// 					Value: to.Ptr("5"),
		// 			}},
		// 			SettingName: to.Ptr("TenantSettingForCapacityDelegatedSwitch"),
		// 			TenantSettingGroup: to.Ptr("Delegation testing"),
		// 			Title: to.Ptr("Capacity delegation test settings"),
		// 			DelegateToWorkspace: to.Ptr(false),
		// 			DelegatedFrom: to.Ptr(admin.DelegatedFromTenant),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleTenantsClient_DeleteCapacityTenantSettingOverride() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := admin.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTenantsClient().DeleteCapacityTenantSettingOverride(ctx, "f51b705f-a409-4d40-9197-c5d5f349e2ef", "TenantSettingForCapacityDelegatedSwitch", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleTenantsClient_UpdateCapacityTenantSettingOverride() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := admin.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTenantsClient().UpdateCapacityTenantSettingOverride(ctx, "f51b705f-a409-4d40-9197-c5d5f349e2ef", "AdminApisIncludeDetailedMetadata", admin.UpdateCapacityTenantSettingOverrideRequest{
		DelegateToWorkspace: to.Ptr(true),
		Enabled:             to.Ptr(true),
		EnabledSecurityGroups: []admin.TenantSettingSecurityGroup{
			{
				Name:    to.Ptr("TestComputeCdsa"),
				GraphID: to.Ptr("f51b705f-a409-4d40-9197-c5d5f349e2f0"),
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UpdateCapacityTenantSettingOverrideResponse = admin.UpdateCapacityTenantSettingOverrideResponse{
	// 	Overrides: []admin.CapacityTenantSetting{
	// 		{
	// 			CanSpecifySecurityGroups: to.Ptr(true),
	// 			Enabled: to.Ptr(true),
	// 			EnabledSecurityGroups: []admin.TenantSettingSecurityGroup{
	// 				{
	// 					Name: to.Ptr("TestComputeCdsa"),
	// 					GraphID: to.Ptr("f51b705f-a409-4d40-9197-c5d5f349e2f0"),
	// 			}},
	// 			Properties: []admin.TenantSettingProperty{
	// 				{
	// 					Name: to.Ptr("testIntProp"),
	// 					Type: to.Ptr(admin.TenantSettingPropertyTypeInteger),
	// 					Value: to.Ptr("5"),
	// 			}},
	// 			SettingName: to.Ptr("TenantSettingForCapacityDelegatedSwitch"),
	// 			TenantSettingGroup: to.Ptr("Delegation testing"),
	// 			Title: to.Ptr("Capacity delegation test settings"),
	// 			DelegateToWorkspace: to.Ptr(true),
	// 			DelegatedFrom: to.Ptr(admin.DelegatedFromTenant),
	// 	}},
	// }
}

// Generated from example definition
func ExampleTenantsClient_NewListDomainsTenantSettingsOverridesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := admin.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTenantsClient().NewListDomainsTenantSettingsOverridesPager(&admin.TenantsClientListDomainsTenantSettingsOverridesOptions{ContinuationToken: nil})
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
		// page.DomainTenantSettingOverrides = admin.DomainTenantSettingOverrides{
		// 	ContinuationToken: to.Ptr("MSwxMDAwMCww"),
		// 	ContinuationURI: to.Ptr("https://api.fabric.microsoft.com/v1/admin/domains/delegatedTenantSettingOverrides?continuationToken=MSwxMDAwMCww"),
		// 	Value: []admin.DomainTenantSettingOverride{
		// 		{
		// 			ID: to.Ptr("f51b705f-a409-4d40-9197-c5d5f349e2ef"),
		// 			TenantSettings: []admin.DomainTenantSetting{
		// 				{
		// 					CanSpecifySecurityGroups: to.Ptr(true),
		// 					Enabled: to.Ptr(true),
		// 					EnabledSecurityGroups: []admin.TenantSettingSecurityGroup{
		// 						{
		// 							Name: to.Ptr("Admin API SP Test"),
		// 							GraphID: to.Ptr("f51b705f-a409-4d40-9197-c5d5f349e2f0"),
		// 						},
		// 						{
		// 							Name: to.Ptr("Admin API Testing"),
		// 							GraphID: to.Ptr("1fecf19f-6e33-41b3-89fa-de8c821f3b79"),
		// 						},
		// 						{
		// 							Name: to.Ptr("Admin Only"),
		// 							GraphID: to.Ptr("64bc10f1-1f1b-4a7e-b7a0-c87d89cba2b4"),
		// 					}},
		// 					SettingName: to.Ptr("TenantSettingForDomainDelegatedSwitch"),
		// 					TenantSettingGroup: to.Ptr("Delegation testing"),
		// 					Title: to.Ptr("Domain delegation test settings"),
		// 					DelegateToWorkspace: to.Ptr(false),
		// 					DelegatedFrom: to.Ptr(admin.DelegatedFromTenant),
		// 			}},
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleTenantsClient_NewListWorkspacesTenantSettingsOverridesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := admin.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTenantsClient().NewListWorkspacesTenantSettingsOverridesPager(&admin.TenantsClientListWorkspacesTenantSettingsOverridesOptions{ContinuationToken: nil})
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
		// page.WorkspaceTenantSettingOverrides = admin.WorkspaceTenantSettingOverrides{
		// 	ContinuationToken: to.Ptr("MSwxMDAwMCww"),
		// 	ContinuationURI: to.Ptr("https://api.fabric.microsoft.com/v1/admin/workspaces/delegatedTenantSettingOverrides?continuationToken=MSwxMDAwMCww"),
		// 	Value: []admin.WorkspaceTenantSettingOverride{
		// 		{
		// 			ID: to.Ptr("f51b705f-a409-4d40-9197-c5d5f349e2ef"),
		// 			TenantSettings: []admin.WorkspaceTenantSetting{
		// 				{
		// 					CanSpecifySecurityGroups: to.Ptr(true),
		// 					Enabled: to.Ptr(true),
		// 					EnabledSecurityGroups: []admin.TenantSettingSecurityGroup{
		// 						{
		// 							Name: to.Ptr("Admin API SP Test"),
		// 							GraphID: to.Ptr("f51b705f-a409-4d40-9197-c5d5f349e2f0"),
		// 						},
		// 						{
		// 							Name: to.Ptr("Admin API Testing"),
		// 							GraphID: to.Ptr("1fecf19f-6e33-41b3-89fa-de8c821f3b79"),
		// 						},
		// 						{
		// 							Name: to.Ptr("Admin Only"),
		// 							GraphID: to.Ptr("64bc10f1-1f1b-4a7e-b7a0-c87d89cba2b4"),
		// 					}},
		// 					SettingName: to.Ptr("TenantSettingForWorkspaceDelegatedSwitch"),
		// 					TenantSettingGroup: to.Ptr("Delegation testing"),
		// 					Title: to.Ptr("Workspace delegation test settings"),
		// 					DelegatedFrom: to.Ptr(admin.DelegatedFromTenant),
		// 			}},
		// 	}},
		// }
	}
}
