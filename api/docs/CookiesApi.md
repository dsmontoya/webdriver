# \CookiesApi

All URIs are relative to *http://localhost:4444/wd/hub*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCookie**](CookiesApi.md#AddCookie) | **Post** /session/{sessionId}/cookie | Add cookie
[**DeleteAllCookies**](CookiesApi.md#DeleteAllCookies) | **Delete** /session/{sessionId}/cookie | Delete all cookies
[**DeleteCookie**](CookiesApi.md#DeleteCookie) | **Delete** /session/{sessionId}/cookie/{name} | Delete cookie
[**GetAllCookies**](CookiesApi.md#GetAllCookies) | **Get** /session/{sessionId}/cookie | Get all cookies
[**GetNamedCookie**](CookiesApi.md#GetNamedCookie) | **Get** /session/{sessionId}/cookie/{name} | Get named cookie



## AddCookie

> EmptyResponse AddCookie(ctx, sessionId).CookieRequest(cookieRequest).Execute()

Add cookie

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
    cookieRequest := *openapiclient.NewCookieRequest(*openapiclient.NewCookie("Name_example", "Value_example", "Path_example", "Domain_example", false, false, "SameSite_example")) // CookieRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CookiesApi.AddCookie(context.Background(), sessionId).CookieRequest(cookieRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CookiesApi.AddCookie``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCookie`: EmptyResponse
    fmt.Fprintf(os.Stdout, "Response from `CookiesApi.AddCookie`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCookieRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cookieRequest** | [**CookieRequest**](CookieRequest.md) |  | 

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


## DeleteAllCookies

> EmptyResponse DeleteAllCookies(ctx, sessionId).Execute()

Delete all cookies

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
    resp, r, err := apiClient.CookiesApi.DeleteAllCookies(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CookiesApi.DeleteAllCookies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAllCookies`: EmptyResponse
    fmt.Fprintf(os.Stdout, "Response from `CookiesApi.DeleteAllCookies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllCookiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmptyResponse**](EmptyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCookie

> EmptyResponse DeleteCookie(ctx, sessionId, name).Execute()

Delete cookie

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
    name := "name_example" // string | Cookie name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CookiesApi.DeleteCookie(context.Background(), sessionId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CookiesApi.DeleteCookie``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCookie`: EmptyResponse
    fmt.Fprintf(os.Stdout, "Response from `CookiesApi.DeleteCookie`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 
**name** | **string** | Cookie name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCookieRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EmptyResponse**](EmptyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllCookies

> CookiesResponse GetAllCookies(ctx, sessionId).Execute()

Get all cookies

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
    resp, r, err := apiClient.CookiesApi.GetAllCookies(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CookiesApi.GetAllCookies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllCookies`: CookiesResponse
    fmt.Fprintf(os.Stdout, "Response from `CookiesApi.GetAllCookies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCookiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CookiesResponse**](CookiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamedCookie

> CookieResponse GetNamedCookie(ctx, sessionId, name).Execute()

Get named cookie

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
    name := "name_example" // string | Cookie name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CookiesApi.GetNamedCookie(context.Background(), sessionId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CookiesApi.GetNamedCookie``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamedCookie`: CookieResponse
    fmt.Fprintf(os.Stdout, "Response from `CookiesApi.GetNamedCookie`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 
**name** | **string** | Cookie name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamedCookieRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CookieResponse**](CookieResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

