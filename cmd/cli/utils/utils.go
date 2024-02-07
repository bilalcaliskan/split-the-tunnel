package utils

import (
	"encoding/json"
	"net"

	"github.com/bilalcaliskan/split-the-tunnel/internal/constants"

	"github.com/pkg/errors"
)

const SocketPath = "/tmp/mydaemon.sock"

var (
	ErrNoArgs      = errors.New("no arguments provided")
	ErrTooManyArgs = errors.New("too many arguments provided")
)

type DaemonResponse struct {
	Success  bool   `json:"success"`
	Response string `json:"response"`
	Error    string `json:"error"`
}

func SendCommandToDaemon(socketPath, command string) (string, error) {
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return "", errors.Wrap(err, constants.FailedToConnectToUnixDomainSocket)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(command + "\n"))
	if err != nil {
		return "", errors.Wrap(err, constants.FailedToWriteToUnixDomainSocket)
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf[:])
	if err != nil {
		return "", err
	}

	var response DaemonResponse
	if err := json.Unmarshal(buf[:n], &response); err != nil {
		return "", err
	}

	var respErr error
	if response.Error != "" {
		respErr = errors.New(response.Error)
	}

	return response.Response, respErr
}
