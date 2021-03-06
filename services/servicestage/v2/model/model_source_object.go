/*
 * ServiceStage
 *
 * ServiceStage的API,包括应用管理和仓库授权管理
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 组件来源。
type SourceObject struct {
	Kind *SourceKind       `json:"kind,omitempty"`
	Spec *SourceOrArtifact `json:"spec,omitempty"`
}

func (o SourceObject) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"SourceObject", string(data)}, " ")
}
