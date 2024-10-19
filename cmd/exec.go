package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wim-vdw/az-secret2env/internal"
)

var execCmd = &cobra.Command{
	Use:     "exec",
	Short:   "Execute a program with secrets injected into environment variables",
	Long:    "Execute a program with secrets injected into environment variables.",
	Aliases: []string{"run", "start", "launch", "invoke"},
	Args:    cobra.MatchAll(cobra.MinimumNArgs(1)),
	RunE:    execProgram,
}

func init() {
	execCmd.Flags().BoolP("show-status", "s", false, "Show status during conversion of environment variables.")
	_ = viper.BindPFlag("show-status", execCmd.Flags().Lookup("show-status"))
	rootCmd.AddCommand(execCmd)
}

func execProgram(cmd *cobra.Command, args []string) error {
	filename := viper.GetString("env-file")
	verboseError := viper.GetBool("verbose")
	showStatus := viper.GetBool("show-status")
	client := internal.NewClient(filename)
	if err := client.LoadEnvs(verboseError); err != nil {
		return err
	}
	if err := client.ConvertSecrets(verboseError, showStatus); err != nil {
		return err
	}
	var runner *exec.Cmd
	if len(args) == 1 {
		runner = exec.Command(args[0])
	} else {
		runner = exec.Command(args[0], args[1:]...)
	}
	runner.Env = os.Environ()
	runner.Stdout = os.Stdout
	runner.Stderr = os.Stderr
	runner.Stdin = os.Stdin
	err := runner.Run()
	if err != nil {
		if verboseError {
			return fmt.Errorf("could not execute program\n%s", err)
		}
		return fmt.Errorf("could not execute program (use --verbose switch for more info)")
	}
	return nil
}
