/*
Copyright © 2024 Mark Kirk
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"
	"tq/pb"

	"github.com/spf13/cobra"
)

var failRate float64

var sleepCmd = &cobra.Command{
	Use:   "sleep [seconds]",
	Short: "Submit a sleep job",
	Long:  `Submit a job that will be queued and sleep for the given number of seconds when run. Useful for testing.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1), integerArgs),
	Run: func(cmd *cobra.Command, args []string) {
		sec, err := strconv.Atoi(args[0])
		if err != nil {
			// Should always be an int since validated
			log.Fatalf("sleep duration was not an integer")
		}
		j := pb.JobSpec{
			Kind:   pb.JobKind_JOB_KIND_SLEEP,
			JobNum: 0,
			Name:   "sleep " + args[0],
			Parms: map[string]string{
				"duration": args[0],
				"failRate": fmt.Sprintf("%f", failRate),
			},
		}
		fmt.Printf("sleep job of %d secs\n", sec)
		submit(j)
	},
}

func init() {
	submitCmd.AddCommand(sleepCmd)
	sleepCmd.Flags().Float64VarP(&failRate, "fail", "f", 0.0, "chance of job failure")
}

// validate args are all integers
func integerArgs(cmd *cobra.Command, args []string) error {
	for _, v := range args {
		_, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("argument must be an integer")
		}
	}
	return nil
}
