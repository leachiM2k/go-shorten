package main

import (
	"github.com/leachim2k/go-shorten/pkg/cli/shorten/cmd"
	"github.com/mrcrgl/pflog/log"
	"github.com/spf13/pflag"
)

// @title Go Shorten API
// @version 1.0
// @description URL Shortener
// @termsOfService http://swagger.io/terms/

// @contact.name leachiM2k
// @contact.url https://github.com/leachim2k/go-shorten
// @contact.email leachiM2k@leachiM2k.de

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
// @query.collection.format multi

// @x-extension-openapi {"example": "value on a json format"}
func main() {
	rootCmd := cmd.NewRootCommand()
	pflag.CommandLine.AddFlagSet(rootCmd.Flags())

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Execution failed: %s", err)
	}
}
