package transfers

import (
	"io"

	"github.com/leonardonatali/file-metadata-api/pkg/config"
	"github.com/leonardonatali/file-metadata-api/pkg/transfers/repository"
)

type TransferService struct {
	TransfersRepository repository.TranfersRepository
	Cfg                 *config.Config
}

func NewTranfersService(tranfersRepository repository.TranfersRepository, cfg *config.Config) *TransferService {
	return &TransferService{
		TransfersRepository: tranfersRepository,
		Cfg:                 cfg,
	}
}

func (s TransferService) UploadFile(path string, content io.Reader) error {
	return s.TransfersRepository.UploadFile(path, content)
}

func (s TransferService) DownloadFile(path string) error {
	return s.TransfersRepository.DownloadFile(path)
}

func (s TransferService) DeleteFile(path string) error {
	return s.TransfersRepository.DeleteFile(path)
}

func (s TransferService) ReplaceFile(path string, content io.Reader) error {
	return s.TransfersRepository.UploadFile(path, content)
}
