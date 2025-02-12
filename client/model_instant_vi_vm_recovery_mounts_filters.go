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

// InstantViVMRecoveryMountsFilters struct for InstantViVMRecoveryMountsFilters
type InstantViVMRecoveryMountsFilters struct {
	// Number of mounts to skip.
	Skip *int32 `json:"skip,omitempty"`
	// Maximum number of mounts to return.
	Limit *int32 `json:"limit,omitempty"`
	OrderColumn *EInstantViVMRecoveryMountsFiltersOrderColumn `json:"orderColumn,omitempty"`
	// Sorts mounts in the ascending order by the `orderColumn` parameter.
	OrderAsc *bool `json:"orderAsc,omitempty"`
	// Filters mounts by the `nameFilter` pattern. The pattern can match any mount parameter. To substitute one or more characters, use the asterisk (*) character at the beginning, at the end or both.
	NameFilter *string `json:"nameFilter,omitempty"`
	StateFilter *EInstantRecoveryMountState `json:"stateFilter,omitempty"`
}

// NewInstantViVMRecoveryMountsFilters instantiates a new InstantViVMRecoveryMountsFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstantViVMRecoveryMountsFilters() *InstantViVMRecoveryMountsFilters {
	this := InstantViVMRecoveryMountsFilters{}
	return &this
}

// NewInstantViVMRecoveryMountsFiltersWithDefaults instantiates a new InstantViVMRecoveryMountsFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstantViVMRecoveryMountsFiltersWithDefaults() *InstantViVMRecoveryMountsFilters {
	this := InstantViVMRecoveryMountsFilters{}
	return &this
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *InstantViVMRecoveryMountsFilters) GetSkip() int32 {
	if o == nil || isNil(o.Skip) {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstantViVMRecoveryMountsFilters) GetSkipOk() (*int32, bool) {
	if o == nil || isNil(o.Skip) {
    return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *InstantViVMRecoveryMountsFilters) HasSkip() bool {
	if o != nil && !isNil(o.Skip) {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *InstantViVMRecoveryMountsFilters) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *InstantViVMRecoveryMountsFilters) GetLimit() int32 {
	if o == nil || isNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstantViVMRecoveryMountsFilters) GetLimitOk() (*int32, bool) {
	if o == nil || isNil(o.Limit) {
    return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *InstantViVMRecoveryMountsFilters) HasLimit() bool {
	if o != nil && !isNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *InstantViVMRecoveryMountsFilters) SetLimit(v int32) {
	o.Limit = &v
}

// GetOrderColumn returns the OrderColumn field value if set, zero value otherwise.
func (o *InstantViVMRecoveryMountsFilters) GetOrderColumn() EInstantViVMRecoveryMountsFiltersOrderColumn {
	if o == nil || isNil(o.OrderColumn) {
		var ret EInstantViVMRecoveryMountsFiltersOrderColumn
		return ret
	}
	return *o.OrderColumn
}

// GetOrderColumnOk returns a tuple with the OrderColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstantViVMRecoveryMountsFilters) GetOrderColumnOk() (*EInstantViVMRecoveryMountsFiltersOrderColumn, bool) {
	if o == nil || isNil(o.OrderColumn) {
    return nil, false
	}
	return o.OrderColumn, true
}

// HasOrderColumn returns a boolean if a field has been set.
func (o *InstantViVMRecoveryMountsFilters) HasOrderColumn() bool {
	if o != nil && !isNil(o.OrderColumn) {
		return true
	}

	return false
}

// SetOrderColumn gets a reference to the given EInstantViVMRecoveryMountsFiltersOrderColumn and assigns it to the OrderColumn field.
func (o *InstantViVMRecoveryMountsFilters) SetOrderColumn(v EInstantViVMRecoveryMountsFiltersOrderColumn) {
	o.OrderColumn = &v
}

// GetOrderAsc returns the OrderAsc field value if set, zero value otherwise.
func (o *InstantViVMRecoveryMountsFilters) GetOrderAsc() bool {
	if o == nil || isNil(o.OrderAsc) {
		var ret bool
		return ret
	}
	return *o.OrderAsc
}

// GetOrderAscOk returns a tuple with the OrderAsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstantViVMRecoveryMountsFilters) GetOrderAscOk() (*bool, bool) {
	if o == nil || isNil(o.OrderAsc) {
    return nil, false
	}
	return o.OrderAsc, true
}

// HasOrderAsc returns a boolean if a field has been set.
func (o *InstantViVMRecoveryMountsFilters) HasOrderAsc() bool {
	if o != nil && !isNil(o.OrderAsc) {
		return true
	}

	return false
}

// SetOrderAsc gets a reference to the given bool and assigns it to the OrderAsc field.
func (o *InstantViVMRecoveryMountsFilters) SetOrderAsc(v bool) {
	o.OrderAsc = &v
}

// GetNameFilter returns the NameFilter field value if set, zero value otherwise.
func (o *InstantViVMRecoveryMountsFilters) GetNameFilter() string {
	if o == nil || isNil(o.NameFilter) {
		var ret string
		return ret
	}
	return *o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstantViVMRecoveryMountsFilters) GetNameFilterOk() (*string, bool) {
	if o == nil || isNil(o.NameFilter) {
    return nil, false
	}
	return o.NameFilter, true
}

// HasNameFilter returns a boolean if a field has been set.
func (o *InstantViVMRecoveryMountsFilters) HasNameFilter() bool {
	if o != nil && !isNil(o.NameFilter) {
		return true
	}

	return false
}

// SetNameFilter gets a reference to the given string and assigns it to the NameFilter field.
func (o *InstantViVMRecoveryMountsFilters) SetNameFilter(v string) {
	o.NameFilter = &v
}

// GetStateFilter returns the StateFilter field value if set, zero value otherwise.
func (o *InstantViVMRecoveryMountsFilters) GetStateFilter() EInstantRecoveryMountState {
	if o == nil || isNil(o.StateFilter) {
		var ret EInstantRecoveryMountState
		return ret
	}
	return *o.StateFilter
}

// GetStateFilterOk returns a tuple with the StateFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstantViVMRecoveryMountsFilters) GetStateFilterOk() (*EInstantRecoveryMountState, bool) {
	if o == nil || isNil(o.StateFilter) {
    return nil, false
	}
	return o.StateFilter, true
}

// HasStateFilter returns a boolean if a field has been set.
func (o *InstantViVMRecoveryMountsFilters) HasStateFilter() bool {
	if o != nil && !isNil(o.StateFilter) {
		return true
	}

	return false
}

// SetStateFilter gets a reference to the given EInstantRecoveryMountState and assigns it to the StateFilter field.
func (o *InstantViVMRecoveryMountsFilters) SetStateFilter(v EInstantRecoveryMountState) {
	o.StateFilter = &v
}

func (o InstantViVMRecoveryMountsFilters) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.StateFilter) {
		toSerialize["stateFilter"] = o.StateFilter
	}
	return json.Marshal(toSerialize)
}

type NullableInstantViVMRecoveryMountsFilters struct {
	value *InstantViVMRecoveryMountsFilters
	isSet bool
}

func (v NullableInstantViVMRecoveryMountsFilters) Get() *InstantViVMRecoveryMountsFilters {
	return v.value
}

func (v *NullableInstantViVMRecoveryMountsFilters) Set(val *InstantViVMRecoveryMountsFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableInstantViVMRecoveryMountsFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableInstantViVMRecoveryMountsFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstantViVMRecoveryMountsFilters(val *InstantViVMRecoveryMountsFilters) *NullableInstantViVMRecoveryMountsFilters {
	return &NullableInstantViVMRecoveryMountsFilters{value: val, isSet: true}
}

func (v NullableInstantViVMRecoveryMountsFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstantViVMRecoveryMountsFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


