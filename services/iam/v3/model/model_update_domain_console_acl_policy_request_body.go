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
type UpdateDomainConsoleAclPolicyRequestBody struct {
	ConsoleAclPolicy *AclPolicyOption `json:"console_acl_policy"`
}

func (o UpdateDomainConsoleAclPolicyRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateDomainConsoleAclPolicyRequestBody", string(data)}, " ")
}
