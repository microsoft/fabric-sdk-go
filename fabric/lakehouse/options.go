// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package lakehouse

// BackgroundJobsClientRunOnDemandTableMaintenanceOptions contains the optional parameters for the BackgroundJobsClient.RunOnDemandTableMaintenance
// method.
type BackgroundJobsClientRunOnDemandTableMaintenanceOptions struct {
	// placeholder for future optional parameters
}

// ItemsClientBeginCreateLakehouseOptions contains the optional parameters for the ItemsClient.BeginCreateLakehouse method.
type ItemsClientBeginCreateLakehouseOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ItemsClientDeleteLakehouseOptions contains the optional parameters for the ItemsClient.DeleteLakehouse method.
type ItemsClientDeleteLakehouseOptions struct {
	// placeholder for future optional parameters
}

// ItemsClientGetLakehouseOptions contains the optional parameters for the ItemsClient.GetLakehouse method.
type ItemsClientGetLakehouseOptions struct {
	// placeholder for future optional parameters
}

// ItemsClientListLakehousesOptions contains the optional parameters for the ItemsClient.NewListLakehousesPager method.
type ItemsClientListLakehousesOptions struct {
	// A token for retrieving the next page of results.
	ContinuationToken *string
}

// ItemsClientUpdateLakehouseOptions contains the optional parameters for the ItemsClient.UpdateLakehouse method.
type ItemsClientUpdateLakehouseOptions struct {
	// placeholder for future optional parameters
}

// TablesClientBeginLoadTableOptions contains the optional parameters for the TablesClient.BeginLoadTable method.
type TablesClientBeginLoadTableOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TablesClientListTablesOptions contains the optional parameters for the TablesClient.NewListTablesPager method.
type TablesClientListTablesOptions struct {
	// Token to retrieve the next page of results, if available.
	ContinuationToken *string

	// The maximum number of results per page to return.
	MaxResults *int32
}
