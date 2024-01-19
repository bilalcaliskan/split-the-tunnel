package root

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bilalcaliskan/split-the-tunnel/cmd/root/options"

	"github.com/bilalcaliskan/split-the-tunnel/internal/logging"
	"github.com/bilalcaliskan/split-the-tunnel/internal/version"
	"github.com/spf13/cobra"
)

var (
	opts *options.RootOptions
	ver  = version.Get()
)

func init() {
	opts = options.GetRootOptions()
	opts.InitFlags(rootCmd)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "split-the-tunnel",
	Short:   "",
	Long:    ``,
	Version: ver.GitVersion,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger()
		logger.Info().Str("appVersion", ver.GitVersion).Str("goVersion", ver.GoVersion).Str("goOS", ver.GoOs).
			Str("goArch", ver.GoArch).Str("gitCommit", ver.GitCommit).Str("buildDate", ver.BuildDate).
			Msg("split-the-tunnel is started!")

		// Setup signal handling for a graceful shutdown
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		// Initialize IPC mechanism
		if err := initIPC(logger); err != nil {
			logger.Fatal().Err(err).Msg("failed to initialize IPC")
		}

		<-sigs
		logger.Info().Msg("received termination signal, shutting down...")

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

const socketPath = "/tmp/mydaemon.sock"

func initIPC(logger zerolog.Logger) error {
	// Check and remove the socket file if it already exists
	if _, err := os.Stat(socketPath); err == nil {
		if err := os.Remove(socketPath); err != nil {
			return err
		}
	}

	// Listen on the UNIX domain socket
	listener, err := net.Listen("unix", socketPath)
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
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				logger.Error().Err(err).Msg("failed to read from IPC connection")
				continue
			}
		}

		// Process the received command
		if err := processCommand(strings.TrimSpace(message), conn); err != nil {
			logger.Error().Err(err).Msg("failed to process command")
			break
		}
	}
}

func processCommand(command string, conn net.Conn) error {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return errors.New("not enough arguments")
	}

	switch parts[0] {
	case "add":
		if len(parts) < 2 {
			err := errors.New("add command requires a domain name")
			_, err = conn.Write([]byte("Error: 'add' command requires a domain name\n"))
			return err
		}
		return handleAddCommand(parts[1], conn)
	case "remove":
		if len(parts) < 2 {
			err := errors.New("remove command requires a domain name")
			_, err = conn.Write([]byte("Error: 'remove' command requires a domain name\n"))
			return err
		}
		return handleRemoveCommand(parts[1], conn)
	case "list":
		return handleListCommand(conn)
	default:
		_, err := conn.Write([]byte("Error: Unknown command\n"))
		return err
	}
}

func handleAddCommand(domain string, conn net.Conn) error {
	// Add logic to handle adding route for domain
	// Placeholder for adding functionality

	//_, err := conn.Write([]byte("Added route for " + domain + "\n"))
	//return err

	fmt.Println("handling add command for domain: " + domain)
	return nil
}

func handleRemoveCommand(domain string, conn net.Conn) error {
	// Add logic to handle removing route for domain
	// Placeholder for removing functionality

	//_, err := conn.Write([]byte("Removed route for " + domain + "\n"))
	//return err

	fmt.Println("handling remove command for domain: " + domain)
	return nil
}

func handleListCommand(conn net.Conn) error {
	// Add logic to list all routes
	// Placeholder for listing functionality

	//_, err := conn.Write([]byte("List of all routes\n"))
	//return err

	fmt.Println("handling list command")
	return nil
}
