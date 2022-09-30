# ScriptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | Pointer to **string** |  | [optional] 
**Args** | Pointer to **[]string** |  | [optional] 

## Methods

### NewScriptRequest

`func NewScriptRequest() *ScriptRequest`

NewScriptRequest instantiates a new ScriptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptRequestWithDefaults

`func NewScriptRequestWithDefaults() *ScriptRequest`

NewScriptRequestWithDefaults instantiates a new ScriptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScript

`func (o *ScriptRequest) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *ScriptRequest) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *ScriptRequest) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *ScriptRequest) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetArgs

`func (o *ScriptRequest) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ScriptRequest) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ScriptRequest) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *ScriptRequest) HasArgs() bool`

HasArgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


