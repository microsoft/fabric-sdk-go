// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package eventstream

// ItemsClientBeginCreateEventstreamOptions contains the optional parameters for the ItemsClient.BeginCreateEventstream method.
type ItemsClientBeginCreateEventstreamOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ItemsClientDeleteEventstreamOptions contains the optional parameters for the ItemsClient.DeleteEventstream method.
type ItemsClientDeleteEventstreamOptions struct {
	// placeholder for future optional parameters
}

// ItemsClientGetEventstreamOptions contains the optional parameters for the ItemsClient.GetEventstream method.
type ItemsClientGetEventstreamOptions struct {
	// placeholder for future optional parameters
}

// ItemsClientListEventstreamsOptions contains the optional parameters for the ItemsClient.NewListEventstreamsPager method.
type ItemsClientListEventstreamsOptions struct {
	// A token for retrieving the next page of results.
	ContinuationToken *string
}

// ItemsClientUpdateEventstreamOptions contains the optional parameters for the ItemsClient.UpdateEventstream method.
type ItemsClientUpdateEventstreamOptions struct {
	// placeholder for future optional parameters
}
