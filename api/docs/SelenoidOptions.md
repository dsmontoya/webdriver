# SelenoidOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableLog** | Pointer to **bool** |  | [optional] 
**EnableVideo** | Pointer to **bool** |  | [optional] 
**EnableVNC** | Pointer to **bool** |  | [optional] 
**Env** | Pointer to **[]string** |  | [optional] 
**Labels** | Pointer to **map[string]string** | Custom container labels | [optional] 
**LogName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**S3KeyPattern** | Pointer to **string** |  | [optional] 
**SessionTimeout** | Pointer to **string** |  | [optional] 
**ScreenResolution** | Pointer to **string** | Screen resolution in format 1920x1080x32 or 1920x1080 | [optional] 
**TimeZone** | Pointer to **string** | Time zone value like Europe/Moscow | [optional] 
**VideoFrameRate** | Pointer to **int64** |  | [optional] 
**VideoName** | Pointer to **string** |  | [optional] 
**VideoScreenSize** | Pointer to **string** | Video screen size in format 1920x1080 | [optional] 

## Methods

### NewSelenoidOptions

`func NewSelenoidOptions() *SelenoidOptions`

NewSelenoidOptions instantiates a new SelenoidOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelenoidOptionsWithDefaults

`func NewSelenoidOptionsWithDefaults() *SelenoidOptions`

NewSelenoidOptionsWithDefaults instantiates a new SelenoidOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableLog

`func (o *SelenoidOptions) GetEnableLog() bool`

GetEnableLog returns the EnableLog field if non-nil, zero value otherwise.

### GetEnableLogOk

`func (o *SelenoidOptions) GetEnableLogOk() (*bool, bool)`

GetEnableLogOk returns a tuple with the EnableLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLog

`func (o *SelenoidOptions) SetEnableLog(v bool)`

SetEnableLog sets EnableLog field to given value.

### HasEnableLog

`func (o *SelenoidOptions) HasEnableLog() bool`

HasEnableLog returns a boolean if a field has been set.

### GetEnableVideo

`func (o *SelenoidOptions) GetEnableVideo() bool`

GetEnableVideo returns the EnableVideo field if non-nil, zero value otherwise.

### GetEnableVideoOk

`func (o *SelenoidOptions) GetEnableVideoOk() (*bool, bool)`

GetEnableVideoOk returns a tuple with the EnableVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVideo

`func (o *SelenoidOptions) SetEnableVideo(v bool)`

SetEnableVideo sets EnableVideo field to given value.

### HasEnableVideo

`func (o *SelenoidOptions) HasEnableVideo() bool`

HasEnableVideo returns a boolean if a field has been set.

### GetEnableVNC

`func (o *SelenoidOptions) GetEnableVNC() bool`

GetEnableVNC returns the EnableVNC field if non-nil, zero value otherwise.

### GetEnableVNCOk

`func (o *SelenoidOptions) GetEnableVNCOk() (*bool, bool)`

GetEnableVNCOk returns a tuple with the EnableVNC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVNC

`func (o *SelenoidOptions) SetEnableVNC(v bool)`

SetEnableVNC sets EnableVNC field to given value.

### HasEnableVNC

`func (o *SelenoidOptions) HasEnableVNC() bool`

HasEnableVNC returns a boolean if a field has been set.

### GetEnv

`func (o *SelenoidOptions) GetEnv() []string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *SelenoidOptions) GetEnvOk() (*[]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *SelenoidOptions) SetEnv(v []string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *SelenoidOptions) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetLabels

`func (o *SelenoidOptions) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SelenoidOptions) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SelenoidOptions) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SelenoidOptions) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLogName

`func (o *SelenoidOptions) GetLogName() string`

GetLogName returns the LogName field if non-nil, zero value otherwise.

### GetLogNameOk

`func (o *SelenoidOptions) GetLogNameOk() (*string, bool)`

GetLogNameOk returns a tuple with the LogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogName

`func (o *SelenoidOptions) SetLogName(v string)`

SetLogName sets LogName field to given value.

### HasLogName

`func (o *SelenoidOptions) HasLogName() bool`

HasLogName returns a boolean if a field has been set.

### GetName

`func (o *SelenoidOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SelenoidOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SelenoidOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SelenoidOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetS3KeyPattern

`func (o *SelenoidOptions) GetS3KeyPattern() string`

GetS3KeyPattern returns the S3KeyPattern field if non-nil, zero value otherwise.

### GetS3KeyPatternOk

`func (o *SelenoidOptions) GetS3KeyPatternOk() (*string, bool)`

GetS3KeyPatternOk returns a tuple with the S3KeyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3KeyPattern

`func (o *SelenoidOptions) SetS3KeyPattern(v string)`

SetS3KeyPattern sets S3KeyPattern field to given value.

### HasS3KeyPattern

`func (o *SelenoidOptions) HasS3KeyPattern() bool`

HasS3KeyPattern returns a boolean if a field has been set.

### GetSessionTimeout

`func (o *SelenoidOptions) GetSessionTimeout() string`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *SelenoidOptions) GetSessionTimeoutOk() (*string, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *SelenoidOptions) SetSessionTimeout(v string)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *SelenoidOptions) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.

### GetScreenResolution

`func (o *SelenoidOptions) GetScreenResolution() string`

GetScreenResolution returns the ScreenResolution field if non-nil, zero value otherwise.

### GetScreenResolutionOk

`func (o *SelenoidOptions) GetScreenResolutionOk() (*string, bool)`

GetScreenResolutionOk returns a tuple with the ScreenResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenResolution

`func (o *SelenoidOptions) SetScreenResolution(v string)`

SetScreenResolution sets ScreenResolution field to given value.

### HasScreenResolution

`func (o *SelenoidOptions) HasScreenResolution() bool`

HasScreenResolution returns a boolean if a field has been set.

### GetTimeZone

`func (o *SelenoidOptions) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *SelenoidOptions) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *SelenoidOptions) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *SelenoidOptions) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetVideoFrameRate

`func (o *SelenoidOptions) GetVideoFrameRate() int64`

GetVideoFrameRate returns the VideoFrameRate field if non-nil, zero value otherwise.

### GetVideoFrameRateOk

`func (o *SelenoidOptions) GetVideoFrameRateOk() (*int64, bool)`

GetVideoFrameRateOk returns a tuple with the VideoFrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoFrameRate

`func (o *SelenoidOptions) SetVideoFrameRate(v int64)`

SetVideoFrameRate sets VideoFrameRate field to given value.

### HasVideoFrameRate

`func (o *SelenoidOptions) HasVideoFrameRate() bool`

HasVideoFrameRate returns a boolean if a field has been set.

### GetVideoName

`func (o *SelenoidOptions) GetVideoName() string`

GetVideoName returns the VideoName field if non-nil, zero value otherwise.

### GetVideoNameOk

`func (o *SelenoidOptions) GetVideoNameOk() (*string, bool)`

GetVideoNameOk returns a tuple with the VideoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoName

`func (o *SelenoidOptions) SetVideoName(v string)`

SetVideoName sets VideoName field to given value.

### HasVideoName

`func (o *SelenoidOptions) HasVideoName() bool`

HasVideoName returns a boolean if a field has been set.

### GetVideoScreenSize

`func (o *SelenoidOptions) GetVideoScreenSize() string`

GetVideoScreenSize returns the VideoScreenSize field if non-nil, zero value otherwise.

### GetVideoScreenSizeOk

`func (o *SelenoidOptions) GetVideoScreenSizeOk() (*string, bool)`

GetVideoScreenSizeOk returns a tuple with the VideoScreenSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoScreenSize

`func (o *SelenoidOptions) SetVideoScreenSize(v string)`

SetVideoScreenSize sets VideoScreenSize field to given value.

### HasVideoScreenSize

`func (o *SelenoidOptions) HasVideoScreenSize() bool`

HasVideoScreenSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


