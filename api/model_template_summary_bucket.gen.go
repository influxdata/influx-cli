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

// TemplateSummaryBucket struct for TemplateSummaryBucket
type TemplateSummaryBucket struct {
	Kind              string                 `json:"kind" yaml:"kind"`
	TemplateMetaName  string                 `json:"templateMetaName" yaml:"templateMetaName"`
	EnvReferences     []TemplateEnvReference `json:"envReferences" yaml:"envReferences"`
	LabelAssociations []TemplateSummaryLabel `json:"labelAssociations" yaml:"labelAssociations"`
	Id                uint64                 `json:"id" yaml:"id"`
	Name              string                 `json:"name" yaml:"name"`
	Description       *string                `json:"description,omitempty" yaml:"description,omitempty"`
	RetentionPeriod   int64                  `json:"retentionPeriod" yaml:"retentionPeriod"`
	SchemaType        *SchemaType            `json:"schemaType,omitempty" yaml:"schemaType,omitempty"`
}

// NewTemplateSummaryBucket instantiates a new TemplateSummaryBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryBucket(kind string, templateMetaName string, envReferences []TemplateEnvReference, labelAssociations []TemplateSummaryLabel, id uint64, name string, retentionPeriod int64) *TemplateSummaryBucket {
	this := TemplateSummaryBucket{}
	this.Kind = kind
	this.TemplateMetaName = templateMetaName
	this.EnvReferences = envReferences
	this.LabelAssociations = labelAssociations
	this.Id = id
	this.Name = name
	this.RetentionPeriod = retentionPeriod
	return &this
}

// NewTemplateSummaryBucketWithDefaults instantiates a new TemplateSummaryBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryBucketWithDefaults() *TemplateSummaryBucket {
	this := TemplateSummaryBucket{}
	return &this
}

// GetKind returns the Kind field value
func (o *TemplateSummaryBucket) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryBucket) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *TemplateSummaryBucket) SetKind(v string) {
	o.Kind = v
}

// GetTemplateMetaName returns the TemplateMetaName field value
func (o *TemplateSummaryBucket) GetTemplateMetaName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateMetaName
}

// GetTemplateMetaNameOk returns a tuple with the TemplateMetaName field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryBucket) GetTemplateMetaNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateMetaName, true
}

// SetTemplateMetaName sets field value
func (o *TemplateSummaryBucket) SetTemplateMetaName(v string) {
	o.TemplateMetaName = v
}

// GetEnvReferences returns the EnvReferences field value
func (o *TemplateSummaryBucket) GetEnvReferences() []TemplateEnvReference {
	if o == nil {
		var ret []TemplateEnvReference
		return ret
	}

	return o.EnvReferences
}

// GetEnvReferencesOk returns a tuple with the EnvReferences field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryBucket) GetEnvReferencesOk() (*[]TemplateEnvReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvReferences, true
}

// SetEnvReferences sets field value
func (o *TemplateSummaryBucket) SetEnvReferences(v []TemplateEnvReference) {
	o.EnvReferences = v
}

// GetLabelAssociations returns the LabelAssociations field value
func (o *TemplateSummaryBucket) GetLabelAssociations() []TemplateSummaryLabel {
	if o == nil {
		var ret []TemplateSummaryLabel
		return ret
	}

	return o.LabelAssociations
}

// GetLabelAssociationsOk returns a tuple with the LabelAssociations field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryBucket) GetLabelAssociationsOk() (*[]TemplateSummaryLabel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelAssociations, true
}

// SetLabelAssociations sets field value
func (o *TemplateSummaryBucket) SetLabelAssociations(v []TemplateSummaryLabel) {
	o.LabelAssociations = v
}

// GetId returns the Id field value
func (o *TemplateSummaryBucket) GetId() uint64 {
	if o == nil {
		var ret uint64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryBucket) GetIdOk() (*uint64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TemplateSummaryBucket) SetId(v uint64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TemplateSummaryBucket) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryBucket) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TemplateSummaryBucket) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TemplateSummaryBucket) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryBucket) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateSummaryBucket) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TemplateSummaryBucket) SetDescription(v string) {
	o.Description = &v
}

// GetRetentionPeriod returns the RetentionPeriod field value
func (o *TemplateSummaryBucket) GetRetentionPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RetentionPeriod
}

// GetRetentionPeriodOk returns a tuple with the RetentionPeriod field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryBucket) GetRetentionPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetentionPeriod, true
}

// SetRetentionPeriod sets field value
func (o *TemplateSummaryBucket) SetRetentionPeriod(v int64) {
	o.RetentionPeriod = v
}

// GetSchemaType returns the SchemaType field value if set, zero value otherwise.
func (o *TemplateSummaryBucket) GetSchemaType() SchemaType {
	if o == nil || o.SchemaType == nil {
		var ret SchemaType
		return ret
	}
	return *o.SchemaType
}

// GetSchemaTypeOk returns a tuple with the SchemaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryBucket) GetSchemaTypeOk() (*SchemaType, bool) {
	if o == nil || o.SchemaType == nil {
		return nil, false
	}
	return o.SchemaType, true
}

// HasSchemaType returns a boolean if a field has been set.
func (o *TemplateSummaryBucket) HasSchemaType() bool {
	if o != nil && o.SchemaType != nil {
		return true
	}

	return false
}

// SetSchemaType gets a reference to the given SchemaType and assigns it to the SchemaType field.
func (o *TemplateSummaryBucket) SetSchemaType(v SchemaType) {
	o.SchemaType = &v
}

func (o TemplateSummaryBucket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["templateMetaName"] = o.TemplateMetaName
	}
	if true {
		toSerialize["envReferences"] = o.EnvReferences
	}
	if true {
		toSerialize["labelAssociations"] = o.LabelAssociations
	}
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
		toSerialize["retentionPeriod"] = o.RetentionPeriod
	}
	if o.SchemaType != nil {
		toSerialize["schemaType"] = o.SchemaType
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryBucket struct {
	value *TemplateSummaryBucket
	isSet bool
}

func (v NullableTemplateSummaryBucket) Get() *TemplateSummaryBucket {
	return v.value
}

func (v *NullableTemplateSummaryBucket) Set(val *TemplateSummaryBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryBucket(val *TemplateSummaryBucket) *NullableTemplateSummaryBucket {
	return &NullableTemplateSummaryBucket{value: val, isSet: true}
}

func (v NullableTemplateSummaryBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
