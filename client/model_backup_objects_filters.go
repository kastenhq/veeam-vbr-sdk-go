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

// BackupObjectsFilters struct for BackupObjectsFilters
type BackupObjectsFilters struct {
	Skip *int32 `json:"skip,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	OrderColumn *EBackupObjectsFiltersOrderColumn `json:"orderColumn,omitempty"`
	OrderAsc *bool `json:"orderAsc,omitempty"`
	NameFilter *string `json:"nameFilter,omitempty"`
	PlatformNameFilter *EPlatformType `json:"platformNameFilter,omitempty"`
	PlatformIdFilter *string `json:"platformIdFilter,omitempty"`
	TypeFilter *string `json:"typeFilter,omitempty"`
	ViTypeFilter *string `json:"viTypeFilter,omitempty"`
}

// NewBackupObjectsFilters instantiates a new BackupObjectsFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupObjectsFilters() *BackupObjectsFilters {
	this := BackupObjectsFilters{}
	return &this
}

// NewBackupObjectsFiltersWithDefaults instantiates a new BackupObjectsFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupObjectsFiltersWithDefaults() *BackupObjectsFilters {
	this := BackupObjectsFilters{}
	return &this
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *BackupObjectsFilters) GetSkip() int32 {
	if o == nil || isNil(o.Skip) {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupObjectsFilters) GetSkipOk() (*int32, bool) {
	if o == nil || isNil(o.Skip) {
    return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *BackupObjectsFilters) HasSkip() bool {
	if o != nil && !isNil(o.Skip) {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *BackupObjectsFilters) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *BackupObjectsFilters) GetLimit() int32 {
	if o == nil || isNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupObjectsFilters) GetLimitOk() (*int32, bool) {
	if o == nil || isNil(o.Limit) {
    return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *BackupObjectsFilters) HasLimit() bool {
	if o != nil && !isNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *BackupObjectsFilters) SetLimit(v int32) {
	o.Limit = &v
}

// GetOrderColumn returns the OrderColumn field value if set, zero value otherwise.
func (o *BackupObjectsFilters) GetOrderColumn() EBackupObjectsFiltersOrderColumn {
	if o == nil || isNil(o.OrderColumn) {
		var ret EBackupObjectsFiltersOrderColumn
		return ret
	}
	return *o.OrderColumn
}

// GetOrderColumnOk returns a tuple with the OrderColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupObjectsFilters) GetOrderColumnOk() (*EBackupObjectsFiltersOrderColumn, bool) {
	if o == nil || isNil(o.OrderColumn) {
    return nil, false
	}
	return o.OrderColumn, true
}

// HasOrderColumn returns a boolean if a field has been set.
func (o *BackupObjectsFilters) HasOrderColumn() bool {
	if o != nil && !isNil(o.OrderColumn) {
		return true
	}

	return false
}

// SetOrderColumn gets a reference to the given EBackupObjectsFiltersOrderColumn and assigns it to the OrderColumn field.
func (o *BackupObjectsFilters) SetOrderColumn(v EBackupObjectsFiltersOrderColumn) {
	o.OrderColumn = &v
}

// GetOrderAsc returns the OrderAsc field value if set, zero value otherwise.
func (o *BackupObjectsFilters) GetOrderAsc() bool {
	if o == nil || isNil(o.OrderAsc) {
		var ret bool
		return ret
	}
	return *o.OrderAsc
}

// GetOrderAscOk returns a tuple with the OrderAsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupObjectsFilters) GetOrderAscOk() (*bool, bool) {
	if o == nil || isNil(o.OrderAsc) {
    return nil, false
	}
	return o.OrderAsc, true
}

// HasOrderAsc returns a boolean if a field has been set.
func (o *BackupObjectsFilters) HasOrderAsc() bool {
	if o != nil && !isNil(o.OrderAsc) {
		return true
	}

	return false
}

// SetOrderAsc gets a reference to the given bool and assigns it to the OrderAsc field.
func (o *BackupObjectsFilters) SetOrderAsc(v bool) {
	o.OrderAsc = &v
}

// GetNameFilter returns the NameFilter field value if set, zero value otherwise.
func (o *BackupObjectsFilters) GetNameFilter() string {
	if o == nil || isNil(o.NameFilter) {
		var ret string
		return ret
	}
	return *o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupObjectsFilters) GetNameFilterOk() (*string, bool) {
	if o == nil || isNil(o.NameFilter) {
    return nil, false
	}
	return o.NameFilter, true
}

// HasNameFilter returns a boolean if a field has been set.
func (o *BackupObjectsFilters) HasNameFilter() bool {
	if o != nil && !isNil(o.NameFilter) {
		return true
	}

	return false
}

// SetNameFilter gets a reference to the given string and assigns it to the NameFilter field.
func (o *BackupObjectsFilters) SetNameFilter(v string) {
	o.NameFilter = &v
}

// GetPlatformNameFilter returns the PlatformNameFilter field value if set, zero value otherwise.
func (o *BackupObjectsFilters) GetPlatformNameFilter() EPlatformType {
	if o == nil || isNil(o.PlatformNameFilter) {
		var ret EPlatformType
		return ret
	}
	return *o.PlatformNameFilter
}

// GetPlatformNameFilterOk returns a tuple with the PlatformNameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupObjectsFilters) GetPlatformNameFilterOk() (*EPlatformType, bool) {
	if o == nil || isNil(o.PlatformNameFilter) {
    return nil, false
	}
	return o.PlatformNameFilter, true
}

// HasPlatformNameFilter returns a boolean if a field has been set.
func (o *BackupObjectsFilters) HasPlatformNameFilter() bool {
	if o != nil && !isNil(o.PlatformNameFilter) {
		return true
	}

	return false
}

// SetPlatformNameFilter gets a reference to the given EPlatformType and assigns it to the PlatformNameFilter field.
func (o *BackupObjectsFilters) SetPlatformNameFilter(v EPlatformType) {
	o.PlatformNameFilter = &v
}

// GetPlatformIdFilter returns the PlatformIdFilter field value if set, zero value otherwise.
func (o *BackupObjectsFilters) GetPlatformIdFilter() string {
	if o == nil || isNil(o.PlatformIdFilter) {
		var ret string
		return ret
	}
	return *o.PlatformIdFilter
}

// GetPlatformIdFilterOk returns a tuple with the PlatformIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupObjectsFilters) GetPlatformIdFilterOk() (*string, bool) {
	if o == nil || isNil(o.PlatformIdFilter) {
    return nil, false
	}
	return o.PlatformIdFilter, true
}

// HasPlatformIdFilter returns a boolean if a field has been set.
func (o *BackupObjectsFilters) HasPlatformIdFilter() bool {
	if o != nil && !isNil(o.PlatformIdFilter) {
		return true
	}

	return false
}

// SetPlatformIdFilter gets a reference to the given string and assigns it to the PlatformIdFilter field.
func (o *BackupObjectsFilters) SetPlatformIdFilter(v string) {
	o.PlatformIdFilter = &v
}

// GetTypeFilter returns the TypeFilter field value if set, zero value otherwise.
func (o *BackupObjectsFilters) GetTypeFilter() string {
	if o == nil || isNil(o.TypeFilter) {
		var ret string
		return ret
	}
	return *o.TypeFilter
}

// GetTypeFilterOk returns a tuple with the TypeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupObjectsFilters) GetTypeFilterOk() (*string, bool) {
	if o == nil || isNil(o.TypeFilter) {
    return nil, false
	}
	return o.TypeFilter, true
}

// HasTypeFilter returns a boolean if a field has been set.
func (o *BackupObjectsFilters) HasTypeFilter() bool {
	if o != nil && !isNil(o.TypeFilter) {
		return true
	}

	return false
}

// SetTypeFilter gets a reference to the given string and assigns it to the TypeFilter field.
func (o *BackupObjectsFilters) SetTypeFilter(v string) {
	o.TypeFilter = &v
}

// GetViTypeFilter returns the ViTypeFilter field value if set, zero value otherwise.
func (o *BackupObjectsFilters) GetViTypeFilter() string {
	if o == nil || isNil(o.ViTypeFilter) {
		var ret string
		return ret
	}
	return *o.ViTypeFilter
}

// GetViTypeFilterOk returns a tuple with the ViTypeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupObjectsFilters) GetViTypeFilterOk() (*string, bool) {
	if o == nil || isNil(o.ViTypeFilter) {
    return nil, false
	}
	return o.ViTypeFilter, true
}

// HasViTypeFilter returns a boolean if a field has been set.
func (o *BackupObjectsFilters) HasViTypeFilter() bool {
	if o != nil && !isNil(o.ViTypeFilter) {
		return true
	}

	return false
}

// SetViTypeFilter gets a reference to the given string and assigns it to the ViTypeFilter field.
func (o *BackupObjectsFilters) SetViTypeFilter(v string) {
	o.ViTypeFilter = &v
}

func (o BackupObjectsFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Skip) {
		toSerialize["skip"] = o.Skip
	}
	if !isNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !isNil(o.OrderColumn) {
		toSerialize["orderColumn"] = o.OrderColumn
	}
	if !isNil(o.OrderAsc) {
		toSerialize["orderAsc"] = o.OrderAsc
	}
	if !isNil(o.NameFilter) {
		toSerialize["nameFilter"] = o.NameFilter
	}
	if !isNil(o.PlatformNameFilter) {
		toSerialize["platformNameFilter"] = o.PlatformNameFilter
	}
	if !isNil(o.PlatformIdFilter) {
		toSerialize["platformIdFilter"] = o.PlatformIdFilter
	}
	if !isNil(o.TypeFilter) {
		toSerialize["typeFilter"] = o.TypeFilter
	}
	if !isNil(o.ViTypeFilter) {
		toSerialize["viTypeFilter"] = o.ViTypeFilter
	}
	return json.Marshal(toSerialize)
}

type NullableBackupObjectsFilters struct {
	value *BackupObjectsFilters
	isSet bool
}

func (v NullableBackupObjectsFilters) Get() *BackupObjectsFilters {
	return v.value
}

func (v *NullableBackupObjectsFilters) Set(val *BackupObjectsFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupObjectsFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupObjectsFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupObjectsFilters(val *BackupObjectsFilters) *NullableBackupObjectsFilters {
	return &NullableBackupObjectsFilters{value: val, isSet: true}
}

func (v NullableBackupObjectsFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupObjectsFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


