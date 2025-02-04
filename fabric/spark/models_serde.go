// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package spark

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// MarshalJSON implements the json.Marshaller interface for type AutoScaleProperties.
func (a AutoScaleProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "enabled", a.Enabled)
	populate(objectMap, "maxNodeCount", a.MaxNodeCount)
	populate(objectMap, "minNodeCount", a.MinNodeCount)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AutoScaleProperties.
func (a *AutoScaleProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "enabled":
			err = unpopulate(val, "Enabled", &a.Enabled)
			delete(rawMsg, key)
		case "maxNodeCount":
			err = unpopulate(val, "MaxNodeCount", &a.MaxNodeCount)
			delete(rawMsg, key)
		case "minNodeCount":
			err = unpopulate(val, "MinNodeCount", &a.MinNodeCount)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AutomaticLogProperties.
func (a AutomaticLogProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "enabled", a.Enabled)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AutomaticLogProperties.
func (a *AutomaticLogProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "enabled":
			err = unpopulate(val, "Enabled", &a.Enabled)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CreateCustomPoolRequest.
func (c CreateCustomPoolRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "autoScale", c.AutoScale)
	populate(objectMap, "dynamicExecutorAllocation", c.DynamicExecutorAllocation)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "nodeFamily", c.NodeFamily)
	populate(objectMap, "nodeSize", c.NodeSize)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CreateCustomPoolRequest.
func (c *CreateCustomPoolRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "autoScale":
			err = unpopulate(val, "AutoScale", &c.AutoScale)
			delete(rawMsg, key)
		case "dynamicExecutorAllocation":
			err = unpopulate(val, "DynamicExecutorAllocation", &c.DynamicExecutorAllocation)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &c.Name)
			delete(rawMsg, key)
		case "nodeFamily":
			err = unpopulate(val, "NodeFamily", &c.NodeFamily)
			delete(rawMsg, key)
		case "nodeSize":
			err = unpopulate(val, "NodeSize", &c.NodeSize)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CustomPool.
func (c CustomPool) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "autoScale", c.AutoScale)
	populate(objectMap, "dynamicExecutorAllocation", c.DynamicExecutorAllocation)
	populate(objectMap, "id", c.ID)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "nodeFamily", c.NodeFamily)
	populate(objectMap, "nodeSize", c.NodeSize)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CustomPool.
func (c *CustomPool) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "autoScale":
			err = unpopulate(val, "AutoScale", &c.AutoScale)
			delete(rawMsg, key)
		case "dynamicExecutorAllocation":
			err = unpopulate(val, "DynamicExecutorAllocation", &c.DynamicExecutorAllocation)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &c.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &c.Name)
			delete(rawMsg, key)
		case "nodeFamily":
			err = unpopulate(val, "NodeFamily", &c.NodeFamily)
			delete(rawMsg, key)
		case "nodeSize":
			err = unpopulate(val, "NodeSize", &c.NodeSize)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &c.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CustomPools.
func (c CustomPools) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "continuationToken", c.ContinuationToken)
	populate(objectMap, "continuationUri", c.ContinuationURI)
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CustomPools.
func (c *CustomPools) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "continuationToken":
			err = unpopulate(val, "ContinuationToken", &c.ContinuationToken)
			delete(rawMsg, key)
		case "continuationUri":
			err = unpopulate(val, "ContinuationURI", &c.ContinuationURI)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &c.Value)
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

// MarshalJSON implements the json.Marshaller interface for type EnvironmentProperties.
func (e EnvironmentProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "name", e.Name)
	populate(objectMap, "runtimeVersion", e.RuntimeVersion)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EnvironmentProperties.
func (e *EnvironmentProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &e.Name)
			delete(rawMsg, key)
		case "runtimeVersion":
			err = unpopulate(val, "RuntimeVersion", &e.RuntimeVersion)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type HighConcurrencyProperties.
func (h HighConcurrencyProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "notebookInteractiveRunEnabled", h.NotebookInteractiveRunEnabled)
	populate(objectMap, "notebookPipelineRunEnabled", h.NotebookPipelineRunEnabled)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type HighConcurrencyProperties.
func (h *HighConcurrencyProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", h, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "notebookInteractiveRunEnabled":
			err = unpopulate(val, "NotebookInteractiveRunEnabled", &h.NotebookInteractiveRunEnabled)
			delete(rawMsg, key)
		case "notebookPipelineRunEnabled":
			err = unpopulate(val, "NotebookPipelineRunEnabled", &h.NotebookPipelineRunEnabled)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", h, err)
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

// MarshalJSON implements the json.Marshaller interface for type JobsProperties.
func (j JobsProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "conservativeJobAdmissionEnabled", j.ConservativeJobAdmissionEnabled)
	populate(objectMap, "sessionTimeoutInMinutes", j.SessionTimeoutInMinutes)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type JobsProperties.
func (j *JobsProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", j, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "conservativeJobAdmissionEnabled":
			err = unpopulate(val, "ConservativeJobAdmissionEnabled", &j.ConservativeJobAdmissionEnabled)
			delete(rawMsg, key)
		case "sessionTimeoutInMinutes":
			err = unpopulate(val, "SessionTimeoutInMinutes", &j.SessionTimeoutInMinutes)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", j, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PoolProperties.
func (p PoolProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "customizeComputeEnabled", p.CustomizeComputeEnabled)
	populate(objectMap, "defaultPool", p.DefaultPool)
	populate(objectMap, "starterPool", p.StarterPool)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PoolProperties.
func (p *PoolProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "customizeComputeEnabled":
			err = unpopulate(val, "CustomizeComputeEnabled", &p.CustomizeComputeEnabled)
			delete(rawMsg, key)
		case "defaultPool":
			err = unpopulate(val, "DefaultPool", &p.DefaultPool)
			delete(rawMsg, key)
		case "starterPool":
			err = unpopulate(val, "StarterPool", &p.StarterPool)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StarterPoolProperties.
func (s StarterPoolProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "maxExecutors", s.MaxExecutors)
	populate(objectMap, "maxNodeCount", s.MaxNodeCount)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StarterPoolProperties.
func (s *StarterPoolProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "maxExecutors":
			err = unpopulate(val, "MaxExecutors", &s.MaxExecutors)
			delete(rawMsg, key)
		case "maxNodeCount":
			err = unpopulate(val, "MaxNodeCount", &s.MaxNodeCount)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UpdateCustomPoolRequest.
func (u UpdateCustomPoolRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "autoScale", u.AutoScale)
	populate(objectMap, "dynamicExecutorAllocation", u.DynamicExecutorAllocation)
	populate(objectMap, "name", u.Name)
	populate(objectMap, "nodeFamily", u.NodeFamily)
	populate(objectMap, "nodeSize", u.NodeSize)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UpdateCustomPoolRequest.
func (u *UpdateCustomPoolRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "autoScale":
			err = unpopulate(val, "AutoScale", &u.AutoScale)
			delete(rawMsg, key)
		case "dynamicExecutorAllocation":
			err = unpopulate(val, "DynamicExecutorAllocation", &u.DynamicExecutorAllocation)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &u.Name)
			delete(rawMsg, key)
		case "nodeFamily":
			err = unpopulate(val, "NodeFamily", &u.NodeFamily)
			delete(rawMsg, key)
		case "nodeSize":
			err = unpopulate(val, "NodeSize", &u.NodeSize)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UpdateWorkspaceSparkSettingsRequest.
func (u UpdateWorkspaceSparkSettingsRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "automaticLog", u.AutomaticLog)
	populate(objectMap, "environment", u.Environment)
	populate(objectMap, "highConcurrency", u.HighConcurrency)
	populate(objectMap, "job", u.Job)
	populate(objectMap, "pool", u.Pool)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UpdateWorkspaceSparkSettingsRequest.
func (u *UpdateWorkspaceSparkSettingsRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "automaticLog":
			err = unpopulate(val, "AutomaticLog", &u.AutomaticLog)
			delete(rawMsg, key)
		case "environment":
			err = unpopulate(val, "Environment", &u.Environment)
			delete(rawMsg, key)
		case "highConcurrency":
			err = unpopulate(val, "HighConcurrency", &u.HighConcurrency)
			delete(rawMsg, key)
		case "job":
			err = unpopulate(val, "Job", &u.Job)
			delete(rawMsg, key)
		case "pool":
			err = unpopulate(val, "Pool", &u.Pool)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type WorkspaceSparkSettings.
func (w WorkspaceSparkSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "automaticLog", w.AutomaticLog)
	populate(objectMap, "environment", w.Environment)
	populate(objectMap, "highConcurrency", w.HighConcurrency)
	populate(objectMap, "job", w.Job)
	populate(objectMap, "pool", w.Pool)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WorkspaceSparkSettings.
func (w *WorkspaceSparkSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "automaticLog":
			err = unpopulate(val, "AutomaticLog", &w.AutomaticLog)
			delete(rawMsg, key)
		case "environment":
			err = unpopulate(val, "Environment", &w.Environment)
			delete(rawMsg, key)
		case "highConcurrency":
			err = unpopulate(val, "HighConcurrency", &w.HighConcurrency)
			delete(rawMsg, key)
		case "job":
			err = unpopulate(val, "Job", &w.Job)
			delete(rawMsg, key)
		case "pool":
			err = unpopulate(val, "Pool", &w.Pool)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
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
