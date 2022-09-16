# AddonInstallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addon** | **string** |  | 
**Temporary** | Pointer to **bool** | temporarily installed extension will be removed at the end of browser session | [optional] [default to true]

## Methods

### NewAddonInstallRequest

`func NewAddonInstallRequest(addon string, ) *AddonInstallRequest`

NewAddonInstallRequest instantiates a new AddonInstallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonInstallRequestWithDefaults

`func NewAddonInstallRequestWithDefaults() *AddonInstallRequest`

NewAddonInstallRequestWithDefaults instantiates a new AddonInstallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddon

`func (o *AddonInstallRequest) GetAddon() string`

GetAddon returns the Addon field if non-nil, zero value otherwise.

### GetAddonOk

`func (o *AddonInstallRequest) GetAddonOk() (*string, bool)`

GetAddonOk returns a tuple with the Addon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddon

`func (o *AddonInstallRequest) SetAddon(v string)`

SetAddon sets Addon field to given value.


### GetTemporary

`func (o *AddonInstallRequest) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *AddonInstallRequest) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *AddonInstallRequest) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.

### HasTemporary

`func (o *AddonInstallRequest) HasTemporary() bool`

HasTemporary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


