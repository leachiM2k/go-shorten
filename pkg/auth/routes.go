package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/leachim2k/go-shorten/pkg/auth/middleware"
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
	callbackUrl.Scheme = "https"

	proto := c.Request.Header.Get("X-Forwarded-Proto")
	if c.Request.TLS == nil || (proto != "" && strings.ToLower(proto) != "https") {
		callbackUrl.Scheme = "http"
	}

	// In this case we use a map to store our secrets, but you can use dotenv or your framework configuration
	// for example, in revel you could use revel.Config.StringDefault(provider + "_clientID", "") etc.
	providerSecrets := map[string]map[string]string{
		"facebook": {
			"clientID":     "***REMOVED***",
			"clientSecret": "***REMOVED***",
		},
		"google": {
			"clientID":     "***REMOVED***",
			"clientSecret": "***REMOVED***",
		},
	}

	for s, m := range providerSecrets {
		callbackUrl.Path = "/auth/" + s + "/callback"
		m["redirectURL"] = callbackUrl.String()
	}

	providerScopes := map[string][]string{
		"facebook": []string{},
		"google":   []string{},
	}

	providerData := providerSecrets[provider]
	actualScopes := providerScopes[provider]
	authURL, err := gocial.New().
		Driver(provider).
		Scopes(actualScopes).
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
	callbackUrl.Scheme = "https"

	proto := c.Request.Header.Get("X-Forwarded-Proto")
	if c.Request.TLS == nil || (proto != "" && strings.ToLower(proto) != "https") {
		callbackUrl.Scheme = "http"
	}
	callbackUrl.Path = "/ui/login/"

	q := url.Values{}
	q.Set("jwt", t2.AccessToken)
	callbackUrl.RawQuery = q.Encode()

	c.Redirect(http.StatusFound, callbackUrl.String())
	//c.JSON(http.StatusOK, t2)
}
