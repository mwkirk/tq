/*
Copyright © 2024 Mark Kirk
*/
package cmd

import (
	"fmt"

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
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// submitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// submitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
