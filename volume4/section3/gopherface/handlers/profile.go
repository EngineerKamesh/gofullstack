package handlers

import (
	"log"
	"net/http"

	"github.com/EngineerKamesh/gofullstack/volume4/section3/gopherface/common"
	"github.com/gorilla/mux"
	"github.com/isomorphicgo/isokit"
)

func ProfileHandler(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		username := vars["username"]

		u, err := env.DB.GetGopherProfile(username)
		if err != nil {
			log.Print(err)
		}

		u.PageTitle = u.Username
		env.TemplateSet.Render("gopherprofile_page", &isokit.RenderParams{Writer: w, Data: u})

	})
}
