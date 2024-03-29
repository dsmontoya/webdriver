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

// StringResponse struct for StringResponse
type StringResponse struct {
	Value string `json:"value"`
}

// NewStringResponse instantiates a new StringResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringResponse(value string) *StringResponse {
	this := StringResponse{}
	this.Value = value
	return &this
}

// NewStringResponseWithDefaults instantiates a new StringResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringResponseWithDefaults() *StringResponse {
	this := StringResponse{}
	return &this
}

// GetValue returns the Value field value
func (o *StringResponse) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *StringResponse) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *StringResponse) SetValue(v string) {
	o.Value = v
}

func (o StringResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NilableStringResponse struct {
	value *StringResponse
	isSet bool
}

func (v NilableStringResponse) Get() *StringResponse {
	return v.value
}

func (v *NilableStringResponse) Set(val *StringResponse) {
	v.value = val
	v.isSet = true
}

func (v NilableStringResponse) IsSet() bool {
	return v.isSet
}

func (v *NilableStringResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNilableStringResponse(val *StringResponse) *NilableStringResponse {
	return &NilableStringResponse{value: val, isSet: true}
}

func (v NilableStringResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NilableStringResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


