package assets

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed all:public
var public embed.FS

func GetAssetsHandler() (http.Handler, error) {
	staticAssets, err := fs.Sub(public, "public")
	if err != nil {
		return nil, err
	}
	return http.FileServer(http.FS(staticAssets)), nil
}
