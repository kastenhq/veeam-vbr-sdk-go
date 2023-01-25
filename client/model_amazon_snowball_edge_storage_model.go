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

// AmazonSnowballEdgeStorageModel AWS Snowball Edge storage.
type AmazonSnowballEdgeStorageModel struct {
	RepositoryModel
	// If *true*, the maximum number of concurrent tasks is limited.
	EnableTaskLimit *bool `json:"enableTaskLimit,omitempty"`
	// Maximum number of concurrent tasks.
	MaxTaskCount *int32 `json:"maxTaskCount,omitempty"`
	Account AmazonSnowballEdgeStorageAccountModel `json:"account"`
	Bucket AmazonSnowballEdgeStorageBucketModel `json:"bucket"`
	MountServer MountServerSettingsModel `json:"mountServer"`
}

// NewAmazonSnowballEdgeStorageModel instantiates a new AmazonSnowballEdgeStorageModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonSnowballEdgeStorageModel(account AmazonSnowballEdgeStorageAccountModel, bucket AmazonSnowballEdgeStorageBucketModel, mountServer MountServerSettingsModel, hostId string, repository LinuxHardenedRepositorySettingsModel, share SmbRepositoryShareSettingsModel, container AzureArchiveStorageContainerModel, proxyAppliance S3CompatibleProxyModel) *AmazonSnowballEdgeStorageModel {
	this := AmazonSnowballEdgeStorageModel{}
	this.HostId = hostId
	this.Repository = repository
	this.MountServer = mountServer
	this.Share = share
	this.Account = account
	this.Container = container
	this.ProxyAppliance = proxyAppliance
	this.Bucket = bucket
	return &this
}

// NewAmazonSnowballEdgeStorageModelWithDefaults instantiates a new AmazonSnowballEdgeStorageModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonSnowballEdgeStorageModelWithDefaults() *AmazonSnowballEdgeStorageModel {
	this := AmazonSnowballEdgeStorageModel{}
	return &this
}

// GetEnableTaskLimit returns the EnableTaskLimit field value if set, zero value otherwise.
func (o *AmazonSnowballEdgeStorageModel) GetEnableTaskLimit() bool {
	if o == nil || isNil(o.EnableTaskLimit) {
		var ret bool
		return ret
	}
	return *o.EnableTaskLimit
}

// GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeStorageModel) GetEnableTaskLimitOk() (*bool, bool) {
	if o == nil || isNil(o.EnableTaskLimit) {
    return nil, false
	}
	return o.EnableTaskLimit, true
}

// HasEnableTaskLimit returns a boolean if a field has been set.
func (o *AmazonSnowballEdgeStorageModel) HasEnableTaskLimit() bool {
	if o != nil && !isNil(o.EnableTaskLimit) {
		return true
	}

	return false
}

// SetEnableTaskLimit gets a reference to the given bool and assigns it to the EnableTaskLimit field.
func (o *AmazonSnowballEdgeStorageModel) SetEnableTaskLimit(v bool) {
	o.EnableTaskLimit = &v
}

// GetMaxTaskCount returns the MaxTaskCount field value if set, zero value otherwise.
func (o *AmazonSnowballEdgeStorageModel) GetMaxTaskCount() int32 {
	if o == nil || isNil(o.MaxTaskCount) {
		var ret int32
		return ret
	}
	return *o.MaxTaskCount
}

// GetMaxTaskCountOk returns a tuple with the MaxTaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeStorageModel) GetMaxTaskCountOk() (*int32, bool) {
	if o == nil || isNil(o.MaxTaskCount) {
    return nil, false
	}
	return o.MaxTaskCount, true
}

// HasMaxTaskCount returns a boolean if a field has been set.
func (o *AmazonSnowballEdgeStorageModel) HasMaxTaskCount() bool {
	if o != nil && !isNil(o.MaxTaskCount) {
		return true
	}

	return false
}

// SetMaxTaskCount gets a reference to the given int32 and assigns it to the MaxTaskCount field.
func (o *AmazonSnowballEdgeStorageModel) SetMaxTaskCount(v int32) {
	o.MaxTaskCount = &v
}

// GetAccount returns the Account field value
func (o *AmazonSnowballEdgeStorageModel) GetAccount() AmazonSnowballEdgeStorageAccountModel {
	if o == nil {
		var ret AmazonSnowballEdgeStorageAccountModel
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeStorageModel) GetAccountOk() (*AmazonSnowballEdgeStorageAccountModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *AmazonSnowballEdgeStorageModel) SetAccount(v AmazonSnowballEdgeStorageAccountModel) {
	o.Account = v
}

// GetBucket returns the Bucket field value
func (o *AmazonSnowballEdgeStorageModel) GetBucket() AmazonSnowballEdgeStorageBucketModel {
	if o == nil {
		var ret AmazonSnowballEdgeStorageBucketModel
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeStorageModel) GetBucketOk() (*AmazonSnowballEdgeStorageBucketModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *AmazonSnowballEdgeStorageModel) SetBucket(v AmazonSnowballEdgeStorageBucketModel) {
	o.Bucket = v
}

// GetMountServer returns the MountServer field value
func (o *AmazonSnowballEdgeStorageModel) GetMountServer() MountServerSettingsModel {
	if o == nil {
		var ret MountServerSettingsModel
		return ret
	}

	return o.MountServer
}

// GetMountServerOk returns a tuple with the MountServer field value
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeStorageModel) GetMountServerOk() (*MountServerSettingsModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MountServer, true
}

// SetMountServer sets field value
func (o *AmazonSnowballEdgeStorageModel) SetMountServer(v MountServerSettingsModel) {
	o.MountServer = v
}

func (o AmazonSnowballEdgeStorageModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedRepositoryModel, errRepositoryModel := json.Marshal(o.RepositoryModel)
	if errRepositoryModel != nil {
		return []byte{}, errRepositoryModel
	}
	errRepositoryModel = json.Unmarshal([]byte(serializedRepositoryModel), &toSerialize)
	if errRepositoryModel != nil {
		return []byte{}, errRepositoryModel
	}
	if !isNil(o.EnableTaskLimit) {
		toSerialize["enableTaskLimit"] = o.EnableTaskLimit
	}
	if !isNil(o.MaxTaskCount) {
		toSerialize["maxTaskCount"] = o.MaxTaskCount
	}
	if true {
		toSerialize["account"] = o.Account
	}
	if true {
		toSerialize["bucket"] = o.Bucket
	}
	if true {
		toSerialize["mountServer"] = o.MountServer
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonSnowballEdgeStorageModel struct {
	value *AmazonSnowballEdgeStorageModel
	isSet bool
}

func (v NullableAmazonSnowballEdgeStorageModel) Get() *AmazonSnowballEdgeStorageModel {
	return v.value
}

func (v *NullableAmazonSnowballEdgeStorageModel) Set(val *AmazonSnowballEdgeStorageModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonSnowballEdgeStorageModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonSnowballEdgeStorageModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonSnowballEdgeStorageModel(val *AmazonSnowballEdgeStorageModel) *NullableAmazonSnowballEdgeStorageModel {
	return &NullableAmazonSnowballEdgeStorageModel{value: val, isSet: true}
}

func (v NullableAmazonSnowballEdgeStorageModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonSnowballEdgeStorageModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


