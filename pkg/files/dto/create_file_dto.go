package dto

import (
	"mime/multipart"
	"strconv"
	"strings"
)

const mimeHeaderPrefix = "name"

type CreateFileDto struct {
	UserID   uint64
	Name     string
	Path     string                `form:"path" binding:"required"`
	File     *multipart.FileHeader `form:"file" binding:"required"`
	Metadata map[string]string
}

func (dto *CreateFileDto) GetFileMimeType() string {

	fileType := dto.File.Header.Get("Content-Type")
	if fileType == "" {
		fileType = "unknown"
	}
	return fileType
}

func (dto *CreateFileDto) GetFilename() string {
	contentDisposition, ok := dto.File.Header["Content-Disposition"]
	if ok {
		for _, value := range contentDisposition {
			if strings.HasPrefix(value, mimeHeaderPrefix) {
				filename := strings.ReplaceAll(
					strings.TrimPrefix(mimeHeaderPrefix, value),
					`"`,
					"",
				)

				if filename != "" {
					return filename
				}
			}
		}
	}

	return dto.File.Filename
}

func (dto *CreateFileDto) LoadMetadata() {
	dto.Metadata = map[string]string{}

	dto.Metadata["filename"] = dto.GetFilename()
	dto.Metadata["path"] = dto.Path
	dto.Metadata["size"] = strconv.FormatInt(dto.File.Size, 10)
	dto.Metadata["type"] = dto.GetFileMimeType()

}
