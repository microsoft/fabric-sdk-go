// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package iruntime

import (
	"errors"
	"net/url"
	"reflect"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

const (
	moduleVersion      = "0.1.0-beta.5"
	defaultApiEndpoint = "https://api.fabric.microsoft.com"
)

// ServiceClient is used to create any client in this module.
// It's a wrapper around azcore.Client and set default endpoint if not specified.
// Don't use this type directly, use NewServiceClient instead.
type ServiceClient struct {
	Internal *azcore.Client
	Endpoint string
}

// NewServiceClient creates a new instance of ServiceClient with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - endpoint - pass nil to accept the default values.
//   - options - pass nil to accept the default values.
func NewServiceClient(credential azcore.TokenCredential, endpoint *string, options *azcore.ClientOptions) (*ServiceClient, error) {
	apiEndpoint, err := getEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	if options == nil {
		options = &azcore.ClientOptions{}
	}

	if reflect.ValueOf(options.Cloud).IsZero() {
		options.Cloud = cloud.AzurePublic
	}

	authPolicy := runtime.NewBearerTokenPolicy(credential, []string{"https://api.fabric.microsoft.com/.default"}, nil)
	plOpts := runtime.PipelineOptions{
		PerRetry: []policy.Policy{authPolicy},
		Tracing: runtime.TracingOptions{
			Namespace: "Microsoft.Fabric",
		},
	}

	client, err := azcore.NewClient("fabric", "v"+moduleVersion, plOpts, options)
	if err != nil {
		return nil, err
	}

	return &ServiceClient{
		Internal: client,
		Endpoint: apiEndpoint,
	}, nil
}

func getEndpoint(endpoint *string) (string, error) {
	apiEndpoint := defaultApiEndpoint
	if endpoint != nil {
		apiEndpoint = *endpoint
	}
	apiEndpoint = strings.ToLower(strings.TrimRight(apiEndpoint, "/"))
	parsedApiEndpoint, err := url.Parse(apiEndpoint)
	if err != nil {
		return "", err
	}

	if !parsedApiEndpoint.IsAbs() {
		return "", errors.New("endpoint must be an absolute URL")
	}

	return parsedApiEndpoint.String(), nil
}
