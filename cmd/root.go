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
	Short: "A CLI tool to compare and display differences between files.",
	Long: `A command-line utility for comparing files and highlighting the differences 
in their content. This tool is ideal for tracking changes between versions of text 
files, configuration files, or source code.

Features:
- Identify added, removed, and unchanged lines between files.
- Output results in a clear and readable format.
- Simple and intuitive command-line interface.

Examples:
- Compare two files:
  $ diff-tool file1.txt file2.txt

This tool leverages the Cobra library to provide a reliable and efficient experience for developers and users alike.`,
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
