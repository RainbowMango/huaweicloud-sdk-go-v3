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

// quota详细信息
type QuotaInfo struct {
	// 查询的资源类型。
	Type string `json:"type"`
	// 已使用的资源配额。
	Used int32 `json:"used"`
	// 查询出的资源的总配额。
	Quota int32 `json:"quota"`
	// 资源的最小配额。
	Min int32 `json:"min"`
	// 资源的最大配额。
	Max int32 `json:"max"`
}

func (o QuotaInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"QuotaInfo", string(data)}, " ")
}
