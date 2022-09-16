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

// Rect struct for Rect
type Rect struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Width float32 `json:"width"`
	Height float32 `json:"height"`
}

// NewRect instantiates a new Rect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRect(x float32, y float32, width float32, height float32) *Rect {
	this := Rect{}
	this.X = x
	this.Y = y
	this.Width = width
	this.Height = height
	return &this
}

// NewRectWithDefaults instantiates a new Rect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRectWithDefaults() *Rect {
	this := Rect{}
	return &this
}

// GetX returns the X field value
func (o *Rect) GetX() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *Rect) GetXOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *Rect) SetX(v float32) {
	o.X = v
}

// GetY returns the Y field value
func (o *Rect) GetY() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *Rect) GetYOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *Rect) SetY(v float32) {
	o.Y = v
}

// GetWidth returns the Width field value
func (o *Rect) GetWidth() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *Rect) GetWidthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *Rect) SetWidth(v float32) {
	o.Width = v
}

// GetHeight returns the Height field value
func (o *Rect) GetHeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *Rect) GetHeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *Rect) SetHeight(v float32) {
	o.Height = v
}

func (o Rect) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["x"] = o.X
	}
	if true {
		toSerialize["y"] = o.Y
	}
	if true {
		toSerialize["width"] = o.Width
	}
	if true {
		toSerialize["height"] = o.Height
	}
	return json.Marshal(toSerialize)
}

type NilableRect struct {
	value *Rect
	isSet bool
}

func (v NilableRect) Get() *Rect {
	return v.value
}

func (v *NilableRect) Set(val *Rect) {
	v.value = val
	v.isSet = true
}

func (v NilableRect) IsSet() bool {
	return v.isSet
}

func (v *NilableRect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNilableRect(val *Rect) *NilableRect {
	return &NilableRect{value: val, isSet: true}
}

func (v NilableRect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NilableRect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

