// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package digitaltwinbuilder

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
