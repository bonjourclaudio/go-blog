package post

import (
	"github.com/claudioontheweb/go-blog/auth"
	"github.com/claudioontheweb/go-blog/router"
)

var PostRoutes = router.RoutePrefix{
	"/posts",
	auth.JwtVerify,
	[]router.Route{
		router.Route{
			"GetPosts",
			"GET",
			"",
			GetPostsHandler,
		},
		router.Route{
			"GetPost",
			"GET",
			"/{postId}",
			GetPostHandler,
		},
		router.Route{
			Name:        "CreatePost",
			Method:      "POST",
			Pattern:     "",
			HandlerFunc: CreatePostHandler,
		},
		router.Route{
			Name:        "DeletePost",
			Method:      "DELETE",
			Pattern:     "/{postId}",
			HandlerFunc: DeletePostHandler,
		},
		router.Route{
			Name:        "GetPostsByUser",
			Method:      "GET",
			Pattern:     "/user/{userId}",
			HandlerFunc: GetPostsByUserHandler,
		},
	},
}


var TagRoutes = router.RoutePrefix{
	"/tags",
	auth.JwtVerify,
	[]router.Route{
		router.Route{
			"GetTags",
			"GET",
			"",
			GetTagsHandler,
		},
		router.Route{
			Name:        "GetTag",
			Method:      "GET",
			Pattern:     "/{tagId}",
			HandlerFunc: GetTagHandler,
		},
		router.Route{
			Name:        "GetTagsOfPost",
			Method:      "GET",
			Pattern:     "/post/{postId}",
			HandlerFunc: GetTagsOfPostHandler,
	},
		router.Route{
			Name:        "CreateTag",
			Method:      "POST",
			Pattern:     "/post/{postId}",
			HandlerFunc: CreateTagHandler,
		},
		router.Route{
			Name:        "DeleteTag",
			Method:      "DELETE",
			Pattern:     "/{tagId}",
			HandlerFunc: DeleteTagHandler,
		},
	},
}

var LikeRoutes = router.RoutePrefix{
	"/likes",
	auth.JwtVerify,
	[]router.Route{
		router.Route{
		Name:        "IncrementLike",
		Method:      "POST",
		Pattern:     "/post/{postId}",
		HandlerFunc: IncrementLikeHandler,
	},
		router.Route{
		Name:        "GetLikesOfPost",
		Method:      "GET",
		Pattern:     "/post/{postId}",
		HandlerFunc: GetLikesOfPostHandler,
	},
		router.Route{
		Name:        "GetLikesOfUser",
		Method:      "GET",
		Pattern:     "/user/{userId}",
		HandlerFunc: GetLikesOfUserHandler,
	},
	},
}