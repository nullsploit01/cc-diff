package cmd

import (
	"context"
	"os"
	"time"

	"github.com/nullsploit01/cc-diff/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ccdiff <file1> <file2>",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		done := make(chan struct{})

		if len(args) != 2 {
			cmd.Usage()
			os.Exit(1)
		}

		go func() {
			d := internal.NewDiff(cmd)
			err := d.FindFileDiff(args[0], args[1])
			if err != nil {
				cmd.OutOrStdout().Write([]byte(err.Error()))
				os.Exit(1)
			}
			close(done)
		}()

		// Wait for completion or timeout
		select {
		case <-done:
		case <-ctx.Done():
			cmd.OutOrStdout().Write([]byte("Process took too long, exiting\n"))
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
