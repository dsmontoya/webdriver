# FindElementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Using** | [**LocatorStrategy**](LocatorStrategy.md) |  | 
**Value** | **string** | Selector value | 

## Methods

### NewFindElementRequest

`func NewFindElementRequest(using LocatorStrategy, value string, ) *FindElementRequest`

NewFindElementRequest instantiates a new FindElementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindElementRequestWithDefaults

`func NewFindElementRequestWithDefaults() *FindElementRequest`

NewFindElementRequestWithDefaults instantiates a new FindElementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsing

`func (o *FindElementRequest) GetUsing() LocatorStrategy`

GetUsing returns the Using field if non-nil, zero value otherwise.

### GetUsingOk

`func (o *FindElementRequest) GetUsingOk() (*LocatorStrategy, bool)`

GetUsingOk returns a tuple with the Using field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsing

`func (o *FindElementRequest) SetUsing(v LocatorStrategy)`

SetUsing sets Using field to given value.


### GetValue

`func (o *FindElementRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FindElementRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FindElementRequest) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


