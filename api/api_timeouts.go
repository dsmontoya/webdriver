/*
Selenium WebDriver

[Selenium WebDriver](https://www.w3.org/TR/webdriver) API specification

API version: 1.0.0
Contact: support@aerokube.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type TimeoutsApi interface {

	/*
	GetTimeouts Get session timeouts

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sessionId Requested session ID
	@return TimeoutsApiGetTimeoutsRequest
	*/
	GetTimeouts(ctx context.Context, sessionId string) TimeoutsApiGetTimeoutsRequest

	// GetTimeoutsExecute executes the request
	//  @return TimeoutsResponse
	GetTimeoutsExecute(r TimeoutsApiGetTimeoutsRequest) (*TimeoutsResponse, *http.Response, error)

	/*
	SetTimeouts Adjusts session timeouts

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sessionId Requested session ID
	@return TimeoutsApiSetTimeoutsRequest
	*/
	SetTimeouts(ctx context.Context, sessionId string) TimeoutsApiSetTimeoutsRequest

	// SetTimeoutsExecute executes the request
	//  @return EmptyResponse
	SetTimeoutsExecute(r TimeoutsApiSetTimeoutsRequest) (*EmptyResponse, *http.Response, error)
}

// TimeoutsApiService TimeoutsApi service
type TimeoutsApiService service

type TimeoutsApiGetTimeoutsRequest struct {
	ctx context.Context
	ApiService TimeoutsApi
	sessionId string
}

func (r TimeoutsApiGetTimeoutsRequest) Execute() (*TimeoutsResponse, *http.Response, error) {
	return r.ApiService.GetTimeoutsExecute(r)
}

/*
GetTimeouts Get session timeouts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sessionId Requested session ID
 @return TimeoutsApiGetTimeoutsRequest
*/
func (a *TimeoutsApiService) GetTimeouts(ctx context.Context, sessionId string) TimeoutsApiGetTimeoutsRequest {
	return TimeoutsApiGetTimeoutsRequest{
		ApiService: a,
		ctx: ctx,
		sessionId: sessionId,
	}
}

// Execute executes the request
//  @return TimeoutsResponse
func (a *TimeoutsApiService) GetTimeoutsExecute(r TimeoutsApiGetTimeoutsRequest) (*TimeoutsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TimeoutsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TimeoutsApiService.GetTimeouts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/session/{sessionId}/timeouts"
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", url.PathEscape(parameterToString(r.sessionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TimeoutsApiSetTimeoutsRequest struct {
	ctx context.Context
	ApiService TimeoutsApi
	sessionId string
	timeouts *Timeouts
}

func (r TimeoutsApiSetTimeoutsRequest) Timeouts(timeouts Timeouts) TimeoutsApiSetTimeoutsRequest {
	r.timeouts = &timeouts
	return r
}

func (r TimeoutsApiSetTimeoutsRequest) Execute() (*EmptyResponse, *http.Response, error) {
	return r.ApiService.SetTimeoutsExecute(r)
}

/*
SetTimeouts Adjusts session timeouts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sessionId Requested session ID
 @return TimeoutsApiSetTimeoutsRequest
*/
func (a *TimeoutsApiService) SetTimeouts(ctx context.Context, sessionId string) TimeoutsApiSetTimeoutsRequest {
	return TimeoutsApiSetTimeoutsRequest{
		ApiService: a,
		ctx: ctx,
		sessionId: sessionId,
	}
}

// Execute executes the request
//  @return EmptyResponse
func (a *TimeoutsApiService) SetTimeoutsExecute(r TimeoutsApiSetTimeoutsRequest) (*EmptyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmptyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TimeoutsApiService.SetTimeouts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/session/{sessionId}/timeouts"
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", url.PathEscape(parameterToString(r.sessionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.timeouts == nil {
		return localVarReturnValue, nil, reportError("timeouts is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.timeouts
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
