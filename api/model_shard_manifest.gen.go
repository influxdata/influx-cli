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

// ShardManifest struct for ShardManifest
type ShardManifest struct {
	Id          int64        `json:"id" yaml:"id"`
	ShardOwners []ShardOwner `json:"shardOwners" yaml:"shardOwners"`
}

// NewShardManifest instantiates a new ShardManifest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardManifest(id int64, shardOwners []ShardOwner) *ShardManifest {
	this := ShardManifest{}
	this.Id = id
	this.ShardOwners = shardOwners
	return &this
}

// NewShardManifestWithDefaults instantiates a new ShardManifest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardManifestWithDefaults() *ShardManifest {
	this := ShardManifest{}
	return &this
}

// GetId returns the Id field value
func (o *ShardManifest) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ShardManifest) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ShardManifest) SetId(v int64) {
	o.Id = v
}

// GetShardOwners returns the ShardOwners field value
func (o *ShardManifest) GetShardOwners() []ShardOwner {
	if o == nil {
		var ret []ShardOwner
		return ret
	}

	return o.ShardOwners
}

// GetShardOwnersOk returns a tuple with the ShardOwners field value
// and a boolean to check if the value has been set.
func (o *ShardManifest) GetShardOwnersOk() (*[]ShardOwner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShardOwners, true
}

// SetShardOwners sets field value
func (o *ShardManifest) SetShardOwners(v []ShardOwner) {
	o.ShardOwners = v
}

func (o ShardManifest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["shardOwners"] = o.ShardOwners
	}
	return json.Marshal(toSerialize)
}

type NullableShardManifest struct {
	value *ShardManifest
	isSet bool
}

func (v NullableShardManifest) Get() *ShardManifest {
	return v.value
}

func (v *NullableShardManifest) Set(val *ShardManifest) {
	v.value = val
	v.isSet = true
}

func (v NullableShardManifest) IsSet() bool {
	return v.isSet
}

func (v *NullableShardManifest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShardManifest(val *ShardManifest) *NullableShardManifest {
	return &NullableShardManifest{value: val, isSet: true}
}

func (v NullableShardManifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShardManifest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
