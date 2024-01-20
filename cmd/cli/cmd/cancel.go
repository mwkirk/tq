/*
Copyright © 2024 Mark Kirk

*/
package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

// cancelCmd represents the cancel command
var cancelCmd = &cobra.Command{
    Use:   "cancel [job number]",
    Short: "Cancel a job",
    Long: `Cancel a job in the queue. The job is removed from the queue if still waiting 
to be dispatched and killed on the worker if running'.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("cancel called")
    },
}

func init() {
    rootCmd.AddCommand(cancelCmd)

    // Here you will define your flags and configuration settings.

    // Cobra supports Persistent Flags which will work for this command
    // and all subcommands, e.g.:
    // cancelCmd.PersistentFlags().String("foo", "", "A help for foo")

    // Cobra supports local flags which will only run when this command
    // is called directly, e.g.:
    // cancelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
