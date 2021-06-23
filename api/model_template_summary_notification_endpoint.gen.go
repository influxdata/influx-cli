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
	Kind              *string                 `json:"kind,omitempty"`
	TemplateMetaName  *string                 `json:"templateMetaName,omitempty"`
	EnvReferences     *[]TemplateEnvReference `json:"envReferences,omitempty"`
	LabelAssociations *[]TemplateSummaryLabel `json:"labelAssociations,omitempty"`
	Id                *string                 `json:"id,omitempty"`
	Name              *string                 `json:"name,omitempty"`
	Description       *string                 `json:"description,omitempty"`
	Status            *string                 `json:"status,omitempty"`
}

// NewTemplateSummaryNotificationEndpoint instantiates a new TemplateSummaryNotificationEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryNotificationEndpoint() *TemplateSummaryNotificationEndpoint {
	this := TemplateSummaryNotificationEndpoint{}
	return &this
}

// NewTemplateSummaryNotificationEndpointWithDefaults instantiates a new TemplateSummaryNotificationEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryNotificationEndpointWithDefaults() *TemplateSummaryNotificationEndpoint {
	this := TemplateSummaryNotificationEndpoint{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *TemplateSummaryNotificationEndpoint) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryNotificationEndpoint) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *TemplateSummaryNotificationEndpoint) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *TemplateSummaryNotificationEndpoint) SetKind(v string) {
	o.Kind = &v
}

// GetTemplateMetaName returns the TemplateMetaName field value if set, zero value otherwise.
func (o *TemplateSummaryNotificationEndpoint) GetTemplateMetaName() string {
	if o == nil || o.TemplateMetaName == nil {
		var ret string
		return ret
	}
	return *o.TemplateMetaName
}

// GetTemplateMetaNameOk returns a tuple with the TemplateMetaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryNotificationEndpoint) GetTemplateMetaNameOk() (*string, bool) {
	if o == nil || o.TemplateMetaName == nil {
		return nil, false
	}
	return o.TemplateMetaName, true
}

// HasTemplateMetaName returns a boolean if a field has been set.
func (o *TemplateSummaryNotificationEndpoint) HasTemplateMetaName() bool {
	if o != nil && o.TemplateMetaName != nil {
		return true
	}

	return false
}

// SetTemplateMetaName gets a reference to the given string and assigns it to the TemplateMetaName field.
func (o *TemplateSummaryNotificationEndpoint) SetTemplateMetaName(v string) {
	o.TemplateMetaName = &v
}

// GetEnvReferences returns the EnvReferences field value if set, zero value otherwise.
func (o *TemplateSummaryNotificationEndpoint) GetEnvReferences() []TemplateEnvReference {
	if o == nil || o.EnvReferences == nil {
		var ret []TemplateEnvReference
		return ret
	}
	return *o.EnvReferences
}

// GetEnvReferencesOk returns a tuple with the EnvReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryNotificationEndpoint) GetEnvReferencesOk() (*[]TemplateEnvReference, bool) {
	if o == nil || o.EnvReferences == nil {
		return nil, false
	}
	return o.EnvReferences, true
}

// HasEnvReferences returns a boolean if a field has been set.
func (o *TemplateSummaryNotificationEndpoint) HasEnvReferences() bool {
	if o != nil && o.EnvReferences != nil {
		return true
	}

	return false
}

// SetEnvReferences gets a reference to the given []TemplateEnvReference and assigns it to the EnvReferences field.
func (o *TemplateSummaryNotificationEndpoint) SetEnvReferences(v []TemplateEnvReference) {
	o.EnvReferences = &v
}

// GetLabelAssociations returns the LabelAssociations field value if set, zero value otherwise.
func (o *TemplateSummaryNotificationEndpoint) GetLabelAssociations() []TemplateSummaryLabel {
	if o == nil || o.LabelAssociations == nil {
		var ret []TemplateSummaryLabel
		return ret
	}
	return *o.LabelAssociations
}

// GetLabelAssociationsOk returns a tuple with the LabelAssociations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryNotificationEndpoint) GetLabelAssociationsOk() (*[]TemplateSummaryLabel, bool) {
	if o == nil || o.LabelAssociations == nil {
		return nil, false
	}
	return o.LabelAssociations, true
}

// HasLabelAssociations returns a boolean if a field has been set.
func (o *TemplateSummaryNotificationEndpoint) HasLabelAssociations() bool {
	if o != nil && o.LabelAssociations != nil {
		return true
	}

	return false
}

// SetLabelAssociations gets a reference to the given []TemplateSummaryLabel and assigns it to the LabelAssociations field.
func (o *TemplateSummaryNotificationEndpoint) SetLabelAssociations(v []TemplateSummaryLabel) {
	o.LabelAssociations = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TemplateSummaryNotificationEndpoint) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryNotificationEndpoint) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TemplateSummaryNotificationEndpoint) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TemplateSummaryNotificationEndpoint) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplateSummaryNotificationEndpoint) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryNotificationEndpoint) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplateSummaryNotificationEndpoint) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplateSummaryNotificationEndpoint) SetName(v string) {
	o.Name = &v
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

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TemplateSummaryNotificationEndpoint) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryNotificationEndpoint) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TemplateSummaryNotificationEndpoint) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TemplateSummaryNotificationEndpoint) SetStatus(v string) {
	o.Status = &v
}

func (o TemplateSummaryNotificationEndpoint) MarshalJSON() ([]byte, error) {
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
	if o.Status != nil {
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
