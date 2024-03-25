/*
Copyright © 2024 Mark Kirk
*/
package cmd

import (
	"fmt"
	"tq/pb"

	"github.com/spf13/cobra"
)

var transcodeCmd = &cobra.Command{
	Use:   "transcode",
	Short: "Submit a transcode job",
	Long: `Submit an FFMPEG transcode job to the queue. The job will transcode the input file to 
the output file using the key info file. Note: file paths must be accessible to the worker.`,
	Run: func(cmd *cobra.Command, args []string) {
		parms := map[string]string{
			"inputPath":   cmd.Flag("inputPath").Value.String(),
			"outputPath":  cmd.Flag("outputPath").Value.String(),
			"keyInfoPath": cmd.Flag("keyInfoPath").Value.String(),
		}
		j := pb.JobSpec{
			Kind:   pb.JobKind_JOB_KIND_FFMPEG,
			JobNum: 0,
			Name:   "transcode",
			Parms:  parms,
		}
		fmt.Printf("transcode job with %v\n", parms)
		submit(j)
	},
}

func init() {
	submitCmd.AddCommand(transcodeCmd)
	transcodeCmd.PersistentFlags().StringP("inputPath", "i", "", "input media path")
	transcodeCmd.PersistentFlags().StringP("outputPath", "o", "", "output media path")
	transcodeCmd.PersistentFlags().StringP("keyInfoPath", "k", "", "key info path")
}
