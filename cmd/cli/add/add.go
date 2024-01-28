package add

import (
	"fmt"

	"github.com/bilalcaliskan/split-the-tunnel/cmd/cli/utils"
	"github.com/bilalcaliskan/split-the-tunnel/internal/logging"
	"github.com/spf13/cobra"
)

// AddCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
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
		logger := logging.GetLogger()

		logger.Info().Any("args", args).Msg("add called")

		for _, arg := range args {
			req := fmt.Sprintf("%s %s", cmd.Name(), arg)
			res, err := utils.SendCommandToDaemon(utils.SocketPath, req)
			if err != nil {
				logger.Error().Str("command", req).Err(err).Msg("error sending command to daemon")
				continue
			}

			logger.Info().Str("command", req).Str("response", res).Msg("successfully processed command")
		}

		return nil
	},
}
