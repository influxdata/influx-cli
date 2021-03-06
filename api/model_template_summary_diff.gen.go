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
	Buckets               []TemplateSummaryDiffBucket               `json:"buckets" yaml:"buckets"`
	Checks                []TemplateSummaryDiffCheck                `json:"checks" yaml:"checks"`
	Dashboards            []TemplateSummaryDiffDashboard            `json:"dashboards" yaml:"dashboards"`
	Labels                []TemplateSummaryDiffLabel                `json:"labels" yaml:"labels"`
	LabelMappings         []TemplateSummaryDiffLabelMapping         `json:"labelMappings" yaml:"labelMappings"`
	NotificationEndpoints []TemplateSummaryDiffNotificationEndpoint `json:"notificationEndpoints" yaml:"notificationEndpoints"`
	NotificationRules     []TemplateSummaryDiffNotificationRule     `json:"notificationRules" yaml:"notificationRules"`
	Tasks                 []TemplateSummaryDiffTask                 `json:"tasks" yaml:"tasks"`
	TelegrafConfigs       []TemplateSummaryDiffTelegraf             `json:"telegrafConfigs" yaml:"telegrafConfigs"`
	Variables             []TemplateSummaryDiffVariable             `json:"variables" yaml:"variables"`
}

// NewTemplateSummaryDiff instantiates a new TemplateSummaryDiff object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryDiff(buckets []TemplateSummaryDiffBucket, checks []TemplateSummaryDiffCheck, dashboards []TemplateSummaryDiffDashboard, labels []TemplateSummaryDiffLabel, labelMappings []TemplateSummaryDiffLabelMapping, notificationEndpoints []TemplateSummaryDiffNotificationEndpoint, notificationRules []TemplateSummaryDiffNotificationRule, tasks []TemplateSummaryDiffTask, telegrafConfigs []TemplateSummaryDiffTelegraf, variables []TemplateSummaryDiffVariable) *TemplateSummaryDiff {
	this := TemplateSummaryDiff{}
	this.Buckets = buckets
	this.Checks = checks
	this.Dashboards = dashboards
	this.Labels = labels
	this.LabelMappings = labelMappings
	this.NotificationEndpoints = notificationEndpoints
	this.NotificationRules = notificationRules
	this.Tasks = tasks
	this.TelegrafConfigs = telegrafConfigs
	this.Variables = variables
	return &this
}

// NewTemplateSummaryDiffWithDefaults instantiates a new TemplateSummaryDiff object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryDiffWithDefaults() *TemplateSummaryDiff {
	this := TemplateSummaryDiff{}
	return &this
}

// GetBuckets returns the Buckets field value
func (o *TemplateSummaryDiff) GetBuckets() []TemplateSummaryDiffBucket {
	if o == nil {
		var ret []TemplateSummaryDiffBucket
		return ret
	}

	return o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiff) GetBucketsOk() (*[]TemplateSummaryDiffBucket, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Buckets, true
}

// SetBuckets sets field value
func (o *TemplateSummaryDiff) SetBuckets(v []TemplateSummaryDiffBucket) {
	o.Buckets = v
}

// GetChecks returns the Checks field value
func (o *TemplateSummaryDiff) GetChecks() []TemplateSummaryDiffCheck {
	if o == nil {
		var ret []TemplateSummaryDiffCheck
		return ret
	}

	return o.Checks
}

// GetChecksOk returns a tuple with the Checks field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiff) GetChecksOk() (*[]TemplateSummaryDiffCheck, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Checks, true
}

// SetChecks sets field value
func (o *TemplateSummaryDiff) SetChecks(v []TemplateSummaryDiffCheck) {
	o.Checks = v
}

// GetDashboards returns the Dashboards field value
func (o *TemplateSummaryDiff) GetDashboards() []TemplateSummaryDiffDashboard {
	if o == nil {
		var ret []TemplateSummaryDiffDashboard
		return ret
	}

	return o.Dashboards
}

// GetDashboardsOk returns a tuple with the Dashboards field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiff) GetDashboardsOk() (*[]TemplateSummaryDiffDashboard, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dashboards, true
}

// SetDashboards sets field value
func (o *TemplateSummaryDiff) SetDashboards(v []TemplateSummaryDiffDashboard) {
	o.Dashboards = v
}

// GetLabels returns the Labels field value
func (o *TemplateSummaryDiff) GetLabels() []TemplateSummaryDiffLabel {
	if o == nil {
		var ret []TemplateSummaryDiffLabel
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiff) GetLabelsOk() (*[]TemplateSummaryDiffLabel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Labels, true
}

// SetLabels sets field value
func (o *TemplateSummaryDiff) SetLabels(v []TemplateSummaryDiffLabel) {
	o.Labels = v
}

// GetLabelMappings returns the LabelMappings field value
func (o *TemplateSummaryDiff) GetLabelMappings() []TemplateSummaryDiffLabelMapping {
	if o == nil {
		var ret []TemplateSummaryDiffLabelMapping
		return ret
	}

	return o.LabelMappings
}

// GetLabelMappingsOk returns a tuple with the LabelMappings field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiff) GetLabelMappingsOk() (*[]TemplateSummaryDiffLabelMapping, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelMappings, true
}

// SetLabelMappings sets field value
func (o *TemplateSummaryDiff) SetLabelMappings(v []TemplateSummaryDiffLabelMapping) {
	o.LabelMappings = v
}

// GetNotificationEndpoints returns the NotificationEndpoints field value
func (o *TemplateSummaryDiff) GetNotificationEndpoints() []TemplateSummaryDiffNotificationEndpoint {
	if o == nil {
		var ret []TemplateSummaryDiffNotificationEndpoint
		return ret
	}

	return o.NotificationEndpoints
}

// GetNotificationEndpointsOk returns a tuple with the NotificationEndpoints field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiff) GetNotificationEndpointsOk() (*[]TemplateSummaryDiffNotificationEndpoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationEndpoints, true
}

// SetNotificationEndpoints sets field value
func (o *TemplateSummaryDiff) SetNotificationEndpoints(v []TemplateSummaryDiffNotificationEndpoint) {
	o.NotificationEndpoints = v
}

// GetNotificationRules returns the NotificationRules field value
func (o *TemplateSummaryDiff) GetNotificationRules() []TemplateSummaryDiffNotificationRule {
	if o == nil {
		var ret []TemplateSummaryDiffNotificationRule
		return ret
	}

	return o.NotificationRules
}

// GetNotificationRulesOk returns a tuple with the NotificationRules field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiff) GetNotificationRulesOk() (*[]TemplateSummaryDiffNotificationRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationRules, true
}

// SetNotificationRules sets field value
func (o *TemplateSummaryDiff) SetNotificationRules(v []TemplateSummaryDiffNotificationRule) {
	o.NotificationRules = v
}

// GetTasks returns the Tasks field value
func (o *TemplateSummaryDiff) GetTasks() []TemplateSummaryDiffTask {
	if o == nil {
		var ret []TemplateSummaryDiffTask
		return ret
	}

	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiff) GetTasksOk() (*[]TemplateSummaryDiffTask, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tasks, true
}

// SetTasks sets field value
func (o *TemplateSummaryDiff) SetTasks(v []TemplateSummaryDiffTask) {
	o.Tasks = v
}

// GetTelegrafConfigs returns the TelegrafConfigs field value
func (o *TemplateSummaryDiff) GetTelegrafConfigs() []TemplateSummaryDiffTelegraf {
	if o == nil {
		var ret []TemplateSummaryDiffTelegraf
		return ret
	}

	return o.TelegrafConfigs
}

// GetTelegrafConfigsOk returns a tuple with the TelegrafConfigs field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiff) GetTelegrafConfigsOk() (*[]TemplateSummaryDiffTelegraf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TelegrafConfigs, true
}

// SetTelegrafConfigs sets field value
func (o *TemplateSummaryDiff) SetTelegrafConfigs(v []TemplateSummaryDiffTelegraf) {
	o.TelegrafConfigs = v
}

// GetVariables returns the Variables field value
func (o *TemplateSummaryDiff) GetVariables() []TemplateSummaryDiffVariable {
	if o == nil {
		var ret []TemplateSummaryDiffVariable
		return ret
	}

	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryDiff) GetVariablesOk() (*[]TemplateSummaryDiffVariable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variables, true
}

// SetVariables sets field value
func (o *TemplateSummaryDiff) SetVariables(v []TemplateSummaryDiffVariable) {
	o.Variables = v
}

func (o TemplateSummaryDiff) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["buckets"] = o.Buckets
	}
	if true {
		toSerialize["checks"] = o.Checks
	}
	if true {
		toSerialize["dashboards"] = o.Dashboards
	}
	if true {
		toSerialize["labels"] = o.Labels
	}
	if true {
		toSerialize["labelMappings"] = o.LabelMappings
	}
	if true {
		toSerialize["notificationEndpoints"] = o.NotificationEndpoints
	}
	if true {
		toSerialize["notificationRules"] = o.NotificationRules
	}
	if true {
		toSerialize["tasks"] = o.Tasks
	}
	if true {
		toSerialize["telegrafConfigs"] = o.TelegrafConfigs
	}
	if true {
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
