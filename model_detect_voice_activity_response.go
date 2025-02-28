/*
Harmony Speech Engine API

API for the Harmony Speech Engine.

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package harmonyspeech

import (
	"encoding/json"
)

// checks if the DetectVoiceActivityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetectVoiceActivityResponse{}

// DetectVoiceActivityResponse DetectVoiceActivityResponse Result text determined from audio and optional language tag
type DetectVoiceActivityResponse struct {
	Id *string `json:"id,omitempty"`
	// the name of the model
	Model   *string `json:"model,omitempty"`
	Created *int32  `json:"created,omitempty"`
	// whether speech activity was detected
	SpeechActivity       *bool                    `json:"speech_activity,omitempty"`
	Timestamps           []map[string]interface{} `json:"timestamps,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DetectVoiceActivityResponse DetectVoiceActivityResponse

// NewDetectVoiceActivityResponse instantiates a new DetectVoiceActivityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetectVoiceActivityResponse() *DetectVoiceActivityResponse {
	this := DetectVoiceActivityResponse{}
	var model string = ""
	this.Model = &model
	var speechActivity bool = false
	this.SpeechActivity = &speechActivity
	return &this
}

// NewDetectVoiceActivityResponseWithDefaults instantiates a new DetectVoiceActivityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetectVoiceActivityResponseWithDefaults() *DetectVoiceActivityResponse {
	this := DetectVoiceActivityResponse{}
	var model string = ""
	this.Model = &model
	var speechActivity bool = false
	this.SpeechActivity = &speechActivity
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DetectVoiceActivityResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectVoiceActivityResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DetectVoiceActivityResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DetectVoiceActivityResponse) SetId(v string) {
	o.Id = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *DetectVoiceActivityResponse) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectVoiceActivityResponse) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *DetectVoiceActivityResponse) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *DetectVoiceActivityResponse) SetModel(v string) {
	o.Model = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *DetectVoiceActivityResponse) GetCreated() int32 {
	if o == nil || IsNil(o.Created) {
		var ret int32
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectVoiceActivityResponse) GetCreatedOk() (*int32, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *DetectVoiceActivityResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given int32 and assigns it to the Created field.
func (o *DetectVoiceActivityResponse) SetCreated(v int32) {
	o.Created = &v
}

// GetSpeechActivity returns the SpeechActivity field value if set, zero value otherwise.
func (o *DetectVoiceActivityResponse) GetSpeechActivity() bool {
	if o == nil || IsNil(o.SpeechActivity) {
		var ret bool
		return ret
	}
	return *o.SpeechActivity
}

// GetSpeechActivityOk returns a tuple with the SpeechActivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectVoiceActivityResponse) GetSpeechActivityOk() (*bool, bool) {
	if o == nil || IsNil(o.SpeechActivity) {
		return nil, false
	}
	return o.SpeechActivity, true
}

// HasSpeechActivity returns a boolean if a field has been set.
func (o *DetectVoiceActivityResponse) HasSpeechActivity() bool {
	if o != nil && !IsNil(o.SpeechActivity) {
		return true
	}

	return false
}

// SetSpeechActivity gets a reference to the given bool and assigns it to the SpeechActivity field.
func (o *DetectVoiceActivityResponse) SetSpeechActivity(v bool) {
	o.SpeechActivity = &v
}

// GetTimestamps returns the Timestamps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DetectVoiceActivityResponse) GetTimestamps() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Timestamps
}

// GetTimestampsOk returns a tuple with the Timestamps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DetectVoiceActivityResponse) GetTimestampsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Timestamps) {
		return nil, false
	}
	return o.Timestamps, true
}

// HasTimestamps returns a boolean if a field has been set.
func (o *DetectVoiceActivityResponse) HasTimestamps() bool {
	if o != nil && !IsNil(o.Timestamps) {
		return true
	}

	return false
}

// SetTimestamps gets a reference to the given []map[string]interface{} and assigns it to the Timestamps field.
func (o *DetectVoiceActivityResponse) SetTimestamps(v []map[string]interface{}) {
	o.Timestamps = v
}

func (o DetectVoiceActivityResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetectVoiceActivityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.SpeechActivity) {
		toSerialize["speech_activity"] = o.SpeechActivity
	}
	if o.Timestamps != nil {
		toSerialize["timestamps"] = o.Timestamps
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DetectVoiceActivityResponse) UnmarshalJSON(data []byte) (err error) {
	varDetectVoiceActivityResponse := _DetectVoiceActivityResponse{}

	err = json.Unmarshal(data, &varDetectVoiceActivityResponse)

	if err != nil {
		return err
	}

	*o = DetectVoiceActivityResponse(varDetectVoiceActivityResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "model")
		delete(additionalProperties, "created")
		delete(additionalProperties, "speech_activity")
		delete(additionalProperties, "timestamps")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDetectVoiceActivityResponse struct {
	value *DetectVoiceActivityResponse
	isSet bool
}

func (v NullableDetectVoiceActivityResponse) Get() *DetectVoiceActivityResponse {
	return v.value
}

func (v *NullableDetectVoiceActivityResponse) Set(val *DetectVoiceActivityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDetectVoiceActivityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDetectVoiceActivityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetectVoiceActivityResponse(val *DetectVoiceActivityResponse) *NullableDetectVoiceActivityResponse {
	return &NullableDetectVoiceActivityResponse{value: val, isSet: true}
}

func (v NullableDetectVoiceActivityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetectVoiceActivityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
