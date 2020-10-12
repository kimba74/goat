package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "goat",
		Short:   "GoAT is a tool to apply values to Go templates",
		Long:    "Go Apply Templates (GoAT).....", //TODO: Long documentation
		Version: "0.1.0",
	}
)

// Execute the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
