package remove

import (
	"fmt"

	"github.com/rs/zerolog"

	"github.com/bilalcaliskan/split-the-tunnel/internal/constants"

	"github.com/bilalcaliskan/split-the-tunnel/cmd/cli/utils"
	"github.com/spf13/cobra"
)

// RemoveCmd represents the remove command
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return utils.ErrNoArgs
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := cmd.Context().Value(constants.LoggerKey{}).(zerolog.Logger)

		logger.Info().
			Str("operation", cmd.Name()).
			Any("args", args).
			Msg(constants.ProcessCommand)

		for _, arg := range args {
			req := fmt.Sprintf("%s %s", cmd.Name(), arg)
			res, err := utils.SendCommandToDaemon(utils.SocketPath, req)
			if err != nil {
				logger.Error().
					Str("command", req).
					Err(err).
					Msg(constants.FailedToProcessCommand)

				continue
			}

			logger.Info().
				Str("command", req).
				Str("response", res).
				Msg(constants.SuccessfullyProcessed)
		}

		return nil
	},
}
