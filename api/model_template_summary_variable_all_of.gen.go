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

// TemplateSummaryVariableAllOf struct for TemplateSummaryVariableAllOf
type TemplateSummaryVariableAllOf struct {
	Id          string                      `json:"id" yaml:"id"`
	Name        string                      `json:"name" yaml:"name"`
	Description *string                     `json:"description,omitempty" yaml:"description,omitempty"`
	Arguments   TemplateSummaryVariableArgs `json:"arguments" yaml:"arguments"`
}

// NewTemplateSummaryVariableAllOf instantiates a new TemplateSummaryVariableAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryVariableAllOf(id string, name string, arguments TemplateSummaryVariableArgs) *TemplateSummaryVariableAllOf {
	this := TemplateSummaryVariableAllOf{}
	this.Id = id
	this.Name = name
	this.Arguments = arguments
	return &this
}

// NewTemplateSummaryVariableAllOfWithDefaults instantiates a new TemplateSummaryVariableAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryVariableAllOfWithDefaults() *TemplateSummaryVariableAllOf {
	this := TemplateSummaryVariableAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *TemplateSummaryVariableAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryVariableAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TemplateSummaryVariableAllOf) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TemplateSummaryVariableAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryVariableAllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TemplateSummaryVariableAllOf) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TemplateSummaryVariableAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryVariableAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateSummaryVariableAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TemplateSummaryVariableAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetArguments returns the Arguments field value
func (o *TemplateSummaryVariableAllOf) GetArguments() TemplateSummaryVariableArgs {
	if o == nil {
		var ret TemplateSummaryVariableArgs
		return ret
	}

	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryVariableAllOf) GetArgumentsOk() (*TemplateSummaryVariableArgs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Arguments, true
}

// SetArguments sets field value
func (o *TemplateSummaryVariableAllOf) SetArguments(v TemplateSummaryVariableArgs) {
	o.Arguments = v
}

func (o TemplateSummaryVariableAllOf) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["arguments"] = o.Arguments
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryVariableAllOf struct {
	value *TemplateSummaryVariableAllOf
	isSet bool
}

func (v NullableTemplateSummaryVariableAllOf) Get() *TemplateSummaryVariableAllOf {
	return v.value
}

func (v *NullableTemplateSummaryVariableAllOf) Set(val *TemplateSummaryVariableAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryVariableAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryVariableAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryVariableAllOf(val *TemplateSummaryVariableAllOf) *NullableTemplateSummaryVariableAllOf {
	return &NullableTemplateSummaryVariableAllOf{value: val, isSet: true}
}

func (v NullableTemplateSummaryVariableAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryVariableAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
