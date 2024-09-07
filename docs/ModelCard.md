# ModelCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "model"]
**Created** | Pointer to **int32** |  | [optional] 
**OwnedBy** | Pointer to **string** |  | [optional] [default to "harmony-ai-solutions"]
**Root** | Pointer to **NullableString** |  | [optional] 
**Parent** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewModelCard

`func NewModelCard(id string, ) *ModelCard`

NewModelCard instantiates a new ModelCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCardWithDefaults

`func NewModelCardWithDefaults() *ModelCard`

NewModelCardWithDefaults instantiates a new ModelCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelCard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelCard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelCard) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *ModelCard) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ModelCard) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ModelCard) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ModelCard) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreated

`func (o *ModelCard) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelCard) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelCard) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ModelCard) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetOwnedBy

`func (o *ModelCard) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *ModelCard) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *ModelCard) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *ModelCard) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetRoot

`func (o *ModelCard) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *ModelCard) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *ModelCard) SetRoot(v string)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *ModelCard) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### SetRootNil

`func (o *ModelCard) SetRootNil(b bool)`

 SetRootNil sets the value for Root to be an explicit nil

### UnsetRoot
`func (o *ModelCard) UnsetRoot()`

UnsetRoot ensures that no value is present for Root, not even an explicit nil
### GetParent

`func (o *ModelCard) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ModelCard) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ModelCard) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ModelCard) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *ModelCard) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *ModelCard) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


