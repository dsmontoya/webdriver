# ErrorResponseValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**ErrorCode**](ErrorCode.md) |  | 
**Message** | **string** |  | 
**Stacktrace** | **string** |  | 
**Data** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewErrorResponseValue

`func NewErrorResponseValue(error_ ErrorCode, message string, stacktrace string, ) *ErrorResponseValue`

NewErrorResponseValue instantiates a new ErrorResponseValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseValueWithDefaults

`func NewErrorResponseValueWithDefaults() *ErrorResponseValue`

NewErrorResponseValueWithDefaults instantiates a new ErrorResponseValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ErrorResponseValue) GetError() ErrorCode`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ErrorResponseValue) GetErrorOk() (*ErrorCode, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ErrorResponseValue) SetError(v ErrorCode)`

SetError sets Error field to given value.


### GetMessage

`func (o *ErrorResponseValue) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponseValue) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponseValue) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStacktrace

`func (o *ErrorResponseValue) GetStacktrace() string`

GetStacktrace returns the Stacktrace field if non-nil, zero value otherwise.

### GetStacktraceOk

`func (o *ErrorResponseValue) GetStacktraceOk() (*string, bool)`

GetStacktraceOk returns a tuple with the Stacktrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacktrace

`func (o *ErrorResponseValue) SetStacktrace(v string)`

SetStacktrace sets Stacktrace field to given value.


### GetData

`func (o *ErrorResponseValue) GetData() map[string]map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ErrorResponseValue) GetDataOk() (*map[string]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ErrorResponseValue) SetData(v map[string]map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ErrorResponseValue) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


