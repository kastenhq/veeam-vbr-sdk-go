# AzureLinuxHelperApplianceSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | ID of a Microsoft Azure subscription credentials. | 
**Location** | Pointer to **string** |  | [optional] 
**StorageAccount** | Pointer to **string** |  | [optional] 
**ResourceGroup** | Pointer to **string** | Resource group associated with the proxy appliance. | [optional] 
**VirtualNetwork** | Pointer to **string** | Network to which the helper appliance is connected. | [optional] 
**Subnet** | Pointer to **string** | Subnet for the proxy appliance. | [optional] 
**SSHPort** | Pointer to **int32** | TCP port used to route requests between the proxy appliance and backup infrastructure components. | [optional] 

## Methods

### NewAzureLinuxHelperApplianceSpecAllOf

`func NewAzureLinuxHelperApplianceSpecAllOf(subscriptionId string, ) *AzureLinuxHelperApplianceSpecAllOf`

NewAzureLinuxHelperApplianceSpecAllOf instantiates a new AzureLinuxHelperApplianceSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureLinuxHelperApplianceSpecAllOfWithDefaults

`func NewAzureLinuxHelperApplianceSpecAllOfWithDefaults() *AzureLinuxHelperApplianceSpecAllOf`

NewAzureLinuxHelperApplianceSpecAllOfWithDefaults instantiates a new AzureLinuxHelperApplianceSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *AzureLinuxHelperApplianceSpecAllOf) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureLinuxHelperApplianceSpecAllOf) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureLinuxHelperApplianceSpecAllOf) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetLocation

`func (o *AzureLinuxHelperApplianceSpecAllOf) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AzureLinuxHelperApplianceSpecAllOf) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AzureLinuxHelperApplianceSpecAllOf) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AzureLinuxHelperApplianceSpecAllOf) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStorageAccount

`func (o *AzureLinuxHelperApplianceSpecAllOf) GetStorageAccount() string`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *AzureLinuxHelperApplianceSpecAllOf) GetStorageAccountOk() (*string, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *AzureLinuxHelperApplianceSpecAllOf) SetStorageAccount(v string)`

SetStorageAccount sets StorageAccount field to given value.

### HasStorageAccount

`func (o *AzureLinuxHelperApplianceSpecAllOf) HasStorageAccount() bool`

HasStorageAccount returns a boolean if a field has been set.

### GetResourceGroup

`func (o *AzureLinuxHelperApplianceSpecAllOf) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AzureLinuxHelperApplianceSpecAllOf) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AzureLinuxHelperApplianceSpecAllOf) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *AzureLinuxHelperApplianceSpecAllOf) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetVirtualNetwork

`func (o *AzureLinuxHelperApplianceSpecAllOf) GetVirtualNetwork() string`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *AzureLinuxHelperApplianceSpecAllOf) GetVirtualNetworkOk() (*string, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *AzureLinuxHelperApplianceSpecAllOf) SetVirtualNetwork(v string)`

SetVirtualNetwork sets VirtualNetwork field to given value.

### HasVirtualNetwork

`func (o *AzureLinuxHelperApplianceSpecAllOf) HasVirtualNetwork() bool`

HasVirtualNetwork returns a boolean if a field has been set.

### GetSubnet

`func (o *AzureLinuxHelperApplianceSpecAllOf) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *AzureLinuxHelperApplianceSpecAllOf) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *AzureLinuxHelperApplianceSpecAllOf) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *AzureLinuxHelperApplianceSpecAllOf) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetSSHPort

`func (o *AzureLinuxHelperApplianceSpecAllOf) GetSSHPort() int32`

GetSSHPort returns the SSHPort field if non-nil, zero value otherwise.

### GetSSHPortOk

`func (o *AzureLinuxHelperApplianceSpecAllOf) GetSSHPortOk() (*int32, bool)`

GetSSHPortOk returns a tuple with the SSHPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSHPort

`func (o *AzureLinuxHelperApplianceSpecAllOf) SetSSHPort(v int32)`

SetSSHPort sets SSHPort field to given value.

### HasSSHPort

`func (o *AzureLinuxHelperApplianceSpecAllOf) HasSSHPort() bool`

HasSSHPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


