// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package warehouse

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// MarshalJSON implements the json.Marshaller interface for type CreateWarehouseRequest.
func (c CreateWarehouseRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", c.Description)
	populate(objectMap, "displayName", c.DisplayName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CreateWarehouseRequest.
func (c *CreateWarehouseRequest) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type Properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "connectionString", p.ConnectionString)
	populateDateTimeRFC3339(objectMap, "createdDate", p.CreatedDate)
	populateDateTimeRFC3339(objectMap, "lastUpdatedTime", p.LastUpdatedTime)
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
		case "connectionString":
			err = unpopulate(val, "ConnectionString", &p.ConnectionString)
			delete(rawMsg, key)
		case "createdDate":
			err = unpopulateDateTimeRFC3339(val, "CreatedDate", &p.CreatedDate)
			delete(rawMsg, key)
		case "lastUpdatedTime":
			err = unpopulateDateTimeRFC3339(val, "LastUpdatedTime", &p.LastUpdatedTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UpdateWarehouseRequest.
func (u UpdateWarehouseRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", u.Description)
	populate(objectMap, "displayName", u.DisplayName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UpdateWarehouseRequest.
func (u *UpdateWarehouseRequest) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type Warehouse.
func (w Warehouse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", w.Description)
	populate(objectMap, "displayName", w.DisplayName)
	populate(objectMap, "id", w.ID)
	populate(objectMap, "properties", w.Properties)
	populate(objectMap, "type", w.Type)
	populate(objectMap, "workspaceId", w.WorkspaceID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Warehouse.
func (w *Warehouse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &w.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &w.DisplayName)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &w.ID)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &w.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &w.Type)
			delete(rawMsg, key)
		case "workspaceId":
			err = unpopulate(val, "WorkspaceID", &w.WorkspaceID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Warehouses.
func (w Warehouses) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "continuationToken", w.ContinuationToken)
	populate(objectMap, "continuationUri", w.ContinuationURI)
	populate(objectMap, "value", w.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Warehouses.
func (w *Warehouses) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "continuationToken":
			err = unpopulate(val, "ContinuationToken", &w.ContinuationToken)
			delete(rawMsg, key)
		case "continuationUri":
			err = unpopulate(val, "ContinuationURI", &w.ContinuationURI)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &w.Value)
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