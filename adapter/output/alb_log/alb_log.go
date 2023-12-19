package alb_log

import "alb-log-json/domain/alb_log_struct"

type OutputALBLogAdapter interface {
	Save(albLog []*alb_log_struct.ALBLogStruct) bool
}
