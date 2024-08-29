/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GenerationOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerationOptions{}

// GenerationOptions struct for GenerationOptions
type GenerationOptions struct {
	Seed   NullableInt32   `json:"seed,omitempty"`
	Style  NullableInt32   `json:"style,omitempty"`
	Speed  NullableFloat32 `json:"speed,omitempty"`
	Pitch  NullableFloat32 `json:"pitch,omitempty"`
	Energy NullableFloat32 `json:"energy,omitempty"`
}

// NewGenerationOptions instantiates a new GenerationOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerationOptions() *GenerationOptions {
	this := GenerationOptions{}
	return &this
}

// NewGenerationOptionsWithDefaults instantiates a new GenerationOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerationOptionsWithDefaults() *GenerationOptions {
	this := GenerationOptions{}
	return &this
}

// GetSeed returns the Seed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenerationOptions) GetSeed() int32 {
	if o == nil || IsNil(o.Seed.Get()) {
		var ret int32
		return ret
	}
	return *o.Seed.Get()
}

// GetSeedOk returns a tuple with the Seed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenerationOptions) GetSeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Seed.Get(), o.Seed.IsSet()
}

// HasSeed returns a boolean if a field has been set.
func (o *GenerationOptions) HasSeed() bool {
	if o != nil && o.Seed.IsSet() {
		return true
	}

	return false
}

// SetSeed gets a reference to the given NullableInt32 and assigns it to the Seed field.
func (o *GenerationOptions) SetSeed(v int32) {
	o.Seed.Set(&v)
}

// SetSeedNil sets the value for Seed to be an explicit nil
func (o *GenerationOptions) SetSeedNil() {
	o.Seed.Set(nil)
}

// UnsetSeed ensures that no value is present for Seed, not even an explicit nil
func (o *GenerationOptions) UnsetSeed() {
	o.Seed.Unset()
}

// GetStyle returns the Style field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenerationOptions) GetStyle() int32 {
	if o == nil || IsNil(o.Style.Get()) {
		var ret int32
		return ret
	}
	return *o.Style.Get()
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenerationOptions) GetStyleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Style.Get(), o.Style.IsSet()
}

// HasStyle returns a boolean if a field has been set.
func (o *GenerationOptions) HasStyle() bool {
	if o != nil && o.Style.IsSet() {
		return true
	}

	return false
}

// SetStyle gets a reference to the given NullableInt32 and assigns it to the Style field.
func (o *GenerationOptions) SetStyle(v int32) {
	o.Style.Set(&v)
}

// SetStyleNil sets the value for Style to be an explicit nil
func (o *GenerationOptions) SetStyleNil() {
	o.Style.Set(nil)
}

// UnsetStyle ensures that no value is present for Style, not even an explicit nil
func (o *GenerationOptions) UnsetStyle() {
	o.Style.Unset()
}

// GetSpeed returns the Speed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenerationOptions) GetSpeed() float32 {
	if o == nil || IsNil(o.Speed.Get()) {
		var ret float32
		return ret
	}
	return *o.Speed.Get()
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenerationOptions) GetSpeedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Speed.Get(), o.Speed.IsSet()
}

// HasSpeed returns a boolean if a field has been set.
func (o *GenerationOptions) HasSpeed() bool {
	if o != nil && o.Speed.IsSet() {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given NullableFloat32 and assigns it to the Speed field.
func (o *GenerationOptions) SetSpeed(v float32) {
	o.Speed.Set(&v)
}

// SetSpeedNil sets the value for Speed to be an explicit nil
func (o *GenerationOptions) SetSpeedNil() {
	o.Speed.Set(nil)
}

// UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
func (o *GenerationOptions) UnsetSpeed() {
	o.Speed.Unset()
}

// GetPitch returns the Pitch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenerationOptions) GetPitch() float32 {
	if o == nil || IsNil(o.Pitch.Get()) {
		var ret float32
		return ret
	}
	return *o.Pitch.Get()
}

// GetPitchOk returns a tuple with the Pitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenerationOptions) GetPitchOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pitch.Get(), o.Pitch.IsSet()
}

// HasPitch returns a boolean if a field has been set.
func (o *GenerationOptions) HasPitch() bool {
	if o != nil && o.Pitch.IsSet() {
		return true
	}

	return false
}

// SetPitch gets a reference to the given NullableFloat32 and assigns it to the Pitch field.
func (o *GenerationOptions) SetPitch(v float32) {
	o.Pitch.Set(&v)
}

// SetPitchNil sets the value for Pitch to be an explicit nil
func (o *GenerationOptions) SetPitchNil() {
	o.Pitch.Set(nil)
}

// UnsetPitch ensures that no value is present for Pitch, not even an explicit nil
func (o *GenerationOptions) UnsetPitch() {
	o.Pitch.Unset()
}

// GetEnergy returns the Energy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenerationOptions) GetEnergy() float32 {
	if o == nil || IsNil(o.Energy.Get()) {
		var ret float32
		return ret
	}
	return *o.Energy.Get()
}

// GetEnergyOk returns a tuple with the Energy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenerationOptions) GetEnergyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Energy.Get(), o.Energy.IsSet()
}

// HasEnergy returns a boolean if a field has been set.
func (o *GenerationOptions) HasEnergy() bool {
	if o != nil && o.Energy.IsSet() {
		return true
	}

	return false
}

// SetEnergy gets a reference to the given NullableFloat32 and assigns it to the Energy field.
func (o *GenerationOptions) SetEnergy(v float32) {
	o.Energy.Set(&v)
}

// SetEnergyNil sets the value for Energy to be an explicit nil
func (o *GenerationOptions) SetEnergyNil() {
	o.Energy.Set(nil)
}

// UnsetEnergy ensures that no value is present for Energy, not even an explicit nil
func (o *GenerationOptions) UnsetEnergy() {
	o.Energy.Unset()
}

func (o GenerationOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerationOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Seed.IsSet() {
		toSerialize["seed"] = o.Seed.Get()
	}
	if o.Style.IsSet() {
		toSerialize["style"] = o.Style.Get()
	}
	if o.Speed.IsSet() {
		toSerialize["speed"] = o.Speed.Get()
	}
	if o.Pitch.IsSet() {
		toSerialize["pitch"] = o.Pitch.Get()
	}
	if o.Energy.IsSet() {
		toSerialize["energy"] = o.Energy.Get()
	}
	return toSerialize, nil
}

type NullableGenerationOptions struct {
	value *GenerationOptions
	isSet bool
}

func (v NullableGenerationOptions) Get() *GenerationOptions {
	return v.value
}

func (v *NullableGenerationOptions) Set(val *GenerationOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerationOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerationOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerationOptions(val *GenerationOptions) *NullableGenerationOptions {
	return &NullableGenerationOptions{value: val, isSet: true}
}

func (v NullableGenerationOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerationOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}