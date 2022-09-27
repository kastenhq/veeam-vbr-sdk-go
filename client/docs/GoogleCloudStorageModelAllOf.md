# GoogleCloudStorageModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableTaskLimit** | Pointer to **bool** |  | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 
**Account** | [**GoogleCloudStorageAccountModel**](GoogleCloudStorageAccountModel.md) |  | 
**Bucket** | [**GoogleCloudStorageBucketModel**](GoogleCloudStorageBucketModel.md) |  | 
**MountServer** | [**MountServerSettingsModel**](MountServerSettingsModel.md) |  | 

## Methods

### NewGoogleCloudStorageModelAllOf

`func NewGoogleCloudStorageModelAllOf(account GoogleCloudStorageAccountModel, bucket GoogleCloudStorageBucketModel, mountServer MountServerSettingsModel, ) *GoogleCloudStorageModelAllOf`

NewGoogleCloudStorageModelAllOf instantiates a new GoogleCloudStorageModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudStorageModelAllOfWithDefaults

`func NewGoogleCloudStorageModelAllOfWithDefaults() *GoogleCloudStorageModelAllOf`

NewGoogleCloudStorageModelAllOfWithDefaults instantiates a new GoogleCloudStorageModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableTaskLimit

`func (o *GoogleCloudStorageModelAllOf) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *GoogleCloudStorageModelAllOf) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *GoogleCloudStorageModelAllOf) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *GoogleCloudStorageModelAllOf) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *GoogleCloudStorageModelAllOf) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *GoogleCloudStorageModelAllOf) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *GoogleCloudStorageModelAllOf) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *GoogleCloudStorageModelAllOf) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetAccount

`func (o *GoogleCloudStorageModelAllOf) GetAccount() GoogleCloudStorageAccountModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GoogleCloudStorageModelAllOf) GetAccountOk() (*GoogleCloudStorageAccountModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GoogleCloudStorageModelAllOf) SetAccount(v GoogleCloudStorageAccountModel)`

SetAccount sets Account field to given value.


### GetBucket

`func (o *GoogleCloudStorageModelAllOf) GetBucket() GoogleCloudStorageBucketModel`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *GoogleCloudStorageModelAllOf) GetBucketOk() (*GoogleCloudStorageBucketModel, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *GoogleCloudStorageModelAllOf) SetBucket(v GoogleCloudStorageBucketModel)`

SetBucket sets Bucket field to given value.


### GetMountServer

`func (o *GoogleCloudStorageModelAllOf) GetMountServer() MountServerSettingsModel`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *GoogleCloudStorageModelAllOf) GetMountServerOk() (*MountServerSettingsModel, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *GoogleCloudStorageModelAllOf) SetMountServer(v MountServerSettingsModel)`

SetMountServer sets MountServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


