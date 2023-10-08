package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/wwmoraes/site/cmd/site/radar"
	"github.com/wwmoraes/site/cmd/site/update"
)

var rootCmd = &cobra.Command{
	Use:     "site",
	Short:   "updates site data",
	Long:    "collection of commands to fetch external information and format it as data usable by Hugo static site builds",
	Example: "site",
}

func init() {
	rootCmd.AddCommand(radar.New())
	rootCmd.AddCommand(update.New())
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
