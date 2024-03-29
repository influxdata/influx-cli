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

// ShardOwner struct for ShardOwner
type ShardOwner struct {
	// The ID of the node that owns the shard.
	NodeID int64 `json:"nodeID" yaml:"nodeID"`
}

// NewShardOwner instantiates a new ShardOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardOwner(nodeID int64) *ShardOwner {
	this := ShardOwner{}
	this.NodeID = nodeID
	return &this
}

// NewShardOwnerWithDefaults instantiates a new ShardOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardOwnerWithDefaults() *ShardOwner {
	this := ShardOwner{}
	return &this
}

// GetNodeID returns the NodeID field value
func (o *ShardOwner) GetNodeID() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NodeID
}

// GetNodeIDOk returns a tuple with the NodeID field value
// and a boolean to check if the value has been set.
func (o *ShardOwner) GetNodeIDOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeID, true
}

// SetNodeID sets field value
func (o *ShardOwner) SetNodeID(v int64) {
	o.NodeID = v
}

func (o ShardOwner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nodeID"] = o.NodeID
	}
	return json.Marshal(toSerialize)
}

type NullableShardOwner struct {
	value *ShardOwner
	isSet bool
}

func (v NullableShardOwner) Get() *ShardOwner {
	return v.value
}

func (v *NullableShardOwner) Set(val *ShardOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableShardOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableShardOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShardOwner(val *ShardOwner) *NullableShardOwner {
	return &NullableShardOwner{value: val, isSet: true}
}

func (v NullableShardOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShardOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
