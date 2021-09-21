package storage

import "github.com/kelseyhightower/envconfig"

type StorageConfig struct {
	BucketName      string `envconfig:"storage_bucket_name"`
	BucketLocation  string `envconfig:"storage_bucket_location"`
	Endpoint        string `envconfig:"storage_endpoint"`
	AccessKeyId     string `envconfig:"storage_access_key_id"`
	SecretAccessKey string `envconfig:"storage_secret_access_key"`
	UseSSL          bool   `envconfig:"storage_use_ssl"`
}

func (c *StorageConfig) Load() error {
	return envconfig.Process("", c)
}
