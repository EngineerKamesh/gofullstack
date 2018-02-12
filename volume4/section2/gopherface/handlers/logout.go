package handlers

import (
	"net/http"

	"github.com/EngineerKamesh/gofullstack/volume4/section2/gopherface/common/authenticate"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	authenticate.ExpireUserSession(w, r)
	authenticate.ExpireSecureCookie(w, r)
}
