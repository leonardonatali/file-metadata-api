package dto

type UpdateFilePathDto struct {
	FileID uint64
	Path   string `binding:"required"`
}
