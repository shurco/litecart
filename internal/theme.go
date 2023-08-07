package app

import (
	"github.com/shurco/litecart/pkg/fsutil"
	"github.com/shurco/litecart/web"
)

// NewTheme is ...
func NewTheme() error {
	if err := fsutil.EmbedExtract(web.Embed(), ""); err != nil {
		return err
	}
	return nil
}
