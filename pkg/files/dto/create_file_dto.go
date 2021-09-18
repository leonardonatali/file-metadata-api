package dto

type CreateFileDto struct {
	UserID      uint64
	Path        string
	Name        string
	Extension   string
	SizeInBytes uint64
	Metadata    map[string]string
}
