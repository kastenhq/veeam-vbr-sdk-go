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

// ESennightOfMonth Sennight of the month.
type ESennightOfMonth string

// List of ESennightOfMonth
const (
	ESENNIGHTOFMONTH_FIRST ESennightOfMonth = "First"
	ESENNIGHTOFMONTH_SECOND ESennightOfMonth = "Second"
	ESENNIGHTOFMONTH_THIRD ESennightOfMonth = "Third"
	ESENNIGHTOFMONTH_FOURTH ESennightOfMonth = "Fourth"
	ESENNIGHTOFMONTH_FIFTH ESennightOfMonth = "Fifth"
	ESENNIGHTOFMONTH_LAST ESennightOfMonth = "Last"
)

// All allowed values of ESennightOfMonth enum
var AllowedESennightOfMonthEnumValues = []ESennightOfMonth{
	"First",
	"Second",
	"Third",
	"Fourth",
	"Fifth",
	"Last",
}

func (v *ESennightOfMonth) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ESennightOfMonth(value)
	for _, existing := range AllowedESennightOfMonthEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ESennightOfMonth", value)
}

// NewESennightOfMonthFromValue returns a pointer to a valid ESennightOfMonth
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewESennightOfMonthFromValue(v string) (*ESennightOfMonth, error) {
	ev := ESennightOfMonth(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ESennightOfMonth: valid values are %v", v, AllowedESennightOfMonthEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ESennightOfMonth) IsValid() bool {
	for _, existing := range AllowedESennightOfMonthEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ESennightOfMonth value
func (v ESennightOfMonth) Ptr() *ESennightOfMonth {
	return &v
}

type NullableESennightOfMonth struct {
	value *ESennightOfMonth
	isSet bool
}

func (v NullableESennightOfMonth) Get() *ESennightOfMonth {
	return v.value
}

func (v *NullableESennightOfMonth) Set(val *ESennightOfMonth) {
	v.value = val
	v.isSet = true
}

func (v NullableESennightOfMonth) IsSet() bool {
	return v.isSet
}

func (v *NullableESennightOfMonth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableESennightOfMonth(val *ESennightOfMonth) *NullableESennightOfMonth {
	return &NullableESennightOfMonth{value: val, isSet: true}
}

func (v NullableESennightOfMonth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableESennightOfMonth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

