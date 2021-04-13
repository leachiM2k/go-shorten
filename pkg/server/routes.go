package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/docgen"
	serverMiddleware "github.com/leachim2k/go-shorten/pkg/server/middleware"
	"github.com/leachim2k/go-shorten/pkg/shorten"
	"net/http"
	"time"
)

func NewGroup(r *gin.Engine) {
	v1 := r.Group("/api")
	{
		shortenRouter(v1)
	}
}

func NewRouter(serveMux *http.ServeMux) {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api", func(r chi.Router) {
		r.Use(serverMiddleware.FullAuditLog)

		//r.Mount("/shorten", shortenRouter())
	})

	serveMux.Handle("/api/", r)
	serveMux.HandleFunc("/docs", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
			ProjectPath: "github.com/go-chi/chi",
			Intro:       "Welcome to the chi/_examples/rest generated docs.",
		}))
	})
}

func shortenRouter(apiRoute *gin.RouterGroup) {
	apiRouting := shorten.NewApiHandler()

	r := apiRoute.Group("/shorten")
	{
		r.GET("/", apiRouting.MissingCodeHandler)
		r.PUT("/", apiRouting.MissingCodeHandler)
		r.DELETE("/", apiRouting.MissingCodeHandler)

		r.POST("/", apiRouting.AddHandler)
		r.GET("/:code", apiRouting.GetHandler)
		r.GET("/handle/:code", apiRouting.HandleCodeHandler)
		r.PUT("/:code", apiRouting.UpdateHandler)
		r.DELETE("/:code", apiRouting.DeleteHandler)
	}

}
