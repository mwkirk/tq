/*
Copyright © 2024 Mark Kirk
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

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

		fmt.Printf("sleep job of %d secs\n", sec)
	},
}

func init() {
	submitCmd.AddCommand(sleepCmd)
}

// validate args are all integers
func integerArgs(cmd *cobra.Command, args []string) error {
	for _, v := range args {
		_, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
	}
	return nil
}
