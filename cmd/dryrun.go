package cmd

import (
	"github.com/spf13/cobra"
)

var dryrunCmd = &cobra.Command{
	Use:     "dry-run",
	Short:   "Simulate secret retrieval without running a program",
	Long:    "Simulate secret retrieval without running a program.",
	Aliases: []string{"test", "validate", "check"},
	RunE:    dryrun,
}

func init() {
	rootCmd.AddCommand(dryrunCmd)
}

func dryrun(cmd *cobra.Command, args []string) error {
	return nil
}
