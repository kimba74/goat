package cmd

import (
	"fmt"
	"os"

	"github.com/kimba74/goat/pkg/goat"
	"github.com/spf13/cobra"
)

var (
	applyCmd = &cobra.Command{
		Use:   "apply [template-file] [value-file]",
		Short: "Apply values to a Go Template",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			// Check if Stdin placeholder was used more than once
			err := goat.CheckMoreThanOnce(args, "-")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Command Apply: can't use Stdin placeholder more than once: %v\n", err)
				os.Exit(1)
			}

			// Load template from first argument
			tmpl, err := goat.ParseTemplateFile(args[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Command Apply: %s\n", err)
				os.Exit(1)
			}

			// Load values from specified files
			vals, err := goat.ParseValuesFile(args[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Command Apply: %s\n", err)
				os.Exit(1)
			}

			// Apply template
			goat.Apply(tmpl, vals)
		},
	}
	fApplyAppend   bool
	fApplyDeepCopy bool
	fApplyOverride bool
)

func init() {
	rootCmd.AddCommand(applyCmd)
}
