package ui

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func NoRoute(ctx *gin.Context) {
	if ctx.Request.URL.Path == "/" {
		ctx.Redirect(http.StatusFound, "/ui/")
		return
	}
}

func NewGroup(r *gin.Engine) {
	r.Any("/ui/*any", func(ctx *gin.Context) {
		path := strings.TrimLeft(ctx.Request.URL.Path, "/ui")
		_, err := AssetInfo(path)
		if err == nil {
			ctx.FileFromFS(path, AssetFile())
			return
		}

		// this is needed because of SPA's self-routing
		ctx.FileFromFS("/", AssetFile())
	})
}
