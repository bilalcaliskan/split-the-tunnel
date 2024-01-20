package main

import (
	"os"

	"github.com/bilalcaliskan/split-the-tunnel/internal/version"
	"github.com/spf13/cobra"
)

var ver = version.Get()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "stt-cli",
	Short:   "",
	Long:    ``,
	Version: ver.GitVersion,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

//const socketPath = "/tmp/mydaemon.sock"

//func sendCommandToDaemon(command string) (string, error) {
//	conn, err := net.Dial("unix", socketPath)
//	if err != nil {
//		return "", err
//	}
//	defer conn.Close()
//
//	_, err = conn.Write([]byte(command + "\n"))
//	if err != nil {
//		return "", err
//	}
//
//	// If you expect a response from the daemon, read it here
//	// For example, using bufio.NewReader(conn).ReadString('\n')
//
//	return "Command sent successfully", nil
//}
