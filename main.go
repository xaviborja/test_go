package main

import (
	"github.com/spf13/cobra"
	"github.com/xaviborja/test_go/cli"
)

func main() {
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd())
	rootCmd.Execute()
}
