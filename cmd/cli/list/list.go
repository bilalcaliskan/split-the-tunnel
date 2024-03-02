package list

import (
	"fmt"
	"strings"

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

		logger.Info().
			Str("operation", cmd.Name()).
			Msg(constants.ProcessCommand)

		res, err := utils.SendCommandToDaemon(utils.SocketPath, cmd.Name())
		if err != nil {
			logger.Error().Str("command", cmd.Name()).Err(err).Msg(constants.FailedToProcessCommand)

			return &utils.CommandError{Err: err, Code: 10}
		}

		logger.Info().Str("command", cmd.Name()).Msg(constants.SuccessfullyProcessed)

		stateMsg := fmt.Sprintf(`|---------------------------------------------------------------------|
|   here is your state:                                               |
|                                                                     |
|   %s   |
|---------------------------------------------------------------------|
`, strings.TrimSpace(res))

		fmt.Print(stateMsg)

		return nil
	},
}
