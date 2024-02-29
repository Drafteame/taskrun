package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Drafteame/taskrun/internal"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print cli version",
	Long:  "Print the version of the cli",
	Run:   printVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func printVersion(cmd *cobra.Command, args []string) {
	_, _ = fmt.Printf("Version: %s\n", internal.Version)
}
