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
	"time"
)

// DeletePredicateRequest The delete predicate request.
type DeletePredicateRequest struct {
	// A timestamp ([RFC3339 date/time format](https://docs.influxdata.com/flux/v0.x/data-types/basic/time/#time-syntax)).
	Start time.Time `json:"start" yaml:"start"`
	// A timestamp ([RFC3339 date/time format](https://docs.influxdata.com/flux/v0.x/data-types/basic/time/#time-syntax)).
	Stop time.Time `json:"stop" yaml:"stop"`
	// An expression in [delete predicate syntax]({{% INFLUXDB_DOCS_URL %}}/reference/syntax/delete-predicate/).
	Predicate *string `json:"predicate,omitempty" yaml:"predicate,omitempty"`
}

// NewDeletePredicateRequest instantiates a new DeletePredicateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeletePredicateRequest(start time.Time, stop time.Time) *DeletePredicateRequest {
	this := DeletePredicateRequest{}
	this.Start = start
	this.Stop = stop
	return &this
}

// NewDeletePredicateRequestWithDefaults instantiates a new DeletePredicateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeletePredicateRequestWithDefaults() *DeletePredicateRequest {
	this := DeletePredicateRequest{}
	return &this
}

// GetStart returns the Start field value
func (o *DeletePredicateRequest) GetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *DeletePredicateRequest) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *DeletePredicateRequest) SetStart(v time.Time) {
	o.Start = v
}

// GetStop returns the Stop field value
func (o *DeletePredicateRequest) GetStop() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Stop
}

// GetStopOk returns a tuple with the Stop field value
// and a boolean to check if the value has been set.
func (o *DeletePredicateRequest) GetStopOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stop, true
}

// SetStop sets field value
func (o *DeletePredicateRequest) SetStop(v time.Time) {
	o.Stop = v
}

// GetPredicate returns the Predicate field value if set, zero value otherwise.
func (o *DeletePredicateRequest) GetPredicate() string {
	if o == nil || o.Predicate == nil {
		var ret string
		return ret
	}
	return *o.Predicate
}

// GetPredicateOk returns a tuple with the Predicate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletePredicateRequest) GetPredicateOk() (*string, bool) {
	if o == nil || o.Predicate == nil {
		return nil, false
	}
	return o.Predicate, true
}

// HasPredicate returns a boolean if a field has been set.
func (o *DeletePredicateRequest) HasPredicate() bool {
	if o != nil && o.Predicate != nil {
		return true
	}

	return false
}

// SetPredicate gets a reference to the given string and assigns it to the Predicate field.
func (o *DeletePredicateRequest) SetPredicate(v string) {
	o.Predicate = &v
}

func (o DeletePredicateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["start"] = o.Start
	}
	if true {
		toSerialize["stop"] = o.Stop
	}
	if o.Predicate != nil {
		toSerialize["predicate"] = o.Predicate
	}
	return json.Marshal(toSerialize)
}

type NullableDeletePredicateRequest struct {
	value *DeletePredicateRequest
	isSet bool
}

func (v NullableDeletePredicateRequest) Get() *DeletePredicateRequest {
	return v.value
}

func (v *NullableDeletePredicateRequest) Set(val *DeletePredicateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeletePredicateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeletePredicateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeletePredicateRequest(val *DeletePredicateRequest) *NullableDeletePredicateRequest {
	return &NullableDeletePredicateRequest{value: val, isSet: true}
}

func (v NullableDeletePredicateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeletePredicateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
