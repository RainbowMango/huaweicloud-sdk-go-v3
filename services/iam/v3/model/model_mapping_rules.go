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
type MappingRules struct {
	// 表示联邦用户在本系统中的用户信息。 user：联邦用户在本系统中的用户名称。group：联邦用户在本系统中所属用户组。
	Local []map[string]RulesLocalAdditional `json:"local"`
	// 表示联邦用户在IdP中的用户信息。由断言属性及运算符组成的表达式，取值由断言决定。
	Remote []RulesRemote `json:"remote"`
}

func (o MappingRules) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"MappingRules", string(data)}, " ")
}