package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/leachim2k/go-shorten/pkg/auth/middleware"
	"github.com/leachim2k/go-shorten/pkg/cli/shorten/options"
	"gopkg.in/danilopolani/gocialite.v1"
	"net/http"
	"net/url"
	"strings"
)

var gocial = gocialite.NewDispatcher()

func NewGroup(r *gin.Engine) {
	v1 := r.Group("/auth")
	{
		utilRouter(v1)
	}
}

func utilRouter(apiRoute *gin.RouterGroup) {
	apiRoute.GET("/", infoHandler)
	apiRoute.GET("/:provider/", redirectHandler)
	apiRoute.GET("/:provider/callback", callbackHandler)
}

func infoHandler(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	claims, err := middleware.GetClaimFromAuthHeader(authHeader)
	if err != nil {
		c.AbortWithError(http.StatusForbidden, err)
		return
	}
	c.JSON(http.StatusOK, claims)
}

// Redirect to correct oAuth URL
func redirectHandler(c *gin.Context) {
	// Retrieve provider from route
	provider := c.Param("provider")

	callbackUrl := *c.Request.URL
	callbackUrl.Host = c.Request.Host
	callbackUrl.Scheme = "http"

	proto := c.Request.Header.Get("X-Forwarded-Proto")
	cfVisitor := c.Request.Header.Get("Cf-Visitor")
	if c.Request.TLS != nil || (proto != "" && strings.ToLower(proto) == "https") || (cfVisitor != "" && strings.Contains(cfVisitor, "https")) {
		callbackUrl.Scheme = "https"
	}

	providerSecrets := map[string]map[string]string{}
	for _, service := range options.Current.AuthServices {
		callbackUrl.Path = "/auth/" + service.Name + "/callback"
		providerSecrets[service.Name] = map[string]string{
			"clientID":     service.ClientId,
			"clientSecret": service.ClientSecret,
			"redirectURL":  callbackUrl.String(),
		}
	}

	providerData := providerSecrets[provider]
	authURL, err := gocial.New().
		Driver(provider).
		Scopes([]string{}).
		Redirect(
			providerData["clientID"],
			providerData["clientSecret"],
			providerData["redirectURL"],
		)

	// Check for errors (usually driver not valid)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// Redirect with authURL
	c.Redirect(http.StatusFound, authURL)
}

// Handle callback of provider
func callbackHandler(c *gin.Context) {
	// Retrieve query params for state and code
	state := c.Query("state")
	code := c.Query("code")
	provider := c.Param("provider")

	// Handle callback and check for errors
	user, token, err := gocial.Handle(state, code)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	t2, err := middleware.BuildJWTToken(user, token, provider)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	callbackUrl := *c.Request.URL
	callbackUrl.Host = c.Request.Host
	callbackUrl.Scheme = "http"

	proto := c.Request.Header.Get("X-Forwarded-Proto")
	cfVisitor := c.Request.Header.Get("Cf-Visitor")
	if c.Request.TLS != nil || (proto != "" && strings.ToLower(proto) == "https") || (cfVisitor != "" && strings.Contains(cfVisitor, "https")) {
		callbackUrl.Scheme = "https"
	}
	callbackUrl.Path = "/ui/login/"

	q := url.Values{}
	q.Set("jwt", t2.AccessToken)
	callbackUrl.RawQuery = q.Encode()

	c.Redirect(http.StatusFound, callbackUrl.String())
}
