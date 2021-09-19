package entities

import (
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
