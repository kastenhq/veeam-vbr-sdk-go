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

// EJobType Type of the job.
type EJobType string

// List of EJobType
const (
	EJOBTYPE_BACKUP EJobType = "Backup"
)

func (v *EJobType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EJobType(value)
	for _, existing := range []EJobType{ "Backup",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EJobType", value)
}

// Ptr returns reference to EJobType value
func (v EJobType) Ptr() *EJobType {
	return &v
}

type NullableEJobType struct {
	value *EJobType
	isSet bool
}

func (v NullableEJobType) Get() *EJobType {
	return v.value
}

func (v *NullableEJobType) Set(val *EJobType) {
	v.value = val
	v.isSet = true
}

func (v NullableEJobType) IsSet() bool {
	return v.isSet
}

func (v *NullableEJobType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEJobType(val *EJobType) *NullableEJobType {
	return &NullableEJobType{value: val, isSet: true}
}

func (v NullableEJobType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEJobType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

