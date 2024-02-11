package ipc

// DaemonResponse is the struct that holds the response of the daemon to the client
type DaemonResponse struct {
	Success  bool   `json:"success"`
	Response string `json:"response"`
	Error    string `json:"error"`
}
