package alb_log

type OutputALBLogAdapter interface {
	Save(albLog []string) bool
}
