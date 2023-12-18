package filestorage

import "context"

type FetchALBLogParam struct {
	Ctx      context.Context
	FilePath string
}

type OutputFileStorageAdapter interface {
	FetchALBLog(param FetchALBLogParam) ([]string, error)
}
