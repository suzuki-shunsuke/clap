package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{ //nolint:gochecknoglobals
	Use:   "clap",
	Short: "clap is a simple installer",
	Long:  `clap is a simple installer. https://github.com/suzuki-shunsuke/clap`,
}

func Execute() error {
	rootCmd.AddCommand(installCmd)
	return rootCmd.Execute()
}
