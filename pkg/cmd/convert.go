package cmd

import (
	"fmt"
	"os"

	"github.com/kimba74/goat/pkg/goat"
	"github.com/spf13/cobra"
)

var (
	convertCmd = &cobra.Command{
		Use:   "convert [file]",
		Short: "Convert the content of a file to another format",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// Decode the values of the provided files
			val, err := goat.ParseValuesFile(args[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "convert command: %s\n", err)
				os.Exit(1)
			}

			// Decode the merge results in the provided format
			if err = goat.Encode(fConvertOutput, val); err != nil {
				fmt.Fprintf(os.Stderr, "convert command: %s\n", err)
				os.Exit(1)
			}
		},
	}
	fConvertOutput string
)

func init() {
	convertCmd.Flags().StringVarP(&fConvertOutput, "output", "o", "yaml", "output format. can be json or yaml")
	rootCmd.AddCommand(convertCmd)
}
