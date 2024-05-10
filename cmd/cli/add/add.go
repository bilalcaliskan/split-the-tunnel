package add

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bilalcaliskan/split-the-tunnel/cmd/cli/utils"
	pb "github.com/bilalcaliskan/split-the-tunnel/pkg/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
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
		// Set up a connection to the server.
		conn, err := grpc.Dial("localhost:50051", grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewRouteManagerClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := c.AddRoute(ctx, &pb.AddRouteRequest{Destination: "example.com"})
		if err != nil {
			log.Fatalf("could not add route: %v", err)
		}
		fmt.Printf("AddRoute Response: %v\n", r)

		// Handle the business error
		if r.GetError() != nil {
			fmt.Printf("Business Error: %v\n", r.GetError().String())
		}

		return nil
		//logger := cmd.Context().Value(constants.LoggerKey{}).(zerolog.Logger)
		//socketPath := cmd.Context().Value(constants.SocketPathKey{}).(string)
		//
		//logger.Info().
		//	Str("operation", cmd.Name()).
		//	Any("args", args).
		//	Msg(constants.ProcessCommand)
		//
		//for _, arg := range args {
		//	req := fmt.Sprintf("%s %s", cmd.Name(), arg)
		//	res, err := utils.SendCommandToDaemon(socketPath, req)
		//	if err != nil {
		//		logger.Error().Str("command", req).Err(err).Msg(constants.FailedToProcessCommand)
		//		continue
		//	}
		//
		//	logger.Info().Str("command", req).Str("response", res).Msg(constants.SuccessfullyProcessed)
		//}
		//
		//return nil
	},
}
