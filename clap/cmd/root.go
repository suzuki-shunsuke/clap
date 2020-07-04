package cmd

import (
	"github.com/spf13/cobra"
	"github.com/suzuki-shunsuke/clap/clap/constant"
)

var rootCmd = &cobra.Command{ //nolint:gochecknoglobals
	Use:     "clap",
	Short:   "clap is a simple installer",
	Long:    `clap is a simple installer. https://github.com/suzuki-shunsuke/clap`,
	Version: constant.Version,
}

func Execute() error {
	rootCmd.AddCommand(installCmd, versionCmd)
	return rootCmd.Execute()
}
