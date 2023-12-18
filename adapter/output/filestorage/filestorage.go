package filestorage

type FetchALBLogParam struct {
}

type OutputFileStorageAdapter interface {
	FetchALBLog(param FetchALBLogParam) []string
}
