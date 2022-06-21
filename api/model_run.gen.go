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

// Run struct for Run
type Run struct {
	Id     *string `json:"id,omitempty" yaml:"id,omitempty"`
	TaskID *string `json:"taskID,omitempty" yaml:"taskID,omitempty"`
	Status *string `json:"status,omitempty" yaml:"status,omitempty"`
	// Time used for run's \"now\" option, RFC3339.
	ScheduledFor *time.Time `json:"scheduledFor,omitempty" yaml:"scheduledFor,omitempty"`
	// An array of logs associated with the run.
	Log *[]LogEvent `json:"log,omitempty" yaml:"log,omitempty"`
	// Flux used for the task
	Flux *string `json:"flux,omitempty" yaml:"flux,omitempty"`
	// Time run started executing, RFC3339Nano.
	StartedAt *time.Time `json:"startedAt,omitempty" yaml:"startedAt,omitempty"`
	// Time run finished executing, RFC3339Nano.
	FinishedAt *time.Time `json:"finishedAt,omitempty" yaml:"finishedAt,omitempty"`
	// Time run was manually requested, RFC3339Nano.
	RequestedAt *time.Time `json:"requestedAt,omitempty" yaml:"requestedAt,omitempty"`
	Links       *RunLinks  `json:"links,omitempty" yaml:"links,omitempty"`
}

// NewRun instantiates a new Run object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRun() *Run {
	this := Run{}
	return &this
}

// NewRunWithDefaults instantiates a new Run object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunWithDefaults() *Run {
	this := Run{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Run) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Run) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Run) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Run) SetId(v string) {
	o.Id = &v
}

// GetTaskID returns the TaskID field value if set, zero value otherwise.
func (o *Run) GetTaskID() string {
	if o == nil || o.TaskID == nil {
		var ret string
		return ret
	}
	return *o.TaskID
}

// GetTaskIDOk returns a tuple with the TaskID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Run) GetTaskIDOk() (*string, bool) {
	if o == nil || o.TaskID == nil {
		return nil, false
	}
	return o.TaskID, true
}

// HasTaskID returns a boolean if a field has been set.
func (o *Run) HasTaskID() bool {
	if o != nil && o.TaskID != nil {
		return true
	}

	return false
}

// SetTaskID gets a reference to the given string and assigns it to the TaskID field.
func (o *Run) SetTaskID(v string) {
	o.TaskID = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Run) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Run) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Run) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Run) SetStatus(v string) {
	o.Status = &v
}

// GetScheduledFor returns the ScheduledFor field value if set, zero value otherwise.
func (o *Run) GetScheduledFor() time.Time {
	if o == nil || o.ScheduledFor == nil {
		var ret time.Time
		return ret
	}
	return *o.ScheduledFor
}

// GetScheduledForOk returns a tuple with the ScheduledFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Run) GetScheduledForOk() (*time.Time, bool) {
	if o == nil || o.ScheduledFor == nil {
		return nil, false
	}
	return o.ScheduledFor, true
}

// HasScheduledFor returns a boolean if a field has been set.
func (o *Run) HasScheduledFor() bool {
	if o != nil && o.ScheduledFor != nil {
		return true
	}

	return false
}

// SetScheduledFor gets a reference to the given time.Time and assigns it to the ScheduledFor field.
func (o *Run) SetScheduledFor(v time.Time) {
	o.ScheduledFor = &v
}

// GetLog returns the Log field value if set, zero value otherwise.
func (o *Run) GetLog() []LogEvent {
	if o == nil || o.Log == nil {
		var ret []LogEvent
		return ret
	}
	return *o.Log
}

// GetLogOk returns a tuple with the Log field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Run) GetLogOk() (*[]LogEvent, bool) {
	if o == nil || o.Log == nil {
		return nil, false
	}
	return o.Log, true
}

// HasLog returns a boolean if a field has been set.
func (o *Run) HasLog() bool {
	if o != nil && o.Log != nil {
		return true
	}

	return false
}

// SetLog gets a reference to the given []LogEvent and assigns it to the Log field.
func (o *Run) SetLog(v []LogEvent) {
	o.Log = &v
}

// GetFlux returns the Flux field value if set, zero value otherwise.
func (o *Run) GetFlux() string {
	if o == nil || o.Flux == nil {
		var ret string
		return ret
	}
	return *o.Flux
}

// GetFluxOk returns a tuple with the Flux field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Run) GetFluxOk() (*string, bool) {
	if o == nil || o.Flux == nil {
		return nil, false
	}
	return o.Flux, true
}

// HasFlux returns a boolean if a field has been set.
func (o *Run) HasFlux() bool {
	if o != nil && o.Flux != nil {
		return true
	}

	return false
}

// SetFlux gets a reference to the given string and assigns it to the Flux field.
func (o *Run) SetFlux(v string) {
	o.Flux = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *Run) GetStartedAt() time.Time {
	if o == nil || o.StartedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Run) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *Run) HasStartedAt() bool {
	if o != nil && o.StartedAt != nil {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *Run) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *Run) GetFinishedAt() time.Time {
	if o == nil || o.FinishedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Run) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || o.FinishedAt == nil {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *Run) HasFinishedAt() bool {
	if o != nil && o.FinishedAt != nil {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *Run) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// GetRequestedAt returns the RequestedAt field value if set, zero value otherwise.
func (o *Run) GetRequestedAt() time.Time {
	if o == nil || o.RequestedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.RequestedAt
}

// GetRequestedAtOk returns a tuple with the RequestedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Run) GetRequestedAtOk() (*time.Time, bool) {
	if o == nil || o.RequestedAt == nil {
		return nil, false
	}
	return o.RequestedAt, true
}

// HasRequestedAt returns a boolean if a field has been set.
func (o *Run) HasRequestedAt() bool {
	if o != nil && o.RequestedAt != nil {
		return true
	}

	return false
}

// SetRequestedAt gets a reference to the given time.Time and assigns it to the RequestedAt field.
func (o *Run) SetRequestedAt(v time.Time) {
	o.RequestedAt = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Run) GetLinks() RunLinks {
	if o == nil || o.Links == nil {
		var ret RunLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Run) GetLinksOk() (*RunLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Run) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RunLinks and assigns it to the Links field.
func (o *Run) SetLinks(v RunLinks) {
	o.Links = &v
}

func (o Run) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.TaskID != nil {
		toSerialize["taskID"] = o.TaskID
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ScheduledFor != nil {
		toSerialize["scheduledFor"] = o.ScheduledFor
	}
	if o.Log != nil {
		toSerialize["log"] = o.Log
	}
	if o.Flux != nil {
		toSerialize["flux"] = o.Flux
	}
	if o.StartedAt != nil {
		toSerialize["startedAt"] = o.StartedAt
	}
	if o.FinishedAt != nil {
		toSerialize["finishedAt"] = o.FinishedAt
	}
	if o.RequestedAt != nil {
		toSerialize["requestedAt"] = o.RequestedAt
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableRun struct {
	value *Run
	isSet bool
}

func (v NullableRun) Get() *Run {
	return v.value
}

func (v *NullableRun) Set(val *Run) {
	v.value = val
	v.isSet = true
}

func (v NullableRun) IsSet() bool {
	return v.isSet
}

func (v *NullableRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRun(val *Run) *NullableRun {
	return &NullableRun{value: val, isSet: true}
}

func (v NullableRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
