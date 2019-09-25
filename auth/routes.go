package auth

import (
	"github.com/claudioontheweb/go-blog/router"
)

var Routes = router.RoutePrefix{
	"/auth",
	nil,
	[]router.Route{
		router.Route{
			"CreateUser",
			"POST",
			"/register",
			CreateUserHandler,
		},
		router.Route{
			"GetAuthor",
			"POST",
			"/login",
			LoginHandler,
		},
	},
}

var UserRoutes = router.RoutePrefix{
	"/user",
	JwtVerify,
	[]router.Route{
		router.Route{
			"GetUsers",
			"GET",
			"",
			GetUsersHandler,
		},
		router.Route{
			"GetUser",
			"GET",
			"/{id}",
			GetUserHandler,
		},
		router.Route{
			"UpdateUser",
			"PUT",
			"/{id}",
			UpdateUserHandler,
		},
		router.Route{
			"DeleteUser",
			"DELETE",
			"/{id}",
			DeleteUserHandler,
		},
	},
}