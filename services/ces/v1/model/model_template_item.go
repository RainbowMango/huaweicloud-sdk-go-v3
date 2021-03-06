/*
 * CES
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 创建自定义告警模板添加的告警规则。
type TemplateItem struct {
	// 告警模板添加的监控指标，如弹性云服务器可添加的监控指标为cpu_util等，各资源的监控指标名称可查看https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html。
	MetricName string                  `json:"metric_name"`
	Condition  *AlarmTemplateCondition `json:"condition"`
	// 设置告警级别，值为1,2,3,4；1为紧急，2为重要，3为次要，4为提示。
	AlarmLevel *int32 `json:"alarm_level,omitempty"`
}

func (o TemplateItem) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"TemplateItem", string(data)}, " ")
}
