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
