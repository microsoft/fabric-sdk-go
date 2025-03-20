// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package spark

// AutoScaleProperties - Autoscale properties.
type AutoScaleProperties struct {
	// REQUIRED; The status of the auto scale. False - Disabled, true - Enabled.
	Enabled *bool

	// REQUIRED; The maximum node count.
	MaxNodeCount *int32

	// REQUIRED; The minimum node count.
	MinNodeCount *int32
}

// AutomaticLogProperties - Automatic Log Properties.
type AutomaticLogProperties struct {
	// REQUIRED; The status of the automatic log. False - Disabled, true - Enabled.
	Enabled *bool
}

// CreateCustomPoolRequest - Create custom pool request payload.
type CreateCustomPoolRequest struct {
	// REQUIRED; Autoscale.
	AutoScale *AutoScaleProperties

	// REQUIRED; Dynamic executor allocation.
	DynamicExecutorAllocation *DynamicExecutorAllocationProperties

	// REQUIRED; Custom pool name.
	// The name must be between 1 and 64 characters long and must contain only letters, numbers, dashes, underscores and spaces.
	// Custom pool names must be unique within the workspace.
	// "Starter Pool" is a reserved custom pool name.
	Name *string

	// REQUIRED; Node family.
	NodeFamily *NodeFamily

	// REQUIRED; Node size.
	NodeSize *NodeSize
}

// CustomPool - Custom pool.
type CustomPool struct {
	// Autoscale.
	AutoScale *AutoScaleProperties

	// Dynamic executor allocation.
	DynamicExecutorAllocation *DynamicExecutorAllocationProperties

	// Custom pool ID.
	ID *string

	// Custom pool name.
	Name *string

	// Node family.
	NodeFamily *NodeFamily

	// Node size.
	NodeSize *NodeSize

	// Custom pool type.
	Type *CustomPoolType
}

type CustomPools struct {
	// REQUIRED; A list of custom pools.
	Value []CustomPool

	// The token for the next result set batch. If there are no more records, it's removed from the response.
	ContinuationToken *string

	// The URI of the next result set batch. If there are no more records, it's removed from the response.
	ContinuationURI *string
}

// DynamicExecutorAllocationProperties - Dynamic executor allocation proerties.
type DynamicExecutorAllocationProperties struct {
	// REQUIRED; The status of the dynamic executor allocation. False - Disabled, true - Enabled.
	Enabled *bool

	// REQUIRED; The maximum executors.
	MaxExecutors *int32

	// REQUIRED; The minimum executors.
	MinExecutors *int32
}

type EnvironmentProperties struct {
	// The name of the default environment. Empty string indicated there is no workspace default environment.
	Name *string

	// Runtime [/fabric/data-engineering/runtime] version. For example: 1.3
	RuntimeVersion *string
}

// HighConcurrencyProperties - High Concurrency Properties.
type HighConcurrencyProperties struct {
	// The status of the high concurrency for notebook interactive run. False - Disabled, true - Enabled.
	NotebookInteractiveRunEnabled *bool

	// The status of the high concurrency for notebook pipeline run. False - Disabled, true - Enabled.
	NotebookPipelineRunEnabled *bool
}

type InstancePool struct {
	// Instance pool ID.
	ID *string

	// Instance pool name.
	Name *string

	// Instance pool type.
	Type *CustomPoolType
}

type JobsProperties struct {
	// Reserve maximum cores for active Spark jobs. When this setting is enabled, your Fabric capacity reserves the maximum number
	// of cores needed for active Spark jobs, ensuring job reliability by making
	// sure that cores are available if a job scales up. When this setting is disabled, jobs are started based on the minimum
	// number of cores needed, letting more jobs run at the same time. For more
	// information see job admission and management [/fabric/data-engineering/job-admission-management]. False - Disabled, true
	// - Enabled.
	ConservativeJobAdmissionEnabled *bool

	// Time to terminate inactive Spark sessions. The maximum is 14 days.
	SessionTimeoutInMinutes *int32
}

type PoolProperties struct {
	// Customize compute configurations for items. False - Disabled, true - Enabled.
	CustomizeComputeEnabled *bool

	// Default pool for workspace. It should be a valid custom pool name. "Starter Pool" means use starter pool.
	DefaultPool *InstancePool

	// Customize starter pool. For more information about configuring starter pool, see configuring starter pool [/fabric/data-engineering/configure-starter-pools].
	StarterPool *StarterPoolProperties
}

// StarterPoolProperties - Custom starter pool.
type StarterPoolProperties struct {
	// The maximum executors count.
	MaxExecutors *int32

	// The maximum node count.
	MaxNodeCount *int32
}

// UpdateCustomPoolRequest - Update custom pool request payload.
type UpdateCustomPoolRequest struct {
	// Autoscale.
	AutoScale *AutoScaleProperties

	// Dynamic executor allocation.
	DynamicExecutorAllocation *DynamicExecutorAllocationProperties

	// Custom pool name.
	// The name must be between 1 and 64 characters long and must contain only letters, numbers, dashes, underscores and spaces.
	// Custom pool names must be unique within the workspace.
	// "Starter Pool" is a reserved custom pool name.
	Name *string

	// Node family.
	NodeFamily *NodeFamily

	// Node size.
	NodeSize *NodeSize
}

// UpdateWorkspaceSparkSettingsRequest - Update workspace Spark settings request payload.
type UpdateWorkspaceSparkSettingsRequest struct {
	// Automatic log settings.
	AutomaticLog *AutomaticLogProperties

	// Environment settings.
	Environment *EnvironmentProperties

	// High concurrency settings.
	HighConcurrency *HighConcurrencyProperties

	// Job management settings.
	Job *JobsProperties

	// Pool settings.
	Pool *PoolProperties
}

// WorkspaceSparkSettings - Workspace Spark settings.
type WorkspaceSparkSettings struct {
	// Automatic log settings.
	AutomaticLog *AutomaticLogProperties

	// Environment settings.
	Environment *EnvironmentProperties

	// High concurrency settings.
	HighConcurrency *HighConcurrencyProperties

	// Job management settings.
	Job *JobsProperties

	// Pool settings.
	Pool *PoolProperties
}
