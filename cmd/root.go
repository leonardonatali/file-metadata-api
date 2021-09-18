package cmd

import (
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use: "run",
}

// Execute executes the root command.
func Execute() error {
	return root.Execute()
}
