package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}

func BasicsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "./templates/basics.html", nil)
	})
}

func LowercaseTextTransformHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var s string

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}

		reqBodyString := string(reqBody)

		err = json.Unmarshal([]byte(reqBodyString), &s)
		if err != nil {
			log.Print("Encountered error when attempting to unmarshal JSON: ", err)
		}

		textBytes, err := json.Marshal(strings.ToLower(s))
		if err != nil {
			log.Print("Encountered error when attempting ot marshal JSON: ", err)
		}
		fmt.Println("textBytes string: ", string(textBytes))
		w.Write(textBytes)

	})

}

func main() {

	r := mux.NewRouter()

	r.Handle("/", BasicsHandler()).Methods("GET")
	r.Handle("/lowercase-text", LowercaseTextTransformHandler())

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", r)
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)

}
