package repository

import "io"

type TranfersRepository interface {
	UploadFile(path string, content io.Reader) error
	DownloadFile(path string) error
	DeleteFile(path string) error
}
