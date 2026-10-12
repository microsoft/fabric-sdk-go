package iruntime

import (
	"fmt"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
)

// ServiceName identifies Microsoft Fabric in a cloud service configuration.
const ServiceName cloud.ServiceName = "microsoftFabric"

const defaultEndpoint = "https://api.fabric.microsoft.com"
const defaultAudience = "https://api.fabric.microsoft.com"

func resolveCloud(configuration cloud.Configuration) (cloud.Configuration, cloud.ServiceConfiguration, error) {
	if reflect.ValueOf(configuration).IsZero() {
		return cloud.AzurePublic, cloud.ServiceConfiguration{
			Endpoint: defaultEndpoint,
			Audience: defaultAudience,
		}, nil
	}

	service, ok := configuration.Services[ServiceName]
	if !ok {
		return configuration, cloud.ServiceConfiguration{}, fmt.Errorf("cloud configuration has no %q service", ServiceName)
	}
	if service.Endpoint == "" {
		return configuration, cloud.ServiceConfiguration{}, fmt.Errorf("cloud configuration service %q has no endpoint", ServiceName)
	}
	if service.Audience == "" {
		return configuration, cloud.ServiceConfiguration{}, fmt.Errorf("cloud configuration service %q has no audience", ServiceName)
	}

	return configuration, service, nil
}
