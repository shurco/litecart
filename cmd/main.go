package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	app "github.com/shurco/litecart/internal"
)

var (
	gitCommit = "00000000"
	version   = "0.0.1"
	buildDate = "14.07.2023"
)

func main() {
	var serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Starts the web server (default to 127.0.0.1:8080)",
		Run: func(cmd *cobra.Command, args []string) {
			if err := app.NewApp(); err != nil {
				os.Exit(1)
			}
		},
	}

	var rootCmd = &cobra.Command{
		Use:   "litecart",
		Short: "LiteCart CLI",
		Long:  "Open Source realtime cart in 1 file",
		FParseErrWhitelist: cobra.FParseErrWhitelist{
			UnknownFlags: true,
		},
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		Version: fmt.Sprintf("LiteCart v%s (%s) from %s", version, gitCommit, buildDate),
	}
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
	
	rootCmd.AddCommand(serveCmd)
	rootCmd.Execute()
}
