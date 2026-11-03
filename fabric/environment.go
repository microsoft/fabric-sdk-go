package fabric

import (
	"github.com/microsoft/fabric-sdk-go/internal/iruntime"
)

// Environment identifies a predefined Microsoft Fabric environment.
// Only the constants defined in this package are supported.
type Environment = iruntime.Environment

const (
	// EnvironmentPublic is the default Public environment.
	EnvironmentPublic Environment = iruntime.EnvironmentPublic
	// EnvironmentGCCHigh is the US Government GCC High environment.
	EnvironmentGCCHigh Environment = iruntime.EnvironmentGCCHigh
)
