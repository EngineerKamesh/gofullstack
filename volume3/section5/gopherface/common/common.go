package common

import (
	"github.com/EngineerKamesh/gofullstack/volume3/section5/gopherface/common/datastore"
	"github.com/isomorphicgo/isokit"
)

type Env struct {
	DB          datastore.Datastore
	TemplateSet *isokit.TemplateSet
}
