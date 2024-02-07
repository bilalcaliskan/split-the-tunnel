package options

import "github.com/spf13/cobra"

var rootOptions = &RootOptions{}

// RootOptions contains frequent command line and application options.
type RootOptions struct {
	DnsServers       string
	CheckIntervalMin int
	Verbose          bool
}

// GetRootOptions returns the pointer of RootOptions
func GetRootOptions() *RootOptions {
	return rootOptions
}

func (opts *RootOptions) InitFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&opts.Verbose, "verbose", "", false, "verbose logging output")
	cmd.Flags().StringVarP(&opts.DnsServers, "dns-servers", "", "", "comma separated dns servers to be used for DNS resolving")
	cmd.Flags().IntVarP(&opts.CheckIntervalMin, "check-interval-min", "", 5, "routing table check interval with collected state, in minutes")
}
