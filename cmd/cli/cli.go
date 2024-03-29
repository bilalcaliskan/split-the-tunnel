package main

import (
	"context"
	"os"
	"path/filepath"

	"github.com/bilalcaliskan/split-the-tunnel/cmd/cli/purge"

	"github.com/bilalcaliskan/split-the-tunnel/internal/constants"
	"github.com/bilalcaliskan/split-the-tunnel/internal/logging"
	"github.com/pkg/errors"

	"github.com/bilalcaliskan/split-the-tunnel/cmd/cli/add"
	"github.com/bilalcaliskan/split-the-tunnel/cmd/cli/list"
	"github.com/bilalcaliskan/split-the-tunnel/cmd/cli/remove"
	"github.com/bilalcaliskan/split-the-tunnel/cmd/cli/utils"
	"github.com/bilalcaliskan/split-the-tunnel/internal/version"

	"github.com/spf13/cobra"
)

var (
	verbose   bool
	workspace string
	ver       = version.Get()
	cliCmd    = &cobra.Command{
		Use:     "stt-cli",
		Short:   "",
		Long:    ``,
		Version: ver.GitVersion,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLogger()
			logger.Info().Str("appVersion", ver.GitVersion).Str("goVersion", ver.GoVersion).Str("goOS", ver.GoOs).
				Str("goArch", ver.GoArch).Str("gitCommit", ver.GitCommit).Str("buildDate", ver.BuildDate).
				Msg("split-the-tunnel cli is started!")

			if verbose {
				logger = logging.WithVerbose()
				logger.Debug().Str("foo", "bar").Msg("this is a dummy log")
			}

			cmd.SetContext(context.WithValue(cmd.Context(), constants.LoggerKey{}, logger))
			cmd.SetContext(context.WithValue(cmd.Context(), constants.SocketPathKey{}, filepath.Join(workspace, constants.SocketFileName)))
		},
	}
)

func main() {
	if err := cliCmd.Execute(); err != nil {
		var cmdErr *utils.CommandError
		if errors.As(err, &cmdErr) {
			os.Exit(cmdErr.Code)
		}
		os.Exit(1)
	}
}

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(errors.Wrap(err, "failed to get user home directory"))
	}

	cliCmd.PersistentFlags().StringVarP(&workspace, "workspace", "w", filepath.Join(homeDir, ".split-the-tunnel"), "workspace directory path")
	cliCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose mode")

	cliCmd.AddCommand(add.AddCmd)
	cliCmd.AddCommand(list.ListCmd)
	cliCmd.AddCommand(remove.RemoveCmd)
	cliCmd.AddCommand(purge.PurgeCmd)
}
