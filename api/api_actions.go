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


type ActionsApi interface {

	/*
	PerformActions Perform actions

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sessionId Requested session ID
	@return ActionsApiPerformActionsRequest
	*/
	PerformActions(ctx context.Context, sessionId string) ActionsApiPerformActionsRequest

	// PerformActionsExecute executes the request
	//  @return EmptyResponse
	PerformActionsExecute(r ActionsApiPerformActionsRequest) (*EmptyResponse, *http.Response, error)

	/*
	ReleaseActions Release actions

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sessionId Requested session ID
	@return ActionsApiReleaseActionsRequest
	*/
	ReleaseActions(ctx context.Context, sessionId string) ActionsApiReleaseActionsRequest

	// ReleaseActionsExecute executes the request
	//  @return EmptyResponse
	ReleaseActionsExecute(r ActionsApiReleaseActionsRequest) (*EmptyResponse, *http.Response, error)
}

// ActionsApiService ActionsApi service
type ActionsApiService service

type ActionsApiPerformActionsRequest struct {
	ctx context.Context
	ApiService ActionsApi
	sessionId string
	actionsRequest *ActionsRequest
}

func (r ActionsApiPerformActionsRequest) ActionsRequest(actionsRequest ActionsRequest) ActionsApiPerformActionsRequest {
	r.actionsRequest = &actionsRequest
	return r
}

func (r ActionsApiPerformActionsRequest) Execute() (*EmptyResponse, *http.Response, error) {
	return r.ApiService.PerformActionsExecute(r)
}

/*
PerformActions Perform actions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sessionId Requested session ID
 @return ActionsApiPerformActionsRequest
*/
func (a *ActionsApiService) PerformActions(ctx context.Context, sessionId string) ActionsApiPerformActionsRequest {
	return ActionsApiPerformActionsRequest{
		ApiService: a,
		ctx: ctx,
		sessionId: sessionId,
	}
}

// Execute executes the request
//  @return EmptyResponse
func (a *ActionsApiService) PerformActionsExecute(r ActionsApiPerformActionsRequest) (*EmptyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmptyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActionsApiService.PerformActions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/session/{sessionId}/actions"
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", url.PathEscape(parameterToString(r.sessionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.actionsRequest == nil {
		return localVarReturnValue, nil, reportError("actionsRequest is required and must be specified")
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
	localVarPostBody = r.actionsRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
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

type ActionsApiReleaseActionsRequest struct {
	ctx context.Context
	ApiService ActionsApi
	sessionId string
}

func (r ActionsApiReleaseActionsRequest) Execute() (*EmptyResponse, *http.Response, error) {
	return r.ApiService.ReleaseActionsExecute(r)
}

/*
ReleaseActions Release actions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sessionId Requested session ID
 @return ActionsApiReleaseActionsRequest
*/
func (a *ActionsApiService) ReleaseActions(ctx context.Context, sessionId string) ActionsApiReleaseActionsRequest {
	return ActionsApiReleaseActionsRequest{
		ApiService: a,
		ctx: ctx,
		sessionId: sessionId,
	}
}

// Execute executes the request
//  @return EmptyResponse
func (a *ActionsApiService) ReleaseActionsExecute(r ActionsApiReleaseActionsRequest) (*EmptyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmptyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActionsApiService.ReleaseActions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/session/{sessionId}/actions"
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
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