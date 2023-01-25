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

// AzureArchiveStorageSpec Microsoft Azure Archive storage.
type AzureArchiveStorageSpec struct {
	RepositorySpec
	Account AzureArchiveStorageAccountModel `json:"account"`
	Container AzureArchiveStorageContainerModel `json:"container"`
	ProxyAppliance AzureStorageProxyModel `json:"proxyAppliance"`
}

// NewAzureArchiveStorageSpec instantiates a new AzureArchiveStorageSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureArchiveStorageSpec(account AzureArchiveStorageAccountModel, container AzureArchiveStorageContainerModel, proxyAppliance AzureStorageProxyModel) *AzureArchiveStorageSpec {
	this := AzureArchiveStorageSpec{}
	this.Account = account
	this.Container = container
	this.ProxyAppliance = proxyAppliance
	return &this
}

// NewAzureArchiveStorageSpecWithDefaults instantiates a new AzureArchiveStorageSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureArchiveStorageSpecWithDefaults() *AzureArchiveStorageSpec {
	this := AzureArchiveStorageSpec{}
	return &this
}

// GetAccount returns the Account field value
func (o *AzureArchiveStorageSpec) GetAccount() AzureArchiveStorageAccountModel {
	if o == nil {
		var ret AzureArchiveStorageAccountModel
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *AzureArchiveStorageSpec) GetAccountOk() (*AzureArchiveStorageAccountModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *AzureArchiveStorageSpec) SetAccount(v AzureArchiveStorageAccountModel) {
	o.Account = v
}

// GetContainer returns the Container field value
func (o *AzureArchiveStorageSpec) GetContainer() AzureArchiveStorageContainerModel {
	if o == nil {
		var ret AzureArchiveStorageContainerModel
		return ret
	}

	return o.Container
}

// GetContainerOk returns a tuple with the Container field value
// and a boolean to check if the value has been set.
func (o *AzureArchiveStorageSpec) GetContainerOk() (*AzureArchiveStorageContainerModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Container, true
}

// SetContainer sets field value
func (o *AzureArchiveStorageSpec) SetContainer(v AzureArchiveStorageContainerModel) {
	o.Container = v
}

// GetProxyAppliance returns the ProxyAppliance field value
func (o *AzureArchiveStorageSpec) GetProxyAppliance() AzureStorageProxyModel {
	if o == nil {
		var ret AzureStorageProxyModel
		return ret
	}

	return o.ProxyAppliance
}

// GetProxyApplianceOk returns a tuple with the ProxyAppliance field value
// and a boolean to check if the value has been set.
func (o *AzureArchiveStorageSpec) GetProxyApplianceOk() (*AzureStorageProxyModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ProxyAppliance, true
}

// SetProxyAppliance sets field value
func (o *AzureArchiveStorageSpec) SetProxyAppliance(v AzureStorageProxyModel) {
	o.ProxyAppliance = v
}

func (o AzureArchiveStorageSpec) MarshalJSON() ([]byte, error) {
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
		toSerialize["account"] = o.Account
	}
	if true {
		toSerialize["container"] = o.Container
	}
	if true {
		toSerialize["proxyAppliance"] = o.ProxyAppliance
	}
	return json.Marshal(toSerialize)
}

type NullableAzureArchiveStorageSpec struct {
	value *AzureArchiveStorageSpec
	isSet bool
}

func (v NullableAzureArchiveStorageSpec) Get() *AzureArchiveStorageSpec {
	return v.value
}

func (v *NullableAzureArchiveStorageSpec) Set(val *AzureArchiveStorageSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureArchiveStorageSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureArchiveStorageSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureArchiveStorageSpec(val *AzureArchiveStorageSpec) *NullableAzureArchiveStorageSpec {
	return &NullableAzureArchiveStorageSpec{value: val, isSet: true}
}

func (v NullableAzureArchiveStorageSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureArchiveStorageSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


