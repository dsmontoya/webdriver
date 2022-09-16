# PointerAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "pause"]
**Button** | Pointer to **int32** | Possible values: &#x60;0&#x60; - left button, &#x60;1&#x60; - middle button, &#x60;2&#x60; - right button | [optional] 
**Pressed** | Pointer to **int32** |  | [optional] 
**X** | Pointer to **int32** |  | [optional] 
**Y** | Pointer to **int32** |  | [optional] 
**Width** | Pointer to **int32** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**Pressure** | Pointer to **int32** |  | [optional] 
**TangentialPressure** | Pointer to **int32** |  | [optional] 
**TiltX** | Pointer to **int32** |  | [optional] 
**TiltY** | Pointer to **int32** |  | [optional] 
**Twist** | Pointer to **int32** |  | [optional] 
**AltitudeAngle** | Pointer to **float32** |  | [optional] 
**AsimuthAngle** | Pointer to **float32** |  | [optional] 
**Origin** | Pointer to [**ActionOrigin**](ActionOrigin.md) |  | [optional] 

## Methods

### NewPointerAction

`func NewPointerAction() *PointerAction`

NewPointerAction instantiates a new PointerAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointerActionWithDefaults

`func NewPointerActionWithDefaults() *PointerAction`

NewPointerActionWithDefaults instantiates a new PointerAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PointerAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PointerAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PointerAction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PointerAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetButton

`func (o *PointerAction) GetButton() int32`

GetButton returns the Button field if non-nil, zero value otherwise.

### GetButtonOk

`func (o *PointerAction) GetButtonOk() (*int32, bool)`

GetButtonOk returns a tuple with the Button field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButton

`func (o *PointerAction) SetButton(v int32)`

SetButton sets Button field to given value.

### HasButton

`func (o *PointerAction) HasButton() bool`

HasButton returns a boolean if a field has been set.

### GetPressed

`func (o *PointerAction) GetPressed() int32`

GetPressed returns the Pressed field if non-nil, zero value otherwise.

### GetPressedOk

`func (o *PointerAction) GetPressedOk() (*int32, bool)`

GetPressedOk returns a tuple with the Pressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPressed

`func (o *PointerAction) SetPressed(v int32)`

SetPressed sets Pressed field to given value.

### HasPressed

`func (o *PointerAction) HasPressed() bool`

HasPressed returns a boolean if a field has been set.

### GetX

`func (o *PointerAction) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *PointerAction) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *PointerAction) SetX(v int32)`

SetX sets X field to given value.

### HasX

`func (o *PointerAction) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *PointerAction) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *PointerAction) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *PointerAction) SetY(v int32)`

SetY sets Y field to given value.

### HasY

`func (o *PointerAction) HasY() bool`

HasY returns a boolean if a field has been set.

### GetWidth

`func (o *PointerAction) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *PointerAction) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *PointerAction) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *PointerAction) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *PointerAction) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *PointerAction) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *PointerAction) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *PointerAction) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetPressure

`func (o *PointerAction) GetPressure() int32`

GetPressure returns the Pressure field if non-nil, zero value otherwise.

### GetPressureOk

`func (o *PointerAction) GetPressureOk() (*int32, bool)`

GetPressureOk returns a tuple with the Pressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPressure

`func (o *PointerAction) SetPressure(v int32)`

SetPressure sets Pressure field to given value.

### HasPressure

`func (o *PointerAction) HasPressure() bool`

HasPressure returns a boolean if a field has been set.

### GetTangentialPressure

`func (o *PointerAction) GetTangentialPressure() int32`

GetTangentialPressure returns the TangentialPressure field if non-nil, zero value otherwise.

### GetTangentialPressureOk

`func (o *PointerAction) GetTangentialPressureOk() (*int32, bool)`

GetTangentialPressureOk returns a tuple with the TangentialPressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTangentialPressure

`func (o *PointerAction) SetTangentialPressure(v int32)`

SetTangentialPressure sets TangentialPressure field to given value.

### HasTangentialPressure

`func (o *PointerAction) HasTangentialPressure() bool`

HasTangentialPressure returns a boolean if a field has been set.

### GetTiltX

`func (o *PointerAction) GetTiltX() int32`

GetTiltX returns the TiltX field if non-nil, zero value otherwise.

### GetTiltXOk

`func (o *PointerAction) GetTiltXOk() (*int32, bool)`

GetTiltXOk returns a tuple with the TiltX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiltX

`func (o *PointerAction) SetTiltX(v int32)`

SetTiltX sets TiltX field to given value.

### HasTiltX

`func (o *PointerAction) HasTiltX() bool`

HasTiltX returns a boolean if a field has been set.

### GetTiltY

`func (o *PointerAction) GetTiltY() int32`

GetTiltY returns the TiltY field if non-nil, zero value otherwise.

### GetTiltYOk

`func (o *PointerAction) GetTiltYOk() (*int32, bool)`

GetTiltYOk returns a tuple with the TiltY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiltY

`func (o *PointerAction) SetTiltY(v int32)`

SetTiltY sets TiltY field to given value.

### HasTiltY

`func (o *PointerAction) HasTiltY() bool`

HasTiltY returns a boolean if a field has been set.

### GetTwist

`func (o *PointerAction) GetTwist() int32`

GetTwist returns the Twist field if non-nil, zero value otherwise.

### GetTwistOk

`func (o *PointerAction) GetTwistOk() (*int32, bool)`

GetTwistOk returns a tuple with the Twist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwist

`func (o *PointerAction) SetTwist(v int32)`

SetTwist sets Twist field to given value.

### HasTwist

`func (o *PointerAction) HasTwist() bool`

HasTwist returns a boolean if a field has been set.

### GetAltitudeAngle

`func (o *PointerAction) GetAltitudeAngle() float32`

GetAltitudeAngle returns the AltitudeAngle field if non-nil, zero value otherwise.

### GetAltitudeAngleOk

`func (o *PointerAction) GetAltitudeAngleOk() (*float32, bool)`

GetAltitudeAngleOk returns a tuple with the AltitudeAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitudeAngle

`func (o *PointerAction) SetAltitudeAngle(v float32)`

SetAltitudeAngle sets AltitudeAngle field to given value.

### HasAltitudeAngle

`func (o *PointerAction) HasAltitudeAngle() bool`

HasAltitudeAngle returns a boolean if a field has been set.

### GetAsimuthAngle

`func (o *PointerAction) GetAsimuthAngle() float32`

GetAsimuthAngle returns the AsimuthAngle field if non-nil, zero value otherwise.

### GetAsimuthAngleOk

`func (o *PointerAction) GetAsimuthAngleOk() (*float32, bool)`

GetAsimuthAngleOk returns a tuple with the AsimuthAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsimuthAngle

`func (o *PointerAction) SetAsimuthAngle(v float32)`

SetAsimuthAngle sets AsimuthAngle field to given value.

### HasAsimuthAngle

`func (o *PointerAction) HasAsimuthAngle() bool`

HasAsimuthAngle returns a boolean if a field has been set.

### GetOrigin

`func (o *PointerAction) GetOrigin() ActionOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *PointerAction) GetOriginOk() (*ActionOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *PointerAction) SetOrigin(v ActionOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *PointerAction) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


