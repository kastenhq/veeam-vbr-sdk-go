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

// AmazonS3CloudBrowserFilters Amazon S3 hierarchy filters.
type AmazonS3CloudBrowserFilters struct {
	// Filters buckets by AWS region where an Amazon S3 data center is located.
	RegionId string `json:"regionId"`
	// Filters buckets by bucket name.
	BucketName *string `json:"bucketName,omitempty"`
}

// NewAmazonS3CloudBrowserFilters instantiates a new AmazonS3CloudBrowserFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonS3CloudBrowserFilters(regionId string) *AmazonS3CloudBrowserFilters {
	this := AmazonS3CloudBrowserFilters{}
	this.RegionId = regionId
	return &this
}

// NewAmazonS3CloudBrowserFiltersWithDefaults instantiates a new AmazonS3CloudBrowserFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonS3CloudBrowserFiltersWithDefaults() *AmazonS3CloudBrowserFilters {
	this := AmazonS3CloudBrowserFilters{}
	return &this
}

// GetRegionId returns the RegionId field value
func (o *AmazonS3CloudBrowserFilters) GetRegionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value
// and a boolean to check if the value has been set.
func (o *AmazonS3CloudBrowserFilters) GetRegionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RegionId, true
}

// SetRegionId sets field value
func (o *AmazonS3CloudBrowserFilters) SetRegionId(v string) {
	o.RegionId = v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *AmazonS3CloudBrowserFilters) GetBucketName() string {
	if o == nil || isNil(o.BucketName) {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonS3CloudBrowserFilters) GetBucketNameOk() (*string, bool) {
	if o == nil || isNil(o.BucketName) {
    return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *AmazonS3CloudBrowserFilters) HasBucketName() bool {
	if o != nil && !isNil(o.BucketName) {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *AmazonS3CloudBrowserFilters) SetBucketName(v string) {
	o.BucketName = &v
}

func (o AmazonS3CloudBrowserFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["regionId"] = o.RegionId
	}
	if !isNil(o.BucketName) {
		toSerialize["bucketName"] = o.BucketName
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonS3CloudBrowserFilters struct {
	value *AmazonS3CloudBrowserFilters
	isSet bool
}

func (v NullableAmazonS3CloudBrowserFilters) Get() *AmazonS3CloudBrowserFilters {
	return v.value
}

func (v *NullableAmazonS3CloudBrowserFilters) Set(val *AmazonS3CloudBrowserFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonS3CloudBrowserFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonS3CloudBrowserFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonS3CloudBrowserFilters(val *AmazonS3CloudBrowserFilters) *NullableAmazonS3CloudBrowserFilters {
	return &NullableAmazonS3CloudBrowserFilters{value: val, isSet: true}
}

func (v NullableAmazonS3CloudBrowserFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonS3CloudBrowserFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


