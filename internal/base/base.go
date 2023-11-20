package base

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
	"github.com/shurco/litecart/pkg/fsutil"
)

// New is ...
func New(dbPath string, migrations embed.FS) (db *sql.DB, err error) {
	if !fsutil.IsFile(dbPath) {
		// create db
		if _, err = fsutil.OpenFile(dbPath, fsutil.FsCWFlags, 0o666); err != nil {
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
