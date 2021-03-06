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

// TemplateSummaryNotificationEndpoint struct for TemplateSummaryNotificationEndpoint
type TemplateSummaryNotificationEndpoint struct {
	Kind              string                 `json:"kind" yaml:"kind"`
	TemplateMetaName  string                 `json:"templateMetaName" yaml:"templateMetaName"`
	EnvReferences     []TemplateEnvReference `json:"envReferences" yaml:"envReferences"`
	LabelAssociations []TemplateSummaryLabel `json:"labelAssociations" yaml:"labelAssociations"`
	Id                uint64                 `json:"id" yaml:"id"`
	Name              string                 `json:"name" yaml:"name"`
	Description       *string                `json:"description,omitempty" yaml:"description,omitempty"`
	Status            string                 `json:"status" yaml:"status"`
}

// NewTemplateSummaryNotificationEndpoint instantiates a new TemplateSummaryNotificationEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryNotificationEndpoint(kind string, templateMetaName string, envReferences []TemplateEnvReference, labelAssociations []TemplateSummaryLabel, id uint64, name string, status string) *TemplateSummaryNotificationEndpoint {
	this := TemplateSummaryNotificationEndpoint{}
	this.Kind = kind
	this.TemplateMetaName = templateMetaName
	this.EnvReferences = envReferences
	this.LabelAssociations = labelAssociations
	this.Id = id
	this.Name = name
	this.Status = status
	return &this
}

// NewTemplateSummaryNotificationEndpointWithDefaults instantiates a new TemplateSummaryNotificationEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryNotificationEndpointWithDefaults() *TemplateSummaryNotificationEndpoint {
	this := TemplateSummaryNotificationEndpoint{}
	return &this
}

// GetKind returns the Kind field value
func (o *TemplateSummaryNotificationEndpoint) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryNotificationEndpoint) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *TemplateSummaryNotificationEndpoint) SetKind(v string) {
	o.Kind = v
}

// GetTemplateMetaName returns the TemplateMetaName field value
func (o *TemplateSummaryNotificationEndpoint) GetTemplateMetaName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateMetaName
}

// GetTemplateMetaNameOk returns a tuple with the TemplateMetaName field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryNotificationEndpoint) GetTemplateMetaNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateMetaName, true
}

// SetTemplateMetaName sets field value
func (o *TemplateSummaryNotificationEndpoint) SetTemplateMetaName(v string) {
	o.TemplateMetaName = v
}

// GetEnvReferences returns the EnvReferences field value
func (o *TemplateSummaryNotificationEndpoint) GetEnvReferences() []TemplateEnvReference {
	if o == nil {
		var ret []TemplateEnvReference
		return ret
	}

	return o.EnvReferences
}

// GetEnvReferencesOk returns a tuple with the EnvReferences field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryNotificationEndpoint) GetEnvReferencesOk() (*[]TemplateEnvReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvReferences, true
}

// SetEnvReferences sets field value
func (o *TemplateSummaryNotificationEndpoint) SetEnvReferences(v []TemplateEnvReference) {
	o.EnvReferences = v
}

// GetLabelAssociations returns the LabelAssociations field value
func (o *TemplateSummaryNotificationEndpoint) GetLabelAssociations() []TemplateSummaryLabel {
	if o == nil {
		var ret []TemplateSummaryLabel
		return ret
	}

	return o.LabelAssociations
}

// GetLabelAssociationsOk returns a tuple with the LabelAssociations field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryNotificationEndpoint) GetLabelAssociationsOk() (*[]TemplateSummaryLabel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelAssociations, true
}

// SetLabelAssociations sets field value
func (o *TemplateSummaryNotificationEndpoint) SetLabelAssociations(v []TemplateSummaryLabel) {
	o.LabelAssociations = v
}

// GetId returns the Id field value
func (o *TemplateSummaryNotificationEndpoint) GetId() uint64 {
	if o == nil {
		var ret uint64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryNotificationEndpoint) GetIdOk() (*uint64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TemplateSummaryNotificationEndpoint) SetId(v uint64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TemplateSummaryNotificationEndpoint) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryNotificationEndpoint) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TemplateSummaryNotificationEndpoint) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TemplateSummaryNotificationEndpoint) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryNotificationEndpoint) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateSummaryNotificationEndpoint) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TemplateSummaryNotificationEndpoint) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value
func (o *TemplateSummaryNotificationEndpoint) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryNotificationEndpoint) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TemplateSummaryNotificationEndpoint) SetStatus(v string) {
	o.Status = v
}

func (o TemplateSummaryNotificationEndpoint) MarshalJSON() ([]byte, error) {
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
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryNotificationEndpoint struct {
	value *TemplateSummaryNotificationEndpoint
	isSet bool
}

func (v NullableTemplateSummaryNotificationEndpoint) Get() *TemplateSummaryNotificationEndpoint {
	return v.value
}

func (v *NullableTemplateSummaryNotificationEndpoint) Set(val *TemplateSummaryNotificationEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryNotificationEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryNotificationEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryNotificationEndpoint(val *TemplateSummaryNotificationEndpoint) *NullableTemplateSummaryNotificationEndpoint {
	return &NullableTemplateSummaryNotificationEndpoint{value: val, isSet: true}
}

func (v NullableTemplateSummaryNotificationEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryNotificationEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
