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

// TemplateSummaryDiff struct for TemplateSummaryDiff
type TemplateSummaryDiff struct {
	Buckets               *[]TemplateSummaryDiffBucket               `json:"buckets,omitempty"`
	Checks                *[]TemplateSummaryDiffCheck                `json:"checks,omitempty"`
	Dashboards            *[]TemplateSummaryDiffDashboard            `json:"dashboards,omitempty"`
	Labels                *[]TemplateSummaryDiffLabel                `json:"labels,omitempty"`
	LabelMappings         *[]TemplateSummaryLabelMapping             `json:"labelMappings,omitempty"`
	NotificationEndpoints *[]TemplateSummaryDiffNotificationEndpoint `json:"notificationEndpoints,omitempty"`
	NotificationRules     *[]TemplateSummaryDiffNotificationRule     `json:"notificationRules,omitempty"`
	Tasks                 *[]TemplateSummaryDiffTask                 `json:"tasks,omitempty"`
	TelegrafConfigs       *[]TemplateSummaryDiffTelegraf             `json:"telegrafConfigs,omitempty"`
	Variables             *[]TemplateSummaryDiffVariable             `json:"variables,omitempty"`
}

// NewTemplateSummaryDiff instantiates a new TemplateSummaryDiff object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryDiff() *TemplateSummaryDiff {
	this := TemplateSummaryDiff{}
	return &this
}

// NewTemplateSummaryDiffWithDefaults instantiates a new TemplateSummaryDiff object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryDiffWithDefaults() *TemplateSummaryDiff {
	this := TemplateSummaryDiff{}
	return &this
}

// GetBuckets returns the Buckets field value if set, zero value otherwise.
func (o *TemplateSummaryDiff) GetBuckets() []TemplateSummaryDiffBucket {
	if o == nil || o.Buckets == nil {
		var ret []TemplateSummaryDiffBucket
		return ret
	}
	return *o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiff) GetBucketsOk() (*[]TemplateSummaryDiffBucket, bool) {
	if o == nil || o.Buckets == nil {
		return nil, false
	}
	return o.Buckets, true
}

// HasBuckets returns a boolean if a field has been set.
func (o *TemplateSummaryDiff) HasBuckets() bool {
	if o != nil && o.Buckets != nil {
		return true
	}

	return false
}

// SetBuckets gets a reference to the given []TemplateSummaryDiffBucket and assigns it to the Buckets field.
func (o *TemplateSummaryDiff) SetBuckets(v []TemplateSummaryDiffBucket) {
	o.Buckets = &v
}

// GetChecks returns the Checks field value if set, zero value otherwise.
func (o *TemplateSummaryDiff) GetChecks() []TemplateSummaryDiffCheck {
	if o == nil || o.Checks == nil {
		var ret []TemplateSummaryDiffCheck
		return ret
	}
	return *o.Checks
}

// GetChecksOk returns a tuple with the Checks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiff) GetChecksOk() (*[]TemplateSummaryDiffCheck, bool) {
	if o == nil || o.Checks == nil {
		return nil, false
	}
	return o.Checks, true
}

// HasChecks returns a boolean if a field has been set.
func (o *TemplateSummaryDiff) HasChecks() bool {
	if o != nil && o.Checks != nil {
		return true
	}

	return false
}

// SetChecks gets a reference to the given []TemplateSummaryDiffCheck and assigns it to the Checks field.
func (o *TemplateSummaryDiff) SetChecks(v []TemplateSummaryDiffCheck) {
	o.Checks = &v
}

// GetDashboards returns the Dashboards field value if set, zero value otherwise.
func (o *TemplateSummaryDiff) GetDashboards() []TemplateSummaryDiffDashboard {
	if o == nil || o.Dashboards == nil {
		var ret []TemplateSummaryDiffDashboard
		return ret
	}
	return *o.Dashboards
}

// GetDashboardsOk returns a tuple with the Dashboards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiff) GetDashboardsOk() (*[]TemplateSummaryDiffDashboard, bool) {
	if o == nil || o.Dashboards == nil {
		return nil, false
	}
	return o.Dashboards, true
}

// HasDashboards returns a boolean if a field has been set.
func (o *TemplateSummaryDiff) HasDashboards() bool {
	if o != nil && o.Dashboards != nil {
		return true
	}

	return false
}

// SetDashboards gets a reference to the given []TemplateSummaryDiffDashboard and assigns it to the Dashboards field.
func (o *TemplateSummaryDiff) SetDashboards(v []TemplateSummaryDiffDashboard) {
	o.Dashboards = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *TemplateSummaryDiff) GetLabels() []TemplateSummaryDiffLabel {
	if o == nil || o.Labels == nil {
		var ret []TemplateSummaryDiffLabel
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiff) GetLabelsOk() (*[]TemplateSummaryDiffLabel, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *TemplateSummaryDiff) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []TemplateSummaryDiffLabel and assigns it to the Labels field.
func (o *TemplateSummaryDiff) SetLabels(v []TemplateSummaryDiffLabel) {
	o.Labels = &v
}

// GetLabelMappings returns the LabelMappings field value if set, zero value otherwise.
func (o *TemplateSummaryDiff) GetLabelMappings() []TemplateSummaryLabelMapping {
	if o == nil || o.LabelMappings == nil {
		var ret []TemplateSummaryLabelMapping
		return ret
	}
	return *o.LabelMappings
}

// GetLabelMappingsOk returns a tuple with the LabelMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiff) GetLabelMappingsOk() (*[]TemplateSummaryLabelMapping, bool) {
	if o == nil || o.LabelMappings == nil {
		return nil, false
	}
	return o.LabelMappings, true
}

// HasLabelMappings returns a boolean if a field has been set.
func (o *TemplateSummaryDiff) HasLabelMappings() bool {
	if o != nil && o.LabelMappings != nil {
		return true
	}

	return false
}

// SetLabelMappings gets a reference to the given []TemplateSummaryLabelMapping and assigns it to the LabelMappings field.
func (o *TemplateSummaryDiff) SetLabelMappings(v []TemplateSummaryLabelMapping) {
	o.LabelMappings = &v
}

// GetNotificationEndpoints returns the NotificationEndpoints field value if set, zero value otherwise.
func (o *TemplateSummaryDiff) GetNotificationEndpoints() []TemplateSummaryDiffNotificationEndpoint {
	if o == nil || o.NotificationEndpoints == nil {
		var ret []TemplateSummaryDiffNotificationEndpoint
		return ret
	}
	return *o.NotificationEndpoints
}

// GetNotificationEndpointsOk returns a tuple with the NotificationEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiff) GetNotificationEndpointsOk() (*[]TemplateSummaryDiffNotificationEndpoint, bool) {
	if o == nil || o.NotificationEndpoints == nil {
		return nil, false
	}
	return o.NotificationEndpoints, true
}

// HasNotificationEndpoints returns a boolean if a field has been set.
func (o *TemplateSummaryDiff) HasNotificationEndpoints() bool {
	if o != nil && o.NotificationEndpoints != nil {
		return true
	}

	return false
}

// SetNotificationEndpoints gets a reference to the given []TemplateSummaryDiffNotificationEndpoint and assigns it to the NotificationEndpoints field.
func (o *TemplateSummaryDiff) SetNotificationEndpoints(v []TemplateSummaryDiffNotificationEndpoint) {
	o.NotificationEndpoints = &v
}

// GetNotificationRules returns the NotificationRules field value if set, zero value otherwise.
func (o *TemplateSummaryDiff) GetNotificationRules() []TemplateSummaryDiffNotificationRule {
	if o == nil || o.NotificationRules == nil {
		var ret []TemplateSummaryDiffNotificationRule
		return ret
	}
	return *o.NotificationRules
}

// GetNotificationRulesOk returns a tuple with the NotificationRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiff) GetNotificationRulesOk() (*[]TemplateSummaryDiffNotificationRule, bool) {
	if o == nil || o.NotificationRules == nil {
		return nil, false
	}
	return o.NotificationRules, true
}

// HasNotificationRules returns a boolean if a field has been set.
func (o *TemplateSummaryDiff) HasNotificationRules() bool {
	if o != nil && o.NotificationRules != nil {
		return true
	}

	return false
}

// SetNotificationRules gets a reference to the given []TemplateSummaryDiffNotificationRule and assigns it to the NotificationRules field.
func (o *TemplateSummaryDiff) SetNotificationRules(v []TemplateSummaryDiffNotificationRule) {
	o.NotificationRules = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *TemplateSummaryDiff) GetTasks() []TemplateSummaryDiffTask {
	if o == nil || o.Tasks == nil {
		var ret []TemplateSummaryDiffTask
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiff) GetTasksOk() (*[]TemplateSummaryDiffTask, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *TemplateSummaryDiff) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []TemplateSummaryDiffTask and assigns it to the Tasks field.
func (o *TemplateSummaryDiff) SetTasks(v []TemplateSummaryDiffTask) {
	o.Tasks = &v
}

// GetTelegrafConfigs returns the TelegrafConfigs field value if set, zero value otherwise.
func (o *TemplateSummaryDiff) GetTelegrafConfigs() []TemplateSummaryDiffTelegraf {
	if o == nil || o.TelegrafConfigs == nil {
		var ret []TemplateSummaryDiffTelegraf
		return ret
	}
	return *o.TelegrafConfigs
}

// GetTelegrafConfigsOk returns a tuple with the TelegrafConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiff) GetTelegrafConfigsOk() (*[]TemplateSummaryDiffTelegraf, bool) {
	if o == nil || o.TelegrafConfigs == nil {
		return nil, false
	}
	return o.TelegrafConfigs, true
}

// HasTelegrafConfigs returns a boolean if a field has been set.
func (o *TemplateSummaryDiff) HasTelegrafConfigs() bool {
	if o != nil && o.TelegrafConfigs != nil {
		return true
	}

	return false
}

// SetTelegrafConfigs gets a reference to the given []TemplateSummaryDiffTelegraf and assigns it to the TelegrafConfigs field.
func (o *TemplateSummaryDiff) SetTelegrafConfigs(v []TemplateSummaryDiffTelegraf) {
	o.TelegrafConfigs = &v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *TemplateSummaryDiff) GetVariables() []TemplateSummaryDiffVariable {
	if o == nil || o.Variables == nil {
		var ret []TemplateSummaryDiffVariable
		return ret
	}
	return *o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiff) GetVariablesOk() (*[]TemplateSummaryDiffVariable, bool) {
	if o == nil || o.Variables == nil {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *TemplateSummaryDiff) HasVariables() bool {
	if o != nil && o.Variables != nil {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []TemplateSummaryDiffVariable and assigns it to the Variables field.
func (o *TemplateSummaryDiff) SetVariables(v []TemplateSummaryDiffVariable) {
	o.Variables = &v
}

func (o TemplateSummaryDiff) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Buckets != nil {
		toSerialize["buckets"] = o.Buckets
	}
	if o.Checks != nil {
		toSerialize["checks"] = o.Checks
	}
	if o.Dashboards != nil {
		toSerialize["dashboards"] = o.Dashboards
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.LabelMappings != nil {
		toSerialize["labelMappings"] = o.LabelMappings
	}
	if o.NotificationEndpoints != nil {
		toSerialize["notificationEndpoints"] = o.NotificationEndpoints
	}
	if o.NotificationRules != nil {
		toSerialize["notificationRules"] = o.NotificationRules
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	if o.TelegrafConfigs != nil {
		toSerialize["telegrafConfigs"] = o.TelegrafConfigs
	}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryDiff struct {
	value *TemplateSummaryDiff
	isSet bool
}

func (v NullableTemplateSummaryDiff) Get() *TemplateSummaryDiff {
	return v.value
}

func (v *NullableTemplateSummaryDiff) Set(val *TemplateSummaryDiff) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryDiff) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryDiff) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryDiff(val *TemplateSummaryDiff) *NullableTemplateSummaryDiff {
	return &NullableTemplateSummaryDiff{value: val, isSet: true}
}

func (v NullableTemplateSummaryDiff) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryDiff) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
