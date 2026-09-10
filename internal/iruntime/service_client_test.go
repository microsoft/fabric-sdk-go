// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// SPDX-License-Identifier: MIT

package iruntime

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

type recordingCredential struct {
	scopes []string
}

func (c *recordingCredential) GetToken(_ context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
	c.scopes = append([]string(nil), options.Scopes...)
	return azcore.AccessToken{Token: "test-token", ExpiresOn: time.Now().Add(time.Hour)}, nil
}

type clientTransport func(*http.Request) (*http.Response, error)

func (f clientTransport) Do(req *http.Request) (*http.Response, error) { return f(req) }

func TestServiceClientCloudConfiguration(t *testing.T) {
	const serviceName cloud.ServiceName = "Microsoft.Fabric"
	const workspacePath = "/v1/workspaces/12345678-1234-1234-1234-123456789abc/items"
	const privatePrefix = "12345678123412341234123456789abc.z12.w."
	customCloud := cloud.Configuration{
		ActiveDirectoryAuthorityHost: cloud.AzureGovernment.ActiveDirectoryAuthorityHost,
		Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
			serviceName: {Endpoint: "https://fabric.example", Audience: "https://fabric-resource.example"},
		},
	}
	endpointOnly := "https://proxy.example"
	for _, tt := range []struct {
		name         string
		cloud        cloud.Configuration
		endpoint     *string
		privateLinks bool
		wantEndpoint string
		wantHost     string
		wantScope    string
	}{
		{"defaults", cloud.Configuration{}, nil, false, defaultApiEndpoint, "api.fabric.microsoft.com", defaultApiEndpoint + "/.default"},
		{"public cloud", cloud.AzurePublic, nil, false, defaultApiEndpoint, "api.fabric.microsoft.com", defaultApiEndpoint + "/.default"},
		{"endpoint override retains audience", cloud.AzurePublic, &endpointOnly, false, endpointOnly, "proxy.example", defaultApiEndpoint + "/.default"},
		{"configured cloud", customCloud, nil, false, "https://fabric.example", "fabric.example", "https://fabric-resource.example/.default"},
		{"explicit endpoint wins", customCloud, &endpointOnly, false, endpointOnly, "proxy.example", "https://fabric-resource.example/.default"},
		{"public private links", cloud.AzurePublic, nil, true, defaultApiEndpoint, privatePrefix + "api.fabric.microsoft.com", defaultApiEndpoint + "/.default"},
		{"configured private links", customCloud, nil, true, "https://fabric.example", privatePrefix + "fabric.example", "https://fabric-resource.example/.default"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			credential := &recordingCredential{}
			requests := 0
			options := &ClientOptions{
				ClientOptions: azcore.ClientOptions{
					Cloud: tt.cloud,
					Transport: clientTransport(func(req *http.Request) (*http.Response, error) {
						requests++
						require.Equal(t, tt.wantHost, req.URL.Host)
						require.Equal(t, workspacePath, req.URL.Path)
						require.Equal(t, "Bearer test-token", req.Header.Get("Authorization"))
						return &http.Response{StatusCode: http.StatusOK, Header: http.Header{}, Body: io.NopCloser(strings.NewReader("{}")), Request: req}, nil
					}),
				},
				UseWorkspacePrivateLinks: tt.privateLinks,
			}
			client, err := NewServiceClient(credential, "0.20.0", tt.endpoint, options)
			require.NoError(t, err)
			require.Equal(t, tt.wantEndpoint, client.Endpoint)
			req, err := runtime.NewRequest(context.Background(), http.MethodGet, client.Endpoint+workspacePath)
			require.NoError(t, err)
			resp, err := client.Internal.Pipeline().Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()
			require.Equal(t, 1, requests)
			require.Equal(t, []string{tt.wantScope}, credential.scopes)
		})
	}
}

func TestServiceClientRequiresSovereignAudience(t *testing.T) {
	for _, config := range []cloud.Configuration{cloud.AzureGovernment, cloud.AzureChina} {
		t.Run(config.ActiveDirectoryAuthorityHost, func(t *testing.T) {
			credential := &recordingCredential{}
			client, err := NewServiceClient(credential, "0.20.0", nil, &ClientOptions{
				ClientOptions: azcore.ClientOptions{Cloud: config},
			})
			require.ErrorContains(t, err, "Microsoft.Fabric")
			require.Nil(t, client)
			require.Empty(t, credential.scopes)
		})
	}
}
