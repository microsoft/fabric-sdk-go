// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package warehouse

import "time"

// CreateWarehouseRequest - Create warehouse request payload.
type CreateWarehouseRequest struct {
	// REQUIRED; The warehouse display name. The display name must follow naming rules according to item type.
	DisplayName *string

	// The warehouse creation payload. Use creationPayload. You can't use definition or creationPayload at the same time.
	CreationPayload *CreationPayload

	// The warehouse description. Maximum length is 256 characters.
	Description *string
}

// CreationPayload - Warehouse item payload
type CreationPayload struct {
	// REQUIRED; The default collation type of the warehouse.
	CollationType *CollationType
}

// Properties - The warehouse item properties.
type Properties struct {
	// REQUIRED; The SQL connection string connected to the workspace containing this warehouse.
	ConnectionString *string

	// REQUIRED; The date and time the warehouse was created.
	CreatedDate *time.Time

	// REQUIRED; The date and time the warehouse was last updated
	LastUpdatedTime *time.Time

	// The collation type of the warehouse.
	CollationType *CollationType
}

// UpdateWarehouseRequest - Update warehouse request.
type UpdateWarehouseRequest struct {
	// The warehouse description. Maximum length is 256 characters.
	Description *string

	// The warehouse display name. The display name must follow naming rules according to item type.
	DisplayName *string
}

// Warehouse - A warehouse object.
type Warehouse struct {
	// REQUIRED; The item type.
	Type *ItemType

	// The item description.
	Description *string

	// The item display name.
	DisplayName *string

	// The warehouse properties.
	Properties *Properties

	// READ-ONLY; The folder ID.
	FolderID *string

	// READ-ONLY; The item ID.
	ID *string

	// READ-ONLY; The workspace ID.
	WorkspaceID *string
}

// Warehouses - A list of warehouses.
type Warehouses struct {
	// REQUIRED; A list of warehouses.
	Value []Warehouse

	// The token for the next result set batch. If there are no more records, it's removed from the response.
	ContinuationToken *string

	// The URI of the next result set batch. If there are no more records, it's removed from the response.
	ContinuationURI *string
}
