package cmd

import (
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:     "exec",
	Short:   "Execute a program with secrets injected into environment variables",
	Long:    "Execute a program with secrets injected into environment variables.",
	Aliases: []string{"run", "start", "launch", "invoke"},
	RunE:    execProgram,
}

func init() {
	rootCmd.AddCommand(execCmd)
}

func execProgram(cmd *cobra.Command, args []string) error {
	return nil
}
