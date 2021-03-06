/*
 * ELB
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListHealthmonitorsResponse struct {
	// 健康检查对象的列表
	Healthmonitors *[]HealthmonitorResp `json:"healthmonitors,omitempty"`
}

func (o ListHealthmonitorsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListHealthmonitorsResponse", string(data)}, " ")
}
