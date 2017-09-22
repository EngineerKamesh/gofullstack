package main

import (
	"strings"

	"github.com/EngineerKamesh/gofullstack/volume3/section5/gopherface/client/handlers"

	"github.com/EngineerKamesh/gofullstack/volume3/section5/gopherface/client/common"

	"github.com/isomorphicgo/isokit"
	"honnef.co/go/js/dom"
)

var D = dom.GetWindow().Document().(dom.HTMLDocument)

func initializeEventHandlers(env *common.Env) {
	println("location: ", env.Window.Location().Href)
	l := strings.Split(env.Window.Location().Href, "/")
	pageName := l[len(l)-1]
	switch pageName {
	case "feed":
		handlers.InitializeFeedEventHandlers(env)
	case "friends":
		handlers.InitializeFriendsEventHandlers(env)
	case "myprofile":
		handlers.InitializeMyProfileEventHandlers(env)
	}
}

func run() {
	templateSetChannel := make(chan *isokit.TemplateSet)
	go isokit.FetchTemplateBundle(templateSetChannel)
	ts := <-templateSetChannel

	env := common.Env{}
	env.TemplateSet = ts
	println("env.TemplateSet: ", env.TemplateSet)

	env.Window = dom.GetWindow()
	env.Document = dom.GetWindow().Document()
	env.PrimaryContent = env.Document.GetElementByID("primaryContent")

	r := isokit.NewRouter()
	r.Handle("/feed", handlers.FeedHandler(&env))
	r.Handle("/friends", handlers.FriendsHandler(&env))
	r.Handle("/myprofile", handlers.MyProfileHandler(&env))
	r.Listen()

	initializeEventHandlers(&env)
}

func main() {
	switch readyState := D.ReadyState(); readyState {
	case "loading":
		D.AddEventListener("DOMContentLoaded", false, func(dom.Event) {
			go run()
		})
	case "interactive", "complete":
		run()
	default:
		println("Unexpected document.ReadyState value!")
	}
}
