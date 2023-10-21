package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/wwmoraes/site/cmd/site/data"
	"github.com/wwmoraes/site/cmd/site/radar"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "site",
		Short: "updates site data",
		Long: "collection of commands to fetch external information and format " +
			"it as data usable by Hugo static site builds",
		Example: "site",
	}

	rootCmd.AddCommand(data.Command())
	rootCmd.AddCommand(radar.Command())

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
