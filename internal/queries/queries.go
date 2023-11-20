package queries

import (
	"database/sql"
	"embed"

	"github.com/shurco/litecart/internal/base"
	_ "modernc.org/sqlite"
)

var db *Base

// Base is ...
type Base struct {
	SettingQueries
	AuthQueries
	InstallQueries
	PageQueries
	ProductQueries
	CartQueries
}

// New is ...
func New(embed embed.FS) (err error) {
	// init database
	var sqlite *sql.DB
	sqlite, err = base.New("./lc_base/data.db", embed)
	if err != nil {
		return
	}

	db = &Base{
		AuthQueries:    AuthQueries{DB: sqlite},
		InstallQueries: InstallQueries{DB: sqlite},
		SettingQueries: SettingQueries{DB: sqlite},
		PageQueries:    PageQueries{DB: sqlite},
		ProductQueries: ProductQueries{DB: sqlite},
		CartQueries:    CartQueries{DB: sqlite},
	}
	return
}

// DB is ...
func DB() *Base {
	if db == nil {
		db = &Base{}
	}
	return db
}
