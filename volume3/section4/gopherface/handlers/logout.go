package handlers

import (
	"net/http"

	"github.com/EngineerKamesh/gofullstack/volume3/section4/gopherface/common/authenticate"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	authenticate.ExpireUserSession(w, r)
	authenticate.ExpireSecureCookie(w, r)
}
