package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/wwmoraes/site/cmd/site/cloudflare"
	"github.com/wwmoraes/site/cmd/site/content"
	"github.com/wwmoraes/site/cmd/site/image"
	"github.com/wwmoraes/site/cmd/site/radar"
	"github.com/wwmoraes/site/cmd/site/vcard"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "site",
		Short: "updates site data",
		Long: "collection of commands to fetch external information and format " +
			"it as data usable by Hugo static site builds",
		Example: "site",
	}

	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)

	rootCmd.AddCommand(content.Command())
	rootCmd.AddCommand(cloudflare.Command())
	rootCmd.AddCommand(image.Command())
	rootCmd.AddCommand(radar.Command())
	rootCmd.AddCommand(vcard.Command())

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
