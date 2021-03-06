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

// Organizations struct for Organizations
type Organizations struct {
	Links *Links          `json:"links,omitempty" yaml:"links,omitempty"`
	Orgs  *[]Organization `json:"orgs,omitempty" yaml:"orgs,omitempty"`
}

// NewOrganizations instantiates a new Organizations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizations() *Organizations {
	this := Organizations{}
	return &this
}

// NewOrganizationsWithDefaults instantiates a new Organizations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsWithDefaults() *Organizations {
	this := Organizations{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Organizations) GetLinks() Links {
	if o == nil || o.Links == nil {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizations) GetLinksOk() (*Links, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Organizations) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *Organizations) SetLinks(v Links) {
	o.Links = &v
}

// GetOrgs returns the Orgs field value if set, zero value otherwise.
func (o *Organizations) GetOrgs() []Organization {
	if o == nil || o.Orgs == nil {
		var ret []Organization
		return ret
	}
	return *o.Orgs
}

// GetOrgsOk returns a tuple with the Orgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizations) GetOrgsOk() (*[]Organization, bool) {
	if o == nil || o.Orgs == nil {
		return nil, false
	}
	return o.Orgs, true
}

// HasOrgs returns a boolean if a field has been set.
func (o *Organizations) HasOrgs() bool {
	if o != nil && o.Orgs != nil {
		return true
	}

	return false
}

// SetOrgs gets a reference to the given []Organization and assigns it to the Orgs field.
func (o *Organizations) SetOrgs(v []Organization) {
	o.Orgs = &v
}

func (o Organizations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Orgs != nil {
		toSerialize["orgs"] = o.Orgs
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizations struct {
	value *Organizations
	isSet bool
}

func (v NullableOrganizations) Get() *Organizations {
	return v.value
}

func (v *NullableOrganizations) Set(val *Organizations) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizations) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizations(val *Organizations) *NullableOrganizations {
	return &NullableOrganizations{value: val, isSet: true}
}

func (v NullableOrganizations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
