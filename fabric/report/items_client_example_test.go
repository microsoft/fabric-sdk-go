// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package report_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"

	"github.com/microsoft/fabric-sdk-go/fabric/report"
)

// Generated from example definition
func ExampleItemsClient_NewListReportsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := report.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewItemsClient().NewListReportsPager("cfafbeb1-8037-4d0c-896e-a46fb27ff229", &report.ItemsClientListReportsOptions{ContinuationToken: nil})
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
		// page.Reports = report.Reports{
		// 	Value: []report.Report{
		// 		{
		// 			Type: to.Ptr(report.ItemTypeReport),
		// 			Description: to.Ptr("A report description."),
		// 			DisplayName: to.Ptr("Report Name 1"),
		// 			ID: to.Ptr("3546052c-ae64-4526-b1a8-52af7761426f"),
		// 			WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
		// 	}},
		// }
	}
}

// Generated from example definition
func ExampleItemsClient_BeginCreateReport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := report.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewItemsClient().BeginCreateReport(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", report.CreateReportRequest{
		Description: to.Ptr("A report description."),
		Definition: &report.Definition{
			Parts: []report.DefinitionPart{
				{
					Path:        to.Ptr("definition.pbir"),
					Payload:     to.Ptr("ew0KICAidmVyc2lvbiI..sYVN0eWxlTGl2ZSINCiAgICB9DQogIH0NCn0="),
					PayloadType: to.Ptr(report.PayloadTypeInlineBase64),
				},
				{
					Path:        to.Ptr("report.json"),
					Payload:     to.Ptr("ewogICJjb25maWciOiA..3aWR0aCI6IDEyODAuMDAKICAgIH0KICBdCn0="),
					PayloadType: to.Ptr(report.PayloadTypeInlineBase64),
				},
				{
					Path:        to.Ptr(".platform"),
					Payload:     to.Ptr("ZG90UGxhdGZvcm1CYXNlNjRTdHJpbmc="),
					PayloadType: to.Ptr(report.PayloadTypeInlineBase64),
				}},
		},
		DisplayName: to.Ptr("Report 1"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition
func ExampleItemsClient_GetReport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := report.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewItemsClient().GetReport(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Report = report.Report{
	// 	Type: to.Ptr(report.ItemTypeReport),
	// 	Description: to.Ptr("A report description."),
	// 	DisplayName: to.Ptr("Report 1"),
	// 	ID: to.Ptr("5b218778-e7a5-4d73-8187-f10824047715"),
	// 	WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff229"),
	// }
}

// Generated from example definition
func ExampleItemsClient_UpdateReport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := report.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewItemsClient().UpdateReport(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff211", "5b218778-e7a5-4d73-8187-f10824047721", report.UpdateReportRequest{
		Description: to.Ptr("A new description for report."),
		DisplayName: to.Ptr("Report's New name"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Report = report.Report{
	// 	Type: to.Ptr(report.ItemTypeReport),
	// 	Description: to.Ptr("A new description for report."),
	// 	DisplayName: to.Ptr("Report's New name"),
	// 	ID: to.Ptr("5b218778-e7a5-4d73-8187-f10824047721"),
	// 	WorkspaceID: to.Ptr("cfafbeb1-8037-4d0c-896e-a46fb27ff211"),
	// }
}

// Generated from example definition
func ExampleItemsClient_DeleteReport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := report.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewItemsClient().DeleteReport(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleItemsClient_BeginGetReportDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := report.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewItemsClient().BeginGetReportDefinition(ctx, "6e335e92-a2a2-4b5a-970a-bd6a89fbb765", "cfafbeb1-8037-4d0c-896e-a46fb27ff229", &report.ItemsClientBeginGetReportDefinitionOptions{Format: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DefinitionResponse = report.DefinitionResponse{
	// 	Definition: &report.Definition{
	// 		Parts: []report.DefinitionPart{
	// 			{
	// 				Path: to.Ptr("report.json"),
	// 				Payload: to.Ptr("QmFzZTY0U3RyaW5n"),
	// 				PayloadType: to.Ptr(report.PayloadTypeInlineBase64),
	// 			},
	// 			{
	// 				Path: to.Ptr("definition.pbir"),
	// 				Payload: to.Ptr("QW5vdGhlckJhc2U2NFN0cmluZw"),
	// 				PayloadType: to.Ptr(report.PayloadTypeInlineBase64),
	// 			},
	// 			{
	// 				Path: to.Ptr(".platform"),
	// 				Payload: to.Ptr("ZG90UGxhdGZvcm1CYXNlNjRTdHJpbmc="),
	// 				PayloadType: to.Ptr(report.PayloadTypeInlineBase64),
	// 		}},
	// 	},
	// }
}

// Generated from example definition
func ExampleItemsClient_BeginUpdateReportDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := report.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewItemsClient().BeginUpdateReportDefinition(ctx, "cfafbeb1-8037-4d0c-896e-a46fb27ff229", "5b218778-e7a5-4d73-8187-f10824047715", report.UpdateReportDefinitionRequest{
		Definition: &report.Definition{
			Parts: []report.DefinitionPart{
				{
					Path:        to.Ptr("report.json"),
					Payload:     to.Ptr("QmFzZTY0U3RyaW5n"),
					PayloadType: to.Ptr(report.PayloadTypeInlineBase64),
				},
				{
					Path:        to.Ptr("definition.pbir"),
					Payload:     to.Ptr("QW5vdGhlckJhc2U2NFN0cmluZw"),
					PayloadType: to.Ptr(report.PayloadTypeInlineBase64),
				},
				{
					Path:        to.Ptr(".platform"),
					Payload:     to.Ptr("ZG90UGxhdGZvcm1CYXNlNjRTdHJpbmc="),
					PayloadType: to.Ptr(report.PayloadTypeInlineBase64),
				}},
		},
	}, &report.ItemsClientBeginUpdateReportDefinitionOptions{UpdateMetadata: to.Ptr(true)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
