# Proxy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProxyType** | **string** |  | 
**ProxyAutoconfigUrl** | Pointer to **string** |  | [optional] 
**FtpProxy** | Pointer to **string** |  | [optional] 
**HttpProxy** | Pointer to **string** |  | [optional] 
**NoProxy** | Pointer to **[]string** |  | [optional] 
**SslProxy** | Pointer to **string** |  | [optional] 
**SocksProxy** | Pointer to **string** |  | [optional] 
**SocksVersion** | Pointer to **int32** |  | [optional] 

## Methods

### NewProxy

`func NewProxy(proxyType string, ) *Proxy`

NewProxy instantiates a new Proxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyWithDefaults

`func NewProxyWithDefaults() *Proxy`

NewProxyWithDefaults instantiates a new Proxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProxyType

`func (o *Proxy) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *Proxy) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *Proxy) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.


### GetProxyAutoconfigUrl

`func (o *Proxy) GetProxyAutoconfigUrl() string`

GetProxyAutoconfigUrl returns the ProxyAutoconfigUrl field if non-nil, zero value otherwise.

### GetProxyAutoconfigUrlOk

`func (o *Proxy) GetProxyAutoconfigUrlOk() (*string, bool)`

GetProxyAutoconfigUrlOk returns a tuple with the ProxyAutoconfigUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAutoconfigUrl

`func (o *Proxy) SetProxyAutoconfigUrl(v string)`

SetProxyAutoconfigUrl sets ProxyAutoconfigUrl field to given value.

### HasProxyAutoconfigUrl

`func (o *Proxy) HasProxyAutoconfigUrl() bool`

HasProxyAutoconfigUrl returns a boolean if a field has been set.

### GetFtpProxy

`func (o *Proxy) GetFtpProxy() string`

GetFtpProxy returns the FtpProxy field if non-nil, zero value otherwise.

### GetFtpProxyOk

`func (o *Proxy) GetFtpProxyOk() (*string, bool)`

GetFtpProxyOk returns a tuple with the FtpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpProxy

`func (o *Proxy) SetFtpProxy(v string)`

SetFtpProxy sets FtpProxy field to given value.

### HasFtpProxy

`func (o *Proxy) HasFtpProxy() bool`

HasFtpProxy returns a boolean if a field has been set.

### GetHttpProxy

`func (o *Proxy) GetHttpProxy() string`

GetHttpProxy returns the HttpProxy field if non-nil, zero value otherwise.

### GetHttpProxyOk

`func (o *Proxy) GetHttpProxyOk() (*string, bool)`

GetHttpProxyOk returns a tuple with the HttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxy

`func (o *Proxy) SetHttpProxy(v string)`

SetHttpProxy sets HttpProxy field to given value.

### HasHttpProxy

`func (o *Proxy) HasHttpProxy() bool`

HasHttpProxy returns a boolean if a field has been set.

### GetNoProxy

`func (o *Proxy) GetNoProxy() []string`

GetNoProxy returns the NoProxy field if non-nil, zero value otherwise.

### GetNoProxyOk

`func (o *Proxy) GetNoProxyOk() (*[]string, bool)`

GetNoProxyOk returns a tuple with the NoProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxy

`func (o *Proxy) SetNoProxy(v []string)`

SetNoProxy sets NoProxy field to given value.

### HasNoProxy

`func (o *Proxy) HasNoProxy() bool`

HasNoProxy returns a boolean if a field has been set.

### GetSslProxy

`func (o *Proxy) GetSslProxy() string`

GetSslProxy returns the SslProxy field if non-nil, zero value otherwise.

### GetSslProxyOk

`func (o *Proxy) GetSslProxyOk() (*string, bool)`

GetSslProxyOk returns a tuple with the SslProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslProxy

`func (o *Proxy) SetSslProxy(v string)`

SetSslProxy sets SslProxy field to given value.

### HasSslProxy

`func (o *Proxy) HasSslProxy() bool`

HasSslProxy returns a boolean if a field has been set.

### GetSocksProxy

`func (o *Proxy) GetSocksProxy() string`

GetSocksProxy returns the SocksProxy field if non-nil, zero value otherwise.

### GetSocksProxyOk

`func (o *Proxy) GetSocksProxyOk() (*string, bool)`

GetSocksProxyOk returns a tuple with the SocksProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocksProxy

`func (o *Proxy) SetSocksProxy(v string)`

SetSocksProxy sets SocksProxy field to given value.

### HasSocksProxy

`func (o *Proxy) HasSocksProxy() bool`

HasSocksProxy returns a boolean if a field has been set.

### GetSocksVersion

`func (o *Proxy) GetSocksVersion() int32`

GetSocksVersion returns the SocksVersion field if non-nil, zero value otherwise.

### GetSocksVersionOk

`func (o *Proxy) GetSocksVersionOk() (*int32, bool)`

GetSocksVersionOk returns a tuple with the SocksVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocksVersion

`func (o *Proxy) SetSocksVersion(v int32)`

SetSocksVersion sets SocksVersion field to given value.

### HasSocksVersion

`func (o *Proxy) HasSocksVersion() bool`

HasSocksVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


