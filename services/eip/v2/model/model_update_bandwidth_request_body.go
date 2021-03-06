/*
 * EIP
 *
 * 云服务接口
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 更新带宽对象的请求体(name,size必须有一个参数)
type UpdateBandwidthRequestBody struct {
	Bandwidth *UpdateBandwidthOption `json:"bandwidth"`
}

func (o UpdateBandwidthRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateBandwidthRequestBody", string(data)}, " ")
}
