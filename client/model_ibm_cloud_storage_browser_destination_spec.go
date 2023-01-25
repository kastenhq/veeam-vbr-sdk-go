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

// IBMCloudStorageBrowserDestinationSpec struct for IBMCloudStorageBrowserDestinationSpec
type IBMCloudStorageBrowserDestinationSpec struct {
	CloudBrowserNewFolderSpec
	// ID of a server you want to use to connect to the object storage. You can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used.
	HostId *string `json:"hostId,omitempty"`
	// Endpoint address and port number of the IBM Cloud object storage.
	ConnectionPoint string `json:"connectionPoint"`
	// Region where the bucket is located.
	RegionId string `json:"regionId"`
	// Name of the bucket where you want to store your backup data.
	BucketName string `json:"bucketName"`
}

// NewIBMCloudStorageBrowserDestinationSpec instantiates a new IBMCloudStorageBrowserDestinationSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIBMCloudStorageBrowserDestinationSpec(connectionPoint string, regionId string, bucketName string) *IBMCloudStorageBrowserDestinationSpec {
	this := IBMCloudStorageBrowserDestinationSpec{}
	this.RegionId = regionId
	this.BucketName = bucketName
	this.ConnectionPoint = connectionPoint
	return &this
}

// NewIBMCloudStorageBrowserDestinationSpecWithDefaults instantiates a new IBMCloudStorageBrowserDestinationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIBMCloudStorageBrowserDestinationSpecWithDefaults() *IBMCloudStorageBrowserDestinationSpec {
	this := IBMCloudStorageBrowserDestinationSpec{}
	return &this
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *IBMCloudStorageBrowserDestinationSpec) GetHostId() string {
	if o == nil || isNil(o.HostId) {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageBrowserDestinationSpec) GetHostIdOk() (*string, bool) {
	if o == nil || isNil(o.HostId) {
    return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *IBMCloudStorageBrowserDestinationSpec) HasHostId() bool {
	if o != nil && !isNil(o.HostId) {
		return true
	}

	return false
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *IBMCloudStorageBrowserDestinationSpec) SetHostId(v string) {
	o.HostId = &v
}

// GetConnectionPoint returns the ConnectionPoint field value
func (o *IBMCloudStorageBrowserDestinationSpec) GetConnectionPoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionPoint
}

// GetConnectionPointOk returns a tuple with the ConnectionPoint field value
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageBrowserDestinationSpec) GetConnectionPointOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConnectionPoint, true
}

// SetConnectionPoint sets field value
func (o *IBMCloudStorageBrowserDestinationSpec) SetConnectionPoint(v string) {
	o.ConnectionPoint = v
}

// GetRegionId returns the RegionId field value
func (o *IBMCloudStorageBrowserDestinationSpec) GetRegionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageBrowserDestinationSpec) GetRegionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RegionId, true
}

// SetRegionId sets field value
func (o *IBMCloudStorageBrowserDestinationSpec) SetRegionId(v string) {
	o.RegionId = v
}

// GetBucketName returns the BucketName field value
func (o *IBMCloudStorageBrowserDestinationSpec) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageBrowserDestinationSpec) GetBucketNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value
func (o *IBMCloudStorageBrowserDestinationSpec) SetBucketName(v string) {
	o.BucketName = v
}

func (o IBMCloudStorageBrowserDestinationSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBrowserNewFolderSpec, errCloudBrowserNewFolderSpec := json.Marshal(o.CloudBrowserNewFolderSpec)
	if errCloudBrowserNewFolderSpec != nil {
		return []byte{}, errCloudBrowserNewFolderSpec
	}
	errCloudBrowserNewFolderSpec = json.Unmarshal([]byte(serializedCloudBrowserNewFolderSpec), &toSerialize)
	if errCloudBrowserNewFolderSpec != nil {
		return []byte{}, errCloudBrowserNewFolderSpec
	}
	if !isNil(o.HostId) {
		toSerialize["hostId"] = o.HostId
	}
	if true {
		toSerialize["connectionPoint"] = o.ConnectionPoint
	}
	if true {
		toSerialize["regionId"] = o.RegionId
	}
	if true {
		toSerialize["bucketName"] = o.BucketName
	}
	return json.Marshal(toSerialize)
}

type NullableIBMCloudStorageBrowserDestinationSpec struct {
	value *IBMCloudStorageBrowserDestinationSpec
	isSet bool
}

func (v NullableIBMCloudStorageBrowserDestinationSpec) Get() *IBMCloudStorageBrowserDestinationSpec {
	return v.value
}

func (v *NullableIBMCloudStorageBrowserDestinationSpec) Set(val *IBMCloudStorageBrowserDestinationSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableIBMCloudStorageBrowserDestinationSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableIBMCloudStorageBrowserDestinationSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIBMCloudStorageBrowserDestinationSpec(val *IBMCloudStorageBrowserDestinationSpec) *NullableIBMCloudStorageBrowserDestinationSpec {
	return &NullableIBMCloudStorageBrowserDestinationSpec{value: val, isSet: true}
}

func (v NullableIBMCloudStorageBrowserDestinationSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIBMCloudStorageBrowserDestinationSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


