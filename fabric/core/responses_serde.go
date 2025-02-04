// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package core

// UnmarshalJSON implements the json.Unmarshaller interface for type GatewaysClientCreateGatewayResponse.
func (g *GatewaysClientCreateGatewayResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalGatewayClassification(data)
	if err != nil {
		return err
	}
	g.GatewayClassification = res
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GatewaysClientGetGatewayResponse.
func (g *GatewaysClientGetGatewayResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalGatewayClassification(data)
	if err != nil {
		return err
	}
	g.GatewayClassification = res
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GatewaysClientUpdateGatewayResponse.
func (g *GatewaysClientUpdateGatewayResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalGatewayClassification(data)
	if err != nil {
		return err
	}
	g.GatewayClassification = res
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GitClientGetMyGitCredentialsResponse.
func (g *GitClientGetMyGitCredentialsResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalGitCredentialsConfigurationResponseClassification(data)
	if err != nil {
		return err
	}
	g.GitCredentialsConfigurationResponseClassification = res
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GitClientUpdateMyGitCredentialsResponse.
func (g *GitClientUpdateMyGitCredentialsResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalGitCredentialsConfigurationResponseClassification(data)
	if err != nil {
		return err
	}
	g.GitCredentialsConfigurationResponseClassification = res
	return nil
}
