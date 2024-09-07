# SpeechToTextResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** | the name of the model | [optional] [default to ""]
**Created** | Pointer to **int32** |  | [optional] 
**Text** | Pointer to **string** | text retrieved from the input audio | [optional] [default to ""]
**Language** | Pointer to **NullableString** |  | [optional] 
**Timestamps** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewSpeechToTextResponse

`func NewSpeechToTextResponse() *SpeechToTextResponse`

NewSpeechToTextResponse instantiates a new SpeechToTextResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpeechToTextResponseWithDefaults

`func NewSpeechToTextResponseWithDefaults() *SpeechToTextResponse`

NewSpeechToTextResponseWithDefaults instantiates a new SpeechToTextResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SpeechToTextResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpeechToTextResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpeechToTextResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SpeechToTextResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModel

`func (o *SpeechToTextResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SpeechToTextResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SpeechToTextResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SpeechToTextResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetCreated

`func (o *SpeechToTextResponse) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SpeechToTextResponse) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SpeechToTextResponse) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SpeechToTextResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetText

`func (o *SpeechToTextResponse) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SpeechToTextResponse) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SpeechToTextResponse) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *SpeechToTextResponse) HasText() bool`

HasText returns a boolean if a field has been set.

### GetLanguage

`func (o *SpeechToTextResponse) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *SpeechToTextResponse) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *SpeechToTextResponse) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *SpeechToTextResponse) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *SpeechToTextResponse) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *SpeechToTextResponse) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetTimestamps

`func (o *SpeechToTextResponse) GetTimestamps() []map[string]interface{}`

GetTimestamps returns the Timestamps field if non-nil, zero value otherwise.

### GetTimestampsOk

`func (o *SpeechToTextResponse) GetTimestampsOk() (*[]map[string]interface{}, bool)`

GetTimestampsOk returns a tuple with the Timestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamps

`func (o *SpeechToTextResponse) SetTimestamps(v []map[string]interface{})`

SetTimestamps sets Timestamps field to given value.

### HasTimestamps

`func (o *SpeechToTextResponse) HasTimestamps() bool`

HasTimestamps returns a boolean if a field has been set.

### SetTimestampsNil

`func (o *SpeechToTextResponse) SetTimestampsNil(b bool)`

 SetTimestampsNil sets the value for Timestamps to be an explicit nil

### UnsetTimestamps
`func (o *SpeechToTextResponse) UnsetTimestamps()`

UnsetTimestamps ensures that no value is present for Timestamps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


