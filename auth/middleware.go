package auth

import (
	"context"
	"encoding/json"
	"github.com/claudioontheweb/go-blog/models"
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
)

//Exception struct
type Exception models.Exception

// JwtVerify Middleware function
func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token = r.Header.Get("x-access-token") // Grab the token from the header

		token = strings.TrimSpace(token)

		if token == "" {
			// Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: "Missing auth token"})
			return
		}
		claims := &models.Token{}

		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "user", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}