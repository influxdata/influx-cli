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

// TemplateSummaryDiffTaskFields struct for TemplateSummaryDiffTaskFields
type TemplateSummaryDiffTaskFields struct {
	Name        string  `json:"name" yaml:"name"`
	Cron        *string `json:"cron,omitempty" yaml:"cron,omitempty"`
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	Every       *string `json:"every,omitempty" yaml:"every,omitempty"`
	Offset      *string `json:"offset,omitempty" yaml:"offset,omitempty"`
	Query       *string `json:"query,omitempty" yaml:"query,omitempty"`
	Status      string  `json:"status" yaml:"status"`
}

// NewTemplateSummaryDiffTaskFields instantiates a new TemplateSummaryDiffTaskFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryDiffTaskFields(name string, status string) *TemplateSummaryDiffTaskFields {
	this := TemplateSummaryDiffTaskFields{}
	this.Name = name
	this.Status = status
	return &this
}

// NewTemplateSummaryDiffTaskFieldsWithDefaults instantiates a new TemplateSummaryDiffTaskFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryDiffTaskFieldsWithDefaults() *TemplateSummaryDiffTaskFields {
	this := TemplateSummaryDiffTaskFields{}
	return &this
}

// GetName returns the Name field value
func (o *TemplateSummaryDiffTaskFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffTaskFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TemplateSummaryDiffTaskFields) SetName(v string) {
	o.Name = v
}

// GetCron returns the Cron field value if set, zero value otherwise.
func (o *TemplateSummaryDiffTaskFields) GetCron() string {
	if o == nil || o.Cron == nil {
		var ret string
		return ret
	}
	return *o.Cron
}

// GetCronOk returns a tuple with the Cron field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffTaskFields) GetCronOk() (*string, bool) {
	if o == nil || o.Cron == nil {
		return nil, false
	}
	return o.Cron, true
}

// HasCron returns a boolean if a field has been set.
func (o *TemplateSummaryDiffTaskFields) HasCron() bool {
	if o != nil && o.Cron != nil {
		return true
	}

	return false
}

// SetCron gets a reference to the given string and assigns it to the Cron field.
func (o *TemplateSummaryDiffTaskFields) SetCron(v string) {
	o.Cron = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TemplateSummaryDiffTaskFields) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffTaskFields) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateSummaryDiffTaskFields) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TemplateSummaryDiffTaskFields) SetDescription(v string) {
	o.Description = &v
}

// GetEvery returns the Every field value if set, zero value otherwise.
func (o *TemplateSummaryDiffTaskFields) GetEvery() string {
	if o == nil || o.Every == nil {
		var ret string
		return ret
	}
	return *o.Every
}

// GetEveryOk returns a tuple with the Every field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffTaskFields) GetEveryOk() (*string, bool) {
	if o == nil || o.Every == nil {
		return nil, false
	}
	return o.Every, true
}

// HasEvery returns a boolean if a field has been set.
func (o *TemplateSummaryDiffTaskFields) HasEvery() bool {
	if o != nil && o.Every != nil {
		return true
	}

	return false
}

// SetEvery gets a reference to the given string and assigns it to the Every field.
func (o *TemplateSummaryDiffTaskFields) SetEvery(v string) {
	o.Every = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *TemplateSummaryDiffTaskFields) GetOffset() string {
	if o == nil || o.Offset == nil {
		var ret string
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffTaskFields) GetOffsetOk() (*string, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *TemplateSummaryDiffTaskFields) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given string and assigns it to the Offset field.
func (o *TemplateSummaryDiffTaskFields) SetOffset(v string) {
	o.Offset = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *TemplateSummaryDiffTaskFields) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffTaskFields) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *TemplateSummaryDiffTaskFields) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *TemplateSummaryDiffTaskFields) SetQuery(v string) {
	o.Query = &v
}

// GetStatus returns the Status field value
func (o *TemplateSummaryDiffTaskFields) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiffTaskFields) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TemplateSummaryDiffTaskFields) SetStatus(v string) {
	o.Status = v
}

func (o TemplateSummaryDiffTaskFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Cron != nil {
		toSerialize["cron"] = o.Cron
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Every != nil {
		toSerialize["every"] = o.Every
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryDiffTaskFields struct {
	value *TemplateSummaryDiffTaskFields
	isSet bool
}

func (v NullableTemplateSummaryDiffTaskFields) Get() *TemplateSummaryDiffTaskFields {
	return v.value
}

func (v *NullableTemplateSummaryDiffTaskFields) Set(val *TemplateSummaryDiffTaskFields) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryDiffTaskFields) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryDiffTaskFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryDiffTaskFields(val *TemplateSummaryDiffTaskFields) *NullableTemplateSummaryDiffTaskFields {
	return &NullableTemplateSummaryDiffTaskFields{value: val, isSet: true}
}

func (v NullableTemplateSummaryDiffTaskFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryDiffTaskFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
