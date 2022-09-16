# WheelAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "pause"]
**X** | Pointer to **int32** |  | [optional] 
**Y** | Pointer to **int32** |  | [optional] 
**DeltaX** | Pointer to **int32** |  | [optional] 
**DeltaY** | Pointer to **int32** |  | [optional] 
**Origin** | Pointer to [**ActionOrigin**](ActionOrigin.md) |  | [optional] 

## Methods

### NewWheelAction

`func NewWheelAction() *WheelAction`

NewWheelAction instantiates a new WheelAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWheelActionWithDefaults

`func NewWheelActionWithDefaults() *WheelAction`

NewWheelActionWithDefaults instantiates a new WheelAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WheelAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WheelAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WheelAction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WheelAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetX

`func (o *WheelAction) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *WheelAction) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *WheelAction) SetX(v int32)`

SetX sets X field to given value.

### HasX

`func (o *WheelAction) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *WheelAction) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *WheelAction) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *WheelAction) SetY(v int32)`

SetY sets Y field to given value.

### HasY

`func (o *WheelAction) HasY() bool`

HasY returns a boolean if a field has been set.

### GetDeltaX

`func (o *WheelAction) GetDeltaX() int32`

GetDeltaX returns the DeltaX field if non-nil, zero value otherwise.

### GetDeltaXOk

`func (o *WheelAction) GetDeltaXOk() (*int32, bool)`

GetDeltaXOk returns a tuple with the DeltaX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaX

`func (o *WheelAction) SetDeltaX(v int32)`

SetDeltaX sets DeltaX field to given value.

### HasDeltaX

`func (o *WheelAction) HasDeltaX() bool`

HasDeltaX returns a boolean if a field has been set.

### GetDeltaY

`func (o *WheelAction) GetDeltaY() int32`

GetDeltaY returns the DeltaY field if non-nil, zero value otherwise.

### GetDeltaYOk

`func (o *WheelAction) GetDeltaYOk() (*int32, bool)`

GetDeltaYOk returns a tuple with the DeltaY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaY

`func (o *WheelAction) SetDeltaY(v int32)`

SetDeltaY sets DeltaY field to given value.

### HasDeltaY

`func (o *WheelAction) HasDeltaY() bool`

HasDeltaY returns a boolean if a field has been set.

### GetOrigin

`func (o *WheelAction) GetOrigin() ActionOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *WheelAction) GetOriginOk() (*ActionOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *WheelAction) SetOrigin(v ActionOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *WheelAction) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


