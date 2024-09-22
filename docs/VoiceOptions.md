# VoiceOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Voice** | Pointer to **string** |  | [optional] [default to "default"]
**Styles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVoiceOptions

`func NewVoiceOptions() *VoiceOptions`

NewVoiceOptions instantiates a new VoiceOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceOptionsWithDefaults

`func NewVoiceOptionsWithDefaults() *VoiceOptions`

NewVoiceOptionsWithDefaults instantiates a new VoiceOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVoice

`func (o *VoiceOptions) GetVoice() string`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *VoiceOptions) GetVoiceOk() (*string, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *VoiceOptions) SetVoice(v string)`

SetVoice sets Voice field to given value.

### HasVoice

`func (o *VoiceOptions) HasVoice() bool`

HasVoice returns a boolean if a field has been set.

### GetStyles

`func (o *VoiceOptions) GetStyles() []string`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *VoiceOptions) GetStylesOk() (*[]string, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *VoiceOptions) SetStyles(v []string)`

SetStyles sets Styles field to given value.

### HasStyles

`func (o *VoiceOptions) HasStyles() bool`

HasStyles returns a boolean if a field has been set.

### SetStylesNil

`func (o *VoiceOptions) SetStylesNil(b bool)`

 SetStylesNil sets the value for Styles to be an explicit nil

### UnsetStyles
`func (o *VoiceOptions) UnsetStyles()`

UnsetStyles ensures that no value is present for Styles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


