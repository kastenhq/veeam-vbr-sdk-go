# AmazonS3StorageModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableTaskLimit** | Pointer to **bool** |  | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 
**Account** | [**AmazonS3StorageAccountModel**](AmazonS3StorageAccountModel.md) |  | 
**Bucket** | [**AmazonS3StorageBucketModel**](AmazonS3StorageBucketModel.md) |  | 
**MountServer** | [**MountServerSettingsModel**](MountServerSettingsModel.md) |  | 
**ProxyAppliance** | Pointer to [**AmazonS3StorageProxyApplianceModel**](AmazonS3StorageProxyApplianceModel.md) |  | [optional] 

## Methods

### NewAmazonS3StorageModelAllOf

`func NewAmazonS3StorageModelAllOf(account AmazonS3StorageAccountModel, bucket AmazonS3StorageBucketModel, mountServer MountServerSettingsModel, ) *AmazonS3StorageModelAllOf`

NewAmazonS3StorageModelAllOf instantiates a new AmazonS3StorageModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonS3StorageModelAllOfWithDefaults

`func NewAmazonS3StorageModelAllOfWithDefaults() *AmazonS3StorageModelAllOf`

NewAmazonS3StorageModelAllOfWithDefaults instantiates a new AmazonS3StorageModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableTaskLimit

`func (o *AmazonS3StorageModelAllOf) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *AmazonS3StorageModelAllOf) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *AmazonS3StorageModelAllOf) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *AmazonS3StorageModelAllOf) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *AmazonS3StorageModelAllOf) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *AmazonS3StorageModelAllOf) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *AmazonS3StorageModelAllOf) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *AmazonS3StorageModelAllOf) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetAccount

`func (o *AmazonS3StorageModelAllOf) GetAccount() AmazonS3StorageAccountModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AmazonS3StorageModelAllOf) GetAccountOk() (*AmazonS3StorageAccountModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AmazonS3StorageModelAllOf) SetAccount(v AmazonS3StorageAccountModel)`

SetAccount sets Account field to given value.


### GetBucket

`func (o *AmazonS3StorageModelAllOf) GetBucket() AmazonS3StorageBucketModel`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *AmazonS3StorageModelAllOf) GetBucketOk() (*AmazonS3StorageBucketModel, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *AmazonS3StorageModelAllOf) SetBucket(v AmazonS3StorageBucketModel)`

SetBucket sets Bucket field to given value.


### GetMountServer

`func (o *AmazonS3StorageModelAllOf) GetMountServer() MountServerSettingsModel`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *AmazonS3StorageModelAllOf) GetMountServerOk() (*MountServerSettingsModel, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *AmazonS3StorageModelAllOf) SetMountServer(v MountServerSettingsModel)`

SetMountServer sets MountServer field to given value.


### GetProxyAppliance

`func (o *AmazonS3StorageModelAllOf) GetProxyAppliance() AmazonS3StorageProxyApplianceModel`

GetProxyAppliance returns the ProxyAppliance field if non-nil, zero value otherwise.

### GetProxyApplianceOk

`func (o *AmazonS3StorageModelAllOf) GetProxyApplianceOk() (*AmazonS3StorageProxyApplianceModel, bool)`

GetProxyApplianceOk returns a tuple with the ProxyAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAppliance

`func (o *AmazonS3StorageModelAllOf) SetProxyAppliance(v AmazonS3StorageProxyApplianceModel)`

SetProxyAppliance sets ProxyAppliance field to given value.

### HasProxyAppliance

`func (o *AmazonS3StorageModelAllOf) HasProxyAppliance() bool`

HasProxyAppliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


