package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mrcrgl/pflog/log"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// GoogleClaims -
type GoogleClaims struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FirstName     string `json:"given_name"`
	LastName      string `json:"family_name"`
	jwt.StandardClaims
}

type AuthUser struct {
	// Id of the account
	Pnum          string `json:"pnum"`
	EmailVerified bool   `json:"emailVerified"`
	Name          string `json:"name"`
	GivenName     string `json:"givenName"`
	FamilyName    string `json:"family_name"`
	Email         string `json:"email"`
}

var googlePublicKey, _ = loadGooglePublicKey()

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		authUser, ok := ctx.Value("authUser").(AuthUser)
		if !ok {
			http.Error(w, http.StatusText(401), 401)
			return
		}
		log.Infof("Found user: %+v", authUser.Pnum)

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}

func JWTAuthenticator(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 {
		c.AbortWithError(http.StatusForbidden, errors.New("detected bearer token, but in invalid format"))
		return
	}

	// Validate the JWT is valid
	claims, err := ValidateGoogleJWT(authHeaderParts[1])
	if err != nil {
		c.AbortWithError(http.StatusForbidden, err)
		return
	}

	c.Set("JWT_CLAIMS", claims)

	return
}

func ValidateGoogleJWT(tokenString string) (GoogleClaims, error) {
	claimsStruct := GoogleClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) {
			pem, err := getGooglePublicKey(fmt.Sprintf("%s", token.Header["kid"]))
			if err != nil {
				return nil, err
			}
			key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
			if err != nil {
				fmt.Printf("-- 2 -- \n")
				return nil, err
			}
			return key, nil
		},
	)

	if err != nil {
		return GoogleClaims{}, err
	}

	claims, ok := token.Claims.(*GoogleClaims)
	if !ok {
		return GoogleClaims{}, errors.New("Invalid Google JWT")
	}

	if claims.Issuer != "accounts.google.com" && claims.Issuer != "https://accounts.google.com" {
		return GoogleClaims{}, errors.New("iss is invalid")
	}

	if claims.Audience != "***REMOVED***" {
		return GoogleClaims{}, errors.New("aud is invalid")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return GoogleClaims{}, errors.New("JWT is expired")
	}

	return *claims, nil
}

func loadGooglePublicKey() ([]byte, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/certs")
	if err != nil {
		return nil, err
	}
	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

func getGooglePublicKey(keyID string) (string, error) {
	var dat []byte = nil
	var err error = nil

	if googlePublicKey == nil {
		dat, err = loadGooglePublicKey()
		if err != nil {
			return "", err
		}
	} else {
		dat = googlePublicKey
	}

	myResp := map[string]string{}
	err = json.Unmarshal(dat, &myResp)
	if err != nil {
		return "", err
	}
	key, ok := myResp[keyID]
	if !ok {
		return "", errors.New("key not found")
	}
	return key, nil
}
