# PerfLoggingPrefs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableNetwork** | Pointer to **bool** |  | [optional] [default to true]
**EnablePage** | Pointer to **bool** |  | [optional] [default to true]
**TraceCategories** | Pointer to **string** |  | [optional] 
**BufferUsageReportingInterval** | Pointer to **int64** |  | [optional] [default to 1000]

## Methods

### NewPerfLoggingPrefs

`func NewPerfLoggingPrefs() *PerfLoggingPrefs`

NewPerfLoggingPrefs instantiates a new PerfLoggingPrefs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerfLoggingPrefsWithDefaults

`func NewPerfLoggingPrefsWithDefaults() *PerfLoggingPrefs`

NewPerfLoggingPrefsWithDefaults instantiates a new PerfLoggingPrefs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableNetwork

`func (o *PerfLoggingPrefs) GetEnableNetwork() bool`

GetEnableNetwork returns the EnableNetwork field if non-nil, zero value otherwise.

### GetEnableNetworkOk

`func (o *PerfLoggingPrefs) GetEnableNetworkOk() (*bool, bool)`

GetEnableNetworkOk returns a tuple with the EnableNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetwork

`func (o *PerfLoggingPrefs) SetEnableNetwork(v bool)`

SetEnableNetwork sets EnableNetwork field to given value.

### HasEnableNetwork

`func (o *PerfLoggingPrefs) HasEnableNetwork() bool`

HasEnableNetwork returns a boolean if a field has been set.

### GetEnablePage

`func (o *PerfLoggingPrefs) GetEnablePage() bool`

GetEnablePage returns the EnablePage field if non-nil, zero value otherwise.

### GetEnablePageOk

`func (o *PerfLoggingPrefs) GetEnablePageOk() (*bool, bool)`

GetEnablePageOk returns a tuple with the EnablePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePage

`func (o *PerfLoggingPrefs) SetEnablePage(v bool)`

SetEnablePage sets EnablePage field to given value.

### HasEnablePage

`func (o *PerfLoggingPrefs) HasEnablePage() bool`

HasEnablePage returns a boolean if a field has been set.

### GetTraceCategories

`func (o *PerfLoggingPrefs) GetTraceCategories() string`

GetTraceCategories returns the TraceCategories field if non-nil, zero value otherwise.

### GetTraceCategoriesOk

`func (o *PerfLoggingPrefs) GetTraceCategoriesOk() (*string, bool)`

GetTraceCategoriesOk returns a tuple with the TraceCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceCategories

`func (o *PerfLoggingPrefs) SetTraceCategories(v string)`

SetTraceCategories sets TraceCategories field to given value.

### HasTraceCategories

`func (o *PerfLoggingPrefs) HasTraceCategories() bool`

HasTraceCategories returns a boolean if a field has been set.

### GetBufferUsageReportingInterval

`func (o *PerfLoggingPrefs) GetBufferUsageReportingInterval() int64`

GetBufferUsageReportingInterval returns the BufferUsageReportingInterval field if non-nil, zero value otherwise.

### GetBufferUsageReportingIntervalOk

`func (o *PerfLoggingPrefs) GetBufferUsageReportingIntervalOk() (*int64, bool)`

GetBufferUsageReportingIntervalOk returns a tuple with the BufferUsageReportingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferUsageReportingInterval

`func (o *PerfLoggingPrefs) SetBufferUsageReportingInterval(v int64)`

SetBufferUsageReportingInterval sets BufferUsageReportingInterval field to given value.

### HasBufferUsageReportingInterval

`func (o *PerfLoggingPrefs) HasBufferUsageReportingInterval() bool`

HasBufferUsageReportingInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


