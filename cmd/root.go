/*
Copyright © 2026 Codurance
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-test-skeleton",
	Short: "A Codurance Go test skeleton",
	Long: `A small Go service skeleton used as the starting point for the
Codurance Go competency exercise. Run "go-test-skeleton serve" to start the
API server.`,
	SilenceUsage: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
