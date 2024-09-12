// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package environment

import "time"

// ComponentPublishInfo - Publish info for each components in environment.
type ComponentPublishInfo struct {
	// Spark libraries publish information.
	SparkLibraries *SparkLibraries

	// Spark settings publish information.
	SparkSettings *SparkSettings
}

// CreateEnvironmentRequest - Create environment request payload.
type CreateEnvironmentRequest struct {
	// REQUIRED; The environment display name.
	DisplayName *string

	// The environment description. Maximum length is 256 characters.
	Description *string
}

// CustomLibraries - Custom libraries.
type CustomLibraries struct {
	// A list of Jar files.
	JarFiles []string

	// A list of Python files.
	PyFiles []string

	// A list of R files.
	RTarFiles []string

	// A list of Wheel files.
	WheelFiles []string
}

// DynamicExecutorAllocationProperties - Dynamic executor allocation proerties.
type DynamicExecutorAllocationProperties struct {
	// REQUIRED; The status of the dynamic executor allocation. False - Disabled, true - Enabled.
	Enabled *bool

	// REQUIRED; The maximum executor number for dynamic allocation and the minimum for this property is 1
	MaxExecutors *int32

	// REQUIRED; The minimum executor number for dynamic allocation and the minimum for this property is 1.
	MinExecutors *int32
}

// Environment - An Environment item.
type Environment struct {
	// REQUIRED; The item type.
	Type *ItemType

	// The item description.
	Description *string

	// The item display name.
	DisplayName *string

	// The environment properties.
	Properties *PublishInfo

	// READ-ONLY; The item ID.
	ID *string

	// READ-ONLY; The workspace ID.
	WorkspaceID *string
}

// Environments - A list of environments.
type Environments struct {
	// REQUIRED; A list of environments.
	Value []Environment

	// The token for the next result set batch. If there are no more records, it's removed from the response.
	ContinuationToken *string

	// The URI of the next result set batch. If there are no more records, it's removed from the response.
	ContinuationURI *string
}

type InstancePool struct {
	// REQUIRED; Instance pool name.
	Name *string

	// REQUIRED; Instance pool type.
	Type *CustomPoolType

	// Instance pool ID.
	ID *string
}

// Libraries - Environment libraries.
type Libraries struct {
	// Custom libraries (.py, .whl, .jar, .tar.gz).
	CustomLibraries *CustomLibraries

	// Feed libraries.
	EnvironmentYml *string
}

// PublishDetails - Details of publish operation.
type PublishDetails struct {
	// Environment component publish information.
	ComponentPublishInfo *ComponentPublishInfo

	// End time of publish operation.
	EndTime *time.Time

	// Start time of publish operation.
	StartTime *time.Time

	// Publish state. Additional state types may be added over time.
	State *PublishState

	// Target verion to be published.
	TargetVersion *string
}

// PublishInfo - Environment publish information.
type PublishInfo struct {
	// REQUIRED; Environment publish operation details.
	PublishDetails *PublishDetails
}

type SparkCompute struct {
	// Spark driver core.
	DriverCores *int32

	// Spark driver memory.
	DriverMemory *string

	// Dynamic executor allocation.
	DynamicExecutorAllocation *DynamicExecutorAllocationProperties

	// Spark executor core.
	ExecutorCores *int32

	// Spark executor memory.
	ExecutorMemory *string

	// Environment pool has to be a valid custom pool. "Starter Pool" means use starter pool.
	InstancePool *InstancePool

	// Runtime version, find the supported fabric runtimes [/fabric/data-engineering/runtime].
	RuntimeVersion *string

	// Spark properties.
	SparkProperties any
}

// SparkLibraries - Spark libraries.
type SparkLibraries struct {
	// Publish state. Additional state types may be added over time.
	State *PublishState
}

// SparkSettings - Spark settings.
type SparkSettings struct {
	// Publish state. Additional state types may be added over time.
	State *PublishState
}

// UpdateEnvironmentRequest - Update environment request.
type UpdateEnvironmentRequest struct {
	// The environment description. Maximum length is 256 characters.
	Description *string

	// The environment display name.
	DisplayName *string
}

type UpdateEnvironmentSparkComputeRequest struct {
	// Spark driver core.
	DriverCores *int32

	// Spark driver memory.
	DriverMemory *string

	// Dynamic executor allocation.
	DynamicExecutorAllocation *DynamicExecutorAllocationProperties

	// Spark executor core.
	ExecutorCores *int32

	// Spark executor memory.
	ExecutorMemory *string

	// Environment pool has to be a valid custom pool. The name for a default workspace pool is Starter Pool.
	InstancePool *InstancePool

	// Runtime version, find the supported fabric runtimes [/fabric/data-engineering/runtime].
	RuntimeVersion *string

	// Spark properties.
	SparkProperties any
}
