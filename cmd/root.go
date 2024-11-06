package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:     "az-secret2env",
	Short:   "Execute a program with environment variables temporarily populated by Azure Key Vault secrets.",
	Version: "v1.0.0",
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
	rootCmd.PersistentFlags().StringP("env-file", "f", "", "Load additional environment variables from a specified file.")
	rootCmd.PersistentFlags().BoolP("verbose", "", false, "Enable verbose output for detailed error handling and diagnostics.")
	rootCmd.SetVersionTemplate("{{ .Version }}\n")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.SilenceUsage = true
	_ = viper.BindPFlag("env-file", rootCmd.PersistentFlags().Lookup("env-file"))
	_ = viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}
