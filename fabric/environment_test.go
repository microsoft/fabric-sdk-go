package fabric_test

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

	"github.com/microsoft/fabric-sdk-go/fabric"
	"github.com/microsoft/fabric-sdk-go/fabric/core"
)

type recordingCredential struct {
	scopes []string
}

func (credential *recordingCredential) GetToken(_ context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
	credential.scopes = options.Scopes
	return azcore.AccessToken{Token: "token", ExpiresOn: time.Now().Add(time.Hour)}, nil
}

type recordingTransport struct {
	host string
}

func (transport *recordingTransport) Do(request *http.Request) (*http.Response, error) {
	transport.host = request.URL.Host
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(`{"value":[]}`)),
		Request:    request,
	}, nil
}

func TestClientFactoriesUseEnvironment(t *testing.T) {
	tests := []struct {
		name         string
		environment  fabric.Environment
		endpoint     string
		cloud        cloud.Configuration
		privateLinks bool
		wantHost     string
		wantScope    string
	}{
		{
			name:      "default Public",
			wantHost:  "api.fabric.microsoft.com",
			wantScope: "https://api.fabric.microsoft.com/.default",
		},
		{
			name:        "explicit Public",
			environment: fabric.EnvironmentPublic,
			wantHost:    "api.fabric.microsoft.com",
			wantScope:   "https://api.fabric.microsoft.com/.default",
		},
		{
			name:        "GCCHigh",
			environment: fabric.EnvironmentGCCHigh,
			wantHost:    "highapi.fabric.microsoft.us",
			wantScope:   "https://api.fabric.microsoft.us/.default",
		},
		{
			name:      "default Public with custom endpoint",
			endpoint:  "https://override.example.com/",
			wantHost:  "override.example.com",
			wantScope: "https://api.fabric.microsoft.com/.default",
		},
		{
			name:        "explicit Public with custom endpoint",
			environment: fabric.EnvironmentPublic,
			endpoint:    "https://override.example.com/",
			wantHost:    "override.example.com",
			wantScope:   "https://api.fabric.microsoft.com/.default",
		},
		{
			name:        "GCCHigh with custom endpoint",
			environment: fabric.EnvironmentGCCHigh,
			endpoint:    "https://override.example.com/",
			wantHost:    "override.example.com",
			wantScope:   "https://api.fabric.microsoft.us/.default",
		},
		{
			name:        "Azure cloud does not override Fabric environment",
			environment: fabric.EnvironmentGCCHigh,
			cloud:       cloud.AzurePublic,
			wantHost:    "highapi.fabric.microsoft.us",
			wantScope:   "https://api.fabric.microsoft.us/.default",
		},
		{
			name: "unrelated cloud service audience is not used",
			cloud: cloud.Configuration{Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
				"otherService": {Endpoint: "https://other.example.com", Audience: "https://other.example.com"},
			}},
			wantHost:  "api.fabric.microsoft.com",
			wantScope: "https://api.fabric.microsoft.com/.default",
		},
		{
			name:         "GCCHigh private links",
			environment:  fabric.EnvironmentGCCHigh,
			privateLinks: true,
			wantHost:     "12345678123412341234123456789abc.z12.w.highapi.fabric.microsoft.us",
			wantScope:    "https://api.fabric.microsoft.us/.default",
		},
		{
			name:         "Public private links with custom endpoint",
			endpoint:     "https://override.example.com",
			privateLinks: true,
			wantHost:     "12345678123412341234123456789abc.z12.w.override.example.com",
			wantScope:    "https://api.fabric.microsoft.com/.default",
		},
	}

	for _, test := range tests {
		for _, sharedClient := range []bool{false, true} {
			factoryName := "direct factory"
			if sharedClient {
				factoryName = "shared Fabric client"
			}
			t.Run(test.name+"/"+factoryName, func(t *testing.T) {
				credential := &recordingCredential{}
				transport := &recordingTransport{}
				options := &fabric.ClientOptions{
					ClientOptions: azcore.ClientOptions{
						Cloud:     test.cloud,
						Transport: transport,
					},
					Environment:              test.environment,
					UseWorkspacePrivateLinks: test.privateLinks,
				}
				var endpoint *string
				if test.endpoint != "" {
					endpoint = &test.endpoint
				}

				var factory *core.ClientFactory
				if sharedClient {
					client, err := fabric.NewClient(credential, endpoint, options)
					if err != nil {
						t.Fatal(err)
					}
					if client.UseWorkspacePrivateLinks != test.privateLinks {
						t.Error("client did not preserve private links setting")
					}
					factory = core.NewClientFactoryWithClient(*client)
				} else {
					var err error
					factory, err = core.NewClientFactory(credential, endpoint, options)
					if err != nil {
						t.Fatal(err)
					}
				}

				pager := factory.NewItemsClient().NewListItemsPager("12345678-1234-1234-1234-123456789abc", nil)
				if _, err := pager.NextPage(context.Background()); err != nil {
					t.Fatal(err)
				}
				if transport.host != test.wantHost {
					t.Errorf("request host = %q, want %q", transport.host, test.wantHost)
				}
				if len(credential.scopes) != 1 || credential.scopes[0] != test.wantScope {
					t.Errorf("requested scopes = %q, want [%q]", credential.scopes, test.wantScope)
				}
			})
		}
	}
}

func TestClientFactoriesRejectInvalidEnvironment(t *testing.T) {
	tests := []struct {
		name        string
		environment fabric.Environment
	}{
		{name: "unsupported environment value", environment: fabric.Environment(2)},
		{name: "unknown environment", environment: fabric.Environment(99)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			options := &fabric.ClientOptions{Environment: test.environment}
			if client, err := fabric.NewClient(nil, nil, options); err == nil || client != nil {
				t.Fatalf("fabric.NewClient() = %v, %v; want nil client and error", client, err)
			}
			if factory, err := core.NewClientFactory(nil, nil, options); err == nil || factory != nil {
				t.Fatalf("core.NewClientFactory() = %v, %v; want nil factory and error", factory, err)
			}
		})
	}
}
