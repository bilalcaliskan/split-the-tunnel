package ipc

type DaemonResponse struct {
	Success  bool   `json:"success"`
	Response string `json:"response"`
	Error    string `json:"error"`
}
