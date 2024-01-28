package remove

import (
	"fmt"
	"strings"

	"github.com/bilalcaliskan/split-the-tunnel/cmd/cli/utils"
	"github.com/bilalcaliskan/split-the-tunnel/internal/logging"

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
		logger := logging.GetLogger()

		argsStr := strings.Join(args, " ")

		logger.Info().Str("args", argsStr).Msg("remove called")

		req := fmt.Sprintf("%s %s", cmd.Name(), argsStr)
		res, err := utils.SendCommandToDaemon(utils.SocketPath, req)
		if err != nil {
			logger.Error().Err(err).Msg("error sending command to daemon")
			return err
		}

		logger.Info().Str("command", req).Str("response", res).Msg("successfully processed command")
		return nil
	},
}
