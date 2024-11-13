// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package admin

// AssignmentMethod - Specifies whether the assigned label was set by an automated process or manually. Additional tenant
// setting property types may be added over time.
type AssignmentMethod string

const (
	// AssignmentMethodPriviledged - The label was set manually.
	AssignmentMethodPriviledged AssignmentMethod = "Priviledged"
	// AssignmentMethodStandard - The label was set by an automated process (default value).
	AssignmentMethodStandard AssignmentMethod = "Standard"
)

// PossibleAssignmentMethodValues returns the possible values for the AssignmentMethod const type.
func PossibleAssignmentMethodValues() []AssignmentMethod {
	return []AssignmentMethod{
		AssignmentMethodPriviledged,
		AssignmentMethodStandard,
	}
}

// Category - The category of the item type. Additional category types may be added over time.
type Category string

const (
	// CategoryItem - Fabric items such as Notebook, Synapse and KQL Database.
	CategoryItem Category = "Item"
)

// PossibleCategoryValues returns the possible values for the Category const type.
func PossibleCategoryValues() []Category {
	return []Category{
		CategoryItem,
	}
}

// ContributorsScopeType - The contributor scope. Additional contributor scopes may be added over time.
type ContributorsScopeType string

const (
	// ContributorsScopeTypeAdminsOnly - Tenant and domain admins only.
	ContributorsScopeTypeAdminsOnly ContributorsScopeType = "AdminsOnly"
	// ContributorsScopeTypeAllTenant - All the tenant's users.
	ContributorsScopeTypeAllTenant ContributorsScopeType = "AllTenant"
	// ContributorsScopeTypeSpecificUsersAndGroups - Specific users and groups.
	ContributorsScopeTypeSpecificUsersAndGroups ContributorsScopeType = "SpecificUsersAndGroups"
)

// PossibleContributorsScopeTypeValues returns the possible values for the ContributorsScopeType const type.
func PossibleContributorsScopeTypeValues() []ContributorsScopeType {
	return []ContributorsScopeType{
		ContributorsScopeTypeAdminsOnly,
		ContributorsScopeTypeAllTenant,
		ContributorsScopeTypeSpecificUsersAndGroups,
	}
}

// DelegatedFrom - The Fabric component (workspace, capacity or domain) that the tenant setting was delegated from. Additional
// DelegatedFrom may be added over time.
type DelegatedFrom string

const (
	// DelegatedFromCapacity - The setting is delegated from a capacity.
	DelegatedFromCapacity DelegatedFrom = "Capacity"
	// DelegatedFromDomain - The setting is delegated from a domain.
	DelegatedFromDomain DelegatedFrom = "Domain"
	// DelegatedFromTenant - The setting is delegated from a tenant.
	DelegatedFromTenant DelegatedFrom = "Tenant"
)

// PossibleDelegatedFromValues returns the possible values for the DelegatedFrom const type.
func PossibleDelegatedFromValues() []DelegatedFrom {
	return []DelegatedFrom{
		DelegatedFromCapacity,
		DelegatedFromDomain,
		DelegatedFromTenant,
	}
}

// DomainRole - Represents the domain members by the principal's request type. Additional request types may be added over
// time.
type DomainRole string

const (
	// DomainRoleAdmins - Domain admins request type.
	DomainRoleAdmins DomainRole = "Admins"
	// DomainRoleContributors - Domain contributors request type.
	DomainRoleContributors DomainRole = "Contributors"
)

// PossibleDomainRoleValues returns the possible values for the DomainRole const type.
func PossibleDomainRoleValues() []DomainRole {
	return []DomainRole{
		DomainRoleAdmins,
		DomainRoleContributors,
	}
}

// ExternalDataShareStatus - The status of a given external data share. Additional ExternalDataShareStatus types may be added
// over time.
type ExternalDataShareStatus string

const (
	// ExternalDataShareStatusActive - The invitation has been accepted by the recipient and the external data share is active.
	ExternalDataShareStatusActive ExternalDataShareStatus = "Active"
	// ExternalDataShareStatusInvitationExpired - The invitation expired and can no longer be accepted by the recipient.
	ExternalDataShareStatusInvitationExpired ExternalDataShareStatus = "InvitationExpired"
	// ExternalDataShareStatusPending - An invitation was created and is now pending acceptance by the recipient.
	ExternalDataShareStatusPending ExternalDataShareStatus = "Pending"
	// ExternalDataShareStatusRevoked - The external data share was revoked.
	ExternalDataShareStatusRevoked ExternalDataShareStatus = "Revoked"
)

// PossibleExternalDataShareStatusValues returns the possible values for the ExternalDataShareStatus const type.
func PossibleExternalDataShareStatusValues() []ExternalDataShareStatus {
	return []ExternalDataShareStatus{
		ExternalDataShareStatusActive,
		ExternalDataShareStatusInvitationExpired,
		ExternalDataShareStatusPending,
		ExternalDataShareStatusRevoked,
	}
}

// GitProviderType - A Git provider type. Additional provider types may be added over time.
type GitProviderType string

const (
	// GitProviderTypeAzureDevOps - Azure DevOps provider
	GitProviderTypeAzureDevOps GitProviderType = "AzureDevOps"
	// GitProviderTypeGitHub - GitHub provider
	GitProviderTypeGitHub GitProviderType = "GitHub"
)

// PossibleGitProviderTypeValues returns the possible values for the GitProviderType const type.
func PossibleGitProviderTypeValues() []GitProviderType {
	return []GitProviderType{
		GitProviderTypeAzureDevOps,
		GitProviderTypeGitHub,
	}
}

// GroupType - The type of the group. Additional group types may be added over time.
type GroupType string

const (
	// GroupTypeDistributionList - Principal is a distribution list.
	GroupTypeDistributionList GroupType = "DistributionList"
	// GroupTypeSecurityGroup - Principal is a security group.
	GroupTypeSecurityGroup GroupType = "SecurityGroup"
	// GroupTypeUnknown - Principal group type is unknown.
	GroupTypeUnknown GroupType = "Unknown"
)

// PossibleGroupTypeValues returns the possible values for the GroupType const type.
func PossibleGroupTypeValues() []GroupType {
	return []GroupType{
		GroupTypeDistributionList,
		GroupTypeSecurityGroup,
		GroupTypeUnknown,
	}
}

// ItemPermissions - Item permissions. Additional item permissions may be added over time.
type ItemPermissions string

const (
	// ItemPermissionsExecute - User can execute and cancel item jobs.
	ItemPermissionsExecute ItemPermissions = "Execute"
	// ItemPermissionsExplore - User can build items on other items.
	ItemPermissionsExplore ItemPermissions = "Explore"
	// ItemPermissionsRead - User can read the metadata about an item.
	ItemPermissionsRead ItemPermissions = "Read"
	// ItemPermissionsReshare - User can share an item with other users.
	ItemPermissionsReshare ItemPermissions = "Reshare"
	// ItemPermissionsWrite - User can perform write operations on an item.
	ItemPermissionsWrite ItemPermissions = "Write"
)

// PossibleItemPermissionsValues returns the possible values for the ItemPermissions const type.
func PossibleItemPermissionsValues() []ItemPermissions {
	return []ItemPermissions{
		ItemPermissionsExecute,
		ItemPermissionsExplore,
		ItemPermissionsRead,
		ItemPermissionsReshare,
		ItemPermissionsWrite,
	}
}

// ItemState - The item state. Additional item states may be added over time.
type ItemState string

const (
	// ItemStateActive - An active item.
	ItemStateActive ItemState = "Active"
)

// PossibleItemStateValues returns the possible values for the ItemState const type.
func PossibleItemStateValues() []ItemState {
	return []ItemState{
		ItemStateActive,
	}
}

// ItemType - The type of the item. Additional item types may be added over time.
type ItemType string

const (
	// ItemTypeDashboard - PowerBI dashboard.
	ItemTypeDashboard ItemType = "Dashboard"
	// ItemTypeDataPipeline - A data pipeline.
	ItemTypeDataPipeline ItemType = "DataPipeline"
	// ItemTypeDatamart - PowerBI datamart.
	ItemTypeDatamart ItemType = "Datamart"
	// ItemTypeEnvironment - An environment.
	ItemTypeEnvironment ItemType = "Environment"
	// ItemTypeEventhouse - An eventhouse.
	ItemTypeEventhouse ItemType = "Eventhouse"
	// ItemTypeEventstream - An eventstream.
	ItemTypeEventstream ItemType = "Eventstream"
	// ItemTypeKQLDashboard - A KQL dashboard.
	ItemTypeKQLDashboard ItemType = "KQLDashboard"
	// ItemTypeKQLDatabase - A KQL database.
	ItemTypeKQLDatabase ItemType = "KQLDatabase"
	// ItemTypeKQLQueryset - A KQL queryset.
	ItemTypeKQLQueryset ItemType = "KQLQueryset"
	// ItemTypeLakehouse - A lakehouse.
	ItemTypeLakehouse ItemType = "Lakehouse"
	// ItemTypeMLExperiment - A machine learning experiment.
	ItemTypeMLExperiment ItemType = "MLExperiment"
	// ItemTypeMLModel - A machine learning model.
	ItemTypeMLModel ItemType = "MLModel"
	// ItemTypeMirroredDatabase - A mirrored database.
	ItemTypeMirroredDatabase ItemType = "MirroredDatabase"
	// ItemTypeMirroredWarehouse - A mirrored warehouse.
	ItemTypeMirroredWarehouse ItemType = "MirroredWarehouse"
	// ItemTypeNotebook - A notebook.
	ItemTypeNotebook ItemType = "Notebook"
	// ItemTypePaginatedReport - PowerBI paginated report.
	ItemTypePaginatedReport ItemType = "PaginatedReport"
	// ItemTypeReport - PowerBI report.
	ItemTypeReport ItemType = "Report"
	// ItemTypeSQLEndpoint - An SQL endpoint.
	ItemTypeSQLEndpoint ItemType = "SQLEndpoint"
	// ItemTypeSemanticModel - PowerBI semantic model.
	ItemTypeSemanticModel ItemType = "SemanticModel"
	// ItemTypeSparkJobDefinition - A spark job definition.
	ItemTypeSparkJobDefinition ItemType = "SparkJobDefinition"
	// ItemTypeWarehouse - A warehouse.
	ItemTypeWarehouse ItemType = "Warehouse"
)

// PossibleItemTypeValues returns the possible values for the ItemType const type.
func PossibleItemTypeValues() []ItemType {
	return []ItemType{
		ItemTypeDashboard,
		ItemTypeDataPipeline,
		ItemTypeDatamart,
		ItemTypeEnvironment,
		ItemTypeEventhouse,
		ItemTypeEventstream,
		ItemTypeKQLDashboard,
		ItemTypeKQLDatabase,
		ItemTypeKQLQueryset,
		ItemTypeLakehouse,
		ItemTypeMLExperiment,
		ItemTypeMLModel,
		ItemTypeMirroredDatabase,
		ItemTypeMirroredWarehouse,
		ItemTypeNotebook,
		ItemTypePaginatedReport,
		ItemTypeReport,
		ItemTypeSQLEndpoint,
		ItemTypeSemanticModel,
		ItemTypeSparkJobDefinition,
		ItemTypeWarehouse,
	}
}

// PrincipalType - The type of the principal. Additional principal types may be added over time.
type PrincipalType string

const (
	// PrincipalTypeGroup - Principal is a security group.
	PrincipalTypeGroup PrincipalType = "Group"
	// PrincipalTypeServicePrincipal - Principal is a Microsoft Entra service principal.
	PrincipalTypeServicePrincipal PrincipalType = "ServicePrincipal"
	// PrincipalTypeServicePrincipalProfile - Principal is a service principal profile.
	PrincipalTypeServicePrincipalProfile PrincipalType = "ServicePrincipalProfile"
	// PrincipalTypeUser - Principal is a Microsoft Entra user principal.
	PrincipalTypeUser PrincipalType = "User"
)

// PossiblePrincipalTypeValues returns the possible values for the PrincipalType const type.
func PossiblePrincipalTypeValues() []PrincipalType {
	return []PrincipalType{
		PrincipalTypeGroup,
		PrincipalTypeServicePrincipal,
		PrincipalTypeServicePrincipalProfile,
		PrincipalTypeUser,
	}
}

// Status - The status of an information protection label change operation. Additional tenant setting property types may be
// added over time.
type Status string

const (
	// StatusFailed - Failed to set a new label. Please retry.
	StatusFailed Status = "Failed"
	// StatusFailedToGetUsageRights - Failed to set a new label. The Fabric item has a sensitivity label with protection settings,
	// and Fabric was unable to verify that the user has sufficient usage rights to change the label.
	StatusFailedToGetUsageRights Status = "FailedToGetUsageRights"
	// StatusInsufficientUsageRights - Failed to set a new label. The Fabric item has a sensitivity label with protection settings,
	// and the admin user (and the delegated user, if provided) doesn't have sufficient usage rights to change the label.
	StatusInsufficientUsageRights Status = "InsufficientUsageRights"
	// StatusNotFound - The Fabric item ID, label or type wasn't found.
	StatusNotFound Status = "NotFound"
	// StatusSucceeded - The Fabric item label was changed.
	StatusSucceeded Status = "Succeeded"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusFailed,
		StatusFailedToGetUsageRights,
		StatusInsufficientUsageRights,
		StatusNotFound,
		StatusSucceeded,
	}
}

// TenantSettingPropertyType - Tenant setting property type. Additional tenant setting property types may be added over time.
type TenantSettingPropertyType string

const (
	// TenantSettingPropertyTypeBoolean - A checkbox in the UI.
	TenantSettingPropertyTypeBoolean TenantSettingPropertyType = "Boolean"
	// TenantSettingPropertyTypeFreeText - UI accepts any string for the text box.
	TenantSettingPropertyTypeFreeText TenantSettingPropertyType = "FreeText"
	// TenantSettingPropertyTypeInteger - UI accepts only integers for the text box.
	TenantSettingPropertyTypeInteger TenantSettingPropertyType = "Integer"
	// TenantSettingPropertyTypeMailEnabledSecurityGroup - UI accepts only email enabled security groups for the text box.
	TenantSettingPropertyTypeMailEnabledSecurityGroup TenantSettingPropertyType = "MailEnabledSecurityGroup"
	// TenantSettingPropertyTypeURL - UI accepts only URLs for the text box.
	TenantSettingPropertyTypeURL TenantSettingPropertyType = "Url"
)

// PossibleTenantSettingPropertyTypeValues returns the possible values for the TenantSettingPropertyType const type.
func PossibleTenantSettingPropertyTypeValues() []TenantSettingPropertyType {
	return []TenantSettingPropertyType{
		TenantSettingPropertyTypeBoolean,
		TenantSettingPropertyTypeFreeText,
		TenantSettingPropertyTypeInteger,
		TenantSettingPropertyTypeMailEnabledSecurityGroup,
		TenantSettingPropertyTypeURL,
	}
}

// WorkspaceRole - A Workspace role. Additional workspace roles may be added over time.
type WorkspaceRole string

const (
	// WorkspaceRoleAdmin - Enables administrative access to the workspace.
	WorkspaceRoleAdmin WorkspaceRole = "Admin"
	// WorkspaceRoleContributor - Enables contribution to the workspace.
	WorkspaceRoleContributor WorkspaceRole = "Contributor"
	// WorkspaceRoleMember - Enables membership access to the workspace.
	WorkspaceRoleMember WorkspaceRole = "Member"
	// WorkspaceRoleViewer - Enables viewing of the workspace.
	WorkspaceRoleViewer WorkspaceRole = "Viewer"
)

// PossibleWorkspaceRoleValues returns the possible values for the WorkspaceRole const type.
func PossibleWorkspaceRoleValues() []WorkspaceRole {
	return []WorkspaceRole{
		WorkspaceRoleAdmin,
		WorkspaceRoleContributor,
		WorkspaceRoleMember,
		WorkspaceRoleViewer,
	}
}

// WorkspaceState - The workspace state. Additional workspace states may be added over time.
type WorkspaceState string

const (
	// WorkspaceStateActive - The workspace is active. Orphaned workspaces are displayed as active.
	WorkspaceStateActive WorkspaceState = "Active"
	// WorkspaceStateDeleted - The workspace is deleted.
	WorkspaceStateDeleted WorkspaceState = "Deleted"
)

// PossibleWorkspaceStateValues returns the possible values for the WorkspaceState const type.
func PossibleWorkspaceStateValues() []WorkspaceState {
	return []WorkspaceState{
		WorkspaceStateActive,
		WorkspaceStateDeleted,
	}
}

// WorkspaceType - A workspace type. Additional workspace types may be added over time.
type WorkspaceType string

const (
	// WorkspaceTypeAdminWorkspace - Admin monitoring workspace. Contains admin reports such as the audit report and the usage
	// and adoption report.
	WorkspaceTypeAdminWorkspace WorkspaceType = "AdminWorkspace"
	// WorkspaceTypePersonal - My folder or My workspace used to manage user items.
	WorkspaceTypePersonal WorkspaceType = "Personal"
	// WorkspaceTypeWorkspace - Workspace used to manage the Fabric items.
	WorkspaceTypeWorkspace WorkspaceType = "Workspace"
)

// PossibleWorkspaceTypeValues returns the possible values for the WorkspaceType const type.
func PossibleWorkspaceTypeValues() []WorkspaceType {
	return []WorkspaceType{
		WorkspaceTypeAdminWorkspace,
		WorkspaceTypePersonal,
		WorkspaceTypeWorkspace,
	}
}
