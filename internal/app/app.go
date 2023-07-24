package app

import (
	"embed"
	"os"

	"github.com/rs/zerolog"

	"github.com/shurco/litecart/internal/app/queries"
	"github.com/shurco/litecart/internal/core"
	"github.com/shurco/litecart/pkg/fsutil"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

// Base is ...
type Base struct {
	core.Core
	DB *queries.Base
}

// NewApp is ...
func NewApp(flags core.Flags) error {
	app := &Base{}
	app.DevMode = true
	app.Log = zerolog.New(os.Stderr).With().Timestamp().Logger()

	if flags.Serve {
		if err := fsutil.MkDirs(0775, "./uploads"); err != nil {
			app.Log.Err(err).Send()
			return err
		}

		// init database
		sqlite, err := app.InitDB("./lc_base/data.db", embedMigrations)
		if err != nil {
			app.Log.Err(err).Send()
			return err
		}

		app.DB = &queries.Base{
			AuthQueries:    queries.AuthQueries{DB: sqlite},
			InstallQueries: queries.InstallQueries{DB: sqlite},
			SettingQueries: queries.SettingQueries{DB: sqlite},
		}

		// init jwt settings
		app.JWT, err = app.DB.SettingJWT()
		if err != nil {
			app.Log.Err(err).Send()
			return err
		}

		// init stripe settings
		app.Stripe, err = app.DB.SettingStripe()
		if err != nil {
			app.Log.Err(err).Send()
			return err
		}
		if app.Stripe.SecretKey != "" {
			app.Stripe.Client = core.InitStripeClient(app.Stripe.SecretKey)
		}

		// web web server
		if err := app.initWebServer(8080); err != nil {
			app.Log.Err(err).Send()
			return err
		}
	}

	return nil
}
