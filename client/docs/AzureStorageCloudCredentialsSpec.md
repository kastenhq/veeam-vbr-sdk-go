# AzureStorageCloudCredentialsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Storage account type. | [optional] 
**SharedKey** | Pointer to **string** | Storage account shared key. | [optional] 
**Tag** | Pointer to **string** | Tag used to identify the credentials record. | [optional] 

## Methods

### NewAzureStorageCloudCredentialsSpec

`func NewAzureStorageCloudCredentialsSpec() *AzureStorageCloudCredentialsSpec`

NewAzureStorageCloudCredentialsSpec instantiates a new AzureStorageCloudCredentialsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureStorageCloudCredentialsSpecWithDefaults

`func NewAzureStorageCloudCredentialsSpecWithDefaults() *AzureStorageCloudCredentialsSpec`

NewAzureStorageCloudCredentialsSpecWithDefaults instantiates a new AzureStorageCloudCredentialsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AzureStorageCloudCredentialsSpec) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AzureStorageCloudCredentialsSpec) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AzureStorageCloudCredentialsSpec) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AzureStorageCloudCredentialsSpec) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSharedKey

`func (o *AzureStorageCloudCredentialsSpec) GetSharedKey() string`

GetSharedKey returns the SharedKey field if non-nil, zero value otherwise.

### GetSharedKeyOk

`func (o *AzureStorageCloudCredentialsSpec) GetSharedKeyOk() (*string, bool)`

GetSharedKeyOk returns a tuple with the SharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedKey

`func (o *AzureStorageCloudCredentialsSpec) SetSharedKey(v string)`

SetSharedKey sets SharedKey field to given value.

### HasSharedKey

`func (o *AzureStorageCloudCredentialsSpec) HasSharedKey() bool`

HasSharedKey returns a boolean if a field has been set.

### GetTag

`func (o *AzureStorageCloudCredentialsSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AzureStorageCloudCredentialsSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AzureStorageCloudCredentialsSpec) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *AzureStorageCloudCredentialsSpec) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


