// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package sparkjobdefinition

// BackgroundJobsClientRunOnDemandSparkJobDefinitionOptions contains the optional parameters for the BackgroundJobsClient.RunOnDemandSparkJobDefinition
// method.
type BackgroundJobsClientRunOnDemandSparkJobDefinitionOptions struct {
	// Run spark job definition request payload with parameters to use.
	RunSparkJobDefinitionRequest *RunSparkJobDefinitionRequest
}

// ItemsClientBeginCreateSparkJobDefinitionOptions contains the optional parameters for the ItemsClient.BeginCreateSparkJobDefinition
// method.
type ItemsClientBeginCreateSparkJobDefinitionOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// ItemsClientBeginGetSparkJobDefinitionDefinitionOptions contains the optional parameters for the ItemsClient.BeginGetSparkJobDefinitionDefinition
// method.
type ItemsClientBeginGetSparkJobDefinitionDefinitionOptions struct {
	// The format of the spark job definition public definition.
	Format *string

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// ItemsClientBeginUpdateSparkJobDefinitionDefinitionOptions contains the optional parameters for the ItemsClient.BeginUpdateSparkJobDefinitionDefinition
// method.
type ItemsClientBeginUpdateSparkJobDefinitionDefinitionOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string

	// When set to true and the .platform file is provided as part of the definition, the item's metadata is updated using the
	// metadata in the .platform file
	UpdateMetadata *bool
}

// ItemsClientDeleteSparkJobDefinitionOptions contains the optional parameters for the ItemsClient.DeleteSparkJobDefinition
// method.
type ItemsClientDeleteSparkJobDefinitionOptions struct {
	// placeholder for future optional parameters
}

// ItemsClientGetSparkJobDefinitionOptions contains the optional parameters for the ItemsClient.GetSparkJobDefinition method.
type ItemsClientGetSparkJobDefinitionOptions struct {
	// placeholder for future optional parameters
}

// ItemsClientListSparkJobDefinitionsOptions contains the optional parameters for the ItemsClient.NewListSparkJobDefinitionsPager
// method.
type ItemsClientListSparkJobDefinitionsOptions struct {
	// A token for retrieving the next page of results.
	ContinuationToken *string
}

// ItemsClientUpdateSparkJobDefinitionOptions contains the optional parameters for the ItemsClient.UpdateSparkJobDefinition
// method.
type ItemsClientUpdateSparkJobDefinitionOptions struct {
	// placeholder for future optional parameters
}
