# \AerokubeApi

All URIs are relative to *http://localhost:4444/wd/hub*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRemoteFile**](AerokubeApi.md#DeleteRemoteFile) | **Delete** /session/{sessionId}/aerokube/download/{fileName} | Deletes file in remote container with browser
[**DownloadRemoteFile**](AerokubeApi.md#DownloadRemoteFile) | **Get** /session/{sessionId}/aerokube/download/{fileName} | Downloads file from remote container with browser
[**GetClipboard**](AerokubeApi.md#GetClipboard) | **Get** /session/{sessionId}/aerokube/clipboard | Returns clipboard contents
[**ListRemoteFiles**](AerokubeApi.md#ListRemoteFiles) | **Get** /session/{sessionId}/aerokube/download | Lists files in remote container with browser
[**UpdateClipboard**](AerokubeApi.md#UpdateClipboard) | **Post** /session/{sessionId}/aerokube/clipboard | Updates clipboard contents



## DeleteRemoteFile

> DeleteRemoteFile(ctx, sessionId, fileName).Execute()

Deletes file in remote container with browser

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
    fileName := "fileName_example" // string | Requested file name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AerokubeApi.DeleteRemoteFile(context.Background(), sessionId, fileName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AerokubeApi.DeleteRemoteFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 
**fileName** | **string** | Requested file name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRemoteFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadRemoteFile

> *os.File DownloadRemoteFile(ctx, sessionId, fileName).Execute()

Downloads file from remote container with browser

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
    fileName := "fileName_example" // string | Requested file name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AerokubeApi.DownloadRemoteFile(context.Background(), sessionId, fileName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AerokubeApi.DownloadRemoteFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadRemoteFile`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AerokubeApi.DownloadRemoteFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 
**fileName** | **string** | Requested file name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadRemoteFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClipboard

> ClipboardData GetClipboard(ctx, sessionId).Execute()

Returns clipboard contents

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
    resp, r, err := apiClient.AerokubeApi.GetClipboard(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AerokubeApi.GetClipboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClipboard`: ClipboardData
    fmt.Fprintf(os.Stdout, "Response from `AerokubeApi.GetClipboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClipboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClipboardData**](ClipboardData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRemoteFiles

> ArrayResponse ListRemoteFiles(ctx, sessionId).Json(json).Execute()

Lists files in remote container with browser

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
    json := true // bool | Information is returned in JSON format

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AerokubeApi.ListRemoteFiles(context.Background(), sessionId).Json(json).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AerokubeApi.ListRemoteFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRemoteFiles`: ArrayResponse
    fmt.Fprintf(os.Stdout, "Response from `AerokubeApi.ListRemoteFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRemoteFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **json** | **bool** | Information is returned in JSON format | 

### Return type

[**ArrayResponse**](ArrayResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClipboard

> UpdateClipboard(ctx, sessionId).ClipboardData(clipboardData).Execute()

Updates clipboard contents

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
    clipboardData := *openapiclient.NewClipboardData("Value_example") // ClipboardData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AerokubeApi.UpdateClipboard(context.Background(), sessionId).ClipboardData(clipboardData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AerokubeApi.UpdateClipboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClipboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clipboardData** | [**ClipboardData**](ClipboardData.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

