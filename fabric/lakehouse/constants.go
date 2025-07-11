// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package lakehouse

// FileFormat - Data file format name. Additional file format types may be added over time.
type FileFormat string

const (
	// FileFormatCSV - CSV format name.
	FileFormatCSV FileFormat = "Csv"
	// FileFormatParquet - Parquet format name.
	FileFormatParquet FileFormat = "Parquet"
)

// PossibleFileFormatValues returns the possible values for the FileFormat const type.
func PossibleFileFormatValues() []FileFormat {
	return []FileFormat{
		FileFormatCSV,
		FileFormatParquet,
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

// ItemReferenceType - The Item reference type. Additional ItemReferenceType types may be added over time.
type ItemReferenceType string

const (
	// ItemReferenceTypeByID - The item is referenced by its ID.
	ItemReferenceTypeByID ItemReferenceType = "ById"
)

// PossibleItemReferenceTypeValues returns the possible values for the ItemReferenceType const type.
func PossibleItemReferenceTypeValues() []ItemReferenceType {
	return []ItemReferenceType{
		ItemReferenceTypeByID,
	}
}

// ItemType - The type of the item. Additional item types may be added over time.
type ItemType string

const (
	// ItemTypeApacheAirflowJob - An ApacheAirflowJob.
	ItemTypeApacheAirflowJob ItemType = "ApacheAirflowJob"
	// ItemTypeCopyJob - A Copy job.
	ItemTypeCopyJob ItemType = "CopyJob"
	// ItemTypeDashboard - PowerBI dashboard.
	ItemTypeDashboard ItemType = "Dashboard"
	// ItemTypeDataPipeline - A data pipeline.
	ItemTypeDataPipeline ItemType = "DataPipeline"
	// ItemTypeDataflow - A Dataflow.
	ItemTypeDataflow ItemType = "Dataflow"
	// ItemTypeDatamart - PowerBI datamart.
	ItemTypeDatamart ItemType = "Datamart"
	// ItemTypeDigitalTwinBuilder - A DigitalTwinBuilder.
	ItemTypeDigitalTwinBuilder ItemType = "DigitalTwinBuilder"
	// ItemTypeDigitalTwinBuilderFlow - A Digital Twin Builder Flow.
	ItemTypeDigitalTwinBuilderFlow ItemType = "DigitalTwinBuilderFlow"
	// ItemTypeEnvironment - An environment.
	ItemTypeEnvironment ItemType = "Environment"
	// ItemTypeEventhouse - An eventhouse.
	ItemTypeEventhouse ItemType = "Eventhouse"
	// ItemTypeEventstream - An eventstream.
	ItemTypeEventstream ItemType = "Eventstream"
	// ItemTypeGraphQLAPI - An API for GraphQL item.
	ItemTypeGraphQLAPI ItemType = "GraphQLApi"
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
	// ItemTypeMirroredAzureDatabricksCatalog - A mirrored azure databricks catalog.
	ItemTypeMirroredAzureDatabricksCatalog ItemType = "MirroredAzureDatabricksCatalog"
	// ItemTypeMirroredDatabase - A mirrored database.
	ItemTypeMirroredDatabase ItemType = "MirroredDatabase"
	// ItemTypeMirroredWarehouse - A mirrored warehouse.
	ItemTypeMirroredWarehouse ItemType = "MirroredWarehouse"
	// ItemTypeMountedDataFactory - A MountedDataFactory.
	ItemTypeMountedDataFactory ItemType = "MountedDataFactory"
	// ItemTypeNotebook - A notebook.
	ItemTypeNotebook ItemType = "Notebook"
	// ItemTypePaginatedReport - PowerBI paginated report.
	ItemTypePaginatedReport ItemType = "PaginatedReport"
	// ItemTypeReflex - A Reflex.
	ItemTypeReflex ItemType = "Reflex"
	// ItemTypeReport - PowerBI report.
	ItemTypeReport ItemType = "Report"
	// ItemTypeSQLDatabase - A SQLDatabase.
	ItemTypeSQLDatabase ItemType = "SQLDatabase"
	// ItemTypeSQLEndpoint - An SQL endpoint.
	ItemTypeSQLEndpoint ItemType = "SQLEndpoint"
	// ItemTypeSemanticModel - PowerBI semantic model.
	ItemTypeSemanticModel ItemType = "SemanticModel"
	// ItemTypeSparkJobDefinition - A spark job definition.
	ItemTypeSparkJobDefinition ItemType = "SparkJobDefinition"
	// ItemTypeVariableLibrary - A VariableLibrary.
	ItemTypeVariableLibrary ItemType = "VariableLibrary"
	// ItemTypeWarehouse - A warehouse.
	ItemTypeWarehouse ItemType = "Warehouse"
	// ItemTypeWarehouseSnapshot - A Warehouse snapshot.
	ItemTypeWarehouseSnapshot ItemType = "WarehouseSnapshot"
)

// PossibleItemTypeValues returns the possible values for the ItemType const type.
func PossibleItemTypeValues() []ItemType {
	return []ItemType{
		ItemTypeApacheAirflowJob,
		ItemTypeCopyJob,
		ItemTypeDashboard,
		ItemTypeDataPipeline,
		ItemTypeDataflow,
		ItemTypeDatamart,
		ItemTypeDigitalTwinBuilder,
		ItemTypeDigitalTwinBuilderFlow,
		ItemTypeEnvironment,
		ItemTypeEventhouse,
		ItemTypeEventstream,
		ItemTypeGraphQLAPI,
		ItemTypeKQLDashboard,
		ItemTypeKQLDatabase,
		ItemTypeKQLQueryset,
		ItemTypeLakehouse,
		ItemTypeMLExperiment,
		ItemTypeMLModel,
		ItemTypeMirroredAzureDatabricksCatalog,
		ItemTypeMirroredDatabase,
		ItemTypeMirroredWarehouse,
		ItemTypeMountedDataFactory,
		ItemTypeNotebook,
		ItemTypePaginatedReport,
		ItemTypeReflex,
		ItemTypeReport,
		ItemTypeSQLDatabase,
		ItemTypeSQLEndpoint,
		ItemTypeSemanticModel,
		ItemTypeSparkJobDefinition,
		ItemTypeVariableLibrary,
		ItemTypeWarehouse,
		ItemTypeWarehouseSnapshot,
	}
}

// JobType - Type of the job. Values are: Unknown, SparkSession, SparkBatch, JupyterSession. Additional LivySessionJobType
// types may be added over time.
type JobType string

const (
	// JobTypeJupyterSession - Job type is from jupyter session.
	JobTypeJupyterSession JobType = "JupyterSession"
	// JobTypeSparkBatch - Job type is from a spark batch.
	JobTypeSparkBatch JobType = "SparkBatch"
	// JobTypeSparkSession - Job type is from a spark session.
	JobTypeSparkSession JobType = "SparkSession"
	// JobTypeUnknown - Job type is unknown.
	JobTypeUnknown JobType = "Unknown"
)

// PossibleJobTypeValues returns the possible values for the JobType const type.
func PossibleJobTypeValues() []JobType {
	return []JobType{
		JobTypeJupyterSession,
		JobTypeSparkBatch,
		JobTypeSparkSession,
		JobTypeUnknown,
	}
}

// ModeType - The load table operation mode, overwrite or append. Additional mode types may be added over time.
type ModeType string

const (
	// ModeTypeAppend - load table in append mode.
	ModeTypeAppend ModeType = "Append"
	// ModeTypeOverwrite - load table in overwrite mode.
	ModeTypeOverwrite ModeType = "Overwrite"
)

// PossibleModeTypeValues returns the possible values for the ModeType const type.
func PossibleModeTypeValues() []ModeType {
	return []ModeType{
		ModeTypeAppend,
		ModeTypeOverwrite,
	}
}

// Origin - Origin of the job. Values are: SubmittedJob, PendingJob. Additional LivySessionOrigin types may be added over
// time.
type Origin string

const (
	// OriginPendingJob - Job is coming from pending data source.
	OriginPendingJob Origin = "PendingJob"
	// OriginSubmittedJob - Job is coming from submitted data source.
	OriginSubmittedJob Origin = "SubmittedJob"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginPendingJob,
		OriginSubmittedJob,
	}
}

// PathType - The type of relativePath, either file or folder. Additional PathType types may be added over time.
type PathType string

const (
	// PathTypeFile - load table from file.
	PathTypeFile PathType = "File"
	// PathTypeFolder - load table from folder.
	PathTypeFolder PathType = "Folder"
)

// PossiblePathTypeValues returns the possible values for the PathType const type.
func PossiblePathTypeValues() []PathType {
	return []PathType{
		PathTypeFile,
		PathTypeFolder,
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

// SQLEndpointProvisioningStatus - The SQL endpoint provisioning status type. Additional SqlEndpointProvisioningStatus types
// may be added over time.
type SQLEndpointProvisioningStatus string

const (
	// SQLEndpointProvisioningStatusFailed - SQL endpoint provisioning failed.
	SQLEndpointProvisioningStatusFailed SQLEndpointProvisioningStatus = "Failed"
	// SQLEndpointProvisioningStatusInProgress - SQL endpoint provisioning is in progress.
	SQLEndpointProvisioningStatusInProgress SQLEndpointProvisioningStatus = "InProgress"
	// SQLEndpointProvisioningStatusSuccess - SQL endpoint provisioning succeeded.
	SQLEndpointProvisioningStatusSuccess SQLEndpointProvisioningStatus = "Success"
)

// PossibleSQLEndpointProvisioningStatusValues returns the possible values for the SQLEndpointProvisioningStatus const type.
func PossibleSQLEndpointProvisioningStatusValues() []SQLEndpointProvisioningStatus {
	return []SQLEndpointProvisioningStatus{
		SQLEndpointProvisioningStatusFailed,
		SQLEndpointProvisioningStatusInProgress,
		SQLEndpointProvisioningStatusSuccess,
	}
}

// State - Current state of the job. Values are: NotStarted, InProgress, Cancelled, Succeeded, Stopped, Failed, Unknown. Additional
// LivySessionState types may be added over time.
type State string

const (
	// StateCancelled - Job got cancelled.
	StateCancelled State = "Cancelled"
	// StateFailed - Job failed or its session timed out.
	StateFailed State = "Failed"
	// StateInProgress - Job is in running or is cancelling state.
	StateInProgress State = "InProgress"
	// StateNotStarted - Job is queued, is starting or in library packaging state.
	StateNotStarted State = "NotStarted"
	// StateSucceeded - Job has stopped or is in success state.
	StateSucceeded State = "Succeeded"
	// StateUnknown - Job is in invalid state.
	StateUnknown State = "Unknown"
)

// PossibleStateValues returns the possible values for the State const type.
func PossibleStateValues() []State {
	return []State{
		StateCancelled,
		StateFailed,
		StateInProgress,
		StateNotStarted,
		StateSucceeded,
		StateUnknown,
	}
}

// TableType - The table type. Additional TableType types may be added over time.
type TableType string

const (
	// TableTypeExternal - External table.
	TableTypeExternal TableType = "External"
	// TableTypeManaged - Managed table.
	TableTypeManaged TableType = "Managed"
)

// PossibleTableTypeValues returns the possible values for the TableType const type.
func PossibleTableTypeValues() []TableType {
	return []TableType{
		TableTypeExternal,
		TableTypeManaged,
	}
}

// TimeUnit - The unit of time for the duration. Additional duration types may be added over time.
type TimeUnit string

const (
	// TimeUnitDays - Duration in days.
	TimeUnitDays TimeUnit = "Days"
	// TimeUnitHours - Duration in hours.
	TimeUnitHours TimeUnit = "Hours"
	// TimeUnitMinutes - Duration in minutes.
	TimeUnitMinutes TimeUnit = "Minutes"
	// TimeUnitSeconds - Duration in seconds.
	TimeUnitSeconds TimeUnit = "Seconds"
)

// PossibleTimeUnitValues returns the possible values for the TimeUnit const type.
func PossibleTimeUnitValues() []TimeUnit {
	return []TimeUnit{
		TimeUnitDays,
		TimeUnitHours,
		TimeUnitMinutes,
		TimeUnitSeconds,
	}
}
