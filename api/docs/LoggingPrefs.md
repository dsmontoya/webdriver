# LoggingPrefs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Browser** | Pointer to [**LogLevel**](LogLevel.md) |  | [optional] 
**Client** | Pointer to [**LogLevel**](LogLevel.md) |  | [optional] 
**Driver** | Pointer to [**LogLevel**](LogLevel.md) |  | [optional] 
**Performance** | Pointer to [**LogLevel**](LogLevel.md) |  | [optional] 
**Profiler** | Pointer to [**LogLevel**](LogLevel.md) |  | [optional] 
**Server** | Pointer to [**LogLevel**](LogLevel.md) |  | [optional] 

## Methods

### NewLoggingPrefs

`func NewLoggingPrefs() *LoggingPrefs`

NewLoggingPrefs instantiates a new LoggingPrefs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingPrefsWithDefaults

`func NewLoggingPrefsWithDefaults() *LoggingPrefs`

NewLoggingPrefsWithDefaults instantiates a new LoggingPrefs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrowser

`func (o *LoggingPrefs) GetBrowser() LogLevel`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *LoggingPrefs) GetBrowserOk() (*LogLevel, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *LoggingPrefs) SetBrowser(v LogLevel)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *LoggingPrefs) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### GetClient

`func (o *LoggingPrefs) GetClient() LogLevel`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *LoggingPrefs) GetClientOk() (*LogLevel, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *LoggingPrefs) SetClient(v LogLevel)`

SetClient sets Client field to given value.

### HasClient

`func (o *LoggingPrefs) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetDriver

`func (o *LoggingPrefs) GetDriver() LogLevel`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *LoggingPrefs) GetDriverOk() (*LogLevel, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *LoggingPrefs) SetDriver(v LogLevel)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *LoggingPrefs) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetPerformance

`func (o *LoggingPrefs) GetPerformance() LogLevel`

GetPerformance returns the Performance field if non-nil, zero value otherwise.

### GetPerformanceOk

`func (o *LoggingPrefs) GetPerformanceOk() (*LogLevel, bool)`

GetPerformanceOk returns a tuple with the Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformance

`func (o *LoggingPrefs) SetPerformance(v LogLevel)`

SetPerformance sets Performance field to given value.

### HasPerformance

`func (o *LoggingPrefs) HasPerformance() bool`

HasPerformance returns a boolean if a field has been set.

### GetProfiler

`func (o *LoggingPrefs) GetProfiler() LogLevel`

GetProfiler returns the Profiler field if non-nil, zero value otherwise.

### GetProfilerOk

`func (o *LoggingPrefs) GetProfilerOk() (*LogLevel, bool)`

GetProfilerOk returns a tuple with the Profiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiler

`func (o *LoggingPrefs) SetProfiler(v LogLevel)`

SetProfiler sets Profiler field to given value.

### HasProfiler

`func (o *LoggingPrefs) HasProfiler() bool`

HasProfiler returns a boolean if a field has been set.

### GetServer

`func (o *LoggingPrefs) GetServer() LogLevel`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *LoggingPrefs) GetServerOk() (*LogLevel, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *LoggingPrefs) SetServer(v LogLevel)`

SetServer sets Server field to given value.

### HasServer

`func (o *LoggingPrefs) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


