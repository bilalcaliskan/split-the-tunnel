package root

import (
	"bufio"
	"errors"
	"io"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bilalcaliskan/split-the-tunnel/cmd/root/options"
	"github.com/bilalcaliskan/split-the-tunnel/internal/logging"
	"github.com/rs/zerolog"

	"github.com/bilalcaliskan/split-the-tunnel/internal/version"
	"github.com/spf13/cobra"
)

const socketPath = "/tmp/mydaemon.sock"

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
		err := initIPC(socketPath, logger)
		if err != nil {
			logger.Error().Err(err).Msg("failed to initialize IPC")
			return err
		}

		logger.Info().Str("socketPath", socketPath).Msg("IPC is initialized")

		defer func() {
			if err := cleanup(socketPath); err != nil {
				logger.Error().Err(err).Msg("failed to cleanup IPC")
			}
		}()

		logger.Info().Msg("daemon is running...")

		// Wait for termination signal
		<-sigs

		logger.Info().Msg("termination signal received")
		logger.Info().Msg("shutting down daemon...")

		return nil
	},
}

func cleanup(path string) error {
	// Perform any cleanup and shutdown tasks here
	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func initIPC(path string, logger zerolog.Logger) error {
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

/*func getDefaultNonVPNGateway() (string, error) {
	file, err := os.Open("/proc/net/route")
	if err != nil {
		return "", fmt.Errorf("failed to open routing info: %w", err)
	}
	defer file.Close()

	var bestGateway string
	highestMetric := -1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) >= 8 && fields[1] == "00000000" {
			metric, err := strconv.Atoi(fields[6])
			if err != nil {
				continue // Ignore lines with invalid metric
			}

			// Looking for the highest metric, assuming it's non-VPN
			if metric > highestMetric {
				highestMetric = metric

				bestGateway, err = parseHexIP(fields[2])
				if err != nil {
					continue // Ignore lines with invalid gateway IPs
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	if bestGateway == "" {
		return "", fmt.Errorf("non-VPN gateway not found")
	}

	return bestGateway, nil
}

func parseHexIP(hexStr string) (string, error) {
	ipBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", fmt.Errorf("failed to decode hex string: %w", err)
	}

	if len(ipBytes) != 4 {
		return "", fmt.Errorf("invalid IP length: %d", len(ipBytes))
	}

	// Reverse the byte order (little endian)
	for i, j := 0, len(ipBytes)-1; i < j; i, j = i+1, j-1 {
		ipBytes[i], ipBytes[j] = ipBytes[j], ipBytes[i]
	}

	return fmt.Sprintf("%d.%d.%d.%d", ipBytes[0], ipBytes[1], ipBytes[2], ipBytes[3]), nil
}

func resolveDomain(domain string) ([]string, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, err
	}

	var ipStrings []string
	for _, ip := range ips {
		ipStrings = append(ipStrings, ip.String())
	}

	return ipStrings, nil
}

func addRoute(ip, gateway string) error {
	cmd := exec.Command("sudo", "ip", "route", "add", ip, "via", gateway)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to add route for %s: %w", ip, err)
	}
	return nil
}*/
