package post

import (
	"github.com/claudioontheweb/go-blog/customRouter"
)

var PostRoutes = customRouter.RoutePrefix{
	IsSecure:  true,
	Prefix: "/posts",
	SubRoutes: []customRouter.Route{
			{
				Name: "GetPosts",
				Method: "GET",
				Pattern: "",
				HandlerFunc: GetPostsHandler,
			},
			{
				Name: "GetPost",
				Method: "GET",
				Pattern: "/{postId}",
				HandlerFunc: GetPostHandler,
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

var TagRoutes = customRouter.RoutePrefix {
	IsSecure: true,
	Prefix:   "/tags",
	SubRoutes: []customRouter.Route{
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

var LikeRoutes = customRouter.RoutePrefix{
	IsSecure: true,
	Prefix:   "/likes",
	SubRoutes: []customRouter.Route{
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
