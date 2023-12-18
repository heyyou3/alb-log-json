package alb_log

import "fmt"

type OutputALBLogStdoutAdapter struct {
}

func (o *OutputALBLogStdoutAdapter) Save(albLog []string) bool {
	for _, log := range albLog {
		fmt.Println(log)
		return true
	}

	return false
}

var _ OutputALBLogAdapter = &OutputALBLogStdoutAdapter{}
