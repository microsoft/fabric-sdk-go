// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package kqldatabase

// ItemsClientCreateKQLDatabaseResponse contains the response from method ItemsClient.BeginCreateKQLDatabase.
type ItemsClientCreateKQLDatabaseResponse struct {
	// A KQL database object.
	KQLDatabase
}

// ItemsClientDeleteKQLDatabaseResponse contains the response from method ItemsClient.DeleteKQLDatabase.
type ItemsClientDeleteKQLDatabaseResponse struct {
	// placeholder for future response values
}

// ItemsClientGetKQLDatabaseDefinitionResponse contains the response from method ItemsClient.BeginGetKQLDatabaseDefinition.
type ItemsClientGetKQLDatabaseDefinitionResponse struct {
	// KQL database public definition response.
	DefinitionResponse
}

// ItemsClientGetKQLDatabaseResponse contains the response from method ItemsClient.GetKQLDatabase.
type ItemsClientGetKQLDatabaseResponse struct {
	// A KQL database object.
	KQLDatabase
}

// ItemsClientListKQLDatabasesResponse contains the response from method ItemsClient.NewListKQLDatabasesPager.
type ItemsClientListKQLDatabasesResponse struct {
	// A list of KQL databases.
	KQLDatabases
}

// ItemsClientUpdateKQLDatabaseDefinitionResponse contains the response from method ItemsClient.BeginUpdateKQLDatabaseDefinition.
type ItemsClientUpdateKQLDatabaseDefinitionResponse struct {
	// placeholder for future response values
}

// ItemsClientUpdateKQLDatabaseResponse contains the response from method ItemsClient.UpdateKQLDatabase.
type ItemsClientUpdateKQLDatabaseResponse struct {
	// A KQL database object.
	KQLDatabase
}
