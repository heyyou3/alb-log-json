package fetch_alb_log

import "alb-log-json/adapter/output/filestorage"

type IFetchALBLogOutput interface {
	FetchALBLog(param filestorage.FetchALBLogParam) ([]string, error)
}
type FetchALBLogOutput struct {
	FileStorageAdapter filestorage.OutputFileStorageAdapter
}

func (o *FetchALBLogOutput) FetchALBLog(param filestorage.FetchALBLogParam) ([]string, error) {
	return o.FileStorageAdapter.FetchALBLog(param)
}
