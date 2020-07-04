package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/suzuki-shunsuke/clap/clap/constant"
)

var versionCmd = &cobra.Command{ //nolint:gochecknoglobals
	Use:   "version",
	Short: "show clap's version",
	Long:  "show clap's version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(constant.Version)
	},
}
