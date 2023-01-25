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

// EncryptionPasswordSpec struct for EncryptionPasswordSpec
type EncryptionPasswordSpec struct {
	// Password for data encryption. If you lose the password, you will not be able to restore it.
	Password string `json:"password"`
	// Hint for the encryption password. Provide a meaningful hint that will help you recall the password.
	Hint string `json:"hint"`
	// Tag for the encryption password.
	Tag *string `json:"tag,omitempty"`
}

// NewEncryptionPasswordSpec instantiates a new EncryptionPasswordSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptionPasswordSpec(password string, hint string) *EncryptionPasswordSpec {
	this := EncryptionPasswordSpec{}
	this.Password = password
	this.Hint = hint
	return &this
}

// NewEncryptionPasswordSpecWithDefaults instantiates a new EncryptionPasswordSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptionPasswordSpecWithDefaults() *EncryptionPasswordSpec {
	this := EncryptionPasswordSpec{}
	return &this
}

// GetPassword returns the Password field value
func (o *EncryptionPasswordSpec) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *EncryptionPasswordSpec) GetPasswordOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *EncryptionPasswordSpec) SetPassword(v string) {
	o.Password = v
}

// GetHint returns the Hint field value
func (o *EncryptionPasswordSpec) GetHint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hint
}

// GetHintOk returns a tuple with the Hint field value
// and a boolean to check if the value has been set.
func (o *EncryptionPasswordSpec) GetHintOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Hint, true
}

// SetHint sets field value
func (o *EncryptionPasswordSpec) SetHint(v string) {
	o.Hint = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *EncryptionPasswordSpec) GetTag() string {
	if o == nil || isNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionPasswordSpec) GetTagOk() (*string, bool) {
	if o == nil || isNil(o.Tag) {
    return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *EncryptionPasswordSpec) HasTag() bool {
	if o != nil && !isNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *EncryptionPasswordSpec) SetTag(v string) {
	o.Tag = &v
}

func (o EncryptionPasswordSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["hint"] = o.Hint
	}
	if !isNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableEncryptionPasswordSpec struct {
	value *EncryptionPasswordSpec
	isSet bool
}

func (v NullableEncryptionPasswordSpec) Get() *EncryptionPasswordSpec {
	return v.value
}

func (v *NullableEncryptionPasswordSpec) Set(val *EncryptionPasswordSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptionPasswordSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptionPasswordSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptionPasswordSpec(val *EncryptionPasswordSpec) *NullableEncryptionPasswordSpec {
	return &NullableEncryptionPasswordSpec{value: val, isSet: true}
}

func (v NullableEncryptionPasswordSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptionPasswordSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


