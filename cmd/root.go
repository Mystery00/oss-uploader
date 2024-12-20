package cmd

import (
	"github.com/spf13/cobra"
	"go-cli-template/env"
	"log/slog"
	"os"
)

var rootCmd = &cobra.Command{
	Use:  "go-cli",
	Long: "go-cli-template",
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&env.Debug, "debug", "D", false, "debug mode")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		slog.Error("detect error:", err.Error())
		os.Exit(1)
	}
}
