# EmbedSpeakerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** | the name of the model | [optional] [default to ""]
**Created** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to **string** | Speaker embedding data for the audio provided in the initial request, encoded in base64 | [optional] 

## Methods

### NewEmbedSpeakerResponse

`func NewEmbedSpeakerResponse() *EmbedSpeakerResponse`

NewEmbedSpeakerResponse instantiates a new EmbedSpeakerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbedSpeakerResponseWithDefaults

`func NewEmbedSpeakerResponseWithDefaults() *EmbedSpeakerResponse`

NewEmbedSpeakerResponseWithDefaults instantiates a new EmbedSpeakerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmbedSpeakerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmbedSpeakerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmbedSpeakerResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmbedSpeakerResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModel

`func (o *EmbedSpeakerResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EmbedSpeakerResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EmbedSpeakerResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EmbedSpeakerResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetCreated

`func (o *EmbedSpeakerResponse) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EmbedSpeakerResponse) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EmbedSpeakerResponse) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EmbedSpeakerResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetData

`func (o *EmbedSpeakerResponse) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EmbedSpeakerResponse) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EmbedSpeakerResponse) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *EmbedSpeakerResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


