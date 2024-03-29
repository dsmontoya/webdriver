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

// NewSessionRequestCapabilities struct for NewSessionRequestCapabilities
type NewSessionRequestCapabilities struct {
	AlwaysMatch *Capabilities `json:"alwaysMatch,omitempty"`
	FirstMatch []Capabilities `json:"firstMatch,omitempty"`
}

// NewNewSessionRequestCapabilities instantiates a new NewSessionRequestCapabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewSessionRequestCapabilities() *NewSessionRequestCapabilities {
	this := NewSessionRequestCapabilities{}
	return &this
}

// NewNewSessionRequestCapabilitiesWithDefaults instantiates a new NewSessionRequestCapabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewSessionRequestCapabilitiesWithDefaults() *NewSessionRequestCapabilities {
	this := NewSessionRequestCapabilities{}
	return &this
}

// GetAlwaysMatch returns the AlwaysMatch field value if set, zero value otherwise.
func (o *NewSessionRequestCapabilities) GetAlwaysMatch() Capabilities {
	if o == nil || o.AlwaysMatch == nil {
		var ret Capabilities
		return ret
	}
	return *o.AlwaysMatch
}

// GetAlwaysMatchOk returns a tuple with the AlwaysMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewSessionRequestCapabilities) GetAlwaysMatchOk() (*Capabilities, bool) {
	if o == nil || o.AlwaysMatch == nil {
		return nil, false
	}
	return o.AlwaysMatch, true
}

// HasAlwaysMatch returns a boolean if a field has been set.
func (o *NewSessionRequestCapabilities) HasAlwaysMatch() bool {
	if o != nil && o.AlwaysMatch != nil {
		return true
	}

	return false
}

// SetAlwaysMatch gets a reference to the given Capabilities and assigns it to the AlwaysMatch field.
func (o *NewSessionRequestCapabilities) SetAlwaysMatch(v Capabilities) {
	o.AlwaysMatch = &v
}

// GetFirstMatch returns the FirstMatch field value if set, zero value otherwise.
func (o *NewSessionRequestCapabilities) GetFirstMatch() []Capabilities {
	if o == nil || o.FirstMatch == nil {
		var ret []Capabilities
		return ret
	}
	return o.FirstMatch
}

// GetFirstMatchOk returns a tuple with the FirstMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewSessionRequestCapabilities) GetFirstMatchOk() ([]Capabilities, bool) {
	if o == nil || o.FirstMatch == nil {
		return nil, false
	}
	return o.FirstMatch, true
}

// HasFirstMatch returns a boolean if a field has been set.
func (o *NewSessionRequestCapabilities) HasFirstMatch() bool {
	if o != nil && o.FirstMatch != nil {
		return true
	}

	return false
}

// SetFirstMatch gets a reference to the given []Capabilities and assigns it to the FirstMatch field.
func (o *NewSessionRequestCapabilities) SetFirstMatch(v []Capabilities) {
	o.FirstMatch = v
}

func (o NewSessionRequestCapabilities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AlwaysMatch != nil {
		toSerialize["alwaysMatch"] = o.AlwaysMatch
	}
	if o.FirstMatch != nil {
		toSerialize["firstMatch"] = o.FirstMatch
	}
	return json.Marshal(toSerialize)
}

type NilableNewSessionRequestCapabilities struct {
	value *NewSessionRequestCapabilities
	isSet bool
}

func (v NilableNewSessionRequestCapabilities) Get() *NewSessionRequestCapabilities {
	return v.value
}

func (v *NilableNewSessionRequestCapabilities) Set(val *NewSessionRequestCapabilities) {
	v.value = val
	v.isSet = true
}

func (v NilableNewSessionRequestCapabilities) IsSet() bool {
	return v.isSet
}

func (v *NilableNewSessionRequestCapabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNilableNewSessionRequestCapabilities(val *NewSessionRequestCapabilities) *NilableNewSessionRequestCapabilities {
	return &NilableNewSessionRequestCapabilities{value: val, isSet: true}
}

func (v NilableNewSessionRequestCapabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NilableNewSessionRequestCapabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


