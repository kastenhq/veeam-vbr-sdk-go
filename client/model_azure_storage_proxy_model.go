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

// AzureStorageProxyModel Azure storage proxy appliance.
type AzureStorageProxyModel struct {
	// ID that Veeam Backup & Replication assigned to the Azure subscription.
	SubscriptionId string `json:"subscriptionId"`
	// Size of the appliance.
	InstanceSize *string `json:"instanceSize,omitempty"`
	// Resource group associated with the proxy appliance.
	ResourceGroup *string `json:"resourceGroup,omitempty"`
	// Network to which the helper appliance is connected.
	VirtualNetwork *string `json:"virtualNetwork,omitempty"`
	// Subnet for the proxy appliance.
	Subnet *string `json:"subnet,omitempty"`
	// TCP port used to route requests between the proxy appliance and backup infrastructure components.
	RedirectorPort *int32 `json:"redirectorPort,omitempty"`
}

// NewAzureStorageProxyModel instantiates a new AzureStorageProxyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureStorageProxyModel(subscriptionId string) *AzureStorageProxyModel {
	this := AzureStorageProxyModel{}
	this.SubscriptionId = subscriptionId
	return &this
}

// NewAzureStorageProxyModelWithDefaults instantiates a new AzureStorageProxyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureStorageProxyModelWithDefaults() *AzureStorageProxyModel {
	this := AzureStorageProxyModel{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *AzureStorageProxyModel) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *AzureStorageProxyModel) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *AzureStorageProxyModel) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetInstanceSize returns the InstanceSize field value if set, zero value otherwise.
func (o *AzureStorageProxyModel) GetInstanceSize() string {
	if o == nil || isNil(o.InstanceSize) {
		var ret string
		return ret
	}
	return *o.InstanceSize
}

// GetInstanceSizeOk returns a tuple with the InstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureStorageProxyModel) GetInstanceSizeOk() (*string, bool) {
	if o == nil || isNil(o.InstanceSize) {
    return nil, false
	}
	return o.InstanceSize, true
}

// HasInstanceSize returns a boolean if a field has been set.
func (o *AzureStorageProxyModel) HasInstanceSize() bool {
	if o != nil && !isNil(o.InstanceSize) {
		return true
	}

	return false
}

// SetInstanceSize gets a reference to the given string and assigns it to the InstanceSize field.
func (o *AzureStorageProxyModel) SetInstanceSize(v string) {
	o.InstanceSize = &v
}

// GetResourceGroup returns the ResourceGroup field value if set, zero value otherwise.
func (o *AzureStorageProxyModel) GetResourceGroup() string {
	if o == nil || isNil(o.ResourceGroup) {
		var ret string
		return ret
	}
	return *o.ResourceGroup
}

// GetResourceGroupOk returns a tuple with the ResourceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureStorageProxyModel) GetResourceGroupOk() (*string, bool) {
	if o == nil || isNil(o.ResourceGroup) {
    return nil, false
	}
	return o.ResourceGroup, true
}

// HasResourceGroup returns a boolean if a field has been set.
func (o *AzureStorageProxyModel) HasResourceGroup() bool {
	if o != nil && !isNil(o.ResourceGroup) {
		return true
	}

	return false
}

// SetResourceGroup gets a reference to the given string and assigns it to the ResourceGroup field.
func (o *AzureStorageProxyModel) SetResourceGroup(v string) {
	o.ResourceGroup = &v
}

// GetVirtualNetwork returns the VirtualNetwork field value if set, zero value otherwise.
func (o *AzureStorageProxyModel) GetVirtualNetwork() string {
	if o == nil || isNil(o.VirtualNetwork) {
		var ret string
		return ret
	}
	return *o.VirtualNetwork
}

// GetVirtualNetworkOk returns a tuple with the VirtualNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureStorageProxyModel) GetVirtualNetworkOk() (*string, bool) {
	if o == nil || isNil(o.VirtualNetwork) {
    return nil, false
	}
	return o.VirtualNetwork, true
}

// HasVirtualNetwork returns a boolean if a field has been set.
func (o *AzureStorageProxyModel) HasVirtualNetwork() bool {
	if o != nil && !isNil(o.VirtualNetwork) {
		return true
	}

	return false
}

// SetVirtualNetwork gets a reference to the given string and assigns it to the VirtualNetwork field.
func (o *AzureStorageProxyModel) SetVirtualNetwork(v string) {
	o.VirtualNetwork = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *AzureStorageProxyModel) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureStorageProxyModel) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *AzureStorageProxyModel) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *AzureStorageProxyModel) SetSubnet(v string) {
	o.Subnet = &v
}

// GetRedirectorPort returns the RedirectorPort field value if set, zero value otherwise.
func (o *AzureStorageProxyModel) GetRedirectorPort() int32 {
	if o == nil || isNil(o.RedirectorPort) {
		var ret int32
		return ret
	}
	return *o.RedirectorPort
}

// GetRedirectorPortOk returns a tuple with the RedirectorPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureStorageProxyModel) GetRedirectorPortOk() (*int32, bool) {
	if o == nil || isNil(o.RedirectorPort) {
    return nil, false
	}
	return o.RedirectorPort, true
}

// HasRedirectorPort returns a boolean if a field has been set.
func (o *AzureStorageProxyModel) HasRedirectorPort() bool {
	if o != nil && !isNil(o.RedirectorPort) {
		return true
	}

	return false
}

// SetRedirectorPort gets a reference to the given int32 and assigns it to the RedirectorPort field.
func (o *AzureStorageProxyModel) SetRedirectorPort(v int32) {
	o.RedirectorPort = &v
}

func (o AzureStorageProxyModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !isNil(o.InstanceSize) {
		toSerialize["instanceSize"] = o.InstanceSize
	}
	if !isNil(o.ResourceGroup) {
		toSerialize["resourceGroup"] = o.ResourceGroup
	}
	if !isNil(o.VirtualNetwork) {
		toSerialize["virtualNetwork"] = o.VirtualNetwork
	}
	if !isNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !isNil(o.RedirectorPort) {
		toSerialize["redirectorPort"] = o.RedirectorPort
	}
	return json.Marshal(toSerialize)
}

type NullableAzureStorageProxyModel struct {
	value *AzureStorageProxyModel
	isSet bool
}

func (v NullableAzureStorageProxyModel) Get() *AzureStorageProxyModel {
	return v.value
}

func (v *NullableAzureStorageProxyModel) Set(val *AzureStorageProxyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureStorageProxyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureStorageProxyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureStorageProxyModel(val *AzureStorageProxyModel) *NullableAzureStorageProxyModel {
	return &NullableAzureStorageProxyModel{value: val, isSet: true}
}

func (v NullableAzureStorageProxyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureStorageProxyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


