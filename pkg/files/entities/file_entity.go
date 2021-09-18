package entities

import "time"

type File struct {
	ID          uint64
	UserID      uint64
	Name        string
	Path        string
	Extension   string
	SizeInBytes uint64
	Metadata    []*FileMetadata
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
