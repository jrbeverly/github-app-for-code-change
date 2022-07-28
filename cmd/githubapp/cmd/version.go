package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version associated with the compiled application
var version string

// Git Commit associated with the compiled application
var commit string

// Name of the tool as shown in the help manual
var tool string

// Initialize the command structure
func init() {
	rootCmd.AddCommand(versionCmd)
}

// Version convention for the application when running 'version'
func versionConvention(tool string, version string, commit string) string {
	if version == "" {
		if commit == "" {
			return fmt.Sprintf("%s/nover git/nostamp", tool)
		} else {
			return fmt.Sprintf("%s/nover git/%s", tool, commit)
		}
	}

	if commit == "" {
		return fmt.Sprintf("%s/%s git/nostamp", tool, version)
	}
	return fmt.Sprintf("%s/%s git/%s", tool, version, commit)
}

// Emit the version of the application
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `All software has versions. This is mine.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(versionConvention(tool, version, commit))
	},
}
