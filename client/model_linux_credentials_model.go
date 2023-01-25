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

// LinuxCredentialsModel struct for LinuxCredentialsModel
type LinuxCredentialsModel struct {
	CredentialsModel
	// Tag used to identify the credentials record.
	Tag *string `json:"tag,omitempty"`
	// SSH port used to connect to a Linux server.
	SSHPort *int32 `json:"SSHPort,omitempty"`
	// If *true*, the permissions of the account are automatically elevated to the root user.
	AutoElevated *bool `json:"autoElevated,omitempty"`
	// If *true*, the account is automatically added to the sudoers file.
	AddToSudoers *bool `json:"addToSudoers,omitempty"`
	// If *true*, the `su` command is used for Linux distributions where the `sudo` command is not available.
	UseSu *bool `json:"useSu,omitempty"`
	// Private key.
	PrivateKey *string `json:"privateKey,omitempty"`
	// Passphrase that protects the private key.
	Passphrase *string `json:"passphrase,omitempty"`
}

// NewLinuxCredentialsModel instantiates a new LinuxCredentialsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinuxCredentialsModel() *LinuxCredentialsModel {
	this := LinuxCredentialsModel{}
	return &this
}

// NewLinuxCredentialsModelWithDefaults instantiates a new LinuxCredentialsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinuxCredentialsModelWithDefaults() *LinuxCredentialsModel {
	this := LinuxCredentialsModel{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *LinuxCredentialsModel) GetTag() string {
	if o == nil || isNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxCredentialsModel) GetTagOk() (*string, bool) {
	if o == nil || isNil(o.Tag) {
    return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *LinuxCredentialsModel) HasTag() bool {
	if o != nil && !isNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *LinuxCredentialsModel) SetTag(v string) {
	o.Tag = &v
}

// GetSSHPort returns the SSHPort field value if set, zero value otherwise.
func (o *LinuxCredentialsModel) GetSSHPort() int32 {
	if o == nil || isNil(o.SSHPort) {
		var ret int32
		return ret
	}
	return *o.SSHPort
}

// GetSSHPortOk returns a tuple with the SSHPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxCredentialsModel) GetSSHPortOk() (*int32, bool) {
	if o == nil || isNil(o.SSHPort) {
    return nil, false
	}
	return o.SSHPort, true
}

// HasSSHPort returns a boolean if a field has been set.
func (o *LinuxCredentialsModel) HasSSHPort() bool {
	if o != nil && !isNil(o.SSHPort) {
		return true
	}

	return false
}

// SetSSHPort gets a reference to the given int32 and assigns it to the SSHPort field.
func (o *LinuxCredentialsModel) SetSSHPort(v int32) {
	o.SSHPort = &v
}

// GetAutoElevated returns the AutoElevated field value if set, zero value otherwise.
func (o *LinuxCredentialsModel) GetAutoElevated() bool {
	if o == nil || isNil(o.AutoElevated) {
		var ret bool
		return ret
	}
	return *o.AutoElevated
}

// GetAutoElevatedOk returns a tuple with the AutoElevated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxCredentialsModel) GetAutoElevatedOk() (*bool, bool) {
	if o == nil || isNil(o.AutoElevated) {
    return nil, false
	}
	return o.AutoElevated, true
}

// HasAutoElevated returns a boolean if a field has been set.
func (o *LinuxCredentialsModel) HasAutoElevated() bool {
	if o != nil && !isNil(o.AutoElevated) {
		return true
	}

	return false
}

// SetAutoElevated gets a reference to the given bool and assigns it to the AutoElevated field.
func (o *LinuxCredentialsModel) SetAutoElevated(v bool) {
	o.AutoElevated = &v
}

// GetAddToSudoers returns the AddToSudoers field value if set, zero value otherwise.
func (o *LinuxCredentialsModel) GetAddToSudoers() bool {
	if o == nil || isNil(o.AddToSudoers) {
		var ret bool
		return ret
	}
	return *o.AddToSudoers
}

// GetAddToSudoersOk returns a tuple with the AddToSudoers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxCredentialsModel) GetAddToSudoersOk() (*bool, bool) {
	if o == nil || isNil(o.AddToSudoers) {
    return nil, false
	}
	return o.AddToSudoers, true
}

// HasAddToSudoers returns a boolean if a field has been set.
func (o *LinuxCredentialsModel) HasAddToSudoers() bool {
	if o != nil && !isNil(o.AddToSudoers) {
		return true
	}

	return false
}

// SetAddToSudoers gets a reference to the given bool and assigns it to the AddToSudoers field.
func (o *LinuxCredentialsModel) SetAddToSudoers(v bool) {
	o.AddToSudoers = &v
}

// GetUseSu returns the UseSu field value if set, zero value otherwise.
func (o *LinuxCredentialsModel) GetUseSu() bool {
	if o == nil || isNil(o.UseSu) {
		var ret bool
		return ret
	}
	return *o.UseSu
}

// GetUseSuOk returns a tuple with the UseSu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxCredentialsModel) GetUseSuOk() (*bool, bool) {
	if o == nil || isNil(o.UseSu) {
    return nil, false
	}
	return o.UseSu, true
}

// HasUseSu returns a boolean if a field has been set.
func (o *LinuxCredentialsModel) HasUseSu() bool {
	if o != nil && !isNil(o.UseSu) {
		return true
	}

	return false
}

// SetUseSu gets a reference to the given bool and assigns it to the UseSu field.
func (o *LinuxCredentialsModel) SetUseSu(v bool) {
	o.UseSu = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *LinuxCredentialsModel) GetPrivateKey() string {
	if o == nil || isNil(o.PrivateKey) {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxCredentialsModel) GetPrivateKeyOk() (*string, bool) {
	if o == nil || isNil(o.PrivateKey) {
    return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *LinuxCredentialsModel) HasPrivateKey() bool {
	if o != nil && !isNil(o.PrivateKey) {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *LinuxCredentialsModel) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *LinuxCredentialsModel) GetPassphrase() string {
	if o == nil || isNil(o.Passphrase) {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxCredentialsModel) GetPassphraseOk() (*string, bool) {
	if o == nil || isNil(o.Passphrase) {
    return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *LinuxCredentialsModel) HasPassphrase() bool {
	if o != nil && !isNil(o.Passphrase) {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *LinuxCredentialsModel) SetPassphrase(v string) {
	o.Passphrase = &v
}

func (o LinuxCredentialsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCredentialsModel, errCredentialsModel := json.Marshal(o.CredentialsModel)
	if errCredentialsModel != nil {
		return []byte{}, errCredentialsModel
	}
	errCredentialsModel = json.Unmarshal([]byte(serializedCredentialsModel), &toSerialize)
	if errCredentialsModel != nil {
		return []byte{}, errCredentialsModel
	}
	if !isNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !isNil(o.SSHPort) {
		toSerialize["SSHPort"] = o.SSHPort
	}
	if !isNil(o.AutoElevated) {
		toSerialize["autoElevated"] = o.AutoElevated
	}
	if !isNil(o.AddToSudoers) {
		toSerialize["addToSudoers"] = o.AddToSudoers
	}
	if !isNil(o.UseSu) {
		toSerialize["useSu"] = o.UseSu
	}
	if !isNil(o.PrivateKey) {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if !isNil(o.Passphrase) {
		toSerialize["passphrase"] = o.Passphrase
	}
	return json.Marshal(toSerialize)
}

type NullableLinuxCredentialsModel struct {
	value *LinuxCredentialsModel
	isSet bool
}

func (v NullableLinuxCredentialsModel) Get() *LinuxCredentialsModel {
	return v.value
}

func (v *NullableLinuxCredentialsModel) Set(val *LinuxCredentialsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLinuxCredentialsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLinuxCredentialsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinuxCredentialsModel(val *LinuxCredentialsModel) *NullableLinuxCredentialsModel {
	return &NullableLinuxCredentialsModel{value: val, isSet: true}
}

func (v NullableLinuxCredentialsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinuxCredentialsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


