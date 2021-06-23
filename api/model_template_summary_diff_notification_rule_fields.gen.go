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

// TemplateSummaryDiffNotificationRuleFields struct for TemplateSummaryDiffNotificationRuleFields
type TemplateSummaryDiffNotificationRuleFields struct {
	Name            *string `json:"name,omitempty"`
	Description     *string `json:"description,omitempty"`
	EndpointName    *string `json:"endpointName,omitempty"`
	EndpointID      *string `json:"endpointID,omitempty"`
	EndpointType    *string `json:"endpointType,omitempty"`
	Every           *string `json:"every,omitempty"`
	Offset          *string `json:"offset,omitempty"`
	MessageTemplate *string `json:"messageTemplate,omitempty"`
}

// NewTemplateSummaryDiffNotificationRuleFields instantiates a new TemplateSummaryDiffNotificationRuleFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryDiffNotificationRuleFields() *TemplateSummaryDiffNotificationRuleFields {
	this := TemplateSummaryDiffNotificationRuleFields{}
	return &this
}

// NewTemplateSummaryDiffNotificationRuleFieldsWithDefaults instantiates a new TemplateSummaryDiffNotificationRuleFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryDiffNotificationRuleFieldsWithDefaults() *TemplateSummaryDiffNotificationRuleFields {
	this := TemplateSummaryDiffNotificationRuleFields{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplateSummaryDiffNotificationRuleFields) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffNotificationRuleFields) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplateSummaryDiffNotificationRuleFields) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplateSummaryDiffNotificationRuleFields) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TemplateSummaryDiffNotificationRuleFields) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffNotificationRuleFields) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateSummaryDiffNotificationRuleFields) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TemplateSummaryDiffNotificationRuleFields) SetDescription(v string) {
	o.Description = &v
}

// GetEndpointName returns the EndpointName field value if set, zero value otherwise.
func (o *TemplateSummaryDiffNotificationRuleFields) GetEndpointName() string {
	if o == nil || o.EndpointName == nil {
		var ret string
		return ret
	}
	return *o.EndpointName
}

// GetEndpointNameOk returns a tuple with the EndpointName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffNotificationRuleFields) GetEndpointNameOk() (*string, bool) {
	if o == nil || o.EndpointName == nil {
		return nil, false
	}
	return o.EndpointName, true
}

// HasEndpointName returns a boolean if a field has been set.
func (o *TemplateSummaryDiffNotificationRuleFields) HasEndpointName() bool {
	if o != nil && o.EndpointName != nil {
		return true
	}

	return false
}

// SetEndpointName gets a reference to the given string and assigns it to the EndpointName field.
func (o *TemplateSummaryDiffNotificationRuleFields) SetEndpointName(v string) {
	o.EndpointName = &v
}

// GetEndpointID returns the EndpointID field value if set, zero value otherwise.
func (o *TemplateSummaryDiffNotificationRuleFields) GetEndpointID() string {
	if o == nil || o.EndpointID == nil {
		var ret string
		return ret
	}
	return *o.EndpointID
}

// GetEndpointIDOk returns a tuple with the EndpointID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffNotificationRuleFields) GetEndpointIDOk() (*string, bool) {
	if o == nil || o.EndpointID == nil {
		return nil, false
	}
	return o.EndpointID, true
}

// HasEndpointID returns a boolean if a field has been set.
func (o *TemplateSummaryDiffNotificationRuleFields) HasEndpointID() bool {
	if o != nil && o.EndpointID != nil {
		return true
	}

	return false
}

// SetEndpointID gets a reference to the given string and assigns it to the EndpointID field.
func (o *TemplateSummaryDiffNotificationRuleFields) SetEndpointID(v string) {
	o.EndpointID = &v
}

// GetEndpointType returns the EndpointType field value if set, zero value otherwise.
func (o *TemplateSummaryDiffNotificationRuleFields) GetEndpointType() string {
	if o == nil || o.EndpointType == nil {
		var ret string
		return ret
	}
	return *o.EndpointType
}

// GetEndpointTypeOk returns a tuple with the EndpointType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffNotificationRuleFields) GetEndpointTypeOk() (*string, bool) {
	if o == nil || o.EndpointType == nil {
		return nil, false
	}
	return o.EndpointType, true
}

// HasEndpointType returns a boolean if a field has been set.
func (o *TemplateSummaryDiffNotificationRuleFields) HasEndpointType() bool {
	if o != nil && o.EndpointType != nil {
		return true
	}

	return false
}

// SetEndpointType gets a reference to the given string and assigns it to the EndpointType field.
func (o *TemplateSummaryDiffNotificationRuleFields) SetEndpointType(v string) {
	o.EndpointType = &v
}

// GetEvery returns the Every field value if set, zero value otherwise.
func (o *TemplateSummaryDiffNotificationRuleFields) GetEvery() string {
	if o == nil || o.Every == nil {
		var ret string
		return ret
	}
	return *o.Every
}

// GetEveryOk returns a tuple with the Every field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffNotificationRuleFields) GetEveryOk() (*string, bool) {
	if o == nil || o.Every == nil {
		return nil, false
	}
	return o.Every, true
}

// HasEvery returns a boolean if a field has been set.
func (o *TemplateSummaryDiffNotificationRuleFields) HasEvery() bool {
	if o != nil && o.Every != nil {
		return true
	}

	return false
}

// SetEvery gets a reference to the given string and assigns it to the Every field.
func (o *TemplateSummaryDiffNotificationRuleFields) SetEvery(v string) {
	o.Every = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *TemplateSummaryDiffNotificationRuleFields) GetOffset() string {
	if o == nil || o.Offset == nil {
		var ret string
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffNotificationRuleFields) GetOffsetOk() (*string, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *TemplateSummaryDiffNotificationRuleFields) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given string and assigns it to the Offset field.
func (o *TemplateSummaryDiffNotificationRuleFields) SetOffset(v string) {
	o.Offset = &v
}

// GetMessageTemplate returns the MessageTemplate field value if set, zero value otherwise.
func (o *TemplateSummaryDiffNotificationRuleFields) GetMessageTemplate() string {
	if o == nil || o.MessageTemplate == nil {
		var ret string
		return ret
	}
	return *o.MessageTemplate
}

// GetMessageTemplateOk returns a tuple with the MessageTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffNotificationRuleFields) GetMessageTemplateOk() (*string, bool) {
	if o == nil || o.MessageTemplate == nil {
		return nil, false
	}
	return o.MessageTemplate, true
}

// HasMessageTemplate returns a boolean if a field has been set.
func (o *TemplateSummaryDiffNotificationRuleFields) HasMessageTemplate() bool {
	if o != nil && o.MessageTemplate != nil {
		return true
	}

	return false
}

// SetMessageTemplate gets a reference to the given string and assigns it to the MessageTemplate field.
func (o *TemplateSummaryDiffNotificationRuleFields) SetMessageTemplate(v string) {
	o.MessageTemplate = &v
}

func (o TemplateSummaryDiffNotificationRuleFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EndpointName != nil {
		toSerialize["endpointName"] = o.EndpointName
	}
	if o.EndpointID != nil {
		toSerialize["endpointID"] = o.EndpointID
	}
	if o.EndpointType != nil {
		toSerialize["endpointType"] = o.EndpointType
	}
	if o.Every != nil {
		toSerialize["every"] = o.Every
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.MessageTemplate != nil {
		toSerialize["messageTemplate"] = o.MessageTemplate
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryDiffNotificationRuleFields struct {
	value *TemplateSummaryDiffNotificationRuleFields
	isSet bool
}

func (v NullableTemplateSummaryDiffNotificationRuleFields) Get() *TemplateSummaryDiffNotificationRuleFields {
	return v.value
}

func (v *NullableTemplateSummaryDiffNotificationRuleFields) Set(val *TemplateSummaryDiffNotificationRuleFields) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryDiffNotificationRuleFields) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryDiffNotificationRuleFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryDiffNotificationRuleFields(val *TemplateSummaryDiffNotificationRuleFields) *NullableTemplateSummaryDiffNotificationRuleFields {
	return &NullableTemplateSummaryDiffNotificationRuleFields{value: val, isSet: true}
}

func (v NullableTemplateSummaryDiffNotificationRuleFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryDiffNotificationRuleFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
