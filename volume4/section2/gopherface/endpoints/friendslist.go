package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/EngineerKamesh/gofullstack/volume4/section2/gopherface/common/authenticate"

	"github.com/EngineerKamesh/gofullstack/volume4/section2/gopherface/common"
)

func FriendsListEndpoint(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		gfSession, err := authenticate.SessionStore.Get(r, "gopherface-session")
		if err != nil {
			log.Print(err)
			return
		}
		uuid := gfSession.Values["uuid"].(string)

		gophers, err := env.DB.FriendsList(uuid)

		if err != nil {
			log.Print(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gophers)

	})
}
