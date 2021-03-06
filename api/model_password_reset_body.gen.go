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

// PasswordResetBody struct for PasswordResetBody
type PasswordResetBody struct {
	Password string `json:"password" yaml:"password"`
}

// NewPasswordResetBody instantiates a new PasswordResetBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordResetBody(password string) *PasswordResetBody {
	this := PasswordResetBody{}
	this.Password = password
	return &this
}

// NewPasswordResetBodyWithDefaults instantiates a new PasswordResetBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordResetBodyWithDefaults() *PasswordResetBody {
	this := PasswordResetBody{}
	return &this
}

// GetPassword returns the Password field value
func (o *PasswordResetBody) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *PasswordResetBody) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *PasswordResetBody) SetPassword(v string) {
	o.Password = v
}

func (o PasswordResetBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullablePasswordResetBody struct {
	value *PasswordResetBody
	isSet bool
}

func (v NullablePasswordResetBody) Get() *PasswordResetBody {
	return v.value
}

func (v *NullablePasswordResetBody) Set(val *PasswordResetBody) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordResetBody) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordResetBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordResetBody(val *PasswordResetBody) *NullablePasswordResetBody {
	return &NullablePasswordResetBody{value: val, isSet: true}
}

func (v NullablePasswordResetBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordResetBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
