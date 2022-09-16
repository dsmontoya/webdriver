# PrintRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orientation** | **string** |  | [default to "portrait"]
**Scale** | **float32** |  | [default to 1.0]
**Background** | **bool** | print page background | [default to false]
**Page** | [**PrintRequestOptionsPage**](PrintRequestOptionsPage.md) |  | 
**Margin** | [**PrintRequestOptionsMargin**](PrintRequestOptionsMargin.md) |  | 
**ShrinkToFit** | **bool** |  | [default to true]
**PageRanges** | **[]string** |  | 

## Methods

### NewPrintRequestOptions

`func NewPrintRequestOptions(orientation string, scale float32, background bool, page PrintRequestOptionsPage, margin PrintRequestOptionsMargin, shrinkToFit bool, pageRanges []string, ) *PrintRequestOptions`

NewPrintRequestOptions instantiates a new PrintRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrintRequestOptionsWithDefaults

`func NewPrintRequestOptionsWithDefaults() *PrintRequestOptions`

NewPrintRequestOptionsWithDefaults instantiates a new PrintRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrientation

`func (o *PrintRequestOptions) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *PrintRequestOptions) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *PrintRequestOptions) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.


### GetScale

`func (o *PrintRequestOptions) GetScale() float32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *PrintRequestOptions) GetScaleOk() (*float32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *PrintRequestOptions) SetScale(v float32)`

SetScale sets Scale field to given value.


### GetBackground

`func (o *PrintRequestOptions) GetBackground() bool`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *PrintRequestOptions) GetBackgroundOk() (*bool, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *PrintRequestOptions) SetBackground(v bool)`

SetBackground sets Background field to given value.


### GetPage

`func (o *PrintRequestOptions) GetPage() PrintRequestOptionsPage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PrintRequestOptions) GetPageOk() (*PrintRequestOptionsPage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PrintRequestOptions) SetPage(v PrintRequestOptionsPage)`

SetPage sets Page field to given value.


### GetMargin

`func (o *PrintRequestOptions) GetMargin() PrintRequestOptionsMargin`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *PrintRequestOptions) GetMarginOk() (*PrintRequestOptionsMargin, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *PrintRequestOptions) SetMargin(v PrintRequestOptionsMargin)`

SetMargin sets Margin field to given value.


### GetShrinkToFit

`func (o *PrintRequestOptions) GetShrinkToFit() bool`

GetShrinkToFit returns the ShrinkToFit field if non-nil, zero value otherwise.

### GetShrinkToFitOk

`func (o *PrintRequestOptions) GetShrinkToFitOk() (*bool, bool)`

GetShrinkToFitOk returns a tuple with the ShrinkToFit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShrinkToFit

`func (o *PrintRequestOptions) SetShrinkToFit(v bool)`

SetShrinkToFit sets ShrinkToFit field to given value.


### GetPageRanges

`func (o *PrintRequestOptions) GetPageRanges() []string`

GetPageRanges returns the PageRanges field if non-nil, zero value otherwise.

### GetPageRangesOk

`func (o *PrintRequestOptions) GetPageRangesOk() (*[]string, bool)`

GetPageRangesOk returns a tuple with the PageRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageRanges

`func (o *PrintRequestOptions) SetPageRanges(v []string)`

SetPageRanges sets PageRanges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


