package handlers

import (
	"net/http"
)

func MyProfileHandler(w http.ResponseWriter, r *http.Request) {

	m := make(map[string]string)
	m["PageTitle"] = "My Profile"
	RenderGatedTemplate(w, "./templates/profile.html", m)

}
