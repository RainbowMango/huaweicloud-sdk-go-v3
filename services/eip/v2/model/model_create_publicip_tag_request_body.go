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

// 创建tag对象的请求体
type CreatePublicipTagRequestBody struct {
	Tag *ResourceTagOption `json:"tag"`
}

func (o CreatePublicipTagRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePublicipTagRequestBody", string(data)}, " ")
}
