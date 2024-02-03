package ipc

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"github.com/bilalcaliskan/split-the-tunnel/internal/constants"

	"github.com/bilalcaliskan/split-the-tunnel/internal/utils"

	"github.com/bilalcaliskan/split-the-tunnel/internal/state"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func InitIPC(path string, logger zerolog.Logger) error {
	// Check and remove the socket file if it already exists
	if _, err := os.Stat(path); err == nil {
		if err := os.Remove(path); err != nil {
			return err
		}
	}

	// Listen on the UNIX domain socket
	listener, err := net.Listen("unix", path)
	if err != nil {
		return err
	}

	go func() {
		defer listener.Close()
		for {
			// Accept new connections
			conn, err := listener.Accept()
			if err != nil {
				logger.Error().Err(err).Msg(constants.FailedToAcceptConnection)
				continue
			}

			// Handle the connection in a new goroutine
			go handleConnection(conn, logger)
		}
	}()

	return nil
}

func handleConnection(conn net.Conn, logger zerolog.Logger) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				logger.Error().Err(err).Msg(constants.FailedToReadFromIPC)
			}

			break
		}

		command := strings.TrimSpace(message)
		logger.Info().Str("command", command).Msg("received command")

		st := new(state.State)
		if err := st.Read("/tmp/state.json"); err != nil {
			logger.Error().Err(err).Msg(constants.FailedToReadState)
			continue
		}

		// get default gateway
		gw, err := utils.GetDefaultNonVPNGateway()
		if err != nil {
			logger.Error().Err(err).Msg(constants.FailedToGetDefaultGateway)
			continue
		}

		processCommand(logger, command, gw, conn, st)
	}
}

func processCommand(logger zerolog.Logger, command, gateway string, conn net.Conn, st *state.State) {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		logger.Error().Msg(constants.EmptyCommandReceived)
		return
	}

	switch parts[0] {
	case "add":
		logger = logger.With().Str("operation", "add").Logger()

		handleAddCommand(logger, gateway, parts[1:], conn, st)
	case "remove":
		logger = logger.With().Str("operation", "remove").Logger()

		handleRemoveCommand(logger, gateway, parts[1:], conn, st)
	case "list":
		logger = logger.With().Str("operation", "list").Logger()

		handleListCommand(logger, conn, st)
	}
}

func handleAddCommand(logger zerolog.Logger, gw string, domains []string, conn net.Conn, st *state.State) {
	logger = logger.With().Str("operation", "add").Logger()

	for _, domain := range domains {
		response := new(DaemonResponse)

		ip, err := utils.ResolveDomain(domain)
		if err != nil {
			response.Success = false
			response.Response = ""
			response.Error = errors.Wrap(err, constants.FailedToResolveDomain).Error()

			responseJson, err := json.Marshal(response)
			if err != nil {
				logger.Error().
					Err(err).
					Str("domain", domain).
					Msg(constants.FailedToMarshalResponse)
				continue
			}

			if _, err := conn.Write(responseJson); err != nil {
				logger.Error().
					Err(err).
					Str("domain", domain).
					Msg(constants.FailedToWriteToUnixDomainSocket)
				continue
			}

			continue
		}

		re := &state.RouteEntry{
			Domain:     domain,
			ResolvedIP: ip[0],
			Gateway:    gw,
		}

		if err := st.AddEntry(re); err != nil {
			response.Success = false
			response.Response = ""
			response.Error = errors.Wrap(err, constants.FailedToWriteRouteEntry).Error()

			responseJson, err := json.Marshal(response)
			if err != nil {
				logger.Error().
					Err(err).
					Str("domain", domain).
					Msg(constants.FailedToMarshalResponse)
				continue
			}

			if _, err := conn.Write(responseJson); err != nil {
				logger.Error().
					Err(err).
					Str("domain", domain).
					Msg(constants.FailedToWriteToUnixDomainSocket)
				continue
			}
		}

		response.Success = false
		response.Response = fmt.Sprintf("added route for " + domain)
		response.Error = ""

		responseJson, err := json.Marshal(response)
		if err != nil {
			logger.Error().
				Err(err).
				Str("domain", domain).
				Msg(constants.FailedToMarshalResponse)
			continue
		}

		// Send a response to the client
		_, err = conn.Write(responseJson)
		if err != nil {
			logger.Error().
				Err(err).
				Str("domain", domain).
				Msg(constants.FailedToWriteToUnixDomainSocket)
			continue
		}
	}
}

func handleRemoveCommand(logger zerolog.Logger, gw string, domains []string, conn net.Conn, st *state.State) {
	for _, domain := range domains {
		response := new(DaemonResponse)

		response.Success = false
		response.Response = ""
		response.Error = fmt.Sprintf("a dummy error for domain %s", domain)

		responseJson, err := json.Marshal(response)
		if err != nil {
			logger.Error().
				Err(err).
				Str("domain", domain).
				Msg(constants.FailedToMarshalResponse)
			continue
		}

		if _, err := conn.Write(responseJson); err != nil {
			logger.Error().
				Err(err).
				Str("domain", domain).
				Msg(constants.FailedToWriteToUnixDomainSocket)
			continue
		}

		continue
	}
}

func handleListCommand(logger zerolog.Logger, conn net.Conn, st *state.State) {
	response := new(DaemonResponse)

	response.Success = false
	response.Response = ""
	response.Error = "a dummy error list command"

	responseJson, err := json.Marshal(response)
	if err != nil {
		logger.Error().
			Err(err).
			Msg(constants.FailedToMarshalResponse)
		return
	}

	if _, err := conn.Write(responseJson); err != nil {
		logger.Error().
			Err(err).
			Msg(constants.FailedToWriteToUnixDomainSocket)
		return
	}
}

func Cleanup(path string) error {
	// Perform any cleanup and shutdown tasks here

	return os.Remove(path)
}
