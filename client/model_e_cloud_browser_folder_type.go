/*
Veeam Backup & Replication REST API

This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br>Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Some request and response bodies refer to reusable schema objects that are defined in the `schemas` section of the REST API specification. Schemas, in turn, may inherit a part of their properties by referring to other schemas, and be polymorphic by referring to multiple alternate schemas.

API version: 1.1-rev0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// ECloudBrowserFolderType Folder type.
type ECloudBrowserFolderType string

// List of ECloudBrowserFolderType
const (
	ECLOUDBROWSERFOLDERTYPE_BACKUP ECloudBrowserFolderType = "backup"
	ECLOUDBROWSERFOLDERTYPE_ARCHIVE ECloudBrowserFolderType = "archive"
)

// All allowed values of ECloudBrowserFolderType enum
var AllowedECloudBrowserFolderTypeEnumValues = []ECloudBrowserFolderType{
	"backup",
	"archive",
}

func (v *ECloudBrowserFolderType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ECloudBrowserFolderType(value)
	for _, existing := range AllowedECloudBrowserFolderTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ECloudBrowserFolderType", value)
}

// NewECloudBrowserFolderTypeFromValue returns a pointer to a valid ECloudBrowserFolderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewECloudBrowserFolderTypeFromValue(v string) (*ECloudBrowserFolderType, error) {
	ev := ECloudBrowserFolderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ECloudBrowserFolderType: valid values are %v", v, AllowedECloudBrowserFolderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ECloudBrowserFolderType) IsValid() bool {
	for _, existing := range AllowedECloudBrowserFolderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ECloudBrowserFolderType value
func (v ECloudBrowserFolderType) Ptr() *ECloudBrowserFolderType {
	return &v
}

type NullableECloudBrowserFolderType struct {
	value *ECloudBrowserFolderType
	isSet bool
}

func (v NullableECloudBrowserFolderType) Get() *ECloudBrowserFolderType {
	return v.value
}

func (v *NullableECloudBrowserFolderType) Set(val *ECloudBrowserFolderType) {
	v.value = val
	v.isSet = true
}

func (v NullableECloudBrowserFolderType) IsSet() bool {
	return v.isSet
}

func (v *NullableECloudBrowserFolderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECloudBrowserFolderType(val *ECloudBrowserFolderType) *NullableECloudBrowserFolderType {
	return &NullableECloudBrowserFolderType{value: val, isSet: true}
}

func (v NullableECloudBrowserFolderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECloudBrowserFolderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

