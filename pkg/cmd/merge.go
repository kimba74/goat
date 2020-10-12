package cmd

import (
	"fmt"
	"os"

	"github.com/kimba74/goat/pkg/goat"
	"github.com/spf13/cobra"
)

var (
	mergeCmd = &cobra.Command{
		Use:   "merge [file]...",
		Short: "Merge the content of multiple files into one",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			// Parse the values of the provided files
			vals, err := goat.ParseValuesFiles(args...)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Command Merge: %s\n", err)
				os.Exit(1)
			}

			// Merge the values into one
			merger, err := goat.Merge(fMergeAppend, fMergeDeepCopy, fMergeOverride, vals...)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Command Merge: %s\n", err)
				os.Exit(1)
			}

			// Decode the merge results in the provided format
			if err = goat.Encode(fMergeOutput, merger); err != nil {
				fmt.Fprintf(os.Stderr, "Command Merge: %s\n", err)
				os.Exit(1)
			}
		},
	}
	fMergeAppend   bool
	fMergeDeepCopy bool
	fMergeOutput   string
	fMergeOverride bool
)

func init() {
	mergeCmd.Flags().BoolVar(&fMergeAppend, "append", false, "append arrays instead of replacing them")
	mergeCmd.Flags().BoolVar(&fMergeDeepCopy, "deepcopy", false, "merge array elements one by one with")
	mergeCmd.Flags().StringVarP(&fMergeOutput, "output", "o", "yaml", "output format. can be json or yaml")
	mergeCmd.Flags().BoolVar(&fMergeOverride, "override", false, "override non-empty destination attributes with non-empty source ones")
	rootCmd.AddCommand(mergeCmd)
}
