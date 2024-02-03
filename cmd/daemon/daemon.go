package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bilalcaliskan/split-the-tunnel/internal/constants"

	"github.com/bilalcaliskan/split-the-tunnel/cmd/daemon/options"
	"github.com/bilalcaliskan/split-the-tunnel/internal/ipc"
	"github.com/bilalcaliskan/split-the-tunnel/internal/logging"
	"github.com/bilalcaliskan/split-the-tunnel/internal/version"
	"github.com/spf13/cobra"
)

const socketPath = "/tmp/mydaemon.sock"

var (
	opts *options.RootOptions
	ver  = version.Get()
)

func init() {
	opts = options.GetRootOptions()
	opts.InitFlags(daemonCmd)
}

// daemonCmd represents the base command when called without any subcommands
var daemonCmd = &cobra.Command{
	Use:     "split-the-tunnel",
	Short:   "",
	Long:    ``,
	Version: ver.GitVersion,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger()
		logger.Info().Str("appVersion", ver.GitVersion).Str("goVersion", ver.GoVersion).Str("goOS", ver.GoOs).
			Str("goArch", ver.GoArch).Str("gitCommit", ver.GitCommit).Str("buildDate", ver.BuildDate).
			Msg(constants.AppStarted)

		// Setup signal handling for a graceful shutdown
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		// Initialize IPC mechanism
		if err := ipc.InitIPC(socketPath, logger); err != nil {
			logger.Error().Err(err).Msg(constants.FailedToInitializeIPC)
			return err
		}

		logger.Info().Str("socket", socketPath).Msg(constants.IPCInitialized)

		defer func() {
			if err := ipc.Cleanup(socketPath); err != nil {
				logger.Error().Err(err).Msg(constants.FailedToCleanupIPC)
			}
		}()

		logger.Info().Msg(constants.DaemonRunning)

		// Wait for termination signal
		<-sigs

		logger.Info().Msg(constants.TermSignalReceived)
		logger.Info().Msg(constants.ShuttingDownDaemon)

		return nil
	},
}

func main() {
	if err := daemonCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
