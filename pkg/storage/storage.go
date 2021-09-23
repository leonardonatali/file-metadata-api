package storage

import (
	"io"
	"time"
)

type StorageService interface {
	// Carrega as configurações
	Load(cfg *StorageConfig) error
	// Verifica se o bucket da configuração existe
	BucketExists() (bool, error)
	// Cria um bucket com o mesmo nome do bucket da configuração
	CreateBucket() error
	// Adiciona o arquivo ao bucket da configuração
	PutFile(content io.Reader, path, mimeType string, size uint64) error
	// Remove um arquivo
	DeleteFile(path string) error
	// Obtém o Link de download
	GetDownloadURL(path, filename, mimeType string, expires time.Duration) (string, error)
	// Move um arquivo de um lugar para o outro
	Move(src, dest string) error
}
