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

// AzureBlobStorageSpecAllOf struct for AzureBlobStorageSpecAllOf
type AzureBlobStorageSpecAllOf struct {
	// If *true*, the maximum number of concurrent tasks is limited.
	EnableTaskLimit *bool `json:"enableTaskLimit,omitempty"`
	// Maximum number of concurrent tasks.
	MaxTaskCount *int32 `json:"maxTaskCount,omitempty"`
	Account AzureBlobStorageAccountModel `json:"account"`
	Container AzureBlobStorageContainerModel `json:"container"`
	ProxyAppliance *AzureStorageProxyModel `json:"proxyAppliance,omitempty"`
	MountServer MountServerSettingsModel `json:"mountServer"`
}

// NewAzureBlobStorageSpecAllOf instantiates a new AzureBlobStorageSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureBlobStorageSpecAllOf(account AzureBlobStorageAccountModel, container AzureBlobStorageContainerModel, mountServer MountServerSettingsModel) *AzureBlobStorageSpecAllOf {
	this := AzureBlobStorageSpecAllOf{}
	this.Account = account
	this.Container = container
	this.MountServer = mountServer
	return &this
}

// NewAzureBlobStorageSpecAllOfWithDefaults instantiates a new AzureBlobStorageSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureBlobStorageSpecAllOfWithDefaults() *AzureBlobStorageSpecAllOf {
	this := AzureBlobStorageSpecAllOf{}
	return &this
}

// GetEnableTaskLimit returns the EnableTaskLimit field value if set, zero value otherwise.
func (o *AzureBlobStorageSpecAllOf) GetEnableTaskLimit() bool {
	if o == nil || isNil(o.EnableTaskLimit) {
		var ret bool
		return ret
	}
	return *o.EnableTaskLimit
}

// GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageSpecAllOf) GetEnableTaskLimitOk() (*bool, bool) {
	if o == nil || isNil(o.EnableTaskLimit) {
    return nil, false
	}
	return o.EnableTaskLimit, true
}

// HasEnableTaskLimit returns a boolean if a field has been set.
func (o *AzureBlobStorageSpecAllOf) HasEnableTaskLimit() bool {
	if o != nil && !isNil(o.EnableTaskLimit) {
		return true
	}

	return false
}

// SetEnableTaskLimit gets a reference to the given bool and assigns it to the EnableTaskLimit field.
func (o *AzureBlobStorageSpecAllOf) SetEnableTaskLimit(v bool) {
	o.EnableTaskLimit = &v
}

// GetMaxTaskCount returns the MaxTaskCount field value if set, zero value otherwise.
func (o *AzureBlobStorageSpecAllOf) GetMaxTaskCount() int32 {
	if o == nil || isNil(o.MaxTaskCount) {
		var ret int32
		return ret
	}
	return *o.MaxTaskCount
}

// GetMaxTaskCountOk returns a tuple with the MaxTaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageSpecAllOf) GetMaxTaskCountOk() (*int32, bool) {
	if o == nil || isNil(o.MaxTaskCount) {
    return nil, false
	}
	return o.MaxTaskCount, true
}

// HasMaxTaskCount returns a boolean if a field has been set.
func (o *AzureBlobStorageSpecAllOf) HasMaxTaskCount() bool {
	if o != nil && !isNil(o.MaxTaskCount) {
		return true
	}

	return false
}

// SetMaxTaskCount gets a reference to the given int32 and assigns it to the MaxTaskCount field.
func (o *AzureBlobStorageSpecAllOf) SetMaxTaskCount(v int32) {
	o.MaxTaskCount = &v
}

// GetAccount returns the Account field value
func (o *AzureBlobStorageSpecAllOf) GetAccount() AzureBlobStorageAccountModel {
	if o == nil {
		var ret AzureBlobStorageAccountModel
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageSpecAllOf) GetAccountOk() (*AzureBlobStorageAccountModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *AzureBlobStorageSpecAllOf) SetAccount(v AzureBlobStorageAccountModel) {
	o.Account = v
}

// GetContainer returns the Container field value
func (o *AzureBlobStorageSpecAllOf) GetContainer() AzureBlobStorageContainerModel {
	if o == nil {
		var ret AzureBlobStorageContainerModel
		return ret
	}

	return o.Container
}

// GetContainerOk returns a tuple with the Container field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageSpecAllOf) GetContainerOk() (*AzureBlobStorageContainerModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Container, true
}

// SetContainer sets field value
func (o *AzureBlobStorageSpecAllOf) SetContainer(v AzureBlobStorageContainerModel) {
	o.Container = v
}

// GetProxyAppliance returns the ProxyAppliance field value if set, zero value otherwise.
func (o *AzureBlobStorageSpecAllOf) GetProxyAppliance() AzureStorageProxyModel {
	if o == nil || isNil(o.ProxyAppliance) {
		var ret AzureStorageProxyModel
		return ret
	}
	return *o.ProxyAppliance
}

// GetProxyApplianceOk returns a tuple with the ProxyAppliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageSpecAllOf) GetProxyApplianceOk() (*AzureStorageProxyModel, bool) {
	if o == nil || isNil(o.ProxyAppliance) {
    return nil, false
	}
	return o.ProxyAppliance, true
}

// HasProxyAppliance returns a boolean if a field has been set.
func (o *AzureBlobStorageSpecAllOf) HasProxyAppliance() bool {
	if o != nil && !isNil(o.ProxyAppliance) {
		return true
	}

	return false
}

// SetProxyAppliance gets a reference to the given AzureStorageProxyModel and assigns it to the ProxyAppliance field.
func (o *AzureBlobStorageSpecAllOf) SetProxyAppliance(v AzureStorageProxyModel) {
	o.ProxyAppliance = &v
}

// GetMountServer returns the MountServer field value
func (o *AzureBlobStorageSpecAllOf) GetMountServer() MountServerSettingsModel {
	if o == nil {
		var ret MountServerSettingsModel
		return ret
	}

	return o.MountServer
}

// GetMountServerOk returns a tuple with the MountServer field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageSpecAllOf) GetMountServerOk() (*MountServerSettingsModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MountServer, true
}

// SetMountServer sets field value
func (o *AzureBlobStorageSpecAllOf) SetMountServer(v MountServerSettingsModel) {
	o.MountServer = v
}

func (o AzureBlobStorageSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
		toSerialize["container"] = o.Container
	}
	if !isNil(o.ProxyAppliance) {
		toSerialize["proxyAppliance"] = o.ProxyAppliance
	}
	if true {
		toSerialize["mountServer"] = o.MountServer
	}
	return json.Marshal(toSerialize)
}

type NullableAzureBlobStorageSpecAllOf struct {
	value *AzureBlobStorageSpecAllOf
	isSet bool
}

func (v NullableAzureBlobStorageSpecAllOf) Get() *AzureBlobStorageSpecAllOf {
	return v.value
}

func (v *NullableAzureBlobStorageSpecAllOf) Set(val *AzureBlobStorageSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureBlobStorageSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureBlobStorageSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureBlobStorageSpecAllOf(val *AzureBlobStorageSpecAllOf) *NullableAzureBlobStorageSpecAllOf {
	return &NullableAzureBlobStorageSpecAllOf{value: val, isSet: true}
}

func (v NullableAzureBlobStorageSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureBlobStorageSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


