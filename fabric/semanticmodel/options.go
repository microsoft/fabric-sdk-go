// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package semanticmodel

// ItemsClientBeginCreateSemanticModelOptions contains the optional parameters for the ItemsClient.BeginCreateSemanticModel
// method.
type ItemsClientBeginCreateSemanticModelOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ItemsClientBeginGetSemanticModelDefinitionOptions contains the optional parameters for the ItemsClient.BeginGetSemanticModelDefinition
// method.
type ItemsClientBeginGetSemanticModelDefinitionOptions struct {
	// The format of the semantic model definition.
	// The following formats are allowed (case sensitive)
	// * - TMDL
	//
	//
	// * - TMSL
	//
	// If not specified, the default is 'TMDL'.
	Format *string

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ItemsClientBeginUpdateSemanticModelDefinitionOptions contains the optional parameters for the ItemsClient.BeginUpdateSemanticModelDefinition
// method.
type ItemsClientBeginUpdateSemanticModelDefinitionOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string

	// When set to true and the .platform file is provided as part of the definition, the item's metadata is updated using the
	// metadata in the .platform file
	UpdateMetadata *bool
}

// ItemsClientDeleteSemanticModelOptions contains the optional parameters for the ItemsClient.DeleteSemanticModel method.
type ItemsClientDeleteSemanticModelOptions struct {
	// placeholder for future optional parameters
}

// ItemsClientGetSemanticModelOptions contains the optional parameters for the ItemsClient.GetSemanticModel method.
type ItemsClientGetSemanticModelOptions struct {
	// placeholder for future optional parameters
}

// ItemsClientListSemanticModelsOptions contains the optional parameters for the ItemsClient.NewListSemanticModelsPager method.
type ItemsClientListSemanticModelsOptions struct {
	// A token for retrieving the next page of results.
	ContinuationToken *string
}

// ItemsClientUpdateSemanticModelOptions contains the optional parameters for the ItemsClient.UpdateSemanticModel method.
type ItemsClientUpdateSemanticModelOptions struct {
	// placeholder for future optional parameters
}