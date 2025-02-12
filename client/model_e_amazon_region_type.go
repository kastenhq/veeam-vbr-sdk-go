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

// EAmazonRegionType AWS region type.
type EAmazonRegionType string

// List of EAmazonRegionType
const (
	EAMAZONREGIONTYPE_CHINA EAmazonRegionType = "China"
	EAMAZONREGIONTYPE_GLOBAL EAmazonRegionType = "Global"
	EAMAZONREGIONTYPE_GOVERNMENT EAmazonRegionType = "Government"
)

// All allowed values of EAmazonRegionType enum
var AllowedEAmazonRegionTypeEnumValues = []EAmazonRegionType{
	"China",
	"Global",
	"Government",
}

func (v *EAmazonRegionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EAmazonRegionType(value)
	for _, existing := range AllowedEAmazonRegionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EAmazonRegionType", value)
}

// NewEAmazonRegionTypeFromValue returns a pointer to a valid EAmazonRegionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEAmazonRegionTypeFromValue(v string) (*EAmazonRegionType, error) {
	ev := EAmazonRegionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EAmazonRegionType: valid values are %v", v, AllowedEAmazonRegionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EAmazonRegionType) IsValid() bool {
	for _, existing := range AllowedEAmazonRegionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EAmazonRegionType value
func (v EAmazonRegionType) Ptr() *EAmazonRegionType {
	return &v
}

type NullableEAmazonRegionType struct {
	value *EAmazonRegionType
	isSet bool
}

func (v NullableEAmazonRegionType) Get() *EAmazonRegionType {
	return v.value
}

func (v *NullableEAmazonRegionType) Set(val *EAmazonRegionType) {
	v.value = val
	v.isSet = true
}

func (v NullableEAmazonRegionType) IsSet() bool {
	return v.isSet
}

func (v *NullableEAmazonRegionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEAmazonRegionType(val *EAmazonRegionType) *NullableEAmazonRegionType {
	return &NullableEAmazonRegionType{value: val, isSet: true}
}

func (v NullableEAmazonRegionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEAmazonRegionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

