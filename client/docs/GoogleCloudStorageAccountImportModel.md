# GoogleCloudStorageAccountImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | [**CredentialsImportModel**](CredentialsImportModel.md) |  | 
**GatewayServer** | [**RepositoryShareGatewayImportSpec**](RepositoryShareGatewayImportSpec.md) |  | 

## Methods

### NewGoogleCloudStorageAccountImportModel

`func NewGoogleCloudStorageAccountImportModel(credentials CredentialsImportModel, gatewayServer RepositoryShareGatewayImportSpec, ) *GoogleCloudStorageAccountImportModel`

NewGoogleCloudStorageAccountImportModel instantiates a new GoogleCloudStorageAccountImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudStorageAccountImportModelWithDefaults

`func NewGoogleCloudStorageAccountImportModelWithDefaults() *GoogleCloudStorageAccountImportModel`

NewGoogleCloudStorageAccountImportModelWithDefaults instantiates a new GoogleCloudStorageAccountImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *GoogleCloudStorageAccountImportModel) GetCredentials() CredentialsImportModel`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GoogleCloudStorageAccountImportModel) GetCredentialsOk() (*CredentialsImportModel, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GoogleCloudStorageAccountImportModel) SetCredentials(v CredentialsImportModel)`

SetCredentials sets Credentials field to given value.


### GetGatewayServer

`func (o *GoogleCloudStorageAccountImportModel) GetGatewayServer() RepositoryShareGatewayImportSpec`

GetGatewayServer returns the GatewayServer field if non-nil, zero value otherwise.

### GetGatewayServerOk

`func (o *GoogleCloudStorageAccountImportModel) GetGatewayServerOk() (*RepositoryShareGatewayImportSpec, bool)`

GetGatewayServerOk returns a tuple with the GatewayServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayServer

`func (o *GoogleCloudStorageAccountImportModel) SetGatewayServer(v RepositoryShareGatewayImportSpec)`

SetGatewayServer sets GatewayServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


