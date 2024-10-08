// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package datapipeline

// CreateDataPipelineRequest - Create data pipeline request payload.
type CreateDataPipelineRequest struct {
	// REQUIRED; The data pipeline display name.
	DisplayName *string

	// The data pipeline description. Maximum length is 256 characters.
	Description *string
}

// DataPipeline - A data pipeline object.
type DataPipeline struct {
	// REQUIRED; The item type.
	Type *ItemType

	// The item description.
	Description *string

	// The item display name.
	DisplayName *string

	// READ-ONLY; The item ID.
	ID *string

	// READ-ONLY; The workspace ID.
	WorkspaceID *string
}

// DataPipelines - A list of data pipelines.
type DataPipelines struct {
	// REQUIRED; A list of data pipelines.
	Value []DataPipeline

	// The token for the next result set batch. If there are no more records, it's removed from the response.
	ContinuationToken *string

	// The URI of the next result set batch. If there are no more records, it's removed from the response.
	ContinuationURI *string
}

// UpdateDataPipelineRequest - Update data pipeline request.
type UpdateDataPipelineRequest struct {
	// The data pipeline description. Maximum length is 256 characters.
	Description *string

	// The data pipeline display name. The display name must follow naming rules according to item type.
	DisplayName *string
}
