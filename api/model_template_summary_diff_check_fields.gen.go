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

// TemplateSummaryDiffCheckFields struct for TemplateSummaryDiffCheckFields
type TemplateSummaryDiffCheckFields struct {
	Name        string  `json:"name" yaml:"name"`
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
}

// NewTemplateSummaryDiffCheckFields instantiates a new TemplateSummaryDiffCheckFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryDiffCheckFields(name string) *TemplateSummaryDiffCheckFields {
	this := TemplateSummaryDiffCheckFields{}
	this.Name = name
	return &this
}

// NewTemplateSummaryDiffCheckFieldsWithDefaults instantiates a new TemplateSummaryDiffCheckFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryDiffCheckFieldsWithDefaults() *TemplateSummaryDiffCheckFields {
	this := TemplateSummaryDiffCheckFields{}
	return &this
}

// GetName returns the Name field value
func (o *TemplateSummaryDiffCheckFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffCheckFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TemplateSummaryDiffCheckFields) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TemplateSummaryDiffCheckFields) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffCheckFields) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateSummaryDiffCheckFields) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TemplateSummaryDiffCheckFields) SetDescription(v string) {
	o.Description = &v
}

func (o TemplateSummaryDiffCheckFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryDiffCheckFields struct {
	value *TemplateSummaryDiffCheckFields
	isSet bool
}

func (v NullableTemplateSummaryDiffCheckFields) Get() *TemplateSummaryDiffCheckFields {
	return v.value
}

func (v *NullableTemplateSummaryDiffCheckFields) Set(val *TemplateSummaryDiffCheckFields) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryDiffCheckFields) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryDiffCheckFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryDiffCheckFields(val *TemplateSummaryDiffCheckFields) *NullableTemplateSummaryDiffCheckFields {
	return &NullableTemplateSummaryDiffCheckFields{value: val, isSet: true}
}

func (v NullableTemplateSummaryDiffCheckFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryDiffCheckFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
