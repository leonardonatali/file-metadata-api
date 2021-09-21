package storage

type StorageStrategy interface {
	Load(cfg *StorageConfig)
	BucketExists() (bool, error)
	CreateBucket(name string) error
	PutFile()
}
