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

// SecretKeysResponseAllOf struct for SecretKeysResponseAllOf
type SecretKeysResponseAllOf struct {
	Links *SecretKeysResponseAllOfLinks `json:"links,omitempty" yaml:"links,omitempty"`
}

// NewSecretKeysResponseAllOf instantiates a new SecretKeysResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretKeysResponseAllOf() *SecretKeysResponseAllOf {
	this := SecretKeysResponseAllOf{}
	return &this
}

// NewSecretKeysResponseAllOfWithDefaults instantiates a new SecretKeysResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretKeysResponseAllOfWithDefaults() *SecretKeysResponseAllOf {
	this := SecretKeysResponseAllOf{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SecretKeysResponseAllOf) GetLinks() SecretKeysResponseAllOfLinks {
	if o == nil || o.Links == nil {
		var ret SecretKeysResponseAllOfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretKeysResponseAllOf) GetLinksOk() (*SecretKeysResponseAllOfLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SecretKeysResponseAllOf) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SecretKeysResponseAllOfLinks and assigns it to the Links field.
func (o *SecretKeysResponseAllOf) SetLinks(v SecretKeysResponseAllOfLinks) {
	o.Links = &v
}

func (o SecretKeysResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableSecretKeysResponseAllOf struct {
	value *SecretKeysResponseAllOf
	isSet bool
}

func (v NullableSecretKeysResponseAllOf) Get() *SecretKeysResponseAllOf {
	return v.value
}

func (v *NullableSecretKeysResponseAllOf) Set(val *SecretKeysResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretKeysResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretKeysResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretKeysResponseAllOf(val *SecretKeysResponseAllOf) *NullableSecretKeysResponseAllOf {
	return &NullableSecretKeysResponseAllOf{value: val, isSet: true}
}

func (v NullableSecretKeysResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretKeysResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
