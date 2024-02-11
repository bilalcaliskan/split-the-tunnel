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

// InitIPC initializes the IPC setup and continuously listens on the given path for incoming connections
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

// handleConnection handles the incoming connection
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

		st := state.NewState()
		if err := st.Read(constants.StateFilePath); err != nil {
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

// processCommand processes the given command and calls the appropriate handler
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

		handleRemoveCommand(logger, parts[1:], conn, st)
	case "list":
		logger = logger.With().Str("operation", "list").Logger()

		handleListCommand(logger, conn, st)
	case "purge":
		logger = logger.With().Str("operation", "purge").Logger()

		handlePurgeCommand(logger, conn, st)
	}
}

// handleAddCommand handles the add command and adds the given domains to the routing table
func handleAddCommand(logger zerolog.Logger, gw string, domains []string, conn net.Conn, st *state.State) {
	logger = logger.With().Str("operation", "add").Logger()

	for _, domain := range domains {
		ips, err := utils.ResolveDomain(domain)
		if err != nil {
			if err := writeResponse(&DaemonResponse{
				Success:  false,
				Response: "",
				Error:    errors.Wrap(err, constants.FailedToResolveDomain).Error(),
			}, conn); err != nil {
				logger.Error().
					Err(err).
					Str("domain", domain).
					Msg(constants.FailedToWriteToUnixDomainSocket)
			}

			continue
		}

		re := state.NewRouteEntry(domain, gw, ips)

		if err := st.AddEntry(re); err != nil {
			logger.Error().Err(err).Str("domain", domain).Msg("failed to add route to state")

			if err := writeResponse(&DaemonResponse{
				Success:  false,
				Response: "",
				Error:    errors.Wrapf(err, "failed to write RouteEntry to state for domain %s", domain).Error(),
			}, conn); err != nil {
				logger.Error().
					Err(err).
					Str("domain", domain).
					Msg(constants.FailedToWriteToUnixDomainSocket)
			}

			continue
		}

		for _, ip := range re.ResolvedIPs {
			if err := utils.AddRoute(ip, re.Gateway); err != nil {
				logger.Error().Err(err).Str("domain", domain).Str("ip", ip).Msg("failed to add route to routing table")

				if err := writeResponse(&DaemonResponse{
					Success:  false,
					Response: "",
					Error:    errors.Wrapf(err, "failed to add route for domain %s to routing table", domain).Error(),
				}, conn); err != nil {
					logger.Error().
						Err(err).
						Str("domain", domain).
						Msg(constants.FailedToWriteToUnixDomainSocket)
				}

				continue
			}

		}

		logger.Info().Str("domain", domain).Msg("successfully added route to routing table")

		if err := writeResponse(&DaemonResponse{
			Success:  true,
			Response: fmt.Sprintf("added route for " + domain),
			Error:    "",
		}, conn); err != nil {
			logger.Error().
				Err(err).
				Str("domain", domain).
				Msg(constants.FailedToWriteToUnixDomainSocket)
		}
	}
}

func handlePurgeCommand(logger zerolog.Logger, conn net.Conn, st *state.State) {
	logger = logger.With().Str("operation", "purge").Logger()

	for _, entry := range st.Entries {
		for _, ip := range entry.ResolvedIPs {
			if err := utils.RemoveRoute(ip); err != nil {
				logger.Error().Err(err).Str("domain", entry.Domain).Str("ip", ip).Msg("failed to remove route from routing table")

				if err := writeResponse(&DaemonResponse{
					Success:  false,
					Response: "",
					Error:    errors.Wrapf(err, "failed to remove route for domain %s from routing table", entry.Domain).Error(),
				}, conn); err != nil {
					logger.Error().
						Err(err).
						Str("domain", entry.Domain).
						Msg(constants.FailedToWriteToUnixDomainSocket)
				}

				continue
			}
		}

		logger.Info().Str("domain", entry.Domain).Msg("successfully removed route from routing table")
	}

	st.Entries = make([]*state.RouteEntry, 0)

	if err := st.Write(constants.StateFilePath); err != nil {
		logger.Error().Err(err).Msg("failed to write state to file")
	}

	if err := writeResponse(&DaemonResponse{
		Success:  true,
		Response: "purged all routes",
		Error:    "",
	}, conn); err != nil {
		logger.Error().
			Err(err).
			Msg(constants.FailedToWriteToUnixDomainSocket)
	}
}

// handleRemoveCommand removes the given domains from the routing table
func handleRemoveCommand(logger zerolog.Logger, domains []string, conn net.Conn, st *state.State) {
	logger = logger.With().Str("operation", "remove").Logger()

	for _, domain := range domains {
		entry := st.GetEntry(domain)
		if entry == nil {
			if err := writeResponse(&DaemonResponse{
				Success:  false,
				Response: "",
				Error:    errors.Wrap(errors.New(constants.EntryNotFound), constants.FailedToRemoveRouteEntry).Error(),
			}, conn); err != nil {
				logger.Error().
					Err(err).
					Str("domain", domain).
					Msg(constants.FailedToWriteToUnixDomainSocket)
			}
			continue
		}

		if err := st.RemoveEntry(domain); err != nil {
			logger.Error().Err(err).Str("domain", domain).Msg("failed to remove route from state")

			if err := writeResponse(&DaemonResponse{
				Success:  false,
				Response: "",
				Error:    errors.Wrap(err, constants.FailedToRemoveRouteEntry).Error(),
			}, conn); err != nil {
				logger.Error().
					Err(err).
					Str("domain", domain).
					Msg(constants.FailedToWriteToUnixDomainSocket)
			}
			continue
		}

		for _, ip := range entry.ResolvedIPs {
			if err := utils.RemoveRoute(ip); err != nil {
				logger.Error().Err(err).Str("domain", domain).Str("ip", ip).Msg("failed to remove route from routing table")

				if err := writeResponse(&DaemonResponse{
					Success:  false,
					Response: "",
					Error:    errors.Wrapf(err, "failed to remove route for domain %s from routing table", domain).Error(),
				}, conn); err != nil {
					logger.Error().
						Err(err).
						Str("domain", domain).
						Msg(constants.FailedToWriteToUnixDomainSocket)
				}

				continue
			}
		}

		logger.Info().Str("domain", domain).Msg("successfully removed route from routing table")

		if err := writeResponse(&DaemonResponse{
			Success:  true,
			Response: fmt.Sprintf("removed route for " + domain),
			Error:    "",
		}, conn); err != nil {
			logger.Error().
				Err(err).
				Str("domain", domain).
				Msg(constants.FailedToWriteToUnixDomainSocket)
		}
	}
}

func handleListCommand(logger zerolog.Logger, conn net.Conn, st *state.State) {
	response := new(DaemonResponse)

	response.Success = false
	response.Response = ""
	response.Error = ""

	// print the state
	for _, entry := range st.Entries {
		response.Response += fmt.Sprintf("Domain: %s, Gateway: %s, IPs: %v\n", entry.Domain, entry.Gateway, entry.ResolvedIPs)
	}

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
