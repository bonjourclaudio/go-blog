package auth

import (
	_ "context"
	_ "encoding/json"
	"github.com/claudioontheweb/go-blog/models"
	_ "github.com/dgrijalva/jwt-go"
	"net/http"
	_ "strings"
)

type Exception models.Exception

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
/*
		// Get Token from Header
		var tk = r.Header.Get("x-access-token")
		tk = strings.TrimSpace(tk)

		// Check if no Token and return 401
		if tk == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(Exception{Message: "Missing auth token"})
			return
		}

		// Declare claims struct
		claims := &models.CustomClaims{}

		// Parse and validate Token
		_, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(Exception{Message: err.Error()})
			return
		}

		// Assign Claims to context and continue with HTTP Request
		ctx := context.WithValue(r.Context(), "user", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
*/
		next.ServeHTTP(w, r)
	})
}