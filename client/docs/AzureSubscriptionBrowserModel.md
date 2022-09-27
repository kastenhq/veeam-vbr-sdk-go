# AzureSubscriptionBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** |  | 
**OriginalSubscriptionId** | Pointer to **string** |  | [optional] 
**Locations** | Pointer to [**[]AzureLocationBrowserModel**](AzureLocationBrowserModel.md) |  | [optional] 

## Methods

### NewAzureSubscriptionBrowserModel

`func NewAzureSubscriptionBrowserModel(subscriptionId string, ) *AzureSubscriptionBrowserModel`

NewAzureSubscriptionBrowserModel instantiates a new AzureSubscriptionBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSubscriptionBrowserModelWithDefaults

`func NewAzureSubscriptionBrowserModelWithDefaults() *AzureSubscriptionBrowserModel`

NewAzureSubscriptionBrowserModelWithDefaults instantiates a new AzureSubscriptionBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *AzureSubscriptionBrowserModel) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureSubscriptionBrowserModel) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureSubscriptionBrowserModel) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetOriginalSubscriptionId

`func (o *AzureSubscriptionBrowserModel) GetOriginalSubscriptionId() string`

GetOriginalSubscriptionId returns the OriginalSubscriptionId field if non-nil, zero value otherwise.

### GetOriginalSubscriptionIdOk

`func (o *AzureSubscriptionBrowserModel) GetOriginalSubscriptionIdOk() (*string, bool)`

GetOriginalSubscriptionIdOk returns a tuple with the OriginalSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSubscriptionId

`func (o *AzureSubscriptionBrowserModel) SetOriginalSubscriptionId(v string)`

SetOriginalSubscriptionId sets OriginalSubscriptionId field to given value.

### HasOriginalSubscriptionId

`func (o *AzureSubscriptionBrowserModel) HasOriginalSubscriptionId() bool`

HasOriginalSubscriptionId returns a boolean if a field has been set.

### GetLocations

`func (o *AzureSubscriptionBrowserModel) GetLocations() []AzureLocationBrowserModel`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *AzureSubscriptionBrowserModel) GetLocationsOk() (*[]AzureLocationBrowserModel, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *AzureSubscriptionBrowserModel) SetLocations(v []AzureLocationBrowserModel)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *AzureSubscriptionBrowserModel) HasLocations() bool`

HasLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


