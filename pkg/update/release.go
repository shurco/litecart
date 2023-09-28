package update

import (
	"errors"
	"strings"
)

type release struct {
	Id        int             `json:"id"`
	Name      string          `json:"name"`
	Tag       string          `json:"tag_name"`
	Published string          `json:"published_at"`
	Url       string          `json:"html_url"`
	Assets    []*ReleaseAsset `json:"assets"`
}

type ReleaseAsset struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Size        int    `json:"size"`
	DownloadUrl string `json:"browser_download_url"`
}

func (r *release) findAssetBySuffix(suffix string) (*ReleaseAsset, error) {
	if suffix != "" {
		for _, asset := range r.Assets {
			if strings.HasSuffix(asset.Name, suffix) {
				return asset, nil
			}
		}
	}

	return nil, errors.New("missing asset containing " + suffix)
}
