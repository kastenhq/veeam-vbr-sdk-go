# IBMCloudStorageModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableTaskLimit** | Pointer to **bool** |  | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 
**Account** | [**IBMCloudStorageAccountModel**](IBMCloudStorageAccountModel.md) |  | 
**Bucket** | [**IBMCloudStorageBucketModel**](IBMCloudStorageBucketModel.md) |  | 
**MountServer** | [**MountServerSettingsModel**](MountServerSettingsModel.md) |  | 
**ProxyAppliance** | Pointer to [**S3CompatibleProxyModel**](S3CompatibleProxyModel.md) |  | [optional] 

## Methods

### NewIBMCloudStorageModelAllOf

`func NewIBMCloudStorageModelAllOf(account IBMCloudStorageAccountModel, bucket IBMCloudStorageBucketModel, mountServer MountServerSettingsModel, ) *IBMCloudStorageModelAllOf`

NewIBMCloudStorageModelAllOf instantiates a new IBMCloudStorageModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIBMCloudStorageModelAllOfWithDefaults

`func NewIBMCloudStorageModelAllOfWithDefaults() *IBMCloudStorageModelAllOf`

NewIBMCloudStorageModelAllOfWithDefaults instantiates a new IBMCloudStorageModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableTaskLimit

`func (o *IBMCloudStorageModelAllOf) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *IBMCloudStorageModelAllOf) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *IBMCloudStorageModelAllOf) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *IBMCloudStorageModelAllOf) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *IBMCloudStorageModelAllOf) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *IBMCloudStorageModelAllOf) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *IBMCloudStorageModelAllOf) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *IBMCloudStorageModelAllOf) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetAccount

`func (o *IBMCloudStorageModelAllOf) GetAccount() IBMCloudStorageAccountModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IBMCloudStorageModelAllOf) GetAccountOk() (*IBMCloudStorageAccountModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IBMCloudStorageModelAllOf) SetAccount(v IBMCloudStorageAccountModel)`

SetAccount sets Account field to given value.


### GetBucket

`func (o *IBMCloudStorageModelAllOf) GetBucket() IBMCloudStorageBucketModel`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *IBMCloudStorageModelAllOf) GetBucketOk() (*IBMCloudStorageBucketModel, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *IBMCloudStorageModelAllOf) SetBucket(v IBMCloudStorageBucketModel)`

SetBucket sets Bucket field to given value.


### GetMountServer

`func (o *IBMCloudStorageModelAllOf) GetMountServer() MountServerSettingsModel`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *IBMCloudStorageModelAllOf) GetMountServerOk() (*MountServerSettingsModel, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *IBMCloudStorageModelAllOf) SetMountServer(v MountServerSettingsModel)`

SetMountServer sets MountServer field to given value.


### GetProxyAppliance

`func (o *IBMCloudStorageModelAllOf) GetProxyAppliance() S3CompatibleProxyModel`

GetProxyAppliance returns the ProxyAppliance field if non-nil, zero value otherwise.

### GetProxyApplianceOk

`func (o *IBMCloudStorageModelAllOf) GetProxyApplianceOk() (*S3CompatibleProxyModel, bool)`

GetProxyApplianceOk returns a tuple with the ProxyAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAppliance

`func (o *IBMCloudStorageModelAllOf) SetProxyAppliance(v S3CompatibleProxyModel)`

SetProxyAppliance sets ProxyAppliance field to given value.

### HasProxyAppliance

`func (o *IBMCloudStorageModelAllOf) HasProxyAppliance() bool`

HasProxyAppliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


