package main

import (
	"log/slog"

	"github.com/keuller/goapp/app"
)

func main() {
	webapp := app.NewApp()
	go webapp.Start()
	slog.Info("Server is up and running at http://localhost:7001")
	webapp.WaitToStop()
}
