package handlers

import (
	"net/http"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("profile"))

}
