// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package core

// GitProviderDetailsClassification provides polymorphic access to related types.
// Call the interface's GetGitProviderDetails() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureDevOpsDetails, *GitHubDetails, *GitProviderDetails
type GitProviderDetailsClassification interface {
	// GetGitProviderDetails returns the GitProviderDetails content of the underlying type.
	GetGitProviderDetails() *GitProviderDetails
}

// ScheduleConfigClassification provides polymorphic access to related types.
// Call the interface's GetScheduleConfig() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CronScheduleConfig, *DailyScheduleConfig, *ScheduleConfig, *WeeklyScheduleConfig
type ScheduleConfigClassification interface {
	// GetScheduleConfig returns the ScheduleConfig content of the underlying type.
	GetScheduleConfig() *ScheduleConfig
}
