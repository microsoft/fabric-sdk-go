// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package environment

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// MarshalJSON implements the json.Marshaller interface for type ComponentPublishInfo.
func (c ComponentPublishInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "sparkLibraries", c.SparkLibraries)
	populate(objectMap, "sparkSettings", c.SparkSettings)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ComponentPublishInfo.
func (c *ComponentPublishInfo) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "sparkLibraries":
			err = unpopulate(val, "SparkLibraries", &c.SparkLibraries)
			delete(rawMsg, key)
		case "sparkSettings":
			err = unpopulate(val, "SparkSettings", &c.SparkSettings)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CreateEnvironmentRequest.
func (c CreateEnvironmentRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", c.Description)
	populate(objectMap, "displayName", c.DisplayName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CreateEnvironmentRequest.
func (c *CreateEnvironmentRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
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

// MarshalJSON implements the json.Marshaller interface for type CustomLibraries.
func (c CustomLibraries) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "jarFiles", c.JarFiles)
	populate(objectMap, "pyFiles", c.PyFiles)
	populate(objectMap, "rTarFiles", c.RTarFiles)
	populate(objectMap, "wheelFiles", c.WheelFiles)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CustomLibraries.
func (c *CustomLibraries) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "jarFiles":
			err = unpopulate(val, "JarFiles", &c.JarFiles)
			delete(rawMsg, key)
		case "pyFiles":
			err = unpopulate(val, "PyFiles", &c.PyFiles)
			delete(rawMsg, key)
		case "rTarFiles":
			err = unpopulate(val, "RTarFiles", &c.RTarFiles)
			delete(rawMsg, key)
		case "wheelFiles":
			err = unpopulate(val, "WheelFiles", &c.WheelFiles)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DynamicExecutorAllocationProperties.
func (d DynamicExecutorAllocationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "enabled", d.Enabled)
	populate(objectMap, "maxExecutors", d.MaxExecutors)
	populate(objectMap, "minExecutors", d.MinExecutors)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DynamicExecutorAllocationProperties.
func (d *DynamicExecutorAllocationProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "enabled":
			err = unpopulate(val, "Enabled", &d.Enabled)
			delete(rawMsg, key)
		case "maxExecutors":
			err = unpopulate(val, "MaxExecutors", &d.MaxExecutors)
			delete(rawMsg, key)
		case "minExecutors":
			err = unpopulate(val, "MinExecutors", &d.MinExecutors)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Environment.
func (e Environment) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", e.Description)
	populate(objectMap, "displayName", e.DisplayName)
	populate(objectMap, "id", e.ID)
	populate(objectMap, "properties", e.Properties)
	populate(objectMap, "type", e.Type)
	populate(objectMap, "workspaceId", e.WorkspaceID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Environment.
func (e *Environment) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &e.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &e.DisplayName)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &e.ID)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &e.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &e.Type)
			delete(rawMsg, key)
		case "workspaceId":
			err = unpopulate(val, "WorkspaceID", &e.WorkspaceID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Environments.
func (e Environments) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "continuationToken", e.ContinuationToken)
	populate(objectMap, "continuationUri", e.ContinuationURI)
	populate(objectMap, "value", e.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Environments.
func (e *Environments) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "continuationToken":
			err = unpopulate(val, "ContinuationToken", &e.ContinuationToken)
			delete(rawMsg, key)
		case "continuationUri":
			err = unpopulate(val, "ContinuationURI", &e.ContinuationURI)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &e.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type InstancePool.
func (i InstancePool) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", i.ID)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "type", i.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type InstancePool.
func (i *InstancePool) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &i.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &i.Name)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Libraries.
func (l Libraries) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "customLibraries", l.CustomLibraries)
	populate(objectMap, "environmentYml", l.EnvironmentYml)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Libraries.
func (l *Libraries) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "customLibraries":
			err = unpopulate(val, "CustomLibraries", &l.CustomLibraries)
			delete(rawMsg, key)
		case "environmentYml":
			err = unpopulate(val, "EnvironmentYml", &l.EnvironmentYml)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PublishDetails.
func (p PublishDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "componentPublishInfo", p.ComponentPublishInfo)
	populateDateTimeRFC3339(objectMap, "endTime", p.EndTime)
	populateDateTimeRFC3339(objectMap, "startTime", p.StartTime)
	populate(objectMap, "state", p.State)
	populate(objectMap, "targetVersion", p.TargetVersion)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PublishDetails.
func (p *PublishDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "componentPublishInfo":
			err = unpopulate(val, "ComponentPublishInfo", &p.ComponentPublishInfo)
			delete(rawMsg, key)
		case "endTime":
			err = unpopulateDateTimeRFC3339(val, "EndTime", &p.EndTime)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateDateTimeRFC3339(val, "StartTime", &p.StartTime)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, "State", &p.State)
			delete(rawMsg, key)
		case "targetVersion":
			err = unpopulate(val, "TargetVersion", &p.TargetVersion)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PublishInfo.
func (p PublishInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "publishDetails", p.PublishDetails)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PublishInfo.
func (p *PublishInfo) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "publishDetails":
			err = unpopulate(val, "PublishDetails", &p.PublishDetails)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SparkCompute.
func (s SparkCompute) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "driverCores", s.DriverCores)
	populate(objectMap, "driverMemory", s.DriverMemory)
	populate(objectMap, "dynamicExecutorAllocation", s.DynamicExecutorAllocation)
	populate(objectMap, "executorCores", s.ExecutorCores)
	populate(objectMap, "executorMemory", s.ExecutorMemory)
	populate(objectMap, "instancePool", s.InstancePool)
	populate(objectMap, "runtimeVersion", s.RuntimeVersion)
	populateAny(objectMap, "sparkProperties", s.SparkProperties)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SparkCompute.
func (s *SparkCompute) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "driverCores":
			err = unpopulate(val, "DriverCores", &s.DriverCores)
			delete(rawMsg, key)
		case "driverMemory":
			err = unpopulate(val, "DriverMemory", &s.DriverMemory)
			delete(rawMsg, key)
		case "dynamicExecutorAllocation":
			err = unpopulate(val, "DynamicExecutorAllocation", &s.DynamicExecutorAllocation)
			delete(rawMsg, key)
		case "executorCores":
			err = unpopulate(val, "ExecutorCores", &s.ExecutorCores)
			delete(rawMsg, key)
		case "executorMemory":
			err = unpopulate(val, "ExecutorMemory", &s.ExecutorMemory)
			delete(rawMsg, key)
		case "instancePool":
			err = unpopulate(val, "InstancePool", &s.InstancePool)
			delete(rawMsg, key)
		case "runtimeVersion":
			err = unpopulate(val, "RuntimeVersion", &s.RuntimeVersion)
			delete(rawMsg, key)
		case "sparkProperties":
			err = unpopulate(val, "SparkProperties", &s.SparkProperties)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SparkLibraries.
func (s SparkLibraries) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "state", s.State)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SparkLibraries.
func (s *SparkLibraries) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "state":
			err = unpopulate(val, "State", &s.State)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SparkSettings.
func (s SparkSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "state", s.State)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SparkSettings.
func (s *SparkSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "state":
			err = unpopulate(val, "State", &s.State)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UpdateEnvironmentRequest.
func (u UpdateEnvironmentRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", u.Description)
	populate(objectMap, "displayName", u.DisplayName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UpdateEnvironmentRequest.
func (u *UpdateEnvironmentRequest) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type UpdateEnvironmentSparkComputeRequest.
func (u UpdateEnvironmentSparkComputeRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "driverCores", u.DriverCores)
	populate(objectMap, "driverMemory", u.DriverMemory)
	populate(objectMap, "dynamicExecutorAllocation", u.DynamicExecutorAllocation)
	populate(objectMap, "executorCores", u.ExecutorCores)
	populate(objectMap, "executorMemory", u.ExecutorMemory)
	populate(objectMap, "instancePool", u.InstancePool)
	populate(objectMap, "runtimeVersion", u.RuntimeVersion)
	populateAny(objectMap, "sparkProperties", u.SparkProperties)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UpdateEnvironmentSparkComputeRequest.
func (u *UpdateEnvironmentSparkComputeRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "driverCores":
			err = unpopulate(val, "DriverCores", &u.DriverCores)
			delete(rawMsg, key)
		case "driverMemory":
			err = unpopulate(val, "DriverMemory", &u.DriverMemory)
			delete(rawMsg, key)
		case "dynamicExecutorAllocation":
			err = unpopulate(val, "DynamicExecutorAllocation", &u.DynamicExecutorAllocation)
			delete(rawMsg, key)
		case "executorCores":
			err = unpopulate(val, "ExecutorCores", &u.ExecutorCores)
			delete(rawMsg, key)
		case "executorMemory":
			err = unpopulate(val, "ExecutorMemory", &u.ExecutorMemory)
			delete(rawMsg, key)
		case "instancePool":
			err = unpopulate(val, "InstancePool", &u.InstancePool)
			delete(rawMsg, key)
		case "runtimeVersion":
			err = unpopulate(val, "RuntimeVersion", &u.RuntimeVersion)
			delete(rawMsg, key)
		case "sparkProperties":
			err = unpopulate(val, "SparkProperties", &u.SparkProperties)
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

func populateAny(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else {
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
