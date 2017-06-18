package main

import (
	"log"
	"os"

	"github.com/EngineerKamesh/gofullstack/volume2/section6/gopherfaceauth/common"
	"github.com/EngineerKamesh/gofullstack/volume2/section6/gopherfaceauth/common/datastore"
	"github.com/EngineerKamesh/gofullstack/volume2/section6/gopherfaceauth/handlers"
	"github.com/EngineerKamesh/gofullstack/volume2/section6/gopherfaceauth/middleware"

	"net/http"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

const (
	WEBSERVERPORT = ":8443"
)

func main() {

	db, err := datastore.NewDatastore(datastore.MYSQL, "gopherface:gopherface@/gopherfacedb")
	//db, err := datastore.NewDatastore(datastore.MONGODB, "localhost:27017")
	//db, err := datastore.NewDatastore(datastore.REDIS, "localhost:6379")

	if err != nil {
		log.Print(err)
	}

	defer db.Close()

	env := common.Env{DB: db}

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	r.Handle("/signup", handlers.SignUpHandler(&env)).Methods("GET", "POST")

	r.Handle("/login", handlers.LoginHandler(&env)).Methods("GET", "POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET", "POST")

	r.Handle("/feed", middleware.GatedContentHandler(handlers.FeedHandler)).Methods("GET")
	r.Handle("/friends", middleware.GatedContentHandler(handlers.FriendsHandler)).Methods("GET")
	r.Handle("/find", middleware.GatedContentHandler(handlers.FindHandler)).Methods("GET,POST")
	r.Handle("/profile", middleware.GatedContentHandler(handlers.MyProfileHandler)).Methods("GET")
	r.Handle("/profile/{username}", middleware.GatedContentHandler(handlers.ProfileHandler)).Methods("GET")
	r.Handle("/postpreview", middleware.GatedContentHandler(handlers.PostPreviewHandler)).Methods("GET", "POST")
	r.Handle("/upload-image", middleware.GatedContentHandler(handlers.UploadImageHandler)).Methods("GET", "POST")
	r.Handle("/upload-video", middleware.GatedContentHandler(handlers.UploadVideoHandler)).Methods("GET", "POST")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	loggedRouter := ghandlers.LoggingHandler(os.Stdout, r)
	stdChain := alice.New(middleware.PanicRecoveryHandler)
	http.Handle("/", stdChain.Then(loggedRouter))

	err = http.ListenAndServeTLS(WEBSERVERPORT, "certs/gopherfacecert.pem", "certs/gopherfacekey.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
