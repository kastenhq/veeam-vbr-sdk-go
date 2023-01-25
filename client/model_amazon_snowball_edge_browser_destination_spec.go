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

// AmazonSnowballEdgeBrowserDestinationSpec struct for AmazonSnowballEdgeBrowserDestinationSpec
type AmazonSnowballEdgeBrowserDestinationSpec struct {
	CloudBrowserNewFolderSpec
	// ID of a server you want to use to connect to the object storage. You can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used.
	HostId *string `json:"hostId,omitempty"`
	// Service point address and port number of the AWS Snowball Edge device.
	ConnectionPoint string `json:"connectionPoint"`
	// For AWS Snowball Edge, the region is *snow*.
	RegionId string `json:"regionId"`
	// Name of the bucket where you want to store your backup data.
	BucketName string `json:"bucketName"`
}

// NewAmazonSnowballEdgeBrowserDestinationSpec instantiates a new AmazonSnowballEdgeBrowserDestinationSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonSnowballEdgeBrowserDestinationSpec(connectionPoint string, regionId string, bucketName string, regionType EAmazonRegionType, containerName string, servicePoint string) *AmazonSnowballEdgeBrowserDestinationSpec {
	this := AmazonSnowballEdgeBrowserDestinationSpec{}
	this.RegionType = regionType
	this.ContainerName = containerName
	this.ServicePoint = servicePoint
	this.RegionId = regionId
	this.BucketName = bucketName
	this.ConnectionPoint = connectionPoint
	return &this
}

// NewAmazonSnowballEdgeBrowserDestinationSpecWithDefaults instantiates a new AmazonSnowballEdgeBrowserDestinationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonSnowballEdgeBrowserDestinationSpecWithDefaults() *AmazonSnowballEdgeBrowserDestinationSpec {
	this := AmazonSnowballEdgeBrowserDestinationSpec{}
	return &this
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *AmazonSnowballEdgeBrowserDestinationSpec) GetHostId() string {
	if o == nil || isNil(o.HostId) {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeBrowserDestinationSpec) GetHostIdOk() (*string, bool) {
	if o == nil || isNil(o.HostId) {
    return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *AmazonSnowballEdgeBrowserDestinationSpec) HasHostId() bool {
	if o != nil && !isNil(o.HostId) {
		return true
	}

	return false
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *AmazonSnowballEdgeBrowserDestinationSpec) SetHostId(v string) {
	o.HostId = &v
}

// GetConnectionPoint returns the ConnectionPoint field value
func (o *AmazonSnowballEdgeBrowserDestinationSpec) GetConnectionPoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionPoint
}

// GetConnectionPointOk returns a tuple with the ConnectionPoint field value
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeBrowserDestinationSpec) GetConnectionPointOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConnectionPoint, true
}

// SetConnectionPoint sets field value
func (o *AmazonSnowballEdgeBrowserDestinationSpec) SetConnectionPoint(v string) {
	o.ConnectionPoint = v
}

// GetRegionId returns the RegionId field value
func (o *AmazonSnowballEdgeBrowserDestinationSpec) GetRegionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeBrowserDestinationSpec) GetRegionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RegionId, true
}

// SetRegionId sets field value
func (o *AmazonSnowballEdgeBrowserDestinationSpec) SetRegionId(v string) {
	o.RegionId = v
}

// GetBucketName returns the BucketName field value
func (o *AmazonSnowballEdgeBrowserDestinationSpec) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeBrowserDestinationSpec) GetBucketNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value
func (o *AmazonSnowballEdgeBrowserDestinationSpec) SetBucketName(v string) {
	o.BucketName = v
}

func (o AmazonSnowballEdgeBrowserDestinationSpec) MarshalJSON() ([]byte, error) {
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
		toSerialize["connectionPoint"] = o.ConnectionPoint
	}
	if true {
		toSerialize["regionId"] = o.RegionId
	}
	if true {
		toSerialize["bucketName"] = o.BucketName
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonSnowballEdgeBrowserDestinationSpec struct {
	value *AmazonSnowballEdgeBrowserDestinationSpec
	isSet bool
}

func (v NullableAmazonSnowballEdgeBrowserDestinationSpec) Get() *AmazonSnowballEdgeBrowserDestinationSpec {
	return v.value
}

func (v *NullableAmazonSnowballEdgeBrowserDestinationSpec) Set(val *AmazonSnowballEdgeBrowserDestinationSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonSnowballEdgeBrowserDestinationSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonSnowballEdgeBrowserDestinationSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonSnowballEdgeBrowserDestinationSpec(val *AmazonSnowballEdgeBrowserDestinationSpec) *NullableAmazonSnowballEdgeBrowserDestinationSpec {
	return &NullableAmazonSnowballEdgeBrowserDestinationSpec{value: val, isSet: true}
}

func (v NullableAmazonSnowballEdgeBrowserDestinationSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonSnowballEdgeBrowserDestinationSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


