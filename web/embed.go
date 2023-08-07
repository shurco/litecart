package web

import "embed"

//go:embed admin/*
//go:embed site/*
var embedWeb embed.FS

// Embed is ...
func Embed() embed.FS {
	return embedWeb
}
