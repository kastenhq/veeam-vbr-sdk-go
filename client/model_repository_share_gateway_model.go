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

// RepositoryShareGatewayModel Settings for the gateway server.
type RepositoryShareGatewayModel struct {
	// If *true*, Veeam Backup & Replication automatically selects a gateway server.
	AutoSelect bool `json:"autoSelect"`
	// Array of gateway server IDs.
	GatewayServerIds []string `json:"gatewayServerIds,omitempty"`
}

// NewRepositoryShareGatewayModel instantiates a new RepositoryShareGatewayModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryShareGatewayModel(autoSelect bool) *RepositoryShareGatewayModel {
	this := RepositoryShareGatewayModel{}
	this.AutoSelect = autoSelect
	return &this
}

// NewRepositoryShareGatewayModelWithDefaults instantiates a new RepositoryShareGatewayModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryShareGatewayModelWithDefaults() *RepositoryShareGatewayModel {
	this := RepositoryShareGatewayModel{}
	return &this
}

// GetAutoSelect returns the AutoSelect field value
func (o *RepositoryShareGatewayModel) GetAutoSelect() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoSelect
}

// GetAutoSelectOk returns a tuple with the AutoSelect field value
// and a boolean to check if the value has been set.
func (o *RepositoryShareGatewayModel) GetAutoSelectOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AutoSelect, true
}

// SetAutoSelect sets field value
func (o *RepositoryShareGatewayModel) SetAutoSelect(v bool) {
	o.AutoSelect = v
}

// GetGatewayServerIds returns the GatewayServerIds field value if set, zero value otherwise.
func (o *RepositoryShareGatewayModel) GetGatewayServerIds() []string {
	if o == nil || isNil(o.GatewayServerIds) {
		var ret []string
		return ret
	}
	return o.GatewayServerIds
}

// GetGatewayServerIdsOk returns a tuple with the GatewayServerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryShareGatewayModel) GetGatewayServerIdsOk() ([]string, bool) {
	if o == nil || isNil(o.GatewayServerIds) {
    return nil, false
	}
	return o.GatewayServerIds, true
}

// HasGatewayServerIds returns a boolean if a field has been set.
func (o *RepositoryShareGatewayModel) HasGatewayServerIds() bool {
	if o != nil && !isNil(o.GatewayServerIds) {
		return true
	}

	return false
}

// SetGatewayServerIds gets a reference to the given []string and assigns it to the GatewayServerIds field.
func (o *RepositoryShareGatewayModel) SetGatewayServerIds(v []string) {
	o.GatewayServerIds = v
}

func (o RepositoryShareGatewayModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["autoSelect"] = o.AutoSelect
	}
	if !isNil(o.GatewayServerIds) {
		toSerialize["gatewayServerIds"] = o.GatewayServerIds
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryShareGatewayModel struct {
	value *RepositoryShareGatewayModel
	isSet bool
}

func (v NullableRepositoryShareGatewayModel) Get() *RepositoryShareGatewayModel {
	return v.value
}

func (v *NullableRepositoryShareGatewayModel) Set(val *RepositoryShareGatewayModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryShareGatewayModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryShareGatewayModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryShareGatewayModel(val *RepositoryShareGatewayModel) *NullableRepositoryShareGatewayModel {
	return &NullableRepositoryShareGatewayModel{value: val, isSet: true}
}

func (v NullableRepositoryShareGatewayModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryShareGatewayModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


