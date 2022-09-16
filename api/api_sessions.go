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


type SessionsApi interface {

	/*
	CreateSession Creates new Selenium session

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SessionsApiCreateSessionRequest
	*/
	CreateSession(ctx context.Context) SessionsApiCreateSessionRequest

	// CreateSessionExecute executes the request
	//  @return NewSessionResponse
	CreateSessionExecute(r SessionsApiCreateSessionRequest) (*NewSessionResponse, *http.Response, error)

	/*
	DeleteSession Deletes existing Selenium session

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sessionId Requested session ID
	@return SessionsApiDeleteSessionRequest
	*/
	DeleteSession(ctx context.Context, sessionId string) SessionsApiDeleteSessionRequest

	// DeleteSessionExecute executes the request
	//  @return EmptyResponse
	DeleteSessionExecute(r SessionsApiDeleteSessionRequest) (*EmptyResponse, *http.Response, error)

	/*
	GetStatus Gets Selenium API status information

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SessionsApiGetStatusRequest
	*/
	GetStatus(ctx context.Context) SessionsApiGetStatusRequest

	// GetStatusExecute executes the request
	//  @return StatusResponse
	GetStatusExecute(r SessionsApiGetStatusRequest) (*StatusResponse, *http.Response, error)
}

// SessionsApiService SessionsApi service
type SessionsApiService service

type SessionsApiCreateSessionRequest struct {
	ctx context.Context
	ApiService SessionsApi
	newSessionRequest *NewSessionRequest
}

func (r SessionsApiCreateSessionRequest) NewSessionRequest(newSessionRequest NewSessionRequest) SessionsApiCreateSessionRequest {
	r.newSessionRequest = &newSessionRequest
	return r
}

func (r SessionsApiCreateSessionRequest) Execute() (*NewSessionResponse, *http.Response, error) {
	return r.ApiService.CreateSessionExecute(r)
}

/*
CreateSession Creates new Selenium session

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SessionsApiCreateSessionRequest
*/
func (a *SessionsApiService) CreateSession(ctx context.Context) SessionsApiCreateSessionRequest {
	return SessionsApiCreateSessionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NewSessionResponse
func (a *SessionsApiService) CreateSessionExecute(r SessionsApiCreateSessionRequest) (*NewSessionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NewSessionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SessionsApiService.CreateSession")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/session"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.newSessionRequest == nil {
		return localVarReturnValue, nil, reportError("newSessionRequest is required and must be specified")
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
	localVarPostBody = r.newSessionRequest
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
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
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

type SessionsApiDeleteSessionRequest struct {
	ctx context.Context
	ApiService SessionsApi
	sessionId string
}

func (r SessionsApiDeleteSessionRequest) Execute() (*EmptyResponse, *http.Response, error) {
	return r.ApiService.DeleteSessionExecute(r)
}

/*
DeleteSession Deletes existing Selenium session

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sessionId Requested session ID
 @return SessionsApiDeleteSessionRequest
*/
func (a *SessionsApiService) DeleteSession(ctx context.Context, sessionId string) SessionsApiDeleteSessionRequest {
	return SessionsApiDeleteSessionRequest{
		ApiService: a,
		ctx: ctx,
		sessionId: sessionId,
	}
}

// Execute executes the request
//  @return EmptyResponse
func (a *SessionsApiService) DeleteSessionExecute(r SessionsApiDeleteSessionRequest) (*EmptyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmptyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SessionsApiService.DeleteSession")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/session/{sessionId}"
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

type SessionsApiGetStatusRequest struct {
	ctx context.Context
	ApiService SessionsApi
}

func (r SessionsApiGetStatusRequest) Execute() (*StatusResponse, *http.Response, error) {
	return r.ApiService.GetStatusExecute(r)
}

/*
GetStatus Gets Selenium API status information

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SessionsApiGetStatusRequest
*/
func (a *SessionsApiService) GetStatus(ctx context.Context) SessionsApiGetStatusRequest {
	return SessionsApiGetStatusRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StatusResponse
func (a *SessionsApiService) GetStatusExecute(r SessionsApiGetStatusRequest) (*StatusResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StatusResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SessionsApiService.GetStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/status"

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
