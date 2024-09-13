// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package semanticmodel

// CreateSemanticModelRequest - Create semantic model request payload.
type CreateSemanticModelRequest struct {
	// REQUIRED; The semantic model public definition.
	Definition *Definition

	// REQUIRED; The semantic model display name. The display name must follow naming rules according to item type.
	DisplayName *string

	// The semantic model description. Maximum length is 256 characters.
	Description *string
}

// Definition - Semantic model public definition [/rest/api/fabric/articles/item-management/definitions/semantic-model-definition]
// object.
type Definition struct {
	// REQUIRED; A list of definition parts.
	Parts []DefinitionPart

	// The format of the item definition.
	Format *string
}

// DefinitionPart - Semantic model definition part object.
type DefinitionPart struct {
	// The semantic model part path.
	Path *string

	// The semantic model part payload.
	Payload *string

	// The payload type.
	PayloadType *PayloadType
}

// DefinitionResponse - Semantic model public definition response.
type DefinitionResponse struct {
	// READ-ONLY; Semantic model public definition [/rest/api/fabric/articles/item-management/definitions/semantic-model-definition]
	// object.
	Definition *Definition
}

// SemanticModel - A semantic model object.
type SemanticModel struct {
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

// SemanticModels - A list of semantic models.
type SemanticModels struct {
	// REQUIRED; A list of semantic models.
	Value []SemanticModel

	// The token for the next result set batch. If there are no more records, it's removed from the response.
	ContinuationToken *string

	// The URI of the next result set batch. If there are no more records, it's removed from the response.
	ContinuationURI *string
}

// UpdateSemanticModelDefinitionRequest - Update semantic model public definition request payload.
type UpdateSemanticModelDefinitionRequest struct {
	// REQUIRED; Semantic model public definition [/rest/api/fabric/articles/item-management/definitions/semantic-model-definition]
	// object.
	Definition *Definition
}

// UpdateSemanticModelRequest - Update semantic model request.
type UpdateSemanticModelRequest struct {
	// The semantic model description. Maximum length is 256 characters.
	Description *string

	// The semantic model display name. The display name must follow naming rules according to item type.
	DisplayName *string
}