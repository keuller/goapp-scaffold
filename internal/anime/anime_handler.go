package anime

import (
	"log/slog"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/keuller/goapp/internal/views"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GetAnimes(res http.ResponseWriter, req *http.Request) {
	page, _ := views.GetPage("./pages/animes/index.html")
	service := NewAnimeService()

	vars := make(jet.VarMap)
	vars.Set("animes", service.GetAll())

	page.Execute(res, vars, nil)
}

func SaveAnime(res http.ResponseWriter, req *http.Request) {
	page, _ := views.GetPage("./pages/animes/anime_row.html")

	completed := false
	name := req.FormValue("name")
	if req.FormValue("complete") == "on" {
		completed = true
	}

	slog.Info("saving anime", "completed", completed)

	vars := make(jet.VarMap)
	newId, _ := gonanoid.Generate("ABCDEFGHIKMNYWX1234567890", 10)
	vars.Set("row", Anime{ID: newId, Name: name, Complete: completed})

	page.Execute(res, vars, nil)
}
