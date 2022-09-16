# \ScreenshotsApi

All URIs are relative to *http://localhost:4444/wd/hub*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TakeElementScreenshot**](ScreenshotsApi.md#TakeElementScreenshot) | **Get** /session/{sessionId}/element/{elementId}/screenshot | Takes element screenshot
[**TakeScreenshot**](ScreenshotsApi.md#TakeScreenshot) | **Get** /session/{sessionId}/screenshot | Takes page screenshot



## TakeElementScreenshot

> StringResponse TakeElementScreenshot(ctx, sessionId, elementId).Execute()

Takes element screenshot

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
    elementId := "elementId_example" // string | Selenium element ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScreenshotsApi.TakeElementScreenshot(context.Background(), sessionId, elementId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScreenshotsApi.TakeElementScreenshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TakeElementScreenshot`: StringResponse
    fmt.Fprintf(os.Stdout, "Response from `ScreenshotsApi.TakeElementScreenshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 
**elementId** | **string** | Selenium element ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTakeElementScreenshotRequest struct via the builder pattern


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


## TakeScreenshot

> StringResponse TakeScreenshot(ctx, sessionId).Execute()

Takes page screenshot

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
    resp, r, err := apiClient.ScreenshotsApi.TakeScreenshot(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScreenshotsApi.TakeScreenshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TakeScreenshot`: StringResponse
    fmt.Fprintf(os.Stdout, "Response from `ScreenshotsApi.TakeScreenshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTakeScreenshotRequest struct via the builder pattern


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

