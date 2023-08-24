package web

import "embed"

//go:embed admin/dist/*
//go:embed site/*.html site/components/* site/layouts/* site/public/*
var embedWeb embed.FS

// Embed is ...
func Embed() embed.FS {
	return embedWeb
}
