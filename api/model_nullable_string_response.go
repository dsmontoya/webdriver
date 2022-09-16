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

// NullableStringResponse struct for NullableStringResponse
type NullableStringResponse struct {
	// String that can sometimes have null value
	Value NilableString `json:"value,omitempty"`
}

// NewNullableStringResponse instantiates a new NullableStringResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNullableStringResponse() *NullableStringResponse {
	this := NullableStringResponse{}
	return &this
}

// NewNullableStringResponseWithDefaults instantiates a new NullableStringResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNullableStringResponseWithDefaults() *NullableStringResponse {
	this := NullableStringResponse{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NullableStringResponse) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullableStringResponse) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *NullableStringResponse) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NilableString and assigns it to the Value field.
func (o *NullableStringResponse) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *NullableStringResponse) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *NullableStringResponse) UnsetValue() {
	o.Value.Unset()
}

func (o NullableStringResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return json.Marshal(toSerialize)
}

type NilableNullableStringResponse struct {
	value *NullableStringResponse
	isSet bool
}

func (v NilableNullableStringResponse) Get() *NullableStringResponse {
	return v.value
}

func (v *NilableNullableStringResponse) Set(val *NullableStringResponse) {
	v.value = val
	v.isSet = true
}

func (v NilableNullableStringResponse) IsSet() bool {
	return v.isSet
}

func (v *NilableNullableStringResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNilableNullableStringResponse(val *NullableStringResponse) *NilableNullableStringResponse {
	return &NilableNullableStringResponse{value: val, isSet: true}
}

func (v NilableNullableStringResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NilableNullableStringResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

