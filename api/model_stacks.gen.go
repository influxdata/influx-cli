/*
 * Subset of Influx API covered by Influx CLI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Stacks struct for Stacks
type Stacks struct {
	Stacks []Stack `json:"stacks" yaml:"stacks"`
}

// NewStacks instantiates a new Stacks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStacks(stacks []Stack) *Stacks {
	this := Stacks{}
	this.Stacks = stacks
	return &this
}

// NewStacksWithDefaults instantiates a new Stacks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStacksWithDefaults() *Stacks {
	this := Stacks{}
	return &this
}

// GetStacks returns the Stacks field value
func (o *Stacks) GetStacks() []Stack {
	if o == nil {
		var ret []Stack
		return ret
	}

	return o.Stacks
}

// GetStacksOk returns a tuple with the Stacks field value
// and a boolean to check if the value has been set.
func (o *Stacks) GetStacksOk() (*[]Stack, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stacks, true
}

// SetStacks sets field value
func (o *Stacks) SetStacks(v []Stack) {
	o.Stacks = v
}

func (o Stacks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["stacks"] = o.Stacks
	}
	return json.Marshal(toSerialize)
}

type NullableStacks struct {
	value *Stacks
	isSet bool
}

func (v NullableStacks) Get() *Stacks {
	return v.value
}

func (v *NullableStacks) Set(val *Stacks) {
	v.value = val
	v.isSet = true
}

func (v NullableStacks) IsSet() bool {
	return v.isSet
}

func (v *NullableStacks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStacks(val *Stacks) *NullableStacks {
	return &NullableStacks{value: val, isSet: true}
}

func (v NullableStacks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStacks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
