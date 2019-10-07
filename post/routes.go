package post

import (
	"github.com/claudioontheweb/go-blog/auth"
	"github.com/claudioontheweb/go-blog/router"
)

var PostRoutes = router.RoutePrefix{
	"/posts",
	auth.JwtVerify,
	[]router.Route{
		{
			"GetPosts",
			"GET",
			"",
			GetPostsHandler,
		},
		{
			"GetPost",
			"GET",
			"/{postId}",
			GetPostHandler,
		},
		{
			Name:        "CreatePost",
			Method:      "POST",
			Pattern:     "",
			HandlerFunc: CreatePostHandler,
		},
		{
			Name:        "DeletePost",
			Method:      "DELETE",
			Pattern:     "/{postId}",
			HandlerFunc: DeletePostHandler,
		},
		{
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
		{
			"GetTags",
			"GET",
			"",
			GetTagsHandler,
		},
		{
			Name:        "GetTag",
			Method:      "GET",
			Pattern:     "/{tagId}",
			HandlerFunc: GetTagHandler,
		},
		{
			Name:        "GetTagsOfPost",
			Method:      "GET",
			Pattern:     "/post/{postId}",
			HandlerFunc: GetTagsOfPostHandler,
		},
		{
			Name:        "CreateTag",
			Method:      "POST",
			Pattern:     "/post/{postId}",
			HandlerFunc: CreateTagHandler,
		},
		{
			Name:        "DeleteTag",
			Method:      "DELETE",
			Pattern:     "/post/{postId}/tag/{tagId}",
			HandlerFunc: DeleteTagHandler,
		},
	},
}

var LikeRoutes = router.RoutePrefix{
	"/likes",
	auth.JwtVerify,
	[]router.Route{
		{
			Name:        "IncrementLike",
			Method:      "POST",
			Pattern:     "/post/{postId}",
			HandlerFunc: IncrementLikeHandler,
		},
		{
			Name:        "GetLikesOfPost",
			Method:      "GET",
			Pattern:     "/post/{postId}",
			HandlerFunc: GetLikesOfPostHandler,
		},
		{
			Name:        "GetLikesOfUser",
			Method:      "GET",
			Pattern:     "/user/{userId}",
			HandlerFunc: GetLikesOfUserHandler,
		},
		{
			Name:		"DeleteLikeHandler",
			Method:		"DELETE",
			Pattern:	"/post/{postId}/like/{likeId}",
			HandlerFunc: DeleteLikeHandler,
		},
	},
}