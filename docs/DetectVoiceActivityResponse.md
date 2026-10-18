# DetectVoiceActivityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** | the name of the model | [optional] [default to ""]
**Created** | Pointer to **int32** |  | [optional] 
**SpeechActivity** | Pointer to **bool** | whether speech activity was detected | [optional] [default to false]
**Timestamps** | Pointer to **[]map[string]interface{}** | list of the detected timestamps | [optional] 

## Methods

### NewDetectVoiceActivityResponse

`func NewDetectVoiceActivityResponse() *DetectVoiceActivityResponse`

NewDetectVoiceActivityResponse instantiates a new DetectVoiceActivityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectVoiceActivityResponseWithDefaults

`func NewDetectVoiceActivityResponseWithDefaults() *DetectVoiceActivityResponse`

NewDetectVoiceActivityResponseWithDefaults instantiates a new DetectVoiceActivityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DetectVoiceActivityResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DetectVoiceActivityResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DetectVoiceActivityResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DetectVoiceActivityResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModel

`func (o *DetectVoiceActivityResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DetectVoiceActivityResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DetectVoiceActivityResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DetectVoiceActivityResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetCreated

`func (o *DetectVoiceActivityResponse) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DetectVoiceActivityResponse) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DetectVoiceActivityResponse) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DetectVoiceActivityResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetSpeechActivity

`func (o *DetectVoiceActivityResponse) GetSpeechActivity() bool`

GetSpeechActivity returns the SpeechActivity field if non-nil, zero value otherwise.

### GetSpeechActivityOk

`func (o *DetectVoiceActivityResponse) GetSpeechActivityOk() (*bool, bool)`

GetSpeechActivityOk returns a tuple with the SpeechActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeechActivity

`func (o *DetectVoiceActivityResponse) SetSpeechActivity(v bool)`

SetSpeechActivity sets SpeechActivity field to given value.

### HasSpeechActivity

`func (o *DetectVoiceActivityResponse) HasSpeechActivity() bool`

HasSpeechActivity returns a boolean if a field has been set.

### GetTimestamps

`func (o *DetectVoiceActivityResponse) GetTimestamps() []map[string]interface{}`

GetTimestamps returns the Timestamps field if non-nil, zero value otherwise.

### GetTimestampsOk

`func (o *DetectVoiceActivityResponse) GetTimestampsOk() (*[]map[string]interface{}, bool)`

GetTimestampsOk returns a tuple with the Timestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamps

`func (o *DetectVoiceActivityResponse) SetTimestamps(v []map[string]interface{})`

SetTimestamps sets Timestamps field to given value.

### HasTimestamps

`func (o *DetectVoiceActivityResponse) HasTimestamps() bool`

HasTimestamps returns a boolean if a field has been set.

### SetTimestampsNil

`func (o *DetectVoiceActivityResponse) SetTimestampsNil(b bool)`

 SetTimestampsNil sets the value for Timestamps to be an explicit nil

### UnsetTimestamps
`func (o *DetectVoiceActivityResponse) UnsetTimestamps()`

UnsetTimestamps ensures that no value is present for Timestamps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


