# \DocumentApi

All URIs are relative to *http://localhost:4444/wd/hub*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecuteScript**](DocumentApi.md#ExecuteScript) | **Post** /session/{sessionId}/execute/sync | Execute script
[**ExecuteScriptAsync**](DocumentApi.md#ExecuteScriptAsync) | **Post** /session/{sessionId}/execute/async | Execute script asynchronously
[**GetPageSource**](DocumentApi.md#GetPageSource) | **Get** /session/{sessionId}/source | Get page source
[**UploadFile**](DocumentApi.md#UploadFile) | **Post** /session/{sessionId}/file | Upload file



## ExecuteScript

> AnyResponse ExecuteScript(ctx, sessionId).ScriptRequest(scriptRequest).Execute()

Execute script

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sessionId := "sessionId_example" // string | Requested session ID
    scriptRequest := *openapiclient.NewScriptRequest() // ScriptRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.ExecuteScript(context.Background(), sessionId).ScriptRequest(scriptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.ExecuteScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteScript`: AnyResponse
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.ExecuteScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scriptRequest** | [**ScriptRequest**](ScriptRequest.md) |  | 

### Return type

[**AnyResponse**](AnyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteScriptAsync

> AnyResponse ExecuteScriptAsync(ctx, sessionId).ScriptRequest(scriptRequest).Execute()

Execute script asynchronously

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sessionId := "sessionId_example" // string | Requested session ID
    scriptRequest := *openapiclient.NewScriptRequest() // ScriptRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.ExecuteScriptAsync(context.Background(), sessionId).ScriptRequest(scriptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.ExecuteScriptAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteScriptAsync`: AnyResponse
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.ExecuteScriptAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteScriptAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scriptRequest** | [**ScriptRequest**](ScriptRequest.md) |  | 

### Return type

[**AnyResponse**](AnyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageSource

> StringResponse GetPageSource(ctx, sessionId).Execute()

Get page source

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sessionId := "sessionId_example" // string | Requested session ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.GetPageSource(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.GetPageSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPageSource`: StringResponse
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.GetPageSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StringResponse**](StringResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFile

> FileUploadResponse UploadFile(ctx, sessionId).FileUploadRequest(fileUploadRequest).Execute()

Upload file

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sessionId := "sessionId_example" // string | Requested session ID
    fileUploadRequest := *openapiclient.NewFileUploadRequest() // FileUploadRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.UploadFile(context.Background(), sessionId).FileUploadRequest(fileUploadRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.UploadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadFile`: FileUploadResponse
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.UploadFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileUploadRequest** | [**FileUploadRequest**](FileUploadRequest.md) |  | 

### Return type

[**FileUploadResponse**](FileUploadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

