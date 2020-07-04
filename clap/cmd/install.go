package cmd

import (
	"github.com/spf13/cobra"
	"github.com/suzuki-shunsuke/clap/clap/install/cli"
	"github.com/suzuki-shunsuke/clap/clap/install/controller"
)

var installCmd = &cobra.Command{ //nolint:gochecknoglobals
	Use:   "install",
	Short: "install binary",
	Long:  `install binary`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// URL <file path in archive>:<install path> ...
		input, err := cli.GetInput(cmd, args)
		if err != nil {
			return err
		}
		paramsNew, err := cli.ParseInput(input)
		if err != nil {
			return err
		}
		ctrl, params, err := controller.New(paramsNew)
		if err != nil {
			return err
		}
		return ctrl.Run(cmd.Context(), params)
	},
}
