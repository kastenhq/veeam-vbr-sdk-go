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

// WasabiCloudStorageBrowserModel struct for WasabiCloudStorageBrowserModel
type WasabiCloudStorageBrowserModel struct {
	CloudBrowserModel
	// ID of a server used to connect to the object storage.
	HostId *string `json:"hostId,omitempty"`
	// Array of regions.
	Regions []WasabiCloudStorageRegionBrowserModel `json:"regions,omitempty"`
}

// NewWasabiCloudStorageBrowserModel instantiates a new WasabiCloudStorageBrowserModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWasabiCloudStorageBrowserModel() *WasabiCloudStorageBrowserModel {
	this := WasabiCloudStorageBrowserModel{}
	return &this
}

// NewWasabiCloudStorageBrowserModelWithDefaults instantiates a new WasabiCloudStorageBrowserModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWasabiCloudStorageBrowserModelWithDefaults() *WasabiCloudStorageBrowserModel {
	this := WasabiCloudStorageBrowserModel{}
	return &this
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *WasabiCloudStorageBrowserModel) GetHostId() string {
	if o == nil || isNil(o.HostId) {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WasabiCloudStorageBrowserModel) GetHostIdOk() (*string, bool) {
	if o == nil || isNil(o.HostId) {
    return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *WasabiCloudStorageBrowserModel) HasHostId() bool {
	if o != nil && !isNil(o.HostId) {
		return true
	}

	return false
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *WasabiCloudStorageBrowserModel) SetHostId(v string) {
	o.HostId = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *WasabiCloudStorageBrowserModel) GetRegions() []WasabiCloudStorageRegionBrowserModel {
	if o == nil || isNil(o.Regions) {
		var ret []WasabiCloudStorageRegionBrowserModel
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WasabiCloudStorageBrowserModel) GetRegionsOk() ([]WasabiCloudStorageRegionBrowserModel, bool) {
	if o == nil || isNil(o.Regions) {
    return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *WasabiCloudStorageBrowserModel) HasRegions() bool {
	if o != nil && !isNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []WasabiCloudStorageRegionBrowserModel and assigns it to the Regions field.
func (o *WasabiCloudStorageBrowserModel) SetRegions(v []WasabiCloudStorageRegionBrowserModel) {
	o.Regions = v
}

func (o WasabiCloudStorageBrowserModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBrowserModel, errCloudBrowserModel := json.Marshal(o.CloudBrowserModel)
	if errCloudBrowserModel != nil {
		return []byte{}, errCloudBrowserModel
	}
	errCloudBrowserModel = json.Unmarshal([]byte(serializedCloudBrowserModel), &toSerialize)
	if errCloudBrowserModel != nil {
		return []byte{}, errCloudBrowserModel
	}
	if !isNil(o.HostId) {
		toSerialize["hostId"] = o.HostId
	}
	if !isNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	return json.Marshal(toSerialize)
}

type NullableWasabiCloudStorageBrowserModel struct {
	value *WasabiCloudStorageBrowserModel
	isSet bool
}

func (v NullableWasabiCloudStorageBrowserModel) Get() *WasabiCloudStorageBrowserModel {
	return v.value
}

func (v *NullableWasabiCloudStorageBrowserModel) Set(val *WasabiCloudStorageBrowserModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWasabiCloudStorageBrowserModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWasabiCloudStorageBrowserModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWasabiCloudStorageBrowserModel(val *WasabiCloudStorageBrowserModel) *NullableWasabiCloudStorageBrowserModel {
	return &NullableWasabiCloudStorageBrowserModel{value: val, isSet: true}
}

func (v NullableWasabiCloudStorageBrowserModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWasabiCloudStorageBrowserModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


