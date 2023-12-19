package fetch_alb_log

type IFetchALBLogInput interface {
	Invoke() ([]*FetchALBLogStruct, error)
}

type FetchALBLogInput struct {
	Usecase IFetchALBLogUsecase
}

func (i *FetchALBLogInput) Invoke() ([]*FetchALBLogStruct, error) {
	// TODO: user input value validation
	return i.Usecase.Invoke()
}
