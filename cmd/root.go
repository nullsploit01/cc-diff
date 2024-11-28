package cmd

import (
	"os"

	"github.com/nullsploit01/cc-diff/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cc-diff",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		a := `Coding Challenges helps you become a better software engineer through that build real applications.
I share a weekly coding challenge aimed at helping software engineers level up their skills through deliberate practice.
I’ve used or am using these coding challenges as exercise to learn a new programming language or technology.
Each challenge will have you writing a full application or tool. Most of which will be based on real world tools and utilities.`

		b := `Helping you become a better software engineer through coding challenges that build real applications.
I share a weekly coding challenge aimed at helping software engineers level up their skills through deliberate practice.
These are challenges that I’ve used or am using as exercises to learn a new programming language or technology.
Each challenge will have you writing a full application or tool. Most of which will be based on real world tools and utilities.`

		d := internal.NewDiff(cmd)
		d.FindLineDiff(a, b)
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
