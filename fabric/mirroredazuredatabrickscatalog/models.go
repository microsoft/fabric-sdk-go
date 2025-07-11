// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package mirroredazuredatabrickscatalog

import "time"

// CreateMirroredAzureDatabricksCatalogRequest - Create MirroredAzureDatabricksCatalog request payload.
type CreateMirroredAzureDatabricksCatalogRequest struct {
	// REQUIRED; The MirroredAzureDatabricksCatalog display name. The display name must follow naming rules according to item
	// type.
	DisplayName *string

	// The MirroredAzureDatabricksCatalog creation payload.
	CreationPayload *CreationPayload

	// The MirroredAzureDatabricksCatalog public definition. Use definition or creationPayload. You can't use both at the same
	// time.
	Definition *PublicDefinition

	// The MirroredAzureDatabricksCatalog description. Maximum length is 256 characters.
	Description *string
}

// CreationPayload - MirroredAzureDatabricksCatalog create item payload.
type CreationPayload struct {
	// REQUIRED; Azure databricks catalog name.
	CatalogName *string

	// REQUIRED; The Azure databricks workspace connection id.
	DatabricksWorkspaceConnectionID *string

	// REQUIRED; Mirroring mode.
	MirroringMode *MirroringModes

	// The storage connection id.
	StorageConnectionID *string
}

// DatabricksCatalog - A catalog from Unity Catalog
type DatabricksCatalog struct {
	// REQUIRED; The type of the catalog.
	CatalogType *CatalogType

	// REQUIRED; The full name of the catalog. This variable depends on name.
	FullName *string

	// REQUIRED; The name of the catalog.
	Name *string

	// REQUIRED; The storage location of the catalog.
	StorageLocation *string
}

// DatabricksCatalogs - A list of catalogs from Unity Catalog.
type DatabricksCatalogs struct {
	// REQUIRED; A list of Catalogs.
	Value []DatabricksCatalog

	// The token for the next result set batch. If there are no more records, it's removed from the response.
	ContinuationToken *string

	// The URI of the next result set batch. If there are no more records, it's removed from the response.
	ContinuationURI *string

	// READ-ONLY; Error is set if unable to fetch catalogs
	Error *ErrorResponse
}

// DatabricksSchema - A schema from Unity Catalog
type DatabricksSchema struct {
	// REQUIRED; The full name of the schema, in the form of catalogname.schemaname.
	FullName *string

	// REQUIRED; The name of the schema, relative to the parent catalog.
	Name *string

	// The storage location of the schema.
	StorageLocation *string
}

// DatabricksSchemas - A list of schemas from Unity Catalog.
type DatabricksSchemas struct {
	// REQUIRED; A list of Databricks schemas.
	Value []DatabricksSchema

	// The token for the next result set batch. If there are no more records, it's removed from the response.
	ContinuationToken *string

	// The URI of the next result set batch. If there are no more records, it's removed from the response.
	ContinuationURI *string

	// READ-ONLY; Error is set if unable to fetch schemas
	Error *ErrorResponse
}

// DatabricksTable - A table from Unity Catalog
type DatabricksTable struct {
	// REQUIRED; The data source format of the table.
	DataSourceFormat *DataSourceFormat

	// REQUIRED; The full name of the table, in the form of catalogname.schemaname.table_name.
	FullName *string

	// REQUIRED; The name of the table, relative to the parent schema.
	Name *string

	// REQUIRED; The storage location of the table.
	StorageLocation *string

	// REQUIRED; The type of the table.
	TableType *TableType
}

// DatabricksTables - A list of tables from Unity Catalog.
type DatabricksTables struct {
	// REQUIRED; A list of Databricks tables.
	Value []DatabricksTable

	// The token for the next result set batch. If there are no more records, it's removed from the response.
	ContinuationToken *string

	// The URI of the next result set batch. If there are no more records, it's removed from the response.
	ContinuationURI *string

	// READ-ONLY; Error is set if unable to fetch tables
	Error *ErrorResponse
}

// DefinitionResponse - API for mirroredAzureDatabricksCatalog public definition response.
type DefinitionResponse struct {
	// READ-ONLY; MirroredAzureDatabricksCatalog public definition object. Refer to this article [/rest/api/fabric/articles/item-management/definitions/mirroredazuredatabrickscatalog-definition]
	// for more details on how
	// to craft a MirroredAzureDatabricksCatalog definition public definition.
	Definition *PublicDefinition
}

// ErrorInfo - The error information.
type ErrorInfo struct {
	// REQUIRED; The error code.
	ErrorCode *string

	// REQUIRED; The error message.
	ErrorMessage *string

	// The error details.
	ErrorDetails *string
}

// ErrorRelatedResource - The error related resource details object.
type ErrorRelatedResource struct {
	// READ-ONLY; The resource ID that's involved in the error.
	ResourceID *string

	// READ-ONLY; The type of the resource that's involved in the error.
	ResourceType *string
}

// ErrorResponse - The error response.
type ErrorResponse struct {
	// READ-ONLY; A specific identifier that provides information about an error condition, allowing for standardized communication
	// between our service and its users.
	ErrorCode *string

	// READ-ONLY; A human readable representation of the error.
	Message *string

	// READ-ONLY; List of additional error details.
	MoreDetails []ErrorResponseDetails

	// READ-ONLY; The error related resource details.
	RelatedResource *ErrorRelatedResource

	// READ-ONLY; ID of the request associated with the error.
	RequestID *string
}

// ErrorResponseDetails - The error response details.
type ErrorResponseDetails struct {
	// READ-ONLY; A specific identifier that provides information about an error condition, allowing for standardized communication
	// between our service and its users.
	ErrorCode *string

	// READ-ONLY; A human readable representation of the error.
	Message *string

	// READ-ONLY; The error related resource details.
	RelatedResource *ErrorRelatedResource
}

// ItemTag - Represents a tag applied on an item.
type ItemTag struct {
	// REQUIRED; The name of the tag.
	DisplayName *string

	// REQUIRED; The tag ID.
	ID *string
}

// MirroredAzureDatabricksCatalog - A MirroredAzureDatabricksCatalog item.
type MirroredAzureDatabricksCatalog struct {
	// REQUIRED; The item type.
	Type *ItemType

	// The item description.
	Description *string

	// The item display name.
	DisplayName *string

	// The MirroredAzureDatabricksCatalog properties.
	Properties *Properties

	// READ-ONLY; The folder ID.
	FolderID *string

	// READ-ONLY; The item ID.
	ID *string

	// READ-ONLY; List of applied tags.
	Tags []ItemTag

	// READ-ONLY; The workspace ID.
	WorkspaceID *string
}

// MirroredAzureDatabricksCatalogs - A list of MirroredAzureDatabricksCatalogs.
type MirroredAzureDatabricksCatalogs struct {
	// REQUIRED; A list of MirroredAzureDatabricksCatalogs.
	Value []MirroredAzureDatabricksCatalog

	// The token for the next result set batch. If there are no more records, it's removed from the response.
	ContinuationToken *string

	// The URI of the next result set batch. If there are no more records, it's removed from the response.
	ContinuationURI *string
}

// Properties - The MirroredAzureDatabricksCatalog properties.
type Properties struct {
	// REQUIRED; Auto sync the catalog. Additional autoSync types may be added over time.
	AutoSync *AutoSync

	// REQUIRED; Azure databricks catalog name.
	CatalogName *string

	// REQUIRED; The Azure databricks workspace connection id.
	DatabricksWorkspaceConnectionID *string

	// REQUIRED; Mirroring mode. Additional mirroringMode may be added over time.
	MirroringMode *MirroringModes

	// REQUIRED; OneLake path to the MirroredAzureDatabricksCatalog tables directory.
	OneLakeTablesPath *string

	// The MirroredAzureDatabricksCatalog sync status.
	MirrorStatus *MirrorStatus

	// An object containing the properties of the SQL endpoint.
	SQLEndpointProperties *SQLEndpointProperties

	// The storage connection id.
	StorageConnectionID *string

	// The MirroredAzureDatabricksCatalog sync status.
	SyncDetails *SyncDetails
}

// PublicDefinition - MirroredAzureDatabricksCatalog public definition object. Refer to this article [/rest/api/fabric/articles/item-management/definitions/mirroredazuredatabrickscatalog-definition]
// for more details on how
// to craft a MirroredAzureDatabricksCatalog definition public definition.
type PublicDefinition struct {
	// REQUIRED; A list of definition parts.
	Parts []PublicDefinitionPart

	// The format of the item definition.
	Format *string
}

// PublicDefinitionPart - MirroredAzureDatabricksCatalog definition part object.
type PublicDefinitionPart struct {
	// REQUIRED; The MirroredAzureDatabricksCatalog part path.
	Path *string

	// REQUIRED; The MirroredAzureDatabricksCatalog part payload.
	Payload *string

	// REQUIRED; The payload type.
	PayloadType *PayloadType
}

// SQLEndpointProperties - An object containing the properties of the SQL endpoint.
type SQLEndpointProperties struct {
	// REQUIRED; SQL endpoint connection string.
	ConnectionString *string

	// REQUIRED; SQL endpoint ID.
	ID *string
}

// SyncDetails - The MirroredAzureDatabricksCatalog mirroring status.
type SyncDetails struct {
	// REQUIRED; The last sync date time in UTC, using the YYYY-MM-DDTHH:mm:ssZ format.
	LastSyncDateTime *time.Time

	// REQUIRED; The sync status. Additional status may be added over time.
	Status *Status

	// The error information.
	ErrorInfo *ErrorInfo
}

// UpdateMirroredAzureDatabricksCatalogRequest - Update MirroredAzureDatabricksCatalog request.
type UpdateMirroredAzureDatabricksCatalogRequest struct {
	// The MirroredAzureDatabricksCatalog description. Maximum length is 256 characters.
	Description *string

	// The MirroredAzureDatabricksCatalog display name. The display name must follow naming rules according to item type.
	DisplayName *string

	// The MirroredAzureDatabricksCatalog updateable properties payload.
	PublicUpdateableExtendedProperties *UpdatePayload
}

// UpdatePayload - MirroredAzureDatabricksCatalog update item payload.
type UpdatePayload struct {
	// Auto sync the catalog.
	AutoSync *AutoSync

	// Mirroring mode.
	MirroringMode *MirroringModes

	// The storage connection id.
	StorageConnectionID *string
}

// UpdatemirroredAzureDatabricksCatalogDefinitionRequest - Update MirroredAzureDatabricksCatalog public definition request
// payload.
type UpdatemirroredAzureDatabricksCatalogDefinitionRequest struct {
	// REQUIRED; MirroredAzureDatabricksCatalog public definition object. Refer to this article [/rest/api/fabric/articles/item-management/definitions/mirroredazuredatabrickscatalog-definition]
	// for more details on how
	// to craft a MirroredAzureDatabricksCatalog definition public definition.
	Definition *PublicDefinition
}
