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

//
type ScopeDomainOption struct {
	// 账号ID，id与name二选一即可。
	Id *string `json:"id,omitempty"`
	// 账号名，id与name二选一即可。
	Name *string `json:"name,omitempty"`
}

func (o ScopeDomainOption) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ScopeDomainOption", string(data)}, " ")
}
