package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// The root command that all commands will be subcommands of
var rootCmd = &cobra.Command{
	Use: tool,
	Run: func(cmd *cobra.Command, args []string) {},
}

// Execute the command processing of the application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
