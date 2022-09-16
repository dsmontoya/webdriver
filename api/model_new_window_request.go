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

// NewWindowRequest struct for NewWindowRequest
type NewWindowRequest struct {
	Type string `json:"type"`
}

// NewNewWindowRequest instantiates a new NewWindowRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewWindowRequest(type_ string) *NewWindowRequest {
	this := NewWindowRequest{}
	this.Type = type_
	return &this
}

// NewNewWindowRequestWithDefaults instantiates a new NewWindowRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewWindowRequestWithDefaults() *NewWindowRequest {
	this := NewWindowRequest{}
	return &this
}

// GetType returns the Type field value
func (o *NewWindowRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NewWindowRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NewWindowRequest) SetType(v string) {
	o.Type = v
}

func (o NewWindowRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NilableNewWindowRequest struct {
	value *NewWindowRequest
	isSet bool
}

func (v NilableNewWindowRequest) Get() *NewWindowRequest {
	return v.value
}

func (v *NilableNewWindowRequest) Set(val *NewWindowRequest) {
	v.value = val
	v.isSet = true
}

func (v NilableNewWindowRequest) IsSet() bool {
	return v.isSet
}

func (v *NilableNewWindowRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNilableNewWindowRequest(val *NewWindowRequest) *NilableNewWindowRequest {
	return &NilableNewWindowRequest{value: val, isSet: true}
}

func (v NilableNewWindowRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NilableNewWindowRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

