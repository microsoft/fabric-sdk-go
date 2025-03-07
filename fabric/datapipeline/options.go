// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package datapipeline

// ItemsClientBeginCreateDataPipelineOptions contains the optional parameters for the ItemsClient.BeginCreateDataPipeline
// method.
type ItemsClientBeginCreateDataPipelineOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// ItemsClientBeginGetDataPipelineDefinitionOptions contains the optional parameters for the ItemsClient.BeginGetDataPipelineDefinition
// method.
type ItemsClientBeginGetDataPipelineDefinitionOptions struct {
	// The format of the data pipeline public definition.
	Format *string

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// ItemsClientBeginUpdateDataPipelineDefinitionOptions contains the optional parameters for the ItemsClient.BeginUpdateDataPipelineDefinition
// method.
type ItemsClientBeginUpdateDataPipelineDefinitionOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string

	// When set to true and the .platform file is provided as part of the definition, the item's metadata is updated using the
	// metadata in the .platform file
	UpdateMetadata *bool
}

// ItemsClientDeleteDataPipelineOptions contains the optional parameters for the ItemsClient.DeleteDataPipeline method.
type ItemsClientDeleteDataPipelineOptions struct {
	// placeholder for future optional parameters
}

// ItemsClientGetDataPipelineOptions contains the optional parameters for the ItemsClient.GetDataPipeline method.
type ItemsClientGetDataPipelineOptions struct {
	// placeholder for future optional parameters
}

// ItemsClientListDataPipelinesOptions contains the optional parameters for the ItemsClient.NewListDataPipelinesPager method.
type ItemsClientListDataPipelinesOptions struct {
	// A token for retrieving the next page of results.
	ContinuationToken *string
}

// ItemsClientUpdateDataPipelineOptions contains the optional parameters for the ItemsClient.UpdateDataPipeline method.
type ItemsClientUpdateDataPipelineOptions struct {
	// placeholder for future optional parameters
}
