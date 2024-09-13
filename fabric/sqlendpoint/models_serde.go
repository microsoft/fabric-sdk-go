// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package sqlendpoint

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// MarshalJSON implements the json.Marshaller interface for type SQLEndpoint.
func (s SQLEndpoint) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", s.Description)
	populate(objectMap, "displayName", s.DisplayName)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "type", s.Type)
	populate(objectMap, "workspaceId", s.WorkspaceID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SQLEndpoint.
func (s *SQLEndpoint) UnmarshalJSON(data []byte) error {
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
		case "id":
			err = unpopulate(val, "ID", &s.ID)
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

// MarshalJSON implements the json.Marshaller interface for type SQLEndpoints.
func (s SQLEndpoints) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "continuationToken", s.ContinuationToken)
	populate(objectMap, "continuationUri", s.ContinuationURI)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SQLEndpoints.
func (s *SQLEndpoints) UnmarshalJSON(data []byte) error {
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