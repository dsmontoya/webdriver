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

// PrintRequestOptionsMargin struct for PrintRequestOptionsMargin
type PrintRequestOptionsMargin struct {
	Top *float32 `json:"top,omitempty"`
	Bottom *float32 `json:"bottom,omitempty"`
	Left *float32 `json:"left,omitempty"`
	Right *float32 `json:"right,omitempty"`
}

// NewPrintRequestOptionsMargin instantiates a new PrintRequestOptionsMargin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrintRequestOptionsMargin() *PrintRequestOptionsMargin {
	this := PrintRequestOptionsMargin{}
	var top float32 = 1
	this.Top = &top
	var bottom float32 = 1
	this.Bottom = &bottom
	var left float32 = 1
	this.Left = &left
	var right float32 = 1
	this.Right = &right
	return &this
}

// NewPrintRequestOptionsMarginWithDefaults instantiates a new PrintRequestOptionsMargin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrintRequestOptionsMarginWithDefaults() *PrintRequestOptionsMargin {
	this := PrintRequestOptionsMargin{}
	var top float32 = 1
	this.Top = &top
	var bottom float32 = 1
	this.Bottom = &bottom
	var left float32 = 1
	this.Left = &left
	var right float32 = 1
	this.Right = &right
	return &this
}

// GetTop returns the Top field value if set, zero value otherwise.
func (o *PrintRequestOptionsMargin) GetTop() float32 {
	if o == nil || o.Top == nil {
		var ret float32
		return ret
	}
	return *o.Top
}

// GetTopOk returns a tuple with the Top field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintRequestOptionsMargin) GetTopOk() (*float32, bool) {
	if o == nil || o.Top == nil {
		return nil, false
	}
	return o.Top, true
}

// HasTop returns a boolean if a field has been set.
func (o *PrintRequestOptionsMargin) HasTop() bool {
	if o != nil && o.Top != nil {
		return true
	}

	return false
}

// SetTop gets a reference to the given float32 and assigns it to the Top field.
func (o *PrintRequestOptionsMargin) SetTop(v float32) {
	o.Top = &v
}

// GetBottom returns the Bottom field value if set, zero value otherwise.
func (o *PrintRequestOptionsMargin) GetBottom() float32 {
	if o == nil || o.Bottom == nil {
		var ret float32
		return ret
	}
	return *o.Bottom
}

// GetBottomOk returns a tuple with the Bottom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintRequestOptionsMargin) GetBottomOk() (*float32, bool) {
	if o == nil || o.Bottom == nil {
		return nil, false
	}
	return o.Bottom, true
}

// HasBottom returns a boolean if a field has been set.
func (o *PrintRequestOptionsMargin) HasBottom() bool {
	if o != nil && o.Bottom != nil {
		return true
	}

	return false
}

// SetBottom gets a reference to the given float32 and assigns it to the Bottom field.
func (o *PrintRequestOptionsMargin) SetBottom(v float32) {
	o.Bottom = &v
}

// GetLeft returns the Left field value if set, zero value otherwise.
func (o *PrintRequestOptionsMargin) GetLeft() float32 {
	if o == nil || o.Left == nil {
		var ret float32
		return ret
	}
	return *o.Left
}

// GetLeftOk returns a tuple with the Left field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintRequestOptionsMargin) GetLeftOk() (*float32, bool) {
	if o == nil || o.Left == nil {
		return nil, false
	}
	return o.Left, true
}

// HasLeft returns a boolean if a field has been set.
func (o *PrintRequestOptionsMargin) HasLeft() bool {
	if o != nil && o.Left != nil {
		return true
	}

	return false
}

// SetLeft gets a reference to the given float32 and assigns it to the Left field.
func (o *PrintRequestOptionsMargin) SetLeft(v float32) {
	o.Left = &v
}

// GetRight returns the Right field value if set, zero value otherwise.
func (o *PrintRequestOptionsMargin) GetRight() float32 {
	if o == nil || o.Right == nil {
		var ret float32
		return ret
	}
	return *o.Right
}

// GetRightOk returns a tuple with the Right field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintRequestOptionsMargin) GetRightOk() (*float32, bool) {
	if o == nil || o.Right == nil {
		return nil, false
	}
	return o.Right, true
}

// HasRight returns a boolean if a field has been set.
func (o *PrintRequestOptionsMargin) HasRight() bool {
	if o != nil && o.Right != nil {
		return true
	}

	return false
}

// SetRight gets a reference to the given float32 and assigns it to the Right field.
func (o *PrintRequestOptionsMargin) SetRight(v float32) {
	o.Right = &v
}

func (o PrintRequestOptionsMargin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Top != nil {
		toSerialize["top"] = o.Top
	}
	if o.Bottom != nil {
		toSerialize["bottom"] = o.Bottom
	}
	if o.Left != nil {
		toSerialize["left"] = o.Left
	}
	if o.Right != nil {
		toSerialize["right"] = o.Right
	}
	return json.Marshal(toSerialize)
}

type NilablePrintRequestOptionsMargin struct {
	value *PrintRequestOptionsMargin
	isSet bool
}

func (v NilablePrintRequestOptionsMargin) Get() *PrintRequestOptionsMargin {
	return v.value
}

func (v *NilablePrintRequestOptionsMargin) Set(val *PrintRequestOptionsMargin) {
	v.value = val
	v.isSet = true
}

func (v NilablePrintRequestOptionsMargin) IsSet() bool {
	return v.isSet
}

func (v *NilablePrintRequestOptionsMargin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNilablePrintRequestOptionsMargin(val *PrintRequestOptionsMargin) *NilablePrintRequestOptionsMargin {
	return &NilablePrintRequestOptionsMargin{value: val, isSet: true}
}

func (v NilablePrintRequestOptionsMargin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NilablePrintRequestOptionsMargin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


