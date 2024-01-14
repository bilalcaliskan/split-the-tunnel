package options

import "github.com/spf13/cobra"

var rootOptions = &RootOptions{}

type (
	OptsKey   struct{}
	LoggerKey struct{}
	DomainKey struct{}
	DnsKey    struct{}
)

// RootOptions contains frequent command line and application options.
type RootOptions struct {
	Domain     string
	DnsServers string
}

// GetRootOptions returns the pointer of RootOptions
func GetRootOptions() *RootOptions {
	return rootOptions
}

func (opts *RootOptions) InitFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&opts.Domain, "domain", "", "", "domain to be used for split tunneling")
	cmd.Flags().StringVarP(&opts.DnsServers, "dns-servers", "", "", "comma separated dns servers to be used for split tunneling")
}
