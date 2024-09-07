# TextToSpeechResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** | the name of the model | [optional] [default to ""]
**Created** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to **string** | Audio Bytes in requested format of the initial request | [optional] 

## Methods

### NewTextToSpeechResponse

`func NewTextToSpeechResponse() *TextToSpeechResponse`

NewTextToSpeechResponse instantiates a new TextToSpeechResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextToSpeechResponseWithDefaults

`func NewTextToSpeechResponseWithDefaults() *TextToSpeechResponse`

NewTextToSpeechResponseWithDefaults instantiates a new TextToSpeechResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TextToSpeechResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TextToSpeechResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TextToSpeechResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TextToSpeechResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModel

`func (o *TextToSpeechResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TextToSpeechResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TextToSpeechResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *TextToSpeechResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetCreated

`func (o *TextToSpeechResponse) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TextToSpeechResponse) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TextToSpeechResponse) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *TextToSpeechResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetData

`func (o *TextToSpeechResponse) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TextToSpeechResponse) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TextToSpeechResponse) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *TextToSpeechResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


