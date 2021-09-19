package entities

type FilesMetadata struct {
	ID     uint64
	File   *File
	FileID uint64
	Key    string
	Value  string
}
