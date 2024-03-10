package config

// Config struct represents the config file
type Config struct {
	CheckIntervalMin int      `toml:"check-interval-min"`
	DnsServers       []string `toml:"dns-servers"`
	Verbose          bool     `toml:"verbose"`
	SocketPath       string   `toml:"socket-path"`
}
