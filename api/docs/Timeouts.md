# Timeouts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | Pointer to **NilableInt64** |  | [optional] 
**PageLoad** | Pointer to **int64** |  | [optional] 
**Implicit** | Pointer to **int64** |  | [optional] 

## Methods

### NewTimeouts

`func NewTimeouts() *Timeouts`

NewTimeouts instantiates a new Timeouts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeoutsWithDefaults

`func NewTimeoutsWithDefaults() *Timeouts`

NewTimeoutsWithDefaults instantiates a new Timeouts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScript

`func (o *Timeouts) GetScript() int64`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *Timeouts) GetScriptOk() (*int64, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *Timeouts) SetScript(v int64)`

SetScript sets Script field to given value.

### HasScript

`func (o *Timeouts) HasScript() bool`

HasScript returns a boolean if a field has been set.

### SetScriptNil

`func (o *Timeouts) SetScriptNil(b bool)`

 SetScriptNil sets the value for Script to be an explicit nil

### UnsetScript
`func (o *Timeouts) UnsetScript()`

UnsetScript ensures that no value is present for Script, not even an explicit nil
### GetPageLoad

`func (o *Timeouts) GetPageLoad() int64`

GetPageLoad returns the PageLoad field if non-nil, zero value otherwise.

### GetPageLoadOk

`func (o *Timeouts) GetPageLoadOk() (*int64, bool)`

GetPageLoadOk returns a tuple with the PageLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageLoad

`func (o *Timeouts) SetPageLoad(v int64)`

SetPageLoad sets PageLoad field to given value.

### HasPageLoad

`func (o *Timeouts) HasPageLoad() bool`

HasPageLoad returns a boolean if a field has been set.

### GetImplicit

`func (o *Timeouts) GetImplicit() int64`

GetImplicit returns the Implicit field if non-nil, zero value otherwise.

### GetImplicitOk

`func (o *Timeouts) GetImplicitOk() (*int64, bool)`

GetImplicitOk returns a tuple with the Implicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicit

`func (o *Timeouts) SetImplicit(v int64)`

SetImplicit sets Implicit field to given value.

### HasImplicit

`func (o *Timeouts) HasImplicit() bool`

HasImplicit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


