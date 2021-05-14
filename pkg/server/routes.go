package server

import (
	"github.com/gin-gonic/gin"
	authMiddleware "github.com/leachim2k/go-shorten/pkg/auth/middleware"
	"github.com/leachim2k/go-shorten/pkg/shorten"
	"github.com/leachim2k/go-shorten/pkg/urlutil"
	"net/http"
	"regexp"
	"strings"
)

func NoRoute(ctx *gin.Context) {
	path := strings.TrimLeft(ctx.Request.URL.Path, "/")
	matched, _ := regexp.MatchString("^[abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-]*$", path)
	if matched {
		apiRouting := shorten.NewApiHandler()
		link, _ := apiRouting.HandleCode(path, ctx.ClientIP(), ctx.Request.UserAgent(), ctx.Request.Referer())
		if link != nil && *link != "" {
			ctx.Redirect(http.StatusFound, *link)
			return
		}
	}
}

func NewGroup(r *gin.Engine) {
	v1 := r.Group("/api")
	{
		shortenRouter(v1)
		utilRouter(v1)
	}
}

func utilRouter(apiRoute *gin.RouterGroup) {
	utilRouting := urlutil.NewApiHandler()

	r := apiRoute.Group("/url")
	{
		r.GET("/meta/", utilRouting.GetUrlMetaHandler)
		r.GET("/qrcode/", utilRouting.GetQrCodeHandler)
	}
}

func shortenRouter(apiRoute *gin.RouterGroup) {
	apiRouting := shorten.NewApiHandler()

	r := apiRoute.Group("/shorten")
	{
		r.GET("/handle/:code", apiRouting.HandleCodeHandler)

		r.Use(authMiddleware.JWTAuthenticator)

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
