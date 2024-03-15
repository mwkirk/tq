/*
Copyright © 2024 Mark Kirk
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"time"
	"tq/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List jobs",
	Long:  `List waiting, running, and completed jobs`,
	Run: func(cmd *cobra.Command, args []string) {
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

	jobFilter := pb.JobFilter{
		JobStateFilter: int32(pb.JobState_JOB_STATE_ALL),
		JobKindFilter:  int32(pb.JobKind_JOB_KIND_TEST | pb.JobKind_JOB_KIND_SLEEP | pb.JobKind_JOB_KIND_FFMPEG),
		JobNums:        nil,
	}
	options := pb.ListOptions{JobFilter: &jobFilter}

	lr, err := c.List(ctx, &pb.ListRequest{Options: &options})
	if err != nil {
		log.Fatalf("failed to list jobs: %s\n", err)
	}

	jobFmt := "%-9d %-6s %-6s %3.0f%% %-10.10s %-20.20s %-30.30s\n"
	headerFmt := "%-9s %-6s %-6s %-3s %-10s %-20.20s %-30.30s\n"
	header := fmt.Sprintf(headerFmt, "job", "state", "kind", "prog", "name", "worker", "msgs")
	fmt.Printf("%s", header)

	printJobs := func(l *pb.JobList) {
		for _, v := range l.Items {
			fmt.Printf(jobFmt, v.JobNum, v.JobState.ShortDesc(), v.Kind.ShortDesc(), v.Progress*100, v.Name, v.Worker,
				v.Msg)
		}
	}

	printJobs(lr.Result.Run)
	printJobs(lr.Result.Wait)
	printJobs(lr.Result.Done)
}
