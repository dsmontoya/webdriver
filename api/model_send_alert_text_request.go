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

// SendAlertTextRequest struct for SendAlertTextRequest
type SendAlertTextRequest struct {
	Text string `json:"text"`
}

// NewSendAlertTextRequest instantiates a new SendAlertTextRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendAlertTextRequest(text string) *SendAlertTextRequest {
	this := SendAlertTextRequest{}
	this.Text = text
	return &this
}

// NewSendAlertTextRequestWithDefaults instantiates a new SendAlertTextRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendAlertTextRequestWithDefaults() *SendAlertTextRequest {
	this := SendAlertTextRequest{}
	return &this
}

// GetText returns the Text field value
func (o *SendAlertTextRequest) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *SendAlertTextRequest) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *SendAlertTextRequest) SetText(v string) {
	o.Text = v
}

func (o SendAlertTextRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NilableSendAlertTextRequest struct {
	value *SendAlertTextRequest
	isSet bool
}

func (v NilableSendAlertTextRequest) Get() *SendAlertTextRequest {
	return v.value
}

func (v *NilableSendAlertTextRequest) Set(val *SendAlertTextRequest) {
	v.value = val
	v.isSet = true
}

func (v NilableSendAlertTextRequest) IsSet() bool {
	return v.isSet
}

func (v *NilableSendAlertTextRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNilableSendAlertTextRequest(val *SendAlertTextRequest) *NilableSendAlertTextRequest {
	return &NilableSendAlertTextRequest{value: val, isSet: true}
}

func (v NilableSendAlertTextRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NilableSendAlertTextRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


