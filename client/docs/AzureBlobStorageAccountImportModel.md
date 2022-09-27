# AzureBlobStorageAccountImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | [**CredentialsImportModel**](CredentialsImportModel.md) |  | 
**RegionType** | [**EAzureRegionType**](EAzureRegionType.md) |  | 
**GatewayServer** | [**RepositoryShareGatewayImportSpec**](RepositoryShareGatewayImportSpec.md) |  | 

## Methods

### NewAzureBlobStorageAccountImportModel

`func NewAzureBlobStorageAccountImportModel(credentials CredentialsImportModel, regionType EAzureRegionType, gatewayServer RepositoryShareGatewayImportSpec, ) *AzureBlobStorageAccountImportModel`

NewAzureBlobStorageAccountImportModel instantiates a new AzureBlobStorageAccountImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureBlobStorageAccountImportModelWithDefaults

`func NewAzureBlobStorageAccountImportModelWithDefaults() *AzureBlobStorageAccountImportModel`

NewAzureBlobStorageAccountImportModelWithDefaults instantiates a new AzureBlobStorageAccountImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *AzureBlobStorageAccountImportModel) GetCredentials() CredentialsImportModel`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AzureBlobStorageAccountImportModel) GetCredentialsOk() (*CredentialsImportModel, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AzureBlobStorageAccountImportModel) SetCredentials(v CredentialsImportModel)`

SetCredentials sets Credentials field to given value.


### GetRegionType

`func (o *AzureBlobStorageAccountImportModel) GetRegionType() EAzureRegionType`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *AzureBlobStorageAccountImportModel) GetRegionTypeOk() (*EAzureRegionType, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *AzureBlobStorageAccountImportModel) SetRegionType(v EAzureRegionType)`

SetRegionType sets RegionType field to given value.


### GetGatewayServer

`func (o *AzureBlobStorageAccountImportModel) GetGatewayServer() RepositoryShareGatewayImportSpec`

GetGatewayServer returns the GatewayServer field if non-nil, zero value otherwise.

### GetGatewayServerOk

`func (o *AzureBlobStorageAccountImportModel) GetGatewayServerOk() (*RepositoryShareGatewayImportSpec, bool)`

GetGatewayServerOk returns a tuple with the GatewayServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayServer

`func (o *AzureBlobStorageAccountImportModel) SetGatewayServer(v RepositoryShareGatewayImportSpec)`

SetGatewayServer sets GatewayServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


