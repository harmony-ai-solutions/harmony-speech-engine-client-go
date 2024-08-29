# AudioOutputOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **NullableString** |  | [optional] 
**SampleRate** | Pointer to **NullableInt32** |  | [optional] 
**Stream** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewAudioOutputOptions

`func NewAudioOutputOptions() *AudioOutputOptions`

NewAudioOutputOptions instantiates a new AudioOutputOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioOutputOptionsWithDefaults

`func NewAudioOutputOptionsWithDefaults() *AudioOutputOptions`

NewAudioOutputOptionsWithDefaults instantiates a new AudioOutputOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *AudioOutputOptions) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *AudioOutputOptions) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *AudioOutputOptions) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *AudioOutputOptions) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *AudioOutputOptions) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *AudioOutputOptions) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetSampleRate

`func (o *AudioOutputOptions) GetSampleRate() int32`

GetSampleRate returns the SampleRate field if non-nil, zero value otherwise.

### GetSampleRateOk

`func (o *AudioOutputOptions) GetSampleRateOk() (*int32, bool)`

GetSampleRateOk returns a tuple with the SampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleRate

`func (o *AudioOutputOptions) SetSampleRate(v int32)`

SetSampleRate sets SampleRate field to given value.

### HasSampleRate

`func (o *AudioOutputOptions) HasSampleRate() bool`

HasSampleRate returns a boolean if a field has been set.

### SetSampleRateNil

`func (o *AudioOutputOptions) SetSampleRateNil(b bool)`

 SetSampleRateNil sets the value for SampleRate to be an explicit nil

### UnsetSampleRate
`func (o *AudioOutputOptions) UnsetSampleRate()`

UnsetSampleRate ensures that no value is present for SampleRate, not even an explicit nil
### GetStream

`func (o *AudioOutputOptions) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *AudioOutputOptions) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *AudioOutputOptions) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *AudioOutputOptions) HasStream() bool`

HasStream returns a boolean if a field has been set.

### SetStreamNil

`func (o *AudioOutputOptions) SetStreamNil(b bool)`

 SetStreamNil sets the value for Stream to be an explicit nil

### UnsetStream
`func (o *AudioOutputOptions) UnsetStream()`

UnsetStream ensures that no value is present for Stream, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


