package gcs

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/logger"
)

type Service struct {
	config *config.Config
}

func New(cfg *config.Config) IService {
	return &Service{
		config: cfg,
	}
}

func (s Service) UploadObject(object, filePath string) (string, error) {
	ctx := context.Background()
	decoded, _ := base64.StdEncoding.DecodeString(s.config.GoogleServiceAccountKey)
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON(decoded))
	if err != nil {
		logger.L.Errorf(err, "[gcs.UploadObject] storage.NewClient() failed: %s\n", err.Error())
		return "", err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	o := client.Bucket(s.config.GoogleNftImageBucket).Object(object)
	// do nothing if object already existed
	o = o.If(storage.Conditions{DoesNotExist: true})

	// open file
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// upload object
	wc := o.NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		logger.L.Errorf(err, "[gcs.UploadObject] io.Copy() failed: %s\n", err.Error())
		return "", err
	}
	if err := wc.Close(); err != nil {
		logger.L.Errorf(err, "[gcs.UploadObject] writer.Close() failed: %s\n", err.Error())
		return "", err
	}
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", s.config.GoogleNftImageBucket, object), nil
}
