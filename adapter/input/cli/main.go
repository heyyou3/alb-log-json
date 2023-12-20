package main

import (
	"alb-log-json/adapter/output/alb_log"
	"alb-log-json/adapter/output/filestorage"
	"alb-log-json/port/fetch_alb_log"
	"context"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
)

var (
	cliVersion string
)

type config struct {
	AWSProfileName string `toml:"aws_profile_name"`
	BucketName     string `toml:"bucket_name"`
	S3Key          string `toml:"s3_key"`
}

func loadConfig() (*config, error) {
	var conf config
	_, err := toml.DecodeFile("./alblogjson-config.toml", &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func run(conf *config) int {
	ctx := context.Background()
	cfg, _ := awsconfig.LoadDefaultConfig(ctx, awsconfig.WithSharedConfigProfile(conf.AWSProfileName))

	fileStorageAdapter := &filestorage.OutputFileStorageS3Adapter{
		S3Client: s3.NewFromConfig(cfg),
		Bucket:   &conf.BucketName,
	}
	input := fetch_alb_log.FetchALBLogInput{Usecase: &fetch_alb_log.FetchALBLogUsecase{FetchALBLogParam: filestorage.FetchALBLogParam{
		Ctx:      ctx,
		FilePath: conf.S3Key,
	}, Output: &fetch_alb_log.FetchALBLogOutput{FileStorageAdapter: fileStorageAdapter}}}
	albLogs, err := input.Invoke()

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return 1
	}

	stdoutAdapter := alb_log.OutputALBLogStdoutAdapter{}
	stdoutAdapter.Save(albLogs)
	return 0
}

func main() {
	var version bool
	flag.BoolVar(&version, "version", false, "show version")
	flag.Parse()

	if version {
		fmt.Println(cliVersion)
		os.Exit(0)
	}

	conf, err := loadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed load toml config.\nCreate `alblogjson-config.toml`.\n")
		os.Exit(1)
	}
	os.Exit(run(conf))
}
