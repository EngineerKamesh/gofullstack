package handlers

import (
	"context"

	"github.com/EngineerKamesh/gofullstack/volume3/section4/gopherface/client/common"

	"github.com/isomorphicgo/isokit"
)

func MyProfileHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {
		println("Client-side My Profile Handler")
		m := make(map[string]string)
		m["PageTitle"] = "My Profile"
		env.TemplateSet.Render("profile_content", &isokit.RenderParams{Data: m, Disposition: isokit.PlacementReplaceInnerContents, Element: env.PrimaryContent})
	})
}
