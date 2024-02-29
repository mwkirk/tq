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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List jobs",
	Long:  `List waiting, running, and completed jobs`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		list()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTqJobClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	lr, err := c.List(ctx, &pb.ListRequest{})
	if err != nil {
		log.Fatalf("failed to list jobs: %s\n", err)
	}
	fmt.Printf("listing jobs %v\n", lr.JobStatus)
}
