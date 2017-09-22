package main

import (
	"log"
	"os"

	"github.com/isomorphicgo/isokit"

	"github.com/EngineerKamesh/gofullstack/volume3/section4/gopherface/common"
	"github.com/EngineerKamesh/gofullstack/volume3/section4/gopherface/common/datastore"
	"github.com/EngineerKamesh/gofullstack/volume3/section4/gopherface/handlers"
	"github.com/EngineerKamesh/gofullstack/volume3/section4/gopherface/middleware"

	"net/http"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

const (
	WEBSERVERPORT = ":8443"
)

var WebAppRoot = os.Getenv("GOPHERFACE_APP_ROOT")

func main() {

	//asyncq.StartTaskDispatcher(9)

	db, err := datastore.NewDatastore(datastore.MYSQL, "gopherface:gopherface@/gopherfacedb")
	//db, err := datastore.NewDatastore(datastore.MONGODB, "localhost:27017")
	//db, err := datastore.NewDatastore(datastore.REDIS, "localhost:6379")

	if err != nil {
		log.Print(err)
	}

	defer db.Close()

	env := common.Env{}
	isokit.TemplateFilesPath = WebAppRoot + "/templates"
	isokit.TemplateFileExtension = ".html"
	ts := isokit.NewTemplateSet()
	ts.GatherTemplates()
	env.TemplateSet = ts
	env.DB = db

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	r.Handle("/signup", handlers.SignUpHandler(&env)).Methods("GET", "POST")

	r.Handle("/login", handlers.LoginHandler(&env)).Methods("GET", "POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET", "POST")

	r.Handle("/feed", middleware.GatedContentHandler(handlers.FeedHandler(&env))).Methods("GET")
	r.Handle("/friends", middleware.GatedContentHandler(handlers.FriendsHandler(&env))).Methods("GET")
	r.Handle("/profile", middleware.GatedContentHandler(handlers.MyProfileHandler(&env))).Methods("GET")
	r.Handle("/find", middleware.GatedContentHandler(handlers.FindHandler)).Methods("GET,POST")
	r.Handle("/profile/{username}", middleware.GatedContentHandler(handlers.ProfileHandler)).Methods("GET")
	r.Handle("/postpreview", middleware.GatedContentHandler(handlers.PostPreviewHandler)).Methods("GET", "POST")
	r.Handle("/upload-image", middleware.GatedContentHandler(handlers.UploadImageHandler)).Methods("GET", "POST")
	r.Handle("/upload-video", middleware.GatedContentHandler(handlers.UploadVideoHandler)).Methods("GET", "POST")

	r.Handle("/js/client.js", isokit.GopherjsScriptHandler(WebAppRoot))
	r.Handle("/js/client.js.map", isokit.GopherjsScriptMapHandler(WebAppRoot))
	r.Handle("/template-bundle", handlers.TemplateBundleHandler(&env))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(WebAppRoot+"/static"))))

	loggedRouter := ghandlers.LoggingHandler(os.Stdout, r)
	stdChain := alice.New(middleware.PanicRecoveryHandler)
	http.Handle("/", stdChain.Then(loggedRouter))

	err = http.ListenAndServeTLS(WEBSERVERPORT, WebAppRoot+"/certs/gopherfacecert.pem", WebAppRoot+"/certs/gopherfacekey.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
