package main

import (
	"context"
	"encoding/json"
	"github.com/claudioontheweb/go-blog/auth"
	"github.com/claudioontheweb/go-blog/comment"
	"github.com/claudioontheweb/go-blog/config"
	"github.com/claudioontheweb/go-blog/customRouter"
	"github.com/claudioontheweb/go-blog/db"
	"github.com/claudioontheweb/go-blog/models"
	"github.com/claudioontheweb/go-blog/post"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strings"
)

func NewRouter() *mux.Router {

	// Init Router
	r := mux.NewRouter()

	// Use global Middlware
	r.Use(commonHandler)

	// Append Routes
	customRouter.Routes = append(customRouter.Routes, post.PostRoutes)
	customRouter.Routes = append(customRouter.Routes, post.TagRoutes)
	customRouter.Routes = append(customRouter.Routes, post.LikeRoutes)
	customRouter.Routes = append(customRouter.Routes, comment.Routes)
	customRouter.Routes = append(customRouter.Routes, auth.Routes)
	customRouter.Routes = append(customRouter.Routes, auth.UserRoutes)

	// Configure each route
	for _, route := range customRouter.Routes {

		// Create subroutes
		routePrefix := r.PathPrefix(route.Prefix).Subrouter()

		// Add Auth Middleware if needed
		if route.IsSecure {
			routePrefix.Use(jwtVerify)
		}

		// Loop through each subroute
		for _, r := range route.SubRoutes {
			var handler http.Handler
			handler = r.HandlerFunc
			routePrefix.Path(r.Pattern).Handler(handler).Methods(r.Method).Name(r.Name)

		}
	}

	// Root Route -> API DOC as HTML
	fileServer := http.FileServer(http.Dir("./docs/"))
	r.Handle("/", http.StripPrefix("/", fileServer))

	return r

}

func main() {

	config.GetConfig()

	port := viper.GetString("PORT")
	router := NewRouter()

	db.DB = db.ConnectDB()
	defer db.DB.Close()

	// Auto migrate Models
	db.DB.AutoMigrate(&auth.User{}, &post.Post{}, &post.Tag{}, &post.Like{}, &comment.Comment{})

	// Enable Log Mode
	db.DB.LogMode(true)

	// Start server
	log.Fatal(http.ListenAndServe(":" + port, router))

}


/*
 	#######################
	MIDDLEWARE
	#######################
 */

// Append Headers
func commonHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")

		next.ServeHTTP(w, r)
	})
}

// Check Auth Token

type Exception models.Exception

func jwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get Token from Header
		var tk = r.Header.Get("x-access-token")
		tk = strings.TrimSpace(tk)

		// If no Token, return 401
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

	})
}