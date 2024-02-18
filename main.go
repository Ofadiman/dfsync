package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
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

var logger = log.NewWithOptions(
	os.Stdout, log.Options{
		ReportTimestamp: false,
	},
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
		logger.Infof("Files will be synchronized from " + from.Value.String() + " directory to " + to.Value.String() + " directory")
	},
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		logger.Fatal(err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		logger.Fatal(err)
	}

	rootCmd.PersistentFlags().StringP(FLAG_FROM, "f", cwd, "directory from which symlinks will be created")
	rootCmd.PersistentFlags().StringP(FLAG_TO, "t", home, "directory to which symlinks will be created")

	if err := rootCmd.Execute(); err != nil {
		logger.Fatal(err)
	}
}
