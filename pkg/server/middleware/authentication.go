package middleware

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/mrcrgl/pflog/log"
	"net/http"
	"os"
	"strings"
)

type AuthUser struct {
	// Id of the account
	Pnum          string `json:"pnum"`
	EmailVerified bool   `json:"emailVerified"`
	Name          string `json:"name"`
	GivenName     string `json:"givenName"`
	FamilyName    string `json:"family_name"`
	Email         string `json:"email"`
}

const JWTIdTokenCookieName = "SIPS-OIDC-AUTH"

// DetectOIDC checks if the request contains a JWT that can be checked
func DetectJWT(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 {
			return "", errors.New("detected bearer token, but in invalid format")
		}
		return authHeaderParts[1], nil
	}
	if cookie, err := r.Cookie(JWTIdTokenCookieName); err == nil {
		return cookie.Value, nil
	}
	return "", nil
}

func SIPSOIDCTokenRetriever(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("APP_ENV") == "dev" {
			dummyUser := AuthUser{Pnum: "p1234567"}
			ctx := context.WithValue(r.Context(), "authUser", dummyUser)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		jwt, err := DetectJWT(r)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		jwtParts := strings.Split(jwt, ".")
		if len(jwtParts) < 3 {
			next.ServeHTTP(w, r)
			return
		}

		rawClaims, err := base64.RawStdEncoding.DecodeString(jwtParts[1])
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		var authUser AuthUser
		err = json.Unmarshal(rawClaims, &authUser)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), "authUser", authUser)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

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
