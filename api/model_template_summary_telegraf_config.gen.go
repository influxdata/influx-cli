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

// TemplateSummaryTelegrafConfig struct for TemplateSummaryTelegrafConfig
type TemplateSummaryTelegrafConfig struct {
	Id          string  `json:"id" yaml:"id"`
	Name        string  `json:"name" yaml:"name"`
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
}

// NewTemplateSummaryTelegrafConfig instantiates a new TemplateSummaryTelegrafConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryTelegrafConfig(id string, name string) *TemplateSummaryTelegrafConfig {
	this := TemplateSummaryTelegrafConfig{}
	this.Id = id
	this.Name = name
	return &this
}

// NewTemplateSummaryTelegrafConfigWithDefaults instantiates a new TemplateSummaryTelegrafConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryTelegrafConfigWithDefaults() *TemplateSummaryTelegrafConfig {
	this := TemplateSummaryTelegrafConfig{}
	return &this
}

// GetId returns the Id field value
func (o *TemplateSummaryTelegrafConfig) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryTelegrafConfig) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TemplateSummaryTelegrafConfig) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TemplateSummaryTelegrafConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryTelegrafConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TemplateSummaryTelegrafConfig) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TemplateSummaryTelegrafConfig) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryTelegrafConfig) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateSummaryTelegrafConfig) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TemplateSummaryTelegrafConfig) SetDescription(v string) {
	o.Description = &v
}

func (o TemplateSummaryTelegrafConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryTelegrafConfig struct {
	value *TemplateSummaryTelegrafConfig
	isSet bool
}

func (v NullableTemplateSummaryTelegrafConfig) Get() *TemplateSummaryTelegrafConfig {
	return v.value
}

func (v *NullableTemplateSummaryTelegrafConfig) Set(val *TemplateSummaryTelegrafConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryTelegrafConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryTelegrafConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryTelegrafConfig(val *TemplateSummaryTelegrafConfig) *NullableTemplateSummaryTelegrafConfig {
	return &NullableTemplateSummaryTelegrafConfig{value: val, isSet: true}
}

func (v NullableTemplateSummaryTelegrafConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryTelegrafConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
