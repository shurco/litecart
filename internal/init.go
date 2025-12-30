package app

import (
	"github.com/shurco/litecart/internal/base"
	"github.com/shurco/litecart/migrations"
	"github.com/shurco/litecart/pkg/fsutil"
)

const (
	dbPath = "./lc_base/data.db"
)

var (
	requiredDirs = []string{"./lc_uploads", "./lc_digitals"}
)

// Init initializes the directory structure and database
func Init() error {
	for _, dir := range requiredDirs {
		if err := fsutil.MkDirs(0o775, dir); err != nil {
			if log != nil {
				log.Err(err).Send()
			}
			return err
		}
	}

	if _, err := base.New(dbPath, migrations.Embed()); err != nil {
		if log != nil {
			log.Err(err).Send()
		}
		return err
	}

	return nil
}

// Migrate performs database migrations
func Migrate() error {
	return base.Migrate(dbPath, migrations.Embed())
}
