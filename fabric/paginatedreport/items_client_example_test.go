// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package paginatedreport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"

	"github.com/microsoft/fabric-sdk-go/fabric/paginatedreport"
)

// Generated from example definition
func ExampleItemsClient_NewListPaginatedReportsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := paginatedreport.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewItemsClient().NewListPaginatedReportsPager("cfafbeb1-8037-4d0c-896e-a46fb27ff229", &paginatedreport.ItemsClientListPaginatedReportsOptions{ContinuationToken: nil})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PaginatedReports = paginatedreport.PaginatedReports{
		// 	Value: []paginatedreport.PaginatedReport{
		// 		{
		// 			Type: to.Ptr(paginatedreport.ItemTypePaginatedReport),
		// 			Description: to.Ptr("A paginated report description."),
		// 			DisplayName: to.Ptr("PaginatedReport Name 1"),
		// 			ID: to.Ptr("3546052c-ae64-4526-b1a8-52af7761426f"),
		// 			WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleItemsClient_UpdatePaginatedReport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := paginatedreport.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewItemsClient().UpdatePaginatedReport(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff233", "5b218778-e7a5-4d73-8187-f10824047731", paginatedreport.UpdatePaginatedReportRequest{
		Description: to.Ptr("A new description for paginated report."),
		DisplayName: to.Ptr("Paginated Report's New name"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PaginatedReport = paginatedreport.PaginatedReport{
	// 	Type: to.Ptr(paginatedreport.ItemTypePaginatedReport),
	// 	Description: to.Ptr("A new description for paginated report."),
	// 	DisplayName: to.Ptr("Paginated Report's New name"),
	// 	ID: to.Ptr("5b218778-e7a5-4d73-8187-f10824047731"),
	// 	WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff233"),
	// }
}