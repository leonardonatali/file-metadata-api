package s3

import (
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/leonardonatali/file-metadata-api/pkg/storage"
)

type S3Service struct {
	client *s3.S3
	cfg    *storage.StorageConfig
}

func (m *S3Service) Load(cfg *storage.StorageConfig) error {
	m.cfg = cfg

	sess, err := session.NewSession(&aws.Config{
		Endpoint:         aws.String(m.cfg.Endpoint),
		DisableSSL:       aws.Bool(!m.cfg.UseSSL),
		Region:           aws.String(m.cfg.BucketRegion),
		Credentials:      credentials.NewStaticCredentials(m.cfg.AccessKeyId, m.cfg.AccessKeyId, m.cfg.SecretAccessKey),
		S3ForcePathStyle: aws.Bool(true),
	})

	if err != nil {
		return err
	}

	m.client = s3.New(sess)
	return nil
}

func (m *S3Service) BucketExists() (bool, error) {
	res, err := m.client.ListBuckets(nil)
	if err != nil {
		return false, err
	}

	for _, bucket := range res.Buckets {
		if bucket.Name != nil && *bucket.Name == m.cfg.BucketName {
			return true, nil
		}
	}

	return false, nil
}

func (m *S3Service) CreateBucket() error {
	_, err := m.client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(m.cfg.BucketName),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String(m.cfg.BucketRegion),
		},
	})

	return err
}

func (m *S3Service) PutFile(content io.Reader, path, mimeType string, size uint64) error {
	_, err := m.client.PutObject(
		&s3.PutObjectInput{
			Bucket:             aws.String(m.cfg.BucketName),
			Body:               aws.ReadSeekCloser(content),
			ContentDisposition: aws.String(mimeType),
			ContentLength:      aws.Int64(int64(size)),
			Key:                aws.String(path),
		})

	if err != nil {
		return err
	}

	return nil
}

func (m *S3Service) DeleteFile(path string) error {
	_, err := m.client.DeleteObject(&s3.DeleteObjectInput{
		Bucket:                    &m.cfg.BucketName,
		BypassGovernanceRetention: aws.Bool(true),
		Key:                       aws.String(path),
	})

	return err
}

func (m *S3Service) GetDownloadURL(path, filename, mimeType string, expires time.Duration) (string, error) {
	res, err := m.client.GetObject(&s3.GetObjectInput{
		Bucket:                     aws.String(m.cfg.BucketName),
		Key:                        &filename,
		ResponseContentDisposition: aws.String(fmt.Sprintf(`attachment; filename="%s"`, filename)),
		ResponseContentType:        aws.String(mimeType),
		ResponseExpires:            aws.Time(time.Now().Add(time.Hour)),
	})

	if err != nil {
		return "", err
	}

	return res.String(), nil
}
