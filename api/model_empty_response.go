/*
Selenium WebDriver

[Selenium WebDriver](https://www.w3.org/TR/webdriver) API specification

API version: 1.0.0
Contact: support@aerokube.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// EmptyResponse struct for EmptyResponse
type EmptyResponse struct {
	// Always null value
	Value NilableString `json:"value,omitempty"`
}

// NewEmptyResponse instantiates a new EmptyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmptyResponse() *EmptyResponse {
	this := EmptyResponse{}
	return &this
}

// NewEmptyResponseWithDefaults instantiates a new EmptyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmptyResponseWithDefaults() *EmptyResponse {
	this := EmptyResponse{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmptyResponse) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmptyResponse) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *EmptyResponse) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NilableString and assigns it to the Value field.
func (o *EmptyResponse) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *EmptyResponse) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *EmptyResponse) UnsetValue() {
	o.Value.Unset()
}

func (o EmptyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return json.Marshal(toSerialize)
}

type NilableEmptyResponse struct {
	value *EmptyResponse
	isSet bool
}

func (v NilableEmptyResponse) Get() *EmptyResponse {
	return v.value
}

func (v *NilableEmptyResponse) Set(val *EmptyResponse) {
	v.value = val
	v.isSet = true
}

func (v NilableEmptyResponse) IsSet() bool {
	return v.isSet
}

func (v *NilableEmptyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNilableEmptyResponse(val *EmptyResponse) *NilableEmptyResponse {
	return &NilableEmptyResponse{value: val, isSet: true}
}

func (v NilableEmptyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NilableEmptyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

