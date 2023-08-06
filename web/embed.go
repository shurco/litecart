package web

import "embed"

//go:embed admin/*
var embedAdmin embed.FS

//go:embed site/*
var embedSite embed.FS

// Admin is ...
func Admin() embed.FS {
	return embedAdmin
}

// Admin is ...
func Site() embed.FS {
	return embedSite
}
