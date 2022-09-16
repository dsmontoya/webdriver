# \ElementsApi

All URIs are relative to *http://localhost:4444/wd/hub*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ElementClear**](ElementsApi.md#ElementClear) | **Post** /session/{sessionId}/element/{elementId}/clear | Clear element
[**ElementClick**](ElementsApi.md#ElementClick) | **Post** /session/{sessionId}/element/{elementId}/click | Click on element
[**ElementSendKeys**](ElementsApi.md#ElementSendKeys) | **Post** /session/{sessionId}/element/{elementId}/value | Send keys to element
[**FindElement**](ElementsApi.md#FindElement) | **Post** /session/{sessionId}/element | Find element
[**FindElementFromElement**](ElementsApi.md#FindElementFromElement) | **Post** /session/{sessionId}/element/{elementId}/element | Find element from element
[**FindElements**](ElementsApi.md#FindElements) | **Post** /session/{sessionId}/elements | Find elements
[**FindElementsFromElement**](ElementsApi.md#FindElementsFromElement) | **Post** /session/{sessionId}/element/{elementId}/elements | Find elements from element
[**GetActiveElement**](ElementsApi.md#GetActiveElement) | **Get** /session/{sessionId}/element/active | Get active element
[**GetElementAttribute**](ElementsApi.md#GetElementAttribute) | **Get** /session/{sessionId}/element/{elementId}/attribute/{name} | Get element attribute
[**GetElementCSSProperty**](ElementsApi.md#GetElementCSSProperty) | **Get** /session/{sessionId}/element/{elementId}/css/{propertyName} | Get element CSS property
[**GetElementComputedLabel**](ElementsApi.md#GetElementComputedLabel) | **Get** /session/{sessionId}/element/{elementId}/computedlabel | Get element computed accessibility label
[**GetElementComputedRole**](ElementsApi.md#GetElementComputedRole) | **Get** /session/{sessionId}/element/{elementId}/computedrole | Get element computed accessibility role
[**GetElementProperty**](ElementsApi.md#GetElementProperty) | **Get** /session/{sessionId}/element/{elementId}/property/{name} | Get element property
[**GetElementRect**](ElementsApi.md#GetElementRect) | **Get** /session/{sessionId}/element/{elementId}/rect | Get element rect
[**GetElementTagName**](ElementsApi.md#GetElementTagName) | **Get** /session/{sessionId}/element/{elementId}/name | Get element tag name
[**GetElementText**](ElementsApi.md#GetElementText) | **Get** /session/{sessionId}/element/{elementId}/text | Get element text
[**IsElementDisplayed**](ElementsApi.md#IsElementDisplayed) | **Get** /session/{sessionId}/element/{elementId}/displayed | Is element displayed
[**IsElementEnabled**](ElementsApi.md#IsElementEnabled) | **Get** /session/{sessionId}/element/{elementId}/enabled | Is element enabled
[**IsElementSelected**](ElementsApi.md#IsElementSelected) | **Get** /session/{sessionId}/element/{elementId}/selected | Is element selected



## ElementClear

> EmptyResponse ElementClear(ctx, sessionId, elementId).RequestBody(requestBody).Execute()

Clear element

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
    elementId := "elementId_example" // string | Requested element ID
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementsApi.ElementClear(context.Background(), sessionId, elementId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.ElementClear``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ElementClear`: EmptyResponse
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.ElementClear`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 
**elementId** | **string** | Requested element ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiElementClearRequest struct via the builder pattern


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


## ElementClick

> EmptyResponse ElementClick(ctx, sessionId, elementId).RequestBody(requestBody).Execute()

Click on element

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
    elementId := "elementId_example" // string | Requested element ID
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementsApi.ElementClick(context.Background(), sessionId, elementId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.ElementClick``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ElementClick`: EmptyResponse
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.ElementClick`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 
**elementId** | **string** | Requested element ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiElementClickRequest struct via the builder pattern


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


## ElementSendKeys

> EmptyResponse ElementSendKeys(ctx, sessionId, elementId).ElementSendKeysRequest(elementSendKeysRequest).Execute()

Send keys to element

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
    elementId := "elementId_example" // string | Requested element ID
    elementSendKeysRequest := *openapiclient.NewElementSendKeysRequest("Text_example") // ElementSendKeysRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementsApi.ElementSendKeys(context.Background(), sessionId, elementId).ElementSendKeysRequest(elementSendKeysRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.ElementSendKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ElementSendKeys`: EmptyResponse
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.ElementSendKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 
**elementId** | **string** | Requested element ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiElementSendKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **elementSendKeysRequest** | [**ElementSendKeysRequest**](ElementSendKeysRequest.md) |  | 

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


## FindElement

> FindElementResponse FindElement(ctx, sessionId).FindElementRequest(findElementRequest).Execute()

Find element

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
    findElementRequest := *openapiclient.NewFindElementRequest(openapiclient.LocatorStrategy("css selector"), "Value_example") // FindElementRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementsApi.FindElement(context.Background(), sessionId).FindElementRequest(findElementRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.FindElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindElement`: FindElementResponse
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.FindElement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **findElementRequest** | [**FindElementRequest**](FindElementRequest.md) |  | 

### Return type

[**FindElementResponse**](FindElementResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindElementFromElement

> FindElementResponse FindElementFromElement(ctx, sessionId, elementId).FindElementRequest(findElementRequest).Execute()

Find element from element

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
    elementId := "elementId_example" // string | Requested element ID
    findElementRequest := *openapiclient.NewFindElementRequest(openapiclient.LocatorStrategy("css selector"), "Value_example") // FindElementRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementsApi.FindElementFromElement(context.Background(), sessionId, elementId).FindElementRequest(findElementRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.FindElementFromElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindElementFromElement`: FindElementResponse
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.FindElementFromElement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 
**elementId** | **string** | Requested element ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindElementFromElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **findElementRequest** | [**FindElementRequest**](FindElementRequest.md) |  | 

### Return type

[**FindElementResponse**](FindElementResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindElements

> FindElementsResponse FindElements(ctx, sessionId).FindElementRequest(findElementRequest).Execute()

Find elements

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
    findElementRequest := *openapiclient.NewFindElementRequest(openapiclient.LocatorStrategy("css selector"), "Value_example") // FindElementRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementsApi.FindElements(context.Background(), sessionId).FindElementRequest(findElementRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.FindElements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindElements`: FindElementsResponse
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.FindElements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindElementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **findElementRequest** | [**FindElementRequest**](FindElementRequest.md) |  | 

### Return type

[**FindElementsResponse**](FindElementsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindElementsFromElement

> FindElementsResponse FindElementsFromElement(ctx, sessionId, elementId).FindElementRequest(findElementRequest).Execute()

Find elements from element

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
    elementId := "elementId_example" // string | Requested element ID
    findElementRequest := *openapiclient.NewFindElementRequest(openapiclient.LocatorStrategy("css selector"), "Value_example") // FindElementRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementsApi.FindElementsFromElement(context.Background(), sessionId, elementId).FindElementRequest(findElementRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.FindElementsFromElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindElementsFromElement`: FindElementsResponse
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.FindElementsFromElement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 
**elementId** | **string** | Requested element ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindElementsFromElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **findElementRequest** | [**FindElementRequest**](FindElementRequest.md) |  | 

### Return type

[**FindElementsResponse**](FindElementsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveElement

> FindElementResponse GetActiveElement(ctx, sessionId).Execute()

Get active element

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
    resp, r, err := apiClient.ElementsApi.GetActiveElement(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.GetActiveElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveElement`: FindElementResponse
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.GetActiveElement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FindElementResponse**](FindElementResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElementAttribute

> NullableStringResponse GetElementAttribute(ctx, sessionId, elementId, name).Execute()

Get element attribute

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
    elementId := "elementId_example" // string | Requested element ID
    name := "name_example" // string | Requested attribute name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementsApi.GetElementAttribute(context.Background(), sessionId, elementId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.GetElementAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetElementAttribute`: NullableStringResponse
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.GetElementAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 
**elementId** | **string** | Requested element ID | 
**name** | **string** | Requested attribute name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**NullableStringResponse**](NullableStringResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElementCSSProperty

> StringResponse GetElementCSSProperty(ctx, sessionId, elementId, propertyName).Execute()

Get element CSS property

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
    elementId := "elementId_example" // string | Requested element ID
    propertyName := "propertyName_example" // string | Requested CSS property name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementsApi.GetElementCSSProperty(context.Background(), sessionId, elementId, propertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.GetElementCSSProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetElementCSSProperty`: StringResponse
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.GetElementCSSProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 
**elementId** | **string** | Requested element ID | 
**propertyName** | **string** | Requested CSS property name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementCSSPropertyRequest struct via the builder pattern


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


## GetElementComputedLabel

> StringResponse GetElementComputedLabel(ctx, sessionId, elementId).Execute()

Get element computed accessibility label

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
    elementId := "elementId_example" // string | Requested element ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementsApi.GetElementComputedLabel(context.Background(), sessionId, elementId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.GetElementComputedLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetElementComputedLabel`: StringResponse
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.GetElementComputedLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 
**elementId** | **string** | Requested element ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementComputedLabelRequest struct via the builder pattern


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


## GetElementComputedRole

> StringResponse GetElementComputedRole(ctx, sessionId, elementId).Execute()

Get element computed accessibility role

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
    elementId := "elementId_example" // string | Requested element ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementsApi.GetElementComputedRole(context.Background(), sessionId, elementId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.GetElementComputedRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetElementComputedRole`: StringResponse
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.GetElementComputedRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 
**elementId** | **string** | Requested element ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementComputedRoleRequest struct via the builder pattern


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


## GetElementProperty

> NullableStringResponse GetElementProperty(ctx, sessionId, elementId, name).Execute()

Get element property

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
    elementId := "elementId_example" // string | Requested element ID
    name := "name_example" // string | Requested property name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementsApi.GetElementProperty(context.Background(), sessionId, elementId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.GetElementProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetElementProperty`: NullableStringResponse
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.GetElementProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 
**elementId** | **string** | Requested element ID | 
**name** | **string** | Requested property name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**NullableStringResponse**](NullableStringResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElementRect

> RectResponse GetElementRect(ctx, sessionId, elementId).Execute()

Get element rect

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
    elementId := "elementId_example" // string | Requested element ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementsApi.GetElementRect(context.Background(), sessionId, elementId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.GetElementRect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetElementRect`: RectResponse
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.GetElementRect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 
**elementId** | **string** | Requested element ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementRectRequest struct via the builder pattern


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


## GetElementTagName

> StringResponse GetElementTagName(ctx, sessionId, elementId).Execute()

Get element tag name

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
    elementId := "elementId_example" // string | Requested element ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementsApi.GetElementTagName(context.Background(), sessionId, elementId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.GetElementTagName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetElementTagName`: StringResponse
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.GetElementTagName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 
**elementId** | **string** | Requested element ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementTagNameRequest struct via the builder pattern


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


## GetElementText

> StringResponse GetElementText(ctx, sessionId, elementId).Execute()

Get element text

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
    elementId := "elementId_example" // string | Requested element ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementsApi.GetElementText(context.Background(), sessionId, elementId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.GetElementText``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetElementText`: StringResponse
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.GetElementText`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 
**elementId** | **string** | Requested element ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementTextRequest struct via the builder pattern


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


## IsElementDisplayed

> BooleanResponse IsElementDisplayed(ctx, sessionId, elementId).Execute()

Is element displayed



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
    elementId := "elementId_example" // string | Requested element ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementsApi.IsElementDisplayed(context.Background(), sessionId, elementId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.IsElementDisplayed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IsElementDisplayed`: BooleanResponse
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.IsElementDisplayed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 
**elementId** | **string** | Requested element ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiIsElementDisplayedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BooleanResponse**](BooleanResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsElementEnabled

> BooleanResponse IsElementEnabled(ctx, sessionId, elementId).Execute()

Is element enabled

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
    elementId := "elementId_example" // string | Requested element ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementsApi.IsElementEnabled(context.Background(), sessionId, elementId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.IsElementEnabled``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IsElementEnabled`: BooleanResponse
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.IsElementEnabled`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 
**elementId** | **string** | Requested element ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiIsElementEnabledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BooleanResponse**](BooleanResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsElementSelected

> BooleanResponse IsElementSelected(ctx, sessionId, elementId).Execute()

Is element selected

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
    elementId := "elementId_example" // string | Requested element ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElementsApi.IsElementSelected(context.Background(), sessionId, elementId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.IsElementSelected``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IsElementSelected`: BooleanResponse
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.IsElementSelected`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Requested session ID | 
**elementId** | **string** | Requested element ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiIsElementSelectedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BooleanResponse**](BooleanResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

