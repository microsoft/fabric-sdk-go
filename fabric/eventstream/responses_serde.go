// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package eventstream

// UnmarshalJSON implements the json.Unmarshaller interface for type TopologyClientGetEventstreamDestinationResponse.
func (t *TopologyClientGetEventstreamDestinationResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalDestinationResponseClassification(data)
	if err != nil {
		return err
	}
	t.DestinationResponseClassification = res
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TopologyClientGetEventstreamSourceResponse.
func (t *TopologyClientGetEventstreamSourceResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalSourceResponseClassification(data)
	if err != nil {
		return err
	}
	t.SourceResponseClassification = res
	return nil
}
