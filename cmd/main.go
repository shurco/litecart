package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	app "github.com/shurco/litecart/internal"
)

var (
	version   = "0.0.1"
	gitCommit = "00000000"
	buildDate = "14.07.2023"
)

var (
	devMode  bool
	proxyMod bool
)

var rootCmd = &cobra.Command{
	Use:                "litecart",
	Short:              "LiteCart CLI",
	Long:               "🛒 litecart - shopping-cart in 1 file",
	Version:            fmt.Sprintf("LiteCart v%s (%s) from %s", version, gitCommit, buildDate),
	FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
	CompletionOptions:  cobra.CompletionOptions{DisableDefaultCmd: true},
}

func main() {
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	rootCmd.AddCommand(cmdServe())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func cmdServe() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve [flags]",
		Short: "Starts the web server (default to 127.0.0.1:8080)",
		Run: func(serveCmd *cobra.Command, args []string) {
			if err := app.NewApp(proxyMod, devMode); err != nil {
				os.Exit(1)
			}
		},
	}
	cmd.PersistentFlags().BoolVar(&proxyMod, "proxy", false, "proxy mode")
	cmd.PersistentFlags().BoolVar(&devMode, "dev", false, "develop mode")
	cmd.PersistentFlags().MarkHidden("dev")

	return cmd
}
