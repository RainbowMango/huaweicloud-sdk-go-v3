/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListSecurityGroupRulesRequest struct {
	Marker          *string `json:"marker,omitempty"`
	Limit           *int32  `json:"limit,omitempty"`
	SecurityGroupId *string `json:"security_group_id,omitempty"`
}

func (o ListSecurityGroupRulesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListSecurityGroupRulesRequest", string(data)}, " ")
}
