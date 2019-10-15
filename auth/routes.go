package auth

import "github.com/claudioontheweb/go-blog/customRouter"

var Routes = customRouter.RoutePrefix{
	IsSecure: false,
	Prefix:   "/auth",
	SubRoutes: []customRouter.Route{
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

var UserRoutes = customRouter.RoutePrefix{
	IsSecure: true,
	Prefix:   "/user",
	SubRoutes: []customRouter.Route{
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
