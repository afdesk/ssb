// Package commands /*
package commands

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ssb",
	Short: "The Static Site Builder",
}

func NewApp(version string) *cobra.Command {
	rootCmd.Version = version
	return rootCmd
}

func init() {}
