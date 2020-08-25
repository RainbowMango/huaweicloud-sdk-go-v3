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

// 弹性ip，同publicips
type EipInfo struct {
	// eip_id
	EipId *string `json:"eip_id,omitempty"`
	// eip_address
	EipAddress *string `json:"eip_address,omitempty"`
}

func (o EipInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"EipInfo", string(data)}, " ")
}