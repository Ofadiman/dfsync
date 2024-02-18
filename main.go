package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	source      string
	destination string
	force       string
	dry         bool
	version     bool
	verbose     bool
	ignore      string
)

const (
	FLAG_FROM = "from"
	FLAG_TO   = "to"
)

var rootCmd = &cobra.Command{
	Use:     "dfsync",
	Short:   "Synchronize dot files",
	Example: "dfsync --source ./src --destination ~/",
	Long:    `Dot files sync is a tool that allows you to painlessly synchronize dot files across multiple environments.`,
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		from := cmd.Flag(FLAG_FROM)
		to := cmd.Flag(FLAG_TO)
		fmt.Printf("Files will be synchronized from %v directory to %v directory.\n", from.Value.String(), to.Value.String())
	},
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	rootCmd.PersistentFlags().StringP(FLAG_FROM, "f", cwd, "directory from which symlinks will be created")
	rootCmd.PersistentFlags().StringP(FLAG_TO, "t", home, "directory to which symlinks will be created")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
