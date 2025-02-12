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

// GoogleCloudStorageBrowserSpec struct for GoogleCloudStorageBrowserSpec
type GoogleCloudStorageBrowserSpec struct {
	CloudBrowserSpec
	// ID of a gateway server you want to use to connect to the object storage. Specify this parameter to check internet connection of the server. As a gateway server you can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used.
	GatewayServerId *string `json:"gatewayServerId,omitempty"`
	Filters *GoogleCloudStorageBrowserFilters `json:"filters,omitempty"`
}

// NewGoogleCloudStorageBrowserSpec instantiates a new GoogleCloudStorageBrowserSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleCloudStorageBrowserSpec(credentialsId string, serviceType ECloudServiceType) *GoogleCloudStorageBrowserSpec {
	this := GoogleCloudStorageBrowserSpec{}
	this.CredentialsId = credentialsId
	this.ServiceType = serviceType
	return &this
}

// NewGoogleCloudStorageBrowserSpecWithDefaults instantiates a new GoogleCloudStorageBrowserSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudStorageBrowserSpecWithDefaults() *GoogleCloudStorageBrowserSpec {
	this := GoogleCloudStorageBrowserSpec{}
	return &this
}

// GetGatewayServerId returns the GatewayServerId field value if set, zero value otherwise.
func (o *GoogleCloudStorageBrowserSpec) GetGatewayServerId() string {
	if o == nil || isNil(o.GatewayServerId) {
		var ret string
		return ret
	}
	return *o.GatewayServerId
}

// GetGatewayServerIdOk returns a tuple with the GatewayServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudStorageBrowserSpec) GetGatewayServerIdOk() (*string, bool) {
	if o == nil || isNil(o.GatewayServerId) {
    return nil, false
	}
	return o.GatewayServerId, true
}

// HasGatewayServerId returns a boolean if a field has been set.
func (o *GoogleCloudStorageBrowserSpec) HasGatewayServerId() bool {
	if o != nil && !isNil(o.GatewayServerId) {
		return true
	}

	return false
}

// SetGatewayServerId gets a reference to the given string and assigns it to the GatewayServerId field.
func (o *GoogleCloudStorageBrowserSpec) SetGatewayServerId(v string) {
	o.GatewayServerId = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *GoogleCloudStorageBrowserSpec) GetFilters() GoogleCloudStorageBrowserFilters {
	if o == nil || isNil(o.Filters) {
		var ret GoogleCloudStorageBrowserFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudStorageBrowserSpec) GetFiltersOk() (*GoogleCloudStorageBrowserFilters, bool) {
	if o == nil || isNil(o.Filters) {
    return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *GoogleCloudStorageBrowserSpec) HasFilters() bool {
	if o != nil && !isNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given GoogleCloudStorageBrowserFilters and assigns it to the Filters field.
func (o *GoogleCloudStorageBrowserSpec) SetFilters(v GoogleCloudStorageBrowserFilters) {
	o.Filters = &v
}

func (o GoogleCloudStorageBrowserSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBrowserSpec, errCloudBrowserSpec := json.Marshal(o.CloudBrowserSpec)
	if errCloudBrowserSpec != nil {
		return []byte{}, errCloudBrowserSpec
	}
	errCloudBrowserSpec = json.Unmarshal([]byte(serializedCloudBrowserSpec), &toSerialize)
	if errCloudBrowserSpec != nil {
		return []byte{}, errCloudBrowserSpec
	}
	if !isNil(o.GatewayServerId) {
		toSerialize["gatewayServerId"] = o.GatewayServerId
	}
	if !isNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleCloudStorageBrowserSpec struct {
	value *GoogleCloudStorageBrowserSpec
	isSet bool
}

func (v NullableGoogleCloudStorageBrowserSpec) Get() *GoogleCloudStorageBrowserSpec {
	return v.value
}

func (v *NullableGoogleCloudStorageBrowserSpec) Set(val *GoogleCloudStorageBrowserSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleCloudStorageBrowserSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleCloudStorageBrowserSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleCloudStorageBrowserSpec(val *GoogleCloudStorageBrowserSpec) *NullableGoogleCloudStorageBrowserSpec {
	return &NullableGoogleCloudStorageBrowserSpec{value: val, isSet: true}
}

func (v NullableGoogleCloudStorageBrowserSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleCloudStorageBrowserSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


