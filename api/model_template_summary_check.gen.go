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

// TemplateSummaryCheck struct for TemplateSummaryCheck
type TemplateSummaryCheck struct {
	Kind              *string                 `json:"kind,omitempty"`
	TemplateMetaName  *string                 `json:"templateMetaName,omitempty"`
	EnvReferences     *[]TemplateEnvReference `json:"envReferences,omitempty"`
	LabelAssociations *[]TemplateSummaryLabel `json:"labelAssociations,omitempty"`
	Id                *string                 `json:"id,omitempty"`
	Name              *string                 `json:"name,omitempty"`
	Description       *string                 `json:"description,omitempty"`
}

// NewTemplateSummaryCheck instantiates a new TemplateSummaryCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryCheck() *TemplateSummaryCheck {
	this := TemplateSummaryCheck{}
	return &this
}

// NewTemplateSummaryCheckWithDefaults instantiates a new TemplateSummaryCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryCheckWithDefaults() *TemplateSummaryCheck {
	this := TemplateSummaryCheck{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *TemplateSummaryCheck) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryCheck) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *TemplateSummaryCheck) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *TemplateSummaryCheck) SetKind(v string) {
	o.Kind = &v
}

// GetTemplateMetaName returns the TemplateMetaName field value if set, zero value otherwise.
func (o *TemplateSummaryCheck) GetTemplateMetaName() string {
	if o == nil || o.TemplateMetaName == nil {
		var ret string
		return ret
	}
	return *o.TemplateMetaName
}

// GetTemplateMetaNameOk returns a tuple with the TemplateMetaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryCheck) GetTemplateMetaNameOk() (*string, bool) {
	if o == nil || o.TemplateMetaName == nil {
		return nil, false
	}
	return o.TemplateMetaName, true
}

// HasTemplateMetaName returns a boolean if a field has been set.
func (o *TemplateSummaryCheck) HasTemplateMetaName() bool {
	if o != nil && o.TemplateMetaName != nil {
		return true
	}

	return false
}

// SetTemplateMetaName gets a reference to the given string and assigns it to the TemplateMetaName field.
func (o *TemplateSummaryCheck) SetTemplateMetaName(v string) {
	o.TemplateMetaName = &v
}

// GetEnvReferences returns the EnvReferences field value if set, zero value otherwise.
func (o *TemplateSummaryCheck) GetEnvReferences() []TemplateEnvReference {
	if o == nil || o.EnvReferences == nil {
		var ret []TemplateEnvReference
		return ret
	}
	return *o.EnvReferences
}

// GetEnvReferencesOk returns a tuple with the EnvReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryCheck) GetEnvReferencesOk() (*[]TemplateEnvReference, bool) {
	if o == nil || o.EnvReferences == nil {
		return nil, false
	}
	return o.EnvReferences, true
}

// HasEnvReferences returns a boolean if a field has been set.
func (o *TemplateSummaryCheck) HasEnvReferences() bool {
	if o != nil && o.EnvReferences != nil {
		return true
	}

	return false
}

// SetEnvReferences gets a reference to the given []TemplateEnvReference and assigns it to the EnvReferences field.
func (o *TemplateSummaryCheck) SetEnvReferences(v []TemplateEnvReference) {
	o.EnvReferences = &v
}

// GetLabelAssociations returns the LabelAssociations field value if set, zero value otherwise.
func (o *TemplateSummaryCheck) GetLabelAssociations() []TemplateSummaryLabel {
	if o == nil || o.LabelAssociations == nil {
		var ret []TemplateSummaryLabel
		return ret
	}
	return *o.LabelAssociations
}

// GetLabelAssociationsOk returns a tuple with the LabelAssociations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryCheck) GetLabelAssociationsOk() (*[]TemplateSummaryLabel, bool) {
	if o == nil || o.LabelAssociations == nil {
		return nil, false
	}
	return o.LabelAssociations, true
}

// HasLabelAssociations returns a boolean if a field has been set.
func (o *TemplateSummaryCheck) HasLabelAssociations() bool {
	if o != nil && o.LabelAssociations != nil {
		return true
	}

	return false
}

// SetLabelAssociations gets a reference to the given []TemplateSummaryLabel and assigns it to the LabelAssociations field.
func (o *TemplateSummaryCheck) SetLabelAssociations(v []TemplateSummaryLabel) {
	o.LabelAssociations = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TemplateSummaryCheck) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryCheck) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TemplateSummaryCheck) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TemplateSummaryCheck) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplateSummaryCheck) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryCheck) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplateSummaryCheck) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplateSummaryCheck) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TemplateSummaryCheck) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryCheck) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateSummaryCheck) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TemplateSummaryCheck) SetDescription(v string) {
	o.Description = &v
}

func (o TemplateSummaryCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.TemplateMetaName != nil {
		toSerialize["templateMetaName"] = o.TemplateMetaName
	}
	if o.EnvReferences != nil {
		toSerialize["envReferences"] = o.EnvReferences
	}
	if o.LabelAssociations != nil {
		toSerialize["labelAssociations"] = o.LabelAssociations
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryCheck struct {
	value *TemplateSummaryCheck
	isSet bool
}

func (v NullableTemplateSummaryCheck) Get() *TemplateSummaryCheck {
	return v.value
}

func (v *NullableTemplateSummaryCheck) Set(val *TemplateSummaryCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryCheck(val *TemplateSummaryCheck) *NullableTemplateSummaryCheck {
	return &NullableTemplateSummaryCheck{value: val, isSet: true}
}

func (v NullableTemplateSummaryCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
