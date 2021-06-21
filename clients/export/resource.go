package export

import (
	"fmt"
	"strings"
)

type ResourceType int

const (
	TypeUnset ResourceType = iota
	TypeBucket
	TypeCheck
	TypeCheckDeadman
	TypeCheckThreshold
	TypeDashboard
	TypeLabel
	TypeNotificationEndpoint
	TypeNotificationEndpointHTTP
	TypeNotificationEndpointPagerDuty
	TypeNotificationEndpointSlack
	TypeNotificationRule
	TypeTask
	TypeTelegraf
	TypeVariable
)

func (r ResourceType) String() string {
	switch r {
	case TypeBucket:
		return "bucket"
	case TypeCheck:
		return "check"
	case TypeCheckDeadman:
		return "checkDeadman"
	case TypeCheckThreshold:
		return "checkThreshold"
	case TypeDashboard:
		return "dashboard"
	case TypeLabel:
		return "label"
	case TypeNotificationEndpoint:
		return "notificationEndpoint"
	case TypeNotificationEndpointHTTP:
		return "notificationEndpointHTTP"
	case TypeNotificationEndpointPagerDuty:
		return "notificationEndpointPagerDuty"
	case TypeNotificationEndpointSlack:
		return "notificationEndpointSlack"
	case TypeNotificationRule:
		return "notificationRule"
	case TypeTask:
		return "task"
	case TypeTelegraf:
		return "telegraf"
	case TypeVariable:
		return "variable"
	case TypeUnset:
		fallthrough
	default:
		return "unset"
	}
}

func (r *ResourceType) Set(v string) error {
	switch strings.ToLower(v) {
	case "bucket":
		*r = TypeBucket
	case "check":
		*r = TypeCheck
	case "checkdeadman":
		*r = TypeCheckDeadman
	case "checkthreshold":
		*r = TypeCheckThreshold
	case "dashboard":
		*r = TypeDashboard
	case "label":
		*r = TypeLabel
	case "notificationendpoint":
		*r = TypeNotificationEndpoint
	case "notificationendpointhttp":
		*r = TypeNotificationEndpointHTTP
	case "notificationendpointpagerduty":
		*r = TypeNotificationEndpointPagerDuty
	case "notificationendpointslack":
		*r = TypeNotificationEndpointSlack
	case "notificationrule":
		*r = TypeNotificationRule
	case "task":
		*r = TypeTask
	case "telegraf":
		*r = TypeTelegraf
	case "variable":
		*r = TypeVariable
	default:
		return fmt.Errorf("unknown resource type: %s", v)
	}
	return nil
}
