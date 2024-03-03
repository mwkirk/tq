/*
Copyright © 2024 Mark Kirk
*/
package cmd

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
	"tq/pb"

	"github.com/spf13/cobra"
)

// submitCmd represents the submit command
var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit a job to the queue",
	Long:  `Submit a job the queue.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("submit called")
	},
}

func init() {
	rootCmd.AddCommand(submitCmd)
}

func submit(jobSpec pb.JobSpec) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTqJobClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	sr, err := c.Submit(ctx, &pb.SubmitRequest{JobSpec: &jobSpec})
	if err != nil {
		log.Fatalf("failed to submit job: %s\n", err)
	}
	fmt.Printf("submitted job %d\n", sr.JobNum)
}
