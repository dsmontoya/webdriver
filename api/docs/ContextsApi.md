# \ContextsApi

All URIs are relative to *http://localhost:4444/wd/hub*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseWindow**](ContextsApi.md#CloseWindow) | **Delete** /session/{sessionId}/window | Closes current window
[**CreateNewWindow**](ContextsApi.md#CreateNewWindow) | **Post** /session/{sessionId}/window/new | Creates a new window or tab
[**FullscreenWindow**](ContextsApi.md#FullscreenWindow) | **Post** /session/{sessionId}/window/fullscreen | Fullscreen window
[**GetWindowHandle**](ContextsApi.md#GetWindowHandle) | **Get** /session/{sessionId}/window | Returns current window handle
[**GetWindowHandles**](ContextsApi.md#GetWindowHandles) | **Get** /session/{sessionId}/window/handles | Returns all window handles
[**GetWindowRect**](ContextsApi.md#GetWindowRect) | **Get** /session/{sessionId}/window/rect | Get window size
[**MaximizeWindow**](ContextsApi.md#MaximizeWindow) | **Post** /session/{sessionId}/window/maximize | Maximize window
[**MinimizeWindow**](ContextsApi.md#MinimizeWindow) | **Post** /session/{sessionId}/window/minimize | Minimize window
[**SetWindowRect**](ContextsApi.md#SetWindowRect) | **Post** /session/{sessionId}/window/rect | Set window rect
[**SwitchToFrame**](ContextsApi.md#SwitchToFrame) | **Post** /session/{sessionId}/frame | Switch to frame
[**SwitchToParentFrame**](ContextsApi.md#SwitchToParentFrame) | **Post** /session/{sessionId}/frame/parent | Switch to parent frame
[**SwitchToWindow**](ContextsApi.md#SwitchToWindow) | **Post** /session/{sessionId}/window | Switches to window



## CloseWindow

> GetWindowHandlesResponse CloseWindow(ctx, sessionId).Execute()

Closes current window

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
    resp, r, err := apiClient.ContextsApi.CloseWindow(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.CloseWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloseWindow`: GetWindowHandlesResponse
    fmt.Fprintf(os.Stdout, "Response from `ContextsApi.CloseWindow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseWindowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWindowHandlesResponse**](GetWindowHandlesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewWindow

> NewWindowResponse CreateNewWindow(ctx, sessionId).NewWindowRequest(newWindowRequest).Execute()

Creates a new window or tab

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
    newWindowRequest := *openapiclient.NewNewWindowRequest("Type_example") // NewWindowRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsApi.CreateNewWindow(context.Background(), sessionId).NewWindowRequest(newWindowRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.CreateNewWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewWindow`: NewWindowResponse
    fmt.Fprintf(os.Stdout, "Response from `ContextsApi.CreateNewWindow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewWindowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newWindowRequest** | [**NewWindowRequest**](NewWindowRequest.md) |  | 

### Return type

[**NewWindowResponse**](NewWindowResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FullscreenWindow

> Rect FullscreenWindow(ctx, sessionId).RequestBody(requestBody).Execute()

Fullscreen window

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
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsApi.FullscreenWindow(context.Background(), sessionId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.FullscreenWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FullscreenWindow`: Rect
    fmt.Fprintf(os.Stdout, "Response from `ContextsApi.FullscreenWindow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFullscreenWindowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]map[string]interface{}** |  | 

### Return type

[**Rect**](Rect.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWindowHandle

> StringResponse GetWindowHandle(ctx, sessionId).Execute()

Returns current window handle

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
    resp, r, err := apiClient.ContextsApi.GetWindowHandle(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.GetWindowHandle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWindowHandle`: StringResponse
    fmt.Fprintf(os.Stdout, "Response from `ContextsApi.GetWindowHandle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWindowHandleRequest struct via the builder pattern


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


## GetWindowHandles

> GetWindowHandlesResponse GetWindowHandles(ctx, sessionId).Execute()

Returns all window handles

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
    resp, r, err := apiClient.ContextsApi.GetWindowHandles(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.GetWindowHandles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWindowHandles`: GetWindowHandlesResponse
    fmt.Fprintf(os.Stdout, "Response from `ContextsApi.GetWindowHandles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWindowHandlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWindowHandlesResponse**](GetWindowHandlesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWindowRect

> RectResponse GetWindowRect(ctx, sessionId).Execute()

Get window size

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
    resp, r, err := apiClient.ContextsApi.GetWindowRect(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.GetWindowRect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWindowRect`: RectResponse
    fmt.Fprintf(os.Stdout, "Response from `ContextsApi.GetWindowRect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWindowRectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RectResponse**](RectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaximizeWindow

> Rect MaximizeWindow(ctx, sessionId).RequestBody(requestBody).Execute()

Maximize window

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
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsApi.MaximizeWindow(context.Background(), sessionId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.MaximizeWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MaximizeWindow`: Rect
    fmt.Fprintf(os.Stdout, "Response from `ContextsApi.MaximizeWindow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMaximizeWindowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]map[string]interface{}** |  | 

### Return type

[**Rect**](Rect.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MinimizeWindow

> Rect MinimizeWindow(ctx, sessionId).RequestBody(requestBody).Execute()

Minimize window

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
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsApi.MinimizeWindow(context.Background(), sessionId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.MinimizeWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MinimizeWindow`: Rect
    fmt.Fprintf(os.Stdout, "Response from `ContextsApi.MinimizeWindow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMinimizeWindowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]map[string]interface{}** |  | 

### Return type

[**Rect**](Rect.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetWindowRect

> Rect SetWindowRect(ctx, sessionId).Rect(rect).Execute()

Set window rect

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
    rect := *openapiclient.NewRect(float32(123), float32(123), float32(123), float32(123)) // Rect | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsApi.SetWindowRect(context.Background(), sessionId).Rect(rect).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.SetWindowRect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetWindowRect`: Rect
    fmt.Fprintf(os.Stdout, "Response from `ContextsApi.SetWindowRect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetWindowRectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rect** | [**Rect**](Rect.md) |  | 

### Return type

[**Rect**](Rect.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwitchToFrame

> EmptyResponse SwitchToFrame(ctx, sessionId).SwitchToFrameRequest(switchToFrameRequest).Execute()

Switch to frame

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
    switchToFrameRequest := *openapiclient.NewSwitchToFrameRequest() // SwitchToFrameRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsApi.SwitchToFrame(context.Background(), sessionId).SwitchToFrameRequest(switchToFrameRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.SwitchToFrame``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SwitchToFrame`: EmptyResponse
    fmt.Fprintf(os.Stdout, "Response from `ContextsApi.SwitchToFrame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwitchToFrameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **switchToFrameRequest** | [**SwitchToFrameRequest**](SwitchToFrameRequest.md) |  | 

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


## SwitchToParentFrame

> EmptyResponse SwitchToParentFrame(ctx, sessionId).RequestBody(requestBody).Execute()

Switch to parent frame

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
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsApi.SwitchToParentFrame(context.Background(), sessionId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.SwitchToParentFrame``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SwitchToParentFrame`: EmptyResponse
    fmt.Fprintf(os.Stdout, "Response from `ContextsApi.SwitchToParentFrame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwitchToParentFrameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]map[string]interface{}** |  | 

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


## SwitchToWindow

> EmptyResponse SwitchToWindow(ctx, sessionId).SwitchToWindowRequest(switchToWindowRequest).Execute()

Switches to window

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
    switchToWindowRequest := *openapiclient.NewSwitchToWindowRequest("Handle_example") // SwitchToWindowRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsApi.SwitchToWindow(context.Background(), sessionId).SwitchToWindowRequest(switchToWindowRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.SwitchToWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SwitchToWindow`: EmptyResponse
    fmt.Fprintf(os.Stdout, "Response from `ContextsApi.SwitchToWindow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwitchToWindowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **switchToWindowRequest** | [**SwitchToWindowRequest**](SwitchToWindowRequest.md) |  | 

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

