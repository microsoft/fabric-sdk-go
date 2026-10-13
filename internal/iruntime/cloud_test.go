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
)

func TestResolveCloud(t *testing.T) {
	tests := []struct {
		name          string
		configuration cloud.Configuration
		endpoint      string
		audience      string
	}{
		{name: "zero configuration", endpoint: "https://api.fabric.microsoft.com", audience: "https://api.fabric.microsoft.com"},
		{
			name: "custom",
			configuration: cloud.Configuration{
				ActiveDirectoryAuthorityHost: "https://login.example.com",
				Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
					ServiceName: {Endpoint: "https://api.fabric.example.com", Audience: "https://fabric.example.com"},
				},
			},
			endpoint: "https://api.fabric.example.com",
			audience: "https://fabric.example.com",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			configuration, service, err := resolveCloud(test.configuration)
			if err != nil {
				t.Fatal(err)
			}
			if service.Endpoint != test.endpoint {
				t.Errorf("endpoint = %q, want %q", service.Endpoint, test.endpoint)
			}
			if service.Audience != test.audience {
				t.Errorf("audience = %q, want %q", service.Audience, test.audience)
			}
			if configuration.ActiveDirectoryAuthorityHost == "" {
				t.Error("resolved configuration has no authority host")
			}
		})
	}
}

func TestResolveCloudErrors(t *testing.T) {
	tests := []struct {
		name          string
		configuration cloud.Configuration
	}{
		{
			name:          "built-in cloud without Fabric service",
			configuration: cloud.AzureGovernment,
		},
		{
			name:          "missing service",
			configuration: cloud.Configuration{Services: map[cloud.ServiceName]cloud.ServiceConfiguration{}},
		},
		{
			name: "missing endpoint",
			configuration: cloud.Configuration{Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
				ServiceName: {Audience: "https://fabric.example.com"},
			}},
		},
		{
			name: "missing audience",
			configuration: cloud.Configuration{Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
				ServiceName: {Endpoint: "https://api.fabric.example.com"},
			}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, _, err := resolveCloud(test.configuration); err == nil {
				t.Fatal("resolveCloud() error = nil")
			}
		})
	}
}

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
		Header:     http.Header{},
		Body:       io.NopCloser(strings.NewReader("")),
		Request:    request,
	}, nil
}

func TestServiceClientUsesCloudEndpointAndAudience(t *testing.T) {
	tests := []struct {
		name          string
		configuration cloud.Configuration
		endpoint      *string
		wantHost      string
		wantScope     string
		wantAuthority string
	}{
		{
			name:          "default Public cloud",
			wantHost:      "api.fabric.microsoft.com",
			wantScope:     "https://api.fabric.microsoft.com/.default",
			wantAuthority: cloud.AzurePublic.ActiveDirectoryAuthorityHost,
		},
		{
			name: "custom with explicit endpoint",
			configuration: cloud.Configuration{
				ActiveDirectoryAuthorityHost: "https://login.example.com",
				Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
					ServiceName: {Endpoint: "https://cloud.example.com", Audience: "https://audience.example.com/"},
				},
			},
			endpoint:      stringPointer("https://override.example.com/"),
			wantHost:      "override.example.com",
			wantScope:     "https://audience.example.com/.default",
			wantAuthority: "https://login.example.com",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			credential := &recordingCredential{}
			transport := &recordingTransport{}
			options := &ClientOptions{ClientOptions: azcore.ClientOptions{
				Cloud:     test.configuration,
				Transport: transport,
			}}

			client, err := NewServiceClient(credential, "1.0.0", test.endpoint, options)
			if err != nil {
				t.Fatal(err)
			}
			request, err := runtime.NewRequest(context.Background(), http.MethodGet, client.Endpoint+"/v1/workspaces")
			if err != nil {
				t.Fatal(err)
			}
			if _, err = client.Internal.Pipeline().Do(request); err != nil {
				t.Fatal(err)
			}

			if transport.host != test.wantHost {
				t.Errorf("request host = %q, want %q", transport.host, test.wantHost)
			}
			if len(credential.scopes) != 1 || credential.scopes[0] != test.wantScope {
				t.Errorf("requested scopes = %q, want [%q]", credential.scopes, test.wantScope)
			}
			if options.Cloud.ActiveDirectoryAuthorityHost != test.wantAuthority {
				t.Errorf("authority host = %q, want %q", options.Cloud.ActiveDirectoryAuthorityHost, test.wantAuthority)
			}
		})
	}
}

func stringPointer(value string) *string {
	return &value
}
