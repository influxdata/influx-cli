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

// MeasurementSchemaCreateRequest Create a new measurement schema
type MeasurementSchemaCreateRequest struct {
	Name string `json:"name" yaml:"name"`
	// An ordered collection of column definitions
	Columns []MeasurementSchemaColumn `json:"columns" yaml:"columns"`
}

// NewMeasurementSchemaCreateRequest instantiates a new MeasurementSchemaCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeasurementSchemaCreateRequest(name string, columns []MeasurementSchemaColumn) *MeasurementSchemaCreateRequest {
	this := MeasurementSchemaCreateRequest{}
	this.Name = name
	this.Columns = columns
	return &this
}

// NewMeasurementSchemaCreateRequestWithDefaults instantiates a new MeasurementSchemaCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeasurementSchemaCreateRequestWithDefaults() *MeasurementSchemaCreateRequest {
	this := MeasurementSchemaCreateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *MeasurementSchemaCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MeasurementSchemaCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MeasurementSchemaCreateRequest) SetName(v string) {
	o.Name = v
}

// GetColumns returns the Columns field value
func (o *MeasurementSchemaCreateRequest) GetColumns() []MeasurementSchemaColumn {
	if o == nil {
		var ret []MeasurementSchemaColumn
		return ret
	}

	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *MeasurementSchemaCreateRequest) GetColumnsOk() (*[]MeasurementSchemaColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value
func (o *MeasurementSchemaCreateRequest) SetColumns(v []MeasurementSchemaColumn) {
	o.Columns = v
}

func (o MeasurementSchemaCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["columns"] = o.Columns
	}
	return json.Marshal(toSerialize)
}

type NullableMeasurementSchemaCreateRequest struct {
	value *MeasurementSchemaCreateRequest
	isSet bool
}

func (v NullableMeasurementSchemaCreateRequest) Get() *MeasurementSchemaCreateRequest {
	return v.value
}

func (v *NullableMeasurementSchemaCreateRequest) Set(val *MeasurementSchemaCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurementSchemaCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurementSchemaCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurementSchemaCreateRequest(val *MeasurementSchemaCreateRequest) *NullableMeasurementSchemaCreateRequest {
	return &NullableMeasurementSchemaCreateRequest{value: val, isSet: true}
}

func (v NullableMeasurementSchemaCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasurementSchemaCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
