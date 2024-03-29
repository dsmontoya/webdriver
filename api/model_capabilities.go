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

// Capabilities struct for Capabilities
type Capabilities struct {
	BrowserName *string `json:"browserName,omitempty"`
	BrowserVersion *string `json:"browserVersion,omitempty"`
	PlatformName *string `json:"platformName,omitempty"`
	AcceptInsecureCerts *bool `json:"acceptInsecureCerts,omitempty"`
	PageLoadStrategy *string `json:"pageLoadStrategy,omitempty"`
	Proxy *Proxy `json:"proxy,omitempty"`
	SetWindowRect *bool `json:"setWindowRect,omitempty"`
	Timeouts *Timeouts `json:"timeouts,omitempty"`
	StrictFileInteractability *bool `json:"strictFileInteractability,omitempty"`
	UnhandledPromptBehavior *string `json:"unhandledPromptBehavior,omitempty"`
	GoogloggingPrefs *LoggingPrefs `json:"goog:loggingPrefs,omitempty"`
	GoogchromeOptions *ChromeOptions `json:"goog:chromeOptions,omitempty"`
	Moonoptions *MoonOptions `json:"moon:options,omitempty"`
	MozfirefoxOptions *FirefoxOptions `json:"moz:firefoxOptions,omitempty"`
	MsedgeOptions *EdgeOptions `json:"ms:edgeOptions,omitempty"`
	OperaOptions *EdgeOptions `json:"operaOptions,omitempty"`
	Selenoidoptions *SelenoidOptions `json:"selenoid:options,omitempty"`
	SafariautomaticInspection *bool `json:"safari:automaticInspection,omitempty"`
	SafariautomaticProfiling *bool `json:"safari:automaticProfiling,omitempty"`
}

// NewCapabilities instantiates a new Capabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilities() *Capabilities {
	this := Capabilities{}
	return &this
}

// NewCapabilitiesWithDefaults instantiates a new Capabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitiesWithDefaults() *Capabilities {
	this := Capabilities{}
	return &this
}

// GetBrowserName returns the BrowserName field value if set, zero value otherwise.
func (o *Capabilities) GetBrowserName() string {
	if o == nil || o.BrowserName == nil {
		var ret string
		return ret
	}
	return *o.BrowserName
}

// GetBrowserNameOk returns a tuple with the BrowserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capabilities) GetBrowserNameOk() (*string, bool) {
	if o == nil || o.BrowserName == nil {
		return nil, false
	}
	return o.BrowserName, true
}

// HasBrowserName returns a boolean if a field has been set.
func (o *Capabilities) HasBrowserName() bool {
	if o != nil && o.BrowserName != nil {
		return true
	}

	return false
}

// SetBrowserName gets a reference to the given string and assigns it to the BrowserName field.
func (o *Capabilities) SetBrowserName(v string) {
	o.BrowserName = &v
}

// GetBrowserVersion returns the BrowserVersion field value if set, zero value otherwise.
func (o *Capabilities) GetBrowserVersion() string {
	if o == nil || o.BrowserVersion == nil {
		var ret string
		return ret
	}
	return *o.BrowserVersion
}

// GetBrowserVersionOk returns a tuple with the BrowserVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capabilities) GetBrowserVersionOk() (*string, bool) {
	if o == nil || o.BrowserVersion == nil {
		return nil, false
	}
	return o.BrowserVersion, true
}

// HasBrowserVersion returns a boolean if a field has been set.
func (o *Capabilities) HasBrowserVersion() bool {
	if o != nil && o.BrowserVersion != nil {
		return true
	}

	return false
}

// SetBrowserVersion gets a reference to the given string and assigns it to the BrowserVersion field.
func (o *Capabilities) SetBrowserVersion(v string) {
	o.BrowserVersion = &v
}

// GetPlatformName returns the PlatformName field value if set, zero value otherwise.
func (o *Capabilities) GetPlatformName() string {
	if o == nil || o.PlatformName == nil {
		var ret string
		return ret
	}
	return *o.PlatformName
}

// GetPlatformNameOk returns a tuple with the PlatformName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capabilities) GetPlatformNameOk() (*string, bool) {
	if o == nil || o.PlatformName == nil {
		return nil, false
	}
	return o.PlatformName, true
}

// HasPlatformName returns a boolean if a field has been set.
func (o *Capabilities) HasPlatformName() bool {
	if o != nil && o.PlatformName != nil {
		return true
	}

	return false
}

// SetPlatformName gets a reference to the given string and assigns it to the PlatformName field.
func (o *Capabilities) SetPlatformName(v string) {
	o.PlatformName = &v
}

// GetAcceptInsecureCerts returns the AcceptInsecureCerts field value if set, zero value otherwise.
func (o *Capabilities) GetAcceptInsecureCerts() bool {
	if o == nil || o.AcceptInsecureCerts == nil {
		var ret bool
		return ret
	}
	return *o.AcceptInsecureCerts
}

// GetAcceptInsecureCertsOk returns a tuple with the AcceptInsecureCerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capabilities) GetAcceptInsecureCertsOk() (*bool, bool) {
	if o == nil || o.AcceptInsecureCerts == nil {
		return nil, false
	}
	return o.AcceptInsecureCerts, true
}

// HasAcceptInsecureCerts returns a boolean if a field has been set.
func (o *Capabilities) HasAcceptInsecureCerts() bool {
	if o != nil && o.AcceptInsecureCerts != nil {
		return true
	}

	return false
}

// SetAcceptInsecureCerts gets a reference to the given bool and assigns it to the AcceptInsecureCerts field.
func (o *Capabilities) SetAcceptInsecureCerts(v bool) {
	o.AcceptInsecureCerts = &v
}

// GetPageLoadStrategy returns the PageLoadStrategy field value if set, zero value otherwise.
func (o *Capabilities) GetPageLoadStrategy() string {
	if o == nil || o.PageLoadStrategy == nil {
		var ret string
		return ret
	}
	return *o.PageLoadStrategy
}

// GetPageLoadStrategyOk returns a tuple with the PageLoadStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capabilities) GetPageLoadStrategyOk() (*string, bool) {
	if o == nil || o.PageLoadStrategy == nil {
		return nil, false
	}
	return o.PageLoadStrategy, true
}

// HasPageLoadStrategy returns a boolean if a field has been set.
func (o *Capabilities) HasPageLoadStrategy() bool {
	if o != nil && o.PageLoadStrategy != nil {
		return true
	}

	return false
}

// SetPageLoadStrategy gets a reference to the given string and assigns it to the PageLoadStrategy field.
func (o *Capabilities) SetPageLoadStrategy(v string) {
	o.PageLoadStrategy = &v
}

// GetProxy returns the Proxy field value if set, zero value otherwise.
func (o *Capabilities) GetProxy() Proxy {
	if o == nil || o.Proxy == nil {
		var ret Proxy
		return ret
	}
	return *o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capabilities) GetProxyOk() (*Proxy, bool) {
	if o == nil || o.Proxy == nil {
		return nil, false
	}
	return o.Proxy, true
}

// HasProxy returns a boolean if a field has been set.
func (o *Capabilities) HasProxy() bool {
	if o != nil && o.Proxy != nil {
		return true
	}

	return false
}

// SetProxy gets a reference to the given Proxy and assigns it to the Proxy field.
func (o *Capabilities) SetProxy(v Proxy) {
	o.Proxy = &v
}

// GetSetWindowRect returns the SetWindowRect field value if set, zero value otherwise.
func (o *Capabilities) GetSetWindowRect() bool {
	if o == nil || o.SetWindowRect == nil {
		var ret bool
		return ret
	}
	return *o.SetWindowRect
}

// GetSetWindowRectOk returns a tuple with the SetWindowRect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capabilities) GetSetWindowRectOk() (*bool, bool) {
	if o == nil || o.SetWindowRect == nil {
		return nil, false
	}
	return o.SetWindowRect, true
}

// HasSetWindowRect returns a boolean if a field has been set.
func (o *Capabilities) HasSetWindowRect() bool {
	if o != nil && o.SetWindowRect != nil {
		return true
	}

	return false
}

// SetSetWindowRect gets a reference to the given bool and assigns it to the SetWindowRect field.
func (o *Capabilities) SetSetWindowRect(v bool) {
	o.SetWindowRect = &v
}

// GetTimeouts returns the Timeouts field value if set, zero value otherwise.
func (o *Capabilities) GetTimeouts() Timeouts {
	if o == nil || o.Timeouts == nil {
		var ret Timeouts
		return ret
	}
	return *o.Timeouts
}

// GetTimeoutsOk returns a tuple with the Timeouts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capabilities) GetTimeoutsOk() (*Timeouts, bool) {
	if o == nil || o.Timeouts == nil {
		return nil, false
	}
	return o.Timeouts, true
}

// HasTimeouts returns a boolean if a field has been set.
func (o *Capabilities) HasTimeouts() bool {
	if o != nil && o.Timeouts != nil {
		return true
	}

	return false
}

// SetTimeouts gets a reference to the given Timeouts and assigns it to the Timeouts field.
func (o *Capabilities) SetTimeouts(v Timeouts) {
	o.Timeouts = &v
}

// GetStrictFileInteractability returns the StrictFileInteractability field value if set, zero value otherwise.
func (o *Capabilities) GetStrictFileInteractability() bool {
	if o == nil || o.StrictFileInteractability == nil {
		var ret bool
		return ret
	}
	return *o.StrictFileInteractability
}

// GetStrictFileInteractabilityOk returns a tuple with the StrictFileInteractability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capabilities) GetStrictFileInteractabilityOk() (*bool, bool) {
	if o == nil || o.StrictFileInteractability == nil {
		return nil, false
	}
	return o.StrictFileInteractability, true
}

// HasStrictFileInteractability returns a boolean if a field has been set.
func (o *Capabilities) HasStrictFileInteractability() bool {
	if o != nil && o.StrictFileInteractability != nil {
		return true
	}

	return false
}

// SetStrictFileInteractability gets a reference to the given bool and assigns it to the StrictFileInteractability field.
func (o *Capabilities) SetStrictFileInteractability(v bool) {
	o.StrictFileInteractability = &v
}

// GetUnhandledPromptBehavior returns the UnhandledPromptBehavior field value if set, zero value otherwise.
func (o *Capabilities) GetUnhandledPromptBehavior() string {
	if o == nil || o.UnhandledPromptBehavior == nil {
		var ret string
		return ret
	}
	return *o.UnhandledPromptBehavior
}

// GetUnhandledPromptBehaviorOk returns a tuple with the UnhandledPromptBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capabilities) GetUnhandledPromptBehaviorOk() (*string, bool) {
	if o == nil || o.UnhandledPromptBehavior == nil {
		return nil, false
	}
	return o.UnhandledPromptBehavior, true
}

// HasUnhandledPromptBehavior returns a boolean if a field has been set.
func (o *Capabilities) HasUnhandledPromptBehavior() bool {
	if o != nil && o.UnhandledPromptBehavior != nil {
		return true
	}

	return false
}

// SetUnhandledPromptBehavior gets a reference to the given string and assigns it to the UnhandledPromptBehavior field.
func (o *Capabilities) SetUnhandledPromptBehavior(v string) {
	o.UnhandledPromptBehavior = &v
}

// GetGoogloggingPrefs returns the GoogloggingPrefs field value if set, zero value otherwise.
func (o *Capabilities) GetGoogloggingPrefs() LoggingPrefs {
	if o == nil || o.GoogloggingPrefs == nil {
		var ret LoggingPrefs
		return ret
	}
	return *o.GoogloggingPrefs
}

// GetGoogloggingPrefsOk returns a tuple with the GoogloggingPrefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capabilities) GetGoogloggingPrefsOk() (*LoggingPrefs, bool) {
	if o == nil || o.GoogloggingPrefs == nil {
		return nil, false
	}
	return o.GoogloggingPrefs, true
}

// HasGoogloggingPrefs returns a boolean if a field has been set.
func (o *Capabilities) HasGoogloggingPrefs() bool {
	if o != nil && o.GoogloggingPrefs != nil {
		return true
	}

	return false
}

// SetGoogloggingPrefs gets a reference to the given LoggingPrefs and assigns it to the GoogloggingPrefs field.
func (o *Capabilities) SetGoogloggingPrefs(v LoggingPrefs) {
	o.GoogloggingPrefs = &v
}

// GetGoogchromeOptions returns the GoogchromeOptions field value if set, zero value otherwise.
func (o *Capabilities) GetGoogchromeOptions() ChromeOptions {
	if o == nil || o.GoogchromeOptions == nil {
		var ret ChromeOptions
		return ret
	}
	return *o.GoogchromeOptions
}

// GetGoogchromeOptionsOk returns a tuple with the GoogchromeOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capabilities) GetGoogchromeOptionsOk() (*ChromeOptions, bool) {
	if o == nil || o.GoogchromeOptions == nil {
		return nil, false
	}
	return o.GoogchromeOptions, true
}

// HasGoogchromeOptions returns a boolean if a field has been set.
func (o *Capabilities) HasGoogchromeOptions() bool {
	if o != nil && o.GoogchromeOptions != nil {
		return true
	}

	return false
}

// SetGoogchromeOptions gets a reference to the given ChromeOptions and assigns it to the GoogchromeOptions field.
func (o *Capabilities) SetGoogchromeOptions(v ChromeOptions) {
	o.GoogchromeOptions = &v
}

// GetMoonoptions returns the Moonoptions field value if set, zero value otherwise.
func (o *Capabilities) GetMoonoptions() MoonOptions {
	if o == nil || o.Moonoptions == nil {
		var ret MoonOptions
		return ret
	}
	return *o.Moonoptions
}

// GetMoonoptionsOk returns a tuple with the Moonoptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capabilities) GetMoonoptionsOk() (*MoonOptions, bool) {
	if o == nil || o.Moonoptions == nil {
		return nil, false
	}
	return o.Moonoptions, true
}

// HasMoonoptions returns a boolean if a field has been set.
func (o *Capabilities) HasMoonoptions() bool {
	if o != nil && o.Moonoptions != nil {
		return true
	}

	return false
}

// SetMoonoptions gets a reference to the given MoonOptions and assigns it to the Moonoptions field.
func (o *Capabilities) SetMoonoptions(v MoonOptions) {
	o.Moonoptions = &v
}

// GetMozfirefoxOptions returns the MozfirefoxOptions field value if set, zero value otherwise.
func (o *Capabilities) GetMozfirefoxOptions() FirefoxOptions {
	if o == nil || o.MozfirefoxOptions == nil {
		var ret FirefoxOptions
		return ret
	}
	return *o.MozfirefoxOptions
}

// GetMozfirefoxOptionsOk returns a tuple with the MozfirefoxOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capabilities) GetMozfirefoxOptionsOk() (*FirefoxOptions, bool) {
	if o == nil || o.MozfirefoxOptions == nil {
		return nil, false
	}
	return o.MozfirefoxOptions, true
}

// HasMozfirefoxOptions returns a boolean if a field has been set.
func (o *Capabilities) HasMozfirefoxOptions() bool {
	if o != nil && o.MozfirefoxOptions != nil {
		return true
	}

	return false
}

// SetMozfirefoxOptions gets a reference to the given FirefoxOptions and assigns it to the MozfirefoxOptions field.
func (o *Capabilities) SetMozfirefoxOptions(v FirefoxOptions) {
	o.MozfirefoxOptions = &v
}

// GetMsedgeOptions returns the MsedgeOptions field value if set, zero value otherwise.
func (o *Capabilities) GetMsedgeOptions() EdgeOptions {
	if o == nil || o.MsedgeOptions == nil {
		var ret EdgeOptions
		return ret
	}
	return *o.MsedgeOptions
}

// GetMsedgeOptionsOk returns a tuple with the MsedgeOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capabilities) GetMsedgeOptionsOk() (*EdgeOptions, bool) {
	if o == nil || o.MsedgeOptions == nil {
		return nil, false
	}
	return o.MsedgeOptions, true
}

// HasMsedgeOptions returns a boolean if a field has been set.
func (o *Capabilities) HasMsedgeOptions() bool {
	if o != nil && o.MsedgeOptions != nil {
		return true
	}

	return false
}

// SetMsedgeOptions gets a reference to the given EdgeOptions and assigns it to the MsedgeOptions field.
func (o *Capabilities) SetMsedgeOptions(v EdgeOptions) {
	o.MsedgeOptions = &v
}

// GetOperaOptions returns the OperaOptions field value if set, zero value otherwise.
func (o *Capabilities) GetOperaOptions() EdgeOptions {
	if o == nil || o.OperaOptions == nil {
		var ret EdgeOptions
		return ret
	}
	return *o.OperaOptions
}

// GetOperaOptionsOk returns a tuple with the OperaOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capabilities) GetOperaOptionsOk() (*EdgeOptions, bool) {
	if o == nil || o.OperaOptions == nil {
		return nil, false
	}
	return o.OperaOptions, true
}

// HasOperaOptions returns a boolean if a field has been set.
func (o *Capabilities) HasOperaOptions() bool {
	if o != nil && o.OperaOptions != nil {
		return true
	}

	return false
}

// SetOperaOptions gets a reference to the given EdgeOptions and assigns it to the OperaOptions field.
func (o *Capabilities) SetOperaOptions(v EdgeOptions) {
	o.OperaOptions = &v
}

// GetSelenoidoptions returns the Selenoidoptions field value if set, zero value otherwise.
func (o *Capabilities) GetSelenoidoptions() SelenoidOptions {
	if o == nil || o.Selenoidoptions == nil {
		var ret SelenoidOptions
		return ret
	}
	return *o.Selenoidoptions
}

// GetSelenoidoptionsOk returns a tuple with the Selenoidoptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capabilities) GetSelenoidoptionsOk() (*SelenoidOptions, bool) {
	if o == nil || o.Selenoidoptions == nil {
		return nil, false
	}
	return o.Selenoidoptions, true
}

// HasSelenoidoptions returns a boolean if a field has been set.
func (o *Capabilities) HasSelenoidoptions() bool {
	if o != nil && o.Selenoidoptions != nil {
		return true
	}

	return false
}

// SetSelenoidoptions gets a reference to the given SelenoidOptions and assigns it to the Selenoidoptions field.
func (o *Capabilities) SetSelenoidoptions(v SelenoidOptions) {
	o.Selenoidoptions = &v
}

// GetSafariautomaticInspection returns the SafariautomaticInspection field value if set, zero value otherwise.
func (o *Capabilities) GetSafariautomaticInspection() bool {
	if o == nil || o.SafariautomaticInspection == nil {
		var ret bool
		return ret
	}
	return *o.SafariautomaticInspection
}

// GetSafariautomaticInspectionOk returns a tuple with the SafariautomaticInspection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capabilities) GetSafariautomaticInspectionOk() (*bool, bool) {
	if o == nil || o.SafariautomaticInspection == nil {
		return nil, false
	}
	return o.SafariautomaticInspection, true
}

// HasSafariautomaticInspection returns a boolean if a field has been set.
func (o *Capabilities) HasSafariautomaticInspection() bool {
	if o != nil && o.SafariautomaticInspection != nil {
		return true
	}

	return false
}

// SetSafariautomaticInspection gets a reference to the given bool and assigns it to the SafariautomaticInspection field.
func (o *Capabilities) SetSafariautomaticInspection(v bool) {
	o.SafariautomaticInspection = &v
}

// GetSafariautomaticProfiling returns the SafariautomaticProfiling field value if set, zero value otherwise.
func (o *Capabilities) GetSafariautomaticProfiling() bool {
	if o == nil || o.SafariautomaticProfiling == nil {
		var ret bool
		return ret
	}
	return *o.SafariautomaticProfiling
}

// GetSafariautomaticProfilingOk returns a tuple with the SafariautomaticProfiling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capabilities) GetSafariautomaticProfilingOk() (*bool, bool) {
	if o == nil || o.SafariautomaticProfiling == nil {
		return nil, false
	}
	return o.SafariautomaticProfiling, true
}

// HasSafariautomaticProfiling returns a boolean if a field has been set.
func (o *Capabilities) HasSafariautomaticProfiling() bool {
	if o != nil && o.SafariautomaticProfiling != nil {
		return true
	}

	return false
}

// SetSafariautomaticProfiling gets a reference to the given bool and assigns it to the SafariautomaticProfiling field.
func (o *Capabilities) SetSafariautomaticProfiling(v bool) {
	o.SafariautomaticProfiling = &v
}

func (o Capabilities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BrowserName != nil {
		toSerialize["browserName"] = o.BrowserName
	}
	if o.BrowserVersion != nil {
		toSerialize["browserVersion"] = o.BrowserVersion
	}
	if o.PlatformName != nil {
		toSerialize["platformName"] = o.PlatformName
	}
	if o.AcceptInsecureCerts != nil {
		toSerialize["acceptInsecureCerts"] = o.AcceptInsecureCerts
	}
	if o.PageLoadStrategy != nil {
		toSerialize["pageLoadStrategy"] = o.PageLoadStrategy
	}
	if o.Proxy != nil {
		toSerialize["proxy"] = o.Proxy
	}
	if o.SetWindowRect != nil {
		toSerialize["setWindowRect"] = o.SetWindowRect
	}
	if o.Timeouts != nil {
		toSerialize["timeouts"] = o.Timeouts
	}
	if o.StrictFileInteractability != nil {
		toSerialize["strictFileInteractability"] = o.StrictFileInteractability
	}
	if o.UnhandledPromptBehavior != nil {
		toSerialize["unhandledPromptBehavior"] = o.UnhandledPromptBehavior
	}
	if o.GoogloggingPrefs != nil {
		toSerialize["goog:loggingPrefs"] = o.GoogloggingPrefs
	}
	if o.GoogchromeOptions != nil {
		toSerialize["goog:chromeOptions"] = o.GoogchromeOptions
	}
	if o.Moonoptions != nil {
		toSerialize["moon:options"] = o.Moonoptions
	}
	if o.MozfirefoxOptions != nil {
		toSerialize["moz:firefoxOptions"] = o.MozfirefoxOptions
	}
	if o.MsedgeOptions != nil {
		toSerialize["ms:edgeOptions"] = o.MsedgeOptions
	}
	if o.OperaOptions != nil {
		toSerialize["operaOptions"] = o.OperaOptions
	}
	if o.Selenoidoptions != nil {
		toSerialize["selenoid:options"] = o.Selenoidoptions
	}
	if o.SafariautomaticInspection != nil {
		toSerialize["safari:automaticInspection"] = o.SafariautomaticInspection
	}
	if o.SafariautomaticProfiling != nil {
		toSerialize["safari:automaticProfiling"] = o.SafariautomaticProfiling
	}
	return json.Marshal(toSerialize)
}

type NilableCapabilities struct {
	value *Capabilities
	isSet bool
}

func (v NilableCapabilities) Get() *Capabilities {
	return v.value
}

func (v *NilableCapabilities) Set(val *Capabilities) {
	v.value = val
	v.isSet = true
}

func (v NilableCapabilities) IsSet() bool {
	return v.isSet
}

func (v *NilableCapabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNilableCapabilities(val *Capabilities) *NilableCapabilities {
	return &NilableCapabilities{value: val, isSet: true}
}

func (v NilableCapabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NilableCapabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


