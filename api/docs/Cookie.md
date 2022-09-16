# Cookie

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | **string** |  | 
**Path** | **string** |  | 
**Domain** | **string** |  | 
**Secure** | **bool** |  | 
**HttpOnly** | **bool** |  | 
**Expiry** | Pointer to **int64** |  | [optional] 
**SameSite** | **string** |  | 

## Methods

### NewCookie

`func NewCookie(name string, value string, path string, domain string, secure bool, httpOnly bool, sameSite string, ) *Cookie`

NewCookie instantiates a new Cookie object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCookieWithDefaults

`func NewCookieWithDefaults() *Cookie`

NewCookieWithDefaults instantiates a new Cookie object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Cookie) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cookie) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cookie) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *Cookie) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Cookie) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Cookie) SetValue(v string)`

SetValue sets Value field to given value.


### GetPath

`func (o *Cookie) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Cookie) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Cookie) SetPath(v string)`

SetPath sets Path field to given value.


### GetDomain

`func (o *Cookie) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Cookie) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Cookie) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetSecure

`func (o *Cookie) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *Cookie) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *Cookie) SetSecure(v bool)`

SetSecure sets Secure field to given value.


### GetHttpOnly

`func (o *Cookie) GetHttpOnly() bool`

GetHttpOnly returns the HttpOnly field if non-nil, zero value otherwise.

### GetHttpOnlyOk

`func (o *Cookie) GetHttpOnlyOk() (*bool, bool)`

GetHttpOnlyOk returns a tuple with the HttpOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpOnly

`func (o *Cookie) SetHttpOnly(v bool)`

SetHttpOnly sets HttpOnly field to given value.


### GetExpiry

`func (o *Cookie) GetExpiry() int64`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *Cookie) GetExpiryOk() (*int64, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *Cookie) SetExpiry(v int64)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *Cookie) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetSameSite

`func (o *Cookie) GetSameSite() string`

GetSameSite returns the SameSite field if non-nil, zero value otherwise.

### GetSameSiteOk

`func (o *Cookie) GetSameSiteOk() (*string, bool)`

GetSameSiteOk returns a tuple with the SameSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSameSite

`func (o *Cookie) SetSameSite(v string)`

SetSameSite sets SameSite field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


