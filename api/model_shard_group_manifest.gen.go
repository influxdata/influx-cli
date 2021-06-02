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

// ShardGroupManifest struct for ShardGroupManifest
type ShardGroupManifest struct {
	Id          int64           `json:"id"`
	StartTime   time.Time       `json:"startTime"`
	EndTime     time.Time       `json:"endTime"`
	DeletedAt   *time.Time      `json:"deletedAt,omitempty"`
	TruncatedAt *time.Time      `json:"truncatedAt,omitempty"`
	Shards      []ShardManifest `json:"shards"`
}

// NewShardGroupManifest instantiates a new ShardGroupManifest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardGroupManifest(id int64, startTime time.Time, endTime time.Time, shards []ShardManifest) *ShardGroupManifest {
	this := ShardGroupManifest{}
	this.Id = id
	this.StartTime = startTime
	this.EndTime = endTime
	this.Shards = shards
	return &this
}

// NewShardGroupManifestWithDefaults instantiates a new ShardGroupManifest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardGroupManifestWithDefaults() *ShardGroupManifest {
	this := ShardGroupManifest{}
	return &this
}

// GetId returns the Id field value
func (o *ShardGroupManifest) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ShardGroupManifest) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ShardGroupManifest) SetId(v int64) {
	o.Id = v
}

// GetStartTime returns the StartTime field value
func (o *ShardGroupManifest) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *ShardGroupManifest) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *ShardGroupManifest) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *ShardGroupManifest) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *ShardGroupManifest) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *ShardGroupManifest) SetEndTime(v time.Time) {
	o.EndTime = v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ShardGroupManifest) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardGroupManifest) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ShardGroupManifest) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *ShardGroupManifest) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetTruncatedAt returns the TruncatedAt field value if set, zero value otherwise.
func (o *ShardGroupManifest) GetTruncatedAt() time.Time {
	if o == nil || o.TruncatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.TruncatedAt
}

// GetTruncatedAtOk returns a tuple with the TruncatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardGroupManifest) GetTruncatedAtOk() (*time.Time, bool) {
	if o == nil || o.TruncatedAt == nil {
		return nil, false
	}
	return o.TruncatedAt, true
}

// HasTruncatedAt returns a boolean if a field has been set.
func (o *ShardGroupManifest) HasTruncatedAt() bool {
	if o != nil && o.TruncatedAt != nil {
		return true
	}

	return false
}

// SetTruncatedAt gets a reference to the given time.Time and assigns it to the TruncatedAt field.
func (o *ShardGroupManifest) SetTruncatedAt(v time.Time) {
	o.TruncatedAt = &v
}

// GetShards returns the Shards field value
func (o *ShardGroupManifest) GetShards() []ShardManifest {
	if o == nil {
		var ret []ShardManifest
		return ret
	}

	return o.Shards
}

// GetShardsOk returns a tuple with the Shards field value
// and a boolean to check if the value has been set.
func (o *ShardGroupManifest) GetShardsOk() (*[]ShardManifest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Shards, true
}

// SetShards sets field value
func (o *ShardGroupManifest) SetShards(v []ShardManifest) {
	o.Shards = v
}

func (o ShardGroupManifest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["startTime"] = o.StartTime
	}
	if true {
		toSerialize["endTime"] = o.EndTime
	}
	if o.DeletedAt != nil {
		toSerialize["deletedAt"] = o.DeletedAt
	}
	if o.TruncatedAt != nil {
		toSerialize["truncatedAt"] = o.TruncatedAt
	}
	if true {
		toSerialize["shards"] = o.Shards
	}
	return json.Marshal(toSerialize)
}

type NullableShardGroupManifest struct {
	value *ShardGroupManifest
	isSet bool
}

func (v NullableShardGroupManifest) Get() *ShardGroupManifest {
	return v.value
}

func (v *NullableShardGroupManifest) Set(val *ShardGroupManifest) {
	v.value = val
	v.isSet = true
}

func (v NullableShardGroupManifest) IsSet() bool {
	return v.isSet
}

func (v *NullableShardGroupManifest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShardGroupManifest(val *ShardGroupManifest) *NullableShardGroupManifest {
	return &NullableShardGroupManifest{value: val, isSet: true}
}

func (v NullableShardGroupManifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShardGroupManifest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}