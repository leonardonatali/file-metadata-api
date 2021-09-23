package repository

import "github.com/leonardonatali/file-metadata-api/pkg/files/entities"

type FilesRepository interface {
	//Cria um arquivo
	CreateFile(file *entities.File) error
	//Retorna todos os arquivos de um usuário
	//(se o path for informado, lista apenas os arquivos de determinado subidretório)
	GetAllFiles(userID uint64, path string) ([]*entities.File, error)
	//Retorna um arquivo pelo ID e ID do Usuário
	GetFile(fileID, userID uint64) (*entities.File, error)
	//Retorna todos os metadados de um arqivo
	GetFileMetadata(fileID uint64) ([]*entities.FilesMetadata, error)
	//Atualiza as informações de um determinado arquivo bem como seus metadados
	UpdateFile(id uint64, path string, metadata []*entities.FilesMetadata) error
	//Atualiza apenas o path de um arquivo
	UpdateFilePath(fileID uint64, path string) error
	//Remove um arquivo
	DeleteFile(userID, fileID uint64) error
	//Retorna todos os metadados de um usuário
	GetAllMetadata(userID uint64) ([]*entities.FilesMetadata, error)
}
