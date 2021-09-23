package entities

import (
	"fmt"
	"strings"
	"time"
)

type File struct {
	ID        uint64
	UserID    uint64
	Name      string
	Path      string
	Metadata  []*FilesMetadata
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (f *File) GetMetadataByKey(key string) *FilesMetadata {
	for _, m := range f.Metadata {
		if m.Key == key {
			return m
		}
	}

	return nil
}

// GetQualifiedName - Retorna o nome do arquivo para upload/download no storage
func (f *File) GetQualifiedName() string {

	path := strings.TrimPrefix(f.Path, "/")
	path = strings.TrimSuffix(path, "/")

	return fmt.Sprintf("/%d/%s/%d_%s", f.UserID, path, f.ID, f.Name)
}
