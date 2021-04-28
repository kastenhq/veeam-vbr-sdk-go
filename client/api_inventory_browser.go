/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 
 *
 * API version: 1.0-rev1
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// InventoryBrowserApiService InventoryBrowserApi service
type InventoryBrowserApiService service

type ApiGetAllInventoryVmwareHostsRequest struct {
	ctx _context.Context
	ApiService *InventoryBrowserApiService
	xApiVersion *string
	skip *int32
	limit *int32
	orderColumn *EViRootFiltersOrderColumn
	orderAsc *bool
	nameFilter *string
}

func (r ApiGetAllInventoryVmwareHostsRequest) XApiVersion(xApiVersion string) ApiGetAllInventoryVmwareHostsRequest {
	r.xApiVersion = &xApiVersion
	return r
}
func (r ApiGetAllInventoryVmwareHostsRequest) Skip(skip int32) ApiGetAllInventoryVmwareHostsRequest {
	r.skip = &skip
	return r
}
func (r ApiGetAllInventoryVmwareHostsRequest) Limit(limit int32) ApiGetAllInventoryVmwareHostsRequest {
	r.limit = &limit
	return r
}
func (r ApiGetAllInventoryVmwareHostsRequest) OrderColumn(orderColumn EViRootFiltersOrderColumn) ApiGetAllInventoryVmwareHostsRequest {
	r.orderColumn = &orderColumn
	return r
}
func (r ApiGetAllInventoryVmwareHostsRequest) OrderAsc(orderAsc bool) ApiGetAllInventoryVmwareHostsRequest {
	r.orderAsc = &orderAsc
	return r
}
func (r ApiGetAllInventoryVmwareHostsRequest) NameFilter(nameFilter string) ApiGetAllInventoryVmwareHostsRequest {
	r.nameFilter = &nameFilter
	return r
}

func (r ApiGetAllInventoryVmwareHostsRequest) Execute() (ViRootsResult, *_nethttp.Response, error) {
	return r.ApiService.GetAllInventoryVmwareHostsExecute(r)
}

/*
 * GetAllInventoryVmwareHosts Get All VMware vSphere Servers
 * The HTTP GET request to the `/api/v1/inventory/vmware/hosts` path allows you to get an array of all VMware vSphere servers added to the backup infrastructure.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetAllInventoryVmwareHostsRequest
 */
func (a *InventoryBrowserApiService) GetAllInventoryVmwareHosts(ctx _context.Context) ApiGetAllInventoryVmwareHostsRequest {
	return ApiGetAllInventoryVmwareHostsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ViRootsResult
 */
func (a *InventoryBrowserApiService) GetAllInventoryVmwareHostsExecute(r ApiGetAllInventoryVmwareHostsRequest) (ViRootsResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ViRootsResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryBrowserApiService.GetAllInventoryVmwareHosts")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/inventory/vmware/hosts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xApiVersion == nil {
		return localVarReturnValue, nil, reportError("xApiVersion is required and must be specified")
	}

	if r.skip != nil {
		localVarQueryParams.Add("skip", parameterToString(*r.skip, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.orderColumn != nil {
		localVarQueryParams.Add("orderColumn", parameterToString(*r.orderColumn, ""))
	}
	if r.orderAsc != nil {
		localVarQueryParams.Add("orderAsc", parameterToString(*r.orderAsc, ""))
	}
	if r.nameFilter != nil {
		localVarQueryParams.Add("nameFilter", parameterToString(*r.nameFilter, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-api-version"] = parameterToString(*r.xApiVersion, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetVmwareHostObjectRequest struct {
	ctx _context.Context
	ApiService *InventoryBrowserApiService
	xApiVersion *string
	name string
	skip *int32
	limit *int32
	orderColumn *EvCentersInventoryFiltersOrderColumn
	orderAsc *bool
	objectIdFilter *string
	hierarchyTypeFilter *EHierarchyType
	nameFilter *string
	typeFilter *EVmwareInventoryType
	parentContainerNameFilter *string
}

func (r ApiGetVmwareHostObjectRequest) XApiVersion(xApiVersion string) ApiGetVmwareHostObjectRequest {
	r.xApiVersion = &xApiVersion
	return r
}
func (r ApiGetVmwareHostObjectRequest) Skip(skip int32) ApiGetVmwareHostObjectRequest {
	r.skip = &skip
	return r
}
func (r ApiGetVmwareHostObjectRequest) Limit(limit int32) ApiGetVmwareHostObjectRequest {
	r.limit = &limit
	return r
}
func (r ApiGetVmwareHostObjectRequest) OrderColumn(orderColumn EvCentersInventoryFiltersOrderColumn) ApiGetVmwareHostObjectRequest {
	r.orderColumn = &orderColumn
	return r
}
func (r ApiGetVmwareHostObjectRequest) OrderAsc(orderAsc bool) ApiGetVmwareHostObjectRequest {
	r.orderAsc = &orderAsc
	return r
}
func (r ApiGetVmwareHostObjectRequest) ObjectIdFilter(objectIdFilter string) ApiGetVmwareHostObjectRequest {
	r.objectIdFilter = &objectIdFilter
	return r
}
func (r ApiGetVmwareHostObjectRequest) HierarchyTypeFilter(hierarchyTypeFilter EHierarchyType) ApiGetVmwareHostObjectRequest {
	r.hierarchyTypeFilter = &hierarchyTypeFilter
	return r
}
func (r ApiGetVmwareHostObjectRequest) NameFilter(nameFilter string) ApiGetVmwareHostObjectRequest {
	r.nameFilter = &nameFilter
	return r
}
func (r ApiGetVmwareHostObjectRequest) TypeFilter(typeFilter EVmwareInventoryType) ApiGetVmwareHostObjectRequest {
	r.typeFilter = &typeFilter
	return r
}
func (r ApiGetVmwareHostObjectRequest) ParentContainerNameFilter(parentContainerNameFilter string) ApiGetVmwareHostObjectRequest {
	r.parentContainerNameFilter = &parentContainerNameFilter
	return r
}

func (r ApiGetVmwareHostObjectRequest) Execute() (VCenterInventoryResult, *_nethttp.Response, error) {
	return r.ApiService.GetVmwareHostObjectExecute(r)
}

/*
 * GetVmwareHostObject Get VMware vSphere Server Objects
 * The HTTP GET request to the `/api/v1/inventory/vmware/hosts/{name}` path allows you to get an array of virtual infrastructure objects of the VMware vSphere server that has the specified `name`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name Name of the VMware vSphere server.
 * @return ApiGetVmwareHostObjectRequest
 */
func (a *InventoryBrowserApiService) GetVmwareHostObject(ctx _context.Context, name string) ApiGetVmwareHostObjectRequest {
	return ApiGetVmwareHostObjectRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

/*
 * Execute executes the request
 * @return VCenterInventoryResult
 */
func (a *InventoryBrowserApiService) GetVmwareHostObjectExecute(r ApiGetVmwareHostObjectRequest) (VCenterInventoryResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  VCenterInventoryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryBrowserApiService.GetVmwareHostObject")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/inventory/vmware/hosts/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", _neturl.PathEscape(parameterToString(r.name, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xApiVersion == nil {
		return localVarReturnValue, nil, reportError("xApiVersion is required and must be specified")
	}

	if r.skip != nil {
		localVarQueryParams.Add("skip", parameterToString(*r.skip, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.orderColumn != nil {
		localVarQueryParams.Add("orderColumn", parameterToString(*r.orderColumn, ""))
	}
	if r.orderAsc != nil {
		localVarQueryParams.Add("orderAsc", parameterToString(*r.orderAsc, ""))
	}
	if r.objectIdFilter != nil {
		localVarQueryParams.Add("objectIdFilter", parameterToString(*r.objectIdFilter, ""))
	}
	if r.hierarchyTypeFilter != nil {
		localVarQueryParams.Add("hierarchyTypeFilter", parameterToString(*r.hierarchyTypeFilter, ""))
	}
	if r.nameFilter != nil {
		localVarQueryParams.Add("nameFilter", parameterToString(*r.nameFilter, ""))
	}
	if r.typeFilter != nil {
		localVarQueryParams.Add("typeFilter", parameterToString(*r.typeFilter, ""))
	}
	if r.parentContainerNameFilter != nil {
		localVarQueryParams.Add("parentContainerNameFilter", parameterToString(*r.parentContainerNameFilter, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-api-version"] = parameterToString(*r.xApiVersion, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
