// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package environment

// CustomPoolType - Custom pool type. Additional CustomPoolType types may be added over time.
type CustomPoolType string

const (
	// CustomPoolTypeCapacity - Capacity custom pool
	CustomPoolTypeCapacity CustomPoolType = "Capacity"
	// CustomPoolTypeWorkspace - Workspace custom pool
	CustomPoolTypeWorkspace CustomPoolType = "Workspace"
)

// PossibleCustomPoolTypeValues returns the possible values for the CustomPoolType const type.
func PossibleCustomPoolTypeValues() []CustomPoolType {
	return []CustomPoolType{
		CustomPoolTypeCapacity,
		CustomPoolTypeWorkspace,
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

// PublishState - Publish state. Additional state types may be added over time.
type PublishState string

const (
	// PublishStateCancelled - Environment publish is in cancelled state.
	PublishStateCancelled PublishState = "Cancelled"
	// PublishStateCancelling - Environment publish is in cancelling state.
	PublishStateCancelling PublishState = "Cancelling"
	// PublishStateFailed - Environment publish is in failed state.
	PublishStateFailed PublishState = "Failed"
	// PublishStateRunning - Environment publish is in running state.
	PublishStateRunning PublishState = "Running"
	// PublishStateSuccess - Environment publish is in success state.
	PublishStateSuccess PublishState = "Success"
	// PublishStateWaiting - Environment publish is in waiting state.
	PublishStateWaiting PublishState = "Waiting"
)

// PossibleEnvironmentPublishStateValues returns the possible values for the PublishState const type.
func PossibleEnvironmentPublishStateValues() []PublishState {
	return []PublishState{
		PublishStateCancelled,
		PublishStateCancelling,
		PublishStateFailed,
		PublishStateRunning,
		PublishStateSuccess,
		PublishStateWaiting,
	}
}
