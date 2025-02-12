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

// AzureArchiveStorageModel Microsoft Azure Archive storage.
type AzureArchiveStorageModel struct {
	RepositoryModel
	Account AzureArchiveStorageAccountModel `json:"account"`
	Container AzureArchiveStorageContainerModel `json:"container"`
	ProxyAppliance AzureStorageProxyModel `json:"proxyAppliance"`
}

// NewAzureArchiveStorageModel instantiates a new AzureArchiveStorageModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureArchiveStorageModel(account AzureArchiveStorageAccountModel, container AzureArchiveStorageContainerModel, proxyAppliance AzureStorageProxyModel, id string, name string, description string, type_ ERepositoryType) *AzureArchiveStorageModel {
	this := AzureArchiveStorageModel{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Type = type_
	this.Account = account
	this.Container = container
	this.ProxyAppliance = proxyAppliance
	return &this
}

// NewAzureArchiveStorageModelWithDefaults instantiates a new AzureArchiveStorageModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureArchiveStorageModelWithDefaults() *AzureArchiveStorageModel {
	this := AzureArchiveStorageModel{}
	return &this
}

// GetAccount returns the Account field value
func (o *AzureArchiveStorageModel) GetAccount() AzureArchiveStorageAccountModel {
	if o == nil {
		var ret AzureArchiveStorageAccountModel
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *AzureArchiveStorageModel) GetAccountOk() (*AzureArchiveStorageAccountModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *AzureArchiveStorageModel) SetAccount(v AzureArchiveStorageAccountModel) {
	o.Account = v
}

// GetContainer returns the Container field value
func (o *AzureArchiveStorageModel) GetContainer() AzureArchiveStorageContainerModel {
	if o == nil {
		var ret AzureArchiveStorageContainerModel
		return ret
	}

	return o.Container
}

// GetContainerOk returns a tuple with the Container field value
// and a boolean to check if the value has been set.
func (o *AzureArchiveStorageModel) GetContainerOk() (*AzureArchiveStorageContainerModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Container, true
}

// SetContainer sets field value
func (o *AzureArchiveStorageModel) SetContainer(v AzureArchiveStorageContainerModel) {
	o.Container = v
}

// GetProxyAppliance returns the ProxyAppliance field value
func (o *AzureArchiveStorageModel) GetProxyAppliance() AzureStorageProxyModel {
	if o == nil {
		var ret AzureStorageProxyModel
		return ret
	}

	return o.ProxyAppliance
}

// GetProxyApplianceOk returns a tuple with the ProxyAppliance field value
// and a boolean to check if the value has been set.
func (o *AzureArchiveStorageModel) GetProxyApplianceOk() (*AzureStorageProxyModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ProxyAppliance, true
}

// SetProxyAppliance sets field value
func (o *AzureArchiveStorageModel) SetProxyAppliance(v AzureStorageProxyModel) {
	o.ProxyAppliance = v
}

func (o AzureArchiveStorageModel) MarshalJSON() ([]byte, error) {
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

type NullableAzureArchiveStorageModel struct {
	value *AzureArchiveStorageModel
	isSet bool
}

func (v NullableAzureArchiveStorageModel) Get() *AzureArchiveStorageModel {
	return v.value
}

func (v *NullableAzureArchiveStorageModel) Set(val *AzureArchiveStorageModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureArchiveStorageModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureArchiveStorageModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureArchiveStorageModel(val *AzureArchiveStorageModel) *NullableAzureArchiveStorageModel {
	return &NullableAzureArchiveStorageModel{value: val, isSet: true}
}

func (v NullableAzureArchiveStorageModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureArchiveStorageModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


