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

// AzureDataBoxStorageAccountModel Account used to access the Azure Data Box storage.
type AzureDataBoxStorageAccountModel struct {
	// Service endpoint address of the Azure Data Box device.
	ServiceEndpoint string `json:"serviceEndpoint"`
	// ID of the cloud credentials record.
	CredentialsId string `json:"credentialsId"`
	ConnectionSettings *ObjectStorageConnectionModel `json:"connectionSettings,omitempty"`
}

// NewAzureDataBoxStorageAccountModel instantiates a new AzureDataBoxStorageAccountModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureDataBoxStorageAccountModel(serviceEndpoint string, credentialsId string) *AzureDataBoxStorageAccountModel {
	this := AzureDataBoxStorageAccountModel{}
	this.ServiceEndpoint = serviceEndpoint
	this.CredentialsId = credentialsId
	return &this
}

// NewAzureDataBoxStorageAccountModelWithDefaults instantiates a new AzureDataBoxStorageAccountModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureDataBoxStorageAccountModelWithDefaults() *AzureDataBoxStorageAccountModel {
	this := AzureDataBoxStorageAccountModel{}
	return &this
}

// GetServiceEndpoint returns the ServiceEndpoint field value
func (o *AzureDataBoxStorageAccountModel) GetServiceEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEndpoint
}

// GetServiceEndpointOk returns a tuple with the ServiceEndpoint field value
// and a boolean to check if the value has been set.
func (o *AzureDataBoxStorageAccountModel) GetServiceEndpointOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServiceEndpoint, true
}

// SetServiceEndpoint sets field value
func (o *AzureDataBoxStorageAccountModel) SetServiceEndpoint(v string) {
	o.ServiceEndpoint = v
}

// GetCredentialsId returns the CredentialsId field value
func (o *AzureDataBoxStorageAccountModel) GetCredentialsId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CredentialsId
}

// GetCredentialsIdOk returns a tuple with the CredentialsId field value
// and a boolean to check if the value has been set.
func (o *AzureDataBoxStorageAccountModel) GetCredentialsIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CredentialsId, true
}

// SetCredentialsId sets field value
func (o *AzureDataBoxStorageAccountModel) SetCredentialsId(v string) {
	o.CredentialsId = v
}

// GetConnectionSettings returns the ConnectionSettings field value if set, zero value otherwise.
func (o *AzureDataBoxStorageAccountModel) GetConnectionSettings() ObjectStorageConnectionModel {
	if o == nil || isNil(o.ConnectionSettings) {
		var ret ObjectStorageConnectionModel
		return ret
	}
	return *o.ConnectionSettings
}

// GetConnectionSettingsOk returns a tuple with the ConnectionSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureDataBoxStorageAccountModel) GetConnectionSettingsOk() (*ObjectStorageConnectionModel, bool) {
	if o == nil || isNil(o.ConnectionSettings) {
    return nil, false
	}
	return o.ConnectionSettings, true
}

// HasConnectionSettings returns a boolean if a field has been set.
func (o *AzureDataBoxStorageAccountModel) HasConnectionSettings() bool {
	if o != nil && !isNil(o.ConnectionSettings) {
		return true
	}

	return false
}

// SetConnectionSettings gets a reference to the given ObjectStorageConnectionModel and assigns it to the ConnectionSettings field.
func (o *AzureDataBoxStorageAccountModel) SetConnectionSettings(v ObjectStorageConnectionModel) {
	o.ConnectionSettings = &v
}

func (o AzureDataBoxStorageAccountModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serviceEndpoint"] = o.ServiceEndpoint
	}
	if true {
		toSerialize["credentialsId"] = o.CredentialsId
	}
	if !isNil(o.ConnectionSettings) {
		toSerialize["connectionSettings"] = o.ConnectionSettings
	}
	return json.Marshal(toSerialize)
}

type NullableAzureDataBoxStorageAccountModel struct {
	value *AzureDataBoxStorageAccountModel
	isSet bool
}

func (v NullableAzureDataBoxStorageAccountModel) Get() *AzureDataBoxStorageAccountModel {
	return v.value
}

func (v *NullableAzureDataBoxStorageAccountModel) Set(val *AzureDataBoxStorageAccountModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureDataBoxStorageAccountModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureDataBoxStorageAccountModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureDataBoxStorageAccountModel(val *AzureDataBoxStorageAccountModel) *NullableAzureDataBoxStorageAccountModel {
	return &NullableAzureDataBoxStorageAccountModel{value: val, isSet: true}
}

func (v NullableAzureDataBoxStorageAccountModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureDataBoxStorageAccountModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


