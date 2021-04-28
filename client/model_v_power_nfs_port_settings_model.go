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
)

// VPowerNFSPortSettingsModel Network ports used by the vPower NFS Service.
type VPowerNFSPortSettingsModel struct {
	// Mount port.
	MountPort *int32 `json:"mountPort,omitempty"`
	// vPower NFS port.
	VPowerNFSPort *int32 `json:"vPowerNFSPort,omitempty"`
}

// NewVPowerNFSPortSettingsModel instantiates a new VPowerNFSPortSettingsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVPowerNFSPortSettingsModel() *VPowerNFSPortSettingsModel {
	this := VPowerNFSPortSettingsModel{}
	return &this
}

// NewVPowerNFSPortSettingsModelWithDefaults instantiates a new VPowerNFSPortSettingsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVPowerNFSPortSettingsModelWithDefaults() *VPowerNFSPortSettingsModel {
	this := VPowerNFSPortSettingsModel{}
	return &this
}

// GetMountPort returns the MountPort field value if set, zero value otherwise.
func (o *VPowerNFSPortSettingsModel) GetMountPort() int32 {
	if o == nil || o.MountPort == nil {
		var ret int32
		return ret
	}
	return *o.MountPort
}

// GetMountPortOk returns a tuple with the MountPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPowerNFSPortSettingsModel) GetMountPortOk() (*int32, bool) {
	if o == nil || o.MountPort == nil {
		return nil, false
	}
	return o.MountPort, true
}

// HasMountPort returns a boolean if a field has been set.
func (o *VPowerNFSPortSettingsModel) HasMountPort() bool {
	if o != nil && o.MountPort != nil {
		return true
	}

	return false
}

// SetMountPort gets a reference to the given int32 and assigns it to the MountPort field.
func (o *VPowerNFSPortSettingsModel) SetMountPort(v int32) {
	o.MountPort = &v
}

// GetVPowerNFSPort returns the VPowerNFSPort field value if set, zero value otherwise.
func (o *VPowerNFSPortSettingsModel) GetVPowerNFSPort() int32 {
	if o == nil || o.VPowerNFSPort == nil {
		var ret int32
		return ret
	}
	return *o.VPowerNFSPort
}

// GetVPowerNFSPortOk returns a tuple with the VPowerNFSPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPowerNFSPortSettingsModel) GetVPowerNFSPortOk() (*int32, bool) {
	if o == nil || o.VPowerNFSPort == nil {
		return nil, false
	}
	return o.VPowerNFSPort, true
}

// HasVPowerNFSPort returns a boolean if a field has been set.
func (o *VPowerNFSPortSettingsModel) HasVPowerNFSPort() bool {
	if o != nil && o.VPowerNFSPort != nil {
		return true
	}

	return false
}

// SetVPowerNFSPort gets a reference to the given int32 and assigns it to the VPowerNFSPort field.
func (o *VPowerNFSPortSettingsModel) SetVPowerNFSPort(v int32) {
	o.VPowerNFSPort = &v
}

func (o VPowerNFSPortSettingsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MountPort != nil {
		toSerialize["mountPort"] = o.MountPort
	}
	if o.VPowerNFSPort != nil {
		toSerialize["vPowerNFSPort"] = o.VPowerNFSPort
	}
	return json.Marshal(toSerialize)
}

type NullableVPowerNFSPortSettingsModel struct {
	value *VPowerNFSPortSettingsModel
	isSet bool
}

func (v NullableVPowerNFSPortSettingsModel) Get() *VPowerNFSPortSettingsModel {
	return v.value
}

func (v *NullableVPowerNFSPortSettingsModel) Set(val *VPowerNFSPortSettingsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVPowerNFSPortSettingsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVPowerNFSPortSettingsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVPowerNFSPortSettingsModel(val *VPowerNFSPortSettingsModel) *NullableVPowerNFSPortSettingsModel {
	return &NullableVPowerNFSPortSettingsModel{value: val, isSet: true}
}

func (v NullableVPowerNFSPortSettingsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVPowerNFSPortSettingsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


