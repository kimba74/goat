package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "goat",
		Short:   "Go Apply Templates (GoAT) is a tool to merge and convert YAML and JSON values files and apply them to Go templates.",
		Version: "v0.1.0-alpha1",
	}
)

// Execute the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
