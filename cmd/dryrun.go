package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wim-vdw/az-secret2env/internal"
)

var dryrunCmd = &cobra.Command{
	Use:     "dry-run",
	Short:   "Simulate secret retrieval without running a program.",
	Long:    "Simulate secret retrieval without running a program.",
	Aliases: []string{"test", "check", "preview"},
	RunE:    dryrun,
}

func init() {
	rootCmd.AddCommand(dryrunCmd)
}

func dryrun(cmd *cobra.Command, args []string) error {
	filename := viper.GetString("env-file")
	verboseError := viper.GetBool("verbose")
	client := internal.NewClient(filename)
	if err := client.LoadEnvs(verboseError); err != nil {
		return err
	}
	if err := client.ConvertSecrets(verboseError, true); err != nil {
		return err
	}
	client.PrintDryRunResults()
	return nil
}
