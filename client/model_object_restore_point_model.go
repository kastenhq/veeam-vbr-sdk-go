/*
Veeam Backup & Replication REST API

This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br>Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Some request and response bodies refer to reusable schema objects that are defined in the `schemas` section of the REST API specification. Schemas, in turn, may inherit a part of their properties by referring to other schemas, and be polymorphic by referring to multiple alternate schemas.

API version: 1.1-rev0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// ObjectRestorePointModel struct for ObjectRestorePointModel
type ObjectRestorePointModel struct {
	// ID of the restore point.
	Id string `json:"id"`
	// Object name.
	Name string `json:"name"`
	PlatformName *EPlatformType `json:"platformName,omitempty"`
	// ID of a platform on which the object was created.
	PlatformId string `json:"platformId"`
	// Date and time when the restore point was created.
	CreationTime time.Time `json:"creationTime"`
	// ID of a backup that contains the restore point.
	BackupId string `json:"backupId"`
	// Array of operations allowed for the restore point.
	AllowedOperations []EObjectRestorePointOperation `json:"allowedOperations"`
}

// NewObjectRestorePointModel instantiates a new ObjectRestorePointModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectRestorePointModel(id string, name string, platformId string, creationTime time.Time, backupId string, allowedOperations []EObjectRestorePointOperation) *ObjectRestorePointModel {
	this := ObjectRestorePointModel{}
	this.Id = id
	this.Name = name
	this.PlatformId = platformId
	this.CreationTime = creationTime
	this.BackupId = backupId
	this.AllowedOperations = allowedOperations
	return &this
}

// NewObjectRestorePointModelWithDefaults instantiates a new ObjectRestorePointModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectRestorePointModelWithDefaults() *ObjectRestorePointModel {
	this := ObjectRestorePointModel{}
	return &this
}

// GetId returns the Id field value
func (o *ObjectRestorePointModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointModel) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ObjectRestorePointModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ObjectRestorePointModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointModel) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ObjectRestorePointModel) SetName(v string) {
	o.Name = v
}

// GetPlatformName returns the PlatformName field value if set, zero value otherwise.
func (o *ObjectRestorePointModel) GetPlatformName() EPlatformType {
	if o == nil || isNil(o.PlatformName) {
		var ret EPlatformType
		return ret
	}
	return *o.PlatformName
}

// GetPlatformNameOk returns a tuple with the PlatformName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointModel) GetPlatformNameOk() (*EPlatformType, bool) {
	if o == nil || isNil(o.PlatformName) {
    return nil, false
	}
	return o.PlatformName, true
}

// HasPlatformName returns a boolean if a field has been set.
func (o *ObjectRestorePointModel) HasPlatformName() bool {
	if o != nil && !isNil(o.PlatformName) {
		return true
	}

	return false
}

// SetPlatformName gets a reference to the given EPlatformType and assigns it to the PlatformName field.
func (o *ObjectRestorePointModel) SetPlatformName(v EPlatformType) {
	o.PlatformName = &v
}

// GetPlatformId returns the PlatformId field value
func (o *ObjectRestorePointModel) GetPlatformId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlatformId
}

// GetPlatformIdOk returns a tuple with the PlatformId field value
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointModel) GetPlatformIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PlatformId, true
}

// SetPlatformId sets field value
func (o *ObjectRestorePointModel) SetPlatformId(v string) {
	o.PlatformId = v
}

// GetCreationTime returns the CreationTime field value
func (o *ObjectRestorePointModel) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointModel) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *ObjectRestorePointModel) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetBackupId returns the BackupId field value
func (o *ObjectRestorePointModel) GetBackupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BackupId
}

// GetBackupIdOk returns a tuple with the BackupId field value
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointModel) GetBackupIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BackupId, true
}

// SetBackupId sets field value
func (o *ObjectRestorePointModel) SetBackupId(v string) {
	o.BackupId = v
}

// GetAllowedOperations returns the AllowedOperations field value
func (o *ObjectRestorePointModel) GetAllowedOperations() []EObjectRestorePointOperation {
	if o == nil {
		var ret []EObjectRestorePointOperation
		return ret
	}

	return o.AllowedOperations
}

// GetAllowedOperationsOk returns a tuple with the AllowedOperations field value
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointModel) GetAllowedOperationsOk() ([]EObjectRestorePointOperation, bool) {
	if o == nil {
    return nil, false
	}
	return o.AllowedOperations, true
}

// SetAllowedOperations sets field value
func (o *ObjectRestorePointModel) SetAllowedOperations(v []EObjectRestorePointOperation) {
	o.AllowedOperations = v
}

func (o ObjectRestorePointModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.PlatformName) {
		toSerialize["platformName"] = o.PlatformName
	}
	if true {
		toSerialize["platformId"] = o.PlatformId
	}
	if true {
		toSerialize["creationTime"] = o.CreationTime
	}
	if true {
		toSerialize["backupId"] = o.BackupId
	}
	if true {
		toSerialize["allowedOperations"] = o.AllowedOperations
	}
	return json.Marshal(toSerialize)
}

type NullableObjectRestorePointModel struct {
	value *ObjectRestorePointModel
	isSet bool
}

func (v NullableObjectRestorePointModel) Get() *ObjectRestorePointModel {
	return v.value
}

func (v *NullableObjectRestorePointModel) Set(val *ObjectRestorePointModel) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectRestorePointModel) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectRestorePointModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectRestorePointModel(val *ObjectRestorePointModel) *NullableObjectRestorePointModel {
	return &NullableObjectRestorePointModel{value: val, isSet: true}
}

func (v NullableObjectRestorePointModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectRestorePointModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


