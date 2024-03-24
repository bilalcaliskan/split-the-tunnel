package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"

	"github.com/bilalcaliskan/split-the-tunnel/internal/state"

	"github.com/bilalcaliskan/split-the-tunnel/internal/constants"

	"github.com/bilalcaliskan/split-the-tunnel/cmd/daemon/options"
	"github.com/bilalcaliskan/split-the-tunnel/internal/ipc"
	"github.com/bilalcaliskan/split-the-tunnel/internal/logging"
	"github.com/bilalcaliskan/split-the-tunnel/internal/version"
	"github.com/spf13/cobra"
)

func init() {
	opts = options.GetRootOptions()
	if err := opts.InitFlags(daemonCmd); err != nil {
		panic(errors.Wrap(err, "failed to initialize flags"))
	}
}

var (
	// opts is the root options of the application
	opts *options.RootOptions
	// ver is the version of the application
	ver = version.Get()
	// daemonCmd represents the base command when called without any subcommands
	daemonCmd = &cobra.Command{
		Use:     "split-the-tunnel",
		Short:   "",
		Long:    ``,
		Version: ver.GitVersion,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := os.MkdirAll(opts.Workspace, 0755); err != nil {
				return errors.Wrap(err, "failed to create workspace directory")
			}

			if err := opts.ReadConfig(); err != nil {
				return errors.Wrap(err, "failed to read config")
			}

			fmt.Println(opts)

			logger := logging.GetLogger().With().Str("job", constants.JobMain).Logger()
			logger.Info().Str("appVersion", ver.GitVersion).Str("goVersion", ver.GoVersion).Str("goOS", ver.GoOs).
				Str("goArch", ver.GoArch).Str("gitCommit", ver.GitCommit).Str("buildDate", ver.BuildDate).
				Msg(constants.AppStarted)

			st := state.NewState(logger, opts.StatePath)

			// initialize IPC for communication between CLI and daemon
			if err := ipc.InitIPC(st, opts.SocketPath, logger); err != nil {
				logger.Error().Err(err).Msg(constants.FailedToInitializeIPC)
				return err
			}

			logger.Info().Str("socket", opts.SocketPath).Msg(constants.IPCInitialized)

			defer func() {
				logger := logger.With().Str("job", constants.JobCleanup).Logger()
				logger.Info().Msg(constants.CleaningUpIPC)
				if err := ipc.Cleanup(opts.SocketPath); err != nil {
					logger.Error().Err(err).Msg(constants.FailedToCleanupIPC)
				}
			}()

			logger.Info().Str("socket", opts.SocketPath).Msg(constants.DaemonRunning)

			go func() {
				// Create a ticker that fires every 5 minutes
				ticker := time.NewTicker(time.Duration(int64(opts.CheckIntervalMin)) * time.Minute)
				logger := logger.With().Str("job", constants.JobIpChangeCheck).Logger()

				for range ticker.C {
					if <-ticker.C; true {
						if err := st.CheckIPChanges(); err != nil {
							logger.Error().Err(err).Msg("failed to check ip changes")
						}
					}
				}
			}()

			// setup signal handling for graceful shutdown
			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

			// Wait for termination signal to gracefully shut down the daemon
			s := <-sigs
			logger.Info().Any("signal", s.String()).Msg(constants.TermSignalReceived)
			logger.Info().Msg(constants.ShuttingDownDaemon)

			return nil
		},
	}
)

func main() {
	if err := daemonCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
