/*
 * ServiceStage
 *
 * ServiceStage的API,包括应用管理和仓库授权管理
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreatePersonalAuthResponse struct {
	Authorization *AuthorizationVo `json:"authorization,omitempty"`
}

func (o CreatePersonalAuthResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePersonalAuthResponse", string(data)}, " ")
}
