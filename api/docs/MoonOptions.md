# MoonOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalFonts** | Pointer to **bool** | Enable additional regional fonts (Chinese, Japanese, Thai and so on) | [optional] 
**Context** | Pointer to **string** | Moon context archive URL | [optional] 
**EnableVideo** | Pointer to **bool** |  | [optional] 
**Env** | Pointer to **[]string** |  | [optional] 
**Hosts** | Pointer to **[]string** | Additional /etc/hosts entries | [optional] 
**Labels** | Pointer to **map[string]string** | Custom pod labels | [optional] 
**LogLevel** | Pointer to [**MoonLogLevel**](MoonLogLevel.md) |  | [optional] 
**MobileDevice** | Pointer to [**MoonMobileDevice**](MoonMobileDevice.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Nameservers** | Pointer to **[]string** | Custom DNS servers | [optional] 
**Pattern** | Pointer to **string** | S3 upload pattern | [optional] 
**SessionTimeout** | Pointer to **string** |  | [optional] 
**ScreenResolution** | Pointer to **string** | Screen resolution in format 1920x1080x32 or 1920x1080 | [optional] 
**VideoFrameRate** | Pointer to **int64** |  | [optional] 
**VideoName** | Pointer to **string** |  | [optional] 
**VideoScreenSize** | Pointer to **string** | Video screen size in format 1920x1080 | [optional] 

## Methods

### NewMoonOptions

`func NewMoonOptions() *MoonOptions`

NewMoonOptions instantiates a new MoonOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoonOptionsWithDefaults

`func NewMoonOptionsWithDefaults() *MoonOptions`

NewMoonOptionsWithDefaults instantiates a new MoonOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalFonts

`func (o *MoonOptions) GetAdditionalFonts() bool`

GetAdditionalFonts returns the AdditionalFonts field if non-nil, zero value otherwise.

### GetAdditionalFontsOk

`func (o *MoonOptions) GetAdditionalFontsOk() (*bool, bool)`

GetAdditionalFontsOk returns a tuple with the AdditionalFonts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFonts

`func (o *MoonOptions) SetAdditionalFonts(v bool)`

SetAdditionalFonts sets AdditionalFonts field to given value.

### HasAdditionalFonts

`func (o *MoonOptions) HasAdditionalFonts() bool`

HasAdditionalFonts returns a boolean if a field has been set.

### GetContext

`func (o *MoonOptions) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *MoonOptions) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *MoonOptions) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *MoonOptions) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetEnableVideo

`func (o *MoonOptions) GetEnableVideo() bool`

GetEnableVideo returns the EnableVideo field if non-nil, zero value otherwise.

### GetEnableVideoOk

`func (o *MoonOptions) GetEnableVideoOk() (*bool, bool)`

GetEnableVideoOk returns a tuple with the EnableVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVideo

`func (o *MoonOptions) SetEnableVideo(v bool)`

SetEnableVideo sets EnableVideo field to given value.

### HasEnableVideo

`func (o *MoonOptions) HasEnableVideo() bool`

HasEnableVideo returns a boolean if a field has been set.

### GetEnv

`func (o *MoonOptions) GetEnv() []string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *MoonOptions) GetEnvOk() (*[]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *MoonOptions) SetEnv(v []string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *MoonOptions) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetHosts

`func (o *MoonOptions) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *MoonOptions) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *MoonOptions) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *MoonOptions) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetLabels

`func (o *MoonOptions) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *MoonOptions) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *MoonOptions) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *MoonOptions) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLogLevel

`func (o *MoonOptions) GetLogLevel() MoonLogLevel`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *MoonOptions) GetLogLevelOk() (*MoonLogLevel, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *MoonOptions) SetLogLevel(v MoonLogLevel)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *MoonOptions) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetMobileDevice

`func (o *MoonOptions) GetMobileDevice() MoonMobileDevice`

GetMobileDevice returns the MobileDevice field if non-nil, zero value otherwise.

### GetMobileDeviceOk

`func (o *MoonOptions) GetMobileDeviceOk() (*MoonMobileDevice, bool)`

GetMobileDeviceOk returns a tuple with the MobileDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileDevice

`func (o *MoonOptions) SetMobileDevice(v MoonMobileDevice)`

SetMobileDevice sets MobileDevice field to given value.

### HasMobileDevice

`func (o *MoonOptions) HasMobileDevice() bool`

HasMobileDevice returns a boolean if a field has been set.

### GetName

`func (o *MoonOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MoonOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MoonOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MoonOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameservers

`func (o *MoonOptions) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *MoonOptions) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *MoonOptions) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *MoonOptions) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetPattern

`func (o *MoonOptions) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *MoonOptions) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *MoonOptions) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *MoonOptions) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetSessionTimeout

`func (o *MoonOptions) GetSessionTimeout() string`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *MoonOptions) GetSessionTimeoutOk() (*string, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *MoonOptions) SetSessionTimeout(v string)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *MoonOptions) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.

### GetScreenResolution

`func (o *MoonOptions) GetScreenResolution() string`

GetScreenResolution returns the ScreenResolution field if non-nil, zero value otherwise.

### GetScreenResolutionOk

`func (o *MoonOptions) GetScreenResolutionOk() (*string, bool)`

GetScreenResolutionOk returns a tuple with the ScreenResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenResolution

`func (o *MoonOptions) SetScreenResolution(v string)`

SetScreenResolution sets ScreenResolution field to given value.

### HasScreenResolution

`func (o *MoonOptions) HasScreenResolution() bool`

HasScreenResolution returns a boolean if a field has been set.

### GetVideoFrameRate

`func (o *MoonOptions) GetVideoFrameRate() int64`

GetVideoFrameRate returns the VideoFrameRate field if non-nil, zero value otherwise.

### GetVideoFrameRateOk

`func (o *MoonOptions) GetVideoFrameRateOk() (*int64, bool)`

GetVideoFrameRateOk returns a tuple with the VideoFrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoFrameRate

`func (o *MoonOptions) SetVideoFrameRate(v int64)`

SetVideoFrameRate sets VideoFrameRate field to given value.

### HasVideoFrameRate

`func (o *MoonOptions) HasVideoFrameRate() bool`

HasVideoFrameRate returns a boolean if a field has been set.

### GetVideoName

`func (o *MoonOptions) GetVideoName() string`

GetVideoName returns the VideoName field if non-nil, zero value otherwise.

### GetVideoNameOk

`func (o *MoonOptions) GetVideoNameOk() (*string, bool)`

GetVideoNameOk returns a tuple with the VideoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoName

`func (o *MoonOptions) SetVideoName(v string)`

SetVideoName sets VideoName field to given value.

### HasVideoName

`func (o *MoonOptions) HasVideoName() bool`

HasVideoName returns a boolean if a field has been set.

### GetVideoScreenSize

`func (o *MoonOptions) GetVideoScreenSize() string`

GetVideoScreenSize returns the VideoScreenSize field if non-nil, zero value otherwise.

### GetVideoScreenSizeOk

`func (o *MoonOptions) GetVideoScreenSizeOk() (*string, bool)`

GetVideoScreenSizeOk returns a tuple with the VideoScreenSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoScreenSize

`func (o *MoonOptions) SetVideoScreenSize(v string)`

SetVideoScreenSize sets VideoScreenSize field to given value.

### HasVideoScreenSize

`func (o *MoonOptions) HasVideoScreenSize() bool`

HasVideoScreenSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


