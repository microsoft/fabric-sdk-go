package iruntime

import (
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
)

func TestResolveEnvironment(t *testing.T) {
	tests := []struct {
		name        string
		environment Environment
		endpoint    string
		audience    string
	}{
		{
			name:     "zero value defaults to Public",
			endpoint: "https://api.fabric.microsoft.com",
			audience: "https://api.fabric.microsoft.com",
		},
		{
			name:        "Public",
			environment: EnvironmentPublic,
			endpoint:    "https://api.fabric.microsoft.com",
			audience:    "https://api.fabric.microsoft.com",
		},
		{
			name:        "GCCHigh",
			environment: EnvironmentGCCHigh,
			endpoint:    "https://highapi.fabric.microsoft.us",
			audience:    "https://api.fabric.microsoft.us",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			configuration, err := resolveEnvironment(test.environment)
			if err != nil {
				t.Fatal(err)
			}
			if configuration.endpoint != test.endpoint {
				t.Errorf("endpoint = %q, want %q", configuration.endpoint, test.endpoint)
			}
			if configuration.audience != test.audience {
				t.Errorf("audience = %q, want %q", configuration.audience, test.audience)
			}
		})
	}
}

func TestServiceClientDefaultEnvironment(t *testing.T) {
	client, err := NewServiceClient(nil, "1.0.0", nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if client.Endpoint != "https://api.fabric.microsoft.com" {
		t.Errorf("endpoint = %q", client.Endpoint)
	}
}

func TestServiceClientRejectsInvalidConfiguration(t *testing.T) {
	tests := []struct {
		name    string
		options ClientOptions
		error   string
	}{
		{
			name:    "unsupported environment value",
			options: ClientOptions{Environment: Environment(2)},
			error:   "unsupported Fabric environment",
		},
		{
			name:    "unknown environment",
			options: ClientOptions{Environment: Environment(99)},
			error:   "unsupported Fabric environment",
		},
		{
			name:    "negative environment",
			options: ClientOptions{Environment: Environment(-1)},
			error:   "unsupported Fabric environment",
		},
		{
			name: "custom Fabric audience",
			options: ClientOptions{ClientOptions: azcore.ClientOptions{
				Cloud: cloud.Configuration{Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
					"microsoftFabric": {Audience: "https://custom.example.com"},
				}},
			}},
			error: "Fabric service configuration",
		},
		{
			name: "custom Fabric endpoint through Cloud",
			options: ClientOptions{ClientOptions: azcore.ClientOptions{
				Cloud: cloud.Configuration{Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
					"microsoftFabric": {Endpoint: "https://custom.example.com"},
				}},
			}},
			error: "Fabric service configuration",
		},
		{
			name: "empty Fabric service",
			options: ClientOptions{ClientOptions: azcore.ClientOptions{
				Cloud: cloud.Configuration{Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
					"microsoftFabric": {},
				}},
			}},
			error: "Fabric service configuration",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			endpoint := "https://override.example.com"
			client, err := NewServiceClient(nil, "1.0.0", &endpoint, &test.options)
			if err == nil || !strings.Contains(err.Error(), test.error) {
				t.Fatalf("NewServiceClient() error = %v, want %q", err, test.error)
			}
			if client != nil {
				t.Fatal("NewServiceClient() returned a client for invalid configuration")
			}
		})
	}
}

func TestServiceClientRejectsInvalidEndpoint(t *testing.T) {
	for _, endpoint := range []string{"", "/relative", "https://%invalid"} {
		t.Run(endpoint, func(t *testing.T) {
			if _, err := NewServiceClient(nil, "1.0.0", &endpoint, nil); err == nil {
				t.Fatal("NewServiceClient() error = nil")
			}
		})
	}
}

func TestServiceClientRejectsInvalidVersion(t *testing.T) {
	client, err := NewServiceClient(nil, "invalid", nil, nil)
	if err == nil || client != nil {
		t.Fatalf("NewServiceClient() = %v, %v; want nil client and error", client, err)
	}
}
