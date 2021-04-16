package server

import (
	"github.com/gin-gonic/gin"
	serverMiddleware "github.com/leachim2k/go-shorten/pkg/server/middleware"
	"github.com/leachim2k/go-shorten/pkg/shorten"
)

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
		r.Use(serverMiddleware.JWTAuthenticator)

		r.PUT("/", apiRouting.MissingCodeHandler)
		r.DELETE("/", apiRouting.MissingCodeHandler)

		r.GET("/", apiRouting.GetAllHandler)

		r.POST("/", apiRouting.AddHandler)
		r.GET("/:code", apiRouting.GetHandler)
		r.GET("/handle/:code", apiRouting.HandleCodeHandler)
		r.PUT("/:code", apiRouting.UpdateHandler)
		r.DELETE("/:code", apiRouting.DeleteHandler)
	}

}
