# \TimeoutsApi

All URIs are relative to *http://localhost:4444/wd/hub*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTimeouts**](TimeoutsApi.md#GetTimeouts) | **Get** /session/{sessionId}/timeouts | Get session timeouts
[**SetTimeouts**](TimeoutsApi.md#SetTimeouts) | **Post** /session/{sessionId}/timeouts | Adjusts session timeouts



## GetTimeouts

> TimeoutsResponse GetTimeouts(ctx, sessionId).Execute()

Get session timeouts

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
    resp, r, err := apiClient.TimeoutsApi.GetTimeouts(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeoutsApi.GetTimeouts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTimeouts`: TimeoutsResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeoutsApi.GetTimeouts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeoutsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TimeoutsResponse**](TimeoutsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTimeouts

> EmptyResponse SetTimeouts(ctx, sessionId).Timeouts(timeouts).Execute()

Adjusts session timeouts

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
    timeouts := *openapiclient.NewTimeouts() // Timeouts | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TimeoutsApi.SetTimeouts(context.Background(), sessionId).Timeouts(timeouts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeoutsApi.SetTimeouts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetTimeouts`: EmptyResponse
    fmt.Fprintf(os.Stdout, "Response from `TimeoutsApi.SetTimeouts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetTimeoutsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeouts** | [**Timeouts**](Timeouts.md) |  | 

### Return type

[**EmptyResponse**](EmptyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

