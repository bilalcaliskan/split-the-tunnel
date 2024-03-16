package constants

const (
	FailedToAcceptConnection          = "failed to accept connection"
	FailedToConnectToUnixDomainSocket = "failed to connect to unix domain socket"
	FailedToOpenRoutingInfoFile       = "failed to open routing info file"
	FailedToWriteToUnixDomainSocket   = "failed to write to unix domain socket"
	FailedToWriteState                = "failed to write state to file"
	FailedToReadFromIPC               = "failed to read from IPC connection"
	FailedToReloadState               = "failed to reload state"
	FailedToGetDefaultGateway         = "failed to get default gateway"
	EmptyCommandReceived              = "empty command received"
	FailedToResolveDomain             = "failed to resolve domain"
	FailedToMarshalResponse           = "failed to marshal response"
	EntryNotFound                     = "route entry not found in state"
	NonVPNGatewayNotFound             = "non-VPN gateway not found"
	FailedToDecodeHex                 = "failed to decode hex string"
	InvalidIpLength                   = "invalid IP length: %d"
	FailedToProcessCommand            = "failed to process command"
	FailedToCleanupIPC                = "failed to cleanup IPC"
	FailedToInitializeIPC             = "failed to initialize IPC"
	FailedToRemoveRouteEntry          = "failed to remove RouteEntry from state"
)
