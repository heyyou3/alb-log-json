package filestorage

import (
	"bytes"
	"compress/gzip"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"golang.org/x/sync/errgroup"
	"strings"
	"sync"
)

var (
	mutex = sync.Mutex{}
)

type OutputFileStorageS3Adapter struct {
	S3Client *s3.Client
	Bucket   *string
}

func (o *OutputFileStorageS3Adapter) fetchS3ObjectKeys(param FetchALBLogParam) ([]*string, error) {
	var results []*string
	res, err := o.S3Client.ListObjectsV2(param.Ctx, &s3.ListObjectsV2Input{Bucket: o.Bucket, Prefix: &param.FilePath})
	if err != nil {
		return nil, err
	}

	for _, object := range res.Contents {
		results = append(results, object.Key)
	}

	return results, nil
}

func (o *OutputFileStorageS3Adapter) getALBLog(ctx context.Context, key *string) (string, error) {
	obj, err := o.S3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: o.Bucket,
		Key:    key,
	})

	if err != nil {
		return "", err
	}
	defer obj.Body.Close()
	reader, err := gzip.NewReader(obj.Body)
	output := bytes.Buffer{}
	output.ReadFrom(reader)

	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func (o *OutputFileStorageS3Adapter) FetchALBLog(param FetchALBLogParam) ([]string, error) {
	var res []string
	objKeys, err := o.fetchS3ObjectKeys(param)

	if err != nil {
		return nil, err
	}

	eg, ctx := errgroup.WithContext(context.Background())
	// NOTE: Parallel Limit 50
	eg.SetLimit(50)
	for _, key := range objKeys {
		key := key
		eg.Go(func() error {
			log, err := o.getALBLog(ctx, key)
			if err != nil {
				return err
			}
			ll := strings.Split(log, "\n")

			mutex.Lock()
			res = append(res, ll...)
			mutex.Unlock()

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return res, nil
}

var _ OutputFileStorageAdapter = &OutputFileStorageS3Adapter{}
