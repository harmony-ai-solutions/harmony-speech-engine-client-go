# VoiceConversionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | the name of the model | [optional] [default to ""]
**SourceAudio** | Pointer to **string** | Binary audio data to be processed, encoded in base64 | [optional] 
**TargetAudio** | Pointer to **NullableString** |  | [optional] 
**TargetEmbedding** | Pointer to **NullableString** |  | [optional] 
**GenerationOptions** | Pointer to [**NullableGenerationOptions**](GenerationOptions.md) |  | [optional] 
**OutputOptions** | Pointer to [**NullableAudioOutputOptions**](AudioOutputOptions.md) |  | [optional] 

## Methods

### NewVoiceConversionRequest

`func NewVoiceConversionRequest() *VoiceConversionRequest`

NewVoiceConversionRequest instantiates a new VoiceConversionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceConversionRequestWithDefaults

`func NewVoiceConversionRequestWithDefaults() *VoiceConversionRequest`

NewVoiceConversionRequestWithDefaults instantiates a new VoiceConversionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *VoiceConversionRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *VoiceConversionRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *VoiceConversionRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *VoiceConversionRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSourceAudio

`func (o *VoiceConversionRequest) GetSourceAudio() string`

GetSourceAudio returns the SourceAudio field if non-nil, zero value otherwise.

### GetSourceAudioOk

`func (o *VoiceConversionRequest) GetSourceAudioOk() (*string, bool)`

GetSourceAudioOk returns a tuple with the SourceAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAudio

`func (o *VoiceConversionRequest) SetSourceAudio(v string)`

SetSourceAudio sets SourceAudio field to given value.

### HasSourceAudio

`func (o *VoiceConversionRequest) HasSourceAudio() bool`

HasSourceAudio returns a boolean if a field has been set.

### GetTargetAudio

`func (o *VoiceConversionRequest) GetTargetAudio() string`

GetTargetAudio returns the TargetAudio field if non-nil, zero value otherwise.

### GetTargetAudioOk

`func (o *VoiceConversionRequest) GetTargetAudioOk() (*string, bool)`

GetTargetAudioOk returns a tuple with the TargetAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAudio

`func (o *VoiceConversionRequest) SetTargetAudio(v string)`

SetTargetAudio sets TargetAudio field to given value.

### HasTargetAudio

`func (o *VoiceConversionRequest) HasTargetAudio() bool`

HasTargetAudio returns a boolean if a field has been set.

### SetTargetAudioNil

`func (o *VoiceConversionRequest) SetTargetAudioNil(b bool)`

 SetTargetAudioNil sets the value for TargetAudio to be an explicit nil

### UnsetTargetAudio
`func (o *VoiceConversionRequest) UnsetTargetAudio()`

UnsetTargetAudio ensures that no value is present for TargetAudio, not even an explicit nil
### GetTargetEmbedding

`func (o *VoiceConversionRequest) GetTargetEmbedding() string`

GetTargetEmbedding returns the TargetEmbedding field if non-nil, zero value otherwise.

### GetTargetEmbeddingOk

`func (o *VoiceConversionRequest) GetTargetEmbeddingOk() (*string, bool)`

GetTargetEmbeddingOk returns a tuple with the TargetEmbedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEmbedding

`func (o *VoiceConversionRequest) SetTargetEmbedding(v string)`

SetTargetEmbedding sets TargetEmbedding field to given value.

### HasTargetEmbedding

`func (o *VoiceConversionRequest) HasTargetEmbedding() bool`

HasTargetEmbedding returns a boolean if a field has been set.

### SetTargetEmbeddingNil

`func (o *VoiceConversionRequest) SetTargetEmbeddingNil(b bool)`

 SetTargetEmbeddingNil sets the value for TargetEmbedding to be an explicit nil

### UnsetTargetEmbedding
`func (o *VoiceConversionRequest) UnsetTargetEmbedding()`

UnsetTargetEmbedding ensures that no value is present for TargetEmbedding, not even an explicit nil
### GetGenerationOptions

`func (o *VoiceConversionRequest) GetGenerationOptions() GenerationOptions`

GetGenerationOptions returns the GenerationOptions field if non-nil, zero value otherwise.

### GetGenerationOptionsOk

`func (o *VoiceConversionRequest) GetGenerationOptionsOk() (*GenerationOptions, bool)`

GetGenerationOptionsOk returns a tuple with the GenerationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationOptions

`func (o *VoiceConversionRequest) SetGenerationOptions(v GenerationOptions)`

SetGenerationOptions sets GenerationOptions field to given value.

### HasGenerationOptions

`func (o *VoiceConversionRequest) HasGenerationOptions() bool`

HasGenerationOptions returns a boolean if a field has been set.

### SetGenerationOptionsNil

`func (o *VoiceConversionRequest) SetGenerationOptionsNil(b bool)`

 SetGenerationOptionsNil sets the value for GenerationOptions to be an explicit nil

### UnsetGenerationOptions
`func (o *VoiceConversionRequest) UnsetGenerationOptions()`

UnsetGenerationOptions ensures that no value is present for GenerationOptions, not even an explicit nil
### GetOutputOptions

`func (o *VoiceConversionRequest) GetOutputOptions() AudioOutputOptions`

GetOutputOptions returns the OutputOptions field if non-nil, zero value otherwise.

### GetOutputOptionsOk

`func (o *VoiceConversionRequest) GetOutputOptionsOk() (*AudioOutputOptions, bool)`

GetOutputOptionsOk returns a tuple with the OutputOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputOptions

`func (o *VoiceConversionRequest) SetOutputOptions(v AudioOutputOptions)`

SetOutputOptions sets OutputOptions field to given value.

### HasOutputOptions

`func (o *VoiceConversionRequest) HasOutputOptions() bool`

HasOutputOptions returns a boolean if a field has been set.

### SetOutputOptionsNil

`func (o *VoiceConversionRequest) SetOutputOptionsNil(b bool)`

 SetOutputOptionsNil sets the value for OutputOptions to be an explicit nil

### UnsetOutputOptions
`func (o *VoiceConversionRequest) UnsetOutputOptions()`

UnsetOutputOptions ensures that no value is present for OutputOptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


