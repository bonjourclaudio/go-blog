package post

import (
"github.com/claudioontheweb/go-blog/router"
)

var PostRoutes = router.RoutePrefix{
	"/posts",
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
			Name:        "IncrementLike",
			Method:      "GET",
			Pattern:     "/{postId}/like",
			HandlerFunc: IncrementLikeHandler,
		},
		router.Route{
			Name:        "GetPostsByAuthor",
			Method:      "GET",
			Pattern:     "/author/{authorId}",
			HandlerFunc: GetPostsByAuthorHandler,
		},
	},
}


var TagRoutes = router.RoutePrefix{
	"/tags",
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
