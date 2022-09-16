# \PrintApi

All URIs are relative to *http://localhost:4444/wd/hub*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrintPage**](PrintApi.md#PrintPage) | **Post** /session/{sessionId}/print | Print page to PDF



## PrintPage

> StringResponse PrintPage(ctx, sessionId).PrintRequest(printRequest).Execute()

Print page to PDF

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
    printRequest := *openapiclient.NewPrintRequest(*openapiclient.NewPrintRequestOptions("Orientation_example", float32(123), false, *openapiclient.NewPrintRequestOptionsPage(), *openapiclient.NewPrintRequestOptionsMargin(), false, []string{"PageRanges_example"})) // PrintRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrintApi.PrintPage(context.Background(), sessionId).PrintRequest(printRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintApi.PrintPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintPage`: StringResponse
    fmt.Fprintf(os.Stdout, "Response from `PrintApi.PrintPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **printRequest** | [**PrintRequest**](PrintRequest.md) |  | 

### Return type

[**StringResponse**](StringResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

