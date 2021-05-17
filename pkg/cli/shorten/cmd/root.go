package cmd

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	_ "github.com/leachim2k/go-shorten/docs"
	"github.com/leachim2k/go-shorten/pkg/auth"
	shortenServer "github.com/leachim2k/go-shorten/pkg/server"
	"github.com/leachim2k/go-shorten/pkg/ui"
	"github.com/mrcrgl/pflog/log"
	"net/http"
	"os"
	"strconv"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/leachim2k/go-shorten/pkg/cli/shorten/options"
	"github.com/leachim2k/go-shorten/pkg/version"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "go-shorten",
		Long: ``,
		Run: func(cmd *cobra.Command, args []string) {
			log.Infof("This is %s [%s]", cmd.Use, version.GetInfo())

			err := options.Current.Validate()
			if err != nil {
				log.Fatalf("config validation failed: %v", err)
			}

			r := gin.New()

			if os.Getenv("APP_ENV") != "dev" {
				gin.SetMode(gin.ReleaseMode)
			}
			r.ForwardedByClientIP = true

			// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
			// By default gin.DefaultWriter = os.Stdout
			r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

				// your custom format
				return fmt.Sprintf("I%s [REQUEST] %s \"%s %s %s %d %s \"%s\" %s\"\n",
					param.TimeStamp.Format(time.RFC3339),
					param.ClientIP,
					param.Method,
					param.Path,
					param.Request.Proto,
					param.StatusCode,
					param.Latency,
					param.Request.UserAgent(),
					param.ErrorMessage,
				)
			}))
			r.Use(gin.ErrorLogger())
			r.Use(gin.Recovery())

			shortenServer.NewGroup(r)
			ui.NewGroup(r)
			auth.NewGroup(r)

			r.GET("/swagger", func(c *gin.Context) {
				c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
			})
			r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

			r.GET("/header/echo", func(context *gin.Context) {
				context.JSON(200, context.Request.Header)
			})

			r.NoRoute(ui.NoRoute, shortenServer.NoRoute)

			err = endless.ListenAndServe(":"+strconv.Itoa(options.Current.Server.Port), r)
			if err != nil {
				log.Fatalf("Failed to boot application: %s", err)
				return
			}
		},
	}

	cmd.Flags().IntVar(&options.Current.Server.Port, "rest-listen-port", options.Current.Server.Port, "tcp port to listen for HTTP requests")
	cmd.Flags().StringVar(&options.Current.Storage.DBUrl, "db-connection", options.Current.Storage.DBUrl, "database connection string")
	cmd.Flags().StringVar(&options.Current.Storage.Engine, "storage", options.Current.Storage.Engine, "storage backend for shorts (memory, postgresql, mysql, file)")

	return cmd
}
