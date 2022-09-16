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

// NewWindowResponseValue struct for NewWindowResponseValue
type NewWindowResponseValue struct {
	Handle string `json:"handle"`
	Type string `json:"type"`
}

// NewNewWindowResponseValue instantiates a new NewWindowResponseValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewWindowResponseValue(handle string, type_ string) *NewWindowResponseValue {
	this := NewWindowResponseValue{}
	this.Handle = handle
	this.Type = type_
	return &this
}

// NewNewWindowResponseValueWithDefaults instantiates a new NewWindowResponseValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewWindowResponseValueWithDefaults() *NewWindowResponseValue {
	this := NewWindowResponseValue{}
	return &this
}

// GetHandle returns the Handle field value
func (o *NewWindowResponseValue) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *NewWindowResponseValue) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value
func (o *NewWindowResponseValue) SetHandle(v string) {
	o.Handle = v
}

// GetType returns the Type field value
func (o *NewWindowResponseValue) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NewWindowResponseValue) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NewWindowResponseValue) SetType(v string) {
	o.Type = v
}

func (o NewWindowResponseValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["handle"] = o.Handle
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NilableNewWindowResponseValue struct {
	value *NewWindowResponseValue
	isSet bool
}

func (v NilableNewWindowResponseValue) Get() *NewWindowResponseValue {
	return v.value
}

func (v *NilableNewWindowResponseValue) Set(val *NewWindowResponseValue) {
	v.value = val
	v.isSet = true
}

func (v NilableNewWindowResponseValue) IsSet() bool {
	return v.isSet
}

func (v *NilableNewWindowResponseValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNilableNewWindowResponseValue(val *NewWindowResponseValue) *NilableNewWindowResponseValue {
	return &NilableNewWindowResponseValue{value: val, isSet: true}
}

func (v NilableNewWindowResponseValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NilableNewWindowResponseValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

