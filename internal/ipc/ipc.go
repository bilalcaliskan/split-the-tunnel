package ipc

import (
	"bufio"
	"errors"
	"io"
	"net"
	"os"
	"strings"

	"github.com/bilalcaliskan/split-the-tunnel/internal/state"
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
				logger.Error().Err(err).Msg("failed to accept connection")
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
				logger.Error().Err(err).Msg("error reading from IPC connection")
			}

			break
		}

		command := strings.TrimSpace(message)
		logger.Info().Str("command", command).Msg("received command")

		st := new(state.State)
		if err := st.Read("/tmp/state.json"); err != nil {
			logger.Error().Str("path", "/tmp/state.json").Err(err).Msg("failed to read state")
			continue
		}

		logger.Info().Any("state", st).Msg("read state")

		if err := processCommand(command, conn); err != nil {
			logger.Error().Str("command", command).Err(err).Msg("error processing command")
			continue
		}

		logger.Info().Str("command", command).Msg("command processed successfully")
	}
}

func processCommand(command string, conn net.Conn) error {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return errors.New("empty command received")
	}

	switch parts[0] {
	case "add":
		if len(parts) < 2 {
			_, err := conn.Write([]byte("'add' command requires at least a domain name\n"))
			if err != nil {
				return err
			}

			return errors.New("'add' command requires at least a domain name")
		}

		return handleAddCommand(parts[1:], conn)
	case "remove":
		if len(parts) < 2 {
			_, err := conn.Write([]byte("'remove' command requires at least a domain name\n"))
			if err != nil {
				return err
			}

			return errors.New("'remove' command requires at least a domain name")
		}

		return handleRemoveCommand(parts[1:], conn)
	case "list":
		if len(parts) != 1 {
			_, err := conn.Write([]byte("'list' command does not accept any arguments\n"))
			if err != nil {
				return err
			}

			return errors.New("'list' command does not accept any arguments")
		}

		return handleListCommand(conn)
	default:
		_, err := conn.Write([]byte("unknown command received\n"))
		return err
	}
}

func handleAddCommand(domains []string, conn net.Conn) error {
	// Add the domain to the routing table
	// ...

	for _, domain := range domains {
		// Send a response to the client
		_, err := conn.Write([]byte("added route for " + domain + "\n"))
		if err != nil {
			return err
		}
	}

	return nil
}

func handleRemoveCommand(domains []string, conn net.Conn) error {
	// Remove the domain from the routing table
	// ...

	for _, domain := range domains {
		// Send a response to the client
		_, err := conn.Write([]byte("removed route for " + domain + "\n"))
		if err != nil {
			return err
		}
	}

	return nil
}

func handleListCommand(conn net.Conn) error {
	// List the domains that we manage from the routing table
	// ...

	// Send a response to the client
	_, err := conn.Write([]byte("listing routes\n"))
	return err
}

func Cleanup(path string) error {
	// Perform any cleanup and shutdown tasks here

	return os.Remove(path)
}
