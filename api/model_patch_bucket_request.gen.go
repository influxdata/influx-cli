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

// PatchBucketRequest An object that contains updated bucket properties to apply.
type PatchBucketRequest struct {
	// The name of the bucket.
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
	// A description of the bucket.
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	// Updates to rules to expire or retain data. No rules means no updates.
	RetentionRules *[]PatchRetentionRule `json:"retentionRules,omitempty" yaml:"retentionRules,omitempty"`
}

// NewPatchBucketRequest instantiates a new PatchBucketRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchBucketRequest() *PatchBucketRequest {
	this := PatchBucketRequest{}
	return &this
}

// NewPatchBucketRequestWithDefaults instantiates a new PatchBucketRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchBucketRequestWithDefaults() *PatchBucketRequest {
	this := PatchBucketRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchBucketRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBucketRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchBucketRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchBucketRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchBucketRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBucketRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchBucketRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchBucketRequest) SetDescription(v string) {
	o.Description = &v
}

// GetRetentionRules returns the RetentionRules field value if set, zero value otherwise.
func (o *PatchBucketRequest) GetRetentionRules() []PatchRetentionRule {
	if o == nil || o.RetentionRules == nil {
		var ret []PatchRetentionRule
		return ret
	}
	return *o.RetentionRules
}

// GetRetentionRulesOk returns a tuple with the RetentionRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBucketRequest) GetRetentionRulesOk() (*[]PatchRetentionRule, bool) {
	if o == nil || o.RetentionRules == nil {
		return nil, false
	}
	return o.RetentionRules, true
}

// HasRetentionRules returns a boolean if a field has been set.
func (o *PatchBucketRequest) HasRetentionRules() bool {
	if o != nil && o.RetentionRules != nil {
		return true
	}

	return false
}

// SetRetentionRules gets a reference to the given []PatchRetentionRule and assigns it to the RetentionRules field.
func (o *PatchBucketRequest) SetRetentionRules(v []PatchRetentionRule) {
	o.RetentionRules = &v
}

func (o PatchBucketRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.RetentionRules != nil {
		toSerialize["retentionRules"] = o.RetentionRules
	}
	return json.Marshal(toSerialize)
}

type NullablePatchBucketRequest struct {
	value *PatchBucketRequest
	isSet bool
}

func (v NullablePatchBucketRequest) Get() *PatchBucketRequest {
	return v.value
}

func (v *NullablePatchBucketRequest) Set(val *PatchBucketRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchBucketRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchBucketRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchBucketRequest(val *PatchBucketRequest) *NullablePatchBucketRequest {
	return &NullablePatchBucketRequest{value: val, isSet: true}
}

func (v NullablePatchBucketRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchBucketRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
