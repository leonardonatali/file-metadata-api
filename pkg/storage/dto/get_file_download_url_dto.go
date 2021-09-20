package dto

type GetFileDownloadURLDto struct {
	UserID   uint64
	Path     string
	Filename string
	Type     string
}

type GetFileDownloadURLResponseDto struct {
	DownloadURL string
}
