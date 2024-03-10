package list

import (
	"os"
	"strings"

	"github.com/bilalcaliskan/split-the-tunnel/internal/state"
	"github.com/olekukonko/tablewriter"

	"github.com/bilalcaliskan/split-the-tunnel/cmd/cli/utils"
	"github.com/bilalcaliskan/split-the-tunnel/internal/constants"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

// ListCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return utils.ErrTooManyArgs
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := cmd.Context().Value(constants.LoggerKey{}).(zerolog.Logger)
		socketPath := cmd.Context().Value(constants.SocketPathKey{}).(string)

		logger.Info().
			Str("operation", cmd.Name()).
			Msg(constants.ProcessCommand)

		res, err := utils.SendCommandToDaemon(socketPath, cmd.Name())
		if err != nil {
			logger.Error().Str("command", cmd.Name()).Err(err).Msg(constants.FailedToProcessCommand)

			return &utils.CommandError{Err: err, Code: 10}
		}

		logger.Info().Str("command", cmd.Name()).Msg(constants.SuccessfullyProcessed)

		domains, err := state.FromStringSlice(res)
		if err != nil {
			logger.Error().Err(err).Msg("failed to parse response")

			return &utils.CommandError{Err: err, Code: 11}
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Domain", "Gateway", "IPs"})
		// Set the Alignment for each column to center
		table.SetColumnAlignment([]int{tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER})
		table.SetBorder(true)  // Set to false if you do not want borders
		table.SetRowLine(true) // Enable row line for more clarity
		table.SetAlignment(tablewriter.ALIGN_CENTER)

		for _, info := range domains {
			table.Append([]string{info.Domain, info.Gateway, strings.Join(info.ResolvedIPs, "\n")})
		}

		table.Render() // Send output

		return nil
	},
}
