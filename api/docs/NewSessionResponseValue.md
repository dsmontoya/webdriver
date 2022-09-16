# NewSessionResponseValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | **string** |  | 
**Capabilities** | [**Capabilities**](Capabilities.md) |  | 

## Methods

### NewNewSessionResponseValue

`func NewNewSessionResponseValue(sessionId string, capabilities Capabilities, ) *NewSessionResponseValue`

NewNewSessionResponseValue instantiates a new NewSessionResponseValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewSessionResponseValueWithDefaults

`func NewNewSessionResponseValueWithDefaults() *NewSessionResponseValue`

NewNewSessionResponseValueWithDefaults instantiates a new NewSessionResponseValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *NewSessionResponseValue) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *NewSessionResponseValue) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *NewSessionResponseValue) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetCapabilities

`func (o *NewSessionResponseValue) GetCapabilities() Capabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *NewSessionResponseValue) GetCapabilitiesOk() (*Capabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *NewSessionResponseValue) SetCapabilities(v Capabilities)`

SetCapabilities sets Capabilities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


