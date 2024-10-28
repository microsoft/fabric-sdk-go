// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package core

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

// NewCapacitiesClient creates a new instance of CapacitiesClient.
func (c *ClientFactory) NewCapacitiesClient() *CapacitiesClient {
	return &CapacitiesClient{
		internal: c.internal.WithClientName("core.CapacitiesClient"),
		endpoint: c.endpoint,
	}
}

// NewDeploymentPipelinesClient creates a new instance of DeploymentPipelinesClient.
func (c *ClientFactory) NewDeploymentPipelinesClient() *DeploymentPipelinesClient {
	return &DeploymentPipelinesClient{
		internal: c.internal.WithClientName("core.DeploymentPipelinesClient"),
		endpoint: c.endpoint,
	}
}

// NewExternalDataSharesClient creates a new instance of ExternalDataSharesClient.
func (c *ClientFactory) NewExternalDataSharesClient() *ExternalDataSharesClient {
	return &ExternalDataSharesClient{
		internal: c.internal.WithClientName("core.ExternalDataSharesClient"),
		endpoint: c.endpoint,
	}
}

// NewGitClient creates a new instance of GitClient.
func (c *ClientFactory) NewGitClient() *GitClient {
	return &GitClient{
		internal: c.internal.WithClientName("core.GitClient"),
		endpoint: c.endpoint,
	}
}

// NewItemsClient creates a new instance of ItemsClient.
func (c *ClientFactory) NewItemsClient() *ItemsClient {
	return &ItemsClient{
		internal: c.internal.WithClientName("core.ItemsClient"),
		endpoint: c.endpoint,
	}
}

// NewJobSchedulerClient creates a new instance of JobSchedulerClient.
func (c *ClientFactory) NewJobSchedulerClient() *JobSchedulerClient {
	return &JobSchedulerClient{
		internal: c.internal.WithClientName("core.JobSchedulerClient"),
		endpoint: c.endpoint,
	}
}

// NewLongRunningOperationsClient creates a new instance of LongRunningOperationsClient.
func (c *ClientFactory) NewLongRunningOperationsClient() *LongRunningOperationsClient {
	return &LongRunningOperationsClient{
		internal: c.internal.WithClientName("core.LongRunningOperationsClient"),
		endpoint: c.endpoint,
	}
}

// NewManagedPrivateEndpointsClient creates a new instance of ManagedPrivateEndpointsClient.
func (c *ClientFactory) NewManagedPrivateEndpointsClient() *ManagedPrivateEndpointsClient {
	return &ManagedPrivateEndpointsClient{
		internal: c.internal.WithClientName("core.ManagedPrivateEndpointsClient"),
		endpoint: c.endpoint,
	}
}

// NewOneLakeDataAccessSecurityClient creates a new instance of OneLakeDataAccessSecurityClient.
func (c *ClientFactory) NewOneLakeDataAccessSecurityClient() *OneLakeDataAccessSecurityClient {
	return &OneLakeDataAccessSecurityClient{
		internal: c.internal.WithClientName("core.OneLakeDataAccessSecurityClient"),
		endpoint: c.endpoint,
	}
}

// NewOneLakeShortcutsClient creates a new instance of OneLakeShortcutsClient.
func (c *ClientFactory) NewOneLakeShortcutsClient() *OneLakeShortcutsClient {
	return &OneLakeShortcutsClient{
		internal: c.internal.WithClientName("core.OneLakeShortcutsClient"),
		endpoint: c.endpoint,
	}
}

// NewWorkspacesClient creates a new instance of WorkspacesClient.
func (c *ClientFactory) NewWorkspacesClient() *WorkspacesClient {
	return &WorkspacesClient{
		internal: c.internal.WithClientName("core.WorkspacesClient"),
		endpoint: c.endpoint,
	}
}
