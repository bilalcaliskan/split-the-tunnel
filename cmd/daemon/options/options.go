package options

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var rootOptions = &RootOptions{}

type RootOptions struct {
	Workspace        string // This field will be managed by the command line argument
	ConfigFile       string // This field will be managed by the command line argument
	DnsServers       string `toml:"dns_servers"`        // This field will be managed by the config file
	CheckIntervalMin int    `toml:"check_interval_min"` // This field will be managed by the config file
	SocketPath       string `toml:"socket_path"`        // This field will be managed by the config file
	StatePath        string `toml:"state_path"`         // This field will be managed by the config file
	Verbose          bool   `toml:"verbose"`            // This field will be managed by the config file
}

// GetRootOptions returns the pointer of RootOptions
func GetRootOptions() *RootOptions {
	return rootOptions
}

func (opts *RootOptions) InitFlags(cmd *cobra.Command) error {
	if err := opts.setFlags(cmd); err != nil {
		return errors.Wrap(err, "failed to set flags")
	}

	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return errors.Wrap(err, "failed to bind flags")
	}

	return nil
}

func (opts *RootOptions) setFlags(cmd *cobra.Command) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get user home directory")
	}

	cmd.Flags().StringVarP(&opts.Workspace, "workspace", "w", filepath.Join(homeDir, ".split-the-tunnel"), "workspace directory path")
	cmd.Flags().StringVarP(&opts.ConfigFile, "config-file", "c", "config.toml", "config file path, will search in workspace")
	cmd.Flags().StringVarP(&opts.SocketPath, "socket-path", "", "ipc.sock", "unix domain socket path in workspace")
	cmd.Flags().StringVarP(&opts.StatePath, "state-path", "", "state.json", "state file path in workspace")
	cmd.Flags().BoolVarP(&opts.Verbose, "verbose", "", false, "verbose logging output")
	cmd.Flags().StringVarP(&opts.DnsServers, "dns-servers", "", "", "comma separated dns servers to be used for DNS resolving")
	cmd.Flags().IntVarP(&opts.CheckIntervalMin, "check-interval-min", "", 5, "routing table check interval with collected state, in minutes")

	return nil
}

func (opts *RootOptions) ReadConfig() error {
	opts.SocketPath = filepath.Join(opts.Workspace, opts.SocketPath)
	opts.ConfigFile = filepath.Join(opts.Workspace, opts.ConfigFile)
	opts.StatePath = filepath.Join(opts.Workspace, opts.StatePath)

	if _, err := os.Stat(opts.ConfigFile); os.IsNotExist(err) {
		log.Printf("config file not found in %s, will use default values\n", opts.ConfigFile)
		return nil
	} else if err != nil {
		return errors.Wrap(err, "failed to access config file")
	}

	viper.SetConfigFile(opts.ConfigFile)
	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "failed to read config file")
	}

	if err := viper.Unmarshal(opts); err != nil {
		return errors.Wrap(err, "failed to unmarshal config file")
	}

	return nil
}
