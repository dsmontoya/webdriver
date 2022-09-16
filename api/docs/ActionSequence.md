# ActionSequence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**ActionSequenceParameters**](ActionSequenceParameters.md) |  | [optional] 
**Actions** | Pointer to [**[]Action**](Action.md) |  | [optional] 

## Methods

### NewActionSequence

`func NewActionSequence() *ActionSequence`

NewActionSequence instantiates a new ActionSequence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionSequenceWithDefaults

`func NewActionSequenceWithDefaults() *ActionSequence`

NewActionSequenceWithDefaults instantiates a new ActionSequence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionSequence) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionSequence) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionSequence) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActionSequence) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ActionSequence) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionSequence) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionSequence) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActionSequence) HasType() bool`

HasType returns a boolean if a field has been set.

### GetParameters

`func (o *ActionSequence) GetParameters() ActionSequenceParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ActionSequence) GetParametersOk() (*ActionSequenceParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ActionSequence) SetParameters(v ActionSequenceParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ActionSequence) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetActions

`func (o *ActionSequence) GetActions() []Action`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ActionSequence) GetActionsOk() (*[]Action, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ActionSequence) SetActions(v []Action)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ActionSequence) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


