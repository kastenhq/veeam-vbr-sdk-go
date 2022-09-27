# AzureBlobStorageModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableTaskLimit** | Pointer to **bool** |  | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 
**Account** | [**AzureBlobStorageAccountModel**](AzureBlobStorageAccountModel.md) |  | 
**Container** | [**AzureBlobStorageContainerModel**](AzureBlobStorageContainerModel.md) |  | 
**ProxyAppliance** | Pointer to [**AzureStorageProxyModel**](AzureStorageProxyModel.md) |  | [optional] 
**MountServer** | [**MountServerSettingsModel**](MountServerSettingsModel.md) |  | 

## Methods

### NewAzureBlobStorageModelAllOf

`func NewAzureBlobStorageModelAllOf(account AzureBlobStorageAccountModel, container AzureBlobStorageContainerModel, mountServer MountServerSettingsModel, ) *AzureBlobStorageModelAllOf`

NewAzureBlobStorageModelAllOf instantiates a new AzureBlobStorageModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureBlobStorageModelAllOfWithDefaults

`func NewAzureBlobStorageModelAllOfWithDefaults() *AzureBlobStorageModelAllOf`

NewAzureBlobStorageModelAllOfWithDefaults instantiates a new AzureBlobStorageModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableTaskLimit

`func (o *AzureBlobStorageModelAllOf) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *AzureBlobStorageModelAllOf) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *AzureBlobStorageModelAllOf) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *AzureBlobStorageModelAllOf) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *AzureBlobStorageModelAllOf) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *AzureBlobStorageModelAllOf) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *AzureBlobStorageModelAllOf) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *AzureBlobStorageModelAllOf) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetAccount

`func (o *AzureBlobStorageModelAllOf) GetAccount() AzureBlobStorageAccountModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AzureBlobStorageModelAllOf) GetAccountOk() (*AzureBlobStorageAccountModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AzureBlobStorageModelAllOf) SetAccount(v AzureBlobStorageAccountModel)`

SetAccount sets Account field to given value.


### GetContainer

`func (o *AzureBlobStorageModelAllOf) GetContainer() AzureBlobStorageContainerModel`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *AzureBlobStorageModelAllOf) GetContainerOk() (*AzureBlobStorageContainerModel, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *AzureBlobStorageModelAllOf) SetContainer(v AzureBlobStorageContainerModel)`

SetContainer sets Container field to given value.


### GetProxyAppliance

`func (o *AzureBlobStorageModelAllOf) GetProxyAppliance() AzureStorageProxyModel`

GetProxyAppliance returns the ProxyAppliance field if non-nil, zero value otherwise.

### GetProxyApplianceOk

`func (o *AzureBlobStorageModelAllOf) GetProxyApplianceOk() (*AzureStorageProxyModel, bool)`

GetProxyApplianceOk returns a tuple with the ProxyAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAppliance

`func (o *AzureBlobStorageModelAllOf) SetProxyAppliance(v AzureStorageProxyModel)`

SetProxyAppliance sets ProxyAppliance field to given value.

### HasProxyAppliance

`func (o *AzureBlobStorageModelAllOf) HasProxyAppliance() bool`

HasProxyAppliance returns a boolean if a field has been set.

### GetMountServer

`func (o *AzureBlobStorageModelAllOf) GetMountServer() MountServerSettingsModel`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *AzureBlobStorageModelAllOf) GetMountServerOk() (*MountServerSettingsModel, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *AzureBlobStorageModelAllOf) SetMountServer(v MountServerSettingsModel)`

SetMountServer sets MountServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


