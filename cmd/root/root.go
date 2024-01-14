package root

import (
	"context"
	"os"

	"github.com/bilalcaliskan/split-the-tunnel/cmd/root/options"

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
	opts.InitFlags(rootCmd)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "split-the-tunnel",
	Short:   "",
	Long:    ``,
	Version: ver.GitVersion,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger()
		logger.Info().Str("appVersion", ver.GitVersion).Str("goVersion", ver.GoVersion).Str("goOS", ver.GoOs).
			Str("goArch", ver.GoArch).Str("gitCommit", ver.GitCommit).Str("buildDate", ver.BuildDate).
			Msg("split-the-tunnel is started!")

		logger.Info().Msg(opts.Domain)
		logger.Info().Msg(opts.DnsServers)

		cmd.SetContext(context.WithValue(cmd.Context(), options.LoggerKey{}, logger))
		cmd.SetContext(context.WithValue(cmd.Context(), options.OptsKey{}, opts))

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
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
