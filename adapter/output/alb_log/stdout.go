package alb_log

import (
	"alb-log-parser/domain/alb_log_struct"
	"encoding/json"
	"fmt"
)

type OutputALBLogStdoutAdapter struct {
}

func (o *OutputALBLogStdoutAdapter) Save(albLog []*alb_log_struct.ALBLogStruct) bool {
	b, err := json.Marshal(albLog)

	if err != nil {
		return false
	}

	fmt.Println(string(b))

	return true
}

var _ OutputALBLogAdapter = &OutputALBLogStdoutAdapter{}
