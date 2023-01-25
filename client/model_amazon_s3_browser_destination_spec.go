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

// AmazonS3BrowserDestinationSpec struct for AmazonS3BrowserDestinationSpec
type AmazonS3BrowserDestinationSpec struct {
	CloudBrowserNewFolderSpec
	// ID of a server you want to use to connect to the object storage. You can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used.
	HostId *string `json:"hostId,omitempty"`
	RegionType EAmazonRegionType `json:"regionType"`
	// AWS region where the Amazon S3 bucket is located.
	RegionId string `json:"regionId"`
	// Name of the bucket where you want to store your backup data.
	BucketName string `json:"bucketName"`
	FolderType *ECloudBrowserFolderType `json:"folderType,omitempty"`
}

// NewAmazonS3BrowserDestinationSpec instantiates a new AmazonS3BrowserDestinationSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonS3BrowserDestinationSpec(regionType EAmazonRegionType, regionId string, bucketName string) *AmazonS3BrowserDestinationSpec {
	this := AmazonS3BrowserDestinationSpec{}
	this.RegionType = regionType
	this.RegionId = regionId
	this.BucketName = bucketName
	return &this
}

// NewAmazonS3BrowserDestinationSpecWithDefaults instantiates a new AmazonS3BrowserDestinationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonS3BrowserDestinationSpecWithDefaults() *AmazonS3BrowserDestinationSpec {
	this := AmazonS3BrowserDestinationSpec{}
	return &this
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *AmazonS3BrowserDestinationSpec) GetHostId() string {
	if o == nil || isNil(o.HostId) {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonS3BrowserDestinationSpec) GetHostIdOk() (*string, bool) {
	if o == nil || isNil(o.HostId) {
    return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *AmazonS3BrowserDestinationSpec) HasHostId() bool {
	if o != nil && !isNil(o.HostId) {
		return true
	}

	return false
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *AmazonS3BrowserDestinationSpec) SetHostId(v string) {
	o.HostId = &v
}

// GetRegionType returns the RegionType field value
func (o *AmazonS3BrowserDestinationSpec) GetRegionType() EAmazonRegionType {
	if o == nil {
		var ret EAmazonRegionType
		return ret
	}

	return o.RegionType
}

// GetRegionTypeOk returns a tuple with the RegionType field value
// and a boolean to check if the value has been set.
func (o *AmazonS3BrowserDestinationSpec) GetRegionTypeOk() (*EAmazonRegionType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RegionType, true
}

// SetRegionType sets field value
func (o *AmazonS3BrowserDestinationSpec) SetRegionType(v EAmazonRegionType) {
	o.RegionType = v
}

// GetRegionId returns the RegionId field value
func (o *AmazonS3BrowserDestinationSpec) GetRegionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value
// and a boolean to check if the value has been set.
func (o *AmazonS3BrowserDestinationSpec) GetRegionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RegionId, true
}

// SetRegionId sets field value
func (o *AmazonS3BrowserDestinationSpec) SetRegionId(v string) {
	o.RegionId = v
}

// GetBucketName returns the BucketName field value
func (o *AmazonS3BrowserDestinationSpec) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *AmazonS3BrowserDestinationSpec) GetBucketNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value
func (o *AmazonS3BrowserDestinationSpec) SetBucketName(v string) {
	o.BucketName = v
}

// GetFolderType returns the FolderType field value if set, zero value otherwise.
func (o *AmazonS3BrowserDestinationSpec) GetFolderType() ECloudBrowserFolderType {
	if o == nil || isNil(o.FolderType) {
		var ret ECloudBrowserFolderType
		return ret
	}
	return *o.FolderType
}

// GetFolderTypeOk returns a tuple with the FolderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonS3BrowserDestinationSpec) GetFolderTypeOk() (*ECloudBrowserFolderType, bool) {
	if o == nil || isNil(o.FolderType) {
    return nil, false
	}
	return o.FolderType, true
}

// HasFolderType returns a boolean if a field has been set.
func (o *AmazonS3BrowserDestinationSpec) HasFolderType() bool {
	if o != nil && !isNil(o.FolderType) {
		return true
	}

	return false
}

// SetFolderType gets a reference to the given ECloudBrowserFolderType and assigns it to the FolderType field.
func (o *AmazonS3BrowserDestinationSpec) SetFolderType(v ECloudBrowserFolderType) {
	o.FolderType = &v
}

func (o AmazonS3BrowserDestinationSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBrowserNewFolderSpec, errCloudBrowserNewFolderSpec := json.Marshal(o.CloudBrowserNewFolderSpec)
	if errCloudBrowserNewFolderSpec != nil {
		return []byte{}, errCloudBrowserNewFolderSpec
	}
	errCloudBrowserNewFolderSpec = json.Unmarshal([]byte(serializedCloudBrowserNewFolderSpec), &toSerialize)
	if errCloudBrowserNewFolderSpec != nil {
		return []byte{}, errCloudBrowserNewFolderSpec
	}
	if !isNil(o.HostId) {
		toSerialize["hostId"] = o.HostId
	}
	if true {
		toSerialize["regionType"] = o.RegionType
	}
	if true {
		toSerialize["regionId"] = o.RegionId
	}
	if true {
		toSerialize["bucketName"] = o.BucketName
	}
	if !isNil(o.FolderType) {
		toSerialize["folderType"] = o.FolderType
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonS3BrowserDestinationSpec struct {
	value *AmazonS3BrowserDestinationSpec
	isSet bool
}

func (v NullableAmazonS3BrowserDestinationSpec) Get() *AmazonS3BrowserDestinationSpec {
	return v.value
}

func (v *NullableAmazonS3BrowserDestinationSpec) Set(val *AmazonS3BrowserDestinationSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonS3BrowserDestinationSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonS3BrowserDestinationSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonS3BrowserDestinationSpec(val *AmazonS3BrowserDestinationSpec) *NullableAmazonS3BrowserDestinationSpec {
	return &NullableAmazonS3BrowserDestinationSpec{value: val, isSet: true}
}

func (v NullableAmazonS3BrowserDestinationSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonS3BrowserDestinationSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


