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

// SchemaType the model 'SchemaType'
type SchemaType string

// List of SchemaType
const (
	SCHEMATYPE_IMPLICIT SchemaType = "implicit"
	SCHEMATYPE_EXPLICIT SchemaType = "explicit"
)

func SchemaTypeValues() []SchemaType {
	return []SchemaType{"implicit", "explicit"}
}

func (v *SchemaType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SchemaType(value)
	for _, existing := range []SchemaType{"implicit", "explicit"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SchemaType", value)
}

// Ptr returns reference to SchemaType value
func (v SchemaType) Ptr() *SchemaType {
	return &v
}

type NullableSchemaType struct {
	value *SchemaType
	isSet bool
}

func (v NullableSchemaType) Get() *SchemaType {
	return v.value
}

func (v *NullableSchemaType) Set(val *SchemaType) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaType) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaType(val *SchemaType) *NullableSchemaType {
	return &NullableSchemaType{value: val, isSet: true}
}

func (v NullableSchemaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
