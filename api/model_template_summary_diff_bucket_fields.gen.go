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

// TemplateSummaryDiffBucketFields struct for TemplateSummaryDiffBucketFields
type TemplateSummaryDiffBucketFields struct {
	Name        string  `json:"name" yaml:"name"`
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	// Rules to expire or retain data.  No rules means data never expires.
	RetentionRules     []RetentionRule          `json:"retentionRules" yaml:"retentionRules"`
	SchemaType         *SchemaType              `json:"schemaType,omitempty" yaml:"schemaType,omitempty"`
	MeasurementSchemas []map[string]interface{} `json:"measurementSchemas" yaml:"measurementSchemas"`
}

// NewTemplateSummaryDiffBucketFields instantiates a new TemplateSummaryDiffBucketFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryDiffBucketFields(name string, retentionRules []RetentionRule, measurementSchemas []map[string]interface{}) *TemplateSummaryDiffBucketFields {
	this := TemplateSummaryDiffBucketFields{}
	this.Name = name
	this.RetentionRules = retentionRules
	this.MeasurementSchemas = measurementSchemas
	return &this
}

// NewTemplateSummaryDiffBucketFieldsWithDefaults instantiates a new TemplateSummaryDiffBucketFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryDiffBucketFieldsWithDefaults() *TemplateSummaryDiffBucketFields {
	this := TemplateSummaryDiffBucketFields{}
	return &this
}

// GetName returns the Name field value
func (o *TemplateSummaryDiffBucketFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffBucketFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TemplateSummaryDiffBucketFields) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TemplateSummaryDiffBucketFields) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffBucketFields) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateSummaryDiffBucketFields) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TemplateSummaryDiffBucketFields) SetDescription(v string) {
	o.Description = &v
}

// GetRetentionRules returns the RetentionRules field value
func (o *TemplateSummaryDiffBucketFields) GetRetentionRules() []RetentionRule {
	if o == nil {
		var ret []RetentionRule
		return ret
	}

	return o.RetentionRules
}

// GetRetentionRulesOk returns a tuple with the RetentionRules field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffBucketFields) GetRetentionRulesOk() (*[]RetentionRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetentionRules, true
}

// SetRetentionRules sets field value
func (o *TemplateSummaryDiffBucketFields) SetRetentionRules(v []RetentionRule) {
	o.RetentionRules = v
}

// GetSchemaType returns the SchemaType field value if set, zero value otherwise.
func (o *TemplateSummaryDiffBucketFields) GetSchemaType() SchemaType {
	if o == nil || o.SchemaType == nil {
		var ret SchemaType
		return ret
	}
	return *o.SchemaType
}

// GetSchemaTypeOk returns a tuple with the SchemaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffBucketFields) GetSchemaTypeOk() (*SchemaType, bool) {
	if o == nil || o.SchemaType == nil {
		return nil, false
	}
	return o.SchemaType, true
}

// HasSchemaType returns a boolean if a field has been set.
func (o *TemplateSummaryDiffBucketFields) HasSchemaType() bool {
	if o != nil && o.SchemaType != nil {
		return true
	}

	return false
}

// SetSchemaType gets a reference to the given SchemaType and assigns it to the SchemaType field.
func (o *TemplateSummaryDiffBucketFields) SetSchemaType(v SchemaType) {
	o.SchemaType = &v
}

// GetMeasurementSchemas returns the MeasurementSchemas field value
func (o *TemplateSummaryDiffBucketFields) GetMeasurementSchemas() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.MeasurementSchemas
}

// GetMeasurementSchemasOk returns a tuple with the MeasurementSchemas field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffBucketFields) GetMeasurementSchemasOk() (*[]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeasurementSchemas, true
}

// SetMeasurementSchemas sets field value
func (o *TemplateSummaryDiffBucketFields) SetMeasurementSchemas(v []map[string]interface{}) {
	o.MeasurementSchemas = v
}

func (o TemplateSummaryDiffBucketFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["retentionRules"] = o.RetentionRules
	}
	if o.SchemaType != nil {
		toSerialize["schemaType"] = o.SchemaType
	}
	if true {
		toSerialize["measurementSchemas"] = o.MeasurementSchemas
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryDiffBucketFields struct {
	value *TemplateSummaryDiffBucketFields
	isSet bool
}

func (v NullableTemplateSummaryDiffBucketFields) Get() *TemplateSummaryDiffBucketFields {
	return v.value
}

func (v *NullableTemplateSummaryDiffBucketFields) Set(val *TemplateSummaryDiffBucketFields) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryDiffBucketFields) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryDiffBucketFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryDiffBucketFields(val *TemplateSummaryDiffBucketFields) *NullableTemplateSummaryDiffBucketFields {
	return &NullableTemplateSummaryDiffBucketFields{value: val, isSet: true}
}

func (v NullableTemplateSummaryDiffBucketFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryDiffBucketFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
