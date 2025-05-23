// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package notebook

// ItemsClientBeginCreateNotebookOptions contains the optional parameters for the ItemsClient.BeginCreateNotebook method.
type ItemsClientBeginCreateNotebookOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// ItemsClientBeginGetNotebookDefinitionOptions contains the optional parameters for the ItemsClient.BeginGetNotebookDefinition
// method.
type ItemsClientBeginGetNotebookDefinitionOptions struct {
	// The format of the notebook public definition. Supported format: ipynb and fabricGitSource. If no format is provided, fabricGitSource
	// is used.
	Format *string

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// ItemsClientBeginUpdateNotebookDefinitionOptions contains the optional parameters for the ItemsClient.BeginUpdateNotebookDefinition
// method.
type ItemsClientBeginUpdateNotebookDefinitionOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string

	// When set to true and the .platform file is provided as part of the definition, the item's metadata is updated using the
	// metadata in the .platform file
	UpdateMetadata *bool
}

// ItemsClientDeleteNotebookOptions contains the optional parameters for the ItemsClient.DeleteNotebook method.
type ItemsClientDeleteNotebookOptions struct {
	// placeholder for future optional parameters
}

// ItemsClientGetNotebookOptions contains the optional parameters for the ItemsClient.GetNotebook method.
type ItemsClientGetNotebookOptions struct {
	// placeholder for future optional parameters
}

// ItemsClientListNotebooksOptions contains the optional parameters for the ItemsClient.NewListNotebooksPager method.
type ItemsClientListNotebooksOptions struct {
	// A token for retrieving the next page of results.
	ContinuationToken *string
}

// ItemsClientUpdateNotebookOptions contains the optional parameters for the ItemsClient.UpdateNotebook method.
type ItemsClientUpdateNotebookOptions struct {
	// placeholder for future optional parameters
}

// LivySessionsClientGetLivySessionOptions contains the optional parameters for the LivySessionsClient.GetLivySession method.
type LivySessionsClientGetLivySessionOptions struct {
	// placeholder for future optional parameters
}

// LivySessionsClientListLivySessionsOptions contains the optional parameters for the LivySessionsClient.NewListLivySessionsPager
// method.
type LivySessionsClientListLivySessionsOptions struct {
	// Token to retrieve the next page of results, if available.
	ContinuationToken *string
}
