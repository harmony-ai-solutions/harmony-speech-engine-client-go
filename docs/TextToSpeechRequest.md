# TextToSpeechRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | the name of the model | [optional] [default to ""]
**Input** | Pointer to **string** | the text to synthesize | [optional] [default to ""]
**Language** | Pointer to **string** | language to synthesize. Only if the model supports language IDs. Please refer to the model&#39;s documentation for which values are availiable. | [optional] 
**Voice** | Pointer to **NullableString** |  | [optional] 
**InputAudio** | Pointer to **NullableString** |  | [optional] 
**InputVadMode** | Pointer to **NullableString** |  | [optional] 
**InputVadData** | Pointer to **NullableString** |  | [optional] 
**InputEmbedding** | Pointer to **NullableString** |  | [optional] 
**GenerationOptions** | Pointer to [**NullableGenerationOptions**](GenerationOptions.md) |  | [optional] 
**OutputOptions** | Pointer to [**NullableAudioOutputOptions**](AudioOutputOptions.md) |  | [optional] 
**PostGenerationFilters** | Pointer to [**[]VoiceConversionRequest**](VoiceConversionRequest.md) |  | [optional] 

## Methods

### NewTextToSpeechRequest

`func NewTextToSpeechRequest() *TextToSpeechRequest`

NewTextToSpeechRequest instantiates a new TextToSpeechRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextToSpeechRequestWithDefaults

`func NewTextToSpeechRequestWithDefaults() *TextToSpeechRequest`

NewTextToSpeechRequestWithDefaults instantiates a new TextToSpeechRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *TextToSpeechRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TextToSpeechRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TextToSpeechRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *TextToSpeechRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetInput

`func (o *TextToSpeechRequest) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *TextToSpeechRequest) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *TextToSpeechRequest) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *TextToSpeechRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetLanguage

`func (o *TextToSpeechRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *TextToSpeechRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *TextToSpeechRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *TextToSpeechRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetVoice

`func (o *TextToSpeechRequest) GetVoice() string`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *TextToSpeechRequest) GetVoiceOk() (*string, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *TextToSpeechRequest) SetVoice(v string)`

SetVoice sets Voice field to given value.

### HasVoice

`func (o *TextToSpeechRequest) HasVoice() bool`

HasVoice returns a boolean if a field has been set.

### SetVoiceNil

`func (o *TextToSpeechRequest) SetVoiceNil(b bool)`

 SetVoiceNil sets the value for Voice to be an explicit nil

### UnsetVoice
`func (o *TextToSpeechRequest) UnsetVoice()`

UnsetVoice ensures that no value is present for Voice, not even an explicit nil
### GetInputAudio

`func (o *TextToSpeechRequest) GetInputAudio() string`

GetInputAudio returns the InputAudio field if non-nil, zero value otherwise.

### GetInputAudioOk

`func (o *TextToSpeechRequest) GetInputAudioOk() (*string, bool)`

GetInputAudioOk returns a tuple with the InputAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputAudio

`func (o *TextToSpeechRequest) SetInputAudio(v string)`

SetInputAudio sets InputAudio field to given value.

### HasInputAudio

`func (o *TextToSpeechRequest) HasInputAudio() bool`

HasInputAudio returns a boolean if a field has been set.

### SetInputAudioNil

`func (o *TextToSpeechRequest) SetInputAudioNil(b bool)`

 SetInputAudioNil sets the value for InputAudio to be an explicit nil

### UnsetInputAudio
`func (o *TextToSpeechRequest) UnsetInputAudio()`

UnsetInputAudio ensures that no value is present for InputAudio, not even an explicit nil
### GetInputVadMode

`func (o *TextToSpeechRequest) GetInputVadMode() string`

GetInputVadMode returns the InputVadMode field if non-nil, zero value otherwise.

### GetInputVadModeOk

`func (o *TextToSpeechRequest) GetInputVadModeOk() (*string, bool)`

GetInputVadModeOk returns a tuple with the InputVadMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputVadMode

`func (o *TextToSpeechRequest) SetInputVadMode(v string)`

SetInputVadMode sets InputVadMode field to given value.

### HasInputVadMode

`func (o *TextToSpeechRequest) HasInputVadMode() bool`

HasInputVadMode returns a boolean if a field has been set.

### SetInputVadModeNil

`func (o *TextToSpeechRequest) SetInputVadModeNil(b bool)`

 SetInputVadModeNil sets the value for InputVadMode to be an explicit nil

### UnsetInputVadMode
`func (o *TextToSpeechRequest) UnsetInputVadMode()`

UnsetInputVadMode ensures that no value is present for InputVadMode, not even an explicit nil
### GetInputVadData

`func (o *TextToSpeechRequest) GetInputVadData() string`

GetInputVadData returns the InputVadData field if non-nil, zero value otherwise.

### GetInputVadDataOk

`func (o *TextToSpeechRequest) GetInputVadDataOk() (*string, bool)`

GetInputVadDataOk returns a tuple with the InputVadData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputVadData

`func (o *TextToSpeechRequest) SetInputVadData(v string)`

SetInputVadData sets InputVadData field to given value.

### HasInputVadData

`func (o *TextToSpeechRequest) HasInputVadData() bool`

HasInputVadData returns a boolean if a field has been set.

### SetInputVadDataNil

`func (o *TextToSpeechRequest) SetInputVadDataNil(b bool)`

 SetInputVadDataNil sets the value for InputVadData to be an explicit nil

### UnsetInputVadData
`func (o *TextToSpeechRequest) UnsetInputVadData()`

UnsetInputVadData ensures that no value is present for InputVadData, not even an explicit nil
### GetInputEmbedding

`func (o *TextToSpeechRequest) GetInputEmbedding() string`

GetInputEmbedding returns the InputEmbedding field if non-nil, zero value otherwise.

### GetInputEmbeddingOk

`func (o *TextToSpeechRequest) GetInputEmbeddingOk() (*string, bool)`

GetInputEmbeddingOk returns a tuple with the InputEmbedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputEmbedding

`func (o *TextToSpeechRequest) SetInputEmbedding(v string)`

SetInputEmbedding sets InputEmbedding field to given value.

### HasInputEmbedding

`func (o *TextToSpeechRequest) HasInputEmbedding() bool`

HasInputEmbedding returns a boolean if a field has been set.

### SetInputEmbeddingNil

`func (o *TextToSpeechRequest) SetInputEmbeddingNil(b bool)`

 SetInputEmbeddingNil sets the value for InputEmbedding to be an explicit nil

### UnsetInputEmbedding
`func (o *TextToSpeechRequest) UnsetInputEmbedding()`

UnsetInputEmbedding ensures that no value is present for InputEmbedding, not even an explicit nil
### GetGenerationOptions

`func (o *TextToSpeechRequest) GetGenerationOptions() GenerationOptions`

GetGenerationOptions returns the GenerationOptions field if non-nil, zero value otherwise.

### GetGenerationOptionsOk

`func (o *TextToSpeechRequest) GetGenerationOptionsOk() (*GenerationOptions, bool)`

GetGenerationOptionsOk returns a tuple with the GenerationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationOptions

`func (o *TextToSpeechRequest) SetGenerationOptions(v GenerationOptions)`

SetGenerationOptions sets GenerationOptions field to given value.

### HasGenerationOptions

`func (o *TextToSpeechRequest) HasGenerationOptions() bool`

HasGenerationOptions returns a boolean if a field has been set.

### SetGenerationOptionsNil

`func (o *TextToSpeechRequest) SetGenerationOptionsNil(b bool)`

 SetGenerationOptionsNil sets the value for GenerationOptions to be an explicit nil

### UnsetGenerationOptions
`func (o *TextToSpeechRequest) UnsetGenerationOptions()`

UnsetGenerationOptions ensures that no value is present for GenerationOptions, not even an explicit nil
### GetOutputOptions

`func (o *TextToSpeechRequest) GetOutputOptions() AudioOutputOptions`

GetOutputOptions returns the OutputOptions field if non-nil, zero value otherwise.

### GetOutputOptionsOk

`func (o *TextToSpeechRequest) GetOutputOptionsOk() (*AudioOutputOptions, bool)`

GetOutputOptionsOk returns a tuple with the OutputOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputOptions

`func (o *TextToSpeechRequest) SetOutputOptions(v AudioOutputOptions)`

SetOutputOptions sets OutputOptions field to given value.

### HasOutputOptions

`func (o *TextToSpeechRequest) HasOutputOptions() bool`

HasOutputOptions returns a boolean if a field has been set.

### SetOutputOptionsNil

`func (o *TextToSpeechRequest) SetOutputOptionsNil(b bool)`

 SetOutputOptionsNil sets the value for OutputOptions to be an explicit nil

### UnsetOutputOptions
`func (o *TextToSpeechRequest) UnsetOutputOptions()`

UnsetOutputOptions ensures that no value is present for OutputOptions, not even an explicit nil
### GetPostGenerationFilters

`func (o *TextToSpeechRequest) GetPostGenerationFilters() []VoiceConversionRequest`

GetPostGenerationFilters returns the PostGenerationFilters field if non-nil, zero value otherwise.

### GetPostGenerationFiltersOk

`func (o *TextToSpeechRequest) GetPostGenerationFiltersOk() (*[]VoiceConversionRequest, bool)`

GetPostGenerationFiltersOk returns a tuple with the PostGenerationFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostGenerationFilters

`func (o *TextToSpeechRequest) SetPostGenerationFilters(v []VoiceConversionRequest)`

SetPostGenerationFilters sets PostGenerationFilters field to given value.

### HasPostGenerationFilters

`func (o *TextToSpeechRequest) HasPostGenerationFilters() bool`

HasPostGenerationFilters returns a boolean if a field has been set.

### SetPostGenerationFiltersNil

`func (o *TextToSpeechRequest) SetPostGenerationFiltersNil(b bool)`

 SetPostGenerationFiltersNil sets the value for PostGenerationFilters to be an explicit nil

### UnsetPostGenerationFilters
`func (o *TextToSpeechRequest) UnsetPostGenerationFilters()`

UnsetPostGenerationFilters ensures that no value is present for PostGenerationFilters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


