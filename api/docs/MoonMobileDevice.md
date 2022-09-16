# MoonMobileDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceName** | **string** |  | 
**Orientation** | Pointer to **string** |  | [optional] [default to "portrait"]

## Methods

### NewMoonMobileDevice

`func NewMoonMobileDevice(deviceName string, ) *MoonMobileDevice`

NewMoonMobileDevice instantiates a new MoonMobileDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoonMobileDeviceWithDefaults

`func NewMoonMobileDeviceWithDefaults() *MoonMobileDevice`

NewMoonMobileDeviceWithDefaults instantiates a new MoonMobileDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceName

`func (o *MoonMobileDevice) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *MoonMobileDevice) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *MoonMobileDevice) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetOrientation

`func (o *MoonMobileDevice) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *MoonMobileDevice) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *MoonMobileDevice) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *MoonMobileDevice) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


