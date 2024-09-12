// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package report

// ItemsClientBeginCreateReportOptions contains the optional parameters for the ItemsClient.BeginCreateReport method.
type ItemsClientBeginCreateReportOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ItemsClientBeginGetReportDefinitionOptions contains the optional parameters for the ItemsClient.BeginGetReportDefinition
// method.
type ItemsClientBeginGetReportDefinitionOptions struct {
	// The format of the report public definition.
	Format *string

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ItemsClientBeginUpdateReportDefinitionOptions contains the optional parameters for the ItemsClient.BeginUpdateReportDefinition
// method.
type ItemsClientBeginUpdateReportDefinitionOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string

	// When set to true and the .platform file is provided as part of the definition, the item's metadata is updated using the
	// metadata in the .platform file
	UpdateMetadata *bool
}

// ItemsClientDeleteReportOptions contains the optional parameters for the ItemsClient.DeleteReport method.
type ItemsClientDeleteReportOptions struct {
	// placeholder for future optional parameters
}

// ItemsClientGetReportOptions contains the optional parameters for the ItemsClient.GetReport method.
type ItemsClientGetReportOptions struct {
	// placeholder for future optional parameters
}

// ItemsClientListReportsOptions contains the optional parameters for the ItemsClient.NewListReportsPager method.
type ItemsClientListReportsOptions struct {
	// A token for retrieving the next page of results.
	ContinuationToken *string
}

// ItemsClientUpdateReportOptions contains the optional parameters for the ItemsClient.UpdateReport method.
type ItemsClientUpdateReportOptions struct {
	// placeholder for future optional parameters
}
