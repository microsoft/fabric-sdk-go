// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package mlexperiment

// CreateMLExperimentRequest - Create machine learning experiment request payload.
type CreateMLExperimentRequest struct {
	// REQUIRED; The machine learning experiment display name. The display name must follow naming rules according to item type.
	DisplayName *string

	// The machine learning experiment description. Maximum length is 256 characters.
	Description *string
}

// MLExperiment - A machine learning experiment object.
type MLExperiment struct {
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

// MLExperiments - A list of machine learning experiments.
type MLExperiments struct {
	// REQUIRED; A list of machine learning experiments.
	Value []MLExperiment

	// The token for the next result set batch. If there are no more records, it's removed from the response.
	ContinuationToken *string

	// The URI of the next result set batch. If there are no more records, it's removed from the response.
	ContinuationURI *string
}

// UpdateMLExperimentRequest - Update machine learning experiment request.
type UpdateMLExperimentRequest struct {
	// The machine learning experiment description. Maximum length is 256 characters.
	Description *string

	// The machine learning experiment display name. The display name must follow naming rules according to item type.
	DisplayName *string
}