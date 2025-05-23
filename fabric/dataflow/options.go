// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package dataflow

// BackgroundJobsClientRunOnDemandApplyChangesOptions contains the optional parameters for the BackgroundJobsClient.RunOnDemandApplyChanges
// method.
type BackgroundJobsClientRunOnDemandApplyChangesOptions struct {
	// placeholder for future optional parameters
}

// BackgroundJobsClientRunOnDemandExecuteOptions contains the optional parameters for the BackgroundJobsClient.RunOnDemandExecute
// method.
type BackgroundJobsClientRunOnDemandExecuteOptions struct {
	// Run on-demand item job request payload.
	RunOnDemandItemJobRequest *RunOnDemandDataflowExecuteJobRequest
}

// BackgroundJobsClientScheduleApplyChangesOptions contains the optional parameters for the BackgroundJobsClient.ScheduleApplyChanges
// method.
type BackgroundJobsClientScheduleApplyChangesOptions struct {
	// placeholder for future optional parameters
}

// BackgroundJobsClientScheduleExecuteOptions contains the optional parameters for the BackgroundJobsClient.ScheduleExecute
// method.
type BackgroundJobsClientScheduleExecuteOptions struct {
	// placeholder for future optional parameters
}

// ItemsClientBeginCreateDataflowOptions contains the optional parameters for the ItemsClient.BeginCreateDataflow method.
type ItemsClientBeginCreateDataflowOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// ItemsClientBeginGetDataflowDefinitionOptions contains the optional parameters for the ItemsClient.BeginGetDataflowDefinition
// method.
type ItemsClientBeginGetDataflowDefinitionOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// ItemsClientBeginUpdateDataflowDefinitionOptions contains the optional parameters for the ItemsClient.BeginUpdateDataflowDefinition
// method.
type ItemsClientBeginUpdateDataflowDefinitionOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string

	// When set to true and the .platform file is provided as part of the definition, the item's metadata is updated using the
	// metadata in the .platform file
	UpdateMetadata *bool
}

// ItemsClientDeleteDataflowOptions contains the optional parameters for the ItemsClient.DeleteDataflow method.
type ItemsClientDeleteDataflowOptions struct {
	// placeholder for future optional parameters
}

// ItemsClientGetDataflowOptions contains the optional parameters for the ItemsClient.GetDataflow method.
type ItemsClientGetDataflowOptions struct {
	// placeholder for future optional parameters
}

// ItemsClientListDataflowsOptions contains the optional parameters for the ItemsClient.NewListDataflowsPager method.
type ItemsClientListDataflowsOptions struct {
	// A token for retrieving the next page of results.
	ContinuationToken *string
}

// ItemsClientUpdateDataflowOptions contains the optional parameters for the ItemsClient.UpdateDataflow method.
type ItemsClientUpdateDataflowOptions struct {
	// placeholder for future optional parameters
}
