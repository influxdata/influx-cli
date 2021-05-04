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
	"errors"
)

// PropertyKey - struct for PropertyKey
type PropertyKey struct {
	Identifier    *Identifier
	StringLiteral *StringLiteral
}

// IdentifierAsPropertyKey is a convenience function that returns Identifier wrapped in PropertyKey
func IdentifierAsPropertyKey(v *Identifier) PropertyKey {
	return PropertyKey{Identifier: v}
}

// StringLiteralAsPropertyKey is a convenience function that returns StringLiteral wrapped in PropertyKey
func StringLiteralAsPropertyKey(v *StringLiteral) PropertyKey {
	return PropertyKey{StringLiteral: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PropertyKey) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Identifier
	err = json.Unmarshal(data, &dst.Identifier)
	if err == nil {
		jsonIdentifier, _ := json.Marshal(dst.Identifier)
		if string(jsonIdentifier) == "{}" { // empty struct
			dst.Identifier = nil
		} else {
			match++
		}
	} else {
		dst.Identifier = nil
	}

	// try to unmarshal data into StringLiteral
	err = json.Unmarshal(data, &dst.StringLiteral)
	if err == nil {
		jsonStringLiteral, _ := json.Marshal(dst.StringLiteral)
		if string(jsonStringLiteral) == "{}" { // empty struct
			dst.StringLiteral = nil
		} else {
			match++
		}
	} else {
		dst.StringLiteral = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Identifier = nil
		dst.StringLiteral = nil

		return errors.New("data matches more than one schema in oneOf(PropertyKey)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return errors.New("data failed to match schemas in oneOf(PropertyKey)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PropertyKey) MarshalJSON() ([]byte, error) {
	if src.Identifier != nil {
		return json.Marshal(&src.Identifier)
	}

	if src.StringLiteral != nil {
		return json.Marshal(&src.StringLiteral)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PropertyKey) GetActualInstance() interface{} {
	if obj.Identifier != nil {
		return obj.Identifier
	}

	if obj.StringLiteral != nil {
		return obj.StringLiteral
	}

	// all schemas are nil
	return nil
}

type NullablePropertyKey struct {
	value *PropertyKey
	isSet bool
}

func (v NullablePropertyKey) Get() *PropertyKey {
	return v.value
}

func (v *NullablePropertyKey) Set(val *PropertyKey) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyKey) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyKey(val *PropertyKey) *NullablePropertyKey {
	return &NullablePropertyKey{value: val, isSet: true}
}

func (v NullablePropertyKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
