package app

import (
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/go-chi/chi/v5"
	"github.com/keuller/goapp/assets"
	"github.com/keuller/goapp/internal/counter"
	"github.com/keuller/goapp/internal/views"
)

func registerAppRoutes(router *chi.Mux) {
	// public assets handler
	if handler, err := assets.GetAssetsHandler(); err == nil {
		router.Handle("/assets/*", http.StripPrefix("/assets", handler))
	}

	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		page, _ := views.GetPage("./pages/index.html")

		vars := make(jet.VarMap)
		vars.Set("value", counter.GetCounterValue())

		page.Execute(res, vars, nil)
	})

	router.Get("/about", func(res http.ResponseWriter, req *http.Request) {
		page, _ := views.GetPage("./pages/about.html")

		page.Execute(res, nil, nil)
	})

	router.Post("/inc", counter.Increment)
	router.Post("/dec", counter.Decrement)
}
