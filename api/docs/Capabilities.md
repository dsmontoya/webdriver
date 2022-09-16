# Capabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrowserName** | Pointer to **string** |  | [optional] 
**BrowserVersion** | Pointer to **string** |  | [optional] 
**PlatformName** | Pointer to **string** |  | [optional] 
**AcceptInsecureCerts** | Pointer to **bool** |  | [optional] 
**PageLoadStrategy** | Pointer to **string** |  | [optional] 
**Proxy** | Pointer to [**Proxy**](Proxy.md) |  | [optional] 
**SetWindowRect** | Pointer to **bool** |  | [optional] 
**Timeouts** | Pointer to [**Timeouts**](Timeouts.md) |  | [optional] 
**StrictFileInteractability** | Pointer to **bool** |  | [optional] 
**UnhandledPromptBehavior** | Pointer to **string** |  | [optional] 
**GoogloggingPrefs** | Pointer to [**LoggingPrefs**](LoggingPrefs.md) |  | [optional] 
**GoogchromeOptions** | Pointer to [**ChromeOptions**](ChromeOptions.md) |  | [optional] 
**Moonoptions** | Pointer to [**MoonOptions**](MoonOptions.md) |  | [optional] 
**MozfirefoxOptions** | Pointer to [**FirefoxOptions**](FirefoxOptions.md) |  | [optional] 
**MsedgeOptions** | Pointer to [**EdgeOptions**](EdgeOptions.md) |  | [optional] 
**OperaOptions** | Pointer to [**EdgeOptions**](EdgeOptions.md) |  | [optional] 
**Selenoidoptions** | Pointer to [**SelenoidOptions**](SelenoidOptions.md) |  | [optional] 
**SafariautomaticInspection** | Pointer to **bool** |  | [optional] 
**SafariautomaticProfiling** | Pointer to **bool** |  | [optional] 

## Methods

### NewCapabilities

`func NewCapabilities() *Capabilities`

NewCapabilities instantiates a new Capabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitiesWithDefaults

`func NewCapabilitiesWithDefaults() *Capabilities`

NewCapabilitiesWithDefaults instantiates a new Capabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrowserName

`func (o *Capabilities) GetBrowserName() string`

GetBrowserName returns the BrowserName field if non-nil, zero value otherwise.

### GetBrowserNameOk

`func (o *Capabilities) GetBrowserNameOk() (*string, bool)`

GetBrowserNameOk returns a tuple with the BrowserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserName

`func (o *Capabilities) SetBrowserName(v string)`

SetBrowserName sets BrowserName field to given value.

### HasBrowserName

`func (o *Capabilities) HasBrowserName() bool`

HasBrowserName returns a boolean if a field has been set.

### GetBrowserVersion

`func (o *Capabilities) GetBrowserVersion() string`

GetBrowserVersion returns the BrowserVersion field if non-nil, zero value otherwise.

### GetBrowserVersionOk

`func (o *Capabilities) GetBrowserVersionOk() (*string, bool)`

GetBrowserVersionOk returns a tuple with the BrowserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserVersion

`func (o *Capabilities) SetBrowserVersion(v string)`

SetBrowserVersion sets BrowserVersion field to given value.

### HasBrowserVersion

`func (o *Capabilities) HasBrowserVersion() bool`

HasBrowserVersion returns a boolean if a field has been set.

### GetPlatformName

`func (o *Capabilities) GetPlatformName() string`

GetPlatformName returns the PlatformName field if non-nil, zero value otherwise.

### GetPlatformNameOk

`func (o *Capabilities) GetPlatformNameOk() (*string, bool)`

GetPlatformNameOk returns a tuple with the PlatformName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformName

`func (o *Capabilities) SetPlatformName(v string)`

SetPlatformName sets PlatformName field to given value.

### HasPlatformName

`func (o *Capabilities) HasPlatformName() bool`

HasPlatformName returns a boolean if a field has been set.

### GetAcceptInsecureCerts

`func (o *Capabilities) GetAcceptInsecureCerts() bool`

GetAcceptInsecureCerts returns the AcceptInsecureCerts field if non-nil, zero value otherwise.

### GetAcceptInsecureCertsOk

`func (o *Capabilities) GetAcceptInsecureCertsOk() (*bool, bool)`

GetAcceptInsecureCertsOk returns a tuple with the AcceptInsecureCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptInsecureCerts

`func (o *Capabilities) SetAcceptInsecureCerts(v bool)`

SetAcceptInsecureCerts sets AcceptInsecureCerts field to given value.

### HasAcceptInsecureCerts

`func (o *Capabilities) HasAcceptInsecureCerts() bool`

HasAcceptInsecureCerts returns a boolean if a field has been set.

### GetPageLoadStrategy

`func (o *Capabilities) GetPageLoadStrategy() string`

GetPageLoadStrategy returns the PageLoadStrategy field if non-nil, zero value otherwise.

### GetPageLoadStrategyOk

`func (o *Capabilities) GetPageLoadStrategyOk() (*string, bool)`

GetPageLoadStrategyOk returns a tuple with the PageLoadStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageLoadStrategy

`func (o *Capabilities) SetPageLoadStrategy(v string)`

SetPageLoadStrategy sets PageLoadStrategy field to given value.

### HasPageLoadStrategy

`func (o *Capabilities) HasPageLoadStrategy() bool`

HasPageLoadStrategy returns a boolean if a field has been set.

### GetProxy

`func (o *Capabilities) GetProxy() Proxy`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *Capabilities) GetProxyOk() (*Proxy, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *Capabilities) SetProxy(v Proxy)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *Capabilities) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetSetWindowRect

`func (o *Capabilities) GetSetWindowRect() bool`

GetSetWindowRect returns the SetWindowRect field if non-nil, zero value otherwise.

### GetSetWindowRectOk

`func (o *Capabilities) GetSetWindowRectOk() (*bool, bool)`

GetSetWindowRectOk returns a tuple with the SetWindowRect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetWindowRect

`func (o *Capabilities) SetSetWindowRect(v bool)`

SetSetWindowRect sets SetWindowRect field to given value.

### HasSetWindowRect

`func (o *Capabilities) HasSetWindowRect() bool`

HasSetWindowRect returns a boolean if a field has been set.

### GetTimeouts

`func (o *Capabilities) GetTimeouts() Timeouts`

GetTimeouts returns the Timeouts field if non-nil, zero value otherwise.

### GetTimeoutsOk

`func (o *Capabilities) GetTimeoutsOk() (*Timeouts, bool)`

GetTimeoutsOk returns a tuple with the Timeouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeouts

`func (o *Capabilities) SetTimeouts(v Timeouts)`

SetTimeouts sets Timeouts field to given value.

### HasTimeouts

`func (o *Capabilities) HasTimeouts() bool`

HasTimeouts returns a boolean if a field has been set.

### GetStrictFileInteractability

`func (o *Capabilities) GetStrictFileInteractability() bool`

GetStrictFileInteractability returns the StrictFileInteractability field if non-nil, zero value otherwise.

### GetStrictFileInteractabilityOk

`func (o *Capabilities) GetStrictFileInteractabilityOk() (*bool, bool)`

GetStrictFileInteractabilityOk returns a tuple with the StrictFileInteractability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrictFileInteractability

`func (o *Capabilities) SetStrictFileInteractability(v bool)`

SetStrictFileInteractability sets StrictFileInteractability field to given value.

### HasStrictFileInteractability

`func (o *Capabilities) HasStrictFileInteractability() bool`

HasStrictFileInteractability returns a boolean if a field has been set.

### GetUnhandledPromptBehavior

`func (o *Capabilities) GetUnhandledPromptBehavior() string`

GetUnhandledPromptBehavior returns the UnhandledPromptBehavior field if non-nil, zero value otherwise.

### GetUnhandledPromptBehaviorOk

`func (o *Capabilities) GetUnhandledPromptBehaviorOk() (*string, bool)`

GetUnhandledPromptBehaviorOk returns a tuple with the UnhandledPromptBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhandledPromptBehavior

`func (o *Capabilities) SetUnhandledPromptBehavior(v string)`

SetUnhandledPromptBehavior sets UnhandledPromptBehavior field to given value.

### HasUnhandledPromptBehavior

`func (o *Capabilities) HasUnhandledPromptBehavior() bool`

HasUnhandledPromptBehavior returns a boolean if a field has been set.

### GetGoogloggingPrefs

`func (o *Capabilities) GetGoogloggingPrefs() LoggingPrefs`

GetGoogloggingPrefs returns the GoogloggingPrefs field if non-nil, zero value otherwise.

### GetGoogloggingPrefsOk

`func (o *Capabilities) GetGoogloggingPrefsOk() (*LoggingPrefs, bool)`

GetGoogloggingPrefsOk returns a tuple with the GoogloggingPrefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogloggingPrefs

`func (o *Capabilities) SetGoogloggingPrefs(v LoggingPrefs)`

SetGoogloggingPrefs sets GoogloggingPrefs field to given value.

### HasGoogloggingPrefs

`func (o *Capabilities) HasGoogloggingPrefs() bool`

HasGoogloggingPrefs returns a boolean if a field has been set.

### GetGoogchromeOptions

`func (o *Capabilities) GetGoogchromeOptions() ChromeOptions`

GetGoogchromeOptions returns the GoogchromeOptions field if non-nil, zero value otherwise.

### GetGoogchromeOptionsOk

`func (o *Capabilities) GetGoogchromeOptionsOk() (*ChromeOptions, bool)`

GetGoogchromeOptionsOk returns a tuple with the GoogchromeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogchromeOptions

`func (o *Capabilities) SetGoogchromeOptions(v ChromeOptions)`

SetGoogchromeOptions sets GoogchromeOptions field to given value.

### HasGoogchromeOptions

`func (o *Capabilities) HasGoogchromeOptions() bool`

HasGoogchromeOptions returns a boolean if a field has been set.

### GetMoonoptions

`func (o *Capabilities) GetMoonoptions() MoonOptions`

GetMoonoptions returns the Moonoptions field if non-nil, zero value otherwise.

### GetMoonoptionsOk

`func (o *Capabilities) GetMoonoptionsOk() (*MoonOptions, bool)`

GetMoonoptionsOk returns a tuple with the Moonoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoonoptions

`func (o *Capabilities) SetMoonoptions(v MoonOptions)`

SetMoonoptions sets Moonoptions field to given value.

### HasMoonoptions

`func (o *Capabilities) HasMoonoptions() bool`

HasMoonoptions returns a boolean if a field has been set.

### GetMozfirefoxOptions

`func (o *Capabilities) GetMozfirefoxOptions() FirefoxOptions`

GetMozfirefoxOptions returns the MozfirefoxOptions field if non-nil, zero value otherwise.

### GetMozfirefoxOptionsOk

`func (o *Capabilities) GetMozfirefoxOptionsOk() (*FirefoxOptions, bool)`

GetMozfirefoxOptionsOk returns a tuple with the MozfirefoxOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMozfirefoxOptions

`func (o *Capabilities) SetMozfirefoxOptions(v FirefoxOptions)`

SetMozfirefoxOptions sets MozfirefoxOptions field to given value.

### HasMozfirefoxOptions

`func (o *Capabilities) HasMozfirefoxOptions() bool`

HasMozfirefoxOptions returns a boolean if a field has been set.

### GetMsedgeOptions

`func (o *Capabilities) GetMsedgeOptions() EdgeOptions`

GetMsedgeOptions returns the MsedgeOptions field if non-nil, zero value otherwise.

### GetMsedgeOptionsOk

`func (o *Capabilities) GetMsedgeOptionsOk() (*EdgeOptions, bool)`

GetMsedgeOptionsOk returns a tuple with the MsedgeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsedgeOptions

`func (o *Capabilities) SetMsedgeOptions(v EdgeOptions)`

SetMsedgeOptions sets MsedgeOptions field to given value.

### HasMsedgeOptions

`func (o *Capabilities) HasMsedgeOptions() bool`

HasMsedgeOptions returns a boolean if a field has been set.

### GetOperaOptions

`func (o *Capabilities) GetOperaOptions() EdgeOptions`

GetOperaOptions returns the OperaOptions field if non-nil, zero value otherwise.

### GetOperaOptionsOk

`func (o *Capabilities) GetOperaOptionsOk() (*EdgeOptions, bool)`

GetOperaOptionsOk returns a tuple with the OperaOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperaOptions

`func (o *Capabilities) SetOperaOptions(v EdgeOptions)`

SetOperaOptions sets OperaOptions field to given value.

### HasOperaOptions

`func (o *Capabilities) HasOperaOptions() bool`

HasOperaOptions returns a boolean if a field has been set.

### GetSelenoidoptions

`func (o *Capabilities) GetSelenoidoptions() SelenoidOptions`

GetSelenoidoptions returns the Selenoidoptions field if non-nil, zero value otherwise.

### GetSelenoidoptionsOk

`func (o *Capabilities) GetSelenoidoptionsOk() (*SelenoidOptions, bool)`

GetSelenoidoptionsOk returns a tuple with the Selenoidoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelenoidoptions

`func (o *Capabilities) SetSelenoidoptions(v SelenoidOptions)`

SetSelenoidoptions sets Selenoidoptions field to given value.

### HasSelenoidoptions

`func (o *Capabilities) HasSelenoidoptions() bool`

HasSelenoidoptions returns a boolean if a field has been set.

### GetSafariautomaticInspection

`func (o *Capabilities) GetSafariautomaticInspection() bool`

GetSafariautomaticInspection returns the SafariautomaticInspection field if non-nil, zero value otherwise.

### GetSafariautomaticInspectionOk

`func (o *Capabilities) GetSafariautomaticInspectionOk() (*bool, bool)`

GetSafariautomaticInspectionOk returns a tuple with the SafariautomaticInspection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafariautomaticInspection

`func (o *Capabilities) SetSafariautomaticInspection(v bool)`

SetSafariautomaticInspection sets SafariautomaticInspection field to given value.

### HasSafariautomaticInspection

`func (o *Capabilities) HasSafariautomaticInspection() bool`

HasSafariautomaticInspection returns a boolean if a field has been set.

### GetSafariautomaticProfiling

`func (o *Capabilities) GetSafariautomaticProfiling() bool`

GetSafariautomaticProfiling returns the SafariautomaticProfiling field if non-nil, zero value otherwise.

### GetSafariautomaticProfilingOk

`func (o *Capabilities) GetSafariautomaticProfilingOk() (*bool, bool)`

GetSafariautomaticProfilingOk returns a tuple with the SafariautomaticProfiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafariautomaticProfiling

`func (o *Capabilities) SetSafariautomaticProfiling(v bool)`

SetSafariautomaticProfiling sets SafariautomaticProfiling field to given value.

### HasSafariautomaticProfiling

`func (o *Capabilities) HasSafariautomaticProfiling() bool`

HasSafariautomaticProfiling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


