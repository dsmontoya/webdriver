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

// ClipboardData struct for ClipboardData
type ClipboardData struct {
	Value string `json:"value"`
	Media *string `json:"media,omitempty"`
}

// NewClipboardData instantiates a new ClipboardData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClipboardData(value string) *ClipboardData {
	this := ClipboardData{}
	this.Value = value
	return &this
}

// NewClipboardDataWithDefaults instantiates a new ClipboardData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClipboardDataWithDefaults() *ClipboardData {
	this := ClipboardData{}
	return &this
}

// GetValue returns the Value field value
func (o *ClipboardData) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ClipboardData) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ClipboardData) SetValue(v string) {
	o.Value = v
}

// GetMedia returns the Media field value if set, zero value otherwise.
func (o *ClipboardData) GetMedia() string {
	if o == nil || o.Media == nil {
		var ret string
		return ret
	}
	return *o.Media
}

// GetMediaOk returns a tuple with the Media field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClipboardData) GetMediaOk() (*string, bool) {
	if o == nil || o.Media == nil {
		return nil, false
	}
	return o.Media, true
}

// HasMedia returns a boolean if a field has been set.
func (o *ClipboardData) HasMedia() bool {
	if o != nil && o.Media != nil {
		return true
	}

	return false
}

// SetMedia gets a reference to the given string and assigns it to the Media field.
func (o *ClipboardData) SetMedia(v string) {
	o.Media = &v
}

func (o ClipboardData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if o.Media != nil {
		toSerialize["media"] = o.Media
	}
	return json.Marshal(toSerialize)
}

type NilableClipboardData struct {
	value *ClipboardData
	isSet bool
}

func (v NilableClipboardData) Get() *ClipboardData {
	return v.value
}

func (v *NilableClipboardData) Set(val *ClipboardData) {
	v.value = val
	v.isSet = true
}

func (v NilableClipboardData) IsSet() bool {
	return v.isSet
}

func (v *NilableClipboardData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNilableClipboardData(val *ClipboardData) *NilableClipboardData {
	return &NilableClipboardData{value: val, isSet: true}
}

func (v NilableClipboardData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NilableClipboardData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


