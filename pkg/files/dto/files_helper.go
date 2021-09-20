package dto

import (
	"mime/multipart"
	"strconv"
	"strings"
)

func GetFileMimeType(file *multipart.FileHeader) string {

	fileType := file.Header.Get("Content-Type")
	if fileType == "" {
		fileType = "unknown"
	}
	return fileType
}

func GetFilename(file *multipart.FileHeader) string {
	contentDisposition, ok := file.Header["Content-Disposition"]
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

	return file.Filename
}

func GetMetadata(file *multipart.FileHeader, path string) map[string]string {
	metadata := map[string]string{}

	metadata["filename"] = GetFilename(file)
	metadata["path"] = path
	metadata["size"] = strconv.FormatInt(file.Size, 10)
	metadata["type"] = GetFileMimeType(file)

	return metadata
}
