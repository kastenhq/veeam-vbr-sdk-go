# IBMCloudStorageAccountImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePoint** | **string** |  | 
**RegionId** | **string** |  | 
**Credentials** | [**CredentialsImportModel**](CredentialsImportModel.md) |  | 
**GatewayServer** | [**RepositoryShareGatewayImportSpec**](RepositoryShareGatewayImportSpec.md) |  | 

## Methods

### NewIBMCloudStorageAccountImportModel

`func NewIBMCloudStorageAccountImportModel(servicePoint string, regionId string, credentials CredentialsImportModel, gatewayServer RepositoryShareGatewayImportSpec, ) *IBMCloudStorageAccountImportModel`

NewIBMCloudStorageAccountImportModel instantiates a new IBMCloudStorageAccountImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIBMCloudStorageAccountImportModelWithDefaults

`func NewIBMCloudStorageAccountImportModelWithDefaults() *IBMCloudStorageAccountImportModel`

NewIBMCloudStorageAccountImportModelWithDefaults instantiates a new IBMCloudStorageAccountImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePoint

`func (o *IBMCloudStorageAccountImportModel) GetServicePoint() string`

GetServicePoint returns the ServicePoint field if non-nil, zero value otherwise.

### GetServicePointOk

`func (o *IBMCloudStorageAccountImportModel) GetServicePointOk() (*string, bool)`

GetServicePointOk returns a tuple with the ServicePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePoint

`func (o *IBMCloudStorageAccountImportModel) SetServicePoint(v string)`

SetServicePoint sets ServicePoint field to given value.


### GetRegionId

`func (o *IBMCloudStorageAccountImportModel) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *IBMCloudStorageAccountImportModel) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *IBMCloudStorageAccountImportModel) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetCredentials

`func (o *IBMCloudStorageAccountImportModel) GetCredentials() CredentialsImportModel`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *IBMCloudStorageAccountImportModel) GetCredentialsOk() (*CredentialsImportModel, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *IBMCloudStorageAccountImportModel) SetCredentials(v CredentialsImportModel)`

SetCredentials sets Credentials field to given value.


### GetGatewayServer

`func (o *IBMCloudStorageAccountImportModel) GetGatewayServer() RepositoryShareGatewayImportSpec`

GetGatewayServer returns the GatewayServer field if non-nil, zero value otherwise.

### GetGatewayServerOk

`func (o *IBMCloudStorageAccountImportModel) GetGatewayServerOk() (*RepositoryShareGatewayImportSpec, bool)`

GetGatewayServerOk returns a tuple with the GatewayServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayServer

`func (o *IBMCloudStorageAccountImportModel) SetGatewayServer(v RepositoryShareGatewayImportSpec)`

SetGatewayServer sets GatewayServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


