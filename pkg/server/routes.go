package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	serverMiddleware "github.com/leachim2k/go-shorten/pkg/server/middleware"
	"github.com/leachim2k/go-shorten/pkg/shorten"
	"net/http"
	"strings"
)

func NoRoute(c *gin.Context) {
	path := strings.TrimLeft(c.Request.URL.Path, "/")
	if path == "" {
		path = "index.html"
	}
	_, err := AssetInfo(path)
	fmt.Printf("err: %#v", err)
	if err == nil {
		c.FileFromFS(c.Request.URL.Path, AssetFile())
		return
	}

	apiRouting := shorten.NewApiHandler()
	link, err := apiRouting.HandleCode(path)
	fmt.Printf("err: %#v", err)
	if link != nil && *link != "" {
		c.Redirect(http.StatusFound, *link)
		return
	}
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
	}

}
