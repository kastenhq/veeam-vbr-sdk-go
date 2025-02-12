/*
Veeam Backup & Replication REST API

This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br>Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Some request and response bodies refer to reusable schema objects that are defined in the `schemas` section of the REST API specification. Schemas, in turn, may inherit a part of their properties by referring to other schemas, and be polymorphic by referring to multiple alternate schemas.

API version: 1.1-rev0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)


// BackupsApiService BackupsApi service
type BackupsApiService service

type ApiGetAllBackupsRequest struct {
	ctx context.Context
	ApiService *BackupsApiService
	xApiVersion *string
	skip *int32
	limit *int32
	orderColumn *EBackupsFiltersOrderColumn
	orderAsc *bool
	nameFilter *string
	createdAfterFilter *time.Time
	createdBeforeFilter *time.Time
	platformIdFilter *string
	jobIdFilter *string
	policyTagFilter *string
}

// Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;.
func (r ApiGetAllBackupsRequest) XApiVersion(xApiVersion string) ApiGetAllBackupsRequest {
	r.xApiVersion = &xApiVersion
	return r
}

// Number of backups to skip.
func (r ApiGetAllBackupsRequest) Skip(skip int32) ApiGetAllBackupsRequest {
	r.skip = &skip
	return r
}

// Maximum number of backups to return.
func (r ApiGetAllBackupsRequest) Limit(limit int32) ApiGetAllBackupsRequest {
	r.limit = &limit
	return r
}

// Sorts backups by one of the backup parameters.
func (r ApiGetAllBackupsRequest) OrderColumn(orderColumn EBackupsFiltersOrderColumn) ApiGetAllBackupsRequest {
	r.orderColumn = &orderColumn
	return r
}

// Sorts backups in the ascending order by the &#x60;orderColumn&#x60; parameter.
func (r ApiGetAllBackupsRequest) OrderAsc(orderAsc bool) ApiGetAllBackupsRequest {
	r.orderAsc = &orderAsc
	return r
}

// Filters backups by the &#x60;nameFilter&#x60; pattern. The pattern can match any backup parameter. To substitute one or more characters, use the asterisk (*) character at the beginning, at the end or both.
func (r ApiGetAllBackupsRequest) NameFilter(nameFilter string) ApiGetAllBackupsRequest {
	r.nameFilter = &nameFilter
	return r
}

// Returns backups that are created after the specified date and time.
func (r ApiGetAllBackupsRequest) CreatedAfterFilter(createdAfterFilter time.Time) ApiGetAllBackupsRequest {
	r.createdAfterFilter = &createdAfterFilter
	return r
}

// Returns backups that are created before the specified date and time.
func (r ApiGetAllBackupsRequest) CreatedBeforeFilter(createdBeforeFilter time.Time) ApiGetAllBackupsRequest {
	r.createdBeforeFilter = &createdBeforeFilter
	return r
}

// Filters backups by ID of the backup object platform.
func (r ApiGetAllBackupsRequest) PlatformIdFilter(platformIdFilter string) ApiGetAllBackupsRequest {
	r.platformIdFilter = &platformIdFilter
	return r
}

// Filters backups by ID of the parent job.
func (r ApiGetAllBackupsRequest) JobIdFilter(jobIdFilter string) ApiGetAllBackupsRequest {
	r.jobIdFilter = &jobIdFilter
	return r
}

// Filters backups by retention policy tag.
func (r ApiGetAllBackupsRequest) PolicyTagFilter(policyTagFilter string) ApiGetAllBackupsRequest {
	r.policyTagFilter = &policyTagFilter
	return r
}

func (r ApiGetAllBackupsRequest) Execute() (*BackupsResult, *http.Response, error) {
	return r.ApiService.GetAllBackupsExecute(r)
}

/*
GetAllBackups Get All Backups

The HTTP GET request to the `/api/v1/backups` path allows you to get an array of all backups that are created on or imported to the backup server.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllBackupsRequest
*/
func (a *BackupsApiService) GetAllBackups(ctx context.Context) ApiGetAllBackupsRequest {
	return ApiGetAllBackupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BackupsResult
func (a *BackupsApiService) GetAllBackupsExecute(r ApiGetAllBackupsRequest) (*BackupsResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BackupsResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupsApiService.GetAllBackups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/backups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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
	if r.createdAfterFilter != nil {
		localVarQueryParams.Add("createdAfterFilter", parameterToString(*r.createdAfterFilter, ""))
	}
	if r.createdBeforeFilter != nil {
		localVarQueryParams.Add("createdBeforeFilter", parameterToString(*r.createdBeforeFilter, ""))
	}
	if r.platformIdFilter != nil {
		localVarQueryParams.Add("platformIdFilter", parameterToString(*r.platformIdFilter, ""))
	}
	if r.jobIdFilter != nil {
		localVarQueryParams.Add("jobIdFilter", parameterToString(*r.jobIdFilter, ""))
	}
	if r.policyTagFilter != nil {
		localVarQueryParams.Add("policyTagFilter", parameterToString(*r.policyTagFilter, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
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
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetBackupRequest struct {
	ctx context.Context
	ApiService *BackupsApiService
	xApiVersion *string
	id string
}

// Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;.
func (r ApiGetBackupRequest) XApiVersion(xApiVersion string) ApiGetBackupRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiGetBackupRequest) Execute() (*BackupModel, *http.Response, error) {
	return r.ApiService.GetBackupExecute(r)
}

/*
GetBackup Get Backup

The HTTP GET request to the `/api/v1/backups/{id}` path allows you to get a backup that has the specified `id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ID of the backup.
 @return ApiGetBackupRequest
*/
func (a *BackupsApiService) GetBackup(ctx context.Context, id string) ApiGetBackupRequest {
	return ApiGetBackupRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BackupModel
func (a *BackupsApiService) GetBackupExecute(r ApiGetBackupRequest) (*BackupModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BackupModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupsApiService.GetBackup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/backups/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xApiVersion == nil {
		return localVarReturnValue, nil, reportError("xApiVersion is required and must be specified")
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
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
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetBackupObjectsRequest struct {
	ctx context.Context
	ApiService *BackupsApiService
	xApiVersion *string
	id string
}

// Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;.
func (r ApiGetBackupObjectsRequest) XApiVersion(xApiVersion string) ApiGetBackupObjectsRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiGetBackupObjectsRequest) Execute() (*BackupObjectsResult, *http.Response, error) {
	return r.ApiService.GetBackupObjectsExecute(r)
}

/*
GetBackupObjects Get Backup Objects

The HTTP GET request to the `/api/v1/backups/{id}/objects` path allows you to get an array of virtual infrastructure objects (VMs and VM containers) that are included in a backup that has the specified ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ID of the backup.
 @return ApiGetBackupObjectsRequest
*/
func (a *BackupsApiService) GetBackupObjects(ctx context.Context, id string) ApiGetBackupObjectsRequest {
	return ApiGetBackupObjectsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BackupObjectsResult
func (a *BackupsApiService) GetBackupObjectsExecute(r ApiGetBackupObjectsRequest) (*BackupObjectsResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BackupObjectsResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupsApiService.GetBackupObjects")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/backups/{id}/objects"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xApiVersion == nil {
		return localVarReturnValue, nil, reportError("xApiVersion is required and must be specified")
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
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
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
