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

// PerfLoggingPrefs struct for PerfLoggingPrefs
type PerfLoggingPrefs struct {
	EnableNetwork *bool `json:"enableNetwork,omitempty"`
	EnablePage *bool `json:"enablePage,omitempty"`
	TraceCategories *string `json:"traceCategories,omitempty"`
	BufferUsageReportingInterval *int64 `json:"bufferUsageReportingInterval,omitempty"`
}

// NewPerfLoggingPrefs instantiates a new PerfLoggingPrefs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerfLoggingPrefs() *PerfLoggingPrefs {
	this := PerfLoggingPrefs{}
	var enableNetwork bool = true
	this.EnableNetwork = &enableNetwork
	var enablePage bool = true
	this.EnablePage = &enablePage
	var bufferUsageReportingInterval int64 = 1000
	this.BufferUsageReportingInterval = &bufferUsageReportingInterval
	return &this
}

// NewPerfLoggingPrefsWithDefaults instantiates a new PerfLoggingPrefs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerfLoggingPrefsWithDefaults() *PerfLoggingPrefs {
	this := PerfLoggingPrefs{}
	var enableNetwork bool = true
	this.EnableNetwork = &enableNetwork
	var enablePage bool = true
	this.EnablePage = &enablePage
	var bufferUsageReportingInterval int64 = 1000
	this.BufferUsageReportingInterval = &bufferUsageReportingInterval
	return &this
}

// GetEnableNetwork returns the EnableNetwork field value if set, zero value otherwise.
func (o *PerfLoggingPrefs) GetEnableNetwork() bool {
	if o == nil || o.EnableNetwork == nil {
		var ret bool
		return ret
	}
	return *o.EnableNetwork
}

// GetEnableNetworkOk returns a tuple with the EnableNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfLoggingPrefs) GetEnableNetworkOk() (*bool, bool) {
	if o == nil || o.EnableNetwork == nil {
		return nil, false
	}
	return o.EnableNetwork, true
}

// HasEnableNetwork returns a boolean if a field has been set.
func (o *PerfLoggingPrefs) HasEnableNetwork() bool {
	if o != nil && o.EnableNetwork != nil {
		return true
	}

	return false
}

// SetEnableNetwork gets a reference to the given bool and assigns it to the EnableNetwork field.
func (o *PerfLoggingPrefs) SetEnableNetwork(v bool) {
	o.EnableNetwork = &v
}

// GetEnablePage returns the EnablePage field value if set, zero value otherwise.
func (o *PerfLoggingPrefs) GetEnablePage() bool {
	if o == nil || o.EnablePage == nil {
		var ret bool
		return ret
	}
	return *o.EnablePage
}

// GetEnablePageOk returns a tuple with the EnablePage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfLoggingPrefs) GetEnablePageOk() (*bool, bool) {
	if o == nil || o.EnablePage == nil {
		return nil, false
	}
	return o.EnablePage, true
}

// HasEnablePage returns a boolean if a field has been set.
func (o *PerfLoggingPrefs) HasEnablePage() bool {
	if o != nil && o.EnablePage != nil {
		return true
	}

	return false
}

// SetEnablePage gets a reference to the given bool and assigns it to the EnablePage field.
func (o *PerfLoggingPrefs) SetEnablePage(v bool) {
	o.EnablePage = &v
}

// GetTraceCategories returns the TraceCategories field value if set, zero value otherwise.
func (o *PerfLoggingPrefs) GetTraceCategories() string {
	if o == nil || o.TraceCategories == nil {
		var ret string
		return ret
	}
	return *o.TraceCategories
}

// GetTraceCategoriesOk returns a tuple with the TraceCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfLoggingPrefs) GetTraceCategoriesOk() (*string, bool) {
	if o == nil || o.TraceCategories == nil {
		return nil, false
	}
	return o.TraceCategories, true
}

// HasTraceCategories returns a boolean if a field has been set.
func (o *PerfLoggingPrefs) HasTraceCategories() bool {
	if o != nil && o.TraceCategories != nil {
		return true
	}

	return false
}

// SetTraceCategories gets a reference to the given string and assigns it to the TraceCategories field.
func (o *PerfLoggingPrefs) SetTraceCategories(v string) {
	o.TraceCategories = &v
}

// GetBufferUsageReportingInterval returns the BufferUsageReportingInterval field value if set, zero value otherwise.
func (o *PerfLoggingPrefs) GetBufferUsageReportingInterval() int64 {
	if o == nil || o.BufferUsageReportingInterval == nil {
		var ret int64
		return ret
	}
	return *o.BufferUsageReportingInterval
}

// GetBufferUsageReportingIntervalOk returns a tuple with the BufferUsageReportingInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfLoggingPrefs) GetBufferUsageReportingIntervalOk() (*int64, bool) {
	if o == nil || o.BufferUsageReportingInterval == nil {
		return nil, false
	}
	return o.BufferUsageReportingInterval, true
}

// HasBufferUsageReportingInterval returns a boolean if a field has been set.
func (o *PerfLoggingPrefs) HasBufferUsageReportingInterval() bool {
	if o != nil && o.BufferUsageReportingInterval != nil {
		return true
	}

	return false
}

// SetBufferUsageReportingInterval gets a reference to the given int64 and assigns it to the BufferUsageReportingInterval field.
func (o *PerfLoggingPrefs) SetBufferUsageReportingInterval(v int64) {
	o.BufferUsageReportingInterval = &v
}

func (o PerfLoggingPrefs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EnableNetwork != nil {
		toSerialize["enableNetwork"] = o.EnableNetwork
	}
	if o.EnablePage != nil {
		toSerialize["enablePage"] = o.EnablePage
	}
	if o.TraceCategories != nil {
		toSerialize["traceCategories"] = o.TraceCategories
	}
	if o.BufferUsageReportingInterval != nil {
		toSerialize["bufferUsageReportingInterval"] = o.BufferUsageReportingInterval
	}
	return json.Marshal(toSerialize)
}

type NilablePerfLoggingPrefs struct {
	value *PerfLoggingPrefs
	isSet bool
}

func (v NilablePerfLoggingPrefs) Get() *PerfLoggingPrefs {
	return v.value
}

func (v *NilablePerfLoggingPrefs) Set(val *PerfLoggingPrefs) {
	v.value = val
	v.isSet = true
}

func (v NilablePerfLoggingPrefs) IsSet() bool {
	return v.isSet
}

func (v *NilablePerfLoggingPrefs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNilablePerfLoggingPrefs(val *PerfLoggingPrefs) *NilablePerfLoggingPrefs {
	return &NilablePerfLoggingPrefs{value: val, isSet: true}
}

func (v NilablePerfLoggingPrefs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NilablePerfLoggingPrefs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


