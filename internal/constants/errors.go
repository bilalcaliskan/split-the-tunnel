package constants

const (
	FailedToAcceptConnection          = "failed to accept connection"
	FailedToConnectToUnixDomainSocket = "failed to connect to unix domain socket"
	FailedToOpenRoutingInfoFile       = "failed to open routing info file"
	FailedToWriteToUnixDomainSocket   = "failed to write to unix domain socket"
	FileDoesNotExist                  = "file does not exist"
	NoArgumentsProvided               = "no arguments provided"
	FailedToReadFromIPC               = "failed to read from IPC connection"
	FailedToReadState                 = "failed to read state"
	FailedToGetDefaultGateway         = "failed to get default gateway"
	EmptyCommandReceived              = "empty command received"
	FailedToResolveDomain             = "failed to resolve domain"
	FailedToMarshalResponse           = "failed to marshal response"
	FailedToWriteRouteEntry           = "failed to write RouteEntry to state"
	EntryNotFound                     = "entry not found"
	NonVPNGatewayNotFound             = "non-VPN gateway not found"
	FailedToDecodeHex                 = "failed to decode hex string"
	InvalidIpLength                   = "invalid IP length: %d"
	FailedToSendCommand               = "failed to send command to daemon"
	FailedToCleanupIPC                = "failed to cleanup IPC"
	FailedToInitializeIPC             = "failed to initialize IPC"
)
