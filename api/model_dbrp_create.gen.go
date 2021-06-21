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

// DBRPCreate struct for DBRPCreate
type DBRPCreate struct {
	// the organization ID that owns this mapping.
	OrgID string `json:"orgID"`
	// the organization that owns this mapping.
	Org *string `json:"org,omitempty"`
	// the bucket ID used as target for the translation.
	BucketID string `json:"bucketID"`
	// InfluxDB v1 database
	Database string `json:"database"`
	// InfluxDB v1 retention policy
	RetentionPolicy *string `json:"retention_policy,omitempty"`
	// Specify if this mapping represents the default retention policy for the database specificed.
	Default *bool `json:"default,omitempty"`
}

// NewDBRPCreate instantiates a new DBRPCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDBRPCreate(orgID string, bucketID string, database string) *DBRPCreate {
	this := DBRPCreate{}
	this.OrgID = orgID
	this.BucketID = bucketID
	this.Database = database
	return &this
}

// NewDBRPCreateWithDefaults instantiates a new DBRPCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDBRPCreateWithDefaults() *DBRPCreate {
	this := DBRPCreate{}
	return &this
}

// GetOrgID returns the OrgID field value
func (o *DBRPCreate) GetOrgID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgID
}

// GetOrgIDOk returns a tuple with the OrgID field value
// and a boolean to check if the value has been set.
func (o *DBRPCreate) GetOrgIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgID, true
}

// SetOrgID sets field value
func (o *DBRPCreate) SetOrgID(v string) {
	o.OrgID = v
}

// GetOrg returns the Org field value if set, zero value otherwise.
func (o *DBRPCreate) GetOrg() string {
	if o == nil || o.Org == nil {
		var ret string
		return ret
	}
	return *o.Org
}

// GetOrgOk returns a tuple with the Org field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBRPCreate) GetOrgOk() (*string, bool) {
	if o == nil || o.Org == nil {
		return nil, false
	}
	return o.Org, true
}

// HasOrg returns a boolean if a field has been set.
func (o *DBRPCreate) HasOrg() bool {
	if o != nil && o.Org != nil {
		return true
	}

	return false
}

// SetOrg gets a reference to the given string and assigns it to the Org field.
func (o *DBRPCreate) SetOrg(v string) {
	o.Org = &v
}

// GetBucketID returns the BucketID field value
func (o *DBRPCreate) GetBucketID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BucketID
}

// GetBucketIDOk returns a tuple with the BucketID field value
// and a boolean to check if the value has been set.
func (o *DBRPCreate) GetBucketIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketID, true
}

// SetBucketID sets field value
func (o *DBRPCreate) SetBucketID(v string) {
	o.BucketID = v
}

// GetDatabase returns the Database field value
func (o *DBRPCreate) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *DBRPCreate) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value
func (o *DBRPCreate) SetDatabase(v string) {
	o.Database = v
}

// GetRetentionPolicy returns the RetentionPolicy field value if set, zero value otherwise.
func (o *DBRPCreate) GetRetentionPolicy() string {
	if o == nil || o.RetentionPolicy == nil {
		var ret string
		return ret
	}
	return *o.RetentionPolicy
}

// GetRetentionPolicyOk returns a tuple with the RetentionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBRPCreate) GetRetentionPolicyOk() (*string, bool) {
	if o == nil || o.RetentionPolicy == nil {
		return nil, false
	}
	return o.RetentionPolicy, true
}

// HasRetentionPolicy returns a boolean if a field has been set.
func (o *DBRPCreate) HasRetentionPolicy() bool {
	if o != nil && o.RetentionPolicy != nil {
		return true
	}

	return false
}

// SetRetentionPolicy gets a reference to the given string and assigns it to the RetentionPolicy field.
func (o *DBRPCreate) SetRetentionPolicy(v string) {
	o.RetentionPolicy = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *DBRPCreate) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBRPCreate) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *DBRPCreate) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *DBRPCreate) SetDefault(v bool) {
	o.Default = &v
}

func (o DBRPCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["orgID"] = o.OrgID
	}
	if o.Org != nil {
		toSerialize["org"] = o.Org
	}
	if true {
		toSerialize["bucketID"] = o.BucketID
	}
	if true {
		toSerialize["database"] = o.Database
	}
	if o.RetentionPolicy != nil {
		toSerialize["retention_policy"] = o.RetentionPolicy
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	return json.Marshal(toSerialize)
}

type NullableDBRPCreate struct {
	value *DBRPCreate
	isSet bool
}

func (v NullableDBRPCreate) Get() *DBRPCreate {
	return v.value
}

func (v *NullableDBRPCreate) Set(val *DBRPCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableDBRPCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableDBRPCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDBRPCreate(val *DBRPCreate) *NullableDBRPCreate {
	return &NullableDBRPCreate{value: val, isSet: true}
}

func (v NullableDBRPCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDBRPCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
