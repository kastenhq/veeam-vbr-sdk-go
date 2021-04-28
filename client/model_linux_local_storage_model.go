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

// LinuxLocalStorageModel Direct attached storage.
type LinuxLocalStorageModel struct {
	RepositoryModel
	// ID of the server that is used as a backup repository.
	HostId string `json:"hostId"`
	Repository LinuxLocalRepositorySettingsModel `json:"repository"`
	MountServer MountServerSettingsModel `json:"mountServer"`
}

// NewLinuxLocalStorageModel instantiates a new LinuxLocalStorageModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinuxLocalStorageModel(hostId string, repository LinuxLocalRepositorySettingsModel, mountServer MountServerSettingsModel, id string, name string, description string, type_ ERepositoryType) *LinuxLocalStorageModel {
	this := LinuxLocalStorageModel{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Type = type_
	this.HostId = hostId
	this.Repository = repository
	this.MountServer = mountServer
	return &this
}

// NewLinuxLocalStorageModelWithDefaults instantiates a new LinuxLocalStorageModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinuxLocalStorageModelWithDefaults() *LinuxLocalStorageModel {
	this := LinuxLocalStorageModel{}
	return &this
}

// GetHostId returns the HostId field value
func (o *LinuxLocalStorageModel) GetHostId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value
// and a boolean to check if the value has been set.
func (o *LinuxLocalStorageModel) GetHostIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HostId, true
}

// SetHostId sets field value
func (o *LinuxLocalStorageModel) SetHostId(v string) {
	o.HostId = v
}

// GetRepository returns the Repository field value
func (o *LinuxLocalStorageModel) GetRepository() LinuxLocalRepositorySettingsModel {
	if o == nil {
		var ret LinuxLocalRepositorySettingsModel
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *LinuxLocalStorageModel) GetRepositoryOk() (*LinuxLocalRepositorySettingsModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *LinuxLocalStorageModel) SetRepository(v LinuxLocalRepositorySettingsModel) {
	o.Repository = v
}

// GetMountServer returns the MountServer field value
func (o *LinuxLocalStorageModel) GetMountServer() MountServerSettingsModel {
	if o == nil {
		var ret MountServerSettingsModel
		return ret
	}

	return o.MountServer
}

// GetMountServerOk returns a tuple with the MountServer field value
// and a boolean to check if the value has been set.
func (o *LinuxLocalStorageModel) GetMountServerOk() (*MountServerSettingsModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MountServer, true
}

// SetMountServer sets field value
func (o *LinuxLocalStorageModel) SetMountServer(v MountServerSettingsModel) {
	o.MountServer = v
}

func (o LinuxLocalStorageModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedRepositoryModel, errRepositoryModel := json.Marshal(o.RepositoryModel)
	if errRepositoryModel != nil {
		return []byte{}, errRepositoryModel
	}
	errRepositoryModel = json.Unmarshal([]byte(serializedRepositoryModel), &toSerialize)
	if errRepositoryModel != nil {
		return []byte{}, errRepositoryModel
	}
	if true {
		toSerialize["hostId"] = o.HostId
	}
	if true {
		toSerialize["repository"] = o.Repository
	}
	if true {
		toSerialize["mountServer"] = o.MountServer
	}
	return json.Marshal(toSerialize)
}

type NullableLinuxLocalStorageModel struct {
	value *LinuxLocalStorageModel
	isSet bool
}

func (v NullableLinuxLocalStorageModel) Get() *LinuxLocalStorageModel {
	return v.value
}

func (v *NullableLinuxLocalStorageModel) Set(val *LinuxLocalStorageModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLinuxLocalStorageModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLinuxLocalStorageModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinuxLocalStorageModel(val *LinuxLocalStorageModel) *NullableLinuxLocalStorageModel {
	return &NullableLinuxLocalStorageModel{value: val, isSet: true}
}

func (v NullableLinuxLocalStorageModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinuxLocalStorageModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


