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
	"fmt"
)

// ColumnSemanticType the model 'ColumnSemanticType'
type ColumnSemanticType string

// List of ColumnSemanticType
const (
	COLUMNSEMANTICTYPE_TIMESTAMP ColumnSemanticType = "timestamp"
	COLUMNSEMANTICTYPE_TAG       ColumnSemanticType = "tag"
	COLUMNSEMANTICTYPE_FIELD     ColumnSemanticType = "field"
)

func ColumnSemanticTypeValues() []ColumnSemanticType {
	return []ColumnSemanticType{"timestamp", "tag", "field"}
}

func (v *ColumnSemanticType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ColumnSemanticType(value)
	for _, existing := range []ColumnSemanticType{"timestamp", "tag", "field"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ColumnSemanticType", value)
}

// Ptr returns reference to ColumnSemanticType value
func (v ColumnSemanticType) Ptr() *ColumnSemanticType {
	return &v
}

type NullableColumnSemanticType struct {
	value *ColumnSemanticType
	isSet bool
}

func (v NullableColumnSemanticType) Get() *ColumnSemanticType {
	return v.value
}

func (v *NullableColumnSemanticType) Set(val *ColumnSemanticType) {
	v.value = val
	v.isSet = true
}

func (v NullableColumnSemanticType) IsSet() bool {
	return v.isSet
}

func (v *NullableColumnSemanticType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableColumnSemanticType(val *ColumnSemanticType) *NullableColumnSemanticType {
	return &NullableColumnSemanticType{value: val, isSet: true}
}

func (v NullableColumnSemanticType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableColumnSemanticType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
