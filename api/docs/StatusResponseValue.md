# StatusResponseValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ready** | **bool** |  | 
**Message** | **string** |  | 

## Methods

### NewStatusResponseValue

`func NewStatusResponseValue(ready bool, message string, ) *StatusResponseValue`

NewStatusResponseValue instantiates a new StatusResponseValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusResponseValueWithDefaults

`func NewStatusResponseValueWithDefaults() *StatusResponseValue`

NewStatusResponseValueWithDefaults instantiates a new StatusResponseValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReady

`func (o *StatusResponseValue) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *StatusResponseValue) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *StatusResponseValue) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetMessage

`func (o *StatusResponseValue) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StatusResponseValue) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StatusResponseValue) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


