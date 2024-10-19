package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wim-vdw/az-secret2env/internal"
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
	filename := viper.GetString("env-file")
	client := internal.NewClient(filename)
	if err := client.LoadEnvs(); err != nil {
		return err
	}
	if err := client.ConvertSecrets(); err != nil {
		return err
	}
	client.PrintDryRunResults()
	return nil
}
