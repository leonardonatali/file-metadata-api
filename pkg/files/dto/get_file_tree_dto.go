package dto

type GetFileTreeDto struct {
	CurrentDir string           `json:"CurrentDir"`
	Children   []GetFileTreeDto `json:"Children,omitempty"`
}
