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

// Query Query influx using the Flux language
type Query struct {
	Extern *File `json:"extern,omitempty"`
	// Query script to execute.
	Query string `json:"query"`
	// The type of query. Must be \"flux\".
	Type    *string  `json:"type,omitempty"`
	Dialect *Dialect `json:"dialect,omitempty"`
	// Specifies the time that should be reported as \"now\" in the query. Default is the server's now time.
	Now *time.Time `json:"now,omitempty"`
}

// NewQuery instantiates a new Query object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuery(query string) *Query {
	this := Query{}
	this.Query = query
	return &this
}

// NewQueryWithDefaults instantiates a new Query object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryWithDefaults() *Query {
	this := Query{}
	return &this
}

// GetExtern returns the Extern field value if set, zero value otherwise.
func (o *Query) GetExtern() File {
	if o == nil || o.Extern == nil {
		var ret File
		return ret
	}
	return *o.Extern
}

// GetExternOk returns a tuple with the Extern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetExternOk() (*File, bool) {
	if o == nil || o.Extern == nil {
		return nil, false
	}
	return o.Extern, true
}

// HasExtern returns a boolean if a field has been set.
func (o *Query) HasExtern() bool {
	if o != nil && o.Extern != nil {
		return true
	}

	return false
}

// SetExtern gets a reference to the given File and assigns it to the Extern field.
func (o *Query) SetExtern(v File) {
	o.Extern = &v
}

// GetQuery returns the Query field value
func (o *Query) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *Query) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *Query) SetQuery(v string) {
	o.Query = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Query) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Query) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Query) SetType(v string) {
	o.Type = &v
}

// GetDialect returns the Dialect field value if set, zero value otherwise.
func (o *Query) GetDialect() Dialect {
	if o == nil || o.Dialect == nil {
		var ret Dialect
		return ret
	}
	return *o.Dialect
}

// GetDialectOk returns a tuple with the Dialect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetDialectOk() (*Dialect, bool) {
	if o == nil || o.Dialect == nil {
		return nil, false
	}
	return o.Dialect, true
}

// HasDialect returns a boolean if a field has been set.
func (o *Query) HasDialect() bool {
	if o != nil && o.Dialect != nil {
		return true
	}

	return false
}

// SetDialect gets a reference to the given Dialect and assigns it to the Dialect field.
func (o *Query) SetDialect(v Dialect) {
	o.Dialect = &v
}

// GetNow returns the Now field value if set, zero value otherwise.
func (o *Query) GetNow() time.Time {
	if o == nil || o.Now == nil {
		var ret time.Time
		return ret
	}
	return *o.Now
}

// GetNowOk returns a tuple with the Now field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetNowOk() (*time.Time, bool) {
	if o == nil || o.Now == nil {
		return nil, false
	}
	return o.Now, true
}

// HasNow returns a boolean if a field has been set.
func (o *Query) HasNow() bool {
	if o != nil && o.Now != nil {
		return true
	}

	return false
}

// SetNow gets a reference to the given time.Time and assigns it to the Now field.
func (o *Query) SetNow(v time.Time) {
	o.Now = &v
}

func (o Query) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Extern != nil {
		toSerialize["extern"] = o.Extern
	}
	if true {
		toSerialize["query"] = o.Query
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Dialect != nil {
		toSerialize["dialect"] = o.Dialect
	}
	if o.Now != nil {
		toSerialize["now"] = o.Now
	}
	return json.Marshal(toSerialize)
}

type NullableQuery struct {
	value *Query
	isSet bool
}

func (v NullableQuery) Get() *Query {
	return v.value
}

func (v *NullableQuery) Set(val *Query) {
	v.value = val
	v.isSet = true
}

func (v NullableQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuery(val *Query) *NullableQuery {
	return &NullableQuery{value: val, isSet: true}
}

func (v NullableQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}