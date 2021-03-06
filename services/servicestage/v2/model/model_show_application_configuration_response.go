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
type ShowApplicationConfigurationResponse struct {
	// 应用配置列表。
	Configuration *[]ApplicationListConfigConfiguration1 `json:"configuration,omitempty"`
}

func (o ShowApplicationConfigurationResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowApplicationConfigurationResponse", string(data)}, " ")
}
