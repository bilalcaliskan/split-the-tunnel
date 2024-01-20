package root

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bilalcaliskan/split-the-tunnel/internal/ipc"
	"github.com/bilalcaliskan/split-the-tunnel/internal/utils"

	"github.com/bilalcaliskan/split-the-tunnel/cmd/root/options"
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
	opts.InitFlags(rootCmd)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "split-the-tunnel",
	Short:   "",
	Long:    ``,
	Version: ver.GitVersion,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger()
		logger.Info().Str("appVersion", ver.GitVersion).Str("goVersion", ver.GoVersion).Str("goOS", ver.GoOs).
			Str("goArch", ver.GoArch).Str("gitCommit", ver.GitCommit).Str("buildDate", ver.BuildDate).
			Msg("split-the-tunnel is started!")

		gateway, err := utils.GetDefaultNonVPNGateway()
		if err != nil {
			logger.Error().Err(err).Msg("failed to get default gateway")
			return err
		}

		logger.Info().Str("gateway", gateway).Msg("found default gateway")

		// Setup signal handling for a graceful shutdown
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		// Initialize IPC mechanism
		if err := ipc.InitIPC(socketPath, logger); err != nil {
			logger.Error().Err(err).Msg("failed to initialize IPC")
			return err
		}

		logger.Info().Str("socket", socketPath).Msg("IPC is initialized")

		defer func() {
			if err := ipc.Cleanup(socketPath); err != nil {
				logger.Error().Err(err).Msg("failed to cleanup IPC")
			}
		}()

		logger.Info().Msg("daemon is running...")

		// Wait for termination signal
		<-sigs

		logger.Info().Msg("termination signal received")
		logger.Info().Msg("shutting down daemon...")

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
