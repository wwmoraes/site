package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/wwmoraes/site/cmd/site/data"
	"github.com/wwmoraes/site/cmd/site/radar"
	"github.com/wwmoraes/site/cmd/site/update"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "site",
		Short: "updates site data",
		Long: "collection of commands to fetch external information and format " +
			"it as data usable by Hugo static site builds",
		Example: "site",
	}

	rootCmd.AddCommand(radar.New())
	rootCmd.AddCommand(update.New())
	rootCmd.AddCommand(data.Command())

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
