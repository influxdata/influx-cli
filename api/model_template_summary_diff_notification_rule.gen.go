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

// TemplateSummaryDiffNotificationRule struct for TemplateSummaryDiffNotificationRule
type TemplateSummaryDiffNotificationRule struct {
	Kind             *string                                    `json:"kind,omitempty"`
	StateStatus      *string                                    `json:"stateStatus,omitempty"`
	Id               *string                                    `json:"id,omitempty"`
	TemplateMetaName *string                                    `json:"templateMetaName,omitempty"`
	New              *TemplateSummaryDiffNotificationRuleFields `json:"new,omitempty"`
	Old              *TemplateSummaryDiffNotificationRuleFields `json:"old,omitempty"`
}

// NewTemplateSummaryDiffNotificationRule instantiates a new TemplateSummaryDiffNotificationRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryDiffNotificationRule() *TemplateSummaryDiffNotificationRule {
	this := TemplateSummaryDiffNotificationRule{}
	return &this
}

// NewTemplateSummaryDiffNotificationRuleWithDefaults instantiates a new TemplateSummaryDiffNotificationRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryDiffNotificationRuleWithDefaults() *TemplateSummaryDiffNotificationRule {
	this := TemplateSummaryDiffNotificationRule{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *TemplateSummaryDiffNotificationRule) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffNotificationRule) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *TemplateSummaryDiffNotificationRule) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *TemplateSummaryDiffNotificationRule) SetKind(v string) {
	o.Kind = &v
}

// GetStateStatus returns the StateStatus field value if set, zero value otherwise.
func (o *TemplateSummaryDiffNotificationRule) GetStateStatus() string {
	if o == nil || o.StateStatus == nil {
		var ret string
		return ret
	}
	return *o.StateStatus
}

// GetStateStatusOk returns a tuple with the StateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffNotificationRule) GetStateStatusOk() (*string, bool) {
	if o == nil || o.StateStatus == nil {
		return nil, false
	}
	return o.StateStatus, true
}

// HasStateStatus returns a boolean if a field has been set.
func (o *TemplateSummaryDiffNotificationRule) HasStateStatus() bool {
	if o != nil && o.StateStatus != nil {
		return true
	}

	return false
}

// SetStateStatus gets a reference to the given string and assigns it to the StateStatus field.
func (o *TemplateSummaryDiffNotificationRule) SetStateStatus(v string) {
	o.StateStatus = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TemplateSummaryDiffNotificationRule) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffNotificationRule) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TemplateSummaryDiffNotificationRule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TemplateSummaryDiffNotificationRule) SetId(v string) {
	o.Id = &v
}

// GetTemplateMetaName returns the TemplateMetaName field value if set, zero value otherwise.
func (o *TemplateSummaryDiffNotificationRule) GetTemplateMetaName() string {
	if o == nil || o.TemplateMetaName == nil {
		var ret string
		return ret
	}
	return *o.TemplateMetaName
}

// GetTemplateMetaNameOk returns a tuple with the TemplateMetaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffNotificationRule) GetTemplateMetaNameOk() (*string, bool) {
	if o == nil || o.TemplateMetaName == nil {
		return nil, false
	}
	return o.TemplateMetaName, true
}

// HasTemplateMetaName returns a boolean if a field has been set.
func (o *TemplateSummaryDiffNotificationRule) HasTemplateMetaName() bool {
	if o != nil && o.TemplateMetaName != nil {
		return true
	}

	return false
}

// SetTemplateMetaName gets a reference to the given string and assigns it to the TemplateMetaName field.
func (o *TemplateSummaryDiffNotificationRule) SetTemplateMetaName(v string) {
	o.TemplateMetaName = &v
}

// GetNew returns the New field value if set, zero value otherwise.
func (o *TemplateSummaryDiffNotificationRule) GetNew() TemplateSummaryDiffNotificationRuleFields {
	if o == nil || o.New == nil {
		var ret TemplateSummaryDiffNotificationRuleFields
		return ret
	}
	return *o.New
}

// GetNewOk returns a tuple with the New field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffNotificationRule) GetNewOk() (*TemplateSummaryDiffNotificationRuleFields, bool) {
	if o == nil || o.New == nil {
		return nil, false
	}
	return o.New, true
}

// HasNew returns a boolean if a field has been set.
func (o *TemplateSummaryDiffNotificationRule) HasNew() bool {
	if o != nil && o.New != nil {
		return true
	}

	return false
}

// SetNew gets a reference to the given TemplateSummaryDiffNotificationRuleFields and assigns it to the New field.
func (o *TemplateSummaryDiffNotificationRule) SetNew(v TemplateSummaryDiffNotificationRuleFields) {
	o.New = &v
}

// GetOld returns the Old field value if set, zero value otherwise.
func (o *TemplateSummaryDiffNotificationRule) GetOld() TemplateSummaryDiffNotificationRuleFields {
	if o == nil || o.Old == nil {
		var ret TemplateSummaryDiffNotificationRuleFields
		return ret
	}
	return *o.Old
}

// GetOldOk returns a tuple with the Old field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffNotificationRule) GetOldOk() (*TemplateSummaryDiffNotificationRuleFields, bool) {
	if o == nil || o.Old == nil {
		return nil, false
	}
	return o.Old, true
}

// HasOld returns a boolean if a field has been set.
func (o *TemplateSummaryDiffNotificationRule) HasOld() bool {
	if o != nil && o.Old != nil {
		return true
	}

	return false
}

// SetOld gets a reference to the given TemplateSummaryDiffNotificationRuleFields and assigns it to the Old field.
func (o *TemplateSummaryDiffNotificationRule) SetOld(v TemplateSummaryDiffNotificationRuleFields) {
	o.Old = &v
}

func (o TemplateSummaryDiffNotificationRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.StateStatus != nil {
		toSerialize["stateStatus"] = o.StateStatus
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

type NullableTemplateSummaryDiffNotificationRule struct {
	value *TemplateSummaryDiffNotificationRule
	isSet bool
}

func (v NullableTemplateSummaryDiffNotificationRule) Get() *TemplateSummaryDiffNotificationRule {
	return v.value
}

func (v *NullableTemplateSummaryDiffNotificationRule) Set(val *TemplateSummaryDiffNotificationRule) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryDiffNotificationRule) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryDiffNotificationRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryDiffNotificationRule(val *TemplateSummaryDiffNotificationRule) *NullableTemplateSummaryDiffNotificationRule {
	return &NullableTemplateSummaryDiffNotificationRule{value: val, isSet: true}
}

func (v NullableTemplateSummaryDiffNotificationRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryDiffNotificationRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
