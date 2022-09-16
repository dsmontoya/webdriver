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

// WebElement struct for WebElement
type WebElement struct {
	Element606611e4A52e4f735466cecf string `json:"element-6066-11e4-a52e-4f735466cecf"`
}

// NewWebElement instantiates a new WebElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebElement(element606611e4A52e4f735466cecf string) *WebElement {
	this := WebElement{}
	this.Element606611e4A52e4f735466cecf = element606611e4A52e4f735466cecf
	return &this
}

// NewWebElementWithDefaults instantiates a new WebElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebElementWithDefaults() *WebElement {
	this := WebElement{}
	return &this
}

// GetElement606611e4A52e4f735466cecf returns the Element606611e4A52e4f735466cecf field value
func (o *WebElement) GetElement606611e4A52e4f735466cecf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Element606611e4A52e4f735466cecf
}

// GetElement606611e4A52e4f735466cecfOk returns a tuple with the Element606611e4A52e4f735466cecf field value
// and a boolean to check if the value has been set.
func (o *WebElement) GetElement606611e4A52e4f735466cecfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Element606611e4A52e4f735466cecf, true
}

// SetElement606611e4A52e4f735466cecf sets field value
func (o *WebElement) SetElement606611e4A52e4f735466cecf(v string) {
	o.Element606611e4A52e4f735466cecf = v
}

func (o WebElement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["element-6066-11e4-a52e-4f735466cecf"] = o.Element606611e4A52e4f735466cecf
	}
	return json.Marshal(toSerialize)
}

type NilableWebElement struct {
	value *WebElement
	isSet bool
}

func (v NilableWebElement) Get() *WebElement {
	return v.value
}

func (v *NilableWebElement) Set(val *WebElement) {
	v.value = val
	v.isSet = true
}

func (v NilableWebElement) IsSet() bool {
	return v.isSet
}

func (v *NilableWebElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNilableWebElement(val *WebElement) *NilableWebElement {
	return &NilableWebElement{value: val, isSet: true}
}

func (v NilableWebElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NilableWebElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


