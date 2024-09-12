// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package dashboard

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"

	"github.com/microsoft/fabric-sdk-go/fabric"
	"github.com/microsoft/fabric-sdk-go/internal/iruntime"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	internal *azcore.Client
	endpoint string
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - endpoint - pass nil to accept the default values.
//   - options - pass nil to accept the default values.
func NewClientFactory(credential azcore.TokenCredential, endpoint *string, options *azcore.ClientOptions) (*ClientFactory, error) {
	sc, err := iruntime.NewServiceClient(credential, endpoint, options)
	if err != nil {
		return nil, err
	}

	return &ClientFactory{
		internal: sc.Internal,
		endpoint: sc.Endpoint,
	}, nil
}

// NewClientFactoryWithClient creates a new instance of ClientFactory with sharable Client.
// The Client will be propagated to any client created from this factory.
//   - client - Client created in the containing module: github.com/microsoft/fabric-sdk-go/fabric
func NewClientFactoryWithClient(client fabric.Client) *ClientFactory {
	return &ClientFactory{
		internal: client.Internal,
		endpoint: client.Endpoint,
	}
}

// NewItemsClient creates a new instance of ItemsClient.
func (c *ClientFactory) NewItemsClient() *ItemsClient {
	return &ItemsClient{
		internal: c.internal.WithClientName("dashboard.ItemsClient"),
		endpoint: c.endpoint,
	}
}
