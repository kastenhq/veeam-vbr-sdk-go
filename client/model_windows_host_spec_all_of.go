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

// WindowsHostSpecAllOf struct for WindowsHostSpecAllOf
type WindowsHostSpecAllOf struct {
	NetworkSettings *WindowsHostPortsModel `json:"networkSettings,omitempty"`
}

// NewWindowsHostSpecAllOf instantiates a new WindowsHostSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWindowsHostSpecAllOf() *WindowsHostSpecAllOf {
	this := WindowsHostSpecAllOf{}
	return &this
}

// NewWindowsHostSpecAllOfWithDefaults instantiates a new WindowsHostSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWindowsHostSpecAllOfWithDefaults() *WindowsHostSpecAllOf {
	this := WindowsHostSpecAllOf{}
	return &this
}

// GetNetworkSettings returns the NetworkSettings field value if set, zero value otherwise.
func (o *WindowsHostSpecAllOf) GetNetworkSettings() WindowsHostPortsModel {
	if o == nil || isNil(o.NetworkSettings) {
		var ret WindowsHostPortsModel
		return ret
	}
	return *o.NetworkSettings
}

// GetNetworkSettingsOk returns a tuple with the NetworkSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsHostSpecAllOf) GetNetworkSettingsOk() (*WindowsHostPortsModel, bool) {
	if o == nil || isNil(o.NetworkSettings) {
    return nil, false
	}
	return o.NetworkSettings, true
}

// HasNetworkSettings returns a boolean if a field has been set.
func (o *WindowsHostSpecAllOf) HasNetworkSettings() bool {
	if o != nil && !isNil(o.NetworkSettings) {
		return true
	}

	return false
}

// SetNetworkSettings gets a reference to the given WindowsHostPortsModel and assigns it to the NetworkSettings field.
func (o *WindowsHostSpecAllOf) SetNetworkSettings(v WindowsHostPortsModel) {
	o.NetworkSettings = &v
}

func (o WindowsHostSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NetworkSettings) {
		toSerialize["networkSettings"] = o.NetworkSettings
	}
	return json.Marshal(toSerialize)
}

type NullableWindowsHostSpecAllOf struct {
	value *WindowsHostSpecAllOf
	isSet bool
}

func (v NullableWindowsHostSpecAllOf) Get() *WindowsHostSpecAllOf {
	return v.value
}

func (v *NullableWindowsHostSpecAllOf) Set(val *WindowsHostSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWindowsHostSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWindowsHostSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWindowsHostSpecAllOf(val *WindowsHostSpecAllOf) *NullableWindowsHostSpecAllOf {
	return &NullableWindowsHostSpecAllOf{value: val, isSet: true}
}

func (v NullableWindowsHostSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWindowsHostSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


