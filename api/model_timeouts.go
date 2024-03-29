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

// Timeouts struct for Timeouts
type Timeouts struct {
	Script NilableInt64 `json:"script,omitempty"`
	PageLoad *int64 `json:"pageLoad,omitempty"`
	Implicit *int64 `json:"implicit,omitempty"`
}

// NewTimeouts instantiates a new Timeouts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeouts() *Timeouts {
	this := Timeouts{}
	return &this
}

// NewTimeoutsWithDefaults instantiates a new Timeouts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeoutsWithDefaults() *Timeouts {
	this := Timeouts{}
	return &this
}

// GetScript returns the Script field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Timeouts) GetScript() int64 {
	if o == nil || o.Script.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Script.Get()
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Timeouts) GetScriptOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Script.Get(), o.Script.IsSet()
}

// HasScript returns a boolean if a field has been set.
func (o *Timeouts) HasScript() bool {
	if o != nil && o.Script.IsSet() {
		return true
	}

	return false
}

// SetScript gets a reference to the given NilableInt64 and assigns it to the Script field.
func (o *Timeouts) SetScript(v int64) {
	o.Script.Set(&v)
}
// SetScriptNil sets the value for Script to be an explicit nil
func (o *Timeouts) SetScriptNil() {
	o.Script.Set(nil)
}

// UnsetScript ensures that no value is present for Script, not even an explicit nil
func (o *Timeouts) UnsetScript() {
	o.Script.Unset()
}

// GetPageLoad returns the PageLoad field value if set, zero value otherwise.
func (o *Timeouts) GetPageLoad() int64 {
	if o == nil || o.PageLoad == nil {
		var ret int64
		return ret
	}
	return *o.PageLoad
}

// GetPageLoadOk returns a tuple with the PageLoad field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeouts) GetPageLoadOk() (*int64, bool) {
	if o == nil || o.PageLoad == nil {
		return nil, false
	}
	return o.PageLoad, true
}

// HasPageLoad returns a boolean if a field has been set.
func (o *Timeouts) HasPageLoad() bool {
	if o != nil && o.PageLoad != nil {
		return true
	}

	return false
}

// SetPageLoad gets a reference to the given int64 and assigns it to the PageLoad field.
func (o *Timeouts) SetPageLoad(v int64) {
	o.PageLoad = &v
}

// GetImplicit returns the Implicit field value if set, zero value otherwise.
func (o *Timeouts) GetImplicit() int64 {
	if o == nil || o.Implicit == nil {
		var ret int64
		return ret
	}
	return *o.Implicit
}

// GetImplicitOk returns a tuple with the Implicit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeouts) GetImplicitOk() (*int64, bool) {
	if o == nil || o.Implicit == nil {
		return nil, false
	}
	return o.Implicit, true
}

// HasImplicit returns a boolean if a field has been set.
func (o *Timeouts) HasImplicit() bool {
	if o != nil && o.Implicit != nil {
		return true
	}

	return false
}

// SetImplicit gets a reference to the given int64 and assigns it to the Implicit field.
func (o *Timeouts) SetImplicit(v int64) {
	o.Implicit = &v
}

func (o Timeouts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Script.IsSet() {
		toSerialize["script"] = o.Script.Get()
	}
	if o.PageLoad != nil {
		toSerialize["pageLoad"] = o.PageLoad
	}
	if o.Implicit != nil {
		toSerialize["implicit"] = o.Implicit
	}
	return json.Marshal(toSerialize)
}

type NilableTimeouts struct {
	value *Timeouts
	isSet bool
}

func (v NilableTimeouts) Get() *Timeouts {
	return v.value
}

func (v *NilableTimeouts) Set(val *Timeouts) {
	v.value = val
	v.isSet = true
}

func (v NilableTimeouts) IsSet() bool {
	return v.isSet
}

func (v *NilableTimeouts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNilableTimeouts(val *Timeouts) *NilableTimeouts {
	return &NilableTimeouts{value: val, isSet: true}
}

func (v NilableTimeouts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NilableTimeouts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


