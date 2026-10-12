package fabric

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"

	"github.com/microsoft/fabric-sdk-go/internal/iruntime"
)

// ServiceName identifies Microsoft Fabric in a cloud service configuration.
// Use it to add Fabric endpoint and audience values to a custom cloud.Configuration.
const ServiceName cloud.ServiceName = iruntime.ServiceName
