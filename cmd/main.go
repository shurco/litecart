package main

import (
	"log"
	"os"

	app "github.com/shurco/litecart/internal"
)

var (
	gitCommit = "00000000"
	version   = "0.0.1"
	buildDate = "14.07.2023"
)

func main() {
	flags := app.Flags{
		Serve: true,
	}

	if err := app.NewApp(flags); err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
	}

	/*
		rootCmd := &cobra.Command{
			Use:     "litecart",
			Short:   "Lightweight online store powered by stripe",
			Long:    "Lightweight online store powered by stripe",
			Args:    cobra.ArbitraryArgs,
			Version: fmt.Sprintf("%s (%s), %s", version, gitCommit, buildDate),
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) < 1 {
					if err := cmd.Usage(); err != nil {
						log.Fatal(err)
					}
					return
				}

				if err := app.NewApp(flags); err != nil {
					log.Printf("%+v", err)
					os.Exit(1)
				}
			},
		}

		pf := rootCmd.PersistentFlags()
		pf.BoolVarP(&flags.Serve, "serve", "s", false, "starts the web server (default to 127.0.0.1:8080)")

		if err := rootCmd.Execute(); err != nil {
			log.Fatal(err)
		}
	*/
}
