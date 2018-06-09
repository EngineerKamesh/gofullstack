package common

import (
	"github.com/EngineerKamesh/gofullstack/volume3/section4/gopherface/common/datastore"
	"go.isomorphicgo.org/go/isokit"
)

type Env struct {
	DB          datastore.Datastore
	TemplateSet *isokit.TemplateSet
}
