package main

import (
	"github.com/claudioontheweb/go-blog/auth"
	"github.com/claudioontheweb/go-blog/comment"
	"github.com/claudioontheweb/go-blog/config"
	"github.com/claudioontheweb/go-blog/db"
	"github.com/claudioontheweb/go-blog/http/middleware"
	"github.com/claudioontheweb/go-blog/post"
	customRouter "github.com/claudioontheweb/go-blog/router"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func NewRouter() *mux.Router {

	// Init Router
	r := mux.NewRouter()

	r.Use(middleware.CommonHandler)

	// Append Post routes
	customRouter.AppRoutes = append(customRouter.AppRoutes, post.PostRoutes)
	// Append Tag routes
	customRouter.AppRoutes = append(customRouter.AppRoutes, post.TagRoutes)
	// Append Likes routes
	customRouter.AppRoutes = append(customRouter.AppRoutes, post.LikeRoutes)
	// Append Comment routes
	customRouter.AppRoutes = append(customRouter.AppRoutes, comment.Routes)
	// Append Auth routes
	customRouter.AppRoutes = append(customRouter.AppRoutes, auth.Routes)
	// Append User routes
	customRouter.AppRoutes = append(customRouter.AppRoutes, auth.UserRoutes)

	for _, route := range customRouter.AppRoutes {

		// Create subroute
		routePrefix := r.PathPrefix(route.Prefix).Subrouter()

		// Use Middleware if exists
		if route.Middleware != nil {
			routePrefix.Use(route.Middleware)
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


	log.Fatal(http.ListenAndServe(":" + port, router))

}
