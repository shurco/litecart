package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	app "github.com/shurco/litecart/internal"
	"github.com/shurco/litecart/pkg/update"
)

var (
	version   = "v0.0.1"
	gitCommit = "00000000"
	buildDate = "14.07.2023"
)

var rootCmd = &cobra.Command{
	Use:                "litecart",
	Short:              "LiteCart CLI",
	Long:               "🛒 litecart - shopping-cart in 1 file",
	Version:            fmt.Sprintf("LiteCart %s (%s) from %s", version, gitCommit, buildDate),
	FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
	CompletionOptions:  cobra.CompletionOptions{DisableDefaultCmd: true},
}

func main() {
	update.SetVersion(&update.Version{
		CurrentVersion: version,
		GitCommit:      gitCommit,
		BuildDate:      buildDate,
	})

	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	rootCmd.AddCommand(cmdInit())
	rootCmd.AddCommand(cmdServe())
	rootCmd.AddCommand(cmdUpdate())
	rootCmd.AddCommand(cmdMigrate())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func cmdServe() *cobra.Command {
	var noSite, devMode bool
	var httpAddr, httpsAddr string
	cmd := &cobra.Command{
		Use:   "serve [flags]",
		Short: "Starts the web server (default to 0.0.0.0:8080)",
		Run: func(serveCmd *cobra.Command, args []string) {
			if err := app.NewApp(httpAddr, httpsAddr, noSite, devMode); err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
		},
	}

	cmd.PersistentFlags().StringVar(
		&httpAddr,
		"http",
		"0.0.0.0:8080",
		"server address",
	)

	// Ports <= 1024 are privileged ports. You can't use them unless you're root or have the explicit
	// permission to use them. See this answer for an explanation or wikipedia or something you trust more.
	// sudo setcap 'cap_net_bind_service=+ep' /opt/yourGoBinary
	cmd.PersistentFlags().StringVar(
		&httpsAddr,
		"https",
		"",
		"https server address (auto TLS)",
	)

	cmd.PersistentFlags().BoolVar(&noSite, "no-site", false, "disable create site")

	cmd.PersistentFlags().BoolVar(&devMode, "dev", false, "develop mode")
	cmd.PersistentFlags().MarkHidden("dev")

	return cmd
}

func cmdInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Creating the basic structure",
		Run: func(serveCmd *cobra.Command, args []string) {
			app.Init()
		},
	}

	return cmd
}

func cmdUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Updating the application to the latest version",
		Run: func(serveCmd *cobra.Command, args []string) {
			cfg := &update.Config{
				Owner:             "shurco",
				Repo:              "litecart",
				CurrentVersion:    version,
				ArchiveExecutable: "litecart",
			}
			err := update.Init(cfg)
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
		},
	}

	return cmd
}

func cmdMigrate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate on the latest version of database schema",
		Run: func(serveCmd *cobra.Command, args []string) {
			if err := app.Migrate(); err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
		},
	}

	return cmd
}
