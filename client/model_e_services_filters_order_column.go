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

// EServicesFiltersOrderColumn the model 'EServicesFiltersOrderColumn'
type EServicesFiltersOrderColumn string

// List of EServicesFiltersOrderColumn
const (
	ESERVICESFILTERSORDERCOLUMN_NAME EServicesFiltersOrderColumn = "Name"
	ESERVICESFILTERSORDERCOLUMN_PORT EServicesFiltersOrderColumn = "Port"
)

func (v *EServicesFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EServicesFiltersOrderColumn(value)
	for _, existing := range []EServicesFiltersOrderColumn{ "Name", "Port",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EServicesFiltersOrderColumn", value)
}

// Ptr returns reference to EServicesFiltersOrderColumn value
func (v EServicesFiltersOrderColumn) Ptr() *EServicesFiltersOrderColumn {
	return &v
}

type NullableEServicesFiltersOrderColumn struct {
	value *EServicesFiltersOrderColumn
	isSet bool
}

func (v NullableEServicesFiltersOrderColumn) Get() *EServicesFiltersOrderColumn {
	return v.value
}

func (v *NullableEServicesFiltersOrderColumn) Set(val *EServicesFiltersOrderColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableEServicesFiltersOrderColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableEServicesFiltersOrderColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEServicesFiltersOrderColumn(val *EServicesFiltersOrderColumn) *NullableEServicesFiltersOrderColumn {
	return &NullableEServicesFiltersOrderColumn{value: val, isSet: true}
}

func (v NullableEServicesFiltersOrderColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEServicesFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

