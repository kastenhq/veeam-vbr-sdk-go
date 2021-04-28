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

// SmbStorageSpec struct for SmbStorageSpec
type SmbStorageSpec struct {
	RepositorySpec
	Share SmbRepositoryShareSettingsModel `json:"share"`
	Repository NetworkRepositorySettingsModel `json:"repository"`
	MountServer MountServerSettingsModel `json:"mountServer"`
}

// NewSmbStorageSpec instantiates a new SmbStorageSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmbStorageSpec(share SmbRepositoryShareSettingsModel, repository NetworkRepositorySettingsModel, mountServer MountServerSettingsModel, name string, description string, type_ ERepositoryType) *SmbStorageSpec {
	this := SmbStorageSpec{}
	this.Name = name
	this.Description = description
	this.Type = type_
	this.Share = share
	this.Repository = repository
	this.MountServer = mountServer
	return &this
}

// NewSmbStorageSpecWithDefaults instantiates a new SmbStorageSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmbStorageSpecWithDefaults() *SmbStorageSpec {
	this := SmbStorageSpec{}
	return &this
}

// GetShare returns the Share field value
func (o *SmbStorageSpec) GetShare() SmbRepositoryShareSettingsModel {
	if o == nil {
		var ret SmbRepositoryShareSettingsModel
		return ret
	}

	return o.Share
}

// GetShareOk returns a tuple with the Share field value
// and a boolean to check if the value has been set.
func (o *SmbStorageSpec) GetShareOk() (*SmbRepositoryShareSettingsModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Share, true
}

// SetShare sets field value
func (o *SmbStorageSpec) SetShare(v SmbRepositoryShareSettingsModel) {
	o.Share = v
}

// GetRepository returns the Repository field value
func (o *SmbStorageSpec) GetRepository() NetworkRepositorySettingsModel {
	if o == nil {
		var ret NetworkRepositorySettingsModel
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *SmbStorageSpec) GetRepositoryOk() (*NetworkRepositorySettingsModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *SmbStorageSpec) SetRepository(v NetworkRepositorySettingsModel) {
	o.Repository = v
}

// GetMountServer returns the MountServer field value
func (o *SmbStorageSpec) GetMountServer() MountServerSettingsModel {
	if o == nil {
		var ret MountServerSettingsModel
		return ret
	}

	return o.MountServer
}

// GetMountServerOk returns a tuple with the MountServer field value
// and a boolean to check if the value has been set.
func (o *SmbStorageSpec) GetMountServerOk() (*MountServerSettingsModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MountServer, true
}

// SetMountServer sets field value
func (o *SmbStorageSpec) SetMountServer(v MountServerSettingsModel) {
	o.MountServer = v
}

func (o SmbStorageSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedRepositorySpec, errRepositorySpec := json.Marshal(o.RepositorySpec)
	if errRepositorySpec != nil {
		return []byte{}, errRepositorySpec
	}
	errRepositorySpec = json.Unmarshal([]byte(serializedRepositorySpec), &toSerialize)
	if errRepositorySpec != nil {
		return []byte{}, errRepositorySpec
	}
	if true {
		toSerialize["share"] = o.Share
	}
	if true {
		toSerialize["repository"] = o.Repository
	}
	if true {
		toSerialize["mountServer"] = o.MountServer
	}
	return json.Marshal(toSerialize)
}

type NullableSmbStorageSpec struct {
	value *SmbStorageSpec
	isSet bool
}

func (v NullableSmbStorageSpec) Get() *SmbStorageSpec {
	return v.value
}

func (v *NullableSmbStorageSpec) Set(val *SmbStorageSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableSmbStorageSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSmbStorageSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmbStorageSpec(val *SmbStorageSpec) *NullableSmbStorageSpec {
	return &NullableSmbStorageSpec{value: val, isSet: true}
}

func (v NullableSmbStorageSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmbStorageSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


