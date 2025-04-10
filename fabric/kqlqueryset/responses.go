// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package kqlqueryset

// ItemsClientCreateKQLQuerysetResponse contains the response from method ItemsClient.BeginCreateKQLQueryset.
type ItemsClientCreateKQLQuerysetResponse struct {
	// A KQL queryset object.
	KQLQueryset
}

// ItemsClientDeleteKQLQuerysetResponse contains the response from method ItemsClient.DeleteKQLQueryset.
type ItemsClientDeleteKQLQuerysetResponse struct {
	// placeholder for future response values
}

// ItemsClientGetKQLQuerysetDefinitionResponse contains the response from method ItemsClient.BeginGetKQLQuerysetDefinition.
type ItemsClientGetKQLQuerysetDefinitionResponse struct {
	// KQL queryset public definition response.
	DefinitionResponse
}

// ItemsClientGetKQLQuerysetResponse contains the response from method ItemsClient.GetKQLQueryset.
type ItemsClientGetKQLQuerysetResponse struct {
	// A KQL queryset object.
	KQLQueryset
}

// ItemsClientListKQLQuerysetsResponse contains the response from method ItemsClient.NewListKQLQuerysetsPager.
type ItemsClientListKQLQuerysetsResponse struct {
	// A list of KQL querysets.
	KQLQuerysets
}

// ItemsClientUpdateKQLQuerysetDefinitionResponse contains the response from method ItemsClient.BeginUpdateKQLQuerysetDefinition.
type ItemsClientUpdateKQLQuerysetDefinitionResponse struct {
	// placeholder for future response values
}

// ItemsClientUpdateKQLQuerysetResponse contains the response from method ItemsClient.UpdateKQLQueryset.
type ItemsClientUpdateKQLQuerysetResponse struct {
	// A KQL queryset object.
	KQLQueryset
}
