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

// NewSessionRequest struct for NewSessionRequest
type NewSessionRequest struct {
	Capabilities NewSessionRequestCapabilities `json:"capabilities"`
}

// NewNewSessionRequest instantiates a new NewSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewSessionRequest(capabilities NewSessionRequestCapabilities) *NewSessionRequest {
	this := NewSessionRequest{}
	this.Capabilities = capabilities
	return &this
}

// NewNewSessionRequestWithDefaults instantiates a new NewSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewSessionRequestWithDefaults() *NewSessionRequest {
	this := NewSessionRequest{}
	return &this
}

// GetCapabilities returns the Capabilities field value
func (o *NewSessionRequest) GetCapabilities() NewSessionRequestCapabilities {
	if o == nil {
		var ret NewSessionRequestCapabilities
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *NewSessionRequest) GetCapabilitiesOk() (*NewSessionRequestCapabilities, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capabilities, true
}

// SetCapabilities sets field value
func (o *NewSessionRequest) SetCapabilities(v NewSessionRequestCapabilities) {
	o.Capabilities = v
}

func (o NewSessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["capabilities"] = o.Capabilities
	}
	return json.Marshal(toSerialize)
}

type NilableNewSessionRequest struct {
	value *NewSessionRequest
	isSet bool
}

func (v NilableNewSessionRequest) Get() *NewSessionRequest {
	return v.value
}

func (v *NilableNewSessionRequest) Set(val *NewSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NilableNewSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NilableNewSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNilableNewSessionRequest(val *NewSessionRequest) *NilableNewSessionRequest {
	return &NilableNewSessionRequest{value: val, isSet: true}
}

func (v NilableNewSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NilableNewSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


