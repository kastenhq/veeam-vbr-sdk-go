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

// GoogleCloudStorageBucketModel Google Cloud object storage bucket where backup data is stored.
type GoogleCloudStorageBucketModel struct {
	// Name of a Google Cloud bucket.
	BucketName string `json:"bucketName"`
	// Name of the folder that the object storage repository is mapped to.
	FolderName string `json:"folderName"`
	StorageConsumptionLimit *ObjectStorageConsumptionLimitModel `json:"storageConsumptionLimit,omitempty"`
	// If *true*, the nearline storage class is used.
	NearlineStorageEnabled *bool `json:"nearlineStorageEnabled,omitempty"`
}

// NewGoogleCloudStorageBucketModel instantiates a new GoogleCloudStorageBucketModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleCloudStorageBucketModel(bucketName string, folderName string) *GoogleCloudStorageBucketModel {
	this := GoogleCloudStorageBucketModel{}
	this.BucketName = bucketName
	this.FolderName = folderName
	return &this
}

// NewGoogleCloudStorageBucketModelWithDefaults instantiates a new GoogleCloudStorageBucketModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudStorageBucketModelWithDefaults() *GoogleCloudStorageBucketModel {
	this := GoogleCloudStorageBucketModel{}
	return &this
}

// GetBucketName returns the BucketName field value
func (o *GoogleCloudStorageBucketModel) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *GoogleCloudStorageBucketModel) GetBucketNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value
func (o *GoogleCloudStorageBucketModel) SetBucketName(v string) {
	o.BucketName = v
}

// GetFolderName returns the FolderName field value
func (o *GoogleCloudStorageBucketModel) GetFolderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FolderName
}

// GetFolderNameOk returns a tuple with the FolderName field value
// and a boolean to check if the value has been set.
func (o *GoogleCloudStorageBucketModel) GetFolderNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FolderName, true
}

// SetFolderName sets field value
func (o *GoogleCloudStorageBucketModel) SetFolderName(v string) {
	o.FolderName = v
}

// GetStorageConsumptionLimit returns the StorageConsumptionLimit field value if set, zero value otherwise.
func (o *GoogleCloudStorageBucketModel) GetStorageConsumptionLimit() ObjectStorageConsumptionLimitModel {
	if o == nil || isNil(o.StorageConsumptionLimit) {
		var ret ObjectStorageConsumptionLimitModel
		return ret
	}
	return *o.StorageConsumptionLimit
}

// GetStorageConsumptionLimitOk returns a tuple with the StorageConsumptionLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudStorageBucketModel) GetStorageConsumptionLimitOk() (*ObjectStorageConsumptionLimitModel, bool) {
	if o == nil || isNil(o.StorageConsumptionLimit) {
    return nil, false
	}
	return o.StorageConsumptionLimit, true
}

// HasStorageConsumptionLimit returns a boolean if a field has been set.
func (o *GoogleCloudStorageBucketModel) HasStorageConsumptionLimit() bool {
	if o != nil && !isNil(o.StorageConsumptionLimit) {
		return true
	}

	return false
}

// SetStorageConsumptionLimit gets a reference to the given ObjectStorageConsumptionLimitModel and assigns it to the StorageConsumptionLimit field.
func (o *GoogleCloudStorageBucketModel) SetStorageConsumptionLimit(v ObjectStorageConsumptionLimitModel) {
	o.StorageConsumptionLimit = &v
}

// GetNearlineStorageEnabled returns the NearlineStorageEnabled field value if set, zero value otherwise.
func (o *GoogleCloudStorageBucketModel) GetNearlineStorageEnabled() bool {
	if o == nil || isNil(o.NearlineStorageEnabled) {
		var ret bool
		return ret
	}
	return *o.NearlineStorageEnabled
}

// GetNearlineStorageEnabledOk returns a tuple with the NearlineStorageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudStorageBucketModel) GetNearlineStorageEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.NearlineStorageEnabled) {
    return nil, false
	}
	return o.NearlineStorageEnabled, true
}

// HasNearlineStorageEnabled returns a boolean if a field has been set.
func (o *GoogleCloudStorageBucketModel) HasNearlineStorageEnabled() bool {
	if o != nil && !isNil(o.NearlineStorageEnabled) {
		return true
	}

	return false
}

// SetNearlineStorageEnabled gets a reference to the given bool and assigns it to the NearlineStorageEnabled field.
func (o *GoogleCloudStorageBucketModel) SetNearlineStorageEnabled(v bool) {
	o.NearlineStorageEnabled = &v
}

func (o GoogleCloudStorageBucketModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bucketName"] = o.BucketName
	}
	if true {
		toSerialize["folderName"] = o.FolderName
	}
	if !isNil(o.StorageConsumptionLimit) {
		toSerialize["storageConsumptionLimit"] = o.StorageConsumptionLimit
	}
	if !isNil(o.NearlineStorageEnabled) {
		toSerialize["nearlineStorageEnabled"] = o.NearlineStorageEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleCloudStorageBucketModel struct {
	value *GoogleCloudStorageBucketModel
	isSet bool
}

func (v NullableGoogleCloudStorageBucketModel) Get() *GoogleCloudStorageBucketModel {
	return v.value
}

func (v *NullableGoogleCloudStorageBucketModel) Set(val *GoogleCloudStorageBucketModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleCloudStorageBucketModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleCloudStorageBucketModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleCloudStorageBucketModel(val *GoogleCloudStorageBucketModel) *NullableGoogleCloudStorageBucketModel {
	return &NullableGoogleCloudStorageBucketModel{value: val, isSet: true}
}

func (v NullableGoogleCloudStorageBucketModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleCloudStorageBucketModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


