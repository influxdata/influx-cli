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

// Links URI pointers for additional paged results.
type Links struct {
	// URI of resource.
	Next *string `json:"next,omitempty" yaml:"next,omitempty"`
	// URI of resource.
	Self string `json:"self" yaml:"self"`
	// URI of resource.
	Prev *string `json:"prev,omitempty" yaml:"prev,omitempty"`
}

// NewLinks instantiates a new Links object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinks(self string) *Links {
	this := Links{}
	this.Self = self
	return &this
}

// NewLinksWithDefaults instantiates a new Links object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksWithDefaults() *Links {
	this := Links{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *Links) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *Links) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *Links) SetNext(v string) {
	o.Next = &v
}

// GetSelf returns the Self field value
func (o *Links) GetSelf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *Links) GetSelfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *Links) SetSelf(v string) {
	o.Self = v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *Links) GetPrev() string {
	if o == nil || o.Prev == nil {
		var ret string
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetPrevOk() (*string, bool) {
	if o == nil || o.Prev == nil {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *Links) HasPrev() bool {
	if o != nil && o.Prev != nil {
		return true
	}

	return false
}

// SetPrev gets a reference to the given string and assigns it to the Prev field.
func (o *Links) SetPrev(v string) {
	o.Prev = &v
}

func (o Links) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if true {
		toSerialize["self"] = o.Self
	}
	if o.Prev != nil {
		toSerialize["prev"] = o.Prev
	}
	return json.Marshal(toSerialize)
}

type NullableLinks struct {
	value *Links
	isSet bool
}

func (v NullableLinks) Get() *Links {
	return v.value
}

func (v *NullableLinks) Set(val *Links) {
	v.value = val
	v.isSet = true
}

func (v NullableLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinks(val *Links) *NullableLinks {
	return &NullableLinks{value: val, isSet: true}
}

func (v NullableLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
