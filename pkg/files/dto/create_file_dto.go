package dto

import (
	"mime/multipart"
)

const mimeHeaderPrefix = "name"

type CreateFileDto struct {
	UserID   uint64
	Name     string
	Path     string                `binding:"required"`
	File     *multipart.FileHeader `binding:"required"`
	Metadata map[string]string
}
