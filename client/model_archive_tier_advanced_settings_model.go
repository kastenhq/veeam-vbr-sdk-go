/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 
 *
 * API version: 1.0-rev1
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ArchiveTierAdvancedSettingsModel Advanced settings of the archive tier.
type ArchiveTierAdvancedSettingsModel struct {
	// If *true*, backups are archived if the remaining retention time is above minimum storage period for the repository.
	CostOptimizedArchiveEnabled *bool `json:"costOptimizedArchiveEnabled,omitempty"`
	// If *true*, each backup is stored as a delta to the previous one.
	ArchiveDeduplicationEnabled *bool `json:"archiveDeduplicationEnabled,omitempty"`
}

// NewArchiveTierAdvancedSettingsModel instantiates a new ArchiveTierAdvancedSettingsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchiveTierAdvancedSettingsModel() *ArchiveTierAdvancedSettingsModel {
	this := ArchiveTierAdvancedSettingsModel{}
	return &this
}

// NewArchiveTierAdvancedSettingsModelWithDefaults instantiates a new ArchiveTierAdvancedSettingsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchiveTierAdvancedSettingsModelWithDefaults() *ArchiveTierAdvancedSettingsModel {
	this := ArchiveTierAdvancedSettingsModel{}
	return &this
}

// GetCostOptimizedArchiveEnabled returns the CostOptimizedArchiveEnabled field value if set, zero value otherwise.
func (o *ArchiveTierAdvancedSettingsModel) GetCostOptimizedArchiveEnabled() bool {
	if o == nil || o.CostOptimizedArchiveEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CostOptimizedArchiveEnabled
}

// GetCostOptimizedArchiveEnabledOk returns a tuple with the CostOptimizedArchiveEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveTierAdvancedSettingsModel) GetCostOptimizedArchiveEnabledOk() (*bool, bool) {
	if o == nil || o.CostOptimizedArchiveEnabled == nil {
		return nil, false
	}
	return o.CostOptimizedArchiveEnabled, true
}

// HasCostOptimizedArchiveEnabled returns a boolean if a field has been set.
func (o *ArchiveTierAdvancedSettingsModel) HasCostOptimizedArchiveEnabled() bool {
	if o != nil && o.CostOptimizedArchiveEnabled != nil {
		return true
	}

	return false
}

// SetCostOptimizedArchiveEnabled gets a reference to the given bool and assigns it to the CostOptimizedArchiveEnabled field.
func (o *ArchiveTierAdvancedSettingsModel) SetCostOptimizedArchiveEnabled(v bool) {
	o.CostOptimizedArchiveEnabled = &v
}

// GetArchiveDeduplicationEnabled returns the ArchiveDeduplicationEnabled field value if set, zero value otherwise.
func (o *ArchiveTierAdvancedSettingsModel) GetArchiveDeduplicationEnabled() bool {
	if o == nil || o.ArchiveDeduplicationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ArchiveDeduplicationEnabled
}

// GetArchiveDeduplicationEnabledOk returns a tuple with the ArchiveDeduplicationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveTierAdvancedSettingsModel) GetArchiveDeduplicationEnabledOk() (*bool, bool) {
	if o == nil || o.ArchiveDeduplicationEnabled == nil {
		return nil, false
	}
	return o.ArchiveDeduplicationEnabled, true
}

// HasArchiveDeduplicationEnabled returns a boolean if a field has been set.
func (o *ArchiveTierAdvancedSettingsModel) HasArchiveDeduplicationEnabled() bool {
	if o != nil && o.ArchiveDeduplicationEnabled != nil {
		return true
	}

	return false
}

// SetArchiveDeduplicationEnabled gets a reference to the given bool and assigns it to the ArchiveDeduplicationEnabled field.
func (o *ArchiveTierAdvancedSettingsModel) SetArchiveDeduplicationEnabled(v bool) {
	o.ArchiveDeduplicationEnabled = &v
}

func (o ArchiveTierAdvancedSettingsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CostOptimizedArchiveEnabled != nil {
		toSerialize["costOptimizedArchiveEnabled"] = o.CostOptimizedArchiveEnabled
	}
	if o.ArchiveDeduplicationEnabled != nil {
		toSerialize["archiveDeduplicationEnabled"] = o.ArchiveDeduplicationEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableArchiveTierAdvancedSettingsModel struct {
	value *ArchiveTierAdvancedSettingsModel
	isSet bool
}

func (v NullableArchiveTierAdvancedSettingsModel) Get() *ArchiveTierAdvancedSettingsModel {
	return v.value
}

func (v *NullableArchiveTierAdvancedSettingsModel) Set(val *ArchiveTierAdvancedSettingsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableArchiveTierAdvancedSettingsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableArchiveTierAdvancedSettingsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchiveTierAdvancedSettingsModel(val *ArchiveTierAdvancedSettingsModel) *NullableArchiveTierAdvancedSettingsModel {
	return &NullableArchiveTierAdvancedSettingsModel{value: val, isSet: true}
}

func (v NullableArchiveTierAdvancedSettingsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchiveTierAdvancedSettingsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


