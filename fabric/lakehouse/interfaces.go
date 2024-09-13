// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package lakehouse

// FileFormatOptionsClassification provides polymorphic access to related types.
// Call the interface's GetFileFormatOptions() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CSV, *FileFormatOptions, *Parquet
type FileFormatOptionsClassification interface {
	// GetFileFormatOptions returns the FileFormatOptions content of the underlying type.
	GetFileFormatOptions() *FileFormatOptions
}