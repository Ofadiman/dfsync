package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "dfsync",
	Short:   "Synchronize dot files",
	Example: "dfsync --source ./src --destination ~/",
	Long:    `Dot files sync is a tool that allows you to painlessly synchronize dot files across multiple environments.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
