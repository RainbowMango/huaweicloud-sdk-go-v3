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
type DeletePermanentAccessKeyRequest struct {
	AccessKey string `json:"access_key"`
}

func (o DeletePermanentAccessKeyRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeletePermanentAccessKeyRequest", string(data)}, " ")
}
