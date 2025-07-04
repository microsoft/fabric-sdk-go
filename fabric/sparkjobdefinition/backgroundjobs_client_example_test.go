// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package sparkjobdefinition_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"

	"github.com/microsoft/fabric-sdk-go/fabric/sparkjobdefinition"
)

// Generated from example definition
func ExampleBackgroundJobsClient_RunOnDemandSparkJobDefinition_runSparkJobDefinitionWithNoRequestBody() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := sparkjobdefinition.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewBackgroundJobsClient().RunOnDemandSparkJobDefinition(ctx, "4b218778-e7a5-4d73-8187-f10824047715", "431e8d7b-4a95-4c02-8ccd-6faef5ba1bd7", "sparkjob", &sparkjobdefinition.BackgroundJobsClientRunOnDemandSparkJobDefinitionOptions{RunSparkJobDefinitionRequest: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition
func ExampleBackgroundJobsClient_RunOnDemandSparkJobDefinition_runSparkJobDefinitionWithRequestBody() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := sparkjobdefinition.NewClientFactory(cred, nil, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewBackgroundJobsClient().RunOnDemandSparkJobDefinition(ctx, "4b218778-e7a5-4d73-8187-f10824047715", "431e8d7b-4a95-4c02-8ccd-6faef5ba1bd7", "sparkjob", &sparkjobdefinition.BackgroundJobsClientRunOnDemandSparkJobDefinitionOptions{RunSparkJobDefinitionRequest: &sparkjobdefinition.RunSparkJobDefinitionRequest{
		ExecutionData: &sparkjobdefinition.ExecutionData{
			AdditionalLibraryUris: []string{
				"abfss://test@onelakecst180.dfs.pbidedicated.windows-int.net/dfsd.Lakehouse/Files/testfile.jar"},
			CommandLineArguments: to.Ptr("firstarg secondarg thirdarg"),
			DefaultLakehouseID: &sparkjobdefinition.ItemReferenceByID{
				ReferenceType: to.Ptr(sparkjobdefinition.ItemReferenceTypeByID),
				ItemID:        to.Ptr("01d11fd3-625b-4c89-880c-3fc0ad19e734"),
				WorkspaceID:   to.Ptr("4b218778-e7a5-4d73-8187-f10824047715"),
			},
			EnvironmentID: &sparkjobdefinition.ItemReferenceByID{
				ReferenceType: to.Ptr(sparkjobdefinition.ItemReferenceTypeByID),
				ItemID:        to.Ptr("937e2b52-320c-4e16-b232-f9907e433a0d"),
				WorkspaceID:   to.Ptr("4b218778-e7a5-4d73-8187-f10824047715"),
			},
			ExecutableFile: to.Ptr("abfss://test@northcentralus-onelake.dfs.fabric.microsoft.com/salesdata.Lakehouse/Files/oneplusoneapp.jar"),
			MainClass:      to.Ptr("com.microsoft.spark.example.OneplusOneApp"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
