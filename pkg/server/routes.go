package server

import (
	"github.com/gin-gonic/gin"
	serverMiddleware "github.com/leachim2k/go-shorten/pkg/server/middleware"
	"github.com/leachim2k/go-shorten/pkg/shorten"
	"net/http"
	"strings"
)

func NoRoute(ctx *gin.Context) {
	path := strings.TrimLeft(ctx.Request.URL.Path, "/")
	if path == "" {
		path = "index.html"
	}
	_, err := AssetInfo(path)
	if err == nil {
		ctx.FileFromFS(ctx.Request.URL.Path, AssetFile())
		return
	}

	apiRouting := shorten.NewApiHandler()
	link, err := apiRouting.HandleCode(path, ctx.ClientIP(), ctx.Request.UserAgent(), ctx.Request.Referer())
	if link != nil && *link != "" {
		ctx.Redirect(http.StatusFound, *link)
		return
	}

	// this is needed because of SPA's self-routing
	ctx.FileFromFS("/", AssetFile())
}

func NewGroup(r *gin.Engine) {
	v1 := r.Group("/api")
	{
		shortenRouter(v1)
	}
}

func shortenRouter(apiRoute *gin.RouterGroup) {
	apiRouting := shorten.NewApiHandler()

	r := apiRoute.Group("/shorten")
	{
		r.GET("/handle/:code", apiRouting.HandleCodeHandler)

		r.Use(serverMiddleware.JWTAuthenticator)

		r.PUT("/", apiRouting.MissingCodeHandler)
		r.DELETE("/", apiRouting.MissingCodeHandler)

		r.GET("/", apiRouting.GetAllHandler)

		r.POST("/", apiRouting.AddHandler)
		r.GET("/:code", apiRouting.GetHandler)
		r.PUT("/:code", apiRouting.UpdateHandler)
		r.DELETE("/:code", apiRouting.DeleteHandler)

		r.GET("/:code/stats", apiRouting.GetStatsHandler)

	}

}
