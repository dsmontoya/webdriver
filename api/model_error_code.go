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
	"fmt"
)

// ErrorCode the model 'ErrorCode'
type ErrorCode string

// List of ErrorCode
const (
	ERRORCODE_ELEMENT_CLICK_INTERCEPTED ErrorCode = "element click intercepted"
	ERRORCODE_ELEMENT_NOT_INTERACTABLE ErrorCode = "element not interactable"
	ERRORCODE_INSECURE_CERTIFICATE ErrorCode = "insecure certificate"
	ERRORCODE_INVALID_ARGUMENT ErrorCode = "invalid argument"
	ERRORCODE_INVALID_COOKIE_DOMAIN ErrorCode = "invalid cookie domain"
	ERRORCODE_INVALID_ELEMENT_STATE ErrorCode = "invalid element state"
	ERRORCODE_INVALID_SELECTOR ErrorCode = "invalid selector"
	ERRORCODE_INVALID_SESSION_ID ErrorCode = "invalid session id"
	ERRORCODE_JAVASCRIPT_ERROR ErrorCode = "javascript error"
	ERRORCODE_MOVE_TARGET_OUT_OF_BOUNDS ErrorCode = "move target out of bounds"
	ERRORCODE_NO_SUCH_ALERT ErrorCode = "no such alert"
	ERRORCODE_NO_SUCH_COOKIE ErrorCode = "no such cookie"
	ERRORCODE_NO_SUCH_ELEMENT ErrorCode = "no such element"
	ERRORCODE_NO_SUCH_FRAME ErrorCode = "no such frame"
	ERRORCODE_NO_SUCH_WINDOW ErrorCode = "no such window"
	ERRORCODE_SCRIPT_TIMEOUT ErrorCode = "script timeout"
	ERRORCODE_SESSION_NOT_CREATED ErrorCode = "session not created"
	ERRORCODE_STALE_ELEMENT_REFERENCE ErrorCode = "stale element reference"
	ERRORCODE_TIMEOUT ErrorCode = "timeout"
	ERRORCODE_UNABLE_TO_SET_COOKIE ErrorCode = "unable to set cookie"
	ERRORCODE_UNABLE_TO_CAPTURE_SCREEN ErrorCode = "unable to capture screen"
	ERRORCODE_UNEXPECTED_ALERT_OPEN ErrorCode = "unexpected alert open"
	ERRORCODE_UNKNOWN_COMMAND ErrorCode = "unknown command"
	ERRORCODE_UNKNOWN_ERROR ErrorCode = "unknown error"
	ERRORCODE_UNKNOWN_METHOD ErrorCode = "unknown method"
	ERRORCODE_UNSUPPORTED_OPERATION ErrorCode = "unsupported operation"
)

// All allowed values of ErrorCode enum
var AllowedErrorCodeEnumValues = []ErrorCode{
	"element click intercepted",
	"element not interactable",
	"insecure certificate",
	"invalid argument",
	"invalid cookie domain",
	"invalid element state",
	"invalid selector",
	"invalid session id",
	"javascript error",
	"move target out of bounds",
	"no such alert",
	"no such cookie",
	"no such element",
	"no such frame",
	"no such window",
	"script timeout",
	"session not created",
	"stale element reference",
	"timeout",
	"unable to set cookie",
	"unable to capture screen",
	"unexpected alert open",
	"unknown command",
	"unknown error",
	"unknown method",
	"unsupported operation",
}

func (v *ErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ErrorCode(value)
	for _, existing := range AllowedErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ErrorCode", value)
}

// NewErrorCodeFromValue returns a pointer to a valid ErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewErrorCodeFromValue(v string) (*ErrorCode, error) {
	ev := ErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ErrorCode: valid values are %v", v, AllowedErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ErrorCode) IsValid() bool {
	for _, existing := range AllowedErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ErrorCode value
func (v ErrorCode) Ptr() *ErrorCode {
	return &v
}

type NullableErrorCode struct {
	value *ErrorCode
	isSet bool
}

func (v NullableErrorCode) Get() *ErrorCode {
	return v.value
}

func (v *NullableErrorCode) Set(val *ErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorCode(val *ErrorCode) *NullableErrorCode {
	return &NullableErrorCode{value: val, isSet: true}
}

func (v NullableErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

