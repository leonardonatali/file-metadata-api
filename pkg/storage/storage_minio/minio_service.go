package storage_minio

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/leonardonatali/file-metadata-api/pkg/storage"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioService struct {
	client *minio.Client
	cfg    *storage.StorageConfig
}

func (m *MinioService) Load(cfg *storage.StorageConfig) error {
	m.cfg = cfg

	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyId, cfg.SecretAccessKey, ""),
		Secure: cfg.UseSSL,
	})

	if err != nil {
		return err
	}

	m.client = client
	return nil
}

func (m *MinioService) BucketExists() (bool, error) {
	return m.client.BucketExists(context.Background(), m.cfg.BucketName)
}

func (m *MinioService) CreateBucket() error {
	return m.client.MakeBucket(
		context.Background(),
		m.cfg.BucketName,
		minio.MakeBucketOptions{Region: m.cfg.BucketLocation,
			ObjectLocking: true,
		},
	)
}

func (m *MinioService) PutFile(content io.Reader, path, mimeType string, size uint64) error {
	_, err := m.client.PutObject(
		context.Background(),
		m.cfg.BucketName,
		path,
		content,
		int64(size),
		minio.PutObjectOptions{
			ContentType:        mimeType,
			ContentDisposition: mimeType,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *MinioService) DeleteFile(path string) error {
	opts := minio.RemoveObjectOptions{GovernanceBypass: true}
	return m.client.RemoveObject(context.Background(), m.cfg.BucketName, path, opts)
}

func (m *MinioService) GetDownloadURL(path, filename, mimeType string, expires time.Duration) (string, error) {
	params := url.Values{}
	params.Set("response-content-disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
	params.Set("content-type", mimeType)

	res, err := m.client.PresignedGetObject(context.Background(), m.cfg.BucketName, path, expires, params)
	if err != nil {
		return "", err
	}

	return res.Query().Encode(), nil
}
