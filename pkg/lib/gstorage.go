package lib

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
	"github.com/rs/zerolog/log"
)

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
}

var GStorageUploader *ClientUploader

func GStorageInit(projectID string, bucketName string) {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		panic("Failed to create client")
	}

	log.Info().Msg("gstorage connected")
	GStorageUploader = &ClientUploader{
		cl:         client,
		projectID:  projectID,
		bucketName: bucketName,
	}
}

func (cu *ClientUploader) UploadFile(ctx context.Context, fileName string, r io.Reader) error {
	bkt := cu.cl.Bucket(cu.bucketName)
	obj := bkt.Object(fileName)
	wc := obj.NewWriter(ctx)

	if _, err := io.Copy(wc, r); err != nil {
		return err
	}

	if err := wc.Close(); err != nil {
		return err
	}

	return nil
}

func (cu *ClientUploader) DeleteFile(ctx context.Context, fileName string) error {
	bkt := cu.cl.Bucket(cu.bucketName)
	obj := bkt.Object(fileName)
	if err := obj.Delete(ctx); err != nil {
		return err
	}
	return nil
}
