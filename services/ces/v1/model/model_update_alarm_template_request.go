/*
 * CES
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateAlarmTemplateRequest struct {
	ContentType string                          `json:"Content-Type"`
	TemplateId  string                          `json:"template_id"`
	Body        *UpdateAlarmTemplateRequestBody `json:"body,omitempty"`
}

func (o UpdateAlarmTemplateRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateAlarmTemplateRequest", string(data)}, " ")
}
