package dto

type ReplaceFileDto struct {
	UserID      uint64
	OldFileID   uint64
	Path        string
	Name        string
	Extension   string
	SizeInBytes uint64
	Metadata    map[string]string
}
