# NewSessionRequestCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlwaysMatch** | Pointer to [**Capabilities**](Capabilities.md) |  | [optional] 
**FirstMatch** | Pointer to [**[]Capabilities**](Capabilities.md) |  | [optional] 

## Methods

### NewNewSessionRequestCapabilities

`func NewNewSessionRequestCapabilities() *NewSessionRequestCapabilities`

NewNewSessionRequestCapabilities instantiates a new NewSessionRequestCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewSessionRequestCapabilitiesWithDefaults

`func NewNewSessionRequestCapabilitiesWithDefaults() *NewSessionRequestCapabilities`

NewNewSessionRequestCapabilitiesWithDefaults instantiates a new NewSessionRequestCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlwaysMatch

`func (o *NewSessionRequestCapabilities) GetAlwaysMatch() Capabilities`

GetAlwaysMatch returns the AlwaysMatch field if non-nil, zero value otherwise.

### GetAlwaysMatchOk

`func (o *NewSessionRequestCapabilities) GetAlwaysMatchOk() (*Capabilities, bool)`

GetAlwaysMatchOk returns a tuple with the AlwaysMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysMatch

`func (o *NewSessionRequestCapabilities) SetAlwaysMatch(v Capabilities)`

SetAlwaysMatch sets AlwaysMatch field to given value.

### HasAlwaysMatch

`func (o *NewSessionRequestCapabilities) HasAlwaysMatch() bool`

HasAlwaysMatch returns a boolean if a field has been set.

### GetFirstMatch

`func (o *NewSessionRequestCapabilities) GetFirstMatch() []Capabilities`

GetFirstMatch returns the FirstMatch field if non-nil, zero value otherwise.

### GetFirstMatchOk

`func (o *NewSessionRequestCapabilities) GetFirstMatchOk() (*[]Capabilities, bool)`

GetFirstMatchOk returns a tuple with the FirstMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstMatch

`func (o *NewSessionRequestCapabilities) SetFirstMatch(v []Capabilities)`

SetFirstMatch sets FirstMatch field to given value.

### HasFirstMatch

`func (o *NewSessionRequestCapabilities) HasFirstMatch() bool`

HasFirstMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


