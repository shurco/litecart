package web

import (
	"embed"
)

const (
	// AdminBuildPath is the path to the admin panel build directory
	AdminBuildPath = "admin/build"
	// SiteBuildPath is the path to the site build directory
	SiteBuildPath = "site/build"
)

//go:embed all:admin/build
var embedAdmin embed.FS

//go:embed all:site/build
var embedSite embed.FS

// EmbedAdmin returns the embedded filesystem for the admin panel build directory
func EmbedAdmin() embed.FS {
	return embedAdmin
}

// EmbedSite returns the embedded filesystem for the site build directory
func EmbedSite() embed.FS {
	return embedSite
}
