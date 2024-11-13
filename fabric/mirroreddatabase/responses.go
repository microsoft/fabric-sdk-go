// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package mirroreddatabase

// ItemsClientCreateMirroredDatabaseResponse contains the response from method ItemsClient.CreateMirroredDatabase.
type ItemsClientCreateMirroredDatabaseResponse struct {
	// A mirrored database item.
	MirroredDatabase
}

// ItemsClientDeleteMirroredDatabaseResponse contains the response from method ItemsClient.DeleteMirroredDatabase.
type ItemsClientDeleteMirroredDatabaseResponse struct {
	// placeholder for future response values
}

// ItemsClientGetMirroredDatabaseDefinitionResponse contains the response from method ItemsClient.GetMirroredDatabaseDefinition.
type ItemsClientGetMirroredDatabaseDefinitionResponse struct {
	// Mirrored database public definition response.
	DefinitionResponse
}

// ItemsClientGetMirroredDatabaseResponse contains the response from method ItemsClient.GetMirroredDatabase.
type ItemsClientGetMirroredDatabaseResponse struct {
	// A mirrored database item.
	MirroredDatabase
}

// ItemsClientListMirroredDatabasesResponse contains the response from method ItemsClient.NewListMirroredDatabasesPager.
type ItemsClientListMirroredDatabasesResponse struct {
	// A list of mirrored databases.
	MirroredDatabases
}

// ItemsClientUpdateMirroredDatabaseDefinitionResponse contains the response from method ItemsClient.UpdateMirroredDatabaseDefinition.
type ItemsClientUpdateMirroredDatabaseDefinitionResponse struct {
	// placeholder for future response values
}

// ItemsClientUpdateMirroredDatabaseResponse contains the response from method ItemsClient.UpdateMirroredDatabase.
type ItemsClientUpdateMirroredDatabaseResponse struct {
	// A mirrored database item.
	MirroredDatabase
}

// MirroringClientGetMirroringStatusResponse contains the response from method MirroringClient.GetMirroringStatus.
type MirroringClientGetMirroringStatusResponse struct {
	// Response of getting mirroring status.
	MirroringStatusResponse
}

// MirroringClientGetTablesMirroringStatusResponse contains the response from method MirroringClient.GetTablesMirroringStatus.
type MirroringClientGetTablesMirroringStatusResponse struct {
	// A paginated list of table mirroring statuses.
	TablesMirroringStatusResponse
}

// MirroringClientStartMirroringResponse contains the response from method MirroringClient.StartMirroring.
type MirroringClientStartMirroringResponse struct {
	// placeholder for future response values
}

// MirroringClientStopMirroringResponse contains the response from method MirroringClient.StopMirroring.
type MirroringClientStopMirroringResponse struct {
	// placeholder for future response values
}
