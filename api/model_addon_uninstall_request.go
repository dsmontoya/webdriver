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

// AddonUninstallRequest struct for AddonUninstallRequest
type AddonUninstallRequest struct {
	Id string `json:"id"`
}

// NewAddonUninstallRequest instantiates a new AddonUninstallRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddonUninstallRequest(id string) *AddonUninstallRequest {
	this := AddonUninstallRequest{}
	this.Id = id
	return &this
}

// NewAddonUninstallRequestWithDefaults instantiates a new AddonUninstallRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddonUninstallRequestWithDefaults() *AddonUninstallRequest {
	this := AddonUninstallRequest{}
	return &this
}

// GetId returns the Id field value
func (o *AddonUninstallRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AddonUninstallRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AddonUninstallRequest) SetId(v string) {
	o.Id = v
}

func (o AddonUninstallRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NilableAddonUninstallRequest struct {
	value *AddonUninstallRequest
	isSet bool
}

func (v NilableAddonUninstallRequest) Get() *AddonUninstallRequest {
	return v.value
}

func (v *NilableAddonUninstallRequest) Set(val *AddonUninstallRequest) {
	v.value = val
	v.isSet = true
}

func (v NilableAddonUninstallRequest) IsSet() bool {
	return v.isSet
}

func (v *NilableAddonUninstallRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNilableAddonUninstallRequest(val *AddonUninstallRequest) *NilableAddonUninstallRequest {
	return &NilableAddonUninstallRequest{value: val, isSet: true}
}

func (v NilableAddonUninstallRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NilableAddonUninstallRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


