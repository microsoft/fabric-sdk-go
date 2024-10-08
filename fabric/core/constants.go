// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package core

// AttributeName - Specifies the name of the attribute that is being evaluated for access permissions. AttributeName can be
// Path or Action. Additional attributeName types may be added over time.
type AttributeName string

const (
	// AttributeNameAction - Attribute name Action
	AttributeNameAction AttributeName = "Action"
	// AttributeNamePath - Attribute name Path
	AttributeNamePath AttributeName = "Path"
)

// PossibleAttributeNameValues returns the possible values for the AttributeName const type.
func PossibleAttributeNameValues() []AttributeName {
	return []AttributeName{
		AttributeNameAction,
		AttributeNamePath,
	}
}

// CapacityAssignmentProgress - A Workspace assignment to capacity progress status. Additional capacity assignment progress
// values may be added over time.
type CapacityAssignmentProgress string

const (
	// CapacityAssignmentProgressCompleted - Last capacity assignment operation was completed successfully.
	CapacityAssignmentProgressCompleted CapacityAssignmentProgress = "Completed"
	// CapacityAssignmentProgressFailed - Capacity assignment operation has encountered an error or failure and was unable to
	// complete.
	CapacityAssignmentProgressFailed CapacityAssignmentProgress = "Failed"
	// CapacityAssignmentProgressInProgress - Capacity assignment operation is currently running and has not yet completed.
	CapacityAssignmentProgressInProgress CapacityAssignmentProgress = "InProgress"
)

// PossibleCapacityAssignmentProgressValues returns the possible values for the CapacityAssignmentProgress const type.
func PossibleCapacityAssignmentProgressValues() []CapacityAssignmentProgress {
	return []CapacityAssignmentProgress{
		CapacityAssignmentProgressCompleted,
		CapacityAssignmentProgressFailed,
		CapacityAssignmentProgressInProgress,
	}
}

// CapacityRegion - The region of the capacity associated with this workspace. Additional capacity region values may be added
// over time.
type CapacityRegion string

const (
	// CapacityRegionAustraliaEast - Australia East region
	CapacityRegionAustraliaEast CapacityRegion = "Australia East"
	// CapacityRegionAustraliaSoutheast - Australia Southeast region
	CapacityRegionAustraliaSoutheast CapacityRegion = "Australia Southeast"
	// CapacityRegionBrazilSouth - Brazil South region
	CapacityRegionBrazilSouth CapacityRegion = "Brazil South"
	// CapacityRegionBrazilSoutheast - Brazil Southeast region
	CapacityRegionBrazilSoutheast CapacityRegion = "Brazil Southeast"
	// CapacityRegionCanadaCentral - Canada Central region
	CapacityRegionCanadaCentral CapacityRegion = "Canada Central"
	// CapacityRegionCanadaEast - Canada East region
	CapacityRegionCanadaEast CapacityRegion = "Canada East"
	// CapacityRegionCentralIndia - Central India region
	CapacityRegionCentralIndia CapacityRegion = "Central India"
	// CapacityRegionCentralUS - Central US region
	CapacityRegionCentralUS CapacityRegion = "Central US"
	// CapacityRegionCentralUSEUAP - Central US EUAP region
	CapacityRegionCentralUSEUAP CapacityRegion = "Central US EUAP"
	// CapacityRegionChinaEast - China East region
	CapacityRegionChinaEast CapacityRegion = "China East"
	// CapacityRegionChinaEast2 - China East 2 region
	CapacityRegionChinaEast2 CapacityRegion = "China East 2"
	// CapacityRegionChinaEast3 - China East 3 region
	CapacityRegionChinaEast3 CapacityRegion = "China East 3"
	// CapacityRegionChinaNorth - China North region
	CapacityRegionChinaNorth CapacityRegion = "China North"
	// CapacityRegionChinaNorth2 - China North 2 region
	CapacityRegionChinaNorth2 CapacityRegion = "China North 2"
	// CapacityRegionChinaNorth3 - China North 3 region
	CapacityRegionChinaNorth3 CapacityRegion = "China North 3"
	// CapacityRegionEastAsia - East Asia region
	CapacityRegionEastAsia CapacityRegion = "East Asia"
	// CapacityRegionEastUS - East US region
	CapacityRegionEastUS CapacityRegion = "East US"
	// CapacityRegionEastUS2 - East US 2 region
	CapacityRegionEastUS2 CapacityRegion = "East US 2"
	// CapacityRegionFranceCentral - France Central region
	CapacityRegionFranceCentral CapacityRegion = "France Central"
	// CapacityRegionFranceSouth - France South region
	CapacityRegionFranceSouth CapacityRegion = "France South"
	// CapacityRegionGermanyCentral - Germany Central region
	CapacityRegionGermanyCentral CapacityRegion = "Germany Central"
	// CapacityRegionGermanyNorth - Germany North region
	CapacityRegionGermanyNorth CapacityRegion = "Germany North"
	// CapacityRegionGermanyNortheast - Germany Northeast region
	CapacityRegionGermanyNortheast CapacityRegion = "Germany Northeast"
	// CapacityRegionGermanyWestCentral - Germany West Central region
	CapacityRegionGermanyWestCentral CapacityRegion = "Germany West Central"
	// CapacityRegionIsraelCentral - Israel Central region
	CapacityRegionIsraelCentral CapacityRegion = "Israel Central"
	// CapacityRegionItalyNorth - Italy North region
	CapacityRegionItalyNorth CapacityRegion = "Italy North"
	// CapacityRegionJapanEast - Japan East region
	CapacityRegionJapanEast CapacityRegion = "Japan East"
	// CapacityRegionJapanWest - Japan West region
	CapacityRegionJapanWest CapacityRegion = "Japan West"
	// CapacityRegionKoreaCentral - Korea Central region
	CapacityRegionKoreaCentral CapacityRegion = "Korea Central"
	// CapacityRegionKoreaSouth - Korea South region
	CapacityRegionKoreaSouth CapacityRegion = "Korea South"
	// CapacityRegionMexicoCentral - Mexico Central region
	CapacityRegionMexicoCentral CapacityRegion = "Mexico Central"
	// CapacityRegionNorthCentralUS - North Central US region
	CapacityRegionNorthCentralUS CapacityRegion = "North Central US"
	// CapacityRegionNorthEurope - North Europe region
	CapacityRegionNorthEurope CapacityRegion = "North Europe"
	// CapacityRegionNorwayEast - Norway East region
	CapacityRegionNorwayEast CapacityRegion = "Norway East"
	// CapacityRegionNorwayWest - Norway West region
	CapacityRegionNorwayWest CapacityRegion = "Norway West"
	// CapacityRegionPolandCentral - Poland Central region
	CapacityRegionPolandCentral CapacityRegion = "Poland Central"
	// CapacityRegionQatarCentral - Qatar Central region
	CapacityRegionQatarCentral CapacityRegion = "Qatar Central"
	// CapacityRegionSouthAfricaNorth - South Africa North region
	CapacityRegionSouthAfricaNorth CapacityRegion = "South Africa North"
	// CapacityRegionSouthAfricaWest - South Africa West region
	CapacityRegionSouthAfricaWest CapacityRegion = "South Africa West"
	// CapacityRegionSouthCentralUS - South Central US region
	CapacityRegionSouthCentralUS CapacityRegion = "South Central US"
	// CapacityRegionSouthIndia - South India region
	CapacityRegionSouthIndia CapacityRegion = "South India"
	// CapacityRegionSoutheastAsia - Southeast Asia region
	CapacityRegionSoutheastAsia CapacityRegion = "Southeast Asia"
	// CapacityRegionSpainCentral - Spain Central region
	CapacityRegionSpainCentral CapacityRegion = "Spain Central"
	// CapacityRegionSwedenCentral - Sweden Central region
	CapacityRegionSwedenCentral CapacityRegion = "Sweden Central"
	// CapacityRegionSwitzerlandNorth - Switzerland North region
	CapacityRegionSwitzerlandNorth CapacityRegion = "Switzerland North"
	// CapacityRegionSwitzerlandWest - Switzerland West region
	CapacityRegionSwitzerlandWest CapacityRegion = "Switzerland West"
	// CapacityRegionUAECentral - UAE Central region
	CapacityRegionUAECentral CapacityRegion = "UAE Central"
	// CapacityRegionUAENorth - UAE North region
	CapacityRegionUAENorth CapacityRegion = "UAE North"
	// CapacityRegionUKSouth - UK South region
	CapacityRegionUKSouth CapacityRegion = "UK South"
	// CapacityRegionUKWest - UK West region
	CapacityRegionUKWest CapacityRegion = "UK West"
	// CapacityRegionWestCentralUS - West Central US region
	CapacityRegionWestCentralUS CapacityRegion = "West Central US"
	// CapacityRegionWestEurope - West Europe region
	CapacityRegionWestEurope CapacityRegion = "West Europe"
	// CapacityRegionWestIndia - West India region
	CapacityRegionWestIndia CapacityRegion = "West India"
	// CapacityRegionWestUS - West US region
	CapacityRegionWestUS CapacityRegion = "West US"
	// CapacityRegionWestUS2 - West US 2 region
	CapacityRegionWestUS2 CapacityRegion = "West US 2"
	// CapacityRegionWestUS3 - West US 3 region
	CapacityRegionWestUS3 CapacityRegion = "West US 3"
)

// PossibleCapacityRegionValues returns the possible values for the CapacityRegion const type.
func PossibleCapacityRegionValues() []CapacityRegion {
	return []CapacityRegion{
		CapacityRegionAustraliaEast,
		CapacityRegionAustraliaSoutheast,
		CapacityRegionBrazilSouth,
		CapacityRegionBrazilSoutheast,
		CapacityRegionCanadaCentral,
		CapacityRegionCanadaEast,
		CapacityRegionCentralIndia,
		CapacityRegionCentralUS,
		CapacityRegionCentralUSEUAP,
		CapacityRegionChinaEast,
		CapacityRegionChinaEast2,
		CapacityRegionChinaEast3,
		CapacityRegionChinaNorth,
		CapacityRegionChinaNorth2,
		CapacityRegionChinaNorth3,
		CapacityRegionEastAsia,
		CapacityRegionEastUS,
		CapacityRegionEastUS2,
		CapacityRegionFranceCentral,
		CapacityRegionFranceSouth,
		CapacityRegionGermanyCentral,
		CapacityRegionGermanyNorth,
		CapacityRegionGermanyNortheast,
		CapacityRegionGermanyWestCentral,
		CapacityRegionIsraelCentral,
		CapacityRegionItalyNorth,
		CapacityRegionJapanEast,
		CapacityRegionJapanWest,
		CapacityRegionKoreaCentral,
		CapacityRegionKoreaSouth,
		CapacityRegionMexicoCentral,
		CapacityRegionNorthCentralUS,
		CapacityRegionNorthEurope,
		CapacityRegionNorwayEast,
		CapacityRegionNorwayWest,
		CapacityRegionPolandCentral,
		CapacityRegionQatarCentral,
		CapacityRegionSouthAfricaNorth,
		CapacityRegionSouthAfricaWest,
		CapacityRegionSouthCentralUS,
		CapacityRegionSouthIndia,
		CapacityRegionSoutheastAsia,
		CapacityRegionSpainCentral,
		CapacityRegionSwedenCentral,
		CapacityRegionSwitzerlandNorth,
		CapacityRegionSwitzerlandWest,
		CapacityRegionUAECentral,
		CapacityRegionUAENorth,
		CapacityRegionUKSouth,
		CapacityRegionUKWest,
		CapacityRegionWestCentralUS,
		CapacityRegionWestEurope,
		CapacityRegionWestIndia,
		CapacityRegionWestUS,
		CapacityRegionWestUS2,
		CapacityRegionWestUS3,
	}
}

// CapacityState - A Capacity state. Additional capacity states may be added over time.
type CapacityState string

const (
	// CapacityStateActive - The capacity is ready to use.
	CapacityStateActive CapacityState = "Active"
	// CapacityStateInactive - The capacity can't be used.
	CapacityStateInactive CapacityState = "Inactive"
)

// PossibleCapacityStateValues returns the possible values for the CapacityState const type.
func PossibleCapacityStateValues() []CapacityState {
	return []CapacityState{
		CapacityStateActive,
		CapacityStateInactive,
	}
}

// ChangeType - A change of an item. Additional changed types may be added over time.
type ChangeType string

const (
	// ChangeTypeAdded - A newly created item.
	ChangeTypeAdded ChangeType = "Added"
	// ChangeTypeDeleted - Item has been deleted.
	ChangeTypeDeleted ChangeType = "Deleted"
	// ChangeTypeModified - Item content has been modified.
	ChangeTypeModified ChangeType = "Modified"
)

// PossibleChangeTypeValues returns the possible values for the ChangeType const type.
func PossibleChangeTypeValues() []ChangeType {
	return []ChangeType{
		ChangeTypeAdded,
		ChangeTypeDeleted,
		ChangeTypeModified,
	}
}

// CommitMode - Modes for the commit operation. Additional modes may be added over time.
type CommitMode string

const (
	// CommitModeAll - Commit all uncommitted changes. The caller is not required to provide the list of items to commit.
	CommitModeAll CommitMode = "All"
	// CommitModeSelective - Commit a specified items list that has uncommitted changes.
	CommitModeSelective CommitMode = "Selective"
)

// PossibleCommitModeValues returns the possible values for the CommitMode const type.
func PossibleCommitModeValues() []CommitMode {
	return []CommitMode{
		CommitModeAll,
		CommitModeSelective,
	}
}

// ConflictResolutionPolicy - Conflict resolution policy. Additional conflict resolution policies may be added over time.
type ConflictResolutionPolicy string

const (
	// ConflictResolutionPolicyPreferRemote - Prefer remote Git side content.
	ConflictResolutionPolicyPreferRemote ConflictResolutionPolicy = "PreferRemote"
	// ConflictResolutionPolicyPreferWorkspace - Prefer workspace side content.
	ConflictResolutionPolicyPreferWorkspace ConflictResolutionPolicy = "PreferWorkspace"
)

// PossibleConflictResolutionPolicyValues returns the possible values for the ConflictResolutionPolicy const type.
func PossibleConflictResolutionPolicyValues() []ConflictResolutionPolicy {
	return []ConflictResolutionPolicy{
		ConflictResolutionPolicyPreferRemote,
		ConflictResolutionPolicyPreferWorkspace,
	}
}

// ConflictResolutionType - Conflict resolution type. Additional conflict resolution types may be added over time.
type ConflictResolutionType string

const (
	// ConflictResolutionTypeWorkspace - Conflict resolution representing the workspace level.
	ConflictResolutionTypeWorkspace ConflictResolutionType = "Workspace"
)

// PossibleConflictResolutionTypeValues returns the possible values for the ConflictResolutionType const type.
func PossibleConflictResolutionTypeValues() []ConflictResolutionType {
	return []ConflictResolutionType{
		ConflictResolutionTypeWorkspace,
	}
}

// ConflictType - A change of an item in both workspace and remote. Additional changed types may be added over time.
type ConflictType string

const (
	// ConflictTypeConflict - There are different changes to the item in the workspace and in remote Git.
	ConflictTypeConflict ConflictType = "Conflict"
	// ConflictTypeNone - There are no changes to the item.
	ConflictTypeNone ConflictType = "None"
	// ConflictTypeSameChanges - There are identical changes to the item in the workspace and in remote Git.
	ConflictTypeSameChanges ConflictType = "SameChanges"
)

// PossibleConflictTypeValues returns the possible values for the ConflictType const type.
func PossibleConflictTypeValues() []ConflictType {
	return []ConflictType{
		ConflictTypeConflict,
		ConflictTypeNone,
		ConflictTypeSameChanges,
	}
}

// ConnectivityType - The connectivity type of the connection. Additional connectivity types may be added over time.
type ConnectivityType string

const (
	// ConnectivityTypeAutomatic - The connection connects through the cloud using an implicit data connection. This option is
	// only available for specific scenarios like semantic models that use Single Sign-On (SSO).”
	ConnectivityTypeAutomatic ConnectivityType = "Automatic"
	// ConnectivityTypeNone - The connection is not bound
	ConnectivityTypeNone ConnectivityType = "None"
	// ConnectivityTypeOnPremisesGateway - The connection connects through an on-premises data gateway.
	ConnectivityTypeOnPremisesGateway ConnectivityType = "OnPremisesGateway"
	// ConnectivityTypeOnPremisesGatewayPersonal - The connection connects through a personal on-premises data gateway.
	ConnectivityTypeOnPremisesGatewayPersonal ConnectivityType = "OnPremisesGatewayPersonal"
	// ConnectivityTypePersonalCloud - The connection connects through the cloud and cannot be shared with others.
	ConnectivityTypePersonalCloud ConnectivityType = "PersonalCloud"
	// ConnectivityTypeShareableCloud - The connection connects through the cloud and can be shared with others.
	ConnectivityTypeShareableCloud ConnectivityType = "ShareableCloud"
	// ConnectivityTypeVirtualNetworkGateway - The connection connects through a virtual network data gateway.
	ConnectivityTypeVirtualNetworkGateway ConnectivityType = "VirtualNetworkGateway"
)

// PossibleConnectivityTypeValues returns the possible values for the ConnectivityType const type.
func PossibleConnectivityTypeValues() []ConnectivityType {
	return []ConnectivityType{
		ConnectivityTypeAutomatic,
		ConnectivityTypeNone,
		ConnectivityTypeOnPremisesGateway,
		ConnectivityTypeOnPremisesGatewayPersonal,
		ConnectivityTypePersonalCloud,
		ConnectivityTypeShareableCloud,
		ConnectivityTypeVirtualNetworkGateway,
	}
}

// DayOfWeek - Days of the week
type DayOfWeek string

const (
	// DayOfWeekFriday - Friday
	DayOfWeekFriday DayOfWeek = "Friday"
	// DayOfWeekMonday - Monday
	DayOfWeekMonday DayOfWeek = "Monday"
	// DayOfWeekSaturday - Saturday
	DayOfWeekSaturday DayOfWeek = "Saturday"
	// DayOfWeekSunday - Sunday
	DayOfWeekSunday DayOfWeek = "Sunday"
	// DayOfWeekThursday - Thursday
	DayOfWeekThursday DayOfWeek = "Thursday"
	// DayOfWeekTuesday - Tuesday
	DayOfWeekTuesday DayOfWeek = "Tuesday"
	// DayOfWeekWednesday - Wednesday
	DayOfWeekWednesday DayOfWeek = "Wednesday"
)

// PossibleDayOfWeekValues returns the possible values for the DayOfWeek const type.
func PossibleDayOfWeekValues() []DayOfWeek {
	return []DayOfWeek{
		DayOfWeekFriday,
		DayOfWeekMonday,
		DayOfWeekSaturday,
		DayOfWeekSunday,
		DayOfWeekThursday,
		DayOfWeekTuesday,
		DayOfWeekWednesday,
	}
}

// DeploymentPipelineOperationStatus - The status of the deployment pipeline operation. Additional statuses may be added over
// time.
type DeploymentPipelineOperationStatus string

const (
	// DeploymentPipelineOperationStatusFailed - The deployment pipeline operation failed.
	DeploymentPipelineOperationStatusFailed DeploymentPipelineOperationStatus = "Failed"
	// DeploymentPipelineOperationStatusNotStarted - The deployment pipeline operation didn't start.
	DeploymentPipelineOperationStatusNotStarted DeploymentPipelineOperationStatus = "NotStarted"
	// DeploymentPipelineOperationStatusRunning - The deployment pipeline operation is running.
	DeploymentPipelineOperationStatusRunning DeploymentPipelineOperationStatus = "Running"
	// DeploymentPipelineOperationStatusSucceeded - The deployment pipeline operation succeeded.
	DeploymentPipelineOperationStatusSucceeded DeploymentPipelineOperationStatus = "Succeeded"
)

// PossibleDeploymentPipelineOperationStatusValues returns the possible values for the DeploymentPipelineOperationStatus const type.
func PossibleDeploymentPipelineOperationStatusValues() []DeploymentPipelineOperationStatus {
	return []DeploymentPipelineOperationStatus{
		DeploymentPipelineOperationStatusFailed,
		DeploymentPipelineOperationStatusNotStarted,
		DeploymentPipelineOperationStatusRunning,
		DeploymentPipelineOperationStatusSucceeded,
	}
}

// DeploymentPipelineOperationType - The operation type. Additional types may be added over time.
type DeploymentPipelineOperationType string

const (
	// DeploymentPipelineOperationTypeDeploy - Deploy content between stages.
	DeploymentPipelineOperationTypeDeploy DeploymentPipelineOperationType = "Deploy"
)

// PossibleDeploymentPipelineOperationTypeValues returns the possible values for the DeploymentPipelineOperationType const type.
func PossibleDeploymentPipelineOperationTypeValues() []DeploymentPipelineOperationType {
	return []DeploymentPipelineOperationType{
		DeploymentPipelineOperationTypeDeploy,
	}
}

// Effect - The effect that a role has on access to the data resource. Currently, the only supported effect type is Permit,
// which grants access to the resource. Additional effect types may be added over time.
type Effect string

const (
	// EffectPermit - the effect type Permit
	EffectPermit Effect = "Permit"
)

// PossibleEffectValues returns the possible values for the Effect const type.
func PossibleEffectValues() []Effect {
	return []Effect{
		EffectPermit,
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

// GitConnectionState - Git connection state. Additional connection state types may be added over time.
type GitConnectionState string

const (
	// GitConnectionStateConnected - Connected state.
	GitConnectionStateConnected GitConnectionState = "Connected"
	// GitConnectionStateConnectedAndInitialized - Connected and initialized state.
	GitConnectionStateConnectedAndInitialized GitConnectionState = "ConnectedAndInitialized"
	// GitConnectionStateNotConnected - Not connected state.
	GitConnectionStateNotConnected GitConnectionState = "NotConnected"
)

// PossibleGitConnectionStateValues returns the possible values for the GitConnectionState const type.
func PossibleGitConnectionStateValues() []GitConnectionState {
	return []GitConnectionState{
		GitConnectionStateConnected,
		GitConnectionStateConnectedAndInitialized,
		GitConnectionStateNotConnected,
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

// InitializationStrategy - The strategy required for an initialization process when content exists on both the remote side
// and the workspace side. Additional strategies may be added over time.
type InitializationStrategy string

const (
	// InitializationStrategyNone - No strategy defined.
	InitializationStrategyNone InitializationStrategy = "None"
	// InitializationStrategyPreferRemote - Prefer remote Git side content.
	InitializationStrategyPreferRemote InitializationStrategy = "PreferRemote"
	// InitializationStrategyPreferWorkspace - Prefer workspace side content.
	InitializationStrategyPreferWorkspace InitializationStrategy = "PreferWorkspace"
)

// PossibleInitializationStrategyValues returns the possible values for the InitializationStrategy const type.
func PossibleInitializationStrategyValues() []InitializationStrategy {
	return []InitializationStrategy{
		InitializationStrategyNone,
		InitializationStrategyPreferRemote,
		InitializationStrategyPreferWorkspace,
	}
}

// InvokeType - The item job invoke type. Additional invokeTypes may be added over time.
type InvokeType string

const (
	// InvokeTypeManual - Job is invoked manually
	InvokeTypeManual InvokeType = "Manual"
	// InvokeTypeScheduled - Job is scheduled
	InvokeTypeScheduled InvokeType = "Scheduled"
)

// PossibleInvokeTypeValues returns the possible values for the InvokeType const type.
func PossibleInvokeTypeValues() []InvokeType {
	return []InvokeType{
		InvokeTypeManual,
		InvokeTypeScheduled,
	}
}

type ItemAccess string

const (
	// ItemAccessExecute - Item Access Execute.
	ItemAccessExecute ItemAccess = "Execute"
	// ItemAccessExplore - Item Access Explore.
	ItemAccessExplore ItemAccess = "Explore"
	// ItemAccessRead - Item Access Read.
	ItemAccessRead ItemAccess = "Read"
	// ItemAccessReadAll - Item Access ReadAll.
	ItemAccessReadAll ItemAccess = "ReadAll"
	// ItemAccessReshare - Item Access Reshare.
	ItemAccessReshare ItemAccess = "Reshare"
	// ItemAccessWrite - Item Access Write.
	ItemAccessWrite ItemAccess = "Write"
)

// PossibleItemAccessValues returns the possible values for the ItemAccess const type.
func PossibleItemAccessValues() []ItemAccess {
	return []ItemAccess{
		ItemAccessExecute,
		ItemAccessExplore,
		ItemAccessRead,
		ItemAccessReadAll,
		ItemAccessReshare,
		ItemAccessWrite,
	}
}

// ItemPreDeploymentDiffState - Specifies if an item is new, different or identical to items in the target stage before deployment.
// Additional states may be added over time.
type ItemPreDeploymentDiffState string

const (
	// ItemPreDeploymentDiffStateDifferent - Before deployment, the item in the source stage wasn't identical to the one in the
	// target stage.
	ItemPreDeploymentDiffStateDifferent ItemPreDeploymentDiffState = "Different"
	// ItemPreDeploymentDiffStateNew - A new deployed item that doesn't exist in the target stage.
	ItemPreDeploymentDiffStateNew ItemPreDeploymentDiffState = "New"
	// ItemPreDeploymentDiffStateNoDifference - Before deployment, the item in the source stage was identical to the one in the
	// target stage.
	ItemPreDeploymentDiffStateNoDifference ItemPreDeploymentDiffState = "NoDifference"
)

// PossibleItemPreDeploymentDiffStateValues returns the possible values for the ItemPreDeploymentDiffState const type.
func PossibleItemPreDeploymentDiffStateValues() []ItemPreDeploymentDiffState {
	return []ItemPreDeploymentDiffState{
		ItemPreDeploymentDiffStateDifferent,
		ItemPreDeploymentDiffStateNew,
		ItemPreDeploymentDiffStateNoDifference,
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

// LongRunningOperationStatus - The current status of the operation. Additional operation statuses may be added over time.
type LongRunningOperationStatus string

const (
	// LongRunningOperationStatusFailed - The operation has failed
	LongRunningOperationStatusFailed LongRunningOperationStatus = "Failed"
	// LongRunningOperationStatusNotStarted - The operation didn't start
	LongRunningOperationStatusNotStarted LongRunningOperationStatus = "NotStarted"
	// LongRunningOperationStatusRunning - The operation is running
	LongRunningOperationStatusRunning LongRunningOperationStatus = "Running"
	// LongRunningOperationStatusSucceeded - The operation has finished successfully
	LongRunningOperationStatusSucceeded LongRunningOperationStatus = "Succeeded"
	// LongRunningOperationStatusUndefined - The status of the operation is undefined
	LongRunningOperationStatusUndefined LongRunningOperationStatus = "Undefined"
)

// PossibleLongRunningOperationStatusValues returns the possible values for the LongRunningOperationStatus const type.
func PossibleLongRunningOperationStatusValues() []LongRunningOperationStatus {
	return []LongRunningOperationStatus{
		LongRunningOperationStatusFailed,
		LongRunningOperationStatusNotStarted,
		LongRunningOperationStatusRunning,
		LongRunningOperationStatusSucceeded,
		LongRunningOperationStatusUndefined,
	}
}

// ObjectType - The type of Microsoft Entra ID object. Additional objectType types may be added over time.
type ObjectType string

const (
	// ObjectTypeGroup - Attribute name Group
	ObjectTypeGroup ObjectType = "Group"
	// ObjectTypeManagedIdentity - Attribute name ManagedIdentity
	ObjectTypeManagedIdentity ObjectType = "ManagedIdentity"
	// ObjectTypeServicePrincipal - Attribute name ServicePrincipal
	ObjectTypeServicePrincipal ObjectType = "ServicePrincipal"
	// ObjectTypeUser - Attribute name User
	ObjectTypeUser ObjectType = "User"
)

// PossibleObjectTypeValues returns the possible values for the ObjectType const type.
func PossibleObjectTypeValues() []ObjectType {
	return []ObjectType{
		ObjectTypeGroup,
		ObjectTypeManagedIdentity,
		ObjectTypeServicePrincipal,
		ObjectTypeUser,
	}
}

// PayloadType - The type of the definition part payload. Additional payload types may be added over time.
type PayloadType string

const (
	// PayloadTypeInlineBase64 - Inline Base 64.
	PayloadTypeInlineBase64 PayloadType = "InlineBase64"
)

// PossiblePayloadTypeValues returns the possible values for the PayloadType const type.
func PossiblePayloadTypeValues() []PayloadType {
	return []PayloadType{
		PayloadTypeInlineBase64,
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

// RequiredAction - Required action after the initialization process has finished. Additional actions may be added over time.
type RequiredAction string

const (
	// RequiredActionCommitToGit - Commit to Git is required.
	RequiredActionCommitToGit RequiredAction = "CommitToGit"
	// RequiredActionNone - No action is required.
	RequiredActionNone RequiredAction = "None"
	// RequiredActionUpdateFromGit - Update from Git is required.
	RequiredActionUpdateFromGit RequiredAction = "UpdateFromGit"
)

// PossibleRequiredActionValues returns the possible values for the RequiredAction const type.
func PossibleRequiredActionValues() []RequiredAction {
	return []RequiredAction{
		RequiredActionCommitToGit,
		RequiredActionNone,
		RequiredActionUpdateFromGit,
	}
}

// ScheduleType - A string represents the type of the plan. Additional planType types may be added over time.
type ScheduleType string

const (
	// ScheduleTypeCron - A type of schedule triggers jobs periodically.
	ScheduleTypeCron ScheduleType = "Cron"
	// ScheduleTypeDaily - A type of schedule triggers jobs daily.
	ScheduleTypeDaily ScheduleType = "Daily"
	// ScheduleTypeWeekly - A type of schedule triggers jobs by the week.
	ScheduleTypeWeekly ScheduleType = "Weekly"
)

// PossibleScheduleTypeValues returns the possible values for the ScheduleType const type.
func PossibleScheduleTypeValues() []ScheduleType {
	return []ScheduleType{
		ScheduleTypeCron,
		ScheduleTypeDaily,
		ScheduleTypeWeekly,
	}
}

type ShortcutConflictPolicy string

const (
	// ShortcutConflictPolicyAbort - When a shortcut with the same name and path already exists the shortcut creation will be
	// cancelled.
	ShortcutConflictPolicyAbort ShortcutConflictPolicy = "Abort"
	// ShortcutConflictPolicyGenerateUniqueName - When a shortcut with the same name and path already exists the shortcut creation
	// will continue with a new unique shortcut name.
	ShortcutConflictPolicyGenerateUniqueName ShortcutConflictPolicy = "GenerateUniqueName"
)

// PossibleShortcutConflictPolicyValues returns the possible values for the ShortcutConflictPolicy const type.
func PossibleShortcutConflictPolicyValues() []ShortcutConflictPolicy {
	return []ShortcutConflictPolicy{
		ShortcutConflictPolicyAbort,
		ShortcutConflictPolicyGenerateUniqueName,
	}
}

// Status - The item job status. Additional statuses may be added over time.
type Status string

const (
	// StatusCancelled - Job cancelled
	StatusCancelled Status = "Cancelled"
	// StatusCompleted - Job completed
	StatusCompleted Status = "Completed"
	// StatusDeduped - A job instance of the same job type is already running and this job instance is skipped
	StatusDeduped Status = "Deduped"
	// StatusFailed - Job failed
	StatusFailed Status = "Failed"
	// StatusInProgress - Job in progress
	StatusInProgress Status = "InProgress"
	// StatusNotStarted - Job not started
	StatusNotStarted Status = "NotStarted"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusCancelled,
		StatusCompleted,
		StatusDeduped,
		StatusFailed,
		StatusInProgress,
		StatusNotStarted,
	}
}

// Type - The type object contains properties like target shortcut account type. Additional types may be added over time.
type Type string

const (
	// TypeAdlsGen2 - AdlsGen2
	TypeAdlsGen2 Type = "AdlsGen2"
	// TypeAmazonS3 - AmazonS3
	TypeAmazonS3 Type = "AmazonS3"
	// TypeDataverse - Dataverse
	TypeDataverse Type = "Dataverse"
	// TypeExternalDataShare - ExternalDataShare
	TypeExternalDataShare Type = "ExternalDataShare"
	// TypeGoogleCloudStorage - GoogleCloudStorage
	TypeGoogleCloudStorage Type = "GoogleCloudStorage"
	// TypeOneLake - OneLake
	TypeOneLake Type = "OneLake"
	// TypeS3Compatible - S3Compatible
	TypeS3Compatible Type = "S3Compatible"
)

// PossibleTypeValues returns the possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{
		TypeAdlsGen2,
		TypeAmazonS3,
		TypeDataverse,
		TypeExternalDataShare,
		TypeGoogleCloudStorage,
		TypeOneLake,
		TypeS3Compatible,
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
