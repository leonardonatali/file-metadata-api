package dto

type CreateFileDto struct {
	Token    string
	Path     string
	Filename string
	Metadata map[string]string
}
