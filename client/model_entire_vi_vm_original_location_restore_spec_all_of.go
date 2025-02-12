/*
Veeam Backup & Replication REST API

This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br>Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Some request and response bodies refer to reusable schema objects that are defined in the `schemas` section of the REST API specification. Schemas, in turn, may inherit a part of their properties by referring to other schemas, and be polymorphic by referring to multiple alternate schemas.

API version: 1.1-rev0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// EntireViVMOriginalLocationRestoreSpecAllOf struct for EntireViVMOriginalLocationRestoreSpecAllOf
type EntireViVMOriginalLocationRestoreSpecAllOf struct {
	// If *true*, Veeam Backup & Replication performs incremental restore.
	QuickRollback *bool `json:"quickRollback,omitempty"`
}

// NewEntireViVMOriginalLocationRestoreSpecAllOf instantiates a new EntireViVMOriginalLocationRestoreSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntireViVMOriginalLocationRestoreSpecAllOf() *EntireViVMOriginalLocationRestoreSpecAllOf {
	this := EntireViVMOriginalLocationRestoreSpecAllOf{}
	return &this
}

// NewEntireViVMOriginalLocationRestoreSpecAllOfWithDefaults instantiates a new EntireViVMOriginalLocationRestoreSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntireViVMOriginalLocationRestoreSpecAllOfWithDefaults() *EntireViVMOriginalLocationRestoreSpecAllOf {
	this := EntireViVMOriginalLocationRestoreSpecAllOf{}
	return &this
}

// GetQuickRollback returns the QuickRollback field value if set, zero value otherwise.
func (o *EntireViVMOriginalLocationRestoreSpecAllOf) GetQuickRollback() bool {
	if o == nil || isNil(o.QuickRollback) {
		var ret bool
		return ret
	}
	return *o.QuickRollback
}

// GetQuickRollbackOk returns a tuple with the QuickRollback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntireViVMOriginalLocationRestoreSpecAllOf) GetQuickRollbackOk() (*bool, bool) {
	if o == nil || isNil(o.QuickRollback) {
    return nil, false
	}
	return o.QuickRollback, true
}

// HasQuickRollback returns a boolean if a field has been set.
func (o *EntireViVMOriginalLocationRestoreSpecAllOf) HasQuickRollback() bool {
	if o != nil && !isNil(o.QuickRollback) {
		return true
	}

	return false
}

// SetQuickRollback gets a reference to the given bool and assigns it to the QuickRollback field.
func (o *EntireViVMOriginalLocationRestoreSpecAllOf) SetQuickRollback(v bool) {
	o.QuickRollback = &v
}

func (o EntireViVMOriginalLocationRestoreSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.QuickRollback) {
		toSerialize["quickRollback"] = o.QuickRollback
	}
	return json.Marshal(toSerialize)
}

type NullableEntireViVMOriginalLocationRestoreSpecAllOf struct {
	value *EntireViVMOriginalLocationRestoreSpecAllOf
	isSet bool
}

func (v NullableEntireViVMOriginalLocationRestoreSpecAllOf) Get() *EntireViVMOriginalLocationRestoreSpecAllOf {
	return v.value
}

func (v *NullableEntireViVMOriginalLocationRestoreSpecAllOf) Set(val *EntireViVMOriginalLocationRestoreSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEntireViVMOriginalLocationRestoreSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEntireViVMOriginalLocationRestoreSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntireViVMOriginalLocationRestoreSpecAllOf(val *EntireViVMOriginalLocationRestoreSpecAllOf) *NullableEntireViVMOriginalLocationRestoreSpecAllOf {
	return &NullableEntireViVMOriginalLocationRestoreSpecAllOf{value: val, isSet: true}
}

func (v NullableEntireViVMOriginalLocationRestoreSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntireViVMOriginalLocationRestoreSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


