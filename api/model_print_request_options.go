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

// PrintRequestOptions struct for PrintRequestOptions
type PrintRequestOptions struct {
	Orientation string `json:"orientation"`
	Scale float32 `json:"scale"`
	// print page background
	Background bool `json:"background"`
	Page PrintRequestOptionsPage `json:"page"`
	Margin PrintRequestOptionsMargin `json:"margin"`
	ShrinkToFit bool `json:"shrinkToFit"`
	PageRanges []string `json:"pageRanges"`
}

// NewPrintRequestOptions instantiates a new PrintRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrintRequestOptions(orientation string, scale float32, background bool, page PrintRequestOptionsPage, margin PrintRequestOptionsMargin, shrinkToFit bool, pageRanges []string) *PrintRequestOptions {
	this := PrintRequestOptions{}
	this.Orientation = orientation
	this.Scale = scale
	this.Background = background
	this.Page = page
	this.Margin = margin
	this.ShrinkToFit = shrinkToFit
	this.PageRanges = pageRanges
	return &this
}

// NewPrintRequestOptionsWithDefaults instantiates a new PrintRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrintRequestOptionsWithDefaults() *PrintRequestOptions {
	this := PrintRequestOptions{}
	var orientation string = "portrait"
	this.Orientation = orientation
	var scale float32 = 1.0
	this.Scale = scale
	var background bool = false
	this.Background = background
	var shrinkToFit bool = true
	this.ShrinkToFit = shrinkToFit
	return &this
}

// GetOrientation returns the Orientation field value
func (o *PrintRequestOptions) GetOrientation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value
// and a boolean to check if the value has been set.
func (o *PrintRequestOptions) GetOrientationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Orientation, true
}

// SetOrientation sets field value
func (o *PrintRequestOptions) SetOrientation(v string) {
	o.Orientation = v
}

// GetScale returns the Scale field value
func (o *PrintRequestOptions) GetScale() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Scale
}

// GetScaleOk returns a tuple with the Scale field value
// and a boolean to check if the value has been set.
func (o *PrintRequestOptions) GetScaleOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scale, true
}

// SetScale sets field value
func (o *PrintRequestOptions) SetScale(v float32) {
	o.Scale = v
}

// GetBackground returns the Background field value
func (o *PrintRequestOptions) GetBackground() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Background
}

// GetBackgroundOk returns a tuple with the Background field value
// and a boolean to check if the value has been set.
func (o *PrintRequestOptions) GetBackgroundOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Background, true
}

// SetBackground sets field value
func (o *PrintRequestOptions) SetBackground(v bool) {
	o.Background = v
}

// GetPage returns the Page field value
func (o *PrintRequestOptions) GetPage() PrintRequestOptionsPage {
	if o == nil {
		var ret PrintRequestOptionsPage
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *PrintRequestOptions) GetPageOk() (*PrintRequestOptionsPage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *PrintRequestOptions) SetPage(v PrintRequestOptionsPage) {
	o.Page = v
}

// GetMargin returns the Margin field value
func (o *PrintRequestOptions) GetMargin() PrintRequestOptionsMargin {
	if o == nil {
		var ret PrintRequestOptionsMargin
		return ret
	}

	return o.Margin
}

// GetMarginOk returns a tuple with the Margin field value
// and a boolean to check if the value has been set.
func (o *PrintRequestOptions) GetMarginOk() (*PrintRequestOptionsMargin, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Margin, true
}

// SetMargin sets field value
func (o *PrintRequestOptions) SetMargin(v PrintRequestOptionsMargin) {
	o.Margin = v
}

// GetShrinkToFit returns the ShrinkToFit field value
func (o *PrintRequestOptions) GetShrinkToFit() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShrinkToFit
}

// GetShrinkToFitOk returns a tuple with the ShrinkToFit field value
// and a boolean to check if the value has been set.
func (o *PrintRequestOptions) GetShrinkToFitOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShrinkToFit, true
}

// SetShrinkToFit sets field value
func (o *PrintRequestOptions) SetShrinkToFit(v bool) {
	o.ShrinkToFit = v
}

// GetPageRanges returns the PageRanges field value
func (o *PrintRequestOptions) GetPageRanges() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PageRanges
}

// GetPageRangesOk returns a tuple with the PageRanges field value
// and a boolean to check if the value has been set.
func (o *PrintRequestOptions) GetPageRangesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PageRanges, true
}

// SetPageRanges sets field value
func (o *PrintRequestOptions) SetPageRanges(v []string) {
	o.PageRanges = v
}

func (o PrintRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["orientation"] = o.Orientation
	}
	if true {
		toSerialize["scale"] = o.Scale
	}
	if true {
		toSerialize["background"] = o.Background
	}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["margin"] = o.Margin
	}
	if true {
		toSerialize["shrinkToFit"] = o.ShrinkToFit
	}
	if true {
		toSerialize["pageRanges"] = o.PageRanges
	}
	return json.Marshal(toSerialize)
}

type NilablePrintRequestOptions struct {
	value *PrintRequestOptions
	isSet bool
}

func (v NilablePrintRequestOptions) Get() *PrintRequestOptions {
	return v.value
}

func (v *NilablePrintRequestOptions) Set(val *PrintRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NilablePrintRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NilablePrintRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNilablePrintRequestOptions(val *PrintRequestOptions) *NilablePrintRequestOptions {
	return &NilablePrintRequestOptions{value: val, isSet: true}
}

func (v NilablePrintRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NilablePrintRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


