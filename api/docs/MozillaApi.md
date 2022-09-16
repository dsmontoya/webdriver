# \MozillaApi

All URIs are relative to *http://localhost:4444/wd/hub*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContext**](MozillaApi.md#GetContext) | **Get** /session/{sessionId}/moz/context | Get current context
[**InstallAddon**](MozillaApi.md#InstallAddon) | **Post** /session/{sessionId}/moz/addon/install | Install addon
[**SetContext**](MozillaApi.md#SetContext) | **Post** /session/{sessionId}/moz/context | Set current context
[**TakeFullScreenshot**](MozillaApi.md#TakeFullScreenshot) | **Get** /session/{sessionId}/moz/screenshot/full | Takes full page screenshot
[**UninstallAddon**](MozillaApi.md#UninstallAddon) | **Post** /session/{sessionId}/moz/addon/uninstall | Uninstall addon



## GetContext

> StringResponse GetContext(ctx, sessionId).Execute()

Get current context

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
    resp, r, err := apiClient.MozillaApi.GetContext(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MozillaApi.GetContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContext`: StringResponse
    fmt.Fprintf(os.Stdout, "Response from `MozillaApi.GetContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContextRequest struct via the builder pattern


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


## InstallAddon

> StringResponse InstallAddon(ctx, sessionId).AddonInstallRequest(addonInstallRequest).Execute()

Install addon

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
    addonInstallRequest := *openapiclient.NewAddonInstallRequest("Addon_example") // AddonInstallRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MozillaApi.InstallAddon(context.Background(), sessionId).AddonInstallRequest(addonInstallRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MozillaApi.InstallAddon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstallAddon`: StringResponse
    fmt.Fprintf(os.Stdout, "Response from `MozillaApi.InstallAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addonInstallRequest** | [**AddonInstallRequest**](AddonInstallRequest.md) |  | 

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


## SetContext

> EmptyResponse SetContext(ctx, sessionId).ContextRequest(contextRequest).Execute()

Set current context

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
    contextRequest := *openapiclient.NewContextRequest(openapiclient.FirefoxContext("content")) // ContextRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MozillaApi.SetContext(context.Background(), sessionId).ContextRequest(contextRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MozillaApi.SetContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetContext`: EmptyResponse
    fmt.Fprintf(os.Stdout, "Response from `MozillaApi.SetContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contextRequest** | [**ContextRequest**](ContextRequest.md) |  | 

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


## TakeFullScreenshot

> StringResponse TakeFullScreenshot(ctx, sessionId).Execute()

Takes full page screenshot

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
    resp, r, err := apiClient.MozillaApi.TakeFullScreenshot(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MozillaApi.TakeFullScreenshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TakeFullScreenshot`: StringResponse
    fmt.Fprintf(os.Stdout, "Response from `MozillaApi.TakeFullScreenshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTakeFullScreenshotRequest struct via the builder pattern


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


## UninstallAddon

> EmptyResponse UninstallAddon(ctx, sessionId).AddonUninstallRequest(addonUninstallRequest).Execute()

Uninstall addon

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
    addonUninstallRequest := *openapiclient.NewAddonUninstallRequest("Id_example") // AddonUninstallRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MozillaApi.UninstallAddon(context.Background(), sessionId).AddonUninstallRequest(addonUninstallRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MozillaApi.UninstallAddon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UninstallAddon`: EmptyResponse
    fmt.Fprintf(os.Stdout, "Response from `MozillaApi.UninstallAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUninstallAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addonUninstallRequest** | [**AddonUninstallRequest**](AddonUninstallRequest.md) |  | 

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

