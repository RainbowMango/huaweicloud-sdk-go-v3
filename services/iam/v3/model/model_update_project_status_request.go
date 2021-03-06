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

// Request Object
type UpdateProjectStatusRequest struct {
	ProjectId string                          `json:"project_id"`
	Body      *UpdateProjectStatusRequestBody `json:"body,omitempty"`
}

func (o UpdateProjectStatusRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateProjectStatusRequest", string(data)}, " ")
}
