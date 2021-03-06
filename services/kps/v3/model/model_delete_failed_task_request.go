/*
 * kps
 *
 * kps v3 版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteFailedTaskRequest struct {
	TaskId string `json:"task_id"`
}

func (o DeleteFailedTaskRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteFailedTaskRequest", string(data)}, " ")
}
