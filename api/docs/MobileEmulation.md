# MobileEmulation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceName** | Pointer to **string** |  | [optional] 
**DeviceMetrics** | Pointer to [**MobileEmulationDeviceMetrics**](MobileEmulationDeviceMetrics.md) |  | [optional] 
**UserAgent** | Pointer to **string** |  | [optional] 

## Methods

### NewMobileEmulation

`func NewMobileEmulation() *MobileEmulation`

NewMobileEmulation instantiates a new MobileEmulation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileEmulationWithDefaults

`func NewMobileEmulationWithDefaults() *MobileEmulation`

NewMobileEmulationWithDefaults instantiates a new MobileEmulation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceName

`func (o *MobileEmulation) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *MobileEmulation) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *MobileEmulation) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *MobileEmulation) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceMetrics

`func (o *MobileEmulation) GetDeviceMetrics() MobileEmulationDeviceMetrics`

GetDeviceMetrics returns the DeviceMetrics field if non-nil, zero value otherwise.

### GetDeviceMetricsOk

`func (o *MobileEmulation) GetDeviceMetricsOk() (*MobileEmulationDeviceMetrics, bool)`

GetDeviceMetricsOk returns a tuple with the DeviceMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMetrics

`func (o *MobileEmulation) SetDeviceMetrics(v MobileEmulationDeviceMetrics)`

SetDeviceMetrics sets DeviceMetrics field to given value.

### HasDeviceMetrics

`func (o *MobileEmulation) HasDeviceMetrics() bool`

HasDeviceMetrics returns a boolean if a field has been set.

### GetUserAgent

`func (o *MobileEmulation) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *MobileEmulation) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *MobileEmulation) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *MobileEmulation) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


