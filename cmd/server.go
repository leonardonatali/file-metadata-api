package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(serve)
}

var serve = &cobra.Command{
	Use:   "serve",
	Short: "Inicia o servidor de gerenciamento de arquivos e metadados",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("servidor iniciando aqui...")
	},
}
