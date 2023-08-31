package web

import (
	"embed"
)

//go:embed all:admin/dist
var embedAdmin embed.FS

//go:embed admin/dist/index.html
var embedAdminIndex embed.FS

//go:embed site/*.html site/layouts/*.html site/public/*
var embedSite embed.FS

// EmbedAdmin is ...
func EmbedAdmin() embed.FS {
	return embedAdmin
}

// EmbedAdminIndex is ...
func EmbedAdminIndex() embed.FS {
	return embedAdminIndex
}

// EmbedSite is ...
func EmbedSite() embed.FS {
	return embedSite
}
