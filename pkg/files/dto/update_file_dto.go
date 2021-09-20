package dto

import "mime/multipart"

type UpdateFileDto struct {
	UserID    uint64
	OldFileID uint64
	File      *multipart.FileHeader `binding:"required"`
	Metadata  map[string]string
}
