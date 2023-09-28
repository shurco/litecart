package app

import (
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/migrations"
	"github.com/shurco/litecart/pkg/fsutil"
)

// Init is ...
func Init() error {
	dirsToCheck := []struct {
		path string
		name string
	}{
		{"./lc_uploads", "lc_uploads"},
		{"./lc_digitals", "lc_digitals"},
	}

	for _, dir := range dirsToCheck {
		if err := fsutil.MkDirs(0775, dir.path); err != nil {
			log.Err(err).Send()
			return err
		}
	}

	if _, err := queries.InitDB("./lc_base/data.db", migrations.Embed()); err != nil {
		log.Err(err).Send()
		return err
	}

	return nil
}
