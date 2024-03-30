/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "Go Quiz!",
	Short:         "Random quiz questions about random things",
	Long:          ``,
	SilenceErrors: true,
	SilenceUsage:  true,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(setNameCmd)
	rootCmd.AddCommand(startQuizCmd)
	rootCmd.Flags().BoolP("toggle", "t", true, "Help message for toggle")
}
