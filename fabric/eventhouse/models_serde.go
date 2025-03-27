// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package eventhouse

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// MarshalJSON implements the json.Marshaller interface for type CreateEventhouseRequest.
func (c CreateEventhouseRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "creationPayload", c.CreationPayload)
	populate(objectMap, "definition", c.Definition)
	populate(objectMap, "description", c.Description)
	populate(objectMap, "displayName", c.DisplayName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CreateEventhouseRequest.
func (c *CreateEventhouseRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "creationPayload":
			err = unpopulate(val, "CreationPayload", &c.CreationPayload)
			delete(rawMsg, key)
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

// MarshalJSON implements the json.Marshaller interface for type CreationPayload.
func (c CreationPayload) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "minimumConsumptionUnits", c.MinimumConsumptionUnits)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CreationPayload.
func (c *CreationPayload) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "minimumConsumptionUnits":
			err = unpopulate(val, "MinimumConsumptionUnits", &c.MinimumConsumptionUnits)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Definition.
func (d Definition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "format", d.Format)
	populate(objectMap, "parts", d.Parts)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Definition.
func (d *Definition) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "format":
			err = unpopulate(val, "Format", &d.Format)
			delete(rawMsg, key)
		case "parts":
			err = unpopulate(val, "Parts", &d.Parts)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DefinitionPart.
func (d DefinitionPart) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "path", d.Path)
	populate(objectMap, "payload", d.Payload)
	populate(objectMap, "payloadType", d.PayloadType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DefinitionPart.
func (d *DefinitionPart) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "path":
			err = unpopulate(val, "Path", &d.Path)
			delete(rawMsg, key)
		case "payload":
			err = unpopulate(val, "Payload", &d.Payload)
			delete(rawMsg, key)
		case "payloadType":
			err = unpopulate(val, "PayloadType", &d.PayloadType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DefinitionResponse.
func (d DefinitionResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "definition", d.Definition)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DefinitionResponse.
func (d *DefinitionResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "definition":
			err = unpopulate(val, "Definition", &d.Definition)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Eventhouse.
func (e Eventhouse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", e.Description)
	populate(objectMap, "displayName", e.DisplayName)
	populate(objectMap, "folderId", e.FolderID)
	populate(objectMap, "id", e.ID)
	populate(objectMap, "properties", e.Properties)
	populate(objectMap, "type", e.Type)
	populate(objectMap, "workspaceId", e.WorkspaceID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Eventhouse.
func (e *Eventhouse) UnmarshalJSON(data []byte) error {
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
		case "folderId":
			err = unpopulate(val, "FolderID", &e.FolderID)
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

// MarshalJSON implements the json.Marshaller interface for type Eventhouses.
func (e Eventhouses) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "continuationToken", e.ContinuationToken)
	populate(objectMap, "continuationUri", e.ContinuationURI)
	populate(objectMap, "value", e.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Eventhouses.
func (e *Eventhouses) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type Properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "databasesItemIds", p.DatabasesItemIDs)
	populate(objectMap, "ingestionServiceUri", p.IngestionServiceURI)
	populate(objectMap, "minimumConsumptionUnits", p.MinimumConsumptionUnits)
	populate(objectMap, "queryServiceUri", p.QueryServiceURI)
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
		case "databasesItemIds":
			err = unpopulate(val, "DatabasesItemIDs", &p.DatabasesItemIDs)
			delete(rawMsg, key)
		case "ingestionServiceUri":
			err = unpopulate(val, "IngestionServiceURI", &p.IngestionServiceURI)
			delete(rawMsg, key)
		case "minimumConsumptionUnits":
			err = unpopulate(val, "MinimumConsumptionUnits", &p.MinimumConsumptionUnits)
			delete(rawMsg, key)
		case "queryServiceUri":
			err = unpopulate(val, "QueryServiceURI", &p.QueryServiceURI)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UpdateEventhouseDefinitionRequest.
func (u UpdateEventhouseDefinitionRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "definition", u.Definition)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UpdateEventhouseDefinitionRequest.
func (u *UpdateEventhouseDefinitionRequest) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type UpdateEventhouseRequest.
func (u UpdateEventhouseRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", u.Description)
	populate(objectMap, "displayName", u.DisplayName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UpdateEventhouseRequest.
func (u *UpdateEventhouseRequest) UnmarshalJSON(data []byte) error {
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
