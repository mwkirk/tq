/*
Copyright © 2024 Mark Kirk
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strconv"
	"time"
	"tq/internal/model"
	"tq/pb"
)

// cancelCmd represents the cancel command
var cancelCmd = &cobra.Command{
	Use:   "cancel [job number]",
	Short: "Cancel a job",
	Long: `Cancel a job in the queue. The job is removed from the queue if still waiting 
to be dispatched and killed on the worker if running'.`,
	Args: cobra.MatchAll(cobra.ExactArgs(1), integerArgs),
	Run: func(cmd *cobra.Command, args []string) {
		n, err := strconv.Atoi(args[0])
		if err != nil {
			// Should always be an int since validated
			log.Fatalf("job number was not an integer")
		}

		fmt.Printf("cancel job %d \n", n)
		cancel(model.JobNumber(n))
	},
}

func init() {
	rootCmd.AddCommand(cancelCmd)
}

func cancel(jobNum model.JobNumber) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTqJobClient(conn)

	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second*120)
	defer ctxCancel()

	options := &pb.CancelOptions{JobNum: uint32(jobNum)}
	cr, err := c.Cancel(ctx, &pb.CancelRequest{Options: options})
	if err != nil {
		log.Fatalf("failed to cancel job: %s\n", err)
	}

	fmt.Printf("canceled job %v\n", cr.Result)
}
