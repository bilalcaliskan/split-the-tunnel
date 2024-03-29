package ipc

import (
	"encoding/json"
	"net"
	"os"

	"github.com/bilalcaliskan/split-the-tunnel/internal/constants"
	"github.com/pkg/errors"
)

// Cleanup performs any cleanup and shutdown tasks
func Cleanup(path string) error {
	// Perform any cleanup and shutdown tasks here

	return os.Remove(path)
}

// writeResponse writes the response to the connection
func writeResponse(response *DaemonResponse, conn net.Conn) error {
	responseJson, err := json.Marshal(response)
	if err != nil {
		return errors.Wrap(err, constants.FailedToMarshalResponse)
	}

	_, err = conn.Write(responseJson)

	return err
}
