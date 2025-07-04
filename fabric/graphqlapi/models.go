// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package graphqlapi

// CreateGraphQLAPIRequest - Create API for GraphQL request payload.
type CreateGraphQLAPIRequest struct {
	// REQUIRED; The API for GraphQL display name. The display name must follow naming rules according to item type.
	DisplayName *string

	// The API for GraphQL public definition.
	Definition *PublicDefinition

	// The API for GraphQL description. Maximum length is 256 characters.
	Description *string

	// The folder ID. If not specified or null, the API for GraphQL is created with the workspace as its folder.
	FolderID *string
}

// DefinitionResponse - API for GraphQL public definition response.
type DefinitionResponse struct {
	// READ-ONLY; API for GraphQL public definition object. To create the definition, see GraphQLApi definition [/rest/api/fabric/articles/item-management/definitions/graphql-api-definition].
	Definition *PublicDefinition
}

// GraphQLAPI - An API for GraphQL item.
type GraphQLAPI struct {
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

// GraphQLApis - A list of API for GraphQL items.
type GraphQLApis struct {
	// REQUIRED; A list of API for GraphQL items.
	Value []GraphQLAPI

	// The token for the next result set batch. If there are no more records, it's removed from the response.
	ContinuationToken *string

	// The URI of the next result set batch. If there are no more records, it's removed from the response.
	ContinuationURI *string
}

// ItemTag - Represents a tag applied on an item.
type ItemTag struct {
	// REQUIRED; The name of the tag.
	DisplayName *string

	// REQUIRED; The tag ID.
	ID *string
}

// PublicDefinition - API for GraphQL public definition object. To create the definition, see GraphQLApi definition [/rest/api/fabric/articles/item-management/definitions/graphql-api-definition].
type PublicDefinition struct {
	// REQUIRED; A list of definition parts.
	Parts []PublicDefinitionPart

	// The format of the item definition.
	Format *string
}

// PublicDefinitionPart - API for GraphQL definition part object.
type PublicDefinitionPart struct {
	// The API for GraphQL definition part path.
	Path *string

	// The API for GraphQL definition part payload.
	Payload *string

	// The payload type.
	PayloadType *PayloadType
}

// UpdateGraphQLAPIDefinitionRequest - Update API for GraphQL public definition request payload.
type UpdateGraphQLAPIDefinitionRequest struct {
	// REQUIRED; API for GraphQL public definition object. To create the definition, see GraphQLApi definition [/rest/api/fabric/articles/item-management/definitions/graphql-api-definition].
	Definition *PublicDefinition
}

// UpdateGraphQLAPIRequest - Update API for GraphQL request.
type UpdateGraphQLAPIRequest struct {
	// The API for GraphQL description. Maximum length is 256 characters.
	Description *string

	// The API for GraphQL display name. The display name must follow naming rules according to item type.
	DisplayName *string
}
