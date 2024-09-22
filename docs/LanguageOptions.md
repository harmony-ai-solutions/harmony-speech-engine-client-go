# LanguageOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | Pointer to **string** |  | [optional] [default to "default"]
**Voices** | Pointer to [**[]VoiceOptions**](VoiceOptions.md) |  | [optional] 

## Methods

### NewLanguageOptions

`func NewLanguageOptions() *LanguageOptions`

NewLanguageOptions instantiates a new LanguageOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanguageOptionsWithDefaults

`func NewLanguageOptionsWithDefaults() *LanguageOptions`

NewLanguageOptionsWithDefaults instantiates a new LanguageOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *LanguageOptions) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *LanguageOptions) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *LanguageOptions) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *LanguageOptions) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetVoices

`func (o *LanguageOptions) GetVoices() []VoiceOptions`

GetVoices returns the Voices field if non-nil, zero value otherwise.

### GetVoicesOk

`func (o *LanguageOptions) GetVoicesOk() (*[]VoiceOptions, bool)`

GetVoicesOk returns a tuple with the Voices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoices

`func (o *LanguageOptions) SetVoices(v []VoiceOptions)`

SetVoices sets Voices field to given value.

### HasVoices

`func (o *LanguageOptions) HasVoices() bool`

HasVoices returns a boolean if a field has been set.

### SetVoicesNil

`func (o *LanguageOptions) SetVoicesNil(b bool)`

 SetVoicesNil sets the value for Voices to be an explicit nil

### UnsetVoices
`func (o *LanguageOptions) UnsetVoices()`

UnsetVoices ensures that no value is present for Voices, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


