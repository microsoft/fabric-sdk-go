// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package core

import "io"

// CapacitiesClientListCapacitiesResponse contains the response from method CapacitiesClient.NewListCapacitiesPager.
type CapacitiesClientListCapacitiesResponse struct {
	Capacities
}

// DeploymentPipelinesClientDeployStageContentResponse contains the response from method DeploymentPipelinesClient.BeginDeployStageContent.
type DeploymentPipelinesClientDeployStageContentResponse struct {
	// A Fabric deployment pipeline operation.
	DeploymentPipelineOperation
}

// DeploymentPipelinesClientGetDeploymentPipelineResponse contains the response from method DeploymentPipelinesClient.GetDeploymentPipeline.
type DeploymentPipelinesClientGetDeploymentPipelineResponse struct {
	// A Fabric deployment pipeline.
	DeploymentPipeline
}

// DeploymentPipelinesClientListDeploymentPipelineStageItemsResponse contains the response from method DeploymentPipelinesClient.NewListDeploymentPipelineStageItemsPager.
type DeploymentPipelinesClientListDeploymentPipelineStageItemsResponse struct {
	// Supported items from a workspace that's assigned to a deployment pipeline stage.
	DeploymentPipelineStageItems
}

// DeploymentPipelinesClientListDeploymentPipelineStagesResponse contains the response from method DeploymentPipelinesClient.NewListDeploymentPipelineStagesPager.
type DeploymentPipelinesClientListDeploymentPipelineStagesResponse struct {
	// A collection of Fabric deployment pipeline stages.
	DeploymentPipelineStages
}

// DeploymentPipelinesClientListDeploymentPipelinesResponse contains the response from method DeploymentPipelinesClient.NewListDeploymentPipelinesPager.
type DeploymentPipelinesClientListDeploymentPipelinesResponse struct {
	// A collection of Fabric deployment pipelines.
	DeploymentPipelines
}

// ExternalDataSharesClientCreateExternalDataShareResponse contains the response from method ExternalDataSharesClient.CreateExternalDataShare.
type ExternalDataSharesClientCreateExternalDataShareResponse struct {
	// An external data share object.
	ExternalDataShare

	// Location contains the information returned from the Location header response.
	Location *string
}

// ExternalDataSharesClientGetExternalDataShareResponse contains the response from method ExternalDataSharesClient.GetExternalDataShare.
type ExternalDataSharesClientGetExternalDataShareResponse struct {
	// An external data share object.
	ExternalDataShare
}

// ExternalDataSharesClientListExternalDataSharesInItemResponse contains the response from method ExternalDataSharesClient.NewListExternalDataSharesInItemPager.
type ExternalDataSharesClientListExternalDataSharesInItemResponse struct {
	// A list of external data shares with a continuation token.
	ExternalDataShares
}

// ExternalDataSharesClientRevokeExternalDataShareResponse contains the response from method ExternalDataSharesClient.RevokeExternalDataShare.
type ExternalDataSharesClientRevokeExternalDataShareResponse struct {
	// placeholder for future response values
}

// GitClientCommitToGitResponse contains the response from method GitClient.BeginCommitToGit.
type GitClientCommitToGitResponse struct {
	// placeholder for future response values
}

// GitClientConnectResponse contains the response from method GitClient.Connect.
type GitClientConnectResponse struct {
	// placeholder for future response values
}

// GitClientDisconnectResponse contains the response from method GitClient.Disconnect.
type GitClientDisconnectResponse struct {
	// placeholder for future response values
}

// GitClientGetConnectionResponse contains the response from method GitClient.GetConnection.
type GitClientGetConnectionResponse struct {
	// Contains the Git connection details.
	GitConnection
}

// GitClientGetMyGitCredentialsResponse contains the response from method GitClient.GetMyGitCredentials.
type GitClientGetMyGitCredentialsResponse struct {
	// The Git credentials configuration.
	GitCredentialsConfigurationResponseClassification
}

// GitClientGetStatusResponse contains the response from method GitClient.BeginGetStatus.
type GitClientGetStatusResponse struct {
	// Contains the status response.
	GitStatusResponse
}

// GitClientInitializeConnectionResponse contains the response from method GitClient.BeginInitializeConnection.
type GitClientInitializeConnectionResponse struct {
	// Contains the initialize Git connection response data.
	InitializeGitConnectionResponse
}

// GitClientUpdateFromGitResponse contains the response from method GitClient.BeginUpdateFromGit.
type GitClientUpdateFromGitResponse struct {
	// placeholder for future response values
}

// GitClientUpdateMyGitCredentialsResponse contains the response from method GitClient.UpdateMyGitCredentials.
type GitClientUpdateMyGitCredentialsResponse struct {
	// The Git credentials configuration.
	GitCredentialsConfigurationResponseClassification
}

// ItemsClientCreateItemResponse contains the response from method ItemsClient.BeginCreateItem.
type ItemsClientCreateItemResponse struct {
	// An item object.
	Item
}

// ItemsClientDeleteItemResponse contains the response from method ItemsClient.DeleteItem.
type ItemsClientDeleteItemResponse struct {
	// placeholder for future response values
}

// ItemsClientGetItemDefinitionResponse contains the response from method ItemsClient.BeginGetItemDefinition.
type ItemsClientGetItemDefinitionResponse struct {
	// Item public definition response.
	ItemDefinitionResponse
}

// ItemsClientGetItemResponse contains the response from method ItemsClient.GetItem.
type ItemsClientGetItemResponse struct {
	// An item object.
	Item
}

// ItemsClientListItemConnectionsResponse contains the response from method ItemsClient.NewListItemConnectionsPager.
type ItemsClientListItemConnectionsResponse struct {
	ItemConnections
}

// ItemsClientListItemsResponse contains the response from method ItemsClient.NewListItemsPager.
type ItemsClientListItemsResponse struct {
	Items
}

// ItemsClientUpdateItemDefinitionResponse contains the response from method ItemsClient.BeginUpdateItemDefinition.
type ItemsClientUpdateItemDefinitionResponse struct {
	// placeholder for future response values
}

// ItemsClientUpdateItemResponse contains the response from method ItemsClient.UpdateItem.
type ItemsClientUpdateItemResponse struct {
	// An item object.
	Item
}

// JobSchedulerClientCancelItemJobInstanceResponse contains the response from method JobSchedulerClient.CancelItemJobInstance.
type JobSchedulerClientCancelItemJobInstanceResponse struct {
	// Location contains the information returned from the Location header response.
	Location *string

	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *int32
}

// JobSchedulerClientCreateItemScheduleResponse contains the response from method JobSchedulerClient.CreateItemSchedule.
type JobSchedulerClientCreateItemScheduleResponse struct {
	// Item schedule.
	ItemSchedule

	// Location contains the information returned from the Location header response.
	Location *string
}

// JobSchedulerClientGetItemJobInstanceResponse contains the response from method JobSchedulerClient.GetItemJobInstance.
type JobSchedulerClientGetItemJobInstanceResponse struct {
	// An object representing item job instance
	ItemJobInstance
}

// JobSchedulerClientGetItemScheduleResponse contains the response from method JobSchedulerClient.GetItemSchedule.
type JobSchedulerClientGetItemScheduleResponse struct {
	// Item schedule.
	ItemSchedule
}

// JobSchedulerClientListItemJobInstancesResponse contains the response from method JobSchedulerClient.NewListItemJobInstancesPager.
type JobSchedulerClientListItemJobInstancesResponse struct {
	ItemJobInstances
}

// JobSchedulerClientListItemSchedulesResponse contains the response from method JobSchedulerClient.ListItemSchedules.
type JobSchedulerClientListItemSchedulesResponse struct {
	// list of schedules for this item.
	ItemSchedules
}

// JobSchedulerClientRunOnDemandItemJobResponse contains the response from method JobSchedulerClient.RunOnDemandItemJob.
type JobSchedulerClientRunOnDemandItemJobResponse struct {
	// Location contains the information returned from the Location header response.
	Location *string

	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *int32
}

// JobSchedulerClientUpdateItemScheduleResponse contains the response from method JobSchedulerClient.UpdateItemSchedule.
type JobSchedulerClientUpdateItemScheduleResponse struct {
	// Item schedule.
	ItemSchedule
}

// LongRunningOperationsClientGetOperationResultResponse contains the response from method LongRunningOperationsClient.GetOperationResult.
type LongRunningOperationsClientGetOperationResultResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser
}

// LongRunningOperationsClientGetOperationStateResponse contains the response from method LongRunningOperationsClient.GetOperationState.
type LongRunningOperationsClientGetOperationStateResponse struct {
	// An object describing the details and current state of a long running operation
	OperationState

	// Location contains the information returned from the Location header response.
	Location *string

	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *int32

	// XMSOperationID contains the information returned from the x-ms-operation-id header response.
	XMSOperationID *string
}

// ManagedPrivateEndpointsClientCreateWorkspaceManagedPrivateEndpointResponse contains the response from method ManagedPrivateEndpointsClient.CreateWorkspaceManagedPrivateEndpoint.
type ManagedPrivateEndpointsClientCreateWorkspaceManagedPrivateEndpointResponse struct {
	// Managed private endpoint.
	ManagedPrivateEndpoint

	// Location contains the information returned from the Location header response.
	Location *string
}

// ManagedPrivateEndpointsClientDeleteWorkspaceManagedPrivateEndpointResponse contains the response from method ManagedPrivateEndpointsClient.DeleteWorkspaceManagedPrivateEndpoint.
type ManagedPrivateEndpointsClientDeleteWorkspaceManagedPrivateEndpointResponse struct {
	// placeholder for future response values
}

// ManagedPrivateEndpointsClientGetWorkspaceManagedPrivateEndpointResponse contains the response from method ManagedPrivateEndpointsClient.GetWorkspaceManagedPrivateEndpoint.
type ManagedPrivateEndpointsClientGetWorkspaceManagedPrivateEndpointResponse struct {
	// Managed private endpoint.
	ManagedPrivateEndpoint
}

// ManagedPrivateEndpointsClientListWorkspaceManagedPrivateEndpointsResponse contains the response from method ManagedPrivateEndpointsClient.NewListWorkspaceManagedPrivateEndpointsPager.
type ManagedPrivateEndpointsClientListWorkspaceManagedPrivateEndpointsResponse struct {
	ManagedPrivateEndpoints
}

// OneLakeDataAccessSecurityClientCreateOrUpdateDataAccessRolesResponse contains the response from method OneLakeDataAccessSecurityClient.CreateOrUpdateDataAccessRoles.
type OneLakeDataAccessSecurityClientCreateOrUpdateDataAccessRolesResponse struct {
	// Etag contains the information returned from the Etag header response.
	Etag *string
}

// OneLakeDataAccessSecurityClientListDataAccessRolesResponse contains the response from method OneLakeDataAccessSecurityClient.ListDataAccessRoles.
type OneLakeDataAccessSecurityClientListDataAccessRolesResponse struct {
	DataAccessRoles

	// Etag contains the information returned from the Etag header response.
	Etag *string
}

// OneLakeShortcutsClientCreateShortcutResponse contains the response from method OneLakeShortcutsClient.CreateShortcut.
type OneLakeShortcutsClientCreateShortcutResponse struct {
	// An object representing a reference that points to other storage locations which can be internal or external to OneLake.
	// Shortcut is defined by name, path where the shortcut is created and target specifying the target storage location.
	Shortcut

	// Location contains the information returned from the Location header response.
	Location *string
}

// OneLakeShortcutsClientDeleteShortcutResponse contains the response from method OneLakeShortcutsClient.DeleteShortcut.
type OneLakeShortcutsClientDeleteShortcutResponse struct {
	// placeholder for future response values
}

// OneLakeShortcutsClientGetShortcutResponse contains the response from method OneLakeShortcutsClient.GetShortcut.
type OneLakeShortcutsClientGetShortcutResponse struct {
	// An object representing a reference that points to other storage locations which can be internal or external to OneLake.
	// Shortcut is defined by name, path where the shortcut is created and target specifying the target storage location.
	Shortcut
}

// OneLakeShortcutsClientListShortcutsResponse contains the response from method OneLakeShortcutsClient.NewListShortcutsPager.
type OneLakeShortcutsClientListShortcutsResponse struct {
	Shortcuts
}

// WorkspacesClientAddWorkspaceRoleAssignmentResponse contains the response from method WorkspacesClient.AddWorkspaceRoleAssignment.
type WorkspacesClientAddWorkspaceRoleAssignmentResponse struct {
	// A workspace role assignment object.
	WorkspaceRoleAssignment

	// Location contains the information returned from the Location header response.
	Location *string
}

// WorkspacesClientAssignToCapacityResponse contains the response from method WorkspacesClient.AssignToCapacity.
type WorkspacesClientAssignToCapacityResponse struct {
	// placeholder for future response values
}

// WorkspacesClientCreateWorkspaceResponse contains the response from method WorkspacesClient.CreateWorkspace.
type WorkspacesClientCreateWorkspaceResponse struct {
	// A workspace object.
	Workspace

	// Location contains the information returned from the Location header response.
	Location *string
}

// WorkspacesClientDeleteWorkspaceResponse contains the response from method WorkspacesClient.DeleteWorkspace.
type WorkspacesClientDeleteWorkspaceResponse struct {
	// placeholder for future response values
}

// WorkspacesClientDeleteWorkspaceRoleAssignmentResponse contains the response from method WorkspacesClient.DeleteWorkspaceRoleAssignment.
type WorkspacesClientDeleteWorkspaceRoleAssignmentResponse struct {
	// placeholder for future response values
}

// WorkspacesClientDeprovisionIdentityResponse contains the response from method WorkspacesClient.BeginDeprovisionIdentity.
type WorkspacesClientDeprovisionIdentityResponse struct {
	// placeholder for future response values
}

// WorkspacesClientGetWorkspaceResponse contains the response from method WorkspacesClient.GetWorkspace.
type WorkspacesClientGetWorkspaceResponse struct {
	// A workspace object.
	WorkspaceInfo
}

// WorkspacesClientGetWorkspaceRoleAssignmentResponse contains the response from method WorkspacesClient.GetWorkspaceRoleAssignment.
type WorkspacesClientGetWorkspaceRoleAssignmentResponse struct {
	// A workspace role assignment object.
	WorkspaceRoleAssignment
}

// WorkspacesClientListWorkspaceRoleAssignmentsResponse contains the response from method WorkspacesClient.NewListWorkspaceRoleAssignmentsPager.
type WorkspacesClientListWorkspaceRoleAssignmentsResponse struct {
	WorkspaceRoleAssignments
}

// WorkspacesClientListWorkspacesResponse contains the response from method WorkspacesClient.NewListWorkspacesPager.
type WorkspacesClientListWorkspacesResponse struct {
	Workspaces
}

// WorkspacesClientProvisionIdentityResponse contains the response from method WorkspacesClient.BeginProvisionIdentity.
type WorkspacesClientProvisionIdentityResponse struct {
	// A workspace identity object.
	WorkspaceIdentity
}

// WorkspacesClientUnassignFromCapacityResponse contains the response from method WorkspacesClient.UnassignFromCapacity.
type WorkspacesClientUnassignFromCapacityResponse struct {
	// placeholder for future response values
}

// WorkspacesClientUpdateWorkspaceResponse contains the response from method WorkspacesClient.UpdateWorkspace.
type WorkspacesClientUpdateWorkspaceResponse struct {
	// A workspace object.
	Workspace
}

// WorkspacesClientUpdateWorkspaceRoleAssignmentResponse contains the response from method WorkspacesClient.UpdateWorkspaceRoleAssignment.
type WorkspacesClientUpdateWorkspaceRoleAssignmentResponse struct {
	// A workspace role assignment object.
	WorkspaceRoleAssignment
}
