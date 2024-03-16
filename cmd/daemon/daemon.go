package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"

	"github.com/pkg/errors"

	"github.com/bilalcaliskan/split-the-tunnel/internal/state"

	"github.com/bilalcaliskan/split-the-tunnel/internal/constants"

	"github.com/bilalcaliskan/split-the-tunnel/cmd/daemon/options"
	"github.com/bilalcaliskan/split-the-tunnel/internal/ipc"
	"github.com/bilalcaliskan/split-the-tunnel/internal/logging"
	"github.com/bilalcaliskan/split-the-tunnel/internal/version"
	"github.com/spf13/cobra"
)

var (
	opts *options.RootOptions
	ver  = version.Get()
)

func init() {
	opts = options.GetRootOptions()
	if err := opts.InitFlags(daemonCmd); err != nil {
		panic(errors.Wrap(err, "failed to initialize flags"))
	}
}

// daemonCmd represents the base command when called without any subcommands
var daemonCmd = &cobra.Command{
	Use:     "split-the-tunnel",
	Short:   "",
	Long:    ``,
	Version: ver.GitVersion,
	RunE: func(cmd *cobra.Command, args []string) error {
		//opts := options.GetRootOptions()
		//if err := opts.InitFlags(cmd); err != nil {
		//	panic(errors.Wrap(err, "failed to initialize flags"))
		//}

		if err := os.MkdirAll(opts.Workspace, 0755); err != nil {
			return errors.Wrap(err, "failed to create workspace directory")
		}

		//if err := viper.BindPFlags(cmd.Flags()); err != nil {
		//	return errors.Wrap(err, "failed to bind flags")
		//}
		//opts = options.GetRootOptions()
		if err := opts.ReadConfig(); err != nil {
			return errors.Wrap(err, "failed to read config")
		}

		fmt.Println(opts)

		fmt.Println("dns_servers:", viper.GetString("dns_servers"))
		fmt.Println("check_interval_min:", viper.GetInt("check_interval_min"))
		fmt.Println("verbose:", viper.GetBool("verbose"))

		logger := logging.GetLogger().With().Str("job", "main").Logger()
		logger.Info().Str("appVersion", ver.GitVersion).Str("goVersion", ver.GoVersion).Str("goOS", ver.GoOs).
			Str("goArch", ver.GoArch).Str("gitCommit", ver.GitCommit).Str("buildDate", ver.BuildDate).
			Msg(constants.AppStarted)

		// Setup signal handling for a graceful shutdown
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		st := state.NewState(logger)

		// Initialize IPC mechanism
		if err := ipc.InitIPC(st, opts.SocketPath, logger); err != nil {
			logger.Error().Err(err).Msg(constants.FailedToInitializeIPC)
			return err
		}

		logger.Info().Str("socket", opts.SocketPath).Msg(constants.IPCInitialized)

		defer func() {
			if err := ipc.Cleanup(opts.SocketPath); err != nil {
				logger.Error().Err(err).Msg(constants.FailedToCleanupIPC)
			}
		}()

		logger.Info().Str("socket", opts.SocketPath).Msg(constants.DaemonRunning)

		go func() {
			// Create a ticker that fires every 5 minutes
			ticker := time.NewTicker(10 * time.Second)
			logger := logger.With().Str("job", "ip-change-check").Logger()

			for range ticker.C {
				if <-ticker.C; true {
					if err := st.CheckIPChanges(); err != nil {
						logger.Error().Err(err).Msg("failed to check ip changes")
					}
				}
			}
		}()

		// Wait for termination signal to gracefully shut down the daemon
		s := <-sigs
		logger.Info().Any("signal", s.String()).Msg(constants.TermSignalReceived)
		logger.Info().Msg(constants.ShuttingDownDaemon)

		return nil
	},
}

func main() {
	if err := daemonCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
