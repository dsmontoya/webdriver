# FirefoxOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AndroidPackage** | Pointer to **string** |  | [optional] 
**AndroidActivity** | Pointer to **string** |  | [optional] 
**AndroidDeviceSerial** | Pointer to **string** |  | [optional] 
**AndroidIntentArguments** | Pointer to **[]string** |  | [optional] 
**Args** | Pointer to **[]string** |  | [optional] 
**Binary** | Pointer to **string** |  | [optional] 
**Env** | Pointer to **map[string]string** |  | [optional] 
**Log** | Pointer to [**FirefoxOptionsLog**](FirefoxOptionsLog.md) |  | [optional] 
**Prefs** | Pointer to [**map[string]PreferenceValue**](PreferenceValue.md) |  | [optional] 
**Profile** | Pointer to **string** |  | [optional] 

## Methods

### NewFirefoxOptions

`func NewFirefoxOptions() *FirefoxOptions`

NewFirefoxOptions instantiates a new FirefoxOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirefoxOptionsWithDefaults

`func NewFirefoxOptionsWithDefaults() *FirefoxOptions`

NewFirefoxOptionsWithDefaults instantiates a new FirefoxOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAndroidPackage

`func (o *FirefoxOptions) GetAndroidPackage() string`

GetAndroidPackage returns the AndroidPackage field if non-nil, zero value otherwise.

### GetAndroidPackageOk

`func (o *FirefoxOptions) GetAndroidPackageOk() (*string, bool)`

GetAndroidPackageOk returns a tuple with the AndroidPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidPackage

`func (o *FirefoxOptions) SetAndroidPackage(v string)`

SetAndroidPackage sets AndroidPackage field to given value.

### HasAndroidPackage

`func (o *FirefoxOptions) HasAndroidPackage() bool`

HasAndroidPackage returns a boolean if a field has been set.

### GetAndroidActivity

`func (o *FirefoxOptions) GetAndroidActivity() string`

GetAndroidActivity returns the AndroidActivity field if non-nil, zero value otherwise.

### GetAndroidActivityOk

`func (o *FirefoxOptions) GetAndroidActivityOk() (*string, bool)`

GetAndroidActivityOk returns a tuple with the AndroidActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidActivity

`func (o *FirefoxOptions) SetAndroidActivity(v string)`

SetAndroidActivity sets AndroidActivity field to given value.

### HasAndroidActivity

`func (o *FirefoxOptions) HasAndroidActivity() bool`

HasAndroidActivity returns a boolean if a field has been set.

### GetAndroidDeviceSerial

`func (o *FirefoxOptions) GetAndroidDeviceSerial() string`

GetAndroidDeviceSerial returns the AndroidDeviceSerial field if non-nil, zero value otherwise.

### GetAndroidDeviceSerialOk

`func (o *FirefoxOptions) GetAndroidDeviceSerialOk() (*string, bool)`

GetAndroidDeviceSerialOk returns a tuple with the AndroidDeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidDeviceSerial

`func (o *FirefoxOptions) SetAndroidDeviceSerial(v string)`

SetAndroidDeviceSerial sets AndroidDeviceSerial field to given value.

### HasAndroidDeviceSerial

`func (o *FirefoxOptions) HasAndroidDeviceSerial() bool`

HasAndroidDeviceSerial returns a boolean if a field has been set.

### GetAndroidIntentArguments

`func (o *FirefoxOptions) GetAndroidIntentArguments() []string`

GetAndroidIntentArguments returns the AndroidIntentArguments field if non-nil, zero value otherwise.

### GetAndroidIntentArgumentsOk

`func (o *FirefoxOptions) GetAndroidIntentArgumentsOk() (*[]string, bool)`

GetAndroidIntentArgumentsOk returns a tuple with the AndroidIntentArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidIntentArguments

`func (o *FirefoxOptions) SetAndroidIntentArguments(v []string)`

SetAndroidIntentArguments sets AndroidIntentArguments field to given value.

### HasAndroidIntentArguments

`func (o *FirefoxOptions) HasAndroidIntentArguments() bool`

HasAndroidIntentArguments returns a boolean if a field has been set.

### GetArgs

`func (o *FirefoxOptions) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *FirefoxOptions) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *FirefoxOptions) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *FirefoxOptions) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetBinary

`func (o *FirefoxOptions) GetBinary() string`

GetBinary returns the Binary field if non-nil, zero value otherwise.

### GetBinaryOk

`func (o *FirefoxOptions) GetBinaryOk() (*string, bool)`

GetBinaryOk returns a tuple with the Binary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinary

`func (o *FirefoxOptions) SetBinary(v string)`

SetBinary sets Binary field to given value.

### HasBinary

`func (o *FirefoxOptions) HasBinary() bool`

HasBinary returns a boolean if a field has been set.

### GetEnv

`func (o *FirefoxOptions) GetEnv() map[string]string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *FirefoxOptions) GetEnvOk() (*map[string]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *FirefoxOptions) SetEnv(v map[string]string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *FirefoxOptions) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetLog

`func (o *FirefoxOptions) GetLog() FirefoxOptionsLog`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *FirefoxOptions) GetLogOk() (*FirefoxOptionsLog, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *FirefoxOptions) SetLog(v FirefoxOptionsLog)`

SetLog sets Log field to given value.

### HasLog

`func (o *FirefoxOptions) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetPrefs

`func (o *FirefoxOptions) GetPrefs() map[string]PreferenceValue`

GetPrefs returns the Prefs field if non-nil, zero value otherwise.

### GetPrefsOk

`func (o *FirefoxOptions) GetPrefsOk() (*map[string]PreferenceValue, bool)`

GetPrefsOk returns a tuple with the Prefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefs

`func (o *FirefoxOptions) SetPrefs(v map[string]PreferenceValue)`

SetPrefs sets Prefs field to given value.

### HasPrefs

`func (o *FirefoxOptions) HasPrefs() bool`

HasPrefs returns a boolean if a field has been set.

### GetProfile

`func (o *FirefoxOptions) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *FirefoxOptions) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *FirefoxOptions) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *FirefoxOptions) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


