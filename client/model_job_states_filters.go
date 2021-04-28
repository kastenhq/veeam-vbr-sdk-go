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
	"time"
)

// JobStatesFilters Filters jobs by the specified parameters.
type JobStatesFilters struct {
	// Skips the specified number of jobs.
	Skip *int32 `json:"skip,omitempty"`
	// Returns the specified number of jobs.
	Limit *int32 `json:"limit,omitempty"`
	OrderColumn *EJobStatesFiltersOrderColumn `json:"orderColumn,omitempty"`
	// Sorts jobs in the ascending order by the `orderColumn` parameter.
	OrderAsc *bool `json:"orderAsc,omitempty"`
	IdFilter *string `json:"idFilter,omitempty"`
	// Filters jobs by the `nameFilter` pattern. The pattern can match any job state parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end.
	NameFilter *string `json:"nameFilter,omitempty"`
	TypeFilter *EJobType `json:"typeFilter,omitempty"`
	LastResultFilter *ESessionResult `json:"lastResultFilter,omitempty"`
	StatusFilter *EJobStatus `json:"statusFilter,omitempty"`
	WorkloadFilter *EJobWorkload `json:"workloadFilter,omitempty"`
	LastRunAfterFilter *time.Time `json:"lastRunAfterFilter,omitempty"`
	LastRunBeforeFilter *time.Time `json:"lastRunBeforeFilter,omitempty"`
	IsHighPriorityJobFilter *bool `json:"isHighPriorityJobFilter,omitempty"`
	RepositoryIdFilter *string `json:"repositoryIdFilter,omitempty"`
	ObjectsCountFilter *int32 `json:"objectsCountFilter,omitempty"`
}

// NewJobStatesFilters instantiates a new JobStatesFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobStatesFilters() *JobStatesFilters {
	this := JobStatesFilters{}
	return &this
}

// NewJobStatesFiltersWithDefaults instantiates a new JobStatesFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobStatesFiltersWithDefaults() *JobStatesFilters {
	this := JobStatesFilters{}
	return &this
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *JobStatesFilters) GetSkip() int32 {
	if o == nil || o.Skip == nil {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStatesFilters) GetSkipOk() (*int32, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *JobStatesFilters) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *JobStatesFilters) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *JobStatesFilters) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStatesFilters) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *JobStatesFilters) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *JobStatesFilters) SetLimit(v int32) {
	o.Limit = &v
}

// GetOrderColumn returns the OrderColumn field value if set, zero value otherwise.
func (o *JobStatesFilters) GetOrderColumn() EJobStatesFiltersOrderColumn {
	if o == nil || o.OrderColumn == nil {
		var ret EJobStatesFiltersOrderColumn
		return ret
	}
	return *o.OrderColumn
}

// GetOrderColumnOk returns a tuple with the OrderColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStatesFilters) GetOrderColumnOk() (*EJobStatesFiltersOrderColumn, bool) {
	if o == nil || o.OrderColumn == nil {
		return nil, false
	}
	return o.OrderColumn, true
}

// HasOrderColumn returns a boolean if a field has been set.
func (o *JobStatesFilters) HasOrderColumn() bool {
	if o != nil && o.OrderColumn != nil {
		return true
	}

	return false
}

// SetOrderColumn gets a reference to the given EJobStatesFiltersOrderColumn and assigns it to the OrderColumn field.
func (o *JobStatesFilters) SetOrderColumn(v EJobStatesFiltersOrderColumn) {
	o.OrderColumn = &v
}

// GetOrderAsc returns the OrderAsc field value if set, zero value otherwise.
func (o *JobStatesFilters) GetOrderAsc() bool {
	if o == nil || o.OrderAsc == nil {
		var ret bool
		return ret
	}
	return *o.OrderAsc
}

// GetOrderAscOk returns a tuple with the OrderAsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStatesFilters) GetOrderAscOk() (*bool, bool) {
	if o == nil || o.OrderAsc == nil {
		return nil, false
	}
	return o.OrderAsc, true
}

// HasOrderAsc returns a boolean if a field has been set.
func (o *JobStatesFilters) HasOrderAsc() bool {
	if o != nil && o.OrderAsc != nil {
		return true
	}

	return false
}

// SetOrderAsc gets a reference to the given bool and assigns it to the OrderAsc field.
func (o *JobStatesFilters) SetOrderAsc(v bool) {
	o.OrderAsc = &v
}

// GetIdFilter returns the IdFilter field value if set, zero value otherwise.
func (o *JobStatesFilters) GetIdFilter() string {
	if o == nil || o.IdFilter == nil {
		var ret string
		return ret
	}
	return *o.IdFilter
}

// GetIdFilterOk returns a tuple with the IdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStatesFilters) GetIdFilterOk() (*string, bool) {
	if o == nil || o.IdFilter == nil {
		return nil, false
	}
	return o.IdFilter, true
}

// HasIdFilter returns a boolean if a field has been set.
func (o *JobStatesFilters) HasIdFilter() bool {
	if o != nil && o.IdFilter != nil {
		return true
	}

	return false
}

// SetIdFilter gets a reference to the given string and assigns it to the IdFilter field.
func (o *JobStatesFilters) SetIdFilter(v string) {
	o.IdFilter = &v
}

// GetNameFilter returns the NameFilter field value if set, zero value otherwise.
func (o *JobStatesFilters) GetNameFilter() string {
	if o == nil || o.NameFilter == nil {
		var ret string
		return ret
	}
	return *o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStatesFilters) GetNameFilterOk() (*string, bool) {
	if o == nil || o.NameFilter == nil {
		return nil, false
	}
	return o.NameFilter, true
}

// HasNameFilter returns a boolean if a field has been set.
func (o *JobStatesFilters) HasNameFilter() bool {
	if o != nil && o.NameFilter != nil {
		return true
	}

	return false
}

// SetNameFilter gets a reference to the given string and assigns it to the NameFilter field.
func (o *JobStatesFilters) SetNameFilter(v string) {
	o.NameFilter = &v
}

// GetTypeFilter returns the TypeFilter field value if set, zero value otherwise.
func (o *JobStatesFilters) GetTypeFilter() EJobType {
	if o == nil || o.TypeFilter == nil {
		var ret EJobType
		return ret
	}
	return *o.TypeFilter
}

// GetTypeFilterOk returns a tuple with the TypeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStatesFilters) GetTypeFilterOk() (*EJobType, bool) {
	if o == nil || o.TypeFilter == nil {
		return nil, false
	}
	return o.TypeFilter, true
}

// HasTypeFilter returns a boolean if a field has been set.
func (o *JobStatesFilters) HasTypeFilter() bool {
	if o != nil && o.TypeFilter != nil {
		return true
	}

	return false
}

// SetTypeFilter gets a reference to the given EJobType and assigns it to the TypeFilter field.
func (o *JobStatesFilters) SetTypeFilter(v EJobType) {
	o.TypeFilter = &v
}

// GetLastResultFilter returns the LastResultFilter field value if set, zero value otherwise.
func (o *JobStatesFilters) GetLastResultFilter() ESessionResult {
	if o == nil || o.LastResultFilter == nil {
		var ret ESessionResult
		return ret
	}
	return *o.LastResultFilter
}

// GetLastResultFilterOk returns a tuple with the LastResultFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStatesFilters) GetLastResultFilterOk() (*ESessionResult, bool) {
	if o == nil || o.LastResultFilter == nil {
		return nil, false
	}
	return o.LastResultFilter, true
}

// HasLastResultFilter returns a boolean if a field has been set.
func (o *JobStatesFilters) HasLastResultFilter() bool {
	if o != nil && o.LastResultFilter != nil {
		return true
	}

	return false
}

// SetLastResultFilter gets a reference to the given ESessionResult and assigns it to the LastResultFilter field.
func (o *JobStatesFilters) SetLastResultFilter(v ESessionResult) {
	o.LastResultFilter = &v
}

// GetStatusFilter returns the StatusFilter field value if set, zero value otherwise.
func (o *JobStatesFilters) GetStatusFilter() EJobStatus {
	if o == nil || o.StatusFilter == nil {
		var ret EJobStatus
		return ret
	}
	return *o.StatusFilter
}

// GetStatusFilterOk returns a tuple with the StatusFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStatesFilters) GetStatusFilterOk() (*EJobStatus, bool) {
	if o == nil || o.StatusFilter == nil {
		return nil, false
	}
	return o.StatusFilter, true
}

// HasStatusFilter returns a boolean if a field has been set.
func (o *JobStatesFilters) HasStatusFilter() bool {
	if o != nil && o.StatusFilter != nil {
		return true
	}

	return false
}

// SetStatusFilter gets a reference to the given EJobStatus and assigns it to the StatusFilter field.
func (o *JobStatesFilters) SetStatusFilter(v EJobStatus) {
	o.StatusFilter = &v
}

// GetWorkloadFilter returns the WorkloadFilter field value if set, zero value otherwise.
func (o *JobStatesFilters) GetWorkloadFilter() EJobWorkload {
	if o == nil || o.WorkloadFilter == nil {
		var ret EJobWorkload
		return ret
	}
	return *o.WorkloadFilter
}

// GetWorkloadFilterOk returns a tuple with the WorkloadFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStatesFilters) GetWorkloadFilterOk() (*EJobWorkload, bool) {
	if o == nil || o.WorkloadFilter == nil {
		return nil, false
	}
	return o.WorkloadFilter, true
}

// HasWorkloadFilter returns a boolean if a field has been set.
func (o *JobStatesFilters) HasWorkloadFilter() bool {
	if o != nil && o.WorkloadFilter != nil {
		return true
	}

	return false
}

// SetWorkloadFilter gets a reference to the given EJobWorkload and assigns it to the WorkloadFilter field.
func (o *JobStatesFilters) SetWorkloadFilter(v EJobWorkload) {
	o.WorkloadFilter = &v
}

// GetLastRunAfterFilter returns the LastRunAfterFilter field value if set, zero value otherwise.
func (o *JobStatesFilters) GetLastRunAfterFilter() time.Time {
	if o == nil || o.LastRunAfterFilter == nil {
		var ret time.Time
		return ret
	}
	return *o.LastRunAfterFilter
}

// GetLastRunAfterFilterOk returns a tuple with the LastRunAfterFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStatesFilters) GetLastRunAfterFilterOk() (*time.Time, bool) {
	if o == nil || o.LastRunAfterFilter == nil {
		return nil, false
	}
	return o.LastRunAfterFilter, true
}

// HasLastRunAfterFilter returns a boolean if a field has been set.
func (o *JobStatesFilters) HasLastRunAfterFilter() bool {
	if o != nil && o.LastRunAfterFilter != nil {
		return true
	}

	return false
}

// SetLastRunAfterFilter gets a reference to the given time.Time and assigns it to the LastRunAfterFilter field.
func (o *JobStatesFilters) SetLastRunAfterFilter(v time.Time) {
	o.LastRunAfterFilter = &v
}

// GetLastRunBeforeFilter returns the LastRunBeforeFilter field value if set, zero value otherwise.
func (o *JobStatesFilters) GetLastRunBeforeFilter() time.Time {
	if o == nil || o.LastRunBeforeFilter == nil {
		var ret time.Time
		return ret
	}
	return *o.LastRunBeforeFilter
}

// GetLastRunBeforeFilterOk returns a tuple with the LastRunBeforeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStatesFilters) GetLastRunBeforeFilterOk() (*time.Time, bool) {
	if o == nil || o.LastRunBeforeFilter == nil {
		return nil, false
	}
	return o.LastRunBeforeFilter, true
}

// HasLastRunBeforeFilter returns a boolean if a field has been set.
func (o *JobStatesFilters) HasLastRunBeforeFilter() bool {
	if o != nil && o.LastRunBeforeFilter != nil {
		return true
	}

	return false
}

// SetLastRunBeforeFilter gets a reference to the given time.Time and assigns it to the LastRunBeforeFilter field.
func (o *JobStatesFilters) SetLastRunBeforeFilter(v time.Time) {
	o.LastRunBeforeFilter = &v
}

// GetIsHighPriorityJobFilter returns the IsHighPriorityJobFilter field value if set, zero value otherwise.
func (o *JobStatesFilters) GetIsHighPriorityJobFilter() bool {
	if o == nil || o.IsHighPriorityJobFilter == nil {
		var ret bool
		return ret
	}
	return *o.IsHighPriorityJobFilter
}

// GetIsHighPriorityJobFilterOk returns a tuple with the IsHighPriorityJobFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStatesFilters) GetIsHighPriorityJobFilterOk() (*bool, bool) {
	if o == nil || o.IsHighPriorityJobFilter == nil {
		return nil, false
	}
	return o.IsHighPriorityJobFilter, true
}

// HasIsHighPriorityJobFilter returns a boolean if a field has been set.
func (o *JobStatesFilters) HasIsHighPriorityJobFilter() bool {
	if o != nil && o.IsHighPriorityJobFilter != nil {
		return true
	}

	return false
}

// SetIsHighPriorityJobFilter gets a reference to the given bool and assigns it to the IsHighPriorityJobFilter field.
func (o *JobStatesFilters) SetIsHighPriorityJobFilter(v bool) {
	o.IsHighPriorityJobFilter = &v
}

// GetRepositoryIdFilter returns the RepositoryIdFilter field value if set, zero value otherwise.
func (o *JobStatesFilters) GetRepositoryIdFilter() string {
	if o == nil || o.RepositoryIdFilter == nil {
		var ret string
		return ret
	}
	return *o.RepositoryIdFilter
}

// GetRepositoryIdFilterOk returns a tuple with the RepositoryIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStatesFilters) GetRepositoryIdFilterOk() (*string, bool) {
	if o == nil || o.RepositoryIdFilter == nil {
		return nil, false
	}
	return o.RepositoryIdFilter, true
}

// HasRepositoryIdFilter returns a boolean if a field has been set.
func (o *JobStatesFilters) HasRepositoryIdFilter() bool {
	if o != nil && o.RepositoryIdFilter != nil {
		return true
	}

	return false
}

// SetRepositoryIdFilter gets a reference to the given string and assigns it to the RepositoryIdFilter field.
func (o *JobStatesFilters) SetRepositoryIdFilter(v string) {
	o.RepositoryIdFilter = &v
}

// GetObjectsCountFilter returns the ObjectsCountFilter field value if set, zero value otherwise.
func (o *JobStatesFilters) GetObjectsCountFilter() int32 {
	if o == nil || o.ObjectsCountFilter == nil {
		var ret int32
		return ret
	}
	return *o.ObjectsCountFilter
}

// GetObjectsCountFilterOk returns a tuple with the ObjectsCountFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStatesFilters) GetObjectsCountFilterOk() (*int32, bool) {
	if o == nil || o.ObjectsCountFilter == nil {
		return nil, false
	}
	return o.ObjectsCountFilter, true
}

// HasObjectsCountFilter returns a boolean if a field has been set.
func (o *JobStatesFilters) HasObjectsCountFilter() bool {
	if o != nil && o.ObjectsCountFilter != nil {
		return true
	}

	return false
}

// SetObjectsCountFilter gets a reference to the given int32 and assigns it to the ObjectsCountFilter field.
func (o *JobStatesFilters) SetObjectsCountFilter(v int32) {
	o.ObjectsCountFilter = &v
}

func (o JobStatesFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Skip != nil {
		toSerialize["skip"] = o.Skip
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.OrderColumn != nil {
		toSerialize["orderColumn"] = o.OrderColumn
	}
	if o.OrderAsc != nil {
		toSerialize["orderAsc"] = o.OrderAsc
	}
	if o.IdFilter != nil {
		toSerialize["idFilter"] = o.IdFilter
	}
	if o.NameFilter != nil {
		toSerialize["nameFilter"] = o.NameFilter
	}
	if o.TypeFilter != nil {
		toSerialize["typeFilter"] = o.TypeFilter
	}
	if o.LastResultFilter != nil {
		toSerialize["lastResultFilter"] = o.LastResultFilter
	}
	if o.StatusFilter != nil {
		toSerialize["statusFilter"] = o.StatusFilter
	}
	if o.WorkloadFilter != nil {
		toSerialize["workloadFilter"] = o.WorkloadFilter
	}
	if o.LastRunAfterFilter != nil {
		toSerialize["lastRunAfterFilter"] = o.LastRunAfterFilter
	}
	if o.LastRunBeforeFilter != nil {
		toSerialize["lastRunBeforeFilter"] = o.LastRunBeforeFilter
	}
	if o.IsHighPriorityJobFilter != nil {
		toSerialize["isHighPriorityJobFilter"] = o.IsHighPriorityJobFilter
	}
	if o.RepositoryIdFilter != nil {
		toSerialize["repositoryIdFilter"] = o.RepositoryIdFilter
	}
	if o.ObjectsCountFilter != nil {
		toSerialize["objectsCountFilter"] = o.ObjectsCountFilter
	}
	return json.Marshal(toSerialize)
}

type NullableJobStatesFilters struct {
	value *JobStatesFilters
	isSet bool
}

func (v NullableJobStatesFilters) Get() *JobStatesFilters {
	return v.value
}

func (v *NullableJobStatesFilters) Set(val *JobStatesFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableJobStatesFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableJobStatesFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobStatesFilters(val *JobStatesFilters) *NullableJobStatesFilters {
	return &NullableJobStatesFilters{value: val, isSet: true}
}

func (v NullableJobStatesFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobStatesFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


