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

// RootOptions contains frequent command line and application options.
type RootOptions struct {
	Workspace        string // This field will be managed by the command line argument
	ConfigFile       string // This field will be managed by the command line argument
	DnsServers       string `toml:"dns-servers"`        // This field will be managed by the config file
	CheckIntervalMin int    `toml:"check-interval-min"` // This field will be managed by the config file
	Verbose          bool   `toml:"verbose"`            // This field will be managed by the config file
}

// GetRootOptions returns the pointer of RootOptions
func GetRootOptions() *RootOptions {
	return rootOptions
}

func (opts *RootOptions) InitFlags(cmd *cobra.Command) error {
	if err := opts.setWorkspace(); err != nil {
		return err
	}

	opts.setFlags(cmd)
	viper.BindPFlags(cmd.Flags())
	viper.SetConfigFile(opts.ConfigFile)
	viper.SetConfigType("toml")

	if err := opts.readConfig(); err != nil {
		return err
	}

	return nil
}

func (opts *RootOptions) setWorkspace() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get user home directory")
	}

	opts.Workspace = filepath.Join(homeDir, ".split-the-tunnel")
	opts.ConfigFile = filepath.Join(opts.Workspace, "config.toml")

	return nil
}

func (opts *RootOptions) setFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&opts.Workspace, "workspace", "w", opts.Workspace, "workspace directory path")
	cmd.Flags().StringVarP(&opts.ConfigFile, "config-file", "c", opts.ConfigFile, "config file path")
	cmd.Flags().BoolVarP(&opts.Verbose, "verbose", "", false, "verbose logging output")
	cmd.Flags().StringVarP(&opts.DnsServers, "dns-servers", "", "", "comma separated dns servers to be used for DNS resolving")
	cmd.Flags().IntVarP(&opts.CheckIntervalMin, "check-interval-min", "", 5, "routing table check interval with collected state, in minutes")
}

func (opts *RootOptions) readConfig() error {
	if _, err := os.Stat(opts.ConfigFile); os.IsNotExist(err) {
		log.Println("Warning: Config file not found. Using default values.")
	} else if err != nil {
		return errors.Wrap(err, "an error occurred while accessing config file")
	} else {
		if err := viper.ReadInConfig(); err != nil {
			return errors.Wrap(err, "an error occurred while reading config file")
		}

		if err := viper.Unmarshal(opts); err != nil {
			return errors.Wrap(err, "an error occurred while unmarshaling config file")
		}
	}

	return nil
}
