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

// ReplicationCreationRequest struct for ReplicationCreationRequest
type ReplicationCreationRequest struct {
	Name              string  `json:"name" yaml:"name"`
	Description       *string `json:"description,omitempty" yaml:"description,omitempty"`
	OrgID             string  `json:"orgID" yaml:"orgID"`
	RemoteID          string  `json:"remoteID" yaml:"remoteID"`
	LocalBucketID     *string `json:"localBucketID,omitempty" yaml:"localBucketID,omitempty"`
	RemoteBucketID    *string `json:"remoteBucketID,omitempty" yaml:"remoteBucketID,omitempty"`
	MaxQueueSizeBytes int64   `json:"maxQueueSizeBytes" yaml:"maxQueueSizeBytes"`
}

// NewReplicationCreationRequest instantiates a new ReplicationCreationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicationCreationRequest(name string, orgID string, remoteID string, maxQueueSizeBytes int64) *ReplicationCreationRequest {
	this := ReplicationCreationRequest{}
	this.Name = name
	this.OrgID = orgID
	this.RemoteID = remoteID
	this.MaxQueueSizeBytes = maxQueueSizeBytes
	return &this
}

// NewReplicationCreationRequestWithDefaults instantiates a new ReplicationCreationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicationCreationRequestWithDefaults() *ReplicationCreationRequest {
	this := ReplicationCreationRequest{}
	var maxQueueSizeBytes int64 = 67108860
	this.MaxQueueSizeBytes = maxQueueSizeBytes
	return &this
}

// GetName returns the Name field value
func (o *ReplicationCreationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReplicationCreationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ReplicationCreationRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ReplicationCreationRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationCreationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ReplicationCreationRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ReplicationCreationRequest) SetDescription(v string) {
	o.Description = &v
}

// GetOrgID returns the OrgID field value
func (o *ReplicationCreationRequest) GetOrgID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgID
}

// GetOrgIDOk returns a tuple with the OrgID field value
// and a boolean to check if the value has been set.
func (o *ReplicationCreationRequest) GetOrgIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgID, true
}

// SetOrgID sets field value
func (o *ReplicationCreationRequest) SetOrgID(v string) {
	o.OrgID = v
}

// GetRemoteID returns the RemoteID field value
func (o *ReplicationCreationRequest) GetRemoteID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteID
}

// GetRemoteIDOk returns a tuple with the RemoteID field value
// and a boolean to check if the value has been set.
func (o *ReplicationCreationRequest) GetRemoteIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteID, true
}

// SetRemoteID sets field value
func (o *ReplicationCreationRequest) SetRemoteID(v string) {
	o.RemoteID = v
}

// GetLocalBucketID returns the LocalBucketID field value if set, zero value otherwise.
func (o *ReplicationCreationRequest) GetLocalBucketID() string {
	if o == nil || o.LocalBucketID == nil {
		var ret string
		return ret
	}
	return *o.LocalBucketID
}

// GetLocalBucketIDOk returns a tuple with the LocalBucketID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationCreationRequest) GetLocalBucketIDOk() (*string, bool) {
	if o == nil || o.LocalBucketID == nil {
		return nil, false
	}
	return o.LocalBucketID, true
}

// HasLocalBucketID returns a boolean if a field has been set.
func (o *ReplicationCreationRequest) HasLocalBucketID() bool {
	if o != nil && o.LocalBucketID != nil {
		return true
	}

	return false
}

// SetLocalBucketID gets a reference to the given string and assigns it to the LocalBucketID field.
func (o *ReplicationCreationRequest) SetLocalBucketID(v string) {
	o.LocalBucketID = &v
}

// GetRemoteBucketID returns the RemoteBucketID field value if set, zero value otherwise.
func (o *ReplicationCreationRequest) GetRemoteBucketID() string {
	if o == nil || o.RemoteBucketID == nil {
		var ret string
		return ret
	}
	return *o.RemoteBucketID
}

// GetRemoteBucketIDOk returns a tuple with the RemoteBucketID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationCreationRequest) GetRemoteBucketIDOk() (*string, bool) {
	if o == nil || o.RemoteBucketID == nil {
		return nil, false
	}
	return o.RemoteBucketID, true
}

// HasRemoteBucketID returns a boolean if a field has been set.
func (o *ReplicationCreationRequest) HasRemoteBucketID() bool {
	if o != nil && o.RemoteBucketID != nil {
		return true
	}

	return false
}

// SetRemoteBucketID gets a reference to the given string and assigns it to the RemoteBucketID field.
func (o *ReplicationCreationRequest) SetRemoteBucketID(v string) {
	o.RemoteBucketID = &v
}

// GetMaxQueueSizeBytes returns the MaxQueueSizeBytes field value
func (o *ReplicationCreationRequest) GetMaxQueueSizeBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxQueueSizeBytes
}

// GetMaxQueueSizeBytesOk returns a tuple with the MaxQueueSizeBytes field value
// and a boolean to check if the value has been set.
func (o *ReplicationCreationRequest) GetMaxQueueSizeBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxQueueSizeBytes, true
}

// SetMaxQueueSizeBytes sets field value
func (o *ReplicationCreationRequest) SetMaxQueueSizeBytes(v int64) {
	o.MaxQueueSizeBytes = v
}

func (o ReplicationCreationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["orgID"] = o.OrgID
	}
	if true {
		toSerialize["remoteID"] = o.RemoteID
	}
	if o.LocalBucketID != nil {
		toSerialize["localBucketID"] = o.LocalBucketID
	}
	if o.RemoteBucketID != nil {
		toSerialize["remoteBucketID"] = o.RemoteBucketID
	}
	if true {
		toSerialize["maxQueueSizeBytes"] = o.MaxQueueSizeBytes
	}
	return json.Marshal(toSerialize)
}

type NullableReplicationCreationRequest struct {
	value *ReplicationCreationRequest
	isSet bool
}

func (v NullableReplicationCreationRequest) Get() *ReplicationCreationRequest {
	return v.value
}

func (v *NullableReplicationCreationRequest) Set(val *ReplicationCreationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicationCreationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicationCreationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicationCreationRequest(val *ReplicationCreationRequest) *NullableReplicationCreationRequest {
	return &NullableReplicationCreationRequest{value: val, isSet: true}
}

func (v NullableReplicationCreationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicationCreationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
