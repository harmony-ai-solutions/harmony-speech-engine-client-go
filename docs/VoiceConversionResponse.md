# VoiceConversionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** | the name of the model | [optional] [default to ""]
**Created** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to **string** | Audio Bytes in requested format of the initial request | [optional] 

## Methods

### NewVoiceConversionResponse

`func NewVoiceConversionResponse() *VoiceConversionResponse`

NewVoiceConversionResponse instantiates a new VoiceConversionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceConversionResponseWithDefaults

`func NewVoiceConversionResponseWithDefaults() *VoiceConversionResponse`

NewVoiceConversionResponseWithDefaults instantiates a new VoiceConversionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VoiceConversionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VoiceConversionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VoiceConversionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VoiceConversionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModel

`func (o *VoiceConversionResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *VoiceConversionResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *VoiceConversionResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *VoiceConversionResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetCreated

`func (o *VoiceConversionResponse) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VoiceConversionResponse) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VoiceConversionResponse) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *VoiceConversionResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetData

`func (o *VoiceConversionResponse) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VoiceConversionResponse) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VoiceConversionResponse) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *VoiceConversionResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


