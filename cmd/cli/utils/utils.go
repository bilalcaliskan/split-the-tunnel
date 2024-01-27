package utils

import (
	"bufio"
	"errors"
	"net"
)

const SocketPath = "/tmp/mydaemon.sock"

var (
	ErrNoArgs      = errors.New("no arguments provided")
	ErrTooManyArgs = errors.New("too many arguments provided")
)

func SendCommandToDaemon(socketPath, command string) (string, error) {
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	_, err = conn.Write([]byte(command + "\n"))
	if err != nil {
		return "", err
	}

	// Read response from daemon
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return "", err
	}

	return response, nil
}
