package comment

import (
	"github.com/claudioontheweb/go-blog/customRouter"
)

var Routes = customRouter.RoutePrefix{
	IsSecure: true,
	Prefix:   "/comments",
	SubRoutes: []customRouter.Route{
		{
			"GetComments",
			"GET",
			"",
			GetCommentsHandler,
		},
		{
			"GetComment",
			"GET",
			"/{commentId}",
			GetCommentHandler,
		},
		{
			"CreateComment",
			"POST",
			"",
			CreateCommentHandler,
		},
		{
			"DeleteComment",
			"DELETE",
			"/{commentId}",
			DeleteCommentHandler,
		},
		{
			"GetCommentsOfPost",
			"GET",
			"/post/{postId}",
			GetCommentsOfPost,
		},
		{
			"GetCommentsOfUser",
			"GET",
			"/user/{userId}",
			GetCommentsOfUser,
		},
	},
}
