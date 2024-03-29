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

// PermissionResource struct for PermissionResource
type PermissionResource struct {
	// A resource type. Identifies the API resource's type (or _kind_).
	Type string `json:"type" yaml:"type"`
	// A resource ID. Identifies a specific resource.
	Id *string `json:"id,omitempty" yaml:"id,omitempty"`
	// The name of the resource. _Note: not all resource types have a `name` property_.
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
	// An organization ID. Identifies the organization that owns the resource.
	OrgID *string `json:"orgID,omitempty" yaml:"orgID,omitempty"`
	// An organization name. The organization that owns the resource.
	Org *string `json:"org,omitempty" yaml:"org,omitempty"`
}

// NewPermissionResource instantiates a new PermissionResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionResource(type_ string) *PermissionResource {
	this := PermissionResource{}
	this.Type = type_
	return &this
}

// NewPermissionResourceWithDefaults instantiates a new PermissionResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionResourceWithDefaults() *PermissionResource {
	this := PermissionResource{}
	return &this
}

// GetType returns the Type field value
func (o *PermissionResource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PermissionResource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PermissionResource) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PermissionResource) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionResource) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PermissionResource) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PermissionResource) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PermissionResource) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionResource) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PermissionResource) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PermissionResource) SetName(v string) {
	o.Name = &v
}

// GetOrgID returns the OrgID field value if set, zero value otherwise.
func (o *PermissionResource) GetOrgID() string {
	if o == nil || o.OrgID == nil {
		var ret string
		return ret
	}
	return *o.OrgID
}

// GetOrgIDOk returns a tuple with the OrgID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionResource) GetOrgIDOk() (*string, bool) {
	if o == nil || o.OrgID == nil {
		return nil, false
	}
	return o.OrgID, true
}

// HasOrgID returns a boolean if a field has been set.
func (o *PermissionResource) HasOrgID() bool {
	if o != nil && o.OrgID != nil {
		return true
	}

	return false
}

// SetOrgID gets a reference to the given string and assigns it to the OrgID field.
func (o *PermissionResource) SetOrgID(v string) {
	o.OrgID = &v
}

// GetOrg returns the Org field value if set, zero value otherwise.
func (o *PermissionResource) GetOrg() string {
	if o == nil || o.Org == nil {
		var ret string
		return ret
	}
	return *o.Org
}

// GetOrgOk returns a tuple with the Org field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionResource) GetOrgOk() (*string, bool) {
	if o == nil || o.Org == nil {
		return nil, false
	}
	return o.Org, true
}

// HasOrg returns a boolean if a field has been set.
func (o *PermissionResource) HasOrg() bool {
	if o != nil && o.Org != nil {
		return true
	}

	return false
}

// SetOrg gets a reference to the given string and assigns it to the Org field.
func (o *PermissionResource) SetOrg(v string) {
	o.Org = &v
}

func (o PermissionResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OrgID != nil {
		toSerialize["orgID"] = o.OrgID
	}
	if o.Org != nil {
		toSerialize["org"] = o.Org
	}
	return json.Marshal(toSerialize)
}

type NullablePermissionResource struct {
	value *PermissionResource
	isSet bool
}

func (v NullablePermissionResource) Get() *PermissionResource {
	return v.value
}

func (v *NullablePermissionResource) Set(val *PermissionResource) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionResource) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionResource(val *PermissionResource) *NullablePermissionResource {
	return &NullablePermissionResource{value: val, isSet: true}
}

func (v NullablePermissionResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
