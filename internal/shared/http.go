package shared

import "net/http"

func IsHtmxRequest(req *http.Request) bool {
	return req.Header.Get("HX-Request") != ""
}

func SetHtmxTarget(res http.ResponseWriter, value string) {
	res.Header().Add("HX-Target", value)
}
