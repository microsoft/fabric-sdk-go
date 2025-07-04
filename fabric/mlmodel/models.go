// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package mlmodel

// CreateMLModelRequest - Create machine learning model request payload.
type CreateMLModelRequest struct {
	// REQUIRED; The machine learning model display name. The display name must follow naming rules according to item type.
	DisplayName *string

	// The machine learning model description. Maximum length is 256 characters.
	Description *string

	// The folder ID. If not specified or null, the machine learning model is created with the workspace as its folder.
	FolderID *string
}

// ItemTag - Represents a tag applied on an item.
type ItemTag struct {
	// REQUIRED; The name of the tag.
	DisplayName *string

	// REQUIRED; The tag ID.
	ID *string
}

// MLModel - A machine learning model object.
type MLModel struct {
	// REQUIRED; The item type.
	Type *ItemType

	// The item description.
	Description *string

	// The item display name.
	DisplayName *string

	// READ-ONLY; The folder ID.
	FolderID *string

	// READ-ONLY; The item ID.
	ID *string

	// READ-ONLY; List of applied tags.
	Tags []ItemTag

	// READ-ONLY; The workspace ID.
	WorkspaceID *string
}

// MLModels - A list of machine learning models.
type MLModels struct {
	// REQUIRED; A list of machine learning models.
	Value []MLModel

	// The token for the next result set batch. If there are no more records, it's removed from the response.
	ContinuationToken *string

	// The URI of the next result set batch. If there are no more records, it's removed from the response.
	ContinuationURI *string
}

// UpdateMLModelRequest - Update machine learning model request.
type UpdateMLModelRequest struct {
	// The machine learning model description. Maximum length is 256 characters.
	Description *string
}
