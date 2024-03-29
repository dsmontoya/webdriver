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

// GetWindowHandlesResponse struct for GetWindowHandlesResponse
type GetWindowHandlesResponse struct {
	Value []string `json:"value"`
}

// NewGetWindowHandlesResponse instantiates a new GetWindowHandlesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWindowHandlesResponse(value []string) *GetWindowHandlesResponse {
	this := GetWindowHandlesResponse{}
	this.Value = value
	return &this
}

// NewGetWindowHandlesResponseWithDefaults instantiates a new GetWindowHandlesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWindowHandlesResponseWithDefaults() *GetWindowHandlesResponse {
	this := GetWindowHandlesResponse{}
	return &this
}

// GetValue returns the Value field value
func (o *GetWindowHandlesResponse) GetValue() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetWindowHandlesResponse) GetValueOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *GetWindowHandlesResponse) SetValue(v []string) {
	o.Value = v
}

func (o GetWindowHandlesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NilableGetWindowHandlesResponse struct {
	value *GetWindowHandlesResponse
	isSet bool
}

func (v NilableGetWindowHandlesResponse) Get() *GetWindowHandlesResponse {
	return v.value
}

func (v *NilableGetWindowHandlesResponse) Set(val *GetWindowHandlesResponse) {
	v.value = val
	v.isSet = true
}

func (v NilableGetWindowHandlesResponse) IsSet() bool {
	return v.isSet
}

func (v *NilableGetWindowHandlesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNilableGetWindowHandlesResponse(val *GetWindowHandlesResponse) *NilableGetWindowHandlesResponse {
	return &NilableGetWindowHandlesResponse{value: val, isSet: true}
}

func (v NilableGetWindowHandlesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NilableGetWindowHandlesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


