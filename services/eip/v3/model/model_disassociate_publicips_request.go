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

// Request Object
type DisassociatePublicipsRequest struct {
	PublicipId string                            `json:"publicip_id"`
	Body       *DisassociatePublicipsRequestBody `json:"body,omitempty"`
}

func (o DisassociatePublicipsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DisassociatePublicipsRequest", string(data)}, " ")
}
