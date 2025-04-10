// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package sparkjobdefinition

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// MarshalJSON implements the json.Marshaller interface for type CreateSparkJobDefinitionRequest.
func (c CreateSparkJobDefinitionRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "definition", c.Definition)
	populate(objectMap, "description", c.Description)
	populate(objectMap, "displayName", c.DisplayName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CreateSparkJobDefinitionRequest.
func (c *CreateSparkJobDefinitionRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "definition":
			err = unpopulate(val, "Definition", &c.Definition)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &c.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &c.DisplayName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Duration.
func (d Duration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "timeUnit", d.TimeUnit)
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Duration.
func (d *Duration) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "timeUnit":
			err = unpopulate(val, "TimeUnit", &d.TimeUnit)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &d.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ExecutionData.
func (e ExecutionData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "additionalLibraryUris", e.AdditionalLibraryUris)
	populate(objectMap, "commandLineArguments", e.CommandLineArguments)
	populate(objectMap, "executableFile", e.ExecutableFile)
	populate(objectMap, "mainClass", e.MainClass)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ExecutionData.
func (e *ExecutionData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "additionalLibraryUris":
			err = unpopulate(val, "AdditionalLibraryUris", &e.AdditionalLibraryUris)
			delete(rawMsg, key)
		case "commandLineArguments":
			err = unpopulate(val, "CommandLineArguments", &e.CommandLineArguments)
			delete(rawMsg, key)
		case "executableFile":
			err = unpopulate(val, "ExecutableFile", &e.ExecutableFile)
			delete(rawMsg, key)
		case "mainClass":
			err = unpopulate(val, "MainClass", &e.MainClass)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ItemReference.
func (i ItemReference) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	objectMap["referenceType"] = i.ReferenceType
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ItemReference.
func (i *ItemReference) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "referenceType":
			err = unpopulate(val, "ReferenceType", &i.ReferenceType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ItemReferenceByID.
func (i ItemReferenceByID) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "itemId", i.ItemID)
	objectMap["referenceType"] = ItemReferenceTypeByID
	populate(objectMap, "workspaceId", i.WorkspaceID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ItemReferenceByID.
func (i *ItemReferenceByID) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "itemId":
			err = unpopulate(val, "ItemID", &i.ItemID)
			delete(rawMsg, key)
		case "referenceType":
			err = unpopulate(val, "ReferenceType", &i.ReferenceType)
			delete(rawMsg, key)
		case "workspaceId":
			err = unpopulate(val, "WorkspaceID", &i.WorkspaceID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LivySession.
func (l LivySession) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "attemptNumber", l.AttemptNumber)
	populate(objectMap, "cancellationReason", l.CancellationReason)
	populate(objectMap, "capacityId", l.CapacityID)
	populate(objectMap, "consumerId", l.ConsumerID)
	populate(objectMap, "creatorItem", l.CreatorItem)
	populateDateTimeRFC3339(objectMap, "endDateTime", l.EndDateTime)
	populate(objectMap, "isHighConcurrency", l.IsHighConcurrency)
	populate(objectMap, "item", l.Item)
	populate(objectMap, "itemName", l.ItemName)
	populate(objectMap, "itemType", l.ItemType)
	populate(objectMap, "jobInstanceId", l.JobInstanceID)
	populate(objectMap, "jobType", l.JobType)
	populate(objectMap, "livyId", l.LivyID)
	populate(objectMap, "livyName", l.LivyName)
	populate(objectMap, "livySessionItemResourceUri", l.LivySessionItemResourceURI)
	populate(objectMap, "maxNumberOfAttempts", l.MaxNumberOfAttempts)
	populate(objectMap, "operationName", l.OperationName)
	populate(objectMap, "origin", l.Origin)
	populate(objectMap, "queuedDuration", l.QueuedDuration)
	populate(objectMap, "runningDuration", l.RunningDuration)
	populate(objectMap, "runtimeVersion", l.RuntimeVersion)
	populate(objectMap, "sparkApplicationId", l.SparkApplicationID)
	populateDateTimeRFC3339(objectMap, "startDateTime", l.StartDateTime)
	populate(objectMap, "state", l.State)
	populateDateTimeRFC3339(objectMap, "submittedDateTime", l.SubmittedDateTime)
	populate(objectMap, "submitter", l.Submitter)
	populate(objectMap, "totalDuration", l.TotalDuration)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LivySession.
func (l *LivySession) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "attemptNumber":
			err = unpopulate(val, "AttemptNumber", &l.AttemptNumber)
			delete(rawMsg, key)
		case "cancellationReason":
			err = unpopulate(val, "CancellationReason", &l.CancellationReason)
			delete(rawMsg, key)
		case "capacityId":
			err = unpopulate(val, "CapacityID", &l.CapacityID)
			delete(rawMsg, key)
		case "consumerId":
			err = unpopulate(val, "ConsumerID", &l.ConsumerID)
			delete(rawMsg, key)
		case "creatorItem":
			err = unpopulate(val, "CreatorItem", &l.CreatorItem)
			delete(rawMsg, key)
		case "endDateTime":
			err = unpopulateDateTimeRFC3339(val, "EndDateTime", &l.EndDateTime)
			delete(rawMsg, key)
		case "isHighConcurrency":
			err = unpopulate(val, "IsHighConcurrency", &l.IsHighConcurrency)
			delete(rawMsg, key)
		case "item":
			err = unpopulate(val, "Item", &l.Item)
			delete(rawMsg, key)
		case "itemName":
			err = unpopulate(val, "ItemName", &l.ItemName)
			delete(rawMsg, key)
		case "itemType":
			err = unpopulate(val, "ItemType", &l.ItemType)
			delete(rawMsg, key)
		case "jobInstanceId":
			err = unpopulate(val, "JobInstanceID", &l.JobInstanceID)
			delete(rawMsg, key)
		case "jobType":
			err = unpopulate(val, "JobType", &l.JobType)
			delete(rawMsg, key)
		case "livyId":
			err = unpopulate(val, "LivyID", &l.LivyID)
			delete(rawMsg, key)
		case "livyName":
			err = unpopulate(val, "LivyName", &l.LivyName)
			delete(rawMsg, key)
		case "livySessionItemResourceUri":
			err = unpopulate(val, "LivySessionItemResourceURI", &l.LivySessionItemResourceURI)
			delete(rawMsg, key)
		case "maxNumberOfAttempts":
			err = unpopulate(val, "MaxNumberOfAttempts", &l.MaxNumberOfAttempts)
			delete(rawMsg, key)
		case "operationName":
			err = unpopulate(val, "OperationName", &l.OperationName)
			delete(rawMsg, key)
		case "origin":
			err = unpopulate(val, "Origin", &l.Origin)
			delete(rawMsg, key)
		case "queuedDuration":
			err = unpopulate(val, "QueuedDuration", &l.QueuedDuration)
			delete(rawMsg, key)
		case "runningDuration":
			err = unpopulate(val, "RunningDuration", &l.RunningDuration)
			delete(rawMsg, key)
		case "runtimeVersion":
			err = unpopulate(val, "RuntimeVersion", &l.RuntimeVersion)
			delete(rawMsg, key)
		case "sparkApplicationId":
			err = unpopulate(val, "SparkApplicationID", &l.SparkApplicationID)
			delete(rawMsg, key)
		case "startDateTime":
			err = unpopulateDateTimeRFC3339(val, "StartDateTime", &l.StartDateTime)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, "State", &l.State)
			delete(rawMsg, key)
		case "submittedDateTime":
			err = unpopulateDateTimeRFC3339(val, "SubmittedDateTime", &l.SubmittedDateTime)
			delete(rawMsg, key)
		case "submitter":
			err = unpopulate(val, "Submitter", &l.Submitter)
			delete(rawMsg, key)
		case "totalDuration":
			err = unpopulate(val, "TotalDuration", &l.TotalDuration)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LivySessions.
func (l LivySessions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "continuationToken", l.ContinuationToken)
	populate(objectMap, "continuationUri", l.ContinuationURI)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LivySessions.
func (l *LivySessions) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "continuationToken":
			err = unpopulate(val, "ContinuationToken", &l.ContinuationToken)
			delete(rawMsg, key)
		case "continuationUri":
			err = unpopulate(val, "ContinuationURI", &l.ContinuationURI)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &l.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Principal.
func (p Principal) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "displayName", p.DisplayName)
	populate(objectMap, "groupDetails", p.GroupDetails)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "servicePrincipalDetails", p.ServicePrincipalDetails)
	populate(objectMap, "servicePrincipalProfileDetails", p.ServicePrincipalProfileDetails)
	populate(objectMap, "type", p.Type)
	populate(objectMap, "userDetails", p.UserDetails)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Principal.
func (p *Principal) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "displayName":
			err = unpopulate(val, "DisplayName", &p.DisplayName)
			delete(rawMsg, key)
		case "groupDetails":
			err = unpopulate(val, "GroupDetails", &p.GroupDetails)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "servicePrincipalDetails":
			err = unpopulate(val, "ServicePrincipalDetails", &p.ServicePrincipalDetails)
			delete(rawMsg, key)
		case "servicePrincipalProfileDetails":
			err = unpopulate(val, "ServicePrincipalProfileDetails", &p.ServicePrincipalProfileDetails)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &p.Type)
			delete(rawMsg, key)
		case "userDetails":
			err = unpopulate(val, "UserDetails", &p.UserDetails)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PrincipalGroupDetails.
func (p PrincipalGroupDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "groupType", p.GroupType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PrincipalGroupDetails.
func (p *PrincipalGroupDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "groupType":
			err = unpopulate(val, "GroupType", &p.GroupType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PrincipalServicePrincipalDetails.
func (p PrincipalServicePrincipalDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "aadAppId", p.AADAppID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PrincipalServicePrincipalDetails.
func (p *PrincipalServicePrincipalDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "aadAppId":
			err = unpopulate(val, "AADAppID", &p.AADAppID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PrincipalServicePrincipalProfileDetails.
func (p PrincipalServicePrincipalProfileDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "parentPrincipal", p.ParentPrincipal)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PrincipalServicePrincipalProfileDetails.
func (p *PrincipalServicePrincipalProfileDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "parentPrincipal":
			err = unpopulate(val, "ParentPrincipal", &p.ParentPrincipal)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PrincipalUserDetails.
func (p PrincipalUserDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "userPrincipalName", p.UserPrincipalName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PrincipalUserDetails.
func (p *PrincipalUserDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "userPrincipalName":
			err = unpopulate(val, "UserPrincipalName", &p.UserPrincipalName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "oneLakeRootPath", p.OneLakeRootPath)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Properties.
func (p *Properties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "oneLakeRootPath":
			err = unpopulate(val, "OneLakeRootPath", &p.OneLakeRootPath)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PublicDefinition.
func (p PublicDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "format", p.Format)
	populate(objectMap, "parts", p.Parts)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PublicDefinition.
func (p *PublicDefinition) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "format":
			err = unpopulate(val, "Format", &p.Format)
			delete(rawMsg, key)
		case "parts":
			err = unpopulate(val, "Parts", &p.Parts)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PublicDefinitionPart.
func (p PublicDefinitionPart) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "path", p.Path)
	populate(objectMap, "payload", p.Payload)
	populate(objectMap, "payloadType", p.PayloadType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PublicDefinitionPart.
func (p *PublicDefinitionPart) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "path":
			err = unpopulate(val, "Path", &p.Path)
			delete(rawMsg, key)
		case "payload":
			err = unpopulate(val, "Payload", &p.Payload)
			delete(rawMsg, key)
		case "payloadType":
			err = unpopulate(val, "PayloadType", &p.PayloadType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Response.
func (r Response) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "definition", r.Definition)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Response.
func (r *Response) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "definition":
			err = unpopulate(val, "Definition", &r.Definition)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RunSparkJobDefinitionRequest.
func (r RunSparkJobDefinitionRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "executionData", r.ExecutionData)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RunSparkJobDefinitionRequest.
func (r *RunSparkJobDefinitionRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "executionData":
			err = unpopulate(val, "ExecutionData", &r.ExecutionData)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SparkJobDefinition.
func (s SparkJobDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", s.Description)
	populate(objectMap, "displayName", s.DisplayName)
	populate(objectMap, "folderId", s.FolderID)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "type", s.Type)
	populate(objectMap, "workspaceId", s.WorkspaceID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SparkJobDefinition.
func (s *SparkJobDefinition) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &s.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &s.DisplayName)
			delete(rawMsg, key)
		case "folderId":
			err = unpopulate(val, "FolderID", &s.FolderID)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &s.ID)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &s.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &s.Type)
			delete(rawMsg, key)
		case "workspaceId":
			err = unpopulate(val, "WorkspaceID", &s.WorkspaceID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SparkJobDefinitions.
func (s SparkJobDefinitions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "continuationToken", s.ContinuationToken)
	populate(objectMap, "continuationUri", s.ContinuationURI)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SparkJobDefinitions.
func (s *SparkJobDefinitions) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "continuationToken":
			err = unpopulate(val, "ContinuationToken", &s.ContinuationToken)
			delete(rawMsg, key)
		case "continuationUri":
			err = unpopulate(val, "ContinuationURI", &s.ContinuationURI)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &s.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UpdateSparkJobDefinitionDefinitionRequest.
func (u UpdateSparkJobDefinitionDefinitionRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "definition", u.Definition)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UpdateSparkJobDefinitionDefinitionRequest.
func (u *UpdateSparkJobDefinitionDefinitionRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "definition":
			err = unpopulate(val, "Definition", &u.Definition)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UpdateSparkJobDefinitionRequest.
func (u UpdateSparkJobDefinitionRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", u.Description)
	populate(objectMap, "displayName", u.DisplayName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UpdateSparkJobDefinitionRequest.
func (u *UpdateSparkJobDefinitionRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &u.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &u.DisplayName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil || string(data) == "null" {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
