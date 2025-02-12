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

// AmazonSnowballEdgeBrowserSpec struct for AmazonSnowballEdgeBrowserSpec
type AmazonSnowballEdgeBrowserSpec struct {
	CloudBrowserSpec
	// To connect to the AWS Snowball Edge device, specify the `snow` value.
	RegionId *string `json:"regionId,omitempty"`
	// Service point address and port number of the AWS Snowball Edge device.
	ServicePoint *string `json:"servicePoint,omitempty"`
	// ID of a gateway server you want to use to connect to the object storage. Specify this parameter to check internet connection of the server. As a gateway server you can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used.
	GatewayServerId *string `json:"gatewayServerId,omitempty"`
}

// NewAmazonSnowballEdgeBrowserSpec instantiates a new AmazonSnowballEdgeBrowserSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonSnowballEdgeBrowserSpec(credentialsId string, serviceType ECloudServiceType) *AmazonSnowballEdgeBrowserSpec {
	this := AmazonSnowballEdgeBrowserSpec{}
	this.CredentialsId = credentialsId
	this.ServiceType = serviceType
	return &this
}

// NewAmazonSnowballEdgeBrowserSpecWithDefaults instantiates a new AmazonSnowballEdgeBrowserSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonSnowballEdgeBrowserSpecWithDefaults() *AmazonSnowballEdgeBrowserSpec {
	this := AmazonSnowballEdgeBrowserSpec{}
	return &this
}

// GetRegionId returns the RegionId field value if set, zero value otherwise.
func (o *AmazonSnowballEdgeBrowserSpec) GetRegionId() string {
	if o == nil || isNil(o.RegionId) {
		var ret string
		return ret
	}
	return *o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeBrowserSpec) GetRegionIdOk() (*string, bool) {
	if o == nil || isNil(o.RegionId) {
    return nil, false
	}
	return o.RegionId, true
}

// HasRegionId returns a boolean if a field has been set.
func (o *AmazonSnowballEdgeBrowserSpec) HasRegionId() bool {
	if o != nil && !isNil(o.RegionId) {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given string and assigns it to the RegionId field.
func (o *AmazonSnowballEdgeBrowserSpec) SetRegionId(v string) {
	o.RegionId = &v
}

// GetServicePoint returns the ServicePoint field value if set, zero value otherwise.
func (o *AmazonSnowballEdgeBrowserSpec) GetServicePoint() string {
	if o == nil || isNil(o.ServicePoint) {
		var ret string
		return ret
	}
	return *o.ServicePoint
}

// GetServicePointOk returns a tuple with the ServicePoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeBrowserSpec) GetServicePointOk() (*string, bool) {
	if o == nil || isNil(o.ServicePoint) {
    return nil, false
	}
	return o.ServicePoint, true
}

// HasServicePoint returns a boolean if a field has been set.
func (o *AmazonSnowballEdgeBrowserSpec) HasServicePoint() bool {
	if o != nil && !isNil(o.ServicePoint) {
		return true
	}

	return false
}

// SetServicePoint gets a reference to the given string and assigns it to the ServicePoint field.
func (o *AmazonSnowballEdgeBrowserSpec) SetServicePoint(v string) {
	o.ServicePoint = &v
}

// GetGatewayServerId returns the GatewayServerId field value if set, zero value otherwise.
func (o *AmazonSnowballEdgeBrowserSpec) GetGatewayServerId() string {
	if o == nil || isNil(o.GatewayServerId) {
		var ret string
		return ret
	}
	return *o.GatewayServerId
}

// GetGatewayServerIdOk returns a tuple with the GatewayServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeBrowserSpec) GetGatewayServerIdOk() (*string, bool) {
	if o == nil || isNil(o.GatewayServerId) {
    return nil, false
	}
	return o.GatewayServerId, true
}

// HasGatewayServerId returns a boolean if a field has been set.
func (o *AmazonSnowballEdgeBrowserSpec) HasGatewayServerId() bool {
	if o != nil && !isNil(o.GatewayServerId) {
		return true
	}

	return false
}

// SetGatewayServerId gets a reference to the given string and assigns it to the GatewayServerId field.
func (o *AmazonSnowballEdgeBrowserSpec) SetGatewayServerId(v string) {
	o.GatewayServerId = &v
}

func (o AmazonSnowballEdgeBrowserSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBrowserSpec, errCloudBrowserSpec := json.Marshal(o.CloudBrowserSpec)
	if errCloudBrowserSpec != nil {
		return []byte{}, errCloudBrowserSpec
	}
	errCloudBrowserSpec = json.Unmarshal([]byte(serializedCloudBrowserSpec), &toSerialize)
	if errCloudBrowserSpec != nil {
		return []byte{}, errCloudBrowserSpec
	}
	if !isNil(o.RegionId) {
		toSerialize["regionId"] = o.RegionId
	}
	if !isNil(o.ServicePoint) {
		toSerialize["servicePoint"] = o.ServicePoint
	}
	if !isNil(o.GatewayServerId) {
		toSerialize["gatewayServerId"] = o.GatewayServerId
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonSnowballEdgeBrowserSpec struct {
	value *AmazonSnowballEdgeBrowserSpec
	isSet bool
}

func (v NullableAmazonSnowballEdgeBrowserSpec) Get() *AmazonSnowballEdgeBrowserSpec {
	return v.value
}

func (v *NullableAmazonSnowballEdgeBrowserSpec) Set(val *AmazonSnowballEdgeBrowserSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonSnowballEdgeBrowserSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonSnowballEdgeBrowserSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonSnowballEdgeBrowserSpec(val *AmazonSnowballEdgeBrowserSpec) *NullableAmazonSnowballEdgeBrowserSpec {
	return &NullableAmazonSnowballEdgeBrowserSpec{value: val, isSet: true}
}

func (v NullableAmazonSnowballEdgeBrowserSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonSnowballEdgeBrowserSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


