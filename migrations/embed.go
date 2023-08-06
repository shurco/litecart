package migrations

import "embed"

//go:embed *.sql
var migrations embed.FS

// Embed is ...
func Embed() embed.FS {
	return migrations
}
