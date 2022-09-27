# AmazonS3StorageAccountImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | [**CredentialsImportModel**](CredentialsImportModel.md) |  | 
**RegionType** | [**EAmazonRegionType**](EAmazonRegionType.md) |  | 
**GatewayServer** | [**RepositoryShareGatewayImportSpec**](RepositoryShareGatewayImportSpec.md) |  | 

## Methods

### NewAmazonS3StorageAccountImportModel

`func NewAmazonS3StorageAccountImportModel(credentials CredentialsImportModel, regionType EAmazonRegionType, gatewayServer RepositoryShareGatewayImportSpec, ) *AmazonS3StorageAccountImportModel`

NewAmazonS3StorageAccountImportModel instantiates a new AmazonS3StorageAccountImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonS3StorageAccountImportModelWithDefaults

`func NewAmazonS3StorageAccountImportModelWithDefaults() *AmazonS3StorageAccountImportModel`

NewAmazonS3StorageAccountImportModelWithDefaults instantiates a new AmazonS3StorageAccountImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *AmazonS3StorageAccountImportModel) GetCredentials() CredentialsImportModel`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AmazonS3StorageAccountImportModel) GetCredentialsOk() (*CredentialsImportModel, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AmazonS3StorageAccountImportModel) SetCredentials(v CredentialsImportModel)`

SetCredentials sets Credentials field to given value.


### GetRegionType

`func (o *AmazonS3StorageAccountImportModel) GetRegionType() EAmazonRegionType`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *AmazonS3StorageAccountImportModel) GetRegionTypeOk() (*EAmazonRegionType, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *AmazonS3StorageAccountImportModel) SetRegionType(v EAmazonRegionType)`

SetRegionType sets RegionType field to given value.


### GetGatewayServer

`func (o *AmazonS3StorageAccountImportModel) GetGatewayServer() RepositoryShareGatewayImportSpec`

GetGatewayServer returns the GatewayServer field if non-nil, zero value otherwise.

### GetGatewayServerOk

`func (o *AmazonS3StorageAccountImportModel) GetGatewayServerOk() (*RepositoryShareGatewayImportSpec, bool)`

GetGatewayServerOk returns a tuple with the GatewayServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayServer

`func (o *AmazonS3StorageAccountImportModel) SetGatewayServer(v RepositoryShareGatewayImportSpec)`

SetGatewayServer sets GatewayServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


