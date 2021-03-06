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
type KeystoneUpdateUserPasswordRequest struct {
	UserId string                                 `json:"user_id"`
	Body   *KeystoneUpdateUserPasswordRequestBody `json:"body,omitempty"`
}

func (o KeystoneUpdateUserPasswordRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneUpdateUserPasswordRequest", string(data)}, " ")
}
