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

// TemplateSummaryDiffLabel struct for TemplateSummaryDiffLabel
type TemplateSummaryDiffLabel struct {
	StateStatus      *string                         `json:"stateStatus,omitempty"`
	Kind             *string                         `json:"kind,omitempty"`
	Id               *string                         `json:"id,omitempty"`
	TemplateMetaName *string                         `json:"templateMetaName,omitempty"`
	New              *TemplateSummaryDiffLabelFields `json:"new,omitempty"`
	Old              *TemplateSummaryDiffLabelFields `json:"old,omitempty"`
}

// NewTemplateSummaryDiffLabel instantiates a new TemplateSummaryDiffLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryDiffLabel() *TemplateSummaryDiffLabel {
	this := TemplateSummaryDiffLabel{}
	return &this
}

// NewTemplateSummaryDiffLabelWithDefaults instantiates a new TemplateSummaryDiffLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryDiffLabelWithDefaults() *TemplateSummaryDiffLabel {
	this := TemplateSummaryDiffLabel{}
	return &this
}

// GetStateStatus returns the StateStatus field value if set, zero value otherwise.
func (o *TemplateSummaryDiffLabel) GetStateStatus() string {
	if o == nil || o.StateStatus == nil {
		var ret string
		return ret
	}
	return *o.StateStatus
}

// GetStateStatusOk returns a tuple with the StateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffLabel) GetStateStatusOk() (*string, bool) {
	if o == nil || o.StateStatus == nil {
		return nil, false
	}
	return o.StateStatus, true
}

// HasStateStatus returns a boolean if a field has been set.
func (o *TemplateSummaryDiffLabel) HasStateStatus() bool {
	if o != nil && o.StateStatus != nil {
		return true
	}

	return false
}

// SetStateStatus gets a reference to the given string and assigns it to the StateStatus field.
func (o *TemplateSummaryDiffLabel) SetStateStatus(v string) {
	o.StateStatus = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *TemplateSummaryDiffLabel) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffLabel) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *TemplateSummaryDiffLabel) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *TemplateSummaryDiffLabel) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TemplateSummaryDiffLabel) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffLabel) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TemplateSummaryDiffLabel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TemplateSummaryDiffLabel) SetId(v string) {
	o.Id = &v
}

// GetTemplateMetaName returns the TemplateMetaName field value if set, zero value otherwise.
func (o *TemplateSummaryDiffLabel) GetTemplateMetaName() string {
	if o == nil || o.TemplateMetaName == nil {
		var ret string
		return ret
	}
	return *o.TemplateMetaName
}

// GetTemplateMetaNameOk returns a tuple with the TemplateMetaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffLabel) GetTemplateMetaNameOk() (*string, bool) {
	if o == nil || o.TemplateMetaName == nil {
		return nil, false
	}
	return o.TemplateMetaName, true
}

// HasTemplateMetaName returns a boolean if a field has been set.
func (o *TemplateSummaryDiffLabel) HasTemplateMetaName() bool {
	if o != nil && o.TemplateMetaName != nil {
		return true
	}

	return false
}

// SetTemplateMetaName gets a reference to the given string and assigns it to the TemplateMetaName field.
func (o *TemplateSummaryDiffLabel) SetTemplateMetaName(v string) {
	o.TemplateMetaName = &v
}

// GetNew returns the New field value if set, zero value otherwise.
func (o *TemplateSummaryDiffLabel) GetNew() TemplateSummaryDiffLabelFields {
	if o == nil || o.New == nil {
		var ret TemplateSummaryDiffLabelFields
		return ret
	}
	return *o.New
}

// GetNewOk returns a tuple with the New field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffLabel) GetNewOk() (*TemplateSummaryDiffLabelFields, bool) {
	if o == nil || o.New == nil {
		return nil, false
	}
	return o.New, true
}

// HasNew returns a boolean if a field has been set.
func (o *TemplateSummaryDiffLabel) HasNew() bool {
	if o != nil && o.New != nil {
		return true
	}

	return false
}

// SetNew gets a reference to the given TemplateSummaryDiffLabelFields and assigns it to the New field.
func (o *TemplateSummaryDiffLabel) SetNew(v TemplateSummaryDiffLabelFields) {
	o.New = &v
}

// GetOld returns the Old field value if set, zero value otherwise.
func (o *TemplateSummaryDiffLabel) GetOld() TemplateSummaryDiffLabelFields {
	if o == nil || o.Old == nil {
		var ret TemplateSummaryDiffLabelFields
		return ret
	}
	return *o.Old
}

// GetOldOk returns a tuple with the Old field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffLabel) GetOldOk() (*TemplateSummaryDiffLabelFields, bool) {
	if o == nil || o.Old == nil {
		return nil, false
	}
	return o.Old, true
}

// HasOld returns a boolean if a field has been set.
func (o *TemplateSummaryDiffLabel) HasOld() bool {
	if o != nil && o.Old != nil {
		return true
	}

	return false
}

// SetOld gets a reference to the given TemplateSummaryDiffLabelFields and assigns it to the Old field.
func (o *TemplateSummaryDiffLabel) SetOld(v TemplateSummaryDiffLabelFields) {
	o.Old = &v
}

func (o TemplateSummaryDiffLabel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StateStatus != nil {
		toSerialize["stateStatus"] = o.StateStatus
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.TemplateMetaName != nil {
		toSerialize["templateMetaName"] = o.TemplateMetaName
	}
	if o.New != nil {
		toSerialize["new"] = o.New
	}
	if o.Old != nil {
		toSerialize["old"] = o.Old
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryDiffLabel struct {
	value *TemplateSummaryDiffLabel
	isSet bool
}

func (v NullableTemplateSummaryDiffLabel) Get() *TemplateSummaryDiffLabel {
	return v.value
}

func (v *NullableTemplateSummaryDiffLabel) Set(val *TemplateSummaryDiffLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryDiffLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryDiffLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryDiffLabel(val *TemplateSummaryDiffLabel) *NullableTemplateSummaryDiffLabel {
	return &NullableTemplateSummaryDiffLabel{value: val, isSet: true}
}

func (v NullableTemplateSummaryDiffLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryDiffLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}