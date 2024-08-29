# SpeechTranscribeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | the name of the model | [optional] [default to ""]
**InputAudio** | Pointer to **NullableString** |  | [optional] 
**GetLanguage** | Pointer to **NullableBool** |  | [optional] 
**GetTimestamps** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewSpeechTranscribeRequest

`func NewSpeechTranscribeRequest() *SpeechTranscribeRequest`

NewSpeechTranscribeRequest instantiates a new SpeechTranscribeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpeechTranscribeRequestWithDefaults

`func NewSpeechTranscribeRequestWithDefaults() *SpeechTranscribeRequest`

NewSpeechTranscribeRequestWithDefaults instantiates a new SpeechTranscribeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *SpeechTranscribeRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SpeechTranscribeRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SpeechTranscribeRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SpeechTranscribeRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetInputAudio

`func (o *SpeechTranscribeRequest) GetInputAudio() string`

GetInputAudio returns the InputAudio field if non-nil, zero value otherwise.

### GetInputAudioOk

`func (o *SpeechTranscribeRequest) GetInputAudioOk() (*string, bool)`

GetInputAudioOk returns a tuple with the InputAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputAudio

`func (o *SpeechTranscribeRequest) SetInputAudio(v string)`

SetInputAudio sets InputAudio field to given value.

### HasInputAudio

`func (o *SpeechTranscribeRequest) HasInputAudio() bool`

HasInputAudio returns a boolean if a field has been set.

### SetInputAudioNil

`func (o *SpeechTranscribeRequest) SetInputAudioNil(b bool)`

 SetInputAudioNil sets the value for InputAudio to be an explicit nil

### UnsetInputAudio
`func (o *SpeechTranscribeRequest) UnsetInputAudio()`

UnsetInputAudio ensures that no value is present for InputAudio, not even an explicit nil
### GetGetLanguage

`func (o *SpeechTranscribeRequest) GetGetLanguage() bool`

GetGetLanguage returns the GetLanguage field if non-nil, zero value otherwise.

### GetGetLanguageOk

`func (o *SpeechTranscribeRequest) GetGetLanguageOk() (*bool, bool)`

GetGetLanguageOk returns a tuple with the GetLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetLanguage

`func (o *SpeechTranscribeRequest) SetGetLanguage(v bool)`

SetGetLanguage sets GetLanguage field to given value.

### HasGetLanguage

`func (o *SpeechTranscribeRequest) HasGetLanguage() bool`

HasGetLanguage returns a boolean if a field has been set.

### SetGetLanguageNil

`func (o *SpeechTranscribeRequest) SetGetLanguageNil(b bool)`

 SetGetLanguageNil sets the value for GetLanguage to be an explicit nil

### UnsetGetLanguage
`func (o *SpeechTranscribeRequest) UnsetGetLanguage()`

UnsetGetLanguage ensures that no value is present for GetLanguage, not even an explicit nil
### GetGetTimestamps

`func (o *SpeechTranscribeRequest) GetGetTimestamps() bool`

GetGetTimestamps returns the GetTimestamps field if non-nil, zero value otherwise.

### GetGetTimestampsOk

`func (o *SpeechTranscribeRequest) GetGetTimestampsOk() (*bool, bool)`

GetGetTimestampsOk returns a tuple with the GetTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetTimestamps

`func (o *SpeechTranscribeRequest) SetGetTimestamps(v bool)`

SetGetTimestamps sets GetTimestamps field to given value.

### HasGetTimestamps

`func (o *SpeechTranscribeRequest) HasGetTimestamps() bool`

HasGetTimestamps returns a boolean if a field has been set.

### SetGetTimestampsNil

`func (o *SpeechTranscribeRequest) SetGetTimestampsNil(b bool)`

 SetGetTimestampsNil sets the value for GetTimestamps to be an explicit nil

### UnsetGetTimestamps
`func (o *SpeechTranscribeRequest) UnsetGetTimestamps()`

UnsetGetTimestamps ensures that no value is present for GetTimestamps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


