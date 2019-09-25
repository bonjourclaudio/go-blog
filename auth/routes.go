package auth

import (
	"github.com/claudioontheweb/go-blog/router"
)

var Routes = router.RoutePrefix{
	"/auth",
	nil,
	[]router.Route{
		{
			"CreateUser",
			"POST",
			"/register",
			CreateUserHandler,
		},
		{
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
		{
			"GetUsers",
			"GET",
			"",
			GetUsersHandler,
		},
		{
			"GetUser",
			"GET",
			"/{id}",
			GetUserHandler,
		},
		{
			"UpdateUser",
			"PUT",
			"/{id}",
			UpdateUserHandler,
		},
		{
			"DeleteUser",
			"DELETE",
			"/{id}",
			DeleteUserHandler,
		},
	},
}