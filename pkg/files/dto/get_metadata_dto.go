package dto

type GetMetadataDto struct {
	FileID uint64 `binding:"required"`
}
