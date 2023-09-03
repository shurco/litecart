package queries

import (
	"embed"

	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"

	"github.com/shurco/litecart/pkg/fsutil"
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

// InitDB is ...
func InitDB(dbPath string, migrations embed.FS) (db *sql.DB, err error) {
	if !fsutil.IsFile(dbPath) {
		// create db
		if _, err = fsutil.OpenFile(dbPath, fsutil.FsCWFlags, 0666); err != nil {
			return
		}

		// first migrate db
		if err = Migrate(dbPath, migrations); err != nil {
			return
		}
	}

	// connect to database
	dsn := fmt.Sprintf("%s?_pragma=busy_timeout(10000)&_pragma=journal_mode(WAL)&_pragma=journal_size_limit(200000000)&_pragma=synchronous(NORMAL)&_pragma=foreign_keys(ON)", dbPath)
	db, err = sql.Open("sqlite", dsn)
	db.Query("PRAGMA auto_vacuum")

	return
}

// Migrate is ...
func Migrate(dbPath string, migrations embed.FS) (err error) {
	goose.SetBaseFS(migrations)
	var db *sql.DB
	db, err = goose.OpenDBWithDriver("sqlite", dbPath)
	if err != nil {
		return
	}
	defer db.Close()

	goose.SetTableName("migrate_db_version")

	err = goose.Up(db, ".")
	return
}

// InitQueries is ...
func InitQueries(embed embed.FS) (err error) {
	// init database
	var sqlite *sql.DB
	sqlite, err = InitDB("./lc_base/data.db", embed)
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
