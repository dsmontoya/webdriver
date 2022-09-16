# Action

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "pause"]
**Duration** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Pressed** | Pointer to **int32** |  | [optional] 
**Button** | Pointer to **int32** | Possible values: &#x60;0&#x60; - left button, &#x60;1&#x60; - middle button, &#x60;2&#x60; - right button | [optional] 
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
**DeltaX** | Pointer to **int32** |  | [optional] 
**DeltaY** | Pointer to **int32** |  | [optional] 

## Methods

### NewAction

`func NewAction() *Action`

NewAction instantiates a new Action object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionWithDefaults

`func NewActionWithDefaults() *Action`

NewActionWithDefaults instantiates a new Action object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Action) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Action) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Action) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Action) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDuration

`func (o *Action) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Action) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Action) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Action) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetValue

`func (o *Action) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Action) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Action) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Action) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPressed

`func (o *Action) GetPressed() int32`

GetPressed returns the Pressed field if non-nil, zero value otherwise.

### GetPressedOk

`func (o *Action) GetPressedOk() (*int32, bool)`

GetPressedOk returns a tuple with the Pressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPressed

`func (o *Action) SetPressed(v int32)`

SetPressed sets Pressed field to given value.

### HasPressed

`func (o *Action) HasPressed() bool`

HasPressed returns a boolean if a field has been set.

### GetButton

`func (o *Action) GetButton() int32`

GetButton returns the Button field if non-nil, zero value otherwise.

### GetButtonOk

`func (o *Action) GetButtonOk() (*int32, bool)`

GetButtonOk returns a tuple with the Button field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButton

`func (o *Action) SetButton(v int32)`

SetButton sets Button field to given value.

### HasButton

`func (o *Action) HasButton() bool`

HasButton returns a boolean if a field has been set.

### GetX

`func (o *Action) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *Action) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *Action) SetX(v int32)`

SetX sets X field to given value.

### HasX

`func (o *Action) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *Action) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *Action) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *Action) SetY(v int32)`

SetY sets Y field to given value.

### HasY

`func (o *Action) HasY() bool`

HasY returns a boolean if a field has been set.

### GetWidth

`func (o *Action) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Action) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Action) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *Action) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *Action) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Action) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Action) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *Action) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetPressure

`func (o *Action) GetPressure() int32`

GetPressure returns the Pressure field if non-nil, zero value otherwise.

### GetPressureOk

`func (o *Action) GetPressureOk() (*int32, bool)`

GetPressureOk returns a tuple with the Pressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPressure

`func (o *Action) SetPressure(v int32)`

SetPressure sets Pressure field to given value.

### HasPressure

`func (o *Action) HasPressure() bool`

HasPressure returns a boolean if a field has been set.

### GetTangentialPressure

`func (o *Action) GetTangentialPressure() int32`

GetTangentialPressure returns the TangentialPressure field if non-nil, zero value otherwise.

### GetTangentialPressureOk

`func (o *Action) GetTangentialPressureOk() (*int32, bool)`

GetTangentialPressureOk returns a tuple with the TangentialPressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTangentialPressure

`func (o *Action) SetTangentialPressure(v int32)`

SetTangentialPressure sets TangentialPressure field to given value.

### HasTangentialPressure

`func (o *Action) HasTangentialPressure() bool`

HasTangentialPressure returns a boolean if a field has been set.

### GetTiltX

`func (o *Action) GetTiltX() int32`

GetTiltX returns the TiltX field if non-nil, zero value otherwise.

### GetTiltXOk

`func (o *Action) GetTiltXOk() (*int32, bool)`

GetTiltXOk returns a tuple with the TiltX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiltX

`func (o *Action) SetTiltX(v int32)`

SetTiltX sets TiltX field to given value.

### HasTiltX

`func (o *Action) HasTiltX() bool`

HasTiltX returns a boolean if a field has been set.

### GetTiltY

`func (o *Action) GetTiltY() int32`

GetTiltY returns the TiltY field if non-nil, zero value otherwise.

### GetTiltYOk

`func (o *Action) GetTiltYOk() (*int32, bool)`

GetTiltYOk returns a tuple with the TiltY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiltY

`func (o *Action) SetTiltY(v int32)`

SetTiltY sets TiltY field to given value.

### HasTiltY

`func (o *Action) HasTiltY() bool`

HasTiltY returns a boolean if a field has been set.

### GetTwist

`func (o *Action) GetTwist() int32`

GetTwist returns the Twist field if non-nil, zero value otherwise.

### GetTwistOk

`func (o *Action) GetTwistOk() (*int32, bool)`

GetTwistOk returns a tuple with the Twist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwist

`func (o *Action) SetTwist(v int32)`

SetTwist sets Twist field to given value.

### HasTwist

`func (o *Action) HasTwist() bool`

HasTwist returns a boolean if a field has been set.

### GetAltitudeAngle

`func (o *Action) GetAltitudeAngle() float32`

GetAltitudeAngle returns the AltitudeAngle field if non-nil, zero value otherwise.

### GetAltitudeAngleOk

`func (o *Action) GetAltitudeAngleOk() (*float32, bool)`

GetAltitudeAngleOk returns a tuple with the AltitudeAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitudeAngle

`func (o *Action) SetAltitudeAngle(v float32)`

SetAltitudeAngle sets AltitudeAngle field to given value.

### HasAltitudeAngle

`func (o *Action) HasAltitudeAngle() bool`

HasAltitudeAngle returns a boolean if a field has been set.

### GetAsimuthAngle

`func (o *Action) GetAsimuthAngle() float32`

GetAsimuthAngle returns the AsimuthAngle field if non-nil, zero value otherwise.

### GetAsimuthAngleOk

`func (o *Action) GetAsimuthAngleOk() (*float32, bool)`

GetAsimuthAngleOk returns a tuple with the AsimuthAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsimuthAngle

`func (o *Action) SetAsimuthAngle(v float32)`

SetAsimuthAngle sets AsimuthAngle field to given value.

### HasAsimuthAngle

`func (o *Action) HasAsimuthAngle() bool`

HasAsimuthAngle returns a boolean if a field has been set.

### GetOrigin

`func (o *Action) GetOrigin() ActionOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Action) GetOriginOk() (*ActionOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Action) SetOrigin(v ActionOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *Action) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetDeltaX

`func (o *Action) GetDeltaX() int32`

GetDeltaX returns the DeltaX field if non-nil, zero value otherwise.

### GetDeltaXOk

`func (o *Action) GetDeltaXOk() (*int32, bool)`

GetDeltaXOk returns a tuple with the DeltaX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaX

`func (o *Action) SetDeltaX(v int32)`

SetDeltaX sets DeltaX field to given value.

### HasDeltaX

`func (o *Action) HasDeltaX() bool`

HasDeltaX returns a boolean if a field has been set.

### GetDeltaY

`func (o *Action) GetDeltaY() int32`

GetDeltaY returns the DeltaY field if non-nil, zero value otherwise.

### GetDeltaYOk

`func (o *Action) GetDeltaYOk() (*int32, bool)`

GetDeltaYOk returns a tuple with the DeltaY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaY

`func (o *Action) SetDeltaY(v int32)`

SetDeltaY sets DeltaY field to given value.

### HasDeltaY

`func (o *Action) HasDeltaY() bool`

HasDeltaY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


