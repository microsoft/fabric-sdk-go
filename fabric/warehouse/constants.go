// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package warehouse

// CollationType - Warehouse collation. Additional collations may be added over time.
type CollationType string

const (
	// CollationTypeLatin1General100BIN2UTF8 - The default - case-sensitive (CS) collation
	CollationTypeLatin1General100BIN2UTF8 CollationType = "Latin1_General_100_BIN2_UTF8"
	// CollationTypeLatin1General100CIASKSWSSCUTF8 - case-insensitive (CI) collation
	CollationTypeLatin1General100CIASKSWSSCUTF8 CollationType = "Latin1_General_100_CI_AS_KS_WS_SC_UTF8"
)

// PossibleCollationTypeValues returns the possible values for the CollationType const type.
func PossibleCollationTypeValues() []CollationType {
	return []CollationType{
		CollationTypeLatin1General100BIN2UTF8,
		CollationTypeLatin1General100CIASKSWSSCUTF8,
	}
}

// ItemType - The type of the item. Additional item types may be added over time.
type ItemType string

const (
	// ItemTypeCopyJob - A Copy job.
	ItemTypeCopyJob ItemType = "CopyJob"
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
)

// PossibleItemTypeValues returns the possible values for the ItemType const type.
func PossibleItemTypeValues() []ItemType {
	return []ItemType{
		ItemTypeCopyJob,
		ItemTypeDashboard,
		ItemTypeDataPipeline,
		ItemTypeDatamart,
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
	}
}
