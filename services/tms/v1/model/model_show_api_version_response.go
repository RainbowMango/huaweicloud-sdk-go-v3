/*
 * TMS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowApiVersionResponse struct {
	Version *VersionDetail `json:"version,omitempty"`
}

func (o ShowApiVersionResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowApiVersionResponse", string(data)}, " ")
}
