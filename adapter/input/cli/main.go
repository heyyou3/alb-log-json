package main

import (
	"alb-log-parser/adapter/output/alb_log"
	"alb-log-parser/adapter/output/filestorage"
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
)

func run() int {
	ctx := context.Background()
	cfg, _ := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile(""))
	bucketName := ""

	fileStorageAdapter := filestorage.OutputFileStorageS3Adapter{
		S3Client: s3.NewFromConfig(cfg),
		Bucket:   &bucketName,
	}

	albLogs, err := fileStorageAdapter.FetchALBLog(filestorage.FetchALBLogParam{FilePath: "", Ctx: ctx})

	if err != nil {
		return 1
	}

	stdoutAdapter := alb_log.OutputALBLogStdoutAdapter{}
	stdoutAdapter.Save(albLogs)
	return 0
}

func main() {
	os.Exit(run())
}
