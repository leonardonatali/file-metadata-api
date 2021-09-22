package storage

import (
	"github.com/kelseyhightower/envconfig"
)

type StorageConfig struct {
	BucketName      string `envconfig:"STORAGE_BUCKET_NAME"`
	BucketRegion    string `envconfig:"STORAGE_BUCKET_REGION"`
	Endpoint        string `envconfig:"STORAGE_ENDPOINT"`
	AccessKeyId     string `envconfig:"STORAGE_ACCESS_KEY_ID"`
	SecretAccessKey string `envconfig:"STORAGE_SECRET_ACCESS_KEY"`
	UseSSL          bool   `envconfig:"STORAGE_USE_SSL"`
}

func (c *StorageConfig) Load() error {
	return envconfig.Process("", c)
}
