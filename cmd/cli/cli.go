package main

import (
	"os"

	"github.com/bilalcaliskan/split-the-tunnel/cmd/cli/add"
	"github.com/bilalcaliskan/split-the-tunnel/cmd/cli/list"
	"github.com/bilalcaliskan/split-the-tunnel/cmd/cli/remove"
	"github.com/bilalcaliskan/split-the-tunnel/cmd/cli/utils"
	"github.com/bilalcaliskan/split-the-tunnel/internal/version"

	"github.com/spf13/cobra"
)

var (
	ver    = version.Get()
	cliCmd = &cobra.Command{
		Use:     "stt-cli",
		Short:   "",
		Long:    ``,
		Version: ver.GitVersion,
	}
)

func main() {
	if err := cliCmd.Execute(); err != nil {
		// extract the response code from the error
		var resCode int
		if cmdErr, ok := err.(*utils.CommandError); ok {
			resCode = cmdErr.Code
		}

		if resCode != 0 {
			os.Exit(resCode)
		}

		os.Exit(1)
	}
}

func init() {
	cliCmd.AddCommand(add.AddCmd)
	cliCmd.AddCommand(list.ListCmd)
	cliCmd.AddCommand(remove.RemoveCmd)
}
