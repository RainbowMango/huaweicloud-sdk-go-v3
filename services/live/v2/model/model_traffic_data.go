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

type TrafficData struct {
	// 采样值，单位为byte。
	Value *int32 `json:"value,omitempty"`
	// 采样时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。
	Time *string `json:"time,omitempty"`
}

func (o TrafficData) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"TrafficData", string(data)}, " ")
}
