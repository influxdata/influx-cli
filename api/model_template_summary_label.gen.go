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

// TemplateSummaryLabel struct for TemplateSummaryLabel
type TemplateSummaryLabel struct {
	Kind             string                              `json:"kind" yaml:"kind"`
	TemplateMetaName *string                             `json:"templateMetaName,omitempty" yaml:"templateMetaName,omitempty"`
	EnvReferences    []TemplateEnvReference              `json:"envReferences" yaml:"envReferences"`
	Id               uint64                              `json:"id" yaml:"id"`
	OrgID            *uint64                             `json:"orgID,omitempty" yaml:"orgID,omitempty"`
	Name             string                              `json:"name" yaml:"name"`
	Properties       TemplateSummaryLabelAllOfProperties `json:"properties" yaml:"properties"`
}

// NewTemplateSummaryLabel instantiates a new TemplateSummaryLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryLabel(kind string, envReferences []TemplateEnvReference, id uint64, name string, properties TemplateSummaryLabelAllOfProperties) *TemplateSummaryLabel {
	this := TemplateSummaryLabel{}
	this.Kind = kind
	this.EnvReferences = envReferences
	this.Id = id
	this.Name = name
	this.Properties = properties
	return &this
}

// NewTemplateSummaryLabelWithDefaults instantiates a new TemplateSummaryLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryLabelWithDefaults() *TemplateSummaryLabel {
	this := TemplateSummaryLabel{}
	return &this
}

// GetKind returns the Kind field value
func (o *TemplateSummaryLabel) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryLabel) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *TemplateSummaryLabel) SetKind(v string) {
	o.Kind = v
}

// GetTemplateMetaName returns the TemplateMetaName field value if set, zero value otherwise.
func (o *TemplateSummaryLabel) GetTemplateMetaName() string {
	if o == nil || o.TemplateMetaName == nil {
		var ret string
		return ret
	}
	return *o.TemplateMetaName
}

// GetTemplateMetaNameOk returns a tuple with the TemplateMetaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryLabel) GetTemplateMetaNameOk() (*string, bool) {
	if o == nil || o.TemplateMetaName == nil {
		return nil, false
	}
	return o.TemplateMetaName, true
}

// HasTemplateMetaName returns a boolean if a field has been set.
func (o *TemplateSummaryLabel) HasTemplateMetaName() bool {
	if o != nil && o.TemplateMetaName != nil {
		return true
	}

	return false
}

// SetTemplateMetaName gets a reference to the given string and assigns it to the TemplateMetaName field.
func (o *TemplateSummaryLabel) SetTemplateMetaName(v string) {
	o.TemplateMetaName = &v
}

// GetEnvReferences returns the EnvReferences field value
func (o *TemplateSummaryLabel) GetEnvReferences() []TemplateEnvReference {
	if o == nil {
		var ret []TemplateEnvReference
		return ret
	}

	return o.EnvReferences
}

// GetEnvReferencesOk returns a tuple with the EnvReferences field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryLabel) GetEnvReferencesOk() (*[]TemplateEnvReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvReferences, true
}

// SetEnvReferences sets field value
func (o *TemplateSummaryLabel) SetEnvReferences(v []TemplateEnvReference) {
	o.EnvReferences = v
}

// GetId returns the Id field value
func (o *TemplateSummaryLabel) GetId() uint64 {
	if o == nil {
		var ret uint64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryLabel) GetIdOk() (*uint64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TemplateSummaryLabel) SetId(v uint64) {
	o.Id = v
}

// GetOrgID returns the OrgID field value if set, zero value otherwise.
func (o *TemplateSummaryLabel) GetOrgID() uint64 {
	if o == nil || o.OrgID == nil {
		var ret uint64
		return ret
	}
	return *o.OrgID
}

// GetOrgIDOk returns a tuple with the OrgID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryLabel) GetOrgIDOk() (*uint64, bool) {
	if o == nil || o.OrgID == nil {
		return nil, false
	}
	return o.OrgID, true
}

// HasOrgID returns a boolean if a field has been set.
func (o *TemplateSummaryLabel) HasOrgID() bool {
	if o != nil && o.OrgID != nil {
		return true
	}

	return false
}

// SetOrgID gets a reference to the given int64 and assigns it to the OrgID field.
func (o *TemplateSummaryLabel) SetOrgID(v uint64) {
	o.OrgID = &v
}

// GetName returns the Name field value
func (o *TemplateSummaryLabel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryLabel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TemplateSummaryLabel) SetName(v string) {
	o.Name = v
}

// GetProperties returns the Properties field value
func (o *TemplateSummaryLabel) GetProperties() TemplateSummaryLabelAllOfProperties {
	if o == nil {
		var ret TemplateSummaryLabelAllOfProperties
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryLabel) GetPropertiesOk() (*TemplateSummaryLabelAllOfProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *TemplateSummaryLabel) SetProperties(v TemplateSummaryLabelAllOfProperties) {
	o.Properties = v
}

func (o TemplateSummaryLabel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if o.TemplateMetaName != nil {
		toSerialize["templateMetaName"] = o.TemplateMetaName
	}
	if true {
		toSerialize["envReferences"] = o.EnvReferences
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.OrgID != nil {
		toSerialize["orgID"] = o.OrgID
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryLabel struct {
	value *TemplateSummaryLabel
	isSet bool
}

func (v NullableTemplateSummaryLabel) Get() *TemplateSummaryLabel {
	return v.value
}

func (v *NullableTemplateSummaryLabel) Set(val *TemplateSummaryLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryLabel(val *TemplateSummaryLabel) *NullableTemplateSummaryLabel {
	return &NullableTemplateSummaryLabel{value: val, isSet: true}
}

func (v NullableTemplateSummaryLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
