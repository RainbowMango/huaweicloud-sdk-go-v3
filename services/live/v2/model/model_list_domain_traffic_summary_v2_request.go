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

// Request Object
type ListDomainTrafficSummaryV2Request struct {
	PlayDomains []string  `json:"play_domains"`
	App         *string   `json:"app,omitempty"`
	Stream      *string   `json:"stream,omitempty"`
	Region      *[]string `json:"region,omitempty"`
	Isp         *[]string `json:"isp,omitempty"`
	StartTime   *string   `json:"start_time,omitempty"`
	EndTime     *string   `json:"end_time,omitempty"`
}

func (o ListDomainTrafficSummaryV2Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListDomainTrafficSummaryV2Request", string(data)}, " ")
}
