// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package eventstream

// CreateEventstreamRequest - Create eventstream request payload.
type CreateEventstreamRequest struct {
	// REQUIRED; The eventstream display name. The display name must follow naming rules according to item type.
	DisplayName *string

	// The eventstream description. Maximum length is 256 characters.
	Description *string
}

// Eventstream - An eventstream object.
type Eventstream struct {
	// REQUIRED; The item type.
	Type *ItemType

	// The item description.
	Description *string

	// The item display name.
	DisplayName *string

	// READ-ONLY; The item ID.
	ID *string

	// READ-ONLY; The workspace ID.
	WorkspaceID *string
}

// Eventstreams - A list of eventstreams.
type Eventstreams struct {
	// REQUIRED; A list of eventstreams.
	Value []Eventstream

	// The token for the next result set batch. If there are no more records, it's removed from the response.
	ContinuationToken *string

	// The URI of the next result set batch. If there are no more records, it's removed from the response.
	ContinuationURI *string
}

// UpdateEventstreamRequest - Update eventstream request.
type UpdateEventstreamRequest struct {
	// The eventstream description. Maximum length is 256 characters.
	Description *string

	// The eventstream display name. The display name must follow naming rules according to item type.
	DisplayName *string
}
