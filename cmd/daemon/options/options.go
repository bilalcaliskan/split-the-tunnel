package options

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var rootOptions = &RootOptions{}

// RootOptions contains frequent command line and application options.
type RootOptions struct {
	Workspace        string
	ConfigFile       string
	DnsServers       string
	CheckIntervalMin int
	Verbose          bool
}

// GetRootOptions returns the pointer of RootOptions
func GetRootOptions() *RootOptions {
	return rootOptions
}

func (opts *RootOptions) InitFlags(cmd *cobra.Command) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get user home directory")
	}

	workspace := filepath.Join(homeDir, ".split-the-tunnel")

	cmd.Flags().StringVarP(&opts.Workspace, "workspace", "w", workspace, "workspace directory path")
	cmd.Flags().StringVarP(&opts.ConfigFile, "config-file", "c", filepath.Join(workspace, "config.toml"), "config file path")
	cmd.Flags().BoolVarP(&opts.Verbose, "verbose", "", false, "verbose logging output")
	cmd.Flags().StringVarP(&opts.DnsServers, "dns-servers", "", "", "comma separated dns servers to be used for DNS resolving")
	cmd.Flags().IntVarP(&opts.CheckIntervalMin, "check-interval-min", "", 5, "routing table check interval with collected state, in minutes")

	return nil
}
