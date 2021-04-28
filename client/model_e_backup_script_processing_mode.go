/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 
 *
 * API version: 1.0-rev1
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// EBackupScriptProcessingMode Scenario for scripts execution.
type EBackupScriptProcessingMode string

// List of EBackupScriptProcessingMode
const (
	EBACKUPSCRIPTPROCESSINGMODE_DISABLE_EXEC EBackupScriptProcessingMode = "disableExec"
	EBACKUPSCRIPTPROCESSINGMODE_IGNORE_EXEC_FAILURES EBackupScriptProcessingMode = "ignoreExecFailures"
	EBACKUPSCRIPTPROCESSINGMODE_REQUIRE_SUCCESS EBackupScriptProcessingMode = "requireSuccess"
)

func (v *EBackupScriptProcessingMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EBackupScriptProcessingMode(value)
	for _, existing := range []EBackupScriptProcessingMode{ "disableExec", "ignoreExecFailures", "requireSuccess",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EBackupScriptProcessingMode", value)
}

// Ptr returns reference to EBackupScriptProcessingMode value
func (v EBackupScriptProcessingMode) Ptr() *EBackupScriptProcessingMode {
	return &v
}

type NullableEBackupScriptProcessingMode struct {
	value *EBackupScriptProcessingMode
	isSet bool
}

func (v NullableEBackupScriptProcessingMode) Get() *EBackupScriptProcessingMode {
	return v.value
}

func (v *NullableEBackupScriptProcessingMode) Set(val *EBackupScriptProcessingMode) {
	v.value = val
	v.isSet = true
}

func (v NullableEBackupScriptProcessingMode) IsSet() bool {
	return v.isSet
}

func (v *NullableEBackupScriptProcessingMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEBackupScriptProcessingMode(val *EBackupScriptProcessingMode) *NullableEBackupScriptProcessingMode {
	return &NullableEBackupScriptProcessingMode{value: val, isSet: true}
}

func (v NullableEBackupScriptProcessingMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEBackupScriptProcessingMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

