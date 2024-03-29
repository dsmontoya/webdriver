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

// ActionSequence struct for ActionSequence
type ActionSequence struct {
	Id *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	Parameters *ActionSequenceParameters `json:"parameters,omitempty"`
	Actions []Action `json:"actions,omitempty"`
}

// NewActionSequence instantiates a new ActionSequence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionSequence() *ActionSequence {
	this := ActionSequence{}
	return &this
}

// NewActionSequenceWithDefaults instantiates a new ActionSequence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionSequenceWithDefaults() *ActionSequence {
	this := ActionSequence{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ActionSequence) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionSequence) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ActionSequence) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ActionSequence) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ActionSequence) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionSequence) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ActionSequence) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ActionSequence) SetType(v string) {
	o.Type = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ActionSequence) GetParameters() ActionSequenceParameters {
	if o == nil || o.Parameters == nil {
		var ret ActionSequenceParameters
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionSequence) GetParametersOk() (*ActionSequenceParameters, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ActionSequence) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given ActionSequenceParameters and assigns it to the Parameters field.
func (o *ActionSequence) SetParameters(v ActionSequenceParameters) {
	o.Parameters = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *ActionSequence) GetActions() []Action {
	if o == nil || o.Actions == nil {
		var ret []Action
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionSequence) GetActionsOk() ([]Action, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *ActionSequence) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []Action and assigns it to the Actions field.
func (o *ActionSequence) SetActions(v []Action) {
	o.Actions = v
}

func (o ActionSequence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	return json.Marshal(toSerialize)
}

type NilableActionSequence struct {
	value *ActionSequence
	isSet bool
}

func (v NilableActionSequence) Get() *ActionSequence {
	return v.value
}

func (v *NilableActionSequence) Set(val *ActionSequence) {
	v.value = val
	v.isSet = true
}

func (v NilableActionSequence) IsSet() bool {
	return v.isSet
}

func (v *NilableActionSequence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNilableActionSequence(val *ActionSequence) *NilableActionSequence {
	return &NilableActionSequence{value: val, isSet: true}
}

func (v NilableActionSequence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NilableActionSequence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


