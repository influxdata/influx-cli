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

// TemplateSummaryError struct for TemplateSummaryError
type TemplateSummaryError struct {
	Kind    *string   `json:"kind,omitempty"`
	Reason  *string   `json:"reason,omitempty"`
	Fields  *[]string `json:"fields,omitempty"`
	Indexes *[]int32  `json:"indexes,omitempty"`
}

// NewTemplateSummaryError instantiates a new TemplateSummaryError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryError() *TemplateSummaryError {
	this := TemplateSummaryError{}
	return &this
}

// NewTemplateSummaryErrorWithDefaults instantiates a new TemplateSummaryError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryErrorWithDefaults() *TemplateSummaryError {
	this := TemplateSummaryError{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *TemplateSummaryError) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryError) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *TemplateSummaryError) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *TemplateSummaryError) SetKind(v string) {
	o.Kind = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *TemplateSummaryError) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryError) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *TemplateSummaryError) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *TemplateSummaryError) SetReason(v string) {
	o.Reason = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *TemplateSummaryError) GetFields() []string {
	if o == nil || o.Fields == nil {
		var ret []string
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryError) GetFieldsOk() (*[]string, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *TemplateSummaryError) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []string and assigns it to the Fields field.
func (o *TemplateSummaryError) SetFields(v []string) {
	o.Fields = &v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *TemplateSummaryError) GetIndexes() []int32 {
	if o == nil || o.Indexes == nil {
		var ret []int32
		return ret
	}
	return *o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryError) GetIndexesOk() (*[]int32, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *TemplateSummaryError) HasIndexes() bool {
	if o != nil && o.Indexes != nil {
		return true
	}

	return false
}

// SetIndexes gets a reference to the given []int32 and assigns it to the Indexes field.
func (o *TemplateSummaryError) SetIndexes(v []int32) {
	o.Indexes = &v
}

func (o TemplateSummaryError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryError struct {
	value *TemplateSummaryError
	isSet bool
}

func (v NullableTemplateSummaryError) Get() *TemplateSummaryError {
	return v.value
}

func (v *NullableTemplateSummaryError) Set(val *TemplateSummaryError) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryError) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryError(val *TemplateSummaryError) *NullableTemplateSummaryError {
	return &NullableTemplateSummaryError{value: val, isSet: true}
}

func (v NullableTemplateSummaryError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}