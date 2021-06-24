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

// Telegraf struct for Telegraf
type Telegraf struct {
	Name        *string                  `json:"name,omitempty" yaml:"name,omitempty"`
	Description *string                  `json:"description,omitempty" yaml:"description,omitempty"`
	Metadata    *TelegrafRequestMetadata `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Config      *string                  `json:"config,omitempty" yaml:"config,omitempty"`
	OrgID       *string                  `json:"orgID,omitempty" yaml:"orgID,omitempty"`
	Id          *string                  `json:"id,omitempty" yaml:"id,omitempty"`
	Links       *TelegrafAllOfLinks      `json:"links,omitempty" yaml:"links,omitempty"`
	Labels      *[]Label                 `json:"labels,omitempty" yaml:"labels,omitempty"`
}

// NewTelegraf instantiates a new Telegraf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelegraf() *Telegraf {
	this := Telegraf{}
	return &this
}

// NewTelegrafWithDefaults instantiates a new Telegraf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelegrafWithDefaults() *Telegraf {
	this := Telegraf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Telegraf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Telegraf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Telegraf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Telegraf) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Telegraf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Telegraf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Telegraf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Telegraf) SetDescription(v string) {
	o.Description = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Telegraf) GetMetadata() TelegrafRequestMetadata {
	if o == nil || o.Metadata == nil {
		var ret TelegrafRequestMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Telegraf) GetMetadataOk() (*TelegrafRequestMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Telegraf) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given TelegrafRequestMetadata and assigns it to the Metadata field.
func (o *Telegraf) SetMetadata(v TelegrafRequestMetadata) {
	o.Metadata = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Telegraf) GetConfig() string {
	if o == nil || o.Config == nil {
		var ret string
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Telegraf) GetConfigOk() (*string, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Telegraf) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given string and assigns it to the Config field.
func (o *Telegraf) SetConfig(v string) {
	o.Config = &v
}

// GetOrgID returns the OrgID field value if set, zero value otherwise.
func (o *Telegraf) GetOrgID() string {
	if o == nil || o.OrgID == nil {
		var ret string
		return ret
	}
	return *o.OrgID
}

// GetOrgIDOk returns a tuple with the OrgID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Telegraf) GetOrgIDOk() (*string, bool) {
	if o == nil || o.OrgID == nil {
		return nil, false
	}
	return o.OrgID, true
}

// HasOrgID returns a boolean if a field has been set.
func (o *Telegraf) HasOrgID() bool {
	if o != nil && o.OrgID != nil {
		return true
	}

	return false
}

// SetOrgID gets a reference to the given string and assigns it to the OrgID field.
func (o *Telegraf) SetOrgID(v string) {
	o.OrgID = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Telegraf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Telegraf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Telegraf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Telegraf) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Telegraf) GetLinks() TelegrafAllOfLinks {
	if o == nil || o.Links == nil {
		var ret TelegrafAllOfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Telegraf) GetLinksOk() (*TelegrafAllOfLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Telegraf) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TelegrafAllOfLinks and assigns it to the Links field.
func (o *Telegraf) SetLinks(v TelegrafAllOfLinks) {
	o.Links = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *Telegraf) GetLabels() []Label {
	if o == nil || o.Labels == nil {
		var ret []Label
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Telegraf) GetLabelsOk() (*[]Label, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *Telegraf) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []Label and assigns it to the Labels field.
func (o *Telegraf) SetLabels(v []Label) {
	o.Labels = &v
}

func (o Telegraf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.OrgID != nil {
		toSerialize["orgID"] = o.OrgID
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullableTelegraf struct {
	value *Telegraf
	isSet bool
}

func (v NullableTelegraf) Get() *Telegraf {
	return v.value
}

func (v *NullableTelegraf) Set(val *Telegraf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelegraf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelegraf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelegraf(val *Telegraf) *NullableTelegraf {
	return &NullableTelegraf{value: val, isSet: true}
}

func (v NullableTelegraf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelegraf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
