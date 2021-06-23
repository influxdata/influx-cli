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

// TemplateSummaryDiffNotificationEndpointFields struct for TemplateSummaryDiffNotificationEndpointFields
type TemplateSummaryDiffNotificationEndpointFields struct {
	Name *string `json:"name,omitempty"`
}

// NewTemplateSummaryDiffNotificationEndpointFields instantiates a new TemplateSummaryDiffNotificationEndpointFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryDiffNotificationEndpointFields() *TemplateSummaryDiffNotificationEndpointFields {
	this := TemplateSummaryDiffNotificationEndpointFields{}
	return &this
}

// NewTemplateSummaryDiffNotificationEndpointFieldsWithDefaults instantiates a new TemplateSummaryDiffNotificationEndpointFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryDiffNotificationEndpointFieldsWithDefaults() *TemplateSummaryDiffNotificationEndpointFields {
	this := TemplateSummaryDiffNotificationEndpointFields{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplateSummaryDiffNotificationEndpointFields) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffNotificationEndpointFields) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplateSummaryDiffNotificationEndpointFields) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplateSummaryDiffNotificationEndpointFields) SetName(v string) {
	o.Name = &v
}

func (o TemplateSummaryDiffNotificationEndpointFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryDiffNotificationEndpointFields struct {
	value *TemplateSummaryDiffNotificationEndpointFields
	isSet bool
}

func (v NullableTemplateSummaryDiffNotificationEndpointFields) Get() *TemplateSummaryDiffNotificationEndpointFields {
	return v.value
}

func (v *NullableTemplateSummaryDiffNotificationEndpointFields) Set(val *TemplateSummaryDiffNotificationEndpointFields) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryDiffNotificationEndpointFields) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryDiffNotificationEndpointFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryDiffNotificationEndpointFields(val *TemplateSummaryDiffNotificationEndpointFields) *NullableTemplateSummaryDiffNotificationEndpointFields {
	return &NullableTemplateSummaryDiffNotificationEndpointFields{value: val, isSet: true}
}

func (v NullableTemplateSummaryDiffNotificationEndpointFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryDiffNotificationEndpointFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
