package iruntime

import (
	"fmt"
)

// Environment identifies a predefined Microsoft Fabric environment.
type Environment int

const (
	EnvironmentPublic Environment = iota
	EnvironmentGCCHigh
)

type environmentConfiguration struct {
	endpoint string
	audience string
}

func resolveEnvironment(environment Environment) (environmentConfiguration, error) {
	switch environment {
	case EnvironmentPublic:
		return environmentConfiguration{
			endpoint: "https://api.fabric.microsoft.com",
			audience: "https://api.fabric.microsoft.com",
		}, nil
	case EnvironmentGCCHigh:
		return environmentConfiguration{
			endpoint: "https://highapi.fabric.microsoft.us",
			audience: "https://api.fabric.microsoft.us",
		}, nil
	default:
		return environmentConfiguration{}, fmt.Errorf("unsupported Fabric environment %d", environment)
	}
}
