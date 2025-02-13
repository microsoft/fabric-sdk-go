// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package mounteddatafactory

// ItemsClientBeginCreateMountedDataFactoryOptions contains the optional parameters for the ItemsClient.BeginCreateMountedDataFactory
// method.
type ItemsClientBeginCreateMountedDataFactoryOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// ItemsClientBeginGetMountedDataFactoryDefinitionOptions contains the optional parameters for the ItemsClient.BeginGetMountedDataFactoryDefinition
// method.
type ItemsClientBeginGetMountedDataFactoryDefinitionOptions struct {
	// The format of the MountedDataFactory public definition.
	Format *string

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// ItemsClientBeginUpdateMountedDataFactoryDefinitionOptions contains the optional parameters for the ItemsClient.BeginUpdateMountedDataFactoryDefinition
// method.
type ItemsClientBeginUpdateMountedDataFactoryDefinitionOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string

	// When set to true and the .platform file is provided as part of the definition, the item's metadata is updated using the
	// metadata in the .platform file
	UpdateMetadata *bool
}

// ItemsClientDeleteMountedDataFactoryOptions contains the optional parameters for the ItemsClient.DeleteMountedDataFactory
// method.
type ItemsClientDeleteMountedDataFactoryOptions struct {
	// placeholder for future optional parameters
}

// ItemsClientGetMountedDataFactoryOptions contains the optional parameters for the ItemsClient.GetMountedDataFactory method.
type ItemsClientGetMountedDataFactoryOptions struct {
	// placeholder for future optional parameters
}

// ItemsClientListMountedDataFactoriesOptions contains the optional parameters for the ItemsClient.NewListMountedDataFactoriesPager
// method.
type ItemsClientListMountedDataFactoriesOptions struct {
	// A token for retrieving the next page of results.
	ContinuationToken *string
}

// ItemsClientUpdateMountedDataFactoryOptions contains the optional parameters for the ItemsClient.UpdateMountedDataFactory
// method.
type ItemsClientUpdateMountedDataFactoryOptions struct {
	// placeholder for future optional parameters
}
