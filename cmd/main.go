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

var devMode bool

var rootCmd = &cobra.Command{
	Use:                "litecart",
	Short:              "LiteCart CLI",
	Long:               "Open Source realtime cart in 1 file",
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
	rootCmd.AddCommand(cmdTheme())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func cmdServe() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve [flags]",
		Short: "Starts the web server (default to 127.0.0.1:8080)",
		Run: func(serveCmd *cobra.Command, args []string) {
			if err := app.NewApp(devMode); err != nil {
				os.Exit(1)
			}
		},
	}
	cmd.PersistentFlags().BoolVar(&devMode, "dev", false, "develop mode")
	cmd.PersistentFlags().MarkHidden("dev")

	return cmd
}

func cmdTheme() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "theme",
		Short: "Create default theme folder",
		Run: func(themeCmd *cobra.Command, args []string) {
			if err := app.NewTheme(); err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
		},
	}

	return cmd
}
