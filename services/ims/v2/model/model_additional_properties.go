/*
 * IMS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 属性值
type AdditionalProperties struct {
	// 类型
	Type string `json:"type"`
}

func (o AdditionalProperties) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AdditionalProperties", string(data)}, " ")
}
