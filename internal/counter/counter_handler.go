package counter

import (
	"net/http"
	"sync/atomic"

	"github.com/CloudyKit/jet/v6"
	"github.com/keuller/goapp/internal/views"
)

var value atomic.Int32

func Increment(res http.ResponseWriter, req *http.Request) {
	page, _ := views.GetPage("./partials/counter.html")

	value.Add(1)

	vars := make(jet.VarMap)
	vars.Set("value", value.Load())
	page.Execute(res, vars, nil)
}

func Decrement(res http.ResponseWriter, req *http.Request) {
	page, _ := views.GetPage("./partials/counter.html")

	value.Add(-1)

	vars := make(jet.VarMap)
	vars.Set("value", value.Load())
	page.Execute(res, vars, nil)
}
