package cmd

import (
	"fmt"
	"log"

	"github.com/leonardonatali/file-metadata-api/pkg/config"
	"github.com/leonardonatali/file-metadata-api/pkg/server"
	"github.com/leonardonatali/file-metadata-api/pkg/storage"
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(serve)
}

var serve = &cobra.Command{
	Use:   "serve",
	Short: "Inicia o servidor de gerenciamento de arquivos e metadados",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Config{}
		storageCfg := storage.StorageConfig{}

		if err := cfg.Load(); err != nil {
			log.Fatalf("cannot load app config: %s", err.Error())
		}

		if err := storageCfg.Load(); err != nil {
			log.Fatalf("cannot load storage config: %s", err.Error())
		}

		srv := server.NewServer(&cfg, &storageCfg)
		srv.Run()

		fmt.Println("servidor iniciando aqui...")
	},
}
