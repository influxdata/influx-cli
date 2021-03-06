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

// TemplateSummaryDiffLabelMapping struct for TemplateSummaryDiffLabelMapping
type TemplateSummaryDiffLabelMapping struct {
	Status                   string `json:"status" yaml:"status"`
	ResourceTemplateMetaName string `json:"resourceTemplateMetaName" yaml:"resourceTemplateMetaName"`
	ResourceName             string `json:"resourceName" yaml:"resourceName"`
	ResourceID               uint64 `json:"resourceID" yaml:"resourceID"`
	ResourceType             string `json:"resourceType" yaml:"resourceType"`
	LabelTemplateMetaName    string `json:"labelTemplateMetaName" yaml:"labelTemplateMetaName"`
	LabelName                string `json:"labelName" yaml:"labelName"`
	LabelID                  uint64 `json:"labelID" yaml:"labelID"`
	StateStatus              string `json:"stateStatus" yaml:"stateStatus"`
}

// NewTemplateSummaryDiffLabelMapping instantiates a new TemplateSummaryDiffLabelMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryDiffLabelMapping(status string, resourceTemplateMetaName string, resourceName string, resourceID uint64, resourceType string, labelTemplateMetaName string, labelName string, labelID uint64, stateStatus string) *TemplateSummaryDiffLabelMapping {
	this := TemplateSummaryDiffLabelMapping{}
	this.Status = status
	this.ResourceTemplateMetaName = resourceTemplateMetaName
	this.ResourceName = resourceName
	this.ResourceID = resourceID
	this.ResourceType = resourceType
	this.LabelTemplateMetaName = labelTemplateMetaName
	this.LabelName = labelName
	this.LabelID = labelID
	this.StateStatus = stateStatus
	return &this
}

// NewTemplateSummaryDiffLabelMappingWithDefaults instantiates a new TemplateSummaryDiffLabelMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryDiffLabelMappingWithDefaults() *TemplateSummaryDiffLabelMapping {
	this := TemplateSummaryDiffLabelMapping{}
	return &this
}

// GetStatus returns the Status field value
func (o *TemplateSummaryDiffLabelMapping) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffLabelMapping) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TemplateSummaryDiffLabelMapping) SetStatus(v string) {
	o.Status = v
}

// GetResourceTemplateMetaName returns the ResourceTemplateMetaName field value
func (o *TemplateSummaryDiffLabelMapping) GetResourceTemplateMetaName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceTemplateMetaName
}

// GetResourceTemplateMetaNameOk returns a tuple with the ResourceTemplateMetaName field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffLabelMapping) GetResourceTemplateMetaNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceTemplateMetaName, true
}

// SetResourceTemplateMetaName sets field value
func (o *TemplateSummaryDiffLabelMapping) SetResourceTemplateMetaName(v string) {
	o.ResourceTemplateMetaName = v
}

// GetResourceName returns the ResourceName field value
func (o *TemplateSummaryDiffLabelMapping) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffLabelMapping) GetResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value
func (o *TemplateSummaryDiffLabelMapping) SetResourceName(v string) {
	o.ResourceName = v
}

// GetResourceID returns the ResourceID field value
func (o *TemplateSummaryDiffLabelMapping) GetResourceID() uint64 {
	if o == nil {
		var ret uint64
		return ret
	}

	return o.ResourceID
}

// GetResourceIDOk returns a tuple with the ResourceID field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffLabelMapping) GetResourceIDOk() (*uint64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceID, true
}

// SetResourceID sets field value
func (o *TemplateSummaryDiffLabelMapping) SetResourceID(v uint64) {
	o.ResourceID = v
}

// GetResourceType returns the ResourceType field value
func (o *TemplateSummaryDiffLabelMapping) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffLabelMapping) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *TemplateSummaryDiffLabelMapping) SetResourceType(v string) {
	o.ResourceType = v
}

// GetLabelTemplateMetaName returns the LabelTemplateMetaName field value
func (o *TemplateSummaryDiffLabelMapping) GetLabelTemplateMetaName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LabelTemplateMetaName
}

// GetLabelTemplateMetaNameOk returns a tuple with the LabelTemplateMetaName field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffLabelMapping) GetLabelTemplateMetaNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelTemplateMetaName, true
}

// SetLabelTemplateMetaName sets field value
func (o *TemplateSummaryDiffLabelMapping) SetLabelTemplateMetaName(v string) {
	o.LabelTemplateMetaName = v
}

// GetLabelName returns the LabelName field value
func (o *TemplateSummaryDiffLabelMapping) GetLabelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LabelName
}

// GetLabelNameOk returns a tuple with the LabelName field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffLabelMapping) GetLabelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelName, true
}

// SetLabelName sets field value
func (o *TemplateSummaryDiffLabelMapping) SetLabelName(v string) {
	o.LabelName = v
}

// GetLabelID returns the LabelID field value
func (o *TemplateSummaryDiffLabelMapping) GetLabelID() uint64 {
	if o == nil {
		var ret uint64
		return ret
	}

	return o.LabelID
}

// GetLabelIDOk returns a tuple with the LabelID field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffLabelMapping) GetLabelIDOk() (*uint64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelID, true
}

// SetLabelID sets field value
func (o *TemplateSummaryDiffLabelMapping) SetLabelID(v uint64) {
	o.LabelID = v
}

// GetStateStatus returns the StateStatus field value
func (o *TemplateSummaryDiffLabelMapping) GetStateStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateStatus
}

// GetStateStatusOk returns a tuple with the StateStatus field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffLabelMapping) GetStateStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateStatus, true
}

// SetStateStatus sets field value
func (o *TemplateSummaryDiffLabelMapping) SetStateStatus(v string) {
	o.StateStatus = v
}

func (o TemplateSummaryDiffLabelMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["resourceTemplateMetaName"] = o.ResourceTemplateMetaName
	}
	if true {
		toSerialize["resourceName"] = o.ResourceName
	}
	if true {
		toSerialize["resourceID"] = o.ResourceID
	}
	if true {
		toSerialize["resourceType"] = o.ResourceType
	}
	if true {
		toSerialize["labelTemplateMetaName"] = o.LabelTemplateMetaName
	}
	if true {
		toSerialize["labelName"] = o.LabelName
	}
	if true {
		toSerialize["labelID"] = o.LabelID
	}
	if true {
		toSerialize["stateStatus"] = o.StateStatus
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryDiffLabelMapping struct {
	value *TemplateSummaryDiffLabelMapping
	isSet bool
}

func (v NullableTemplateSummaryDiffLabelMapping) Get() *TemplateSummaryDiffLabelMapping {
	return v.value
}

func (v *NullableTemplateSummaryDiffLabelMapping) Set(val *TemplateSummaryDiffLabelMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryDiffLabelMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryDiffLabelMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryDiffLabelMapping(val *TemplateSummaryDiffLabelMapping) *NullableTemplateSummaryDiffLabelMapping {
	return &NullableTemplateSummaryDiffLabelMapping{value: val, isSet: true}
}

func (v NullableTemplateSummaryDiffLabelMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryDiffLabelMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
