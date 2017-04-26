package main

import (
	"fmt"
	"net/http"

	"github.com/EngineerKamesh/gofullstack/volume1/section5/validationkit"
)

func helloGopherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Gopher!")
}

func checkUsernameSyntaxHandler(w http.ResponseWriter, r *http.Request) {

	var usernameSyntaxResult bool
	username := r.URL.Query().Get("username")

	if username == "" {
		http.Error(w, "Username not provided!", http.StatusInternalServerError)
	} else {
		usernameSyntaxResult = validationkit.CheckUsernameSyntax(username)
		fmt.Fprintf(w, "Sytnax Check Result for %v is %v", username, usernameSyntaxResult)
	}

}

func main() {
	http.HandleFunc("/hello-gopher", helloGopherHandler)
	http.HandleFunc("/username-syntax-checker", checkUsernameSyntaxHandler)
	http.ListenAndServe(":8080", nil)
}
