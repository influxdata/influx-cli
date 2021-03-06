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

// BucketLinks struct for BucketLinks
type BucketLinks struct {
	// URI of resource.
	Labels *string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// URI of resource.
	Members *string `json:"members,omitempty" yaml:"members,omitempty"`
	// URI of resource.
	Org *string `json:"org,omitempty" yaml:"org,omitempty"`
	// URI of resource.
	Owners *string `json:"owners,omitempty" yaml:"owners,omitempty"`
	// URI of resource.
	Self *string `json:"self,omitempty" yaml:"self,omitempty"`
	// URI of resource.
	Write *string `json:"write,omitempty" yaml:"write,omitempty"`
}

// NewBucketLinks instantiates a new BucketLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucketLinks() *BucketLinks {
	this := BucketLinks{}
	return &this
}

// NewBucketLinksWithDefaults instantiates a new BucketLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketLinksWithDefaults() *BucketLinks {
	this := BucketLinks{}
	return &this
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *BucketLinks) GetLabels() string {
	if o == nil || o.Labels == nil {
		var ret string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketLinks) GetLabelsOk() (*string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *BucketLinks) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given string and assigns it to the Labels field.
func (o *BucketLinks) SetLabels(v string) {
	o.Labels = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *BucketLinks) GetMembers() string {
	if o == nil || o.Members == nil {
		var ret string
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketLinks) GetMembersOk() (*string, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *BucketLinks) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given string and assigns it to the Members field.
func (o *BucketLinks) SetMembers(v string) {
	o.Members = &v
}

// GetOrg returns the Org field value if set, zero value otherwise.
func (o *BucketLinks) GetOrg() string {
	if o == nil || o.Org == nil {
		var ret string
		return ret
	}
	return *o.Org
}

// GetOrgOk returns a tuple with the Org field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketLinks) GetOrgOk() (*string, bool) {
	if o == nil || o.Org == nil {
		return nil, false
	}
	return o.Org, true
}

// HasOrg returns a boolean if a field has been set.
func (o *BucketLinks) HasOrg() bool {
	if o != nil && o.Org != nil {
		return true
	}

	return false
}

// SetOrg gets a reference to the given string and assigns it to the Org field.
func (o *BucketLinks) SetOrg(v string) {
	o.Org = &v
}

// GetOwners returns the Owners field value if set, zero value otherwise.
func (o *BucketLinks) GetOwners() string {
	if o == nil || o.Owners == nil {
		var ret string
		return ret
	}
	return *o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketLinks) GetOwnersOk() (*string, bool) {
	if o == nil || o.Owners == nil {
		return nil, false
	}
	return o.Owners, true
}

// HasOwners returns a boolean if a field has been set.
func (o *BucketLinks) HasOwners() bool {
	if o != nil && o.Owners != nil {
		return true
	}

	return false
}

// SetOwners gets a reference to the given string and assigns it to the Owners field.
func (o *BucketLinks) SetOwners(v string) {
	o.Owners = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *BucketLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *BucketLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *BucketLinks) SetSelf(v string) {
	o.Self = &v
}

// GetWrite returns the Write field value if set, zero value otherwise.
func (o *BucketLinks) GetWrite() string {
	if o == nil || o.Write == nil {
		var ret string
		return ret
	}
	return *o.Write
}

// GetWriteOk returns a tuple with the Write field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketLinks) GetWriteOk() (*string, bool) {
	if o == nil || o.Write == nil {
		return nil, false
	}
	return o.Write, true
}

// HasWrite returns a boolean if a field has been set.
func (o *BucketLinks) HasWrite() bool {
	if o != nil && o.Write != nil {
		return true
	}

	return false
}

// SetWrite gets a reference to the given string and assigns it to the Write field.
func (o *BucketLinks) SetWrite(v string) {
	o.Write = &v
}

func (o BucketLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	if o.Org != nil {
		toSerialize["org"] = o.Org
	}
	if o.Owners != nil {
		toSerialize["owners"] = o.Owners
	}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Write != nil {
		toSerialize["write"] = o.Write
	}
	return json.Marshal(toSerialize)
}

type NullableBucketLinks struct {
	value *BucketLinks
	isSet bool
}

func (v NullableBucketLinks) Get() *BucketLinks {
	return v.value
}

func (v *NullableBucketLinks) Set(val *BucketLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketLinks(val *BucketLinks) *NullableBucketLinks {
	return &NullableBucketLinks{value: val, isSet: true}
}

func (v NullableBucketLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
