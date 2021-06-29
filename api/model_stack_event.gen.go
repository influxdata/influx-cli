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
	"time"
)

// StackEvent struct for StackEvent
type StackEvent struct {
	EventType   string               `json:"eventType" yaml:"eventType"`
	Name        string               `json:"name" yaml:"name"`
	Description *string              `json:"description,omitempty" yaml:"description,omitempty"`
	Sources     []string             `json:"sources" yaml:"sources"`
	Resources   []StackEventResource `json:"resources" yaml:"resources"`
	Urls        []string             `json:"urls" yaml:"urls"`
	UpdatedAt   time.Time            `json:"updatedAt" yaml:"updatedAt"`
}

// NewStackEvent instantiates a new StackEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStackEvent(eventType string, name string, sources []string, resources []StackEventResource, urls []string, updatedAt time.Time) *StackEvent {
	this := StackEvent{}
	this.EventType = eventType
	this.Name = name
	this.Sources = sources
	this.Resources = resources
	this.Urls = urls
	this.UpdatedAt = updatedAt
	return &this
}

// NewStackEventWithDefaults instantiates a new StackEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStackEventWithDefaults() *StackEvent {
	this := StackEvent{}
	return &this
}

// GetEventType returns the EventType field value
func (o *StackEvent) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *StackEvent) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *StackEvent) SetEventType(v string) {
	o.EventType = v
}

// GetName returns the Name field value
func (o *StackEvent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StackEvent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StackEvent) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StackEvent) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackEvent) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StackEvent) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StackEvent) SetDescription(v string) {
	o.Description = &v
}

// GetSources returns the Sources field value
func (o *StackEvent) GetSources() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *StackEvent) GetSourcesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value
func (o *StackEvent) SetSources(v []string) {
	o.Sources = v
}

// GetResources returns the Resources field value
func (o *StackEvent) GetResources() []StackEventResource {
	if o == nil {
		var ret []StackEventResource
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *StackEvent) GetResourcesOk() (*[]StackEventResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *StackEvent) SetResources(v []StackEventResource) {
	o.Resources = v
}

// GetUrls returns the Urls field value
func (o *StackEvent) GetUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *StackEvent) GetUrlsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *StackEvent) SetUrls(v []string) {
	o.Urls = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *StackEvent) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *StackEvent) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *StackEvent) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o StackEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["sources"] = o.Sources
	}
	if true {
		toSerialize["resources"] = o.Resources
	}
	if true {
		toSerialize["urls"] = o.Urls
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableStackEvent struct {
	value *StackEvent
	isSet bool
}

func (v NullableStackEvent) Get() *StackEvent {
	return v.value
}

func (v *NullableStackEvent) Set(val *StackEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableStackEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableStackEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStackEvent(val *StackEvent) *NullableStackEvent {
	return &NullableStackEvent{value: val, isSet: true}
}

func (v NullableStackEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStackEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
