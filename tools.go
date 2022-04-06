package main

import (
	"github.com/spf13/cobra"
	"os"
)

var mainCmd = &cobra.Command{Use: "tools"}

func main() {
	mainCmd.AddCommand(parseCid())
	mainCmd.AddCommand(bootstrap())
	if mainCmd.Execute() != nil {
		os.Exit(1)
	}
}
