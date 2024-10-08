// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// SPDX-License-Identifier: MIT

package fake

import (
	"io"
	"net/http"
	"strings"

	"github.com/microsoft/fabric-sdk-go/fabric/core"
)

func SetResponseError(httpStatusCode int, errorCode, errorMessage string) error {
	errResp := core.ErrorResponse{
		ErrorCode: &errorCode,
		Message:   &errorMessage,
	}

	errRespBytes, err := errResp.MarshalJSON()
	if err != nil {
		return err
	}

	headers := http.Header{}
	headers.Set("x-ms-public-api-error-code", errorCode)
	resp := &http.Response{
		Status:     http.StatusText(httpStatusCode),
		StatusCode: httpStatusCode,
		Header:     headers,
		Body:       io.NopCloser(strings.NewReader(string(errRespBytes))),
	}

	return core.NewResponseError(resp)
}
