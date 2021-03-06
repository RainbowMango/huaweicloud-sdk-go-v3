/*
 * cloudide
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
type StartInstanceResponse struct {
	// 返回值
	Result *string `json:"result,omitempty"`
	// 状态
	Status *string `json:"status,omitempty"`
}

func (o StartInstanceResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"StartInstanceResponse", string(data)}, " ")
}
