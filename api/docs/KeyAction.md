# KeyAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "pause"]
**Value** | Pointer to **string** |  | [optional] 
**Pressed** | Pointer to **bool** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 

## Methods

### NewKeyAction

`func NewKeyAction() *KeyAction`

NewKeyAction instantiates a new KeyAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyActionWithDefaults

`func NewKeyActionWithDefaults() *KeyAction`

NewKeyActionWithDefaults instantiates a new KeyAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *KeyAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeyAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeyAction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KeyAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *KeyAction) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KeyAction) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KeyAction) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *KeyAction) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPressed

`func (o *KeyAction) GetPressed() bool`

GetPressed returns the Pressed field if non-nil, zero value otherwise.

### GetPressedOk

`func (o *KeyAction) GetPressedOk() (*bool, bool)`

GetPressedOk returns a tuple with the Pressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPressed

`func (o *KeyAction) SetPressed(v bool)`

SetPressed sets Pressed field to given value.

### HasPressed

`func (o *KeyAction) HasPressed() bool`

HasPressed returns a boolean if a field has been set.

### GetDuration

`func (o *KeyAction) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *KeyAction) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *KeyAction) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *KeyAction) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


