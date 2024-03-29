# Go API client for api

[Selenium WebDriver](https://www.w3.org/TR/webdriver) API specification

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: com.github.dsmontoya.webdriver.GoCustomGenerator

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import api "github.com/dsmontoya/webdriver"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), api.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), api.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), api.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), api.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:4444/wd/hub*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ActionsApi* | [**PerformActions**](docs/ActionsApi.md#performactions) | **Post** /session/{sessionId}/actions | Perform actions
*ActionsApi* | [**ReleaseActions**](docs/ActionsApi.md#releaseactions) | **Delete** /session/{sessionId}/actions | Release actions
*AerokubeApi* | [**DeleteRemoteFile**](docs/AerokubeApi.md#deleteremotefile) | **Delete** /session/{sessionId}/aerokube/download/{fileName} | Deletes file in remote container with browser
*AerokubeApi* | [**DownloadRemoteFile**](docs/AerokubeApi.md#downloadremotefile) | **Get** /session/{sessionId}/aerokube/download/{fileName} | Downloads file from remote container with browser
*AerokubeApi* | [**GetClipboard**](docs/AerokubeApi.md#getclipboard) | **Get** /session/{sessionId}/aerokube/clipboard | Returns clipboard contents
*AerokubeApi* | [**ListRemoteFiles**](docs/AerokubeApi.md#listremotefiles) | **Get** /session/{sessionId}/aerokube/download | Lists files in remote container with browser
*AerokubeApi* | [**UpdateClipboard**](docs/AerokubeApi.md#updateclipboard) | **Post** /session/{sessionId}/aerokube/clipboard | Updates clipboard contents
*ContextsApi* | [**CloseWindow**](docs/ContextsApi.md#closewindow) | **Delete** /session/{sessionId}/window | Closes current window
*ContextsApi* | [**CreateNewWindow**](docs/ContextsApi.md#createnewwindow) | **Post** /session/{sessionId}/window/new | Creates a new window or tab
*ContextsApi* | [**FullscreenWindow**](docs/ContextsApi.md#fullscreenwindow) | **Post** /session/{sessionId}/window/fullscreen | Fullscreen window
*ContextsApi* | [**GetWindowHandle**](docs/ContextsApi.md#getwindowhandle) | **Get** /session/{sessionId}/window | Returns current window handle
*ContextsApi* | [**GetWindowHandles**](docs/ContextsApi.md#getwindowhandles) | **Get** /session/{sessionId}/window/handles | Returns all window handles
*ContextsApi* | [**GetWindowRect**](docs/ContextsApi.md#getwindowrect) | **Get** /session/{sessionId}/window/rect | Get window size
*ContextsApi* | [**MaximizeWindow**](docs/ContextsApi.md#maximizewindow) | **Post** /session/{sessionId}/window/maximize | Maximize window
*ContextsApi* | [**MinimizeWindow**](docs/ContextsApi.md#minimizewindow) | **Post** /session/{sessionId}/window/minimize | Minimize window
*ContextsApi* | [**SetWindowRect**](docs/ContextsApi.md#setwindowrect) | **Post** /session/{sessionId}/window/rect | Set window rect
*ContextsApi* | [**SwitchToFrame**](docs/ContextsApi.md#switchtoframe) | **Post** /session/{sessionId}/frame | Switch to frame
*ContextsApi* | [**SwitchToParentFrame**](docs/ContextsApi.md#switchtoparentframe) | **Post** /session/{sessionId}/frame/parent | Switch to parent frame
*ContextsApi* | [**SwitchToWindow**](docs/ContextsApi.md#switchtowindow) | **Post** /session/{sessionId}/window | Switches to window
*CookiesApi* | [**AddCookie**](docs/CookiesApi.md#addcookie) | **Post** /session/{sessionId}/cookie | Add cookie
*CookiesApi* | [**DeleteAllCookies**](docs/CookiesApi.md#deleteallcookies) | **Delete** /session/{sessionId}/cookie | Delete all cookies
*CookiesApi* | [**DeleteCookie**](docs/CookiesApi.md#deletecookie) | **Delete** /session/{sessionId}/cookie/{name} | Delete cookie
*CookiesApi* | [**GetAllCookies**](docs/CookiesApi.md#getallcookies) | **Get** /session/{sessionId}/cookie | Get all cookies
*CookiesApi* | [**GetNamedCookie**](docs/CookiesApi.md#getnamedcookie) | **Get** /session/{sessionId}/cookie/{name} | Get named cookie
*DocumentApi* | [**ExecuteScript**](docs/DocumentApi.md#executescript) | **Post** /session/{sessionId}/execute/sync | Execute script
*DocumentApi* | [**ExecuteScriptAsync**](docs/DocumentApi.md#executescriptasync) | **Post** /session/{sessionId}/execute/async | Execute script asynchronously
*DocumentApi* | [**GetPageSource**](docs/DocumentApi.md#getpagesource) | **Get** /session/{sessionId}/source | Get page source
*DocumentApi* | [**UploadFile**](docs/DocumentApi.md#uploadfile) | **Post** /session/{sessionId}/file | Upload file
*ElementsApi* | [**ElementClear**](docs/ElementsApi.md#elementclear) | **Post** /session/{sessionId}/element/{elementId}/clear | Clear element
*ElementsApi* | [**ElementClick**](docs/ElementsApi.md#elementclick) | **Post** /session/{sessionId}/element/{elementId}/click | Click on element
*ElementsApi* | [**ElementSendKeys**](docs/ElementsApi.md#elementsendkeys) | **Post** /session/{sessionId}/element/{elementId}/value | Send keys to element
*ElementsApi* | [**FindElement**](docs/ElementsApi.md#findelement) | **Post** /session/{sessionId}/element | Find element
*ElementsApi* | [**FindElementFromElement**](docs/ElementsApi.md#findelementfromelement) | **Post** /session/{sessionId}/element/{elementId}/element | Find element from element
*ElementsApi* | [**FindElements**](docs/ElementsApi.md#findelements) | **Post** /session/{sessionId}/elements | Find elements
*ElementsApi* | [**FindElementsFromElement**](docs/ElementsApi.md#findelementsfromelement) | **Post** /session/{sessionId}/element/{elementId}/elements | Find elements from element
*ElementsApi* | [**GetActiveElement**](docs/ElementsApi.md#getactiveelement) | **Get** /session/{sessionId}/element/active | Get active element
*ElementsApi* | [**GetElementAttribute**](docs/ElementsApi.md#getelementattribute) | **Get** /session/{sessionId}/element/{elementId}/attribute/{name} | Get element attribute
*ElementsApi* | [**GetElementCSSProperty**](docs/ElementsApi.md#getelementcssproperty) | **Get** /session/{sessionId}/element/{elementId}/css/{propertyName} | Get element CSS property
*ElementsApi* | [**GetElementComputedLabel**](docs/ElementsApi.md#getelementcomputedlabel) | **Get** /session/{sessionId}/element/{elementId}/computedlabel | Get element computed accessibility label
*ElementsApi* | [**GetElementComputedRole**](docs/ElementsApi.md#getelementcomputedrole) | **Get** /session/{sessionId}/element/{elementId}/computedrole | Get element computed accessibility role
*ElementsApi* | [**GetElementProperty**](docs/ElementsApi.md#getelementproperty) | **Get** /session/{sessionId}/element/{elementId}/property/{name} | Get element property
*ElementsApi* | [**GetElementRect**](docs/ElementsApi.md#getelementrect) | **Get** /session/{sessionId}/element/{elementId}/rect | Get element rect
*ElementsApi* | [**GetElementTagName**](docs/ElementsApi.md#getelementtagname) | **Get** /session/{sessionId}/element/{elementId}/name | Get element tag name
*ElementsApi* | [**GetElementText**](docs/ElementsApi.md#getelementtext) | **Get** /session/{sessionId}/element/{elementId}/text | Get element text
*ElementsApi* | [**IsElementDisplayed**](docs/ElementsApi.md#iselementdisplayed) | **Get** /session/{sessionId}/element/{elementId}/displayed | Is element displayed
*ElementsApi* | [**IsElementEnabled**](docs/ElementsApi.md#iselementenabled) | **Get** /session/{sessionId}/element/{elementId}/enabled | Is element enabled
*ElementsApi* | [**IsElementSelected**](docs/ElementsApi.md#iselementselected) | **Get** /session/{sessionId}/element/{elementId}/selected | Is element selected
*MozillaApi* | [**GetContext**](docs/MozillaApi.md#getcontext) | **Get** /session/{sessionId}/moz/context | Get current context
*MozillaApi* | [**InstallAddon**](docs/MozillaApi.md#installaddon) | **Post** /session/{sessionId}/moz/addon/install | Install addon
*MozillaApi* | [**SetContext**](docs/MozillaApi.md#setcontext) | **Post** /session/{sessionId}/moz/context | Set current context
*MozillaApi* | [**TakeFullScreenshot**](docs/MozillaApi.md#takefullscreenshot) | **Get** /session/{sessionId}/moz/screenshot/full | Takes full page screenshot
*MozillaApi* | [**UninstallAddon**](docs/MozillaApi.md#uninstalladdon) | **Post** /session/{sessionId}/moz/addon/uninstall | Uninstall addon
*NavigationApi* | [**GetCurrentUrl**](docs/NavigationApi.md#getcurrenturl) | **Get** /session/{sessionId}/url | Returns current URL
*NavigationApi* | [**GetPageTitle**](docs/NavigationApi.md#getpagetitle) | **Get** /session/{sessionId}/title | Returns current page title
*NavigationApi* | [**NavigateBack**](docs/NavigationApi.md#navigateback) | **Post** /session/{sessionId}/back | Navigates to the previous page
*NavigationApi* | [**NavigateForward**](docs/NavigationApi.md#navigateforward) | **Post** /session/{sessionId}/forward | Navigates to the next page
*NavigationApi* | [**NavigateTo**](docs/NavigationApi.md#navigateto) | **Post** /session/{sessionId}/url | Navigates to URL
*NavigationApi* | [**RefreshPage**](docs/NavigationApi.md#refreshpage) | **Post** /session/{sessionId}/refresh | Reloads current page
*PrintApi* | [**PrintPage**](docs/PrintApi.md#printpage) | **Post** /session/{sessionId}/print | Print page to PDF
*PromptsApi* | [**AcceptAlert**](docs/PromptsApi.md#acceptalert) | **Post** /session/{sessionId}/alert/accept | Accept alert
*PromptsApi* | [**DismissAlert**](docs/PromptsApi.md#dismissalert) | **Post** /session/{sessionId}/alert/dismiss | Dismiss alert
*PromptsApi* | [**GetAlertText**](docs/PromptsApi.md#getalerttext) | **Get** /session/{sessionId}/alert/text | Get alert text
*PromptsApi* | [**SendAlertText**](docs/PromptsApi.md#sendalerttext) | **Post** /session/{sessionId}/alert/text | Send alert text
*ScreenshotsApi* | [**TakeElementScreenshot**](docs/ScreenshotsApi.md#takeelementscreenshot) | **Get** /session/{sessionId}/element/{elementId}/screenshot | Takes element screenshot
*ScreenshotsApi* | [**TakeScreenshot**](docs/ScreenshotsApi.md#takescreenshot) | **Get** /session/{sessionId}/screenshot | Takes page screenshot
*SessionsApi* | [**CreateSession**](docs/SessionsApi.md#createsession) | **Post** /session | Creates new Selenium session
*SessionsApi* | [**DeleteSession**](docs/SessionsApi.md#deletesession) | **Delete** /session/{sessionId} | Deletes existing Selenium session
*SessionsApi* | [**GetStatus**](docs/SessionsApi.md#getstatus) | **Get** /status | Gets Selenium API status information
*TimeoutsApi* | [**GetTimeouts**](docs/TimeoutsApi.md#gettimeouts) | **Get** /session/{sessionId}/timeouts | Get session timeouts
*TimeoutsApi* | [**SetTimeouts**](docs/TimeoutsApi.md#settimeouts) | **Post** /session/{sessionId}/timeouts | Adjusts session timeouts


## Documentation For Models

 - [Action](docs/Action.md)
 - [ActionOrigin](docs/ActionOrigin.md)
 - [ActionSequence](docs/ActionSequence.md)
 - [ActionSequenceParameters](docs/ActionSequenceParameters.md)
 - [ActionsRequest](docs/ActionsRequest.md)
 - [AddonInstallRequest](docs/AddonInstallRequest.md)
 - [AddonUninstallRequest](docs/AddonUninstallRequest.md)
 - [AnyResponse](docs/AnyResponse.md)
 - [ArrayResponse](docs/ArrayResponse.md)
 - [BooleanResponse](docs/BooleanResponse.md)
 - [Capabilities](docs/Capabilities.md)
 - [ChromeOptions](docs/ChromeOptions.md)
 - [ChromiumLogLevel](docs/ChromiumLogLevel.md)
 - [ClipboardData](docs/ClipboardData.md)
 - [ContextRequest](docs/ContextRequest.md)
 - [Cookie](docs/Cookie.md)
 - [CookieRequest](docs/CookieRequest.md)
 - [CookieResponse](docs/CookieResponse.md)
 - [CookiesResponse](docs/CookiesResponse.md)
 - [EdgeOptions](docs/EdgeOptions.md)
 - [ElementSendKeysRequest](docs/ElementSendKeysRequest.md)
 - [EmptyResponse](docs/EmptyResponse.md)
 - [ErrorCode](docs/ErrorCode.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [ErrorResponseValue](docs/ErrorResponseValue.md)
 - [FileUploadRequest](docs/FileUploadRequest.md)
 - [FileUploadResponse](docs/FileUploadResponse.md)
 - [FindElementRequest](docs/FindElementRequest.md)
 - [FindElementResponse](docs/FindElementResponse.md)
 - [FindElementResponseValue](docs/FindElementResponseValue.md)
 - [FindElementsResponse](docs/FindElementsResponse.md)
 - [FirefoxContext](docs/FirefoxContext.md)
 - [FirefoxLogLevel](docs/FirefoxLogLevel.md)
 - [FirefoxOptions](docs/FirefoxOptions.md)
 - [FirefoxOptionsLog](docs/FirefoxOptionsLog.md)
 - [FrameId](docs/FrameId.md)
 - [GetWindowHandlesResponse](docs/GetWindowHandlesResponse.md)
 - [KeyAction](docs/KeyAction.md)
 - [LocatorStrategy](docs/LocatorStrategy.md)
 - [LogLevel](docs/LogLevel.md)
 - [LogType](docs/LogType.md)
 - [LoggingPrefs](docs/LoggingPrefs.md)
 - [MobileEmulation](docs/MobileEmulation.md)
 - [MobileEmulationDeviceMetrics](docs/MobileEmulationDeviceMetrics.md)
 - [MoonLogLevel](docs/MoonLogLevel.md)
 - [MoonMobileDevice](docs/MoonMobileDevice.md)
 - [MoonOptions](docs/MoonOptions.md)
 - [NewSessionRequest](docs/NewSessionRequest.md)
 - [NewSessionRequestCapabilities](docs/NewSessionRequestCapabilities.md)
 - [NewSessionResponse](docs/NewSessionResponse.md)
 - [NewSessionResponseValue](docs/NewSessionResponseValue.md)
 - [NewWindowRequest](docs/NewWindowRequest.md)
 - [NewWindowResponse](docs/NewWindowResponse.md)
 - [NewWindowResponseValue](docs/NewWindowResponseValue.md)
 - [NullAction](docs/NullAction.md)
 - [NullableStringResponse](docs/NullableStringResponse.md)
 - [OperaOptions](docs/OperaOptions.md)
 - [PerfLoggingPrefs](docs/PerfLoggingPrefs.md)
 - [PointerAction](docs/PointerAction.md)
 - [PreferenceValue](docs/PreferenceValue.md)
 - [PrintRequest](docs/PrintRequest.md)
 - [PrintRequestOptions](docs/PrintRequestOptions.md)
 - [PrintRequestOptionsMargin](docs/PrintRequestOptionsMargin.md)
 - [PrintRequestOptionsPage](docs/PrintRequestOptionsPage.md)
 - [Proxy](docs/Proxy.md)
 - [Rect](docs/Rect.md)
 - [RectResponse](docs/RectResponse.md)
 - [ScriptRequest](docs/ScriptRequest.md)
 - [SelenoidOptions](docs/SelenoidOptions.md)
 - [SendAlertTextRequest](docs/SendAlertTextRequest.md)
 - [StatusResponse](docs/StatusResponse.md)
 - [StatusResponseValue](docs/StatusResponseValue.md)
 - [StringOrigin](docs/StringOrigin.md)
 - [StringResponse](docs/StringResponse.md)
 - [SwitchToFrameRequest](docs/SwitchToFrameRequest.md)
 - [SwitchToWindowRequest](docs/SwitchToWindowRequest.md)
 - [Timeouts](docs/Timeouts.md)
 - [TimeoutsResponse](docs/TimeoutsResponse.md)
 - [UrlRequest](docs/UrlRequest.md)
 - [WebElement](docs/WebElement.md)
 - [WebElementOrigin](docs/WebElementOrigin.md)
 - [WheelAction](docs/WheelAction.md)


## Documentation For Authorization



### BasicAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@aerokube.com

