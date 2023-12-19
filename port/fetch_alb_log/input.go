package fetch_alb_log

import "alb-log-json/domain/alb_log_struct"

type IFetchALBLogInput interface {
	Invoke() ([]*alb_log_struct.ALBLogStruct, error)
}

type FetchALBLogInput struct {
	Usecase IFetchALBLogUsecase
}

func (i *FetchALBLogInput) Invoke() ([]*alb_log_struct.ALBLogStruct, error) {
	// TODO: user input value validation
	return i.Usecase.Invoke()
}
