# ChromeOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **[]string** |  | [optional] 
**Binary** | Pointer to **string** |  | [optional] 
**DebuggerAddress** | Pointer to **string** |  | [optional] 
**Detach** | Pointer to **bool** |  | [optional] 
**ExcludeSwitches** | Pointer to **[]string** |  | [optional] 
**Extensions** | Pointer to **[]string** |  | [optional] 
**LocalState** | Pointer to [**map[string]PreferenceValue**](PreferenceValue.md) |  | [optional] 
**MinidumpPath** | Pointer to **string** |  | [optional] 
**MobileEmulation** | Pointer to [**MobileEmulation**](MobileEmulation.md) |  | [optional] 
**PerfLoggingPrefs** | Pointer to [**PerfLoggingPrefs**](PerfLoggingPrefs.md) |  | [optional] 
**Prefs** | Pointer to [**[]PreferenceValue**](PreferenceValue.md) |  | [optional] 
**WindowTypes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewChromeOptions

`func NewChromeOptions() *ChromeOptions`

NewChromeOptions instantiates a new ChromeOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChromeOptionsWithDefaults

`func NewChromeOptionsWithDefaults() *ChromeOptions`

NewChromeOptionsWithDefaults instantiates a new ChromeOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *ChromeOptions) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ChromeOptions) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ChromeOptions) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *ChromeOptions) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetBinary

`func (o *ChromeOptions) GetBinary() string`

GetBinary returns the Binary field if non-nil, zero value otherwise.

### GetBinaryOk

`func (o *ChromeOptions) GetBinaryOk() (*string, bool)`

GetBinaryOk returns a tuple with the Binary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinary

`func (o *ChromeOptions) SetBinary(v string)`

SetBinary sets Binary field to given value.

### HasBinary

`func (o *ChromeOptions) HasBinary() bool`

HasBinary returns a boolean if a field has been set.

### GetDebuggerAddress

`func (o *ChromeOptions) GetDebuggerAddress() string`

GetDebuggerAddress returns the DebuggerAddress field if non-nil, zero value otherwise.

### GetDebuggerAddressOk

`func (o *ChromeOptions) GetDebuggerAddressOk() (*string, bool)`

GetDebuggerAddressOk returns a tuple with the DebuggerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebuggerAddress

`func (o *ChromeOptions) SetDebuggerAddress(v string)`

SetDebuggerAddress sets DebuggerAddress field to given value.

### HasDebuggerAddress

`func (o *ChromeOptions) HasDebuggerAddress() bool`

HasDebuggerAddress returns a boolean if a field has been set.

### GetDetach

`func (o *ChromeOptions) GetDetach() bool`

GetDetach returns the Detach field if non-nil, zero value otherwise.

### GetDetachOk

`func (o *ChromeOptions) GetDetachOk() (*bool, bool)`

GetDetachOk returns a tuple with the Detach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetach

`func (o *ChromeOptions) SetDetach(v bool)`

SetDetach sets Detach field to given value.

### HasDetach

`func (o *ChromeOptions) HasDetach() bool`

HasDetach returns a boolean if a field has been set.

### GetExcludeSwitches

`func (o *ChromeOptions) GetExcludeSwitches() []string`

GetExcludeSwitches returns the ExcludeSwitches field if non-nil, zero value otherwise.

### GetExcludeSwitchesOk

`func (o *ChromeOptions) GetExcludeSwitchesOk() (*[]string, bool)`

GetExcludeSwitchesOk returns a tuple with the ExcludeSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSwitches

`func (o *ChromeOptions) SetExcludeSwitches(v []string)`

SetExcludeSwitches sets ExcludeSwitches field to given value.

### HasExcludeSwitches

`func (o *ChromeOptions) HasExcludeSwitches() bool`

HasExcludeSwitches returns a boolean if a field has been set.

### GetExtensions

`func (o *ChromeOptions) GetExtensions() []string`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *ChromeOptions) GetExtensionsOk() (*[]string, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *ChromeOptions) SetExtensions(v []string)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *ChromeOptions) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetLocalState

`func (o *ChromeOptions) GetLocalState() map[string]PreferenceValue`

GetLocalState returns the LocalState field if non-nil, zero value otherwise.

### GetLocalStateOk

`func (o *ChromeOptions) GetLocalStateOk() (*map[string]PreferenceValue, bool)`

GetLocalStateOk returns a tuple with the LocalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalState

`func (o *ChromeOptions) SetLocalState(v map[string]PreferenceValue)`

SetLocalState sets LocalState field to given value.

### HasLocalState

`func (o *ChromeOptions) HasLocalState() bool`

HasLocalState returns a boolean if a field has been set.

### GetMinidumpPath

`func (o *ChromeOptions) GetMinidumpPath() string`

GetMinidumpPath returns the MinidumpPath field if non-nil, zero value otherwise.

### GetMinidumpPathOk

`func (o *ChromeOptions) GetMinidumpPathOk() (*string, bool)`

GetMinidumpPathOk returns a tuple with the MinidumpPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinidumpPath

`func (o *ChromeOptions) SetMinidumpPath(v string)`

SetMinidumpPath sets MinidumpPath field to given value.

### HasMinidumpPath

`func (o *ChromeOptions) HasMinidumpPath() bool`

HasMinidumpPath returns a boolean if a field has been set.

### GetMobileEmulation

`func (o *ChromeOptions) GetMobileEmulation() MobileEmulation`

GetMobileEmulation returns the MobileEmulation field if non-nil, zero value otherwise.

### GetMobileEmulationOk

`func (o *ChromeOptions) GetMobileEmulationOk() (*MobileEmulation, bool)`

GetMobileEmulationOk returns a tuple with the MobileEmulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileEmulation

`func (o *ChromeOptions) SetMobileEmulation(v MobileEmulation)`

SetMobileEmulation sets MobileEmulation field to given value.

### HasMobileEmulation

`func (o *ChromeOptions) HasMobileEmulation() bool`

HasMobileEmulation returns a boolean if a field has been set.

### GetPerfLoggingPrefs

`func (o *ChromeOptions) GetPerfLoggingPrefs() PerfLoggingPrefs`

GetPerfLoggingPrefs returns the PerfLoggingPrefs field if non-nil, zero value otherwise.

### GetPerfLoggingPrefsOk

`func (o *ChromeOptions) GetPerfLoggingPrefsOk() (*PerfLoggingPrefs, bool)`

GetPerfLoggingPrefsOk returns a tuple with the PerfLoggingPrefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfLoggingPrefs

`func (o *ChromeOptions) SetPerfLoggingPrefs(v PerfLoggingPrefs)`

SetPerfLoggingPrefs sets PerfLoggingPrefs field to given value.

### HasPerfLoggingPrefs

`func (o *ChromeOptions) HasPerfLoggingPrefs() bool`

HasPerfLoggingPrefs returns a boolean if a field has been set.

### GetPrefs

`func (o *ChromeOptions) GetPrefs() []PreferenceValue`

GetPrefs returns the Prefs field if non-nil, zero value otherwise.

### GetPrefsOk

`func (o *ChromeOptions) GetPrefsOk() (*[]PreferenceValue, bool)`

GetPrefsOk returns a tuple with the Prefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefs

`func (o *ChromeOptions) SetPrefs(v []PreferenceValue)`

SetPrefs sets Prefs field to given value.

### HasPrefs

`func (o *ChromeOptions) HasPrefs() bool`

HasPrefs returns a boolean if a field has been set.

### GetWindowTypes

`func (o *ChromeOptions) GetWindowTypes() []string`

GetWindowTypes returns the WindowTypes field if non-nil, zero value otherwise.

### GetWindowTypesOk

`func (o *ChromeOptions) GetWindowTypesOk() (*[]string, bool)`

GetWindowTypesOk returns a tuple with the WindowTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowTypes

`func (o *ChromeOptions) SetWindowTypes(v []string)`

SetWindowTypes sets WindowTypes field to given value.

### HasWindowTypes

`func (o *ChromeOptions) HasWindowTypes() bool`

HasWindowTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


