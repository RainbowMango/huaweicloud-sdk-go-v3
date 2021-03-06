/*
 * IAM
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
type KeystoneListMappingsResponse struct {
	Links *Links `json:"links,omitempty"`
	// 映射信息列表。
	Mappings *[]MappingResult `json:"mappings,omitempty"`
}

func (o KeystoneListMappingsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneListMappingsResponse", string(data)}, " ")
}
