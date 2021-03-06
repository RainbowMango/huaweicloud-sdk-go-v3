/*
 * devcloudpipeline
 *
 * devcloud pipeline api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type StartPipelineParameters struct {
	// 启动流水线时的构建参数
	BuildParams *[]StartPipelineBuildParams `json:"build_params,omitempty"`
}

func (o StartPipelineParameters) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"StartPipelineParameters", string(data)}, " ")
}
