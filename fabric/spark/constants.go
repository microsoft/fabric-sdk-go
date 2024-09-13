// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package spark

// CustomPoolType - Custom pool type. Additional CustomPoolType types may be added over time.
type CustomPoolType string

const (
	// CustomPoolTypeCapacity - Capacity level custom pool
	CustomPoolTypeCapacity CustomPoolType = "Capacity"
	// CustomPoolTypeWorkspace - Workspace level custom pool
	CustomPoolTypeWorkspace CustomPoolType = "Workspace"
)

// PossibleCustomPoolTypeValues returns the possible values for the CustomPoolType const type.
func PossibleCustomPoolTypeValues() []CustomPoolType {
	return []CustomPoolType{
		CustomPoolTypeCapacity,
		CustomPoolTypeWorkspace,
	}
}

// NodeFamily - Node family. Additional NodeFamily types may be added over time.
type NodeFamily string

const (
	// NodeFamilyMemoryOptimized - Memory optimized
	NodeFamilyMemoryOptimized NodeFamily = "MemoryOptimized"
)

// PossibleNodeFamilyValues returns the possible values for the NodeFamily const type.
func PossibleNodeFamilyValues() []NodeFamily {
	return []NodeFamily{
		NodeFamilyMemoryOptimized,
	}
}

// NodeSize - Node size [/fabric/data-engineering/spark-compute#node-sizes]. Additional NodeSize types may be added over time.
type NodeSize string

const (
	// NodeSizeLarge - Large node size
	NodeSizeLarge NodeSize = "Large"
	// NodeSizeMedium - Medium node size
	NodeSizeMedium NodeSize = "Medium"
	// NodeSizeSmall - Small node size
	NodeSizeSmall NodeSize = "Small"
	// NodeSizeXLarge - XLarge node size
	NodeSizeXLarge NodeSize = "XLarge"
	// NodeSizeXXLarge - XXLarge node size
	NodeSizeXXLarge NodeSize = "XXLarge"
)

// PossibleNodeSizeValues returns the possible values for the NodeSize const type.
func PossibleNodeSizeValues() []NodeSize {
	return []NodeSize{
		NodeSizeLarge,
		NodeSizeMedium,
		NodeSizeSmall,
		NodeSizeXLarge,
		NodeSizeXXLarge,
	}
}