package dto

import "mime/multipart"

type UpdateFileDto struct {
	UserID    uint64
	OldFileID uint64
	Path      string                `binding:"required"`
	File      *multipart.FileHeader `binding:"required"`
	Metadata  map[string]string
}
