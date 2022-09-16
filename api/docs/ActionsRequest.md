# ActionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | [**[]ActionSequence**](ActionSequence.md) | A list of actions to be performed | 

## Methods

### NewActionsRequest

`func NewActionsRequest(actions []ActionSequence, ) *ActionsRequest`

NewActionsRequest instantiates a new ActionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsRequestWithDefaults

`func NewActionsRequestWithDefaults() *ActionsRequest`

NewActionsRequestWithDefaults instantiates a new ActionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ActionsRequest) GetActions() []ActionSequence`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ActionsRequest) GetActionsOk() (*[]ActionSequence, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ActionsRequest) SetActions(v []ActionSequence)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


