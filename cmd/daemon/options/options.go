package options

import (
	"os"
	"path/filepath"

	"github.com/bilalcaliskan/split-the-tunnel/internal/constants"

	"github.com/spf13/viper"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var rootOptions = &RootOptions{}

type RootOptions struct {
	// Workspace is the directory path where the application will store its data
	Workspace string
	// ConfigFile is the path of the configuration file, which will be searched in the Workspace
	ConfigFile string
	// SocketPath is the path of the socket file, which will be stored in the Workspace
	SocketPath string
	// StatePath is the path of the state file, which will be stored in the Workspace
	StatePath string

	// DnsServers is the list of DNS servers to be used for DNS resolving
	DnsServers string `toml:"dnsservers"`
	// CheckIntervalMin is the interval in minutes to check the routing table with the collected state.State
	CheckIntervalMin int `toml:"checkintervalmin"`
	// Verbose is the flag to enable verbose logging output
	Verbose bool `toml:"verbose"`
}

// GetRootOptions returns the pointer of RootOptions
func GetRootOptions() *RootOptions {
	return rootOptions
}

// InitFlags initializes the flags of the root command
func (opts *RootOptions) InitFlags(cmd *cobra.Command) error {
	if err := opts.setFlags(cmd); err != nil {
		return errors.Wrap(err, "failed to set flags")
	}

	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return errors.Wrap(err, "failed to bind flags")
	}

	return nil
}

// setFlags sets the flags of the root command
func (opts *RootOptions) setFlags(cmd *cobra.Command) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get user home directory")
	}

	cmd.Flags().StringVarP(&opts.Workspace, "workspace", "w", filepath.Join(homeDir, ".split-the-tunnel"), "workspace directory path")
	cmd.Flags().StringVarP(&opts.ConfigFile, "config-file", "c", "config.toml", "config file path, will search in workspace")
	cmd.Flags().BoolVarP(&opts.Verbose, "verbose", "", false, "verbose logging output")
	cmd.Flags().StringVarP(&opts.DnsServers, "dns-servers", "", "", "comma separated dns servers to be used for DNS resolving")
	cmd.Flags().IntVarP(&opts.CheckIntervalMin, "check-interval-min", "", 5, "routing table check interval with collected state, in minutes")

	return nil
}

// ReadConfig reads the configuration file and unmarshalls it into RootOptions
func (opts *RootOptions) ReadConfig() error {
	viper.SetConfigType("toml")
	viper.SetConfigFile(filepath.Join(opts.Workspace, opts.ConfigFile))
	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "failed to read config file")
	}

	if err := viper.Unmarshal(opts); err != nil {
		return errors.Wrap(err, "failed to unmarshal config file")
	}

	opts.ConfigFile = filepath.Join(opts.Workspace, opts.ConfigFile)
	opts.StatePath = filepath.Join(opts.Workspace, constants.StateFileName)
	opts.SocketPath = filepath.Join(opts.Workspace, constants.SocketFileName)

	return nil
}
