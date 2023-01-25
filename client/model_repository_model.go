/*
Veeam Backup & Replication REST API

This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br>Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Some request and response bodies refer to reusable schema objects that are defined in the `schemas` section of the REST API specification. Schemas, in turn, may inherit a part of their properties by referring to other schemas, and be polymorphic by referring to multiple alternate schemas.

API version: 1.1-rev0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// RepositoryModel - struct for RepositoryModel
type RepositoryModel struct {
	AmazonS3GlacierStorageModel *AmazonS3GlacierStorageModel
	AmazonS3StorageModel *AmazonS3StorageModel
	AmazonSnowballEdgeStorageModel *AmazonSnowballEdgeStorageModel
	AzureArchiveStorageModel *AzureArchiveStorageModel
	AzureBlobStorageModel *AzureBlobStorageModel
	AzureDataBoxStorageModel *AzureDataBoxStorageModel
	GoogleCloudStorageModel *GoogleCloudStorageModel
	IBMCloudStorageModel *IBMCloudStorageModel
	LinuxHardenedStorageModel *LinuxHardenedStorageModel
	LinuxLocalStorageModel *LinuxLocalStorageModel
	NfsStorageModel *NfsStorageModel
	S3CompatibleStorageModel *S3CompatibleStorageModel
	SmbStorageModel *SmbStorageModel
	WasabiCloudStorageModel *WasabiCloudStorageModel
	WindowsLocalStorageModel *WindowsLocalStorageModel
}

// AmazonS3GlacierStorageModelAsRepositoryModel is a convenience function that returns AmazonS3GlacierStorageModel wrapped in RepositoryModel
func AmazonS3GlacierStorageModelAsRepositoryModel(v *AmazonS3GlacierStorageModel) RepositoryModel {
	return RepositoryModel{
		AmazonS3GlacierStorageModel: v,
	}
}

// AmazonS3StorageModelAsRepositoryModel is a convenience function that returns AmazonS3StorageModel wrapped in RepositoryModel
func AmazonS3StorageModelAsRepositoryModel(v *AmazonS3StorageModel) RepositoryModel {
	return RepositoryModel{
		AmazonS3StorageModel: v,
	}
}

// AmazonSnowballEdgeStorageModelAsRepositoryModel is a convenience function that returns AmazonSnowballEdgeStorageModel wrapped in RepositoryModel
func AmazonSnowballEdgeStorageModelAsRepositoryModel(v *AmazonSnowballEdgeStorageModel) RepositoryModel {
	return RepositoryModel{
		AmazonSnowballEdgeStorageModel: v,
	}
}

// AzureArchiveStorageModelAsRepositoryModel is a convenience function that returns AzureArchiveStorageModel wrapped in RepositoryModel
func AzureArchiveStorageModelAsRepositoryModel(v *AzureArchiveStorageModel) RepositoryModel {
	return RepositoryModel{
		AzureArchiveStorageModel: v,
	}
}

// AzureBlobStorageModelAsRepositoryModel is a convenience function that returns AzureBlobStorageModel wrapped in RepositoryModel
func AzureBlobStorageModelAsRepositoryModel(v *AzureBlobStorageModel) RepositoryModel {
	return RepositoryModel{
		AzureBlobStorageModel: v,
	}
}

// AzureDataBoxStorageModelAsRepositoryModel is a convenience function that returns AzureDataBoxStorageModel wrapped in RepositoryModel
func AzureDataBoxStorageModelAsRepositoryModel(v *AzureDataBoxStorageModel) RepositoryModel {
	return RepositoryModel{
		AzureDataBoxStorageModel: v,
	}
}

// GoogleCloudStorageModelAsRepositoryModel is a convenience function that returns GoogleCloudStorageModel wrapped in RepositoryModel
func GoogleCloudStorageModelAsRepositoryModel(v *GoogleCloudStorageModel) RepositoryModel {
	return RepositoryModel{
		GoogleCloudStorageModel: v,
	}
}

// IBMCloudStorageModelAsRepositoryModel is a convenience function that returns IBMCloudStorageModel wrapped in RepositoryModel
func IBMCloudStorageModelAsRepositoryModel(v *IBMCloudStorageModel) RepositoryModel {
	return RepositoryModel{
		IBMCloudStorageModel: v,
	}
}

// LinuxHardenedStorageModelAsRepositoryModel is a convenience function that returns LinuxHardenedStorageModel wrapped in RepositoryModel
func LinuxHardenedStorageModelAsRepositoryModel(v *LinuxHardenedStorageModel) RepositoryModel {
	return RepositoryModel{
		LinuxHardenedStorageModel: v,
	}
}

// LinuxLocalStorageModelAsRepositoryModel is a convenience function that returns LinuxLocalStorageModel wrapped in RepositoryModel
func LinuxLocalStorageModelAsRepositoryModel(v *LinuxLocalStorageModel) RepositoryModel {
	return RepositoryModel{
		LinuxLocalStorageModel: v,
	}
}

// NfsStorageModelAsRepositoryModel is a convenience function that returns NfsStorageModel wrapped in RepositoryModel
func NfsStorageModelAsRepositoryModel(v *NfsStorageModel) RepositoryModel {
	return RepositoryModel{
		NfsStorageModel: v,
	}
}

// S3CompatibleStorageModelAsRepositoryModel is a convenience function that returns S3CompatibleStorageModel wrapped in RepositoryModel
func S3CompatibleStorageModelAsRepositoryModel(v *S3CompatibleStorageModel) RepositoryModel {
	return RepositoryModel{
		S3CompatibleStorageModel: v,
	}
}

// SmbStorageModelAsRepositoryModel is a convenience function that returns SmbStorageModel wrapped in RepositoryModel
func SmbStorageModelAsRepositoryModel(v *SmbStorageModel) RepositoryModel {
	return RepositoryModel{
		SmbStorageModel: v,
	}
}

// WasabiCloudStorageModelAsRepositoryModel is a convenience function that returns WasabiCloudStorageModel wrapped in RepositoryModel
func WasabiCloudStorageModelAsRepositoryModel(v *WasabiCloudStorageModel) RepositoryModel {
	return RepositoryModel{
		WasabiCloudStorageModel: v,
	}
}

// WindowsLocalStorageModelAsRepositoryModel is a convenience function that returns WindowsLocalStorageModel wrapped in RepositoryModel
func WindowsLocalStorageModelAsRepositoryModel(v *WindowsLocalStorageModel) RepositoryModel {
	return RepositoryModel{
		WindowsLocalStorageModel: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *RepositoryModel) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AmazonS3GlacierStorageModel
	err = newStrictDecoder(data).Decode(&dst.AmazonS3GlacierStorageModel)
	if err == nil {
		jsonAmazonS3GlacierStorageModel, _ := json.Marshal(dst.AmazonS3GlacierStorageModel)
		if string(jsonAmazonS3GlacierStorageModel) == "{}" { // empty struct
			dst.AmazonS3GlacierStorageModel = nil
		} else {
			match++
		}
	} else {
		dst.AmazonS3GlacierStorageModel = nil
	}

	// try to unmarshal data into AmazonS3StorageModel
	err = newStrictDecoder(data).Decode(&dst.AmazonS3StorageModel)
	if err == nil {
		jsonAmazonS3StorageModel, _ := json.Marshal(dst.AmazonS3StorageModel)
		if string(jsonAmazonS3StorageModel) == "{}" { // empty struct
			dst.AmazonS3StorageModel = nil
		} else {
			match++
		}
	} else {
		dst.AmazonS3StorageModel = nil
	}

	// try to unmarshal data into AmazonSnowballEdgeStorageModel
	err = newStrictDecoder(data).Decode(&dst.AmazonSnowballEdgeStorageModel)
	if err == nil {
		jsonAmazonSnowballEdgeStorageModel, _ := json.Marshal(dst.AmazonSnowballEdgeStorageModel)
		if string(jsonAmazonSnowballEdgeStorageModel) == "{}" { // empty struct
			dst.AmazonSnowballEdgeStorageModel = nil
		} else {
			match++
		}
	} else {
		dst.AmazonSnowballEdgeStorageModel = nil
	}

	// try to unmarshal data into AzureArchiveStorageModel
	err = newStrictDecoder(data).Decode(&dst.AzureArchiveStorageModel)
	if err == nil {
		jsonAzureArchiveStorageModel, _ := json.Marshal(dst.AzureArchiveStorageModel)
		if string(jsonAzureArchiveStorageModel) == "{}" { // empty struct
			dst.AzureArchiveStorageModel = nil
		} else {
			match++
		}
	} else {
		dst.AzureArchiveStorageModel = nil
	}

	// try to unmarshal data into AzureBlobStorageModel
	err = newStrictDecoder(data).Decode(&dst.AzureBlobStorageModel)
	if err == nil {
		jsonAzureBlobStorageModel, _ := json.Marshal(dst.AzureBlobStorageModel)
		if string(jsonAzureBlobStorageModel) == "{}" { // empty struct
			dst.AzureBlobStorageModel = nil
		} else {
			match++
		}
	} else {
		dst.AzureBlobStorageModel = nil
	}

	// try to unmarshal data into AzureDataBoxStorageModel
	err = newStrictDecoder(data).Decode(&dst.AzureDataBoxStorageModel)
	if err == nil {
		jsonAzureDataBoxStorageModel, _ := json.Marshal(dst.AzureDataBoxStorageModel)
		if string(jsonAzureDataBoxStorageModel) == "{}" { // empty struct
			dst.AzureDataBoxStorageModel = nil
		} else {
			match++
		}
	} else {
		dst.AzureDataBoxStorageModel = nil
	}

	// try to unmarshal data into GoogleCloudStorageModel
	err = newStrictDecoder(data).Decode(&dst.GoogleCloudStorageModel)
	if err == nil {
		jsonGoogleCloudStorageModel, _ := json.Marshal(dst.GoogleCloudStorageModel)
		if string(jsonGoogleCloudStorageModel) == "{}" { // empty struct
			dst.GoogleCloudStorageModel = nil
		} else {
			match++
		}
	} else {
		dst.GoogleCloudStorageModel = nil
	}

	// try to unmarshal data into IBMCloudStorageModel
	err = newStrictDecoder(data).Decode(&dst.IBMCloudStorageModel)
	if err == nil {
		jsonIBMCloudStorageModel, _ := json.Marshal(dst.IBMCloudStorageModel)
		if string(jsonIBMCloudStorageModel) == "{}" { // empty struct
			dst.IBMCloudStorageModel = nil
		} else {
			match++
		}
	} else {
		dst.IBMCloudStorageModel = nil
	}

	// try to unmarshal data into LinuxHardenedStorageModel
	err = newStrictDecoder(data).Decode(&dst.LinuxHardenedStorageModel)
	if err == nil {
		jsonLinuxHardenedStorageModel, _ := json.Marshal(dst.LinuxHardenedStorageModel)
		if string(jsonLinuxHardenedStorageModel) == "{}" { // empty struct
			dst.LinuxHardenedStorageModel = nil
		} else {
			match++
		}
	} else {
		dst.LinuxHardenedStorageModel = nil
	}

	// try to unmarshal data into LinuxLocalStorageModel
	err = newStrictDecoder(data).Decode(&dst.LinuxLocalStorageModel)
	if err == nil {
		jsonLinuxLocalStorageModel, _ := json.Marshal(dst.LinuxLocalStorageModel)
		if string(jsonLinuxLocalStorageModel) == "{}" { // empty struct
			dst.LinuxLocalStorageModel = nil
		} else {
			match++
		}
	} else {
		dst.LinuxLocalStorageModel = nil
	}

	// try to unmarshal data into NfsStorageModel
	err = newStrictDecoder(data).Decode(&dst.NfsStorageModel)
	if err == nil {
		jsonNfsStorageModel, _ := json.Marshal(dst.NfsStorageModel)
		if string(jsonNfsStorageModel) == "{}" { // empty struct
			dst.NfsStorageModel = nil
		} else {
			match++
		}
	} else {
		dst.NfsStorageModel = nil
	}

	// try to unmarshal data into S3CompatibleStorageModel
	err = newStrictDecoder(data).Decode(&dst.S3CompatibleStorageModel)
	if err == nil {
		jsonS3CompatibleStorageModel, _ := json.Marshal(dst.S3CompatibleStorageModel)
		if string(jsonS3CompatibleStorageModel) == "{}" { // empty struct
			dst.S3CompatibleStorageModel = nil
		} else {
			match++
		}
	} else {
		dst.S3CompatibleStorageModel = nil
	}

	// try to unmarshal data into SmbStorageModel
	err = newStrictDecoder(data).Decode(&dst.SmbStorageModel)
	if err == nil {
		jsonSmbStorageModel, _ := json.Marshal(dst.SmbStorageModel)
		if string(jsonSmbStorageModel) == "{}" { // empty struct
			dst.SmbStorageModel = nil
		} else {
			match++
		}
	} else {
		dst.SmbStorageModel = nil
	}

	// try to unmarshal data into WasabiCloudStorageModel
	err = newStrictDecoder(data).Decode(&dst.WasabiCloudStorageModel)
	if err == nil {
		jsonWasabiCloudStorageModel, _ := json.Marshal(dst.WasabiCloudStorageModel)
		if string(jsonWasabiCloudStorageModel) == "{}" { // empty struct
			dst.WasabiCloudStorageModel = nil
		} else {
			match++
		}
	} else {
		dst.WasabiCloudStorageModel = nil
	}

	// try to unmarshal data into WindowsLocalStorageModel
	err = newStrictDecoder(data).Decode(&dst.WindowsLocalStorageModel)
	if err == nil {
		jsonWindowsLocalStorageModel, _ := json.Marshal(dst.WindowsLocalStorageModel)
		if string(jsonWindowsLocalStorageModel) == "{}" { // empty struct
			dst.WindowsLocalStorageModel = nil
		} else {
			match++
		}
	} else {
		dst.WindowsLocalStorageModel = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AmazonS3GlacierStorageModel = nil
		dst.AmazonS3StorageModel = nil
		dst.AmazonSnowballEdgeStorageModel = nil
		dst.AzureArchiveStorageModel = nil
		dst.AzureBlobStorageModel = nil
		dst.AzureDataBoxStorageModel = nil
		dst.GoogleCloudStorageModel = nil
		dst.IBMCloudStorageModel = nil
		dst.LinuxHardenedStorageModel = nil
		dst.LinuxLocalStorageModel = nil
		dst.NfsStorageModel = nil
		dst.S3CompatibleStorageModel = nil
		dst.SmbStorageModel = nil
		dst.WasabiCloudStorageModel = nil
		dst.WindowsLocalStorageModel = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RepositoryModel)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RepositoryModel)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RepositoryModel) MarshalJSON() ([]byte, error) {
	if src.AmazonS3GlacierStorageModel != nil {
		return json.Marshal(&src.AmazonS3GlacierStorageModel)
	}

	if src.AmazonS3StorageModel != nil {
		return json.Marshal(&src.AmazonS3StorageModel)
	}

	if src.AmazonSnowballEdgeStorageModel != nil {
		return json.Marshal(&src.AmazonSnowballEdgeStorageModel)
	}

	if src.AzureArchiveStorageModel != nil {
		return json.Marshal(&src.AzureArchiveStorageModel)
	}

	if src.AzureBlobStorageModel != nil {
		return json.Marshal(&src.AzureBlobStorageModel)
	}

	if src.AzureDataBoxStorageModel != nil {
		return json.Marshal(&src.AzureDataBoxStorageModel)
	}

	if src.GoogleCloudStorageModel != nil {
		return json.Marshal(&src.GoogleCloudStorageModel)
	}

	if src.IBMCloudStorageModel != nil {
		return json.Marshal(&src.IBMCloudStorageModel)
	}

	if src.LinuxHardenedStorageModel != nil {
		return json.Marshal(&src.LinuxHardenedStorageModel)
	}

	if src.LinuxLocalStorageModel != nil {
		return json.Marshal(&src.LinuxLocalStorageModel)
	}

	if src.NfsStorageModel != nil {
		return json.Marshal(&src.NfsStorageModel)
	}

	if src.S3CompatibleStorageModel != nil {
		return json.Marshal(&src.S3CompatibleStorageModel)
	}

	if src.SmbStorageModel != nil {
		return json.Marshal(&src.SmbStorageModel)
	}

	if src.WasabiCloudStorageModel != nil {
		return json.Marshal(&src.WasabiCloudStorageModel)
	}

	if src.WindowsLocalStorageModel != nil {
		return json.Marshal(&src.WindowsLocalStorageModel)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RepositoryModel) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AmazonS3GlacierStorageModel != nil {
		return obj.AmazonS3GlacierStorageModel
	}

	if obj.AmazonS3StorageModel != nil {
		return obj.AmazonS3StorageModel
	}

	if obj.AmazonSnowballEdgeStorageModel != nil {
		return obj.AmazonSnowballEdgeStorageModel
	}

	if obj.AzureArchiveStorageModel != nil {
		return obj.AzureArchiveStorageModel
	}

	if obj.AzureBlobStorageModel != nil {
		return obj.AzureBlobStorageModel
	}

	if obj.AzureDataBoxStorageModel != nil {
		return obj.AzureDataBoxStorageModel
	}

	if obj.GoogleCloudStorageModel != nil {
		return obj.GoogleCloudStorageModel
	}

	if obj.IBMCloudStorageModel != nil {
		return obj.IBMCloudStorageModel
	}

	if obj.LinuxHardenedStorageModel != nil {
		return obj.LinuxHardenedStorageModel
	}

	if obj.LinuxLocalStorageModel != nil {
		return obj.LinuxLocalStorageModel
	}

	if obj.NfsStorageModel != nil {
		return obj.NfsStorageModel
	}

	if obj.S3CompatibleStorageModel != nil {
		return obj.S3CompatibleStorageModel
	}

	if obj.SmbStorageModel != nil {
		return obj.SmbStorageModel
	}

	if obj.WasabiCloudStorageModel != nil {
		return obj.WasabiCloudStorageModel
	}

	if obj.WindowsLocalStorageModel != nil {
		return obj.WindowsLocalStorageModel
	}

	// all schemas are nil
	return nil
}

type NullableRepositoryModel struct {
	value *RepositoryModel
	isSet bool
}

func (v NullableRepositoryModel) Get() *RepositoryModel {
	return v.value
}

func (v *NullableRepositoryModel) Set(val *RepositoryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryModel(val *RepositoryModel) *NullableRepositoryModel {
	return &NullableRepositoryModel{value: val, isSet: true}
}

func (v NullableRepositoryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


