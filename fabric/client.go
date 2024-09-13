// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package fabric

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"

	"github.com/microsoft/fabric-sdk-go/internal/iruntime"
)

// Client is a Client factory used to create any client in this module.
// Don't use this type directly, use NewClient instead.
type Client struct {
	Internal *azcore.Client
	Endpoint string
}

// NewClient creates a new instance of Client with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - endpoint - pass nil to accept the default values.
//   - options - pass nil to accept the default values.
func NewClient(credential azcore.TokenCredential, endpoint *string, options *azcore.ClientOptions) (*Client, error) {
	sc, err := iruntime.NewServiceClient(credential, endpoint, options)
	if err != nil {
		return nil, err
	}

	return &Client{
		Internal: sc.Internal,
		Endpoint: sc.Endpoint,
	}, nil
}