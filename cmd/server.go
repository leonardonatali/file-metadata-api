package cmd

import (
	"fmt"
	"log"

	"github.com/leonardonatali/file-metadata-api/pkg/config"
	"github.com/leonardonatali/file-metadata-api/pkg/database/migrations"
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(serve)
}

var serve = &cobra.Command{
	Use:   "serve",
	Short: "Inicia o servidor de gerenciamento de arquivos e metadados",
	Run: func(cmd *cobra.Command, args []string) {
		var cfg config.Config

		if err := cfg.Load(); err != nil {
			log.Fatalf("cannot load app config: %s", err.Error())
		}

		if err := migrations.Migrate(&cfg); err != nil {
			log.Fatalf("cannot run database migration: %s", err.Error())
		}

		fmt.Println("servidor iniciando aqui...")
	},
}
