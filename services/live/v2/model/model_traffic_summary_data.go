/*
 * LiveAPI
 *
 * 直播服务源站所有接口
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type TrafficSummaryData struct {
	// 流量，单位为byte。
	Value *int64 `json:"value,omitempty"`
	// 域名。
	Domain *string `json:"domain,omitempty"`
}

func (o TrafficSummaryData) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"TrafficSummaryData", string(data)}, " ")
}
