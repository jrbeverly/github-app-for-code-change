package cmd

import (
	"log"

	"github.com/jrbeverly/github-app-for-code-change/pkg/actions"
	"github.com/jrbeverly/github-app-for-code-change/pkg/config"
	"github.com/spf13/cobra"
)

var configPath string

// Initialize the command structure
func init() {
	startCmd.Flags().StringVar(&configPath, "config", "", "Path to the configuration file")
	startCmd.MarkFlagRequired("config")
	rootCmd.AddCommand(startCmd)
}

// Start the server for processing events
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server for processing events",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.ReadV1Configuration(configPath)
		if err != nil {
			log.Fatal(err)
		}

		actions.Start(cfg)
	},
}
