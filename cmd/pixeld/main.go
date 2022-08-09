package main

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	Node = "node"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "pixel",
		Short: "pixel ",
	}

	rootCmd.AddCommand(
		QueryCmd(),
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
