/*
 * ELB
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
type ShowListenerRequest struct {
	ListenerId string `json:"listener_id"`
}

func (o ShowListenerRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowListenerRequest", string(data)}, " ")
}
