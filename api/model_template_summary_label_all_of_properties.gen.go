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

// TemplateSummaryLabelAllOfProperties struct for TemplateSummaryLabelAllOfProperties
type TemplateSummaryLabelAllOfProperties struct {
	Color       *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewTemplateSummaryLabelAllOfProperties instantiates a new TemplateSummaryLabelAllOfProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryLabelAllOfProperties() *TemplateSummaryLabelAllOfProperties {
	this := TemplateSummaryLabelAllOfProperties{}
	return &this
}

// NewTemplateSummaryLabelAllOfPropertiesWithDefaults instantiates a new TemplateSummaryLabelAllOfProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryLabelAllOfPropertiesWithDefaults() *TemplateSummaryLabelAllOfProperties {
	this := TemplateSummaryLabelAllOfProperties{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *TemplateSummaryLabelAllOfProperties) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryLabelAllOfProperties) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *TemplateSummaryLabelAllOfProperties) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *TemplateSummaryLabelAllOfProperties) SetColor(v string) {
	o.Color = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TemplateSummaryLabelAllOfProperties) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryLabelAllOfProperties) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateSummaryLabelAllOfProperties) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TemplateSummaryLabelAllOfProperties) SetDescription(v string) {
	o.Description = &v
}

func (o TemplateSummaryLabelAllOfProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryLabelAllOfProperties struct {
	value *TemplateSummaryLabelAllOfProperties
	isSet bool
}

func (v NullableTemplateSummaryLabelAllOfProperties) Get() *TemplateSummaryLabelAllOfProperties {
	return v.value
}

func (v *NullableTemplateSummaryLabelAllOfProperties) Set(val *TemplateSummaryLabelAllOfProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryLabelAllOfProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryLabelAllOfProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryLabelAllOfProperties(val *TemplateSummaryLabelAllOfProperties) *NullableTemplateSummaryLabelAllOfProperties {
	return &NullableTemplateSummaryLabelAllOfProperties{value: val, isSet: true}
}

func (v NullableTemplateSummaryLabelAllOfProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryLabelAllOfProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}