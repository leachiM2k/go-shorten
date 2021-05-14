package middleware

import (
	"errors"
	"fmt"
	"github.com/danilopolani/gocialite/structs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"strings"
	"time"
)

type AuthCustomClaims struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Provider string `json:"p"`
	jwt.StandardClaims
}

const (
	issuer = "GoShorten"
)

func GetSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func JWTAuthenticator(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	claims, err := GetClaimFromAuthHeader(authHeader)
	if err != nil {
		c.AbortWithError(http.StatusForbidden, err)
		return
	}
	c.Set("JWT_CLAIMS", claims)
}

func GetClaimFromAuthHeader(authHeader string) (*AuthCustomClaims, error) {
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, fmt.Errorf("invalid Token Type (Bearer only)")
	}
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 {
		return nil, fmt.Errorf("detected bearer token, but in invalid format")
	}

	// Validate the JWT is valid
	claims, err := ValidateJWT(authHeaderParts[1])
	if err != nil {
		return nil, err
	}

	return claims, err
}

func ValidateJWT(tokenString string) (*AuthCustomClaims, error) {
	claimsStruct := AuthCustomClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}
			return []byte(GetSecretKey()), nil
			/*
				pem, err := getGooglePublicKey(fmt.Sprintf("%s", token.Header["kid"]))
				if err != nil {
					return nil, err
				}
				key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
				if err != nil {
					return nil, err
				}
				return key, nil
			*/
		},
	)

	if err != nil {
		return &AuthCustomClaims{}, err
	}

	claims, ok := token.Claims.(*AuthCustomClaims)
	if !ok {
		return &AuthCustomClaims{}, errors.New("Invalid Google JWT")
	}

	if claims.Issuer != issuer {
		return &AuthCustomClaims{}, errors.New("iss is invalid")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return &AuthCustomClaims{}, errors.New("JWT is expired")
	}

	return claims, nil
}

func BuildJWTToken(user *structs.User, token *oauth2.Token, provider string) (*oauth2.Token, error) {
	providerPrefix := map[string]string{
		"facebook": "f",
		"google":   "g",
	}

	claims := &AuthCustomClaims{
		Name:     user.FullName,
		Email:    user.Email,
		Provider: provider,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: token.Expiry.Unix(),
			Id:        user.ID,
			Subject:   providerPrefix[provider] + "_" + user.ID,
			IssuedAt:  time.Now().Unix(),
			Issuer:    issuer,
		},
	}
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := tkn.SignedString([]byte(GetSecretKey()))
	if err != nil {
		return nil, err
	}

	t2 := oauth2.Token{
		AccessToken: t,
		TokenType:   "Bearer",
		Expiry:      token.Expiry,
	}
	return &t2, nil
}
