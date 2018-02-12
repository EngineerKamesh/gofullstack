package handlers

import (
	"net/http"

	"github.com/EngineerKamesh/gofullstack/volume4/section2/gopherface/common"
	"github.com/EngineerKamesh/gofullstack/volume4/section2/gopherface/forms"

	"github.com/isomorphicgo/isokit"
)

func FeedHandler(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		formParams := isokit.FormParams{ResponseWriter: w, Request: r}
		p := forms.NewSocialMediaPostForm(&formParams)
		p.PageTitle = "Feed"
		env.TemplateSet.Render("feed_page", &isokit.RenderParams{Writer: w, Data: p})
	})
}
