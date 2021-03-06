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

// MeasurementSchemaList A list of measurement schemas returning summary information
type MeasurementSchemaList struct {
	MeasurementSchemas []MeasurementSchema `json:"measurementSchemas" yaml:"measurementSchemas"`
}

// NewMeasurementSchemaList instantiates a new MeasurementSchemaList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeasurementSchemaList(measurementSchemas []MeasurementSchema) *MeasurementSchemaList {
	this := MeasurementSchemaList{}
	this.MeasurementSchemas = measurementSchemas
	return &this
}

// NewMeasurementSchemaListWithDefaults instantiates a new MeasurementSchemaList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeasurementSchemaListWithDefaults() *MeasurementSchemaList {
	this := MeasurementSchemaList{}
	return &this
}

// GetMeasurementSchemas returns the MeasurementSchemas field value
func (o *MeasurementSchemaList) GetMeasurementSchemas() []MeasurementSchema {
	if o == nil {
		var ret []MeasurementSchema
		return ret
	}

	return o.MeasurementSchemas
}

// GetMeasurementSchemasOk returns a tuple with the MeasurementSchemas field value
// and a boolean to check if the value has been set.
func (o *MeasurementSchemaList) GetMeasurementSchemasOk() (*[]MeasurementSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeasurementSchemas, true
}

// SetMeasurementSchemas sets field value
func (o *MeasurementSchemaList) SetMeasurementSchemas(v []MeasurementSchema) {
	o.MeasurementSchemas = v
}

func (o MeasurementSchemaList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["measurementSchemas"] = o.MeasurementSchemas
	}
	return json.Marshal(toSerialize)
}

type NullableMeasurementSchemaList struct {
	value *MeasurementSchemaList
	isSet bool
}

func (v NullableMeasurementSchemaList) Get() *MeasurementSchemaList {
	return v.value
}

func (v *NullableMeasurementSchemaList) Set(val *MeasurementSchemaList) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurementSchemaList) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurementSchemaList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurementSchemaList(val *MeasurementSchemaList) *NullableMeasurementSchemaList {
	return &NullableMeasurementSchemaList{value: val, isSet: true}
}

func (v NullableMeasurementSchemaList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasurementSchemaList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
