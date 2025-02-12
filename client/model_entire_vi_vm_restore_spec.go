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

// EntireViVMRestoreSpec struct for EntireViVMRestoreSpec
type EntireViVMRestoreSpec struct {
	// ID of the restore point.
	ObjectRestorePointId string `json:"objectRestorePointId"`
	Type EEntireVMRestoreModeType `json:"type"`
	RestoreProxies *RestoreProxySpec `json:"restoreProxies,omitempty"`
	SecureRestore *SecureRestoreSpec `json:"secureRestore,omitempty"`
	// If *true*, Veeam Backup & Replication starts the restored VM on the target host.
	PowerUp *bool `json:"powerUp,omitempty"`
	// Reason for restoring the VM.
	Reason *string `json:"reason,omitempty"`
}

// NewEntireViVMRestoreSpec instantiates a new EntireViVMRestoreSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntireViVMRestoreSpec(objectRestorePointId string, type_ EEntireVMRestoreModeType) *EntireViVMRestoreSpec {
	this := EntireViVMRestoreSpec{}
	this.ObjectRestorePointId = objectRestorePointId
	this.Type = type_
	return &this
}

// NewEntireViVMRestoreSpecWithDefaults instantiates a new EntireViVMRestoreSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntireViVMRestoreSpecWithDefaults() *EntireViVMRestoreSpec {
	this := EntireViVMRestoreSpec{}
	return &this
}

// GetObjectRestorePointId returns the ObjectRestorePointId field value
func (o *EntireViVMRestoreSpec) GetObjectRestorePointId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectRestorePointId
}

// GetObjectRestorePointIdOk returns a tuple with the ObjectRestorePointId field value
// and a boolean to check if the value has been set.
func (o *EntireViVMRestoreSpec) GetObjectRestorePointIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ObjectRestorePointId, true
}

// SetObjectRestorePointId sets field value
func (o *EntireViVMRestoreSpec) SetObjectRestorePointId(v string) {
	o.ObjectRestorePointId = v
}

// GetType returns the Type field value
func (o *EntireViVMRestoreSpec) GetType() EEntireVMRestoreModeType {
	if o == nil {
		var ret EEntireVMRestoreModeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EntireViVMRestoreSpec) GetTypeOk() (*EEntireVMRestoreModeType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EntireViVMRestoreSpec) SetType(v EEntireVMRestoreModeType) {
	o.Type = v
}

// GetRestoreProxies returns the RestoreProxies field value if set, zero value otherwise.
func (o *EntireViVMRestoreSpec) GetRestoreProxies() RestoreProxySpec {
	if o == nil || isNil(o.RestoreProxies) {
		var ret RestoreProxySpec
		return ret
	}
	return *o.RestoreProxies
}

// GetRestoreProxiesOk returns a tuple with the RestoreProxies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntireViVMRestoreSpec) GetRestoreProxiesOk() (*RestoreProxySpec, bool) {
	if o == nil || isNil(o.RestoreProxies) {
    return nil, false
	}
	return o.RestoreProxies, true
}

// HasRestoreProxies returns a boolean if a field has been set.
func (o *EntireViVMRestoreSpec) HasRestoreProxies() bool {
	if o != nil && !isNil(o.RestoreProxies) {
		return true
	}

	return false
}

// SetRestoreProxies gets a reference to the given RestoreProxySpec and assigns it to the RestoreProxies field.
func (o *EntireViVMRestoreSpec) SetRestoreProxies(v RestoreProxySpec) {
	o.RestoreProxies = &v
}

// GetSecureRestore returns the SecureRestore field value if set, zero value otherwise.
func (o *EntireViVMRestoreSpec) GetSecureRestore() SecureRestoreSpec {
	if o == nil || isNil(o.SecureRestore) {
		var ret SecureRestoreSpec
		return ret
	}
	return *o.SecureRestore
}

// GetSecureRestoreOk returns a tuple with the SecureRestore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntireViVMRestoreSpec) GetSecureRestoreOk() (*SecureRestoreSpec, bool) {
	if o == nil || isNil(o.SecureRestore) {
    return nil, false
	}
	return o.SecureRestore, true
}

// HasSecureRestore returns a boolean if a field has been set.
func (o *EntireViVMRestoreSpec) HasSecureRestore() bool {
	if o != nil && !isNil(o.SecureRestore) {
		return true
	}

	return false
}

// SetSecureRestore gets a reference to the given SecureRestoreSpec and assigns it to the SecureRestore field.
func (o *EntireViVMRestoreSpec) SetSecureRestore(v SecureRestoreSpec) {
	o.SecureRestore = &v
}

// GetPowerUp returns the PowerUp field value if set, zero value otherwise.
func (o *EntireViVMRestoreSpec) GetPowerUp() bool {
	if o == nil || isNil(o.PowerUp) {
		var ret bool
		return ret
	}
	return *o.PowerUp
}

// GetPowerUpOk returns a tuple with the PowerUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntireViVMRestoreSpec) GetPowerUpOk() (*bool, bool) {
	if o == nil || isNil(o.PowerUp) {
    return nil, false
	}
	return o.PowerUp, true
}

// HasPowerUp returns a boolean if a field has been set.
func (o *EntireViVMRestoreSpec) HasPowerUp() bool {
	if o != nil && !isNil(o.PowerUp) {
		return true
	}

	return false
}

// SetPowerUp gets a reference to the given bool and assigns it to the PowerUp field.
func (o *EntireViVMRestoreSpec) SetPowerUp(v bool) {
	o.PowerUp = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *EntireViVMRestoreSpec) GetReason() string {
	if o == nil || isNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntireViVMRestoreSpec) GetReasonOk() (*string, bool) {
	if o == nil || isNil(o.Reason) {
    return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *EntireViVMRestoreSpec) HasReason() bool {
	if o != nil && !isNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *EntireViVMRestoreSpec) SetReason(v string) {
	o.Reason = &v
}

func (o EntireViVMRestoreSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objectRestorePointId"] = o.ObjectRestorePointId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.RestoreProxies) {
		toSerialize["restoreProxies"] = o.RestoreProxies
	}
	if !isNil(o.SecureRestore) {
		toSerialize["secureRestore"] = o.SecureRestore
	}
	if !isNil(o.PowerUp) {
		toSerialize["powerUp"] = o.PowerUp
	}
	if !isNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableEntireViVMRestoreSpec struct {
	value *EntireViVMRestoreSpec
	isSet bool
}

func (v NullableEntireViVMRestoreSpec) Get() *EntireViVMRestoreSpec {
	return v.value
}

func (v *NullableEntireViVMRestoreSpec) Set(val *EntireViVMRestoreSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableEntireViVMRestoreSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableEntireViVMRestoreSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntireViVMRestoreSpec(val *EntireViVMRestoreSpec) *NullableEntireViVMRestoreSpec {
	return &NullableEntireViVMRestoreSpec{value: val, isSet: true}
}

func (v NullableEntireViVMRestoreSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntireViVMRestoreSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


