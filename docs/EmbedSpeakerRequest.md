# EmbedSpeakerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | the name of the model | [optional] [default to ""]
**InputAudio** | Pointer to **string** | Binary audio data to be processed, encoded in base64 | [optional] 

## Methods

### NewEmbedSpeakerRequest

`func NewEmbedSpeakerRequest() *EmbedSpeakerRequest`

NewEmbedSpeakerRequest instantiates a new EmbedSpeakerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbedSpeakerRequestWithDefaults

`func NewEmbedSpeakerRequestWithDefaults() *EmbedSpeakerRequest`

NewEmbedSpeakerRequestWithDefaults instantiates a new EmbedSpeakerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *EmbedSpeakerRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EmbedSpeakerRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EmbedSpeakerRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EmbedSpeakerRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetInputAudio

`func (o *EmbedSpeakerRequest) GetInputAudio() string`

GetInputAudio returns the InputAudio field if non-nil, zero value otherwise.

### GetInputAudioOk

`func (o *EmbedSpeakerRequest) GetInputAudioOk() (*string, bool)`

GetInputAudioOk returns a tuple with the InputAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputAudio

`func (o *EmbedSpeakerRequest) SetInputAudio(v string)`

SetInputAudio sets InputAudio field to given value.

### HasInputAudio

`func (o *EmbedSpeakerRequest) HasInputAudio() bool`

HasInputAudio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


