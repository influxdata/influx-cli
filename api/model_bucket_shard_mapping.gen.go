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

// BucketShardMapping struct for BucketShardMapping
type BucketShardMapping struct {
	OldId int64 `json:"oldId" yaml:"oldId"`
	NewId int64 `json:"newId" yaml:"newId"`
}

// NewBucketShardMapping instantiates a new BucketShardMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucketShardMapping(oldId int64, newId int64) *BucketShardMapping {
	this := BucketShardMapping{}
	this.OldId = oldId
	this.NewId = newId
	return &this
}

// NewBucketShardMappingWithDefaults instantiates a new BucketShardMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketShardMappingWithDefaults() *BucketShardMapping {
	this := BucketShardMapping{}
	return &this
}

// GetOldId returns the OldId field value
func (o *BucketShardMapping) GetOldId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OldId
}

// GetOldIdOk returns a tuple with the OldId field value
// and a boolean to check if the value has been set.
func (o *BucketShardMapping) GetOldIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OldId, true
}

// SetOldId sets field value
func (o *BucketShardMapping) SetOldId(v int64) {
	o.OldId = v
}

// GetNewId returns the NewId field value
func (o *BucketShardMapping) GetNewId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NewId
}

// GetNewIdOk returns a tuple with the NewId field value
// and a boolean to check if the value has been set.
func (o *BucketShardMapping) GetNewIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewId, true
}

// SetNewId sets field value
func (o *BucketShardMapping) SetNewId(v int64) {
	o.NewId = v
}

func (o BucketShardMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["oldId"] = o.OldId
	}
	if true {
		toSerialize["newId"] = o.NewId
	}
	return json.Marshal(toSerialize)
}

type NullableBucketShardMapping struct {
	value *BucketShardMapping
	isSet bool
}

func (v NullableBucketShardMapping) Get() *BucketShardMapping {
	return v.value
}

func (v *NullableBucketShardMapping) Set(val *BucketShardMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketShardMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketShardMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketShardMapping(val *BucketShardMapping) *NullableBucketShardMapping {
	return &NullableBucketShardMapping{value: val, isSet: true}
}

func (v NullableBucketShardMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketShardMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
