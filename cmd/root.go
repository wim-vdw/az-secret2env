package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "az-secret2env",
	Short:   "Execute a program with environment variables temporarily populated by Azure Key Vault secrets.",
	Version: "v1.0.0-alpha",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Display this help message.")
	rootCmd.Flags().BoolP("version", "v", false, "Display version info.")
	rootCmd.SetVersionTemplate("{{ .Version }}\n")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SilenceUsage = true
}
