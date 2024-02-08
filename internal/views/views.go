package views

import (
	"github.com/CloudyKit/jet/v6"
)

var (
	views = jet.NewSet(
		jet.NewOSFileSystemLoader("./internal/views"),
		jet.InDevelopmentMode(), // remove in production
	)
)

func GetPage(pageName string) (*jet.Template, error) {
	return views.GetTemplate(pageName)
}
