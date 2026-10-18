# DetectVoiceActivityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | the name of the model | [optional] [default to ""]
**InputAudio** | Pointer to **string** | Binary audio data of the reference speaker for synthesizing the text, encoded in base64 | [optional] 
**GetTimestamps** | Pointer to **NullableBool** | whether to return the word timestamps. Check model description if supported. | [optional] 
**Threshold** | Pointer to **NullableFloat32** | Speech threshold. Probabilities above this value are considered as speech. | [optional] 
**MinSpeechDurationMs** | Pointer to **NullableInt32** | Final speech chunks shorter than this duration (in ms) are discarded. | [optional] 
**MinSilenceDurationMs** | Pointer to **NullableInt32** | Minimum silence duration (in ms) before separating speech chunks. | [optional] 
**SpeechPadMs** | Pointer to **NullableInt32** | Speech chunks are padded by this duration (in ms) on each side. | [optional] 
**ReturnSeconds** | Pointer to **NullableBool** | Whether to return timestamps in seconds (default: samples). | [optional] 

## Methods

### NewDetectVoiceActivityRequest

`func NewDetectVoiceActivityRequest() *DetectVoiceActivityRequest`

NewDetectVoiceActivityRequest instantiates a new DetectVoiceActivityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectVoiceActivityRequestWithDefaults

`func NewDetectVoiceActivityRequestWithDefaults() *DetectVoiceActivityRequest`

NewDetectVoiceActivityRequestWithDefaults instantiates a new DetectVoiceActivityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *DetectVoiceActivityRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DetectVoiceActivityRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DetectVoiceActivityRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DetectVoiceActivityRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetInputAudio

`func (o *DetectVoiceActivityRequest) GetInputAudio() string`

GetInputAudio returns the InputAudio field if non-nil, zero value otherwise.

### GetInputAudioOk

`func (o *DetectVoiceActivityRequest) GetInputAudioOk() (*string, bool)`

GetInputAudioOk returns a tuple with the InputAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputAudio

`func (o *DetectVoiceActivityRequest) SetInputAudio(v string)`

SetInputAudio sets InputAudio field to given value.

### HasInputAudio

`func (o *DetectVoiceActivityRequest) HasInputAudio() bool`

HasInputAudio returns a boolean if a field has been set.

### GetGetTimestamps

`func (o *DetectVoiceActivityRequest) GetGetTimestamps() bool`

GetGetTimestamps returns the GetTimestamps field if non-nil, zero value otherwise.

### GetGetTimestampsOk

`func (o *DetectVoiceActivityRequest) GetGetTimestampsOk() (*bool, bool)`

GetGetTimestampsOk returns a tuple with the GetTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetTimestamps

`func (o *DetectVoiceActivityRequest) SetGetTimestamps(v bool)`

SetGetTimestamps sets GetTimestamps field to given value.

### HasGetTimestamps

`func (o *DetectVoiceActivityRequest) HasGetTimestamps() bool`

HasGetTimestamps returns a boolean if a field has been set.

### SetGetTimestampsNil

`func (o *DetectVoiceActivityRequest) SetGetTimestampsNil(b bool)`

 SetGetTimestampsNil sets the value for GetTimestamps to be an explicit nil

### UnsetGetTimestamps
`func (o *DetectVoiceActivityRequest) UnsetGetTimestamps()`

UnsetGetTimestamps ensures that no value is present for GetTimestamps, not even an explicit nil
### GetThreshold

`func (o *DetectVoiceActivityRequest) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *DetectVoiceActivityRequest) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *DetectVoiceActivityRequest) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *DetectVoiceActivityRequest) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### SetThresholdNil

`func (o *DetectVoiceActivityRequest) SetThresholdNil(b bool)`

 SetThresholdNil sets the value for Threshold to be an explicit nil

### UnsetThreshold
`func (o *DetectVoiceActivityRequest) UnsetThreshold()`

UnsetThreshold ensures that no value is present for Threshold, not even an explicit nil
### GetMinSpeechDurationMs

`func (o *DetectVoiceActivityRequest) GetMinSpeechDurationMs() int32`

GetMinSpeechDurationMs returns the MinSpeechDurationMs field if non-nil, zero value otherwise.

### GetMinSpeechDurationMsOk

`func (o *DetectVoiceActivityRequest) GetMinSpeechDurationMsOk() (*int32, bool)`

GetMinSpeechDurationMsOk returns a tuple with the MinSpeechDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSpeechDurationMs

`func (o *DetectVoiceActivityRequest) SetMinSpeechDurationMs(v int32)`

SetMinSpeechDurationMs sets MinSpeechDurationMs field to given value.

### HasMinSpeechDurationMs

`func (o *DetectVoiceActivityRequest) HasMinSpeechDurationMs() bool`

HasMinSpeechDurationMs returns a boolean if a field has been set.

### SetMinSpeechDurationMsNil

`func (o *DetectVoiceActivityRequest) SetMinSpeechDurationMsNil(b bool)`

 SetMinSpeechDurationMsNil sets the value for MinSpeechDurationMs to be an explicit nil

### UnsetMinSpeechDurationMs
`func (o *DetectVoiceActivityRequest) UnsetMinSpeechDurationMs()`

UnsetMinSpeechDurationMs ensures that no value is present for MinSpeechDurationMs, not even an explicit nil
### GetMinSilenceDurationMs

`func (o *DetectVoiceActivityRequest) GetMinSilenceDurationMs() int32`

GetMinSilenceDurationMs returns the MinSilenceDurationMs field if non-nil, zero value otherwise.

### GetMinSilenceDurationMsOk

`func (o *DetectVoiceActivityRequest) GetMinSilenceDurationMsOk() (*int32, bool)`

GetMinSilenceDurationMsOk returns a tuple with the MinSilenceDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSilenceDurationMs

`func (o *DetectVoiceActivityRequest) SetMinSilenceDurationMs(v int32)`

SetMinSilenceDurationMs sets MinSilenceDurationMs field to given value.

### HasMinSilenceDurationMs

`func (o *DetectVoiceActivityRequest) HasMinSilenceDurationMs() bool`

HasMinSilenceDurationMs returns a boolean if a field has been set.

### SetMinSilenceDurationMsNil

`func (o *DetectVoiceActivityRequest) SetMinSilenceDurationMsNil(b bool)`

 SetMinSilenceDurationMsNil sets the value for MinSilenceDurationMs to be an explicit nil

### UnsetMinSilenceDurationMs
`func (o *DetectVoiceActivityRequest) UnsetMinSilenceDurationMs()`

UnsetMinSilenceDurationMs ensures that no value is present for MinSilenceDurationMs, not even an explicit nil
### GetSpeechPadMs

`func (o *DetectVoiceActivityRequest) GetSpeechPadMs() int32`

GetSpeechPadMs returns the SpeechPadMs field if non-nil, zero value otherwise.

### GetSpeechPadMsOk

`func (o *DetectVoiceActivityRequest) GetSpeechPadMsOk() (*int32, bool)`

GetSpeechPadMsOk returns a tuple with the SpeechPadMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeechPadMs

`func (o *DetectVoiceActivityRequest) SetSpeechPadMs(v int32)`

SetSpeechPadMs sets SpeechPadMs field to given value.

### HasSpeechPadMs

`func (o *DetectVoiceActivityRequest) HasSpeechPadMs() bool`

HasSpeechPadMs returns a boolean if a field has been set.

### SetSpeechPadMsNil

`func (o *DetectVoiceActivityRequest) SetSpeechPadMsNil(b bool)`

 SetSpeechPadMsNil sets the value for SpeechPadMs to be an explicit nil

### UnsetSpeechPadMs
`func (o *DetectVoiceActivityRequest) UnsetSpeechPadMs()`

UnsetSpeechPadMs ensures that no value is present for SpeechPadMs, not even an explicit nil
### GetReturnSeconds

`func (o *DetectVoiceActivityRequest) GetReturnSeconds() bool`

GetReturnSeconds returns the ReturnSeconds field if non-nil, zero value otherwise.

### GetReturnSecondsOk

`func (o *DetectVoiceActivityRequest) GetReturnSecondsOk() (*bool, bool)`

GetReturnSecondsOk returns a tuple with the ReturnSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnSeconds

`func (o *DetectVoiceActivityRequest) SetReturnSeconds(v bool)`

SetReturnSeconds sets ReturnSeconds field to given value.

### HasReturnSeconds

`func (o *DetectVoiceActivityRequest) HasReturnSeconds() bool`

HasReturnSeconds returns a boolean if a field has been set.

### SetReturnSecondsNil

`func (o *DetectVoiceActivityRequest) SetReturnSecondsNil(b bool)`

 SetReturnSecondsNil sets the value for ReturnSeconds to be an explicit nil

### UnsetReturnSeconds
`func (o *DetectVoiceActivityRequest) UnsetReturnSeconds()`

UnsetReturnSeconds ensures that no value is present for ReturnSeconds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


